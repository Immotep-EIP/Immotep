import { Form, Button } from 'antd'

import useNavigation from '@/hooks/useNavigation/useNavigation'
import '@/views/Login/Login.css'
import EmailInput from '@/components/Inputs/EmailInput'
import PasswordInput from '@/components/Inputs/PasswordInput'
import LoginTitle from '@/components/Login/LoginTitle'
import LoginOptions from '@/components/Login/LoginOptions'
import AuthFooter from '@/components/Footers/AuthFooter'

const { Item } = Form

type FieldType = {
  email: string
  password: string
  remember: boolean
}

const LoginPage = () => {
  const { goToHome, goToSignup } = useNavigation()
  const onFinish = (values: FieldType) => {
    // eslint-disable-next-line no-console
    console.log(values)
  }

  return (
    <div className="loginContainer">
      <LoginTitle />
      <Form
        name="login"
        layout="vertical"
        onFinish={onFinish}
        initialValues={{ remember: false }}
      >
        <EmailInput />
        <PasswordInput />
        <LoginOptions goToHome={goToHome} />
        <Item>
          <Button type="primary" htmlType="submit">
            Sign in
          </Button>
        </Item>
        <AuthFooter
          goTo={goToSignup}
          text="Don't have an account ?"
          buttonText="Sign up"
        />
      </Form>
    </div>
  )
}

export default LoginPage