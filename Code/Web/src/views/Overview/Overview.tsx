import React, { useState } from 'react'
import { Responsive, WidthProvider } from 'react-grid-layout'
import { Button } from 'antd'
import { UserOutlined } from '@ant-design/icons'

import { useTranslation } from 'react-i18next'
import PageTitle from '@/components/PageText/Title.tsx'
import AddWidgetModal from '@/components/Overview/AddWidgetModal.tsx'
import UserInfoWidget from '@/components/Widgets/UserInfoWidget.tsx'
import MaintenanceWidget from '@/components/Widgets/MaintenanceWidget.tsx'
import { Layout, Widget, addWidgetType } from '@/interfaces/Widgets/Widgets.ts'
import PageMeta from '@/components/PageMeta/PageMeta'
import style from './Overview.module.css'
import '@/../node_modules/react-grid-layout/css/styles.css'
import '@/../node_modules/react-resizable/css/styles.css'

const ResponsiveGridLayout = WidthProvider(Responsive)

const WidgetTemplate: React.FC<{
  logo?: React.ReactElement
  children: React.ReactNode
}> = ({ logo, children }) => (
  <div className={style.widgetContainer}>
    <div className={style.widgetHeader}>{logo}</div>
    <div className={style.widgetContent}>{children}</div>
  </div>
)

const Overview: React.FC = () => {
  const { t } = useTranslation()
  const [isModalOpen, setIsModalOpen] = useState(false)

  const layouts: { lg: Widget[] } = {
    lg: [
      {
        i: '0',
        name: 'Widget 1',
        logo: <UserOutlined />,
        x: 0,
        y: 0,
        w: 2,
        h: 2,
        children: <UserInfoWidget height={2} />,
        minW: 2,
        maxW: 3,
        minH: 2,
        maxH: 3
      },
      {
        i: '1',
        name: 'Maintenance',
        logo: <UserOutlined />,
        x: 2,
        y: 0,
        w: 3,
        h: 4,
        children: <MaintenanceWidget height={4} />,
        minW: 3,
        maxW: 6,
        minH: 4,
        maxH: 6
      }
    ]
  }

  const [widgets, setWidgets] = useState(layouts.lg)

  const showModal = () => setIsModalOpen(true)
  const handleCancel = () => setIsModalOpen(false)

  const handleAddWidget = (widget: addWidgetType) => {
    let widgetContent: React.ReactNode = null

    switch (widget.types) {
      case 'UserInfoWidget':
        widgetContent = <UserInfoWidget height={widget.height} />
        break
      case 'MaintenanceWidget':
        widgetContent = <MaintenanceWidget height={widget.height} />
        break
      default:
        widgetContent = <div> </div>
        break
    }

    const newWidget = {
      i: String(widgets.length),
      name: widget.name,
      logo: <UserOutlined />,
      x: 0,
      y: Infinity,
      w: widget.width,
      h: widget.height,
      children: widgetContent
    }

    setWidgets([...widgets, newWidget])
  }

  const handleLayoutChange = (layout: Layout[]) => {
    const updatedWidgets = widgets.map(widget => {
      const layoutItem = layout.find(item => item.i === widget.i)
      const oldHeight = widget.h

      if (layoutItem && layoutItem.h !== oldHeight) {
        const updatedWidget = {
          ...widget,
          x: layoutItem.x,
          y: layoutItem.y,
          w: layoutItem.w,
          h: layoutItem.h
        }

        const updatedChildren = React.isValidElement(widget.children)
          ? React.cloneElement(widget.children as React.ReactElement, {
              height: updatedWidget.h
            })
          : widget.children

        return { ...updatedWidget, children: updatedChildren }
      }

      return widget
    })

    setWidgets(updatedWidgets)
  }

  return (
    <>
      <PageMeta
        title={t('pages.overview.document_title')}
        description={t('pages.overview.document_description')}
        keywords="overview, dashboard, Immotep"
      />
      <div className={style.pageContainer}>
        <div className={style.pageHeader}>
          <PageTitle title={t('pages.overview.title')} size="title" />
          <Button type="primary" onClick={showModal}>
            {t('components.button.add_widget')}
          </Button>
        </div>
        <ResponsiveGridLayout
          className={style.gridLayout}
          layouts={{ lg: widgets }}
          breakpoints={{ lg: 1200, md: 996, sm: 768, xs: 480, xxs: 0 }}
          cols={{ lg: 12, md: 10, sm: 6, xs: 4, xxs: 2 }}
          rowHeight={80}
          isResizable
          onResize={handleLayoutChange}
          draggableHandle={`.${style.widgetHeader}`}
        >
          {widgets.map((widget: Widget) => (
            <div key={widget.i} data-grid={widget}>
              <WidgetTemplate logo={widget.logo}>
                {widget.children}
              </WidgetTemplate>
            </div>
          ))}
        </ResponsiveGridLayout>
        <AddWidgetModal
          isOpen={isModalOpen}
          onClose={handleCancel}
          onAddWidget={handleAddWidget}
        />
      </div>
    </>
  )
}

export default Overview
