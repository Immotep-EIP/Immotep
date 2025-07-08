import React from 'react'
import { Responsive, WidthProvider } from 'react-grid-layout'
import { useTranslation } from 'react-i18next'

import { Typography, Popover, Space, Tag } from 'antd'
import { HomeOutlined, PlusOutlined } from '@ant-design/icons'

import { Button, Badge } from '@/components/common'

import { getRoomColor, isValidRoomType } from '@/utils/types/roomTypes'

import { Room } from '@/interfaces/Property/Inventory/Room/Room'
import { PlanViewProps } from '@/interfaces/Property/Inventory/Views/PlanView'

import style from './PlanView.module.css'

const ResponsiveGridLayout = WidthProvider(Responsive)
const { Text } = Typography

const PlanView: React.FC<PlanViewProps> = ({
  inventory,
  layouts,
  setLayouts,
  showModal,
  handleDeleteFurniture
}) => {
  const { t } = useTranslation()

  const getRoomClassName = (roomType: string) => {
    const baseClass = style.roomBlock
    if (roomType && isValidRoomType(roomType)) {
      return `${baseClass} ${style.customRoom}`
    }
    return baseClass
  }

  const renderRoomInPlan = (room: Room) => {
    const roomColor =
      room.type && isValidRoomType(room.type)
        ? getRoomColor(room.type)
        : '#1890ff'
    return (
      <div
        className={getRoomClassName(room.type)}
        style={{
          borderColor: roomColor,
          backgroundColor: `${roomColor}15`
        }}
        role="button"
        tabIndex={0}
        aria-label={t('components.inventory.plan.room_details', {
          room: room.name,
          count: room.furniture.length
        })}
      >
        <HomeOutlined
          style={{ fontSize: 24, marginBottom: 8, color: roomColor }}
          aria-hidden="true"
        />
        <Text strong>{room.name}</Text>
        <Badge
          count={room.furniture.length}
          style={{ backgroundColor: '#52c41a' }}
          aria-label={t('components.inventory.plan.items_count', {
            count: room.furniture.length
          })}
        />
      </div>
    )
  }

  const defaultLayouts = {
    lg: inventory.rooms.map((room, index) => ({
      i: room.id,
      x: (index % 3) * 4,
      y: Math.floor(index / 3) * 4,
      w: 4,
      h: 4,
      minW: 3,
      maxW: 6,
      minH: 3,
      maxH: 6,
      static: false
    }))
  }

  return (
    <main>
      <section aria-labelledby="plan-view-title">
        <h2 id="plan-view-title" className="sr-only">
          {t('components.inventory.plan.title')}
        </h2>

        <div className={style.planContainer}>
          <ResponsiveGridLayout
            className={style.planGrid}
            layouts={layouts.lg ? layouts : defaultLayouts}
            breakpoints={{ lg: 1200, md: 996, sm: 768, xs: 480, xxs: 0 }}
            cols={{ lg: 24, md: 20, sm: 16, xs: 12, xxs: 8 }}
            rowHeight={60}
            margin={[16, 16]}
            isResizable
            isDraggable
            onLayoutChange={layouts => {
              setLayouts(layouts)
            }}
            containerPadding={[24, 24]}
            compactType={null}
          >
            {inventory.rooms.map(room => {
              const content = (
                <div style={{ maxWidth: 300 }}>
                  <Text
                    strong
                    style={{ fontSize: 16, marginBottom: 8, display: 'block' }}
                  >
                    {room.name}
                  </Text>
                  <div
                    role="group"
                    aria-labelledby={`room-${room.id}-furniture-title`}
                  >
                    <h3
                      id={`room-${room.id}-furniture-title`}
                      className="sr-only"
                    >
                      {t('components.inventory.plan.furniture_in_room', {
                        room: room.name
                      })}
                    </h3>
                    <Space direction="vertical" style={{ width: '100%' }}>
                      {room.furniture.map(stuff => (
                        <Tag
                          key={stuff.id}
                          closable
                          onClose={() =>
                            handleDeleteFurniture(room.id, stuff.id)
                          }
                          style={{ margin: '2px 0' }}
                          aria-label={t(
                            'components.inventory.plan.delete_item',
                            {
                              item: t(
                                `components.inventory.furniture.${stuff.name}`,
                                {
                                  defaultValue: stuff.name
                                }
                              ),
                              quantity: stuff.quantity
                            }
                          )}
                        >
                          {t(`components.inventory.furniture.${stuff.name}`, {
                            defaultValue: stuff.name
                          })}
                          ({stuff.quantity})
                        </Tag>
                      ))}
                      <Button
                        type="dashed"
                        icon={<PlusOutlined aria-hidden="true" />}
                        size="small"
                        onClick={() => showModal('addStuff', room.id)}
                        block
                        aria-label={t(
                          'components.inventory.plan.add_item_to_room',
                          { room: room.name }
                        )}
                      >
                        {t('components.button.add_item')}
                      </Button>
                    </Space>
                  </div>
                </div>
              )

              return (
                <div key={room.id}>
                  <Popover
                    content={content}
                    title={null}
                    trigger="click"
                    placement="right"
                    arrowPointAtCenter
                  >
                    {renderRoomInPlan(room)}
                  </Popover>
                </div>
              )
            })}
          </ResponsiveGridLayout>
          <div
            className={style.addRoomBlock}
            aria-label={t('components.inventory.plan.add_room')}
            onClick={() => showModal('addRoom')}
            onKeyDown={e => {
              if (e.key === 'Enter' || e.key === ' ') {
                showModal('addRoom')
              }
            }}
            role="button"
            tabIndex={0}
          >
            <PlusOutlined style={{ fontSize: 24 }} />
            <Text>{t('components.button.add_room')}</Text>
          </div>
        </div>
      </section>
    </main>
  )
}

export default PlanView
