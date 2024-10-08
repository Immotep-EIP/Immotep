import { Button, Form, Checkbox } from 'antd'
import '@/components/Login/LoginOptions.css'

const LoginOptions = ({ goToHome }: { goToHome: () => void }) => (
  <div className="optionContainer">
    <Form.Item name="remember" valuePropName="checked">
      <Checkbox>Keep me signed</Checkbox>
    </Form.Item>
    <Button type="link" onClick={goToHome}>
      Forgot password ?
    </Button>
  </div>
)

export default LoginOptions
