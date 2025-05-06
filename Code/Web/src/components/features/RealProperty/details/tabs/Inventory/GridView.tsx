import React from 'react'
import {
  Card,
  Row,
  Col,
  Space,
  Typography,
  Badge,
  Button,
  Tag,
  Tooltip
} from 'antd'
import {
  HomeOutlined,
  PlusOutlined,
  EditOutlined,
  DeleteOutlined
} from '@ant-design/icons'
import { useTranslation } from 'react-i18next'

import { Room } from '@/interfaces/Property/Inventory/Room/Room'
import { getRoomColor, isValidRoomType } from '@/utils/types/roomTypes'
import { GridViewProps } from '@/interfaces/Property/Inventory/Views/GridView'
import './GridView.css'

const { Text } = Typography

const GridView: React.FC<GridViewProps> = ({
  inventory,
  showModal,
  handleDeleteRoom,
  handleDeleteFurniture
}) => {
  const { t } = useTranslation()

  const getTypeColor = (type: string | undefined) => {
    if (type && isValidRoomType(type)) {
      return getRoomColor(type)
    }
    return '#1890ff'
  }

  return (
    <Row gutter={[24, 24]}>
      {inventory.rooms.map((room: Room) => (
        <Col xs={24} sm={12} md={8} key={room.id}>
          <Card
            className="room-card"
            title={
              <Space>
                <HomeOutlined
                  style={{
                    fontSize: '18px',
                    color: getTypeColor(room.type)
                  }}
                />
                <Text strong>{room.name}</Text>
                <Badge
                  count={room.furniture.length}
                  style={{
                    backgroundColor: '#52c41a',
                    boxShadow: '0 0 0 2px #fff'
                  }}
                />
              </Space>
            }
            extra={
              <Space>
                <Tooltip title={t('components.tooltip.edit')}>
                  <Button
                    type="text"
                    icon={<EditOutlined />}
                    size="small"
                    className="action-button"
                  />
                </Tooltip>
                <Tooltip title={t('components.tooltip.delete')}>
                  <Button
                    type="text"
                    danger
                    icon={<DeleteOutlined />}
                    size="small"
                    onClick={() => handleDeleteRoom(room.id)}
                    className="action-button"
                  />
                </Tooltip>
              </Space>
            }
            actions={[
              <Button
                key="add"
                type="link"
                icon={<PlusOutlined />}
                onClick={() => showModal('addStuff', room.id)}
                style={{ color: getTypeColor(room.type) }}
              >
                {t('components.button.add_item')}
              </Button>
            ]}
            style={{
              borderRadius: '8px',
              boxShadow: '0 2px 8px rgba(0,0,0,0.08)',
              transition: 'all 0.3s ease',
              borderTop: `3px solid ${getTypeColor(room.type)}`
            }}
          >
            <Space direction="vertical" style={{ width: '100%' }} size="middle">
              {room.furniture.map(stuff => (
                <div key={stuff.id} className="furniture-item">
                  <Row align="middle" justify="space-between">
                    <Col flex="auto">
                      <Text>{stuff.name}</Text>
                    </Col>
                    <Col flex="100px" style={{ textAlign: 'right' }}>
                      <Tag
                        color="blue"
                        style={{
                          margin: 0,
                          minWidth: '60px',
                          textAlign: 'center'
                        }}
                      >
                        {stuff.quantity}
                      </Tag>
                    </Col>
                    <Col flex="40px" style={{ textAlign: 'right' }}>
                      <Button
                        type="text"
                        danger
                        icon={<DeleteOutlined />}
                        size="small"
                        onClick={() => handleDeleteFurniture(room.id, stuff.id)}
                        className="action-button"
                      />
                    </Col>
                  </Row>
                </div>
              ))}
            </Space>
          </Card>
        </Col>
      ))}
    </Row>
  )
}

export default GridView
