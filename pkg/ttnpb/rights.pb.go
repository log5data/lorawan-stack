// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/rights.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Right is the enum that defines all the different rights to do something in
// the network.
type Right int32

const (
	// RIGHT_INVALID is the invalid right and should not be used.
	// It can denote a parsing error.
	RIGHT_INVALID Right = 0
	// RIGHT_USER_PROFILE_READ is the right to read the user's account.
	RIGHT_USER_PROFILE_READ Right = 1
	// RIGHT_USER_PROFILE_WRITE is the right to edit the user's account.
	RIGHT_USER_PROFILE_WRITE Right = 2
	// RIGHT_USER_DELETE is the right to delete the user's account.
	RIGHT_USER_DELETE Right = 3
	// RIGHT_USER_AUTHORIZEDCLIENTS is the right to view and manage the authorized
	// clients of the user: list, authorize and de-authorize clients.
	RIGHT_USER_AUTHORIZEDCLIENTS Right = 4
	// RIGHT_USER_APPLICATIONS_LIST is the right to list all of the applications
	// the user is a collaborator of.
	RIGHT_USER_APPLICATIONS_LIST Right = 5
	// RIGHT_USER_APPLICATIONS_CREATE is the right to create an application in
	// the users account.
	RIGHT_USER_APPLICATIONS_CREATE Right = 6
	// RIGHT_USER_APPLICATIONS is the right to manage applications in the users account.
	RIGHT_USER_APPLICATIONS Right = 8
	// RIGHT_USER_GATEWAYS_LIST is the right to list gateways the user is collaborator of.
	RIGHT_USER_GATEWAYS_LIST Right = 9
	// RIGHT_USER_GATEWAYS_CREATE is the right to register a gateway in the users account.
	RIGHT_USER_GATEWAYS_CREATE Right = 10
	// RIGHT_USER_APPLICATIONS is the right to manage applications in the users account.
	RIGHT_USER_GATEWAYS Right = 11
	// RIGHT_USER_CLIENTS_LIST is the right to list the OAuth clients of the user.
	RIGHT_USER_CLIENTS_LIST Right = 12
	// RIGHT_USER_CLIENTS_CREATE is the right to create a new OAuth client in the users account.
	RIGHT_USER_CLIENTS_CREATE Right = 13
	// RIGHT_USER_CLIENTS_MANAGE is the right to update and delete OAuth clients of the user.
	RIGHT_USER_CLIENTS_MANAGE Right = 14
	// RIGHT_APPLICATION_INFO is the right to see basic information about the application.
	// This does not include API keys.
	RIGHT_APPLICATION_INFO Right = 31
	// RIGHT_APPLICATION_SETTINGS_BASIC is the right to edit basic settings of the application.
	RIGHT_APPLICATION_SETTINGS_BASIC Right = 32
	// RIGHT_APPLICATION_SETTINGS_KEYS is the right to view and edit application
	// API keys.
	RIGHT_APPLICATION_SETTINGS_KEYS Right = 33
	// RIGHT_APPLICATION_SETTINGS_COLLABORATORS is the right to manage the
	// collaborators of the application.
	RIGHT_APPLICATION_SETTINGS_COLLABORATORS Right = 34
	// RIGHT_APPLICATION_DELETE is the right to delete the application.
	RIGHT_APPLICATION_DELETE Right = 35
	// RIGHT_APPLICATION_DEVICES_READ is the right to see the devices of an application.
	RIGHT_APPLICATION_DEVICES_READ Right = 36
	// RIGHT_APPLICATION_DEVICES_WRITE is the right to register devices to an application.
	RIGHT_APPLICATION_DEVICES_WRITE Right = 37
	// RIGHT_APPLICATION_TRAFFIC_READ is the right to read traffic of the
	// application (uplink and downlink).
	RIGHT_APPLICATION_TRAFFIC_READ Right = 38
	// RIGHT_APPLICATION_TRAFFIC_WRITE is the right to write traffic of the
	// application (uplink and downlink).
	RIGHT_APPLICATION_TRAFFIC_WRITE Right = 39
	// RIGHT_GATEWAY_INFO is the right to see basic information about the gateway.
	// This does not include API keys.
	RIGHT_GATEWAY_INFO Right = 51
	// RIGHT_GATEWAY_SETTINGS_BASIC is the right to manage basic settings of the gateway.
	RIGHT_GATEWAY_SETTINGS_BASIC Right = 52
	// RIGHT_GATEWAY_SETTINGS_KEYS is the right to view and edit API keys of the gateway.
	RIGHT_GATEWAY_SETTINGS_KEYS Right = 53
	// RIGHT_GATEWAY_SETTINGS_COLLABORATORS is the right to manage collaborators
	// of the gateway.
	RIGHT_GATEWAY_SETTINGS_COLLABORATORS Right = 54
	// RIGHT_GATEWAY_DELETE is the right to delete the gateway.
	RIGHT_GATEWAY_DELETE Right = 55
	// RIGHT_GATEWAY_TRAFFIC is the right to view traffic of the gateway.
	RIGHT_GATEWAY_TRAFFIC Right = 56
	// RIGHT_GATEWAY_STATUS is the right to view the status of the gateway.
	RIGHT_GATEWAY_STATUS Right = 57
	// RIGHT_GATEWAY_LOCATION is the right to view the location of the gateway.
	RIGHT_GATEWAY_LOCATION Right = 58
)

