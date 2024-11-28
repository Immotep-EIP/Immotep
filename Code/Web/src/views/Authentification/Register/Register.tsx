import React from 'react'
import { useParams } from 'react-router-dom';
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
  const { contractId } = useParams();

  const { t } = useTranslation()

  const onFinish: FormProps<UserRegister>['onFinish'] = async values => {
    try {
      const { password, confirmPassword } = values
      if (password === confirmPassword) {
        const userInfo = {
          ...values,
          contractId,
        }
        await register(userInfo)
        message.success(t('pages.register.registrationSuccess'))
        form.resetFields()
        goToLogin()
      } else message.error(t('pages.register.passwordsNotMatch'))
    } catch (err: any) {
      if (err.response.status === 409)
        message.error(t('pages.register.emailAlreadyUsed'))
      console.error(t('pages.register.registrationError'), err)
    }
  }

  const onFinishFailed: FormProps<UserRegister>['onFinishFailed'] = () => {
    message.error(t('pages.register.fillFields'))
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
          label={t('components.input.firstName.label')}
          name="firstname"
          rules={[
            {
              required: true,
              message: t('components.input.firstName.error'),
              pattern: /^[A-Za-zÀ-ÿ]+([ '-][A-Za-zÀ-ÿ]+)*$/
            }
          ]}
        >
          <Input
            className="input"
            size="middle"
            placeholder={t('components.input.firstName.placeholder')}
          />
        </Form.Item>

        <Form.Item
          label={t('components.input.lastName.label')}
          name="lastname"
          rules={[
            {
              required: true,
              message: t('components.input.lastName.error'),
              pattern: /^[A-Za-zÀ-ÿ]+([ '-][A-Za-zÀ-ÿ]+)*$/
            }
          ]}
        >
          <Input
            className="input"
            size="middle"
            placeholder={t('components.input.lastName.placeholder')}
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
          rules={[{ required: true, message: t('components.input.password.error') }]}
        >
          <Input.Password
            className="input"
            size="middle"
            placeholder={t('components.input.password.placeholder')}
          />
        </Form.Item>

        <Form.Item
          label={t('components.input.confirmPassword.label')}
          name="confirmPassword"
          rules={[{ required: true, message: t('components.input.confirmPassword.error') }]}
        >
          <Input.Password
            className="input"
            size="middle"
            placeholder={t('components.input.confirmPassword.placeholder')}
          />
        </Form.Item>
        <Form.Item name="termAgree" valuePropName="checked">
          <div className={style.optionsContainer}>
            <Checkbox>
              {t('pages.register.agreeTerms')}
            </Checkbox>
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
            {t('components.button.signUp')}
          </Button>
        </Form.Item>

        <div className={style.dontHaveAccountContainer}>
          <span className={style.footerText}>
            {t('pages.register.alreadyHaveAccount')}
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
            {t('components.button.signIn')}
          </span>
        </div>
      </Form>
    </AuthentificationPage>
  )
}

export default Register
