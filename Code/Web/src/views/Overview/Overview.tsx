import React, { useState } from 'react'
import { Responsive, WidthProvider } from 'react-grid-layout'
import { Button } from 'antd'

import { useTranslation } from 'react-i18next'
import MoveWidgetIcon from '@/assets/icons/move.png'
import PageTitle from '@/components/PageText/Title.tsx'
import PropertiesNumber from '@/components/Widgets/PropertiesNumber.tsx'
import PropertiesRepartition from '@/components/Widgets/PropertiesRepartition.tsx'
import LastMessages from '@/components/Widgets/LastMessages.tsx'
import { Layout, Widget } from '@/interfaces/Widgets/Widgets.ts'
import PageMeta from '@/components/PageMeta/PageMeta'
import style from './Overview.module.css'
import '@/../node_modules/react-grid-layout/css/styles.css'
import '@/../node_modules/react-resizable/css/styles.css'

const ResponsiveGridLayout = WidthProvider(Responsive)

const WidgetTemplate: React.FC<{
  areWidgetsMovable: boolean
  children: React.ReactNode
}> = ({ areWidgetsMovable, children }) => (
  <div className={style.widgetContainer}>
    {areWidgetsMovable && (
      <div className={style.moveWidgetIcon}>
        <img src={MoveWidgetIcon} alt="move widget" style={{ width: '17px' }} />
      </div>
    )}
    <div className={style.widgetContent}>{children}</div>
  </div>
)

const Overview: React.FC = () => {
  const { t } = useTranslation()
  const [areWidgetsMovable, setAreWidgetsMovable] = useState(false)

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
        name: 'LastMessages',
        x: 0,
        y: 1,
        w: 2,
        h: 2,
        children: <LastMessages height={2} />
      }
    ]
  }

  const [widgets, setWidgets] = useState(layouts.lg)

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
          {!areWidgetsMovable && (
            <Button
              type="primary"
              onClick={() => setAreWidgetsMovable(!areWidgetsMovable)}
            >
              {t('components.button.edit_widgets_position')}
            </Button>
          )}
          {areWidgetsMovable && (
            <div className={style.editButtonsContainer}>
              <Button
                type="primary"
                danger
                onClick={() => setAreWidgetsMovable(!areWidgetsMovable)}
              >
                {t('components.button.cancel')}
              </Button>
              <Button
                type="primary"
                onClick={() => setAreWidgetsMovable(!areWidgetsMovable)}
              >
                {t('components.button.save')}
              </Button>
            </div>
          )}
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
            isDraggable={areWidgetsMovable}
            draggableHandle={`.${style.moveWidgetIcon}`}
            preventCollision
            compactType={null}
          >
            {widgets.map((widget: Widget) => (
              <div key={widget.i} data-grid={widget}>
                <WidgetTemplate areWidgetsMovable={areWidgetsMovable}>
                  {widget.children}
                </WidgetTemplate>
              </div>
            ))}
          </ResponsiveGridLayout>
        </div>
      </div>
    </>
  )
}

export default Overview
