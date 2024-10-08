import { Typography } from 'antd'
import '@/components/Login/global.css'

const { Title, Text } = Typography

const LoginTitle = () => (
  <>
    <Title level={3}>Welcome back</Title>
    <Text>Please enter your details to sign in.</Text>
  </>
)

export default LoginTitle
