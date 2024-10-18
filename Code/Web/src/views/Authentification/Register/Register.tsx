import React from 'react'
import { Button, Input, Form, message, Checkbox } from 'antd'
import type { FormProps } from 'antd'
import AuthentificationPage from '@/components/AuthentificationPage/AuthentificationPage'
import useNavigation from '@/hooks/useNavigation/useNavigation'
import '@/App.css'
import { register } from '@/services/api/Authentification/AuthApi'
import { UserRegister } from '@/interfaces/User/User'
import style from './Register.module.css'

const Register: React.FC = () => {
  const { goToLogin } = useNavigation()
  const [form] = Form.useForm()

  const onFinish: FormProps<UserRegister>['onFinish'] = async values => {
    try {
      const { password, confirmPassword } = values
      if (password === confirmPassword) {
        await register(values)
        message.success('Registration success', 5)
        form.resetFields()
        setTimeout(() => {
          goToLogin()
        }, 1000)
      } else message.error('Please confirm your password', 5)
    } catch (err: any) {
      if (err.response.status === 409)
        message.error('Email already exist. Please try again.')
      console.error('Registration error:', err)
    }
  }

  const onFinishFailed: FormProps<UserRegister>['onFinishFailed'] = () => {
    message.error('An error occured, please try again')
  }

  return (
    <AuthentificationPage
      title="Create your account"
      subtitle="Join immotep for your peace of mind !"
    >
      <Form
        form={form}
        name="basic"
        initialValues={{ termAgree: false }}
        onFinish={onFinish}
        onFinishFailed={onFinishFailed}
        autoComplete="off"
        layout="vertical"
        style={{ width: '90%', maxWidth: '400px' }}
      >
        <Form.Item
          label="First name"
          name="firstname"
          rules={[
            {
              required: true,
              message: 'Please input your first name!',
              pattern: /^[A-Za-zÀ-ÿ]+([ '-][A-Za-zÀ-ÿ]+)*$/
            }
          ]}
        >
          <Input
            className="input"
            size="large"
            placeholder="Enter your first name"
          />
        </Form.Item>

        <Form.Item
          label="Last name"
          name="lastname"
          rules={[
            {
              required: true,
              message: 'Please input your name!',
              pattern: /^[A-Za-zÀ-ÿ]+([ '-][A-Za-zÀ-ÿ]+)*$/
            }
          ]}
        >
          <Input className="input" size="large" placeholder="Enter your name" />
        </Form.Item>

        <Form.Item
          label="Email"
          name="email"
          rules={[
            {
              required: true,
              message: 'Please input your email!',
              type: 'email'
            }
          ]}
        >
          <Input
            className="input"
            size="large"
            placeholder="Enter your email"
          />
        </Form.Item>

        <Form.Item
          label="Password"
          name="password"
          rules={[{ required: true, message: 'Please input your password!' }]}
        >
          <Input.Password
            className="input"
            size="large"
            placeholder="Enter your password"
          />
        </Form.Item>

        <Form.Item
          label="Confirm password"
          name="confirmPassword"
          rules={[{ required: true, message: 'Please confirm your password!' }]}
        >
          <Input.Password
            className="input"
            size="large"
            placeholder="Confirm your password"
          />
        </Form.Item>
        <Form.Item name="termAgree" valuePropName="checked">
          <div className={style.optionsContainer}>
            <Checkbox>I agree to all Term, Privacy Policy and Fees</Checkbox>
          </div>
        </Form.Item>
        <Form.Item>
          <Button
            className="submitButton"
            htmlType="submit"
            size="large"
            color="default"
            variant="solid"
          >
            Sign up
          </Button>
        </Form.Item>

        <div className={style.dontHaveAccountContainer}>
          <span className={style.footerText}>Already have an account?</span>
          <span
            className={style.footerLink}
            onClick={goToLogin}
            role="link"
            tabIndex={0}
            onKeyDown={e => {
              if (e.key === 'Enter') goToLogin()
            }}
          >
            Sign in
          </span>
        </div>
      </Form>
    </AuthentificationPage>
  )
}

export default Register
