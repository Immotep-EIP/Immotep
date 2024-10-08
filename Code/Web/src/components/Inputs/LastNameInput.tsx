import { Form, Input } from 'antd'
import '@/components/Inputs/global.css'

const LastNameInput = () => (
  <Form.Item
    name="lastname"
    label="Last name"
    rules={[
      {
        required: true,
        message: 'Please enter your last name !'
      }
    ]}
  >
    <Input className="inputLogin" placeholder="Enter your last name" />
  </Form.Item>
)

export default LastNameInput
