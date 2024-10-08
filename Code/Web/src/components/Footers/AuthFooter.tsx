import { Typography, Button } from 'antd'
import '@/components/Login/global.css'

const AuthFooter = ({
  goTo,
  text,
  buttonText
}: {
  goTo: () => void
  text: string
  buttonText: string
}) => (
  <Typography.Text>
    {text}
    <Button type="link" onClick={goTo}>
      {buttonText}
    </Button>
  </Typography.Text>
)

export default AuthFooter
