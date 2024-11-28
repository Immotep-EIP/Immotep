import React from "react";
import { useTranslation } from "react-i18next";
import { Button, Input, Form, message } from "antd";
import type { FormProps } from 'antd';
import AuthentificationPage from "@/components/AuthentificationPage/AuthentificationPage";
import useNavigation from '@/hooks/useNavigation/useNavigation'
import '@/App.css';

type FieldType = {
  email?: string;
  emailConfirmation?: string;
};

const ForgotPassword: React.FC = () => {
  const { goToLogin } = useNavigation();
  const { t } = useTranslation();

  const onFinish: FormProps<FieldType>['onFinish'] = (values) => {
    if (values.email !== values.emailConfirmation) {
      message.error(t('pages.forgotPassword.emailsDontMatch'));
    } else {
      message.success(t('pages.forgotPassword.sendEmailSuccess'));
      goToLogin();
    }
  };

  const onFinishFailed: FormProps<FieldType>['onFinishFailed'] = () => {
    message.error(t('pages.forgotPassword.fillFields'));
  };

  return (
    <AuthentificationPage
      title={t('pages.forgotPassword.title')}
      subtitle={t('pages.forgotPassword.description')}
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
          rules={[{ required: true, message: t('components.input.email.error') }]}
        >
          <Input
            className='input'
            size="large"
            placeholder={t('components.input.email.placeholder')}
          />
        </Form.Item>

        <Form.Item
          label={t('components.input.emailConfirmation.label')}
          name="emailConfirmation"
          rules={[{ required: true, message: t('components.input.emailConfirmation.error') }]}
        >
          <Input
            className='input'
            size="large"
            placeholder={t('components.input.emailConfirmation.placeholder')}
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
          >
            {t('components.button.sendEmail')}
          </Button>
        </Form.Item>

      </Form>
    </AuthentificationPage>
  )
};

export default ForgotPassword;
