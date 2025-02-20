import React, { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { CloseOutlined } from '@ant-design/icons'
import { Button, Form, Input, InputNumber, Modal, message } from 'antd'

import addIcon from '@/assets/icons/plus.png'
import { usePropertyId } from '@/context/propertyIdContext'
import GetRoomsByProperty from '@/services/api/Owner/Properties/Rooms/GetRoomsByProperty'
import CreateRoomByProperty from '@/services/api/Owner/Properties/Rooms/CreateRoomByProperty'
import GetFurnituresByRoom from '@/services/api/Owner/Properties/Rooms/Furnitures/GetFurnituresByRoom'
import CreateFurnitureByRoom from '@/services/api/Owner/Properties/Rooms/Furnitures/CreateFurnitureByRoom'
import DeleteFurnitureByRoom from '@/services/api/Owner/Properties/Rooms/Furnitures/DeleteFurnitureByRoom'
import DeleteRoomByPropertyById from '@/services/api/Owner/Properties/Rooms/DeleteRoomByPropertyById'
import style from './3InventoryTab.module.css'

const InventoryTab: React.FC = () => {
  const { t } = useTranslation()
  const [formAddRoom] = Form.useForm()
  const [formAddStuff] = Form.useForm()
  const id = usePropertyId()
  const [isModalAddRoomOpen, setIsModalAddRoomOpen] = useState(false)
  const [isModalAddStuffOpen, setIsModalAddStuffOpen] = useState(false)

  const [inventory, setInventory] = useState<
    {
      roomId: string
      roomName: string
      stuffs: { name: string; id: string | undefined }[]
    }[]
  >([])
  const [selectedRoomName, setSelectedRoomName] = useState<string | null>(null)

  useEffect(() => {
    const fetchInventory = async () => {
      if (!id) {
        message.error('Property ID is missing.')
        return
      }

      try {
        const rooms = await GetRoomsByProperty(id)
        const inventoryData = await Promise.all(
          rooms.map(async room => {
            try {
              const furnitures = await GetFurnituresByRoom(id, room.id)
              return {
                roomId: room.id,
                roomName: room.name,
                stuffs: furnitures.map(furniture => ({
                  name: furniture.name,
                  id: furniture.id
                }))
              }
            } catch (error) {
              console.error(
                `Error fetching furniture for room ${room.name}:`,
                error
              )
              return { roomId: room.id, roomName: room.name, stuffs: [] }
            }
          })
        )
        setInventory(inventoryData)
      } catch (error) {
        console.error('Error fetching rooms:', error)
        message.error('Failed to load rooms.')
      }
    }

    fetchInventory()
  }, [id])

  const showModal = (modal: string, roomName?: string) => {
    if (modal === 'addRoom') {
      setIsModalAddRoomOpen(true)
    } else if (modal === 'addStuff' && roomName) {
      setSelectedRoomName(roomName)
      setIsModalAddStuffOpen(true)
    }
  }

  const handleAddRoom = () => {
    formAddRoom
      .validateFields()
      .then(async () => {
        const { roomName } = formAddRoom.getFieldsValue()

        if (!id) {
          console.error('Property ID is missing')
          return
        }

        try {
          const newRoom = await CreateRoomByProperty(id, roomName)

          setInventory([
            ...inventory,
            { roomId: newRoom.id, roomName: newRoom.name, stuffs: [] }
          ])

          formAddRoom.resetFields()
          setIsModalAddRoomOpen(false)
        } catch (error) {
          console.error('Error creating room:', error)
        }
      })
      .catch(info => {
        console.error('Validate Failed:', info)
      })
  }

  const handleAddStuff = () => {
    formAddStuff
      .validateFields()
      .then(async () => {
        const { stuffName, itemQuantity } = formAddStuff.getFieldsValue()

        if (selectedRoomName && id) {
          const room = inventory.find(r => r.roomName === selectedRoomName)
          if (room) {
            try {
              const newFurniture = await CreateFurnitureByRoom(
                id,
                room.roomId,
                {
                  name: stuffName,
                  quantity: itemQuantity
                }
              )

              setInventory(
                inventory.map(r =>
                  r.roomName === selectedRoomName
                    ? {
                        ...r,
                        stuffs: [
                          ...r.stuffs,
                          {
                            name: newFurniture.name,
                            id: newFurniture.id,
                            quantity: newFurniture.quantity
                          }
                        ]
                      }
                    : r
                )
              )

              formAddStuff.resetFields()
              setSelectedRoomName(null)
              setIsModalAddStuffOpen(false)
            } catch (error) {
              console.error('Error creating furniture:', error)
              message.error('Failed to add item.')
            }
          }
        }
      })
      .catch(info => {
        console.error('Validate Failed:', info)
      })
  }

  const handleCancel = (modal: string) => {
    if (modal === 'addRoom') {
      formAddRoom.resetFields()
      setIsModalAddRoomOpen(false)
    } else if (modal === 'addStuff') {
      formAddStuff.resetFields()
      setSelectedRoomName(null)
      setIsModalAddStuffOpen(false)
    }
  }

  const removeRoom = (roomName: string, roomId: string) => {
    Modal.confirm({
      title: t(
        'pages.real_property_details.tabs.inventory.remove_room_confirmation'
      ),
      onOk: async () => {
        if (!id) {
          message.error('Property ID is missing.')
          return
        }

        try {
          await DeleteRoomByPropertyById(id, roomId)
          setInventory(inventory.filter(r => r.roomName !== roomName))
          message.success(
            t('pages.real_property_details.tabs.inventory.remove_room_success')
          )
        } catch (error) {
          console.error('Error deleting room:', error)
          message.error('Failed to delete room.')
        }
      }
    })
  }

  const removeStuff = (
    roomName: string,
    roomId: string,
    furnitureId: string
  ) => {
    Modal.confirm({
      title: t(
        'pages.real_property_details.tabs.inventory.remove_stuff_confirmation'
      ),
      onOk: async () => {
        if (!id) {
          message.error('Property ID is missing.')
          return
        }

        try {
          await DeleteFurnitureByRoom(id, roomId, furnitureId)
          setInventory(
            inventory.map(r =>
              r.roomName === roomName
                ? {
                    ...r,
                    stuffs: r.stuffs.filter(stuff => stuff.id !== furnitureId)
                  }
                : r
            )
          )
          message.success(
            t('pages.real_property_details.tabs.inventory.remove_stuff_success')
          )
        } catch (error) {
          console.error('Error deleting furniture:', error)
          message.error('Failed to delete item.')
        }
      }
    })
  }

  return (
    <div className={style.tabContent}>
      <div className={style.buttonAddContainer}>
        <Button type="primary" onClick={() => showModal('addRoom')}>
          {t('components.button.add_room')}
        </Button>
      </div>
      <Modal
        title={t(
          'pages.real_property_details.tabs.inventory.add_room_modal_title'
        )}
        open={isModalAddRoomOpen}
        onOk={handleAddRoom}
        onCancel={() => handleCancel('addRoom')}
      >
        <Form form={formAddRoom} layout="vertical" name="add_room_form">
          <Form.Item
            label={t('components.input.room_name.label')}
            name="roomName"
            rules={[
              { required: true, message: t('components.input.room_name.error') }
            ]}
          >
            <Input
              placeholder={t('components.input.room_name.placeholder')}
              maxLength={20}
              count={{
                show: true,
                max: 20
              }}
            />
          </Form.Item>
        </Form>
      </Modal>

      <div className={style.roomsContainer}>
        {inventory.map(room => (
          <div key={room.roomName} className={style.roomContainer}>
            <div className={style.roomHeader}>
              <span>{room.roomName}</span>
              <CloseOutlined
                onClick={() => removeRoom(room.roomName, room.roomId)}
              />
            </div>
            <div className={style.stuffsContainer}>
              {room?.stuffs?.map(stuff => (
                <div key={stuff.id} className={style.stuffCard}>
                  <span>{stuff.name}</span>
                  <CloseOutlined
                    onClick={() =>
                      removeStuff(room.roomName, room.roomId, stuff.id || '')
                    }
                    className={style.removeStuffIcon}
                    style={{ width: '20px', height: '20px' }}
                  />
                </div>
              ))}
              <div
                className={style.stuffCardAdd}
                onClick={() => showModal('addStuff', room.roomName)}
                onKeyDown={e => {
                  if (e.key === 'Enter') {
                    showModal('addStuff', room.roomName)
                  }
                }}
                role="button"
                tabIndex={0}
              >
                <img
                  src={addIcon}
                  alt="Add"
                  style={{ width: '20px', height: '20px' }}
                />
              </div>
            </div>
          </div>
        ))}
        <Modal
          title={t(
            'pages.real_property_details.tabs.inventory.add_stuff_modal_title'
          )}
          open={isModalAddStuffOpen}
          onOk={handleAddStuff}
          onCancel={() => handleCancel('addStuff')}
        >
          <Form form={formAddStuff} layout="vertical" name="add_stuff_form">
            <Form.Item
              label={t('components.input.stuff_name.label')}
              name="stuffName"
              rules={[
                {
                  required: true,
                  message: t('components.input.stuff_name.error')
                }
              ]}
            >
              <Input
                placeholder={t('components.input.stuff_name.placeholder')}
                maxLength={20}
                count={{
                  show: true,
                  max: 20
                }}
              />
            </Form.Item>
            <Form.Item
              label={t('components.input.item_quantity.label')}
              name="itemQuantity"
              rules={[
                {
                  required: true,
                  message: t('components.input.item_quantity.error')
                },
                {
                  type: 'number',
                  min: 1,
                  max: 1000,
                  message: t(
                    'components.input.item_quantity.validation_message'
                  )
                }
              ]}
            >
              <InputNumber
                type="number"
                placeholder={t('components.input.item_quantity.placeholder')}
                max={1000}
                min={1}
                style={{ width: '100%' }}
              />
            </Form.Item>
          </Form>
        </Modal>
      </div>
    </div>
  )
}

export default InventoryTab
