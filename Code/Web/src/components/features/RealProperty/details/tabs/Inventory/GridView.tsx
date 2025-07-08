import React from 'react'
import { useTranslation } from 'react-i18next'

import { Row, Col, Space, Typography, Tag, Tooltip } from 'antd'
import {
  HomeOutlined,
  PlusOutlined,
  EditOutlined,
  DeleteOutlined
} from '@ant-design/icons'

import { Button, Card, Badge } from '@/components/common'

import { getRoomColor, isValidRoomType } from '@/utils/types/roomTypes'

import { Room } from '@/interfaces/Property/Inventory/Room/Room'
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
    <section aria-labelledby="inventory-grid-title">
      <h2 id="inventory-grid-title" className="sr-only">
        {t('components.inventory.view_mode.grid')}
      </h2>

      <Row
        gutter={[24, 24]}
        role="grid"
        aria-label={t('components.inventory.view_mode.grid')}
      >
        {inventory.rooms.map((room: Room) => (
          <Col xs={24} sm={12} md={8} key={room.id} role="gridcell">
            <Card
              className="room-card"
              role="article"
              aria-labelledby={`room-title-${room.id}`}
              title={
                <Space>
                  <HomeOutlined
                    style={{
                      fontSize: '18px',
                      color: getTypeColor(room.type)
                    }}
                    aria-hidden="true"
                  />
                  <Text strong id={`room-title-${room.id}`}>
                    {room.name}
                  </Text>
                  <Badge
                    count={room.furniture.length}
                    style={{
                      backgroundColor: '#52c41a',
                      boxShadow: '0 0 0 2px #fff'
                    }}
                    aria-label={`${room.furniture.length} ${t('components.inventory.furniture.items')}`}
                  />
                </Space>
              }
              extra={
                <Space
                  role="toolbar"
                  aria-label={`${t('components.button.edit')} ${room.name}`}
                >
                  <Tooltip title={t('components.tooltip.edit')}>
                    <Button
                      type="text"
                      icon={<EditOutlined aria-hidden="true" />}
                      size="small"
                      className="action-button"
                      aria-label={`${t('components.tooltip.edit')} ${room.name}`}
                    />
                  </Tooltip>
                  <Tooltip title={t('components.tooltip.delete')}>
                    <Button
                      type="text"
                      danger
                      icon={<DeleteOutlined aria-hidden="true" />}
                      size="small"
                      onClick={() => handleDeleteRoom(room.id)}
                      className="action-button"
                      aria-label={`${t('components.tooltip.delete')} ${room.name}`}
                    />
                  </Tooltip>
                </Space>
              }
              actions={[
                <Button
                  key="add"
                  type="link"
                  icon={<PlusOutlined aria-hidden="true" />}
                  onClick={() => showModal('addStuff', room.id)}
                  style={{ color: getTypeColor(room.type) }}
                  aria-label={`${t('components.button.add_item')} - ${room.name}`}
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
              <section aria-labelledby={`furniture-list-${room.id}`}>
                <h3 id={`furniture-list-${room.id}`} className="sr-only">
                  {t(
                    'pages.real_property_details.tabs.inventory.list_object_name'
                  )}
                </h3>

                <div
                  role="list"
                  style={{ width: '100%' }}
                  className="furniture-list"
                >
                  {room.furniture.map(stuff => (
                    <div
                      key={stuff.id}
                      className="furniture-item"
                      role="listitem"
                    >
                      <Row align="middle" justify="space-between">
                        <Col flex="auto">
                          <Text id={`furniture-name-${stuff.id}`}>
                            {t(`components.inventory.furniture.${stuff.name}`, {
                              defaultValue: stuff.name
                            })}
                          </Text>
                        </Col>
                        <Col flex="100px" style={{ textAlign: 'right' }}>
                          <Tag
                            color="blue"
                            style={{
                              margin: 0,
                              minWidth: '60px',
                              textAlign: 'center'
                            }}
                            aria-label={`${t('components.inventory.grid.quantity')}: ${stuff.quantity}`}
                          >
                            {stuff.quantity}
                          </Tag>
                        </Col>
                        <Col flex="40px" style={{ textAlign: 'right' }}>
                          <Button
                            type="text"
                            danger
                            icon={<DeleteOutlined aria-hidden="true" />}
                            size="small"
                            onClick={() =>
                              handleDeleteFurniture(room.id, stuff.id)
                            }
                            className="action-button"
                            aria-label={`${t('components.button.delete_item')} ${t(`components.inventory.furniture.${stuff.name}`, { defaultValue: stuff.name })}`}
                            aria-describedby={`furniture-name-${stuff.id}`}
                          />
                        </Col>
                      </Row>
                    </div>
                  ))}
                </div>
              </section>
            </Card>
          </Col>
        ))}
      </Row>
    </section>
  )
}

export default GridView
