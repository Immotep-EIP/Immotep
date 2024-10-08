import { ReactNode, useState } from 'react'
import { Layout, Menu, Image, Typography } from 'antd'
import {
  UserOutlined,
  SettingOutlined,
  MailOutlined,
  HomeOutlined,
  AppstoreOutlined
} from '@ant-design/icons'
import type { MenuProps } from 'antd'

import useNavigation from '@/hooks/useNavigation/useNavigation'
import immotepLogo from '@/assets/icons/ImmotepLogo.svg'
import '@/components/AuthLayout/AuthLayout.css'

const { Header, Content, Sider } = Layout

const { Title } = Typography

interface LayoutProps {
  children: ReactNode
}

type MenuItem = Required<MenuProps>['items'][number]

const AuthLayout = ({ children }: LayoutProps) => {
  const { goToDashboard } = useNavigation()
  const [collapsed, setCollapsed] = useState(false)

  const items: MenuItem[] = [
    {
      key: 'overview',
      icon: <AppstoreOutlined className="menuIcon" />,
      label: 'Overview',
      onClick: goToDashboard
    },
    {
      key: 'realProperty',
      label: 'Real property',
      icon: <HomeOutlined className="menuIcon" />
    },
    {
      key: 'message',
      label: 'Messages',
      icon: <MailOutlined className="menuIcon" />
    }
  ]
  return (
    <Layout className="authLayout">
      <Header className="authHeader">
        <Image
          preview={false}
          src={immotepLogo}
          alt="Immotep logo"
          className="immotepLogo"
        />
        <Title level={3}>Immotep</Title>
        <SettingOutlined className="settingsIcon" onClick={() => {}} />
        <UserOutlined style={{ fontSize: '25px' }} onClick={() => {}} />
      </Header>
      <Layout>
        <Sider
          collapsible
          collapsed={collapsed}
          onCollapse={value => setCollapsed(value)}
        >
          <Menu
            mode="inline"
            style={{ height: '100%', borderRight: 0 }}
            items={items}
          />
        </Sider>
        <Content className="authMainContent">{children}</Content>
      </Layout>
    </Layout>
  )
}

export default AuthLayout
