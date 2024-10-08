import EmailInput from '@/components/Inputs/EmailInput'
import PasswordInput from '@/components/Inputs/PasswordInput'
import PasswordConfirmationInput from '@/components/Inputs/PasswordConfirmationInput'
import LastNameInput from '@/components/Inputs/LastNameInput'
import FirstNameInput from '@/components/Inputs/FirstNameInput'

const RegisterInputs = () => (
  <>
    <LastNameInput />
    <FirstNameInput />
    <EmailInput />
    <PasswordInput />
    <PasswordConfirmationInput />
  </>
)

export default RegisterInputs
