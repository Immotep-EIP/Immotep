import React from "react";
import { Button, Input, Form, Radio, message } from "antd";
import type { FormProps } from 'antd';
import AuthentificationPage from "@/components/AuthentificationPage/AuthentificationPage";
import useNavigation from '@/hooks/useNavigation/useNavigation'
import '@/App.css';
import style from './Register.module.css';

type FieldType = {
  name?: string;
  firstName?: string;
  email?: string;
  password?: string;
};

const Register: React.FC = () => {
  const { goToLogin, goToOverview } = useNavigation();

  const onFinish: FormProps<FieldType>['onFinish'] = (values) => {
    if (!values) {
      goToOverview();
    }
  };

  const onFinishFailed: FormProps<FieldType>['onFinishFailed'] = () => {
    message.error('An error occured, please try again');
  };

  return (
    <AuthentificationPage
      title="Create your account"
      subtitle="Join immotep for your peace of mind !"
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
          label="First name"
          name="firstName"
          rules={[{ required: true, message: 'Please input your first name!' }]}
        >
          <Input className='input' size="large" placeholder="Enter your first name" />
        </Form.Item>

        <Form.Item
          label="Last name"
          name="name"
          rules={[{ required: true, message: 'Please input your name!' }]}
        >
          <Input className='input' size="large" placeholder="Enter your name" />
        </Form.Item>

        <Form.Item
          label="Email"
          name="email"
          rules={[{ required: true, message: 'Please input your email!' }]}
        >
          <Input className='input' size="large" placeholder="Enter your email" />
        </Form.Item>

        <Form.Item
          label="Password"
          name="password"
          rules={[{ required: true, message: 'Please input your password!' }]}
        >
          <Input.Password className='input' size="large" placeholder="Enter your password" />
        </Form.Item>

        <Form.Item
          label="Confirm password"
          name="password"
          rules={[{ required: true, message: 'Please confirm your password!' }]}
        >
          <Input.Password className='input' size="large" placeholder="Confirm your password" />
        </Form.Item>

        <div className={style.optionsContainer}>
          <Radio>I agree to all Term, Privacy Policy and Fees</Radio>
        </div>


        <Form.Item>
          <Button
            className="submitButton"
            // htmlType="submit"
            onClick={goToOverview}
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
          onKeyDown={(e) => { if (e.key === 'Enter') goToLogin(); }}
        >
          Sign in
        </span>
      </div>

      </Form>
    </AuthentificationPage>
  )
};

export default Register;
