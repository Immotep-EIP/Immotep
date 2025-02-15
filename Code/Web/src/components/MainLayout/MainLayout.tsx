import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { useNavigate, Outlet, useLocation } from 'react-router-dom'
import { Layout, Menu } from 'antd'

import Property from '@/assets/icons/realProperty.png'
import Overview from '@/assets/icons/overview.png'
import Messages from '@/assets/icons/messages.png'
import Immotep from '@/assets/icons/ImmotepLogo.svg'
import Settings from '@/assets/icons/settings.png'
import NavigationEnum from '@/enums/NavigationEnum'
import type { MenuProps } from 'antd'
import { t } from 'i18next'
import style from './MainLayout.module.css'

const { Content } = Layout

const items = [
  {
    label: t('components.button.overview'),
    key: NavigationEnum.OVERVIEW,
    icon: <img src={Overview} alt="Overview" className={style.menuIcon} />
  },
  {
    label: t('components.button.real_property'),
    key: NavigationEnum.REAL_PROPERTY,
    icon: <img src={Property} alt="Real Property" className={style.menuIcon} />
  },
  {
    label: t('components.button.messages'),
    key: NavigationEnum.MESSAGES,
    icon: <img src={Messages} alt="Messages" className={style.menuIcon} />
  },
  {
    label: t('components.button.settings'),
    key: NavigationEnum.SETTINGS,
    icon: <img src={Settings} alt="Settings" className={style.menuIcon} />
  }
]

const MainLayout: React.FC = () => {
  const navigate = useNavigate()
  const [menuOpen, setMenuOpen] = useState(false)

  const location = useLocation()
  const currentLocation = `/${location.pathname.split('/')[1] || ''}`

  const toggleMenu = () => {
    setMenuOpen(!menuOpen)
  }

  const { t } = useTranslation()

  const translatedItems = items.map(item => ({
    ...item,
    label: t(item.label)
  }))

  const onClick: MenuProps['onClick'] = e => {
    navigate(e.key)
  }

  return (
    <div className={style.pageContainer}>
      <div className={style.headerContainer}>
        <div className={style.leftPartHeader}>
          <img src={Immotep} alt="logo Immotep" className={style.headerLogo} />
          <span className={style.headerTitle}>Immotep</span>
        </div>
        <div className={style.rightPartHeader}>
          <div
            className={`${style.menuToggleButton} ${menuOpen ? style.open : ''}`}
            onClick={toggleMenu}
            onKeyDown={e => {
              if (e.key === 'Enter' || e.key === ' ') {
                toggleMenu()
              }
            }}
            tabIndex={0}
            role="button"
            aria-label="Toggle menu"
          >
            <div className={style.burgerLine} />
            <div className={style.burgerLine} />
            <div className={style.burgerLine} />
          </div>

          {menuOpen && (
            <div className={style.menuDropdown}>
              {translatedItems.map(item => (
                <div
                  key={item.key}
                  className={
                    currentLocation === item.key
                      ? style.menuItemActive
                      : style.menuItem
                  }
                  onClick={() => {
                    navigate(item.key)
                    setMenuOpen(false)
                  }}
                  onKeyDown={e => {
                    if (e.key === 'Enter' || e.key === ' ') {
                      navigate(item.key)
                      setMenuOpen(false)
                    }
                  }}
                  tabIndex={0}
                  role="button"
                  aria-label={t(item.label)}
                >
                  {item.icon}
                  <span>{item.label}</span>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>

      <Layout
        style={{
          height: '100%',
          width: '100%',
          display: 'flex',
          flexDirection: 'row'
        }}
      >
        <Menu
          onClick={onClick}
          style={{ width: 256, paddingTop: '70px' }}
          defaultSelectedKeys={[NavigationEnum.OVERVIEW]}
          selectedKeys={[currentLocation]}
          mode="inline"
          items={items}
          className={style.menu}
        />
        <Content style={{ height: '100%', width: '100%', paddingTop: '70px' }}>
          <Outlet />
        </Content>
      </Layout>
    </div>
  )
}

export default MainLayout
