import React from 'react'
import { useTranslation } from 'react-i18next'

import { Empty, Spin, Table, Typography } from 'antd'
import type { TableProps } from 'antd'

import { usePropertyId } from '@/context/propertyIdContext'
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
  const propertyId = usePropertyId()
  const { goToDamageDetails } = useNavigation()
  const { damages, loading, error } = useDamages(propertyId || '', status || '')

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
      <div className={style.loadingContainer}>
        <Spin size="large" />
      </div>
    )
  }

  if (status === 'available') {
    return (
      <div className={style.tabContentEmpty}>
        <Empty
          image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
          imageStyle={{ height: 60 }}
          description={
            <Typography.Text>
              {t('pages.real_property.error.no_tenant_linked')}
            </Typography.Text>
          }
        />
      </div>
    )
  }
  if (error) return <div>{error}</div>

  return (
    <div className={style.tabContent}>
      <Table<DataType>
        columns={columns}
        dataSource={transformedData}
        pagination={false}
        bordered
        style={{ width: '100%', cursor: 'pointer' }}
        onRow={record => ({
          onClick: event => {
            if (record.key) {
              if (!propertyId) {
                event.stopPropagation()
                return
              }
              goToDamageDetails(propertyId, record.key)
            }
          }
        })}
      />
    </div>
  )
}

export default DamageTab
