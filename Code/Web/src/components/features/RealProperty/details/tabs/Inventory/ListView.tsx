import React from 'react'
import { useTranslation } from 'react-i18next'

import { Row, Col, Space, Typography, Tag, Tooltip } from 'antd'
import {
  HomeOutlined,
  PlusOutlined,
  DeleteOutlined,
  EditOutlined
} from '@ant-design/icons'

import { Button, Card, Badge } from '@/components/common'

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
    <main>
      <section aria-labelledby="inventory-list-title">
        <h2 id="inventory-list-title" className="sr-only">
          {t('components.inventory.list.title')}
        </h2>

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
                role="article"
                aria-labelledby={`room-${room.id}-title`}
              >
                <Row align="middle" justify="space-between">
                  <Col>
                    <Space size="middle">
                      <HomeOutlined
                        style={{ fontSize: '18px', color: roomColor }}
                        aria-hidden="true"
                      />
                      <Text
                        strong
                        style={{ fontSize: '16px' }}
                        id={`room-${room.id}-title`}
                      >
                        {room.name}
                      </Text>
                      <Badge
                        count={room.furniture.length}
                        style={{
                          backgroundColor: '#52c41a',
                          boxShadow: '0 0 0 2px #fff'
                        }}
                        aria-label={t('components.inventory.list.items_count', {
                          count: room.furniture.length
                        })}
                      />
                    </Space>
                  </Col>
                  <Col>
                    <div
                      role="toolbar"
                      aria-label={t('components.inventory.list.room_actions', {
                        room: room.name
                      })}
                    >
                      <Space size="middle">
                        <Tooltip title={t('components.tooltip.edit')}>
                          <Button
                            type="text"
                            icon={<EditOutlined aria-hidden="true" />}
                            size="small"
                            className="action-button"
                            aria-label={t(
                              'components.inventory.list.edit_room',
                              { room: room.name }
                            )}
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
                            aria-label={t(
                              'components.inventory.list.delete_room',
                              { room: room.name }
                            )}
                          />
                        </Tooltip>
                        <Button
                          type="primary"
                          ghost
                          icon={<PlusOutlined aria-hidden="true" />}
                          size="small"
                          onClick={() => showModal('addStuff', room.id)}
                          className="add-button"
                          style={{ color: roomColor, borderColor: roomColor }}
                          aria-label={t(
                            'components.inventory.list.add_item_to_room',
                            { room: room.name }
                          )}
                        >
                          {t('components.button.add_item')}
                        </Button>
                      </Space>
                    </div>
                  </Col>
                </Row>
                <section
                  className="furniture-list"
                  aria-labelledby={`furniture-${room.id}-title`}
                >
                  <h3 id={`furniture-${room.id}-title`} className="sr-only">
                    {t('components.inventory.list.furniture_in_room', {
                      room: room.name
                    })}
                  </h3>
                  <ul style={{ listStyle: 'none', padding: 0, margin: 0 }}>
                    {room.furniture.map(stuff => (
                      <li key={stuff.id} className="furniture-item">
                        <Row align="middle" justify="space-between">
                          <Col flex="auto">
                            <Text>
                              {t(
                                `components.inventory.furniture.${stuff.name}`,
                                {
                                  defaultValue: stuff.name
                                }
                              )}
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
                              aria-label={t(
                                'components.inventory.list.quantity',
                                { quantity: stuff.quantity }
                              )}
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
                              aria-label={t(
                                'components.inventory.list.delete_item',
                                {
                                  item: t(
                                    `components.inventory.furniture.${stuff.name}`,
                                    {
                                      defaultValue: stuff.name
                                    }
                                  )
                                }
                              )}
                            />
                          </Col>
                        </Row>
                      </li>
                    ))}
                  </ul>
                </section>
              </Card.Grid>
            )
          })}
        </Card>
      </section>
    </main>
  )
}

export default ListView
