import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Space, Table, Tag, Modal } from 'antd'
import type { TableProps } from 'antd'
import { EyeOutlined } from '@ant-design/icons'
import style from './3DamageTab.module.css'

interface Picture {
  id: string
  url: string
}

interface DataType {
  key: number
  date: string
  comment: string
  intervention_date: string
  room: string
  priority: string
  pictures: Picture[]
}

const data: DataType[] = [
  {
    key: 1,
    date: '2021-07-01',
    comment: 'Broken window',
    intervention_date: '2021-07-02',
    room: 'Living room',
    priority: 'High',
    pictures: [
      {
        id: '1',
        url: 'https://comarquesnord.cat/wp-content/uploads/placeholder-image-150x150.png'
      }
    ]
  },
  {
    key: 2,
    date: '2021-07-02',
    comment: 'Broken door',
    intervention_date: '2021-07-03',
    room: 'Kitchen',
    priority: 'Low',
    pictures: [
      {
        id: '2',
        url: 'https://comarquesnord.cat/wp-content/uploads/placeholder-image-150x150.png'
      },
      {
        id: '3',
        url: 'https://comarquesnord.cat/wp-content/uploads/placeholder-image-150x150.png'
      },
      {
        id: '4',
        url: 'https://comarquesnord.cat/wp-content/uploads/placeholder-image-150x150.png'
      }
    ]
  },
  {
    key: 3,
    date: '2021-07-03',
    comment: 'Broken wall',
    intervention_date: '2021-07-04',
    room: 'Bedroom',
    priority: 'Medium',
    pictures: [
      {
        id: '5',
        url: 'https://comarquesnord.cat/wp-content/uploads/placeholder-image-150x150.png'
      },
      {
        id: '6',
        url: 'https://comarquesnord.cat/wp-content/uploads/placeholder-image-150x150.png'
      }
    ]
  }
]

const DamageTab: React.FC = () => {
  const { t } = useTranslation()

  const [isModalVisible, setIsModalVisible] = useState(false)
  const [selectedImage, setSelectedImage] = useState<string | null>(null)

  const handleImageClick = (image: string) => {
    setSelectedImage(image)
    setIsModalVisible(true)
  }

  const handleCloseModal = () => {
    setIsModalVisible(false)
    setSelectedImage(null)
  }

  const getPriorityColor = (priority: string): string => {
    if (priority === 'High') return 'red'
    if (priority === 'Medium') return 'yellow'
    return 'green'
  }

  const columns: TableProps<DataType>['columns'] = [
    {
      title: t('pages.real_property_details.tabs.damage.table.date'),
      dataIndex: 'date',
      key: 'date',
      render: text => text
    },
    {
      title: t('pages.real_property_details.tabs.damage.table.comment'),
      dataIndex: 'comment',
      key: 'comment',
      render: text => text
    },
    {
      title: t(
        'pages.real_property_details.tabs.damage.table.intervention_date'
      ),
      dataIndex: 'intervention_date',
      key: 'intervention_date',
      render: text => text
    },
    {
      title: t('pages.real_property_details.tabs.damage.table.room'),
      dataIndex: 'room',
      key: 'room',
      render: text => text
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
      render: pictures => (
        <Space>
          {pictures.map((picture: Picture) => (
            <div
              key={picture.id}
              style={{ position: 'relative', width: '45px', height: '45px' }}
            >
              <img
                src={picture.url}
                alt="damage"
                style={{
                  width: '100%',
                  height: '100%',
                  objectFit: 'cover',
                  borderRadius: '5px',
                  cursor: 'pointer'
                }}
              />
              <div
                className={style.hoverOverlay}
                role="button"
                tabIndex={0}
                onClick={() => handleImageClick(picture.url)}
                onKeyPress={e =>
                  e.key === 'Enter' && handleImageClick(picture.url)
                }
              >
                <EyeOutlined style={{ color: 'white', fontSize: '20px' }} />
              </div>
            </div>
          ))}
        </Space>
      )
    }
  ]

  return (
    <div className={style.tabContent}>
      <div style={{ width: '100%' }}>
        <Table<DataType>
          columns={columns}
          dataSource={data}
          pagination={false}
        />
      </div>
      <Modal
        title={t('pages.real_property_details.tabs.damage.modal.title')}
        open={isModalVisible}
        footer={null}
        onCancel={handleCloseModal}
      >
        <img
          src={selectedImage || ''}
          alt="Selected"
          style={{ width: '100%', borderRadius: '5px' }}
        />
      </Modal>
    </div>
  )
}

export default DamageTab
