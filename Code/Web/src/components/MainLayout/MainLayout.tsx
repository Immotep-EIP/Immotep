import React, { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { useNavigate, useLocation, Outlet } from 'react-router-dom'
import { Button, Layout, Menu, Tooltip } from 'antd'
import RealProprtyIcon from '@/assets/icons/realProperty.png'
import Overview from '@/assets/icons/overview.png'
import Messages from '@/assets/icons/messages.png'
import Immotep from '@/assets/icons/ImmotepLogo.svg'
import Settings from '@/assets/icons/settings.png'
// import Me from '@/assets/icons/me.png'
import NavigationEnum from '@/enums/NavigationEnum'
import useNavigation from '@/hooks/useNavigation/useNavigation'
import style from './MainLayout.module.css'

const { Content, Sider } = Layout

const items = [
  {
    label: 'components.button.overview',
    key: NavigationEnum.OVERVIEW,
    icon: <img src={Overview} alt="Overview" className={style.menuIcon} />
  },
  {
    label: 'components.button.realProperty',
    key: NavigationEnum.REAL_PROPERTY,
    icon: (
      <img
        src={RealProprtyIcon}
        alt="Real Property"
        className={style.menuIcon}
      />
    )
  },
  {
    label: 'components.button.messages',
    key: NavigationEnum.MESSAGES,
    icon: <img src={Messages} alt="Messages" className={style.menuIcon} />
  }
]

const MainLayout: React.FC = () => {
  const screenWidth = window.innerWidth
  const [collapsed, setCollapsed] = useState(!(screenWidth > 500))
  const [selectedTab, setSelectedTab] = useState('')
  const location = useLocation()
  const navigate = useNavigate()
  const { goToSettings/* , goToMyProfile */ } = useNavigation()

  const { t } = useTranslation()

  useEffect(() => {
    setSelectedTab(location.pathname)
  }, [location.pathname])

  const translatedItems = items.map(item => ({
    ...item,
    label: t(item.label)
  }))

  return (
    <div className={style.pageContainer}>
      <div className={style.headerContainer}>
        <div className={style.leftPartHeader}>
          <img src={Immotep} alt="logo Immotep" className={style.headerLogo} />
          <span className={style.headerTitle}>Immotep</span>
        </div>
        <div className={style.rightPartHeader}>
          <Tooltip title={t('components.button.settings')} placement="bottom">
            <Button
              shape="circle"
              style={{ marginRight: 10 }}
              color="default"
              variant="solid"
              size="middle"
              onClick={() => goToSettings()}
            >
              <img
                src={Settings}
                alt="Settings"
                style={{ width: 17, height: 17 }}
              />
            </Button>
          </Tooltip>
          {/* <Tooltip title={t('components.button.myProfile')} placement="bottom">
            <Button
              shape="circle"
              style={{ marginRight: 10 }}
              color="default"
              variant="solid"
              size="middle"
              onClick={() => goToMyProfile()}
            >
              <img src={Me} alt="Me" style={{ width: 17, height: 17 }} />
            </Button>
          </Tooltip> */}
        </div>
      </div>
      <Layout style={{ height: '100%', width: '100%' }}>
        <Sider
          collapsible={screenWidth > 500}
          collapsed={collapsed}
          onCollapse={value => setCollapsed(value)}
          theme="light"
        >
          <Menu
            theme="light"
            selectedKeys={[selectedTab]}
            mode="inline"
            items={translatedItems}
            onClick={e => {
              setSelectedTab(e.key)
              navigate(e.key)
            }}
            style={{
              height: '100%',
              boxShadow: '0 4px 8px rgba(0, 0, 0, 0.15)'
            }}
          />
        </Sider>
        <Layout>
          <Content style={{ height: '100%', width: '100%' }}>
            <Outlet />
          </Content>
        </Layout>
      </Layout>
    </div>
  )
}

export default MainLayout
