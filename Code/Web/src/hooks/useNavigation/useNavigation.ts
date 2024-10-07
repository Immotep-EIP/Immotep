import { useNavigate } from 'react-router-dom';

const useNavigation = () => {
  const navigate = useNavigate();

  const goToLogin = () => {
    navigate('/login');
  };

  const goToSignup = () => {
    navigate('/register');
  };

  const goToHome = () => {
    navigate('/');
  };

  return {
    goToLogin,
    goToSignup,
    goToHome,
  };
};

export default useNavigation;
