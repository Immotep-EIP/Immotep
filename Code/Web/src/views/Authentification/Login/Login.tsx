import React, { useEffect } from 'react'
import { useTranslation } from 'react-i18next'

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

  const { t } = useTranslation()

  useEffect(() => {
    if (sessionStorage.getItem('access_token') &&
      sessionStorage.getItem('refresh_token') &&
      sessionStorage.getItem('expires_in')
    ) {
      sessionStorage.removeItem('access_token')
      sessionStorage.removeItem('refresh_token')
      sessionStorage.removeItem('expires_in')
    }

    if (
      localStorage.getItem('access_token') &&
      localStorage.getItem('refresh_token') &&
      localStorage.getItem('expires_in')
    ) {
      goToOverview()
    }
  }, [])

  const onFinish: FormProps<UserToken>['onFinish'] = async values => {
    try {
      const loginValues = {
        ...values,
        grant_type: 'password'
      }
      loginValues.grant_type = 'password'
      await login(loginValues)
      message.success(t('pages.login.connectionSuccess'))
      goToOverview()
    } catch (error: any) {
      if (error.response.status === 401)
        message.error(t('pages.login.connectionError'))
    }
  }

  const onFinishFailed: FormProps<UserToken>['onFinishFailed'] = () => {
    message.error(t('pages.login.fillFields'))
  }

  return (
    <AuthentificationPage
      title={t('pages.login.title')}
      subtitle={t('pages.login.description')}
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
          label={t('components.input.email.label')}
          name="username"
          rules={[{ required: true, message: t('components.input.email.error') }]}
        >
          <Input
            className="input"
            size="middle"
            placeholder={t('components.input.email.placeholder')}
          />
        </Form.Item>

        <Form.Item
          label={t('components.input.password.label')}
          name="password"
          rules={[{ required: true, message: t('components.input.password.error') }]}
        >
          <Input.Password
            className="input"
            size="middle"
            placeholder={t('components.input.password.placeholder')}
          />
        </Form.Item>

        <div className={style.optionsContainer}>
          <Form.Item name="rememberMe" valuePropName="checked">
            <Checkbox>{t('components.button.rememberMe')}</Checkbox>
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
            {t('components.button.askForgotPassword')}
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
            {t('components.button.signIn')}
          </Button>
        </Form.Item>

        <div className={style.dontHaveAccountContainer}>
          <span className={style.footerText}>
            {t('pages.login.dontHaveAccount')}
          </span>
          <span
            className={style.footerLink}
            onClick={goToSignup}
            role="link"
            tabIndex={0}
            onKeyDown={e => {
              if (e.key === 'Enter') goToSignup()
            }}
          >
            {t('components.button.signUp')}
          </span>
        </div>
      </Form>
    </AuthentificationPage>
  )
}

export default Login
