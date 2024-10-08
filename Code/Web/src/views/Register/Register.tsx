import { Form, Button, Checkbox } from 'antd'

import useNavigation from '@/hooks/useNavigation/useNavigation'
import RegisterInputs from '@/components/Register/RegisterInputs'
import RegisterTitle from '@/components/Register/RegisterTitle'
import AuthFooter from '@/components/Footers/AuthFooter'

import '@/views/Register/Register.css'

const { Item } = Form

type RegistrationInput = {
  lastname: string
  firstname: string
  email: string
  password: string
  confirmPassword: string
  agree: boolean
}

const Register = () => {
  const { goToLogin } = useNavigation()

  const onFinish = (values: RegistrationInput) => {
    // eslint-disable-next-line no-console
    console.log(values)
  }

  return (
    <div className="registerContainer">
      <RegisterTitle />
      <Form
        name="register"
        layout="vertical"
        onFinish={onFinish}
        initialValues={{ agree: false }}
      >
        <RegisterInputs />
        <Item name="agree" valuePropName="checked">
          <Checkbox>I agree to all Term, Privacy Policy and Fees</Checkbox>
        </Item>
        <Item>
          <Button type="primary" htmlType="submit">
            Sign in
          </Button>
        </Item>
        <AuthFooter
          goTo={goToLogin}
          text="Already have an account ?"
          buttonText="Log in"
        />
      </Form>
    </div>
  )
}

export default Register
