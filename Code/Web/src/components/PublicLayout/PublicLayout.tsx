import { ReactNode } from 'react'

import { Layout, Image, Typography } from 'antd'
import immotepLogo from '@/assets/icons/ImmotepLogo.svg'
import '@/components/PublicLayout/PublicLayout.css'

const { Header, Content } = Layout

const { Title } = Typography
interface LayoutProps {
  children: ReactNode
}

const PublicLayout = ({ children }: LayoutProps) => (
  <Layout className="pubLayout">
    <Header className="pubHeader">
      <Image
        preview={false}
        src={immotepLogo}
        alt="Immotep logo"
        className="immotepLogo"
      />
      <Title level={3}>Immotep Public</Title>
    </Header>
    <Content className="pubMainContent">{children}</Content>
  </Layout>
)

export default PublicLayout
