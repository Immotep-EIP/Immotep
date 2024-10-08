import { Button, Typography } from 'antd';

import '@/views/Home/Home.css';
import useNavigation from '@/hooks/useNavigation/useNavigation';

const { Title } = Typography;

const Home = () => {
  const { goToLogin, goToSignup } = useNavigation();

  return (
    <div className="container">
      <Title level={1}>Bienvenue sur Immotep</Title>
      <div className="buttonContainer">
        <Button type="primary" onClick={goToLogin}>
          Se connecter
        </Button>
        <Button type="default" onClick={goToSignup}>
          S&apos;inscrire
        </Button>
      </div>
    </div>
  );
};

export default Home;
