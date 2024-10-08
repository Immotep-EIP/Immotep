import { Form, Input } from 'antd'
import '@/components/Inputs/global.css'

const FirstNameInput = () => (
  <Form.Item
    name="firstname"
    label="First name"
    rules={[
      {
        required: true,
        message: 'Please enter your first name !'
      }
    ]}
  >
    <Input className="inputLogin" placeholder="Enter your first name" />
  </Form.Item>
)

export default FirstNameInput
