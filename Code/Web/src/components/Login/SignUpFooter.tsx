import { Typography, Button } from 'antd'
import '@/components/Login/global.css'

const SignUpFooter = ({ goToSignup }: { goToSignup: () => void }) => (
  <Typography.Text>
    Don’t have an account ?
    <Button type="link" onClick={goToSignup}>
      Sign up
    </Button>
  </Typography.Text>
)

export default SignUpFooter
