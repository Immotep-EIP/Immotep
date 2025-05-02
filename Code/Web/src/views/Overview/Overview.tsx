import React, { useState } from 'react'
import { Responsive, WidthProvider } from 'react-grid-layout'
import { Button } from 'antd'
import { UserOutlined } from '@ant-design/icons'

import { useTranslation } from 'react-i18next'
import PageTitle from '@/components/PageText/Title.tsx'
import AddWidgetModal from '@/components/Overview/AddWidgetModal.tsx'
import MaintenanceWidget from '@/components/Widgets/MaintenanceWidget.tsx'
import PropertiesNumber from '@/components/Widgets/PropertiesNumber.tsx'
import PropertiesRepartition from '@/components/Widgets/PropertiesRepartition.tsx'
import PropertiesDamages from '@/components/Widgets/PropertiesDamages.tsx'
import { Layout, Widget, addWidgetType } from '@/interfaces/Widgets/Widgets.ts'
import PageMeta from '@/components/PageMeta/PageMeta'
import style from './Overview.module.css'
import '@/../node_modules/react-grid-layout/css/styles.css'
import '@/../node_modules/react-resizable/css/styles.css'

const ResponsiveGridLayout = WidthProvider(Responsive)

const WidgetTemplate: React.FC<{
  children: React.ReactNode
}> = ({ children }) => (
  <div className={style.widgetContainer}>
    <div className={style.widgetContent}>{children}</div>
  </div>
)

const Overview: React.FC = () => {
  const { t } = useTranslation()
  const [isModalOpen, setIsModalOpen] = useState(false)

  const layouts: { lg: Widget[] } = {
    lg: [
      {
        i: '1',
        name: 'PropertiesNumber',
        x: 0,
        y: 0,
        w: 1,
        h: 1,
        children: <PropertiesNumber height={1} />
      },
      {
        i: '2',
        name: 'PropertiesRepartition',
        x: 1,
        y: 0,
        w: 1,
        h: 1,
        children: <PropertiesRepartition height={1} />
      },
      {
        i: '3',
        name: 'PropertiesDamages',
        x: 0,
        y: 1,
        w: 2,
        h: 2,
        children: <PropertiesDamages height={2} />
      }
      // {
      //   i: '2',
      //   name: 'Maintenance',
      //   logo: (
      //     <img
      //       src={MoveWidgetIcon}
      //       alt="move widget"
      //       style={{ width: '23px' }}
      //     />
      //   ),
      //   x: 0,
      //   y: 2,
      //   w: 3,
      //   h: 4,
      //   children: <MaintenanceWidget height={4} />,
      // }
    ]
  }

  const [widgets, setWidgets] = useState(layouts.lg)

  const showModal = () => setIsModalOpen(true)
  const handleCancel = () => setIsModalOpen(false)

  const handleAddWidget = (widget: addWidgetType) => {
    let widgetContent: React.ReactNode = null

    switch (widget.types) {
      case 'PropertiesNumber':
        widgetContent = <PropertiesNumber height={widget.height} />
        break
      case 'MaintenanceWidget':
        widgetContent = <MaintenanceWidget height={widget.height} />
        break
      case 'PropertiesRepartition':
        widgetContent = <PropertiesRepartition height={widget.height} />
        break
      case 'PropertiesDamages':
        widgetContent = <PropertiesDamages height={widget.height} />
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
        keywords="overview, dashboard, Keyz"
      />
      <div className={style.pageContainer}>
        <div className={style.pageHeader}>
          <PageTitle title={t('pages.overview.title')} size="title" />
          <Button type="primary" onClick={showModal}>
            {t('components.button.add_widget')}
          </Button>
        </div>
        <div className={style.contentContainer}>
          <ResponsiveGridLayout
            className={style.gridLayout}
            layouts={{ lg: widgets }}
            breakpoints={{ lg: 1200, md: 996, sm: 768, xs: 480, xxs: 0 }}
            cols={{ lg: 12, md: 10, sm: 6, xs: 4, xxs: 2 }}
            rowHeight={120}
            isResizable={false}
            onResize={handleLayoutChange}
            draggableHandle={`.${style.widgetContainer}`}
            preventCollision
            compactType={null}
          >
            {widgets.map((widget: Widget) => (
              <div key={widget.i} data-grid={widget}>
                <WidgetTemplate>{widget.children}</WidgetTemplate>
              </div>
            ))}
          </ResponsiveGridLayout>
        </div>
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
