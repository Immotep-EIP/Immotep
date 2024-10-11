import React from "react";
import { Button, Input, Form, Radio, message } from "antd";
import type { FormProps } from 'antd';
import AuthentificationPage from "@/components/AuthentificationPage/AuthentificationPage";
import useNavigation from '@/hooks/useNavigation/useNavigation'
import '@/App.css';
import style from './Login.module.css';

type FieldType = {
  email?: string;
  password?: string;
  remember?: string;
};

const Login: React.FC = () => {
  const { goToSignup, goToOverview } = useNavigation();

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
      title="Welcome back !"
      subtitle="Please enter your details to sign in."
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
          label="Password"
          name="password"
          rules={[{ required: true, message: 'Please input your password!' }]}
        >
          <Input.Password className='input' size="large" placeholder="Enter your password" />
        </Form.Item>

        <div className={style.optionsContainer}>
          <Radio>Remember me</Radio>
          <span className={style.forgotPassword}>Forgot password ?</span>
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
          onKeyDown={(e) => { if (e.key === 'Enter') goToSignup(); }}
        >
          Sign up
        </span>
      </div>

      </Form>
    </AuthentificationPage>
  )
};

export default Login;
