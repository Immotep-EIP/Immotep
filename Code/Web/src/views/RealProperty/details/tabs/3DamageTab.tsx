import React from 'react'
import { useTranslation } from 'react-i18next'

import { Spin, Table, Typography } from 'antd'
import type { TableProps } from 'antd'

import { Empty } from '@/components/common'
import { usePropertyContext } from '@/context/propertyContext'
import useDamages from '@/hooks/Property/useDamages'
import useNavigation from '@/hooks/Navigation/useNavigation'
import StatusTag from '@/components/common/Tag/StatusTag'

import style from './3DamageTab.module.css'

interface DataType {
  key: string
  date: string
  comment: string
  intervention_date: string
  room: string
  priority: string
  pictures: string[]
}

interface DamageTabProps {
  status?: string
}

const DamageTab: React.FC<DamageTabProps> = ({ status }) => {
  const { t } = useTranslation()
  const { property, selectedLease } = usePropertyContext()
  const { goToDamageDetails } = useNavigation()
  const { damages, loading, error } = useDamages(
    property?.id || '',
    status || ''
  )

  const transformedData: DataType[] =
    damages?.map(item => ({
      key: String(item.id),
      date: new Date(item.created_at).toISOString().split('T')[0],
      comment:
        item.comment || t('pages.real_property_details.tabs.damage.no_comment'),
      intervention_date: item.fix_planned_at
        ? new Date(item.fix_planned_at).toISOString().split('T')[0]
        : '-',
      room:
        item.room_name ||
        t('pages.real_property_details.tabs.damage.unknown_room'),
      priority: item.priority.charAt(0).toUpperCase() + item.priority.slice(1),
      pictures: item?.pictures?.map(picture => picture)
    })) || []

  const columns: TableProps<DataType>['columns'] = [
    {
      title: t('pages.real_property_details.tabs.damage.table.date'),
      dataIndex: 'date',
      key: 'date'
    },
    {
      title: t('pages.real_property_details.tabs.damage.table.comment'),
      dataIndex: 'comment',
      key: 'comment'
    },
    {
      title: t(
        'pages.real_property_details.tabs.damage.table.intervention_date'
      ),
      dataIndex: 'intervention_date',
      key: 'intervention_date'
    },
    {
      title: t('pages.real_property_details.tabs.damage.table.room'),
      dataIndex: 'room',
      key: 'room'
    },
    {
      title: t('pages.real_property_details.tabs.damage.table.priority'),
      dataIndex: 'priority',
      key: 'priority',
      render: text => (
        <StatusTag
          value={text}
          colorMap={{
            urgent: 'red',
            high: 'red',
            medium: 'yellow',
            low: 'green'
          }}
        />
      )
    },
    {
      title: t('pages.real_property_details.tabs.damage.table.pictures'),
      dataIndex: 'pictures',
      key: 'pictures',
      render: record => (
        <div className={style.imagesContainer}>
          {!record || record.length === 0 ? (
            <Typography.Text>
              {t('pages.real_property_details.tabs.damage.no_pictures')}
            </Typography.Text>
          ) : (
            <Typography.Text>
              {record?.length}{' '}
              {t('pages.real_property_details.tabs.damage.pictures')}
            </Typography.Text>
          )}
        </div>
      )
    }
  ]

  if (loading) {
    return (
      <section
        className={style.loadingContainer}
        role="status"
        aria-live="polite"
        aria-labelledby="damage-loading-title"
      >
        <h2 id="damage-loading-title" className="sr-only">
          {t('pages.real_property_details.tabs.damage.loading_title')}
        </h2>
        <Spin size="large" />
      </section>
    )
  }

  if (
    (status === 'available' && !selectedLease) ||
    (status === 'invite sent' && !selectedLease)
  ) {
    return (
      <section
        className={style.tabContentEmpty}
        role="status"
        aria-live="polite"
        aria-labelledby="damage-empty-title"
      >
        <h2 id="damage-empty-title" className="sr-only">
          {t('pages.real_property_details.tabs.damage.empty_title')}
        </h2>
        <Empty description={t('pages.real_property.error.no_tenant_linked')} />
      </section>
    )
  }

  if (error) {
    return (
      <section
        role="alert"
        aria-live="assertive"
        aria-labelledby="damage-error-title"
      >
        <h2 id="damage-error-title" className="sr-only">
          {t('pages.real_property_details.tabs.damage.error_title')}
        </h2>
        <div>{error}</div>
      </section>
    )
  }

  return (
    <main className={style.tabContent} aria-labelledby="damage-tab-title">
      <h2 id="damage-tab-title" className="sr-only">
        {t('pages.real_property_details.tabs.damage.tab_title')}
      </h2>
      <h3 id="damage-table-title" className="sr-only">
        {t('pages.real_property_details.tabs.damage.table_title')}
      </h3>
      <Table<DataType>
        columns={columns}
        dataSource={transformedData}
        pagination={false}
        bordered
        style={{ width: '100%' }}
        onRow={record => ({
          onClick: event => {
            if (record.key) {
              if (!property?.id) {
                event.stopPropagation()
                return
              }
              goToDamageDetails(property.id, record.key)
            }
          },
          onKeyDown: event => {
            if ((event.key === 'Enter' || event.key === ' ') && record.key) {
              event.preventDefault()
              if (!property?.id) {
                return
              }
              goToDamageDetails(property.id, record.key)
            }
          },
          tabIndex: 0,
          role: 'button',
          'aria-label': t('pages.real_property_details.tabs.damage.row_aria', {
            date: record.date,
            room: record.room,
            priority: record.priority
          })
        })}
        aria-label={t('pages.real_property_details.tabs.damage.table_aria')}
      />
    </main>
  )
}

export default DamageTab
