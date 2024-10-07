import { ReactNode } from 'react'
// import { Link } from 'react-router-dom';

import { Layout, Image, Typography } from 'antd'
import immotepLogo from '@/assets/icons/ImmotepLogo.svg'
import '@/components/PublicLayout/PublicLayout.css'

const { Header, Content } = Layout

const { Title } = Typography
interface LayoutProps {
  children: ReactNode
}

const PublicLayout = ({ children }: LayoutProps) => (
  <Layout className="layout">
    <Header className="header">
      <Image
        preview={false}
        src={immotepLogo}
        alt="Immotep logo"
        className="immotepLogo"
      />
      <Title level={3}>Immotep Public</Title>
    </Header>
    <Content className="mainContent">{children}</Content>
  </Layout>
)

export default PublicLayout
