import { ReactNode } from 'react'

import { Layout, Image, Typography } from 'antd'
import immotepLogo from '@/assets/icons/ImmotepLogo.svg'
import '@/components/AuthLayout/AuthLayout.css'

const { Header, Content } = Layout

const { Title } = Typography

interface LayoutProps {
  children: ReactNode
}

const AuthLayout = ({ children }: LayoutProps) => (
  <Layout className="layout">
    <Header className="header">
      <Image
        preview={false}
        src={immotepLogo}
        alt="Immotep logo"
        className="immotepLogo"
      />
      <Title level={3}>Immotep Auth</Title>
    </Header>
    <Content style={{ padding: '20px' }}>{children}</Content>
  </Layout>
)

export default AuthLayout
