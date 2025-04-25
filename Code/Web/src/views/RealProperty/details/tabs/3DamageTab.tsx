import React from 'react'
import { useTranslation } from 'react-i18next'
import { Empty, Spin, Table, Tag, Typography } from 'antd'
import type { TableProps } from 'antd'
import style from './3DamageTab.module.css'
import useDamages from '@/hooks/Property/useDamages'
import { usePropertyId } from '@/context/propertyIdContext'

interface DataType {
  key: string
  date: string
  comment: string
  intervention_date: string
  room: string
  priority: string
  pictures: number
}

interface DamageTabProps {
  status?: string
}

const DamageTab: React.FC<DamageTabProps> = ({ status }) => {
  const { t } = useTranslation()
  const propertyId = usePropertyId()
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
      pictures: item.pictures.length
    })) || []

  const getPriorityColor = (priority: string): string => {
    switch (priority) {
      case 'High':
        return 'red'
      case 'Medium':
        return 'yellow'
      default:
        return 'green'
    }
  }

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
        <Tag color={getPriorityColor(text)}>
          {t(
            `pages.real_property_details.tabs.damage.priority.${text.toLowerCase()}`
          )}
        </Tag>
      )
    },
    {
      title: t('pages.real_property_details.tabs.damage.table.pictures'),
      dataIndex: 'pictures',
      key: 'pictures',
      render: text => (
        <div>
          {text}{' '}
          {t(
            'pages.real_property_details.tabs.damage.table.available_pictures'
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
          styles={{ image: { height: 60 } }}
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
        style={{ width: '100%' }}
        // onRow={(record, rowIndex) => {
        //   return {
        //     onClick: event => {
        //       console.log('Row clicked:', record)
        //     }
        //   }
        // }}
      />
    </div>
  )
}

export default DamageTab
