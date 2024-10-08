import { Form, Input } from 'antd'
import '@/components/Inputs/global.css'

const PasswordConfirmationInput = () => (
  <Form.Item
    name="confirmPassword"
    label="Password confirmation"
    rules={[{ required: true, message: 'Please confirm your password !' }]}
  >
    <Input.Password className="inputLogin" placeholder="Enter your password" />
  </Form.Item>
)

export default PasswordConfirmationInput
