import React, { useState } from 'react'
import { useParams } from 'react-router-dom'
import { useTranslation } from 'react-i18next'

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
  const { contractId } = useParams()
  const [loading, setLoading] = useState(false)

  const { t } = useTranslation()

  const onFinish: FormProps<UserRegister>['onFinish'] = async values => {
    try {
      setLoading(true)
      const { password, confirmPassword } = values
      if (password === confirmPassword) {
        const userInfo = {
          ...values,
          contractId
        }
        await register(userInfo)
        message.success(t('pages.register.register_success'))
        form.resetFields()
        setLoading(false)
        goToLogin()
      } else message.error(t('pages.register.confirm_password_error'))
    } catch (err: any) {
      if (err.response.status === 409)
        message.error(t('pages.register.email_already_used'))
      console.error(t('pages.register.register_error'), err)
      setLoading(false)
    }
  }

  const onFinishFailed: FormProps<UserRegister>['onFinishFailed'] = () => {
    message.error(t('pages.register.fill_fields'))
  }

  return (
    <AuthentificationPage
      title={t('pages.register.title')}
      subtitle={t('pages.register.description')}
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
          label={t('components.input.first_name.label')}
          name="firstname"
          rules={[
            {
              required: true,
              message: t('components.input.first_name.error'),
              pattern: /^[A-Za-zÀ-ÿ]+([ '-][A-Za-zÀ-ÿ]+)*$/
            }
          ]}
        >
          <Input
            className="input"
            size="middle"
            placeholder={t('components.input.first_name.placeholder')}
          />
        </Form.Item>

        <Form.Item
          label={t('components.input.last_name.label')}
          name="lastname"
          rules={[
            {
              required: true,
              message: t('components.input.last_name.error'),
              pattern: /^[A-Za-zÀ-ÿ]+([ '-][A-Za-zÀ-ÿ]+)*$/
            }
          ]}
        >
          <Input
            className="input"
            size="middle"
            placeholder={t('components.input.last_name.placeholder')}
          />
        </Form.Item>

        <Form.Item
          label={t('components.input.email.label')}
          name="email"
          rules={[
            {
              required: true,
              message: t('components.input.email.error'),
              type: 'email'
            }
          ]}
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
          rules={[
            { required: true, message: t('components.input.password.error') }
          ]}
        >
          <Input.Password
            className="input"
            size="middle"
            placeholder={t('components.input.password.placeholder')}
          />
        </Form.Item>

        <Form.Item
          label={t('components.input.confirm_password.label')}
          name="confirmPassword"
          rules={[
            {
              required: true,
              message: t('components.input.confirm_password.error')
            }
          ]}
        >
          <Input.Password
            className="input"
            size="middle"
            placeholder={t('components.input.confirm_password.placeholder')}
          />
        </Form.Item>
        <Form.Item name="termAgree" valuePropName="checked">
          <div className={style.optionsContainer}>
            <Checkbox>{t('pages.register.agree_terms')}</Checkbox>
          </div>
        </Form.Item>
        <Form.Item>
          <Button
            className="submitButton"
            htmlType="submit"
            size="large"
            color="default"
            variant="solid"
            loading={loading}
          >
            {t('components.button.sign_up')}
          </Button>
        </Form.Item>

        <div className={style.dontHaveAccountContainer}>
          <span className={style.footerText}>
            {t('pages.register.already_have_account')}
          </span>
          <span
            className={style.footerLink}
            onClick={goToLogin}
            role="link"
            tabIndex={0}
            onKeyDown={e => {
              if (e.key === 'Enter') goToLogin()
            }}
          >
            {t('components.button.sign_in')}
          </span>
        </div>
      </Form>
    </AuthentificationPage>
  )
}

export default Register