var Right_name = map[int32]string{
	0:  "RIGHT_INVALID",
	1:  "RIGHT_USER_PROFILE_READ",
	2:  "RIGHT_USER_PROFILE_WRITE",
	3:  "RIGHT_USER_DELETE",
	4:  "RIGHT_USER_AUTHORIZEDCLIENTS",
	5:  "RIGHT_USER_APPLICATIONS_LIST",
	6:  "RIGHT_USER_APPLICATIONS_CREATE",
	8:  "RIGHT_USER_APPLICATIONS",
	9:  "RIGHT_USER_GATEWAYS_LIST",
	10: "RIGHT_USER_GATEWAYS_CREATE",
	11: "RIGHT_USER_GATEWAYS",
	12: "RIGHT_USER_CLIENTS_LIST",
	13: "RIGHT_USER_CLIENTS_CREATE",
	14: "RIGHT_USER_CLIENTS_MANAGE",
	31: "RIGHT_APPLICATION_INFO",
	32: "RIGHT_APPLICATION_SETTINGS_BASIC",
	33: "RIGHT_APPLICATION_SETTINGS_KEYS",
	34: "RIGHT_APPLICATION_SETTINGS_COLLABORATORS",
	35: "RIGHT_APPLICATION_DELETE",
	36: "RIGHT_APPLICATION_DEVICES_READ",
	37: "RIGHT_APPLICATION_DEVICES_WRITE",
	38: "RIGHT_APPLICATION_TRAFFIC_READ",
	39: "RIGHT_APPLICATION_TRAFFIC_WRITE",
	51: "RIGHT_GATEWAY_INFO",
	52: "RIGHT_GATEWAY_SETTINGS_BASIC",
	53: "RIGHT_GATEWAY_SETTINGS_KEYS",
	54: "RIGHT_GATEWAY_SETTINGS_COLLABORATORS",
	55: "RIGHT_GATEWAY_DELETE",
	56: "RIGHT_GATEWAY_TRAFFIC",
	57: "RIGHT_GATEWAY_STATUS",
	58: "RIGHT_GATEWAY_LOCATION",
}
var Right_value = map[string]int32{
	"RIGHT_INVALID":                            0,
	"RIGHT_USER_PROFILE_READ":                  1,
	"RIGHT_USER_PROFILE_WRITE":                 2,
	"RIGHT_USER_DELETE":                        3,
	"RIGHT_USER_AUTHORIZEDCLIENTS":             4,
	"RIGHT_USER_APPLICATIONS_LIST":             5,
	"RIGHT_USER_APPLICATIONS_CREATE":           6,
	"RIGHT_USER_APPLICATIONS":                  8,
	"RIGHT_USER_GATEWAYS_LIST":                 9,
	"RIGHT_USER_GATEWAYS_CREATE":               10,
	"RIGHT_USER_GATEWAYS":                      11,
	"RIGHT_USER_CLIENTS_LIST":                  12,
	"RIGHT_USER_CLIENTS_CREATE":                13,
	"RIGHT_USER_CLIENTS_MANAGE":                14,
	"RIGHT_APPLICATION_INFO":                   31,
	"RIGHT_APPLICATION_SETTINGS_BASIC":         32,
	"RIGHT_APPLICATION_SETTINGS_KEYS":          33,
	"RIGHT_APPLICATION_SETTINGS_COLLABORATORS": 34,
	"RIGHT_APPLICATION_DELETE":                 35,
	"RIGHT_APPLICATION_DEVICES_READ":           36,
	"RIGHT_APPLICATION_DEVICES_WRITE":          37,
	"RIGHT_APPLICATION_TRAFFIC_READ":           38,
	"RIGHT_APPLICATION_TRAFFIC_WRITE":          39,
	"RIGHT_GATEWAY_INFO":                       51,
	"RIGHT_GATEWAY_SETTINGS_BASIC":             52,
	"RIGHT_GATEWAY_SETTINGS_KEYS":              53,
	"RIGHT_GATEWAY_SETTINGS_COLLABORATORS":     54,
	"RIGHT_GATEWAY_DELETE":                     55,
	"RIGHT_GATEWAY_TRAFFIC":                    56,
	"RIGHT_GATEWAY_STATUS":                     57,
	"RIGHT_GATEWAY_LOCATION":                   58,
}

