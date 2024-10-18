import React from 'react'
import { Button, Input, Form, message, Checkbox } from 'antd'
import type { FormProps } from 'antd'

import AuthentificationPage from '@/components/AuthentificationPage/AuthentificationPage'
import useNavigation from '@/hooks/useNavigation/useNavigation'
import '@/App.css'
import { useAuth } from '@/context/authContext'
import { UserToken } from '@/interfaces/User/User'
import style from './Login.module.css'

const Login: React.FC = () => {
  const { goToSignup, goToOverview, goToForgotPassword } = useNavigation()
  const { login } = useAuth()

  const onFinish: FormProps<UserToken>['onFinish'] = async values => {
    try {
      const loginValues = {
        ...values,
        grant_type: 'password'
      }
      loginValues.grant_type = 'password'
      await login(loginValues)
      message.success('Login successful')
      goToOverview()
    } catch (error: any) {
      if (error.response.status === 401)
        message.error('Login failed, please try again !')
    }
  }

  const onFinishFailed: FormProps<UserToken>['onFinishFailed'] = () => {
    message.error('An error occured, please try again')
  }

  return (
    <AuthentificationPage
      title="Welcome back !"
      subtitle="Please enter your details to sign in."
    >
      <Form
        name="basic"
        initialValues={{ rememberMe: false }}
        onFinish={onFinish}
        onFinishFailed={onFinishFailed}
        autoComplete="off"
        layout="vertical"
        style={{ width: '90%', maxWidth: '400px' }}
      >
        <Form.Item
          label="Email"
          name="username"
          rules={[{ required: true, message: 'Please input your email!' }]}
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

        <div className={style.optionsContainer}>
          <Form.Item name="rememberMe" valuePropName="checked">
            <Checkbox>Remember me</Checkbox>
          </Form.Item>
          <span
            className={style.footerLink}
            onClick={goToForgotPassword}
            role="link"
            tabIndex={0}
            onKeyDown={e => {
              if (e.key === 'Enter') goToForgotPassword()
            }}
          >
            Forgot password ?
          </span>
        </div>

        <Form.Item>
          <Button
            className="submitButton"
            htmlType="submit"
            size="large"
            color="default"
            variant="solid"
          >
            Sign in
          </Button>
        </Form.Item>

        <div className={style.dontHaveAccountContainer}>
          <span className={style.footerText}>Don&apos;t have an account?</span>
          <span
            className={style.footerLink}
            onClick={goToSignup}
            role="link"
            tabIndex={0}
            onKeyDown={e => {
              if (e.key === 'Enter') goToSignup()
            }}
          >
            Sign up
          </span>
        </div>
      </Form>
    </AuthentificationPage>
  )
}

export default Login
