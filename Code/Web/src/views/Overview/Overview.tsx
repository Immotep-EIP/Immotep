import React, { useState, useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { Responsive, WidthProvider } from 'react-grid-layout'

import { Spin } from 'antd'

import useDashboard from '@/hooks/Dashboard/useDashboard'
import PageTitle from '@/components/ui/PageText/Title.tsx'
import PropertiesNumber from '@/components/features/Overview/Widgets/PropertiesNumber/PropertiesNumber'
import PropertiesRepartition from '@/components/features/Overview/Widgets/PropertiesRepartition/PropertiesRepartition'
import PageMeta from '@/components/ui/PageMeta/PageMeta'
import Reminders from '@/components/features/Overview/Widgets/Reminders/Reminders'
import OpenDamages from '@/components/features/Overview/Widgets/OpenDamages/OpenDamages'
import DamagesRepartition from '@/components/features/Overview/Widgets/DamagesRepartition/DamagesRepartition'

import { Layout, Widget } from '@/interfaces/Widgets/Widgets.ts'

import '@/../node_modules/react-grid-layout/css/styles.css'
import '@/../node_modules/react-resizable/css/styles.css'
import style from './Overview.module.css'

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
  const {
    reminders,
    properties,
    open_damages: openDamages,
    loading,
    error
  } = useDashboard()

  const initialLayouts: { lg: Widget[] } = {
    lg: [
      {
        i: '1',
        name: 'PropertiesNumber',
        x: 0,
        y: 0,
        w: 2,
        h: 1,
        children: (
          <PropertiesNumber
            properties={properties}
            loading={loading}
            error={error}
            height={1}
          />
        )
      },
      {
        i: '2',
        name: 'PropertiesRepartition',
        x: 2,
        y: 0,
        w: 2,
        h: 1,
        children: (
          <PropertiesRepartition
            properties={properties}
            loading={loading}
            error={error}
            height={1}
          />
        )
      },
      // {
      //   i: '3',
      //   name: 'LastMessages',
      //   x: 0,
      //   y: 1,
      //   w: 4,
      //   h: 2,
      //   children: <LastMessages height={2} />
      // },
      {
        i: '4',
        name: 'Reminders',
        x: 7,
        y: 0,
        w: 3,
        h: 3,
        children: (
          <Reminders
            reminders={reminders}
            loading={loading}
            error={error}
            height={3}
          />
        )
      },
      {
        i: '5',
        name: 'OpenDamages',
        x: 0,
        y: 0,
        w: 3,
        h: 3,
        children: (
          <OpenDamages
            openDamages={openDamages}
            loading={loading}
            error={error}
            height={3}
          />
        )
      },
      {
        i: '6',
        name: 'DamagesRepartition',
        x: 6,
        y: 0,
        w: 2,
        h: 1,
        children: (
          <DamagesRepartition
            openDamages={openDamages}
            loading={loading}
            error={error}
            height={1}
          />
        )
      }
    ]
  }

  const [widgets, setWidgets] = useState(initialLayouts.lg)

  useEffect(() => {
    if (properties) {
      setWidgets(prevWidgets =>
        prevWidgets.map(widget => {
          if (widget.name === 'PropertiesNumber') {
            return {
              ...widget,
              children: (
                <PropertiesNumber
                  properties={properties}
                  loading={loading}
                  error={error}
                  height={widget.h}
                />
              )
            }
          }
          if (widget.name === 'PropertiesRepartition') {
            return {
              ...widget,
              children: (
                <PropertiesRepartition
                  properties={properties}
                  loading={loading}
                  error={error}
                  height={widget.h}
                />
              )
            }
          }
          if (widget.name === 'Reminders') {
            return {
              ...widget,
              children: (
                <Reminders
                  reminders={reminders}
                  loading={loading}
                  error={error}
                  height={widget.h}
                />
              )
            }
          }
          if (widget.name === 'OpenDamages') {
            return {
              ...widget,
              children: (
                <OpenDamages
                  openDamages={openDamages}
                  loading={loading}
                  error={error}
                  height={widget.h}
                />
              )
            }
          }
          if (widget.name === 'DamagesRepartition') {
            return {
              ...widget,
              children: (
                <DamagesRepartition
                  openDamages={openDamages}
                  loading={loading}
                  error={error}
                  height={widget.h}
                />
              )
            }
          }
          return widget
        })
      )
    }
  }, [properties, loading, error, reminders])

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
      <main
        className={style.pageContainer}
        aria-labelledby="overview-page-title"
      >
        <header className={style.pageHeader}>
          <PageTitle
            title={t('pages.overview.title')}
            size="title"
            id="overview-page-title"
          />
        </header>

        {loading ? (
          <section
            role="status"
            aria-live="polite"
            aria-labelledby="loading-title"
          >
            <h2 id="loading-title" className="sr-only">
              {t('pages.overview.loading_title')}
            </h2>
            <Spin
              size="large"
              style={{
                position: 'absolute',
                top: '50%',
                left: '50%',
                transform: 'translate(-50%, -50%)'
              }}
            />
          </section>
        ) : (
          <section
            className={style.contentContainer}
            aria-labelledby="dashboard-widgets-title"
          >
            <h2 id="dashboard-widgets-title" className="sr-only">
              {t('pages.overview.widgets_title')}
            </h2>
            <ResponsiveGridLayout
              className={style.gridLayout}
              layouts={{ lg: widgets }}
              breakpoints={{ lg: 768, md: 768, sm: 768, xs: 480, xxs: 0 }}
              cols={{ lg: 12, md: 10, sm: 6, xs: 4, xxs: 2 }}
              rowHeight={120}
              isResizable={false}
              onResize={handleLayoutChange}
              draggableHandle={`.${style.moveWidgetIcon}`}
              preventCollision
              compactType={null}
            >
              {widgets.map((widget: Widget) => (
                <article
                  key={widget.i}
                  data-grid={widget}
                  aria-labelledby={`widget-${widget.i}-title`}
                >
                  <h3 id={`widget-${widget.i}-title`} className="sr-only">
                    {t(`pages.overview.widgets.${widget.name.toLowerCase()}`)}
                  </h3>
                  <WidgetTemplate>{widget.children}</WidgetTemplate>
                </article>
              ))}
            </ResponsiveGridLayout>
          </section>
        )}
      </main>
    </>
  )
}

export default Overview
