// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React, { Component } from 'react'
import * as Yup from 'yup'
import { defineMessages } from 'react-intl'
import bind from 'autobind-decorator'

import Form from '../../../components/form'
import DeviceTemplateFormatSelect from '../../containers/device-template-format-select'
import FileInput from '../../../components/file-input'
import Checkbox from '../../../components/checkbox'
import SubmitBar from '../../../components/submit-bar'
import SubmitButton from '../../../components/submit-button'
import sharedMessages from '../../../lib/shared-messages'
import Message from '../../../lib/components/message'
import PropTypes from '../../../lib/prop-types'
import { dict as componentDict } from '../../../constants/components'

import style from './device-import-form.styl'

const m = defineMessages({
  fileImport: 'File Import',
  file: 'File',
  formatInfo: 'Format Information',
  createDevices: 'Create Devices',
  selectAFile: 'Please select a template file',
  fileInfoPlaceholder: 'Please select a template format',
  claimAuthCode: 'Set claim authentication code',
  targetedComponents: 'Targeted Components',
})

const validationSchema = Yup.object({
  format_id: Yup.string().required(sharedMessages.validateRequired),
  data: Yup.string().required(m.selectAFile),
  set_claim_auth_code: Yup.boolean(),
  components: Yup.object({
    is: Yup.boolean().required(),
    as: Yup.boolean(),
    js: Yup.boolean(),
    ns: Yup.boolean(),
  }).required(sharedMessages.validateRequired),
})

export default class DeviceBulkCreateForm extends Component {
  static propTypes = {
    components: PropTypes.components.isRequired,
    initialValues: PropTypes.shape({
      format_id: PropTypes.string,
      data: PropTypes.string,
      set_claim_auth_code: PropTypes.bool,
      components: PropTypes.shape({
        is: PropTypes.bool,
        ns: PropTypes.bool,
        js: PropTypes.bool,
        as: PropTypes.bool,
      }),
    }).isRequired,
    onSubmit: PropTypes.func.isRequired,
  }

  constructor(props) {
    super(props)

    const { initialValues } = props

    this.state = {
      allowedFileExtensions: undefined,
      formatDescription: undefined,
      formatSelected: false,
      jsSelected: Boolean(initialValues.components.js),
    }
    this.formRef = React.createRef()
  }

  @bind
  handleSelectChange(value) {
    const newState = { formatSelected: true }
    if (value && value.fileExtensions && value.fileExtensions instanceof Array) {
      newState.allowedFileExtensions = value.fileExtensions.join(',')
    }
    if (value && value.description) {
      newState.formatDescription = value.description
    }
    this.setState(newState)
  }

  @bind
  handleComponentChange(value) {
    const { jsSelected } = this.state
    const { state } = this.formRef.current
    const { js } = value

    if (js !== jsSelected) {
      this.setState({ jsSelected: js }, () => {
        if (state.values.set_claim_auth_code) {
          const { setFieldValue } = this.formRef.current

          // `claim_authentication_code` is stored in JS, so if JS option is not selected
          // we dont want to include it in the paylaod
          setFieldValue('set_claim_auth_code', false)
        }
      })
    }
  }

  render() {
    const { initialValues, onSubmit, components } = this.props
    const { allowedFileExtensions, formatSelected, formatDescription, jsSelected } = this.state
    return (
      <Form
        formikRef={this.formRef}
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        submitEnabledWhenInvalid
        initialValues={initialValues}
      >
        <Message component="h4" content={m.fileImport} />
        <DeviceTemplateFormatSelect onChange={this.handleSelectChange} name="format_id" required />
        <Form.InfoField disabled={!formatSelected} title={m.formatInfo}>
          {formatDescription ? formatDescription : <Message content={m.fileInfoPlaceholder} />}
        </Form.InfoField>
        <hr className={style.hRule} />
        <Form.Field
          disabled={!formatSelected}
          title={m.file}
          accept={allowedFileExtensions}
          component={FileInput}
          name="data"
          required
        />
        <Form.Field
          onChange={this.handleComponentChange}
          component={Checkbox.Group}
          name="components"
          title={m.targetedComponents}
          required
          horizontal={false}
          disabled={!formatSelected}
        >
          {components.map(component => (
            <Checkbox
              disabled={component === 'is'}
              key={component}
              name={component}
              label={componentDict[component]}
            />
          ))}
        </Form.Field>
        <Form.Field
          disabled={!formatSelected || !jsSelected}
          title={m.claimAuthCode}
          component={Checkbox}
          name="set_claim_auth_code"
        />
        <SubmitBar>
          <Form.Submit component={SubmitButton} message={m.createDevices} />
        </SubmitBar>
      </Form>
    )
  }
}
