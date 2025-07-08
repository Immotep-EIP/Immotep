import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'

import { Form, message } from 'antd'
import type { FormProps } from 'antd'

import { Button, Input } from '@/components/common'
import AuthentificationPage from '@/components/features/authentication/AuthentificationPage/AuthentificationPage'
import useNavigation from '@/hooks/Navigation/useNavigation'
import PageMeta from '@/components/ui/PageMeta/PageMeta'

import '@/App.css'

type FieldType = {
  email?: string
  emailConfirmation?: string
}

const ForgotPassword: React.FC = () => {
  const { goToLogin } = useNavigation()
  const { t } = useTranslation()
  const [loading, setLoading] = useState(false)

  const onFinish: FormProps<FieldType>['onFinish'] = values => {
    setLoading(true)
    if (values.email !== values.emailConfirmation) {
      message.error(t('pages.forgot_password.emails_dont_match'))
      setLoading(false)
    } else {
      message.success(t('pages.forgot_password.send_email_success'))
      setLoading(false)
      goToLogin()
    }
  }

  const onFinishFailed: FormProps<FieldType>['onFinishFailed'] = () => {
    message.error(t('pages.forgot_password.fill_fields'))
  }

  return (
    <>
      <PageMeta
        title={t('pages.forgot_password.document_title')}
        description={t('pages.forgot_password.document_description')}
        keywords="forgot password, reset, authentication, Keyz"
      />
      <AuthentificationPage
        title={t('pages.forgot_password.title')}
        subtitle={t('pages.forgot_password.description')}
      >
        <section
          style={{ width: '90%', maxWidth: '400px' }}
          aria-labelledby="forgot-password-form-title"
        >
          <h2 id="forgot-password-form-title" className="sr-only">
            {t('pages.forgot_password.form_title')}
          </h2>

          <Form
            name="forgot-password-form"
            initialValues={{ remember: true }}
            onFinish={onFinish}
            onFinishFailed={onFinishFailed}
            autoComplete="off"
            layout="vertical"
            aria-labelledby="forgot-password-form-title"
            noValidate
          >
            <Form.Item
              label={t('components.input.email.label')}
              name="email"
              rules={[
                { required: true, message: t('components.input.email.error') },
                {
                  type: 'email',
                  message: t('components.input.email.valid_email')
                }
              ]}
            >
              <Input
                className="input"
                size="large"
                type="email"
                placeholder={t('components.input.email.placeholder')}
                aria-label={t('components.input.email.label')}
                aria-required="true"
                aria-describedby="email-help"
                id="forgot-password-email"
                autoComplete="email"
              />
            </Form.Item>

            <Form.Item
              label={t('components.input.email_confirmation.label')}
              name="emailConfirmation"
              rules={[
                {
                  required: true,
                  message: t('components.input.email_confirmation.error')
                },
                {
                  type: 'email',
                  message: t('components.input.email.valid_email')
                }
              ]}
            >
              <Input
                className="input"
                size="large"
                type="email"
                placeholder={t(
                  'components.input.email_confirmation.placeholder'
                )}
                aria-label={t('components.input.email_confirmation.label')}
                aria-required="true"
                aria-describedby="email-confirmation-help"
                id="forgot-password-email-confirmation"
                autoComplete="email"
              />
            </Form.Item>

            <Form.Item>
              <Button
                className="submitButton"
                htmlType="submit"
                size="large"
                color="default"
                variant="solid"
                loading={loading}
                aria-describedby="submit-button-help"
                disabled={loading}
              >
                {loading
                  ? t('pages.forgot_password.sending')
                  : t('components.button.send_email')}
              </Button>
              <div id="submit-button-help" className="sr-only">
                {t('pages.forgot_password.submit_help')}
              </div>
            </Form.Item>
          </Form>
        </section>
      </AuthentificationPage>
    </>
  )
}

export default ForgotPassword
