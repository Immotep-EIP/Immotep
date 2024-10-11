import React, { useEffect, useState } from 'react';
import { useNavigate, useLocation, Outlet } from 'react-router-dom';
import { Button, Layout, Menu } from 'antd';
import RealProprtyIcon from '@/assets/icons/realProperty.png';
import Overview from '@/assets/icons/overview.png';
import Messages from '@/assets/icons/messages.png';
import Immotep from '@/assets/icons/ImmotepLogo.svg';
import Settings from '@/assets/icons/settings.png';
import Me from '@/assets/icons/me.png';
import NavigationEnum from '@/enums/NavigationEnum';
import useNavigation from '@/hooks/useNavigation/useNavigation';
import style from './MainLayout.module.css';

const { Content, Sider } = Layout;

const items = [
  {
    label: 'Overview',
    key: NavigationEnum.OVERVIEW,
    icon: <img src={Overview} alt="Overview" className={style.menuIcon} /> },
  {
    label: 'Real property',
    key: NavigationEnum.REAL_PROPERTY,
    icon: <img src={RealProprtyIcon} alt="Real Property" className={style.menuIcon} /> },
  {
    label: 'Messages',
    key: NavigationEnum.MESSAGES,
    icon: <img src={Messages} alt="Messages" className={style.menuIcon} /> }
];

const MainLayout: React.FC = () => {
  const [collapsed, setCollapsed] = useState(false);
  const [selectedTab, setSelectedTab] = useState('');
  const location = useLocation();
  const navigate = useNavigate();
  const { goToSettings, goToMyProfile } = useNavigation();

  useEffect(() => {
    setSelectedTab(location.pathname);
  }, [location.pathname]);

  return (
    <div className={style.pageContainer}>
      <div className={style.headerContainer}>
        <div className={style.leftPartHeader}>
          <img src={Immotep} alt="logo Immotep" className={style.headerLogo} />
          <span className={style.headerTitle}>Immotep</span>
        </div>
        <div className={style.rightPartHeader}>
          <Button
            shape='circle'
            style={{ marginRight: 10 }}
            color="default"
            variant="solid"
            size='large'
            onClick={() => goToSettings()}
          >
            <img src={Settings} alt="Settings" style={{ width: 20, height: 20 }} />
          </Button>
          <Button
            shape='circle'
            style={{ marginRight: 10 }}
            color="default"
            variant="solid"
            size='large'
            onClick={() => goToMyProfile()}
          >
            <img src={Me} alt="Me" style={{ width: 20, height: 20 }} />
          </Button>
        </div>
      </div>
      <Layout>
        <Sider collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)} theme="light">
          <Menu
            theme="light"
            selectedKeys={[selectedTab]}
            mode="inline"
            items={items}
            onClick={(e) => {
              setSelectedTab(e.key);
              navigate(e.key);
            }}
            style={{ height: '100%', boxShadow: '0 4px 8px rgba(0, 0, 0, 0.15)' }}
          />
        </Sider>
        <Layout>
          <Content style={{ margin: '0 16px', padding: 24, minHeight: 360 }}>
            <Outlet />
          </Content>
        </Layout>
      </Layout>
    </div>
  );
};

export default MainLayout;