func (Right) EnumDescriptor() ([]byte, []int) { return fileDescriptorRights, []int{0} }

type APIKey struct {
	// key is the API key itself which is generated by the Identity Server.
	// It is unique, immutable and read-only.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// name is the API key name. The name is an alias and it is non-unique within
	// the application.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// rights are the rights this API key bears.
	Rights []Right `protobuf:"varint,3,rep,packed,name=rights,enum=ttn.v3.Right" json:"rights,omitempty"`
}

func (m *APIKey) Reset()                    { *m = APIKey{} }
func (m *APIKey) String() string            { return proto.CompactTextString(m) }
func (*APIKey) ProtoMessage()               {}
func (*APIKey) Descriptor() ([]byte, []int) { return fileDescriptorRights, []int{0} }

func (m *APIKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *APIKey) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *APIKey) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

type APIKeyMask struct {
	Name   bool `protobuf:"varint,1,opt,name=name,proto3" json:"name,omitempty"`
	Rights bool `protobuf:"varint,2,opt,name=rights,proto3" json:"rights,omitempty"`
}

func (m *APIKeyMask) Reset()                    { *m = APIKeyMask{} }
func (m *APIKeyMask) String() string            { return proto.CompactTextString(m) }
func (*APIKeyMask) ProtoMessage()               {}
func (*APIKeyMask) Descriptor() ([]byte, []int) { return fileDescriptorRights, []int{1} }

func (m *APIKeyMask) GetName() bool {
	if m != nil {
		return m.Name
	}
	return false
}

func (m *APIKeyMask) GetRights() bool {
	if m != nil {
		return m.Rights
	}
	return false
}

