import { Form, Input } from 'antd'
import '@/components/Inputs/global.css'

const PasswordInput = () => (
  <Form.Item
    name="password"
    label="Password"
    rules={[{ required: true, message: 'Please enter your password !' }]}
  >
    <Input.Password className="inputLogin" placeholder="Enter your password" />
  </Form.Item>
)

export default PasswordInput
