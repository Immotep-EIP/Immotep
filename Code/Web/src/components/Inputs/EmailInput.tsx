import { Form, Input } from 'antd'
import '@/components/Inputs/global.css'

const EmailInput = () => (
  <Form.Item
    name="email"
    label="Email"
    rules={[
      {
        required: true,
        message: 'Please enter your email !'
      }
    ]}
  >
    <Input className="inputLogin" placeholder="Enter your email" />
  </Form.Item>
)

export default EmailInput