func init() {
	proto.RegisterType((*APIKey)(nil), "ttn.v3.APIKey")
	golang_proto.RegisterType((*APIKey)(nil), "ttn.v3.APIKey")
	proto.RegisterType((*APIKeyMask)(nil), "ttn.v3.APIKeyMask")
	golang_proto.RegisterType((*APIKeyMask)(nil), "ttn.v3.APIKeyMask")
	proto.RegisterEnum("ttn.v3.Right", Right_name, Right_value)
	golang_proto.RegisterEnum("ttn.v3.Right", Right_name, Right_value)
}
func (x Right) String() string {
	s, ok := Right_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *APIKey) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*APIKey)
	if !ok {
		that2, ok := that.(APIKey)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *APIKey")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *APIKey but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *APIKey but is not nil && this == nil")
	}
	if this.Key != that1.Key {
		return fmt.Errorf("Key this(%v) Not Equal that(%v)", this.Key, that1.Key)
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if len(this.Rights) != len(that1.Rights) {
		return fmt.Errorf("Rights this(%v) Not Equal that(%v)", len(this.Rights), len(that1.Rights))
	}
	for i := range this.Rights {
		if this.Rights[i] != that1.Rights[i] {
			return fmt.Errorf("Rights this[%v](%v) Not Equal that[%v](%v)", i, this.Rights[i], i, that1.Rights[i])
		}
	}
	return nil
}
func (this *APIKey) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*APIKey)
	if !ok {
		that2, ok := that.(APIKey)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if len(this.Rights) != len(that1.Rights) {
		return false
	}
	for i := range this.Rights {
		if this.Rights[i] != that1.Rights[i] {
			return false
		}
	}
	return true
}
func (this *APIKeyMask) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*APIKeyMask)
	if !ok {
		that2, ok := that.(APIKeyMask)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *APIKeyMask")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *APIKeyMask but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *APIKeyMask but is not nil && this == nil")
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if this.Rights != that1.Rights {
		return fmt.Errorf("Rights this(%v) Not Equal that(%v)", this.Rights, that1.Rights)
	}
	return nil
}
func (this *APIKeyMask) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*APIKeyMask)
	if !ok {
		that2, ok := that.(APIKeyMask)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Rights != that1.Rights {
		return false
	}
	return true
}
func (m *APIKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *APIKey) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRights(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRights(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Rights) > 0 {
		dAtA2 := make([]byte, len(m.Rights)*10)
		var j1 int
		for _, num := range m.Rights {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRights(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	return i, nil
}

func (m *APIKeyMask) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *APIKeyMask) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name {
		dAtA[i] = 0x8
		i++
		if m.Name {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Rights {
		dAtA[i] = 0x10
		i++
		if m.Rights {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintRights(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedAPIKey(r randyRights, easy bool) *APIKey {
	this := &APIKey{}
	this.Key = randStringRights(r)
	this.Name = randStringRights(r)
	v1 := r.Intn(10)
	this.Rights = make([]Right, v1)
	for i := 0; i < v1; i++ {
		this.Rights[i] = Right([]int32{0, 1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12, 13, 14, 31, 32, 33, 34, 35, 36, 37, 38, 39, 51, 52, 53, 54, 55, 56, 57, 58}[r.Intn(31)])
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedAPIKeyMask(r randyRights, easy bool) *APIKeyMask {
	this := &APIKeyMask{}
	this.Name = bool(r.Intn(2) == 0)
	this.Rights = bool(r.Intn(2) == 0)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyRights interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneRights(r randyRights) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringRights(r randyRights) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneRights(r)
	}
	return string(tmps)
}
func randUnrecognizedRights(r randyRights, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldRights(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldRights(dAtA []byte, r randyRights, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateRights(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateRights(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateRights(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateRights(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateRights(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateRights(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateRights(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *APIKey) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRights(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRights(uint64(l))
	}
	if len(m.Rights) > 0 {
		l = 0
		for _, e := range m.Rights {
			l += sovRights(uint64(e))
		}
		n += 1 + sovRights(uint64(l)) + l
	}
	return n
}

func (m *APIKeyMask) Size() (n int) {
	var l int
	_ = l
	if m.Name {
		n += 2
	}
	if m.Rights {
		n += 2
	}
	return n
}

func sovRights(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRights(x uint64) (n int) {
	return sovRights((x << 1) ^ uint64((int64(x) >> 63)))
}
func (m *APIKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRights
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: APIKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: APIKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRights
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRights
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRights
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRights
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v Right
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRights
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (Right(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Rights = append(m.Rights, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRights
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthRights
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v Right
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRights
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (Right(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Rights = append(m.Rights, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Rights", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRights(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRights
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *APIKeyMask) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRights
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: APIKeyMask: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: APIKeyMask: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRights
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Name = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rights", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRights
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Rights = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRights(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRights
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRights(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRights
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRights
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRights
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRights
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRights
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRights(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRights = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRights   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/rights.proto", fileDescriptorRights)
}
func init() {
	golang_proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/rights.proto", fileDescriptorRights)
}

var fileDescriptorRights = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x3f, 0x50, 0x1a, 0x4f,
	0x14, 0xc7, 0x6f, 0x45, 0xf9, 0xe9, 0xfe, 0xa2, 0xb3, 0x6e, 0xa2, 0x22, 0x9a, 0x27, 0x41, 0x4d,
	0x48, 0x26, 0x81, 0x4c, 0xcc, 0x1f, 0x93, 0xee, 0x84, 0x05, 0x77, 0x3c, 0x39, 0xe6, 0x6e, 0xd1,
	0xd1, 0x86, 0x91, 0x0c, 0x01, 0x86, 0x11, 0x18, 0x3d, 0x93, 0xb1, 0xb3, 0xb4, 0x4c, 0x99, 0x2e,
	0x99, 0x49, 0x63, 0x69, 0x69, 0x69, 0x69, 0x69, 0x69, 0x95, 0x91, 0xbb, 0xc6, 0xd2, 0xd2, 0x32,
	0xc3, 0xdd, 0x21, 0x7f, 0x04, 0x53, 0xdd, 0xee, 0xfb, 0x7e, 0xf7, 0xb3, 0x6f, 0xdf, 0x7b, 0x73,
	0xf8, 0x75, 0xbe, 0x68, 0x14, 0xf6, 0xb2, 0xe1, 0xcf, 0x95, 0xed, 0x88, 0x28, 0xe4, 0x44, 0xa1,
	0x58, 0xce, 0xef, 0x26, 0x73, 0xc6, 0xb7, 0xca, 0x4e, 0x29, 0x62, 0x18, 0xe5, 0xc8, 0x56, 0xb5,
	0x18, 0xd9, 0x29, 0xe6, 0x0b, 0xc6, 0x6e, 0xb8, 0xba, 0x53, 0x31, 0x2a, 0xd4, 0x6b, 0x18, 0xe5,
	0xf0, 0xd7, 0x05, 0xff, 0xab, 0x96, 0x93, 0xf9, 0x4a, 0xbe, 0x12, 0xb1, 0xe5, 0xec, 0xde, 0x17,
	0x7b, 0x67, 0x6f, 0xec, 0x95, 0x73, 0x2c, 0x98, 0xc6, 0x5e, 0x39, 0xc5, 0x57, 0x72, 0xfb, 0x94,
	0x60, 0x4f, 0x29, 0xb7, 0xef, 0x43, 0x01, 0x14, 0x1a, 0xd2, 0xea, 0x4b, 0x4a, 0x71, 0x7f, 0x79,
	0x6b, 0x3b, 0xe7, 0xeb, 0xb3, 0x43, 0xf6, 0x9a, 0xce, 0x63, 0xaf, 0x73, 0xad, 0xcf, 0x13, 0xf0,
	0x84, 0x46, 0xde, 0x0c, 0x87, 0x9d, 0x7b, 0xc3, 0x5a, 0x3d, 0xaa, 0xb9, 0x62, 0x70, 0x11, 0x63,
	0x07, 0xbb, 0xba, 0xb5, 0x5b, 0xba, 0x05, 0xd5, 0xd9, 0x83, 0x2e, 0x68, 0xfc, 0x16, 0xd4, 0x67,
	0x47, 0xdd, 0xdd, 0x8b, 0x3f, 0xff, 0xe1, 0x01, 0x9b, 0x45, 0x47, 0xf1, 0xb0, 0xc6, 0x13, 0xcb,
	0x22, 0xc3, 0x93, 0x6b, 0xb2, 0xc2, 0x63, 0x44, 0xa2, 0x53, 0x78, 0xc2, 0x09, 0xa5, 0x75, 0xa6,
	0x65, 0x52, 0x9a, 0x1a, 0xe7, 0x0a, 0xcb, 0x68, 0x4c, 0x8e, 0x11, 0x44, 0xa7, 0xb1, 0xaf, 0x8b,
	0xb8, 0xae, 0x71, 0xc1, 0x48, 0x1f, 0x1d, 0xc3, 0xa3, 0x2d, 0x6a, 0x8c, 0x29, 0x4c, 0x30, 0xe2,
	0xa1, 0x01, 0x3c, 0xdd, 0x12, 0x96, 0xd3, 0x62, 0x59, 0xd5, 0xf8, 0x26, 0x8b, 0x45, 0x15, 0xce,
	0x92, 0x42, 0x27, 0xfd, 0x9d, 0x8e, 0x54, 0x4a, 0xe1, 0x51, 0x59, 0x70, 0x35, 0xa9, 0x67, 0x14,
	0xae, 0x0b, 0x32, 0x40, 0x83, 0x18, 0x7a, 0x39, 0xa2, 0x1a, 0x93, 0x05, 0x23, 0xde, 0x8e, 0xcc,
	0x5b, 0x3d, 0x64, 0xb0, 0x23, 0xf3, 0x84, 0x2c, 0xd8, 0xba, 0xbc, 0xe1, 0xe2, 0x87, 0x28, 0x60,
	0x7f, 0x37, 0xd5, 0x45, 0x63, 0x3a, 0x81, 0x1f, 0x76, 0xd1, 0xc9, 0xff, 0x1d, 0x77, 0xba, 0x2f,
	0x72, 0xa8, 0x0f, 0xe8, 0x63, 0x3c, 0xd9, 0x45, 0x74, 0xa1, 0xc3, 0x3d, 0xe4, 0x55, 0x39, 0x29,
	0x27, 0x18, 0x19, 0xa1, 0x7e, 0x3c, 0xee, 0xc8, 0x2d, 0x2f, 0xc9, 0xf0, 0x64, 0x5c, 0x25, 0x33,
	0x74, 0x0e, 0x07, 0xee, 0x6a, 0x3a, 0x13, 0x82, 0x27, 0x13, 0x7a, 0x66, 0x49, 0xd6, 0x79, 0x94,
	0x04, 0xe8, 0x2c, 0x9e, 0xb9, 0xc7, 0xb5, 0xc2, 0x36, 0x74, 0xf2, 0x84, 0xbe, 0xc4, 0xa1, 0x7b,
	0x4c, 0x51, 0x55, 0x51, 0xe4, 0x25, 0x55, 0x93, 0x85, 0xaa, 0xe9, 0x24, 0xd8, 0x2c, 0x63, 0xab,
	0xdb, 0xed, 0xf4, 0x6c, 0xb3, 0x4b, 0xed, 0xea, 0x1a, 0x8f, 0x32, 0xdd, 0x19, 0xa1, 0xb9, 0xee,
	0x49, 0x35, 0x3c, 0xce, 0x24, 0xcd, 0x77, 0x07, 0x09, 0x4d, 0x8e, 0xc7, 0x79, 0xd4, 0x01, 0x3d,
	0xed, 0x0e, 0x6a, 0x78, 0x1c, 0xd0, 0x33, 0x3a, 0x8e, 0xa9, 0x63, 0x72, 0x7b, 0xe6, 0x14, 0x70,
	0xa1, 0x39, 0x71, 0x8d, 0x78, 0x47, 0xf1, 0xde, 0xd2, 0x19, 0x3c, 0xd5, 0xc3, 0x61, 0x17, 0xee,
	0x1d, 0x0d, 0xe1, 0xb9, 0x1e, 0x86, 0xf6, 0xa2, 0xbd, 0xa7, 0x3e, 0xfc, 0xa8, 0xdd, 0xe9, 0x16,
	0xec, 0x03, 0x9d, 0xc4, 0x63, 0xed, 0x8a, 0x9b, 0x3f, 0x59, 0xbc, 0x7b, 0x48, 0x17, 0xb2, 0x48,
	0xeb, 0xe4, 0x63, 0x73, 0x30, 0x1a, 0x8a, 0xa2, 0x3a, 0xaf, 0x27, 0x9f, 0xfc, 0xfd, 0x87, 0xbf,
	0x41, 0x5a, 0xfa, 0x89, 0xce, 0x6a, 0x80, 0xce, 0x6b, 0x80, 0x2e, 0x6a, 0x80, 0x2e, 0x6b, 0x80,
	0xae, 0x6a, 0x20, 0x5d, 0xd7, 0x40, 0xba, 0xa9, 0x01, 0x3a, 0x30, 0x41, 0x3a, 0x34, 0x41, 0x3a,
	0x32, 0x01, 0x1d, 0x9b, 0x20, 0x9d, 0x98, 0x80, 0x4e, 0x4d, 0x40, 0x67, 0x26, 0xa0, 0x73, 0x13,
	0xd0, 0x85, 0x09, 0xd2, 0xa5, 0x09, 0xe8, 0xca, 0x04, 0xe9, 0xda, 0x04, 0x74, 0x63, 0x82, 0x74,
	0x60, 0x81, 0x74, 0x68, 0x01, 0xfa, 0x6e, 0x81, 0xf4, 0xc3, 0x02, 0xf4, 0xcb, 0x02, 0xe9, 0xc8,
	0x02, 0xe9, 0xd8, 0x02, 0x74, 0x62, 0x01, 0x3a, 0xb5, 0x00, 0x6d, 0x3e, 0xff, 0xd7, 0x1f, 0xb5,
	0x5a, 0xca, 0xd7, 0xbf, 0xd5, 0x6c, 0xd6, 0x6b, 0xff, 0x1a, 0x17, 0xfe, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xd3, 0xcf, 0xd8, 0x32, 0x85, 0x05, 0x00, 0x00,
}
