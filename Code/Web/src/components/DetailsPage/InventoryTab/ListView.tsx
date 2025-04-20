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
  DeleteOutlined,
  EditOutlined
} from '@ant-design/icons'
import { useTranslation } from 'react-i18next'
import { getRoomColor, isValidRoomType } from '@/utils/types/roomTypes'
import { ListViewProps } from '@/interfaces/Property/Inventory/Views/ListView'
import './ListView.css'

const { Text } = Typography

const ListView: React.FC<ListViewProps> = ({
  inventory,
  showModal,
  handleDeleteRoom,
  handleDeleteFurniture
}) => {
  const { t } = useTranslation()

  return (
    <Card className="list-container">
      {inventory.rooms.map(room => {
        const roomColor =
          room.type && isValidRoomType(room.type)
            ? getRoomColor(room.type)
            : '#1890ff'
        return (
          <Card.Grid
            key={room.id}
            className="room-item"
            style={{
              width: '100%',
              padding: '20px',
              marginBottom: '16px',
              borderRadius: '8px',
              backgroundColor: '#fafafa',
              borderLeft: `4px solid ${roomColor}`,
              transition: 'all 0.3s ease'
            }}
          >
            <Row align="middle" justify="space-between">
              <Col>
                <Space size="middle">
                  <HomeOutlined
                    style={{ fontSize: '18px', color: roomColor }}
                  />
                  <Text strong style={{ fontSize: '16px' }}>
                    {room.name}
                  </Text>
                  <Badge
                    count={room.furniture.length}
                    style={{
                      backgroundColor: '#52c41a',
                      boxShadow: '0 0 0 2px #fff'
                    }}
                  />
                </Space>
              </Col>
              <Col>
                <Space size="middle">
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
                  <Button
                    type="primary"
                    ghost
                    icon={<PlusOutlined />}
                    size="small"
                    onClick={() => showModal('addStuff', room.id)}
                    className="add-button"
                    style={{ color: roomColor, borderColor: roomColor }}
                  >
                    {t('components.button.add_item')}
                  </Button>
                </Space>
              </Col>
            </Row>
            <div className="furniture-list">
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
            </div>
          </Card.Grid>
        )
      })}
    </Card>
  )
}

export default ListView
