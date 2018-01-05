// Copyright © 2018 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package claims

import (
	"context"

	"github.com/TheThingsNetwork/ttn/pkg/auth"
	"github.com/TheThingsNetwork/ttn/pkg/auth/oauth"
	"github.com/TheThingsNetwork/ttn/pkg/errors"
	"github.com/TheThingsNetwork/ttn/pkg/identityserver/types"
	"github.com/TheThingsNetwork/ttn/pkg/rpcmetadata"
	"github.com/TheThingsNetwork/ttn/pkg/ttnpb"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

// TokenKeyInfoProvider is the interface that validates and introspects OAuth
// access tokens and API keys.
type TokenKeyInfoProvider interface {
	// TokenInfo returns the access data of an OAuth access token.
	// It returns error if token is expired.
	TokenInfo(accessToken string) (*types.AccessData, error)

	// KeyInfo returns the entityID an API key belongs to and its rights.
	// The Resource Server must check that the entityID of the API key matches
	// with the resource that is trying to being access to.
	KeyInfo(key string) (string, *ttnpb.APIKey, error)
}

// claims constructs the claims based on the authentication values in the request
// metadata. Mainly there are three scenarios:
//   - Authentication values are empty: empty claims are returned
//   - A token is provided: it is introspected (and therefore validated) through
//       the TokenInfoProvider and then claims are built based on this.
//   - A key is provided: it is introspected (and therefore validated) through
//       the KeyInfoProvider and then claims are built based on this.
func claims(ctx context.Context, tokenkey TokenKeyInfoProvider) (*auth.Claims, error) {
	md := rpcmetadata.FromIncomingContext(ctx)

	if md.AuthType == "" && md.AuthValue == "" {
		return new(auth.Claims), nil
	}

	if md.AuthType != "Bearer" {
		return nil, errors.Errorf("Expected authentication type to be Bearer but got `%s", md.AuthType)
	}

	header, payload, err := auth.DecodeTokenOrKey(md.AuthValue)
	if err != nil {
		return nil, err
	}

	var claims *auth.Claims
	switch header.Type {
	case auth.Token:
		data, err := tokenkey.TokenInfo(md.AuthValue)
		if err != nil {
			return nil, err
		}

		rights, err := oauth.ParseScope(data.Scope)
		if err != nil {
			return nil, err
		}

		claims = &auth.Claims{
			EntityID:   data.UserID,
			EntityType: auth.EntityUser,
			Source:     auth.Token,
			Rights:     rights,
		}
	case auth.Key:
		entityID, key, err := tokenkey.KeyInfo(md.AuthValue)
		if err != nil {
			return nil, err
		}

		claims = &auth.Claims{
			EntityID: entityID,
			Source:   auth.Key,
			Rights:   key.Rights,
		}

		switch payload.Type {
		case auth.ApplicationKey:
			claims.EntityType = auth.EntityApplication
		case auth.GatewayKey:
			claims.EntityType = auth.EntityGateway
		case auth.UserKey:
			claims.EntityType = auth.EntityUser
		default:
			return nil, errors.Errorf("Invalid API key type `%s`", payload.Type)
		}
	default:
		return nil, errors.New("Invalid authentication value")
	}

	return claims, nil
}

// UnaryServerInterceptor returns a new unary server interceptor that construct
// the claims based on the authentication value in the request metadata.
// Empty claims are injected if authentication is missing in the request metadata.
func UnaryServerInterceptor(tokenkey TokenKeyInfoProvider) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		c, err := claims(ctx, tokenkey)
		if err != nil {
			return nil, err
		}

		return handler(NewContext(ctx, c), req)
	}
}

// StreamServerInterceptor returns a new unary server interceptor that construct
// the claims based on the authentication value in the request metadata.
// Empty claims are injected if authentication is missing in the request metadata.
func StreamServerInterceptor(tokenkey TokenKeyInfoProvider) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		c, err := claims(stream.Context(), tokenkey)
		if err != nil {
			return err
		}

		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = NewContext(stream.Context(), c)

		return handler(srv, wrapped)
	}
}
