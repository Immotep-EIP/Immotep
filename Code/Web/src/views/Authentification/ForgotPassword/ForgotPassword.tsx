import React from "react";
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

  const onFinish: FormProps<FieldType>['onFinish'] = (values) => {
    if (values.email !== values.emailConfirmation) {
      message.error('Email and email confirmation do not match');
      return;
    }
    if (!values) {
      goToLogin();
    }
  };

  const onFinishFailed: FormProps<FieldType>['onFinishFailed'] = () => {
    message.error('An error occured, please try again');
  };

  return (
    <AuthentificationPage
      title="Forgot password"
      subtitle="Please enter your email to reset your password."
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
          label="Email"
          name="email"
          rules={[{ required: true, message: 'Please input your email!' }]}
        >
          <Input className='input' size="large" placeholder="Enter your email" />
        </Form.Item>

        <Form.Item
          label="Email confirmation"
          name="emailConfirmation"
          rules={[{ required: true, message: 'Please input your email confirmation!' }]}
        >
          <Input className='input' size="large" placeholder="Enter your email confirmation" />
        </Form.Item>

        <Form.Item>
          <Button
            className="submitButton"
            // htmlType="submit"
            onClick={goToLogin}
            size="large"
            color="default"
            variant="solid"
          >
            Send email
          </Button>
        </Form.Item>

      </Form>
    </AuthentificationPage>
  )
};

export default ForgotPassword;
