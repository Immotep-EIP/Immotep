import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { useTranslation } from 'react-i18next'

import { Form, Input as AntInput, message, Checkbox } from 'antd'
import type { FormProps } from 'antd'

import { useAuth } from '@/context/authContext'
import useNavigation from '@/hooks/Navigation/useNavigation'
import { Button, Input } from '@/components/common'
import DividedPage from '@/components/layout/DividedPage/DividedPage'
import PageMeta from '@/components/ui/PageMeta/PageMeta'
import PageTitle from '@/components/ui/PageText/Title'
import AcceptInvite from '@/services/api/Tenant/AcceptInvite'

import { UserTokenPayload } from '@/interfaces/User/User'

import backgroundImg from '@/assets/images/building.jpg'
import '@/App.css'
import style from './Login.module.css'

const Login: React.FC = () => {
  const {
    goToSignup,
    goToOverview,
    goToForgotPassword,
    goToSuccessLoginTenant
  } = useNavigation()
  const { login } = useAuth()
  const [loading, setLoading] = useState(false)
  const { leaseId } = useParams()

  const { t } = useTranslation()

  useEffect(() => {
    if (
      sessionStorage.getItem('access_token') &&
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

  const onFinish: FormProps<UserTokenPayload>['onFinish'] = async values => {
    setLoading(true)
    try {
      const loginValues = {
        ...values,
        grant_type: 'password'
      }
      loginValues.grant_type = 'password'
      await login(loginValues, leaseId)
      message.success(t('pages.login.connection_success'))
      setLoading(false)
      if (leaseId) {
        await AcceptInvite(leaseId)
        goToSuccessLoginTenant()
      } else {
        goToOverview()
      }
    } catch (error: any) {
      if (error?.response?.status === 401) {
        message.error(t('pages.login.connection_error'))
        setLoading(false)
      }
      setLoading(false)
    }
  }

  const onFinishFailed: FormProps<UserTokenPayload>['onFinishFailed'] = () => {
    message.error(t('pages.login.fill_fields'))
  }

  return (
    <>
      <PageMeta
        title={t('pages.login.document_title')}
        description={t('pages.login.document_description')}
        keywords="login, authentication, Keyz"
      />
      <DividedPage
        childrenLeft={
          <img
            src={backgroundImg}
            alt="background"
            className={style.backgroundImg}
          />
        }
        childrenRight={
          <>
            <PageTitle title={t('pages.login.title')} size="title" />
            <PageTitle title={t('pages.login.description')} size="subtitle" />
            <section
              style={{ width: '90%', maxWidth: '400px' }}
              aria-labelledby="login-form-title"
            >
              <h2 id="login-form-title" className="sr-only">
                {t('pages.login.form_title')}
              </h2>
              <Form
                name="login"
                initialValues={{ rememberMe: false }}
                onFinish={onFinish}
                onFinishFailed={onFinishFailed}
                autoComplete="on"
                layout="vertical"
                aria-labelledby="login-form-title"
                noValidate
              >
                <legend className="sr-only">
                  {t('pages.login.form_legend')}
                </legend>

                <Form.Item
                  label={t('components.input.email.label')}
                  name="username"
                  rules={[
                    {
                      required: true,
                      message: t('components.input.email.error')
                    },
                    {
                      type: 'email',
                      message: t('components.input.email.valid_email')
                    }
                  ]}
                >
                  <Input
                    id="login-email"
                    type="email"
                    className="input"
                    size="middle"
                    placeholder={t('components.input.email.placeholder')}
                    aria-label={t('components.input.email.label')}
                    aria-required="true"
                    aria-describedby="login-email-help"
                    autoComplete="username"
                  />
                </Form.Item>
                <div id="login-email-help" className="sr-only">
                  {t('pages.login.email_help')}
                </div>

                <Form.Item
                  label={t('components.input.password.label')}
                  name="password"
                  rules={[
                    {
                      required: true,
                      message: t('components.input.password.error')
                    }
                  ]}
                >
                  <AntInput.Password
                    id="login-password"
                    className="input"
                    size="middle"
                    placeholder={t('components.input.password.placeholder')}
                    aria-label={t('components.input.password.label')}
                    aria-required="true"
                    aria-describedby="login-password-help"
                    autoComplete="current-password"
                  />
                </Form.Item>
                <div id="login-password-help" className="sr-only">
                  {t('pages.login.password_help')}
                </div>

                <div className={style.optionsContainer}>
                  <Form.Item name="rememberMe" valuePropName="checked">
                    <Checkbox id="login-remember">
                      {t('components.button.remember_me')}
                    </Checkbox>
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
                    {t('components.button.ask_forgot_password')}
                  </span>
                </div>

                <Form.Item>
                  <Button
                    className="submitButton"
                    htmlType="submit"
                    size="large"
                    color="default"
                    variant="solid"
                    loading={loading}
                    disabled={loading}
                    aria-describedby="login-submit-help"
                  >
                    {loading
                      ? t('components.button.signing_in')
                      : t('components.button.sign_in')}
                  </Button>
                </Form.Item>
                <div id="login-submit-help" className="sr-only">
                  {t('pages.login.submit_help')}
                </div>
              </Form>

              <div className={style.dontHaveAccountContainer}>
                <span className={style.footerText}>
                  {t('pages.login.dont_have_account')}
                </span>
                <Button
                  type="link"
                  style={{ border: 'none', padding: 0 }}
                  className={style.footerLink}
                  onClick={goToSignup}
                  aria-label={t('components.button.sign_up')}
                >
                  {t('components.button.sign_up')}
                </Button>
              </div>
            </section>
          </>
        }
      />
    </>
  )
}

export default Login
