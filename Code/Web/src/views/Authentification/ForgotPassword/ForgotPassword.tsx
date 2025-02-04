import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Button, Input, Form, message } from 'antd'
import type { FormProps } from 'antd'
import AuthentificationPage from '@/components/AuthentificationPage/AuthentificationPage'
import useNavigation from '@/hooks/useNavigation/useNavigation'
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
    <AuthentificationPage
      title={t('pages.forgot_password.title')}
      subtitle={t('pages.forgot_password.description')}
    >
      <Form
        name="basic"
        initialValues={{ remember: true }}
        onFinish={onFinish}
        onFinishFailed={onFinishFailed}
        autoComplete="off"
        layout="vertical"
        style={{ width: '90%', maxWidth: '400px' }}
      >
        <Form.Item
          label={t('components.input.email.label')}
          name="email"
          rules={[
            { required: true, message: t('components.input.email.error') }
          ]}
        >
          <Input
            className="input"
            size="large"
            placeholder={t('components.input.email.placeholder')}
            aria-label={t('components.input.email.placeholder')}
          />
        </Form.Item>

        <Form.Item
          label={t('components.input.email_confirmation.label')}
          name="emailConfirmation"
          rules={[
            {
              required: true,
              message: t('components.input.email_confirmation.error')
            }
          ]}
        >
          <Input
            className="input"
            size="large"
            placeholder={t('components.input.email_confirmation.placeholder')}
            aria-label={t('components.input.email_confirmation.placeholder')}
          />
        </Form.Item>

        <Form.Item>
          <Button
            className="submitButton"
            htmlType="submit"
            // onClick={goToLogin}
            size="large"
            color="default"
            variant="solid"
            loading={loading}
            aria-label={t('components.button.send_email')}
          >
            {t('components.button.send_email')}
          </Button>
        </Form.Item>
      </Form>
    </AuthentificationPage>
  )
}

export default ForgotPassword
