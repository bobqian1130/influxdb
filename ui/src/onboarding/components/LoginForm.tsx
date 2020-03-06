// Libraries
import React, {FC, useState, ChangeEvent} from 'react'
import {
  Button,
  ButtonShape,
  ButtonType,
  Columns,
  ComponentColor,
  ComponentSize,
  ComponentStatus,
  Form,
  Grid,
  Input,
  InputType,
  VisibilityInput,
} from '@influxdata/clockface'

// Types
import {FormFieldValidation} from 'src/types'

interface Props {
  buttonStatus: ComponentStatus
  emailValidation: FormFieldValidation
  email: string
  passwordValidation: FormFieldValidation
  password: string
  handleInputChange: (e: ChangeEvent<HTMLInputElement>) => void
  handleForgotPasswordClick: () => void
}

export const LoginForm: FC<Props> = ({
  buttonStatus,
  emailValidation,
  email,
  passwordValidation,
  password,
  handleInputChange,
  handleForgotPasswordClick,
}) => {
  const [visible, toggleIcon] = useState(false)
  return (
    <>
      <Grid>
        <Grid.Row className="sign-up--form-padded-row">
          <Grid.Column widthXS={Columns.Twelve}>
            <Form.Element
              label="Work Email Address"
              required={true}
              errorMessage={emailValidation.errorMessage}
            >
              <Input
                name="email"
                value={email}
                type={InputType.Email}
                size={ComponentSize.Large}
                status={
                  emailValidation.isValid
                    ? ComponentStatus.Error
                    : ComponentStatus.Default
                }
                onChange={handleInputChange}
              />
            </Form.Element>
          </Grid.Column>
        </Grid.Row>
        <Grid.Row>
          <Grid.Column widthXS={Columns.Twelve}>
            <Form.Element
              label="Password"
              required={true}
              errorMessage={passwordValidation.errorMessage}
            >
              <VisibilityInput
                name="password"
                value={password}
                size={ComponentSize.Large}
                onChange={handleInputChange}
                visible={visible}
                status={
                  passwordValidation.isValid
                    ? ComponentStatus.Error
                    : ComponentStatus.Default
                }
                onToggleClick={() => toggleIcon(!visible)}
              />
            </Form.Element>
          </Grid.Column>
        </Grid.Row>
      </Grid>
      <a onClick={handleForgotPasswordClick} className="login--forgot-password">
        Forgot Password?
      </a>
      <Button
        className="create-account--button"
        text="Login"
        color={ComponentColor.Primary}
        size={ComponentSize.Large}
        type={ButtonType.Submit}
        status={buttonStatus}
        shape={ButtonShape.StretchToFit}
      />
    </>
  )
}
