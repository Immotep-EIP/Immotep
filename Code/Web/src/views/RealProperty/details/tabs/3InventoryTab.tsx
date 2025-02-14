import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { CloseOutlined } from '@ant-design/icons'
import {
  Button,
  Empty,
  Form,
  Input,
  InputNumber,
  Modal,
  Typography
} from 'antd'
import addIcon from '@/assets/icons/plus.png'
import { usePropertyId } from '@/context/propertyIdContext'
import useInventory from '@/hooks/useInventory/useInventory'
import CardInventoryLoader from '@/components/Loader/CardInventoryLoader'
import style from './3InventoryTab.module.css'

const InventoryTab: React.FC = () => {
  const { t } = useTranslation()
  const [formAddRoom] = Form.useForm()
  const [formAddStuff] = Form.useForm()
  const id = usePropertyId()
  const [isModalAddRoomOpen, setIsModalAddRoomOpen] = useState(false)
  const [isModalAddStuffOpen, setIsModalAddStuffOpen] = useState(false)
  const [selectedRoomId, setSelectedRoomId] = useState<string | null>(null)

  const {
    inventory,
    isLoading,
    error,
    createRoom,
    createFurniture,
    deleteRoom,
    deleteFurniture
  } = useInventory(id || '')

  const showModal = (modal: string, roomId?: string) => {
    if (modal === 'addRoom') {
      setIsModalAddRoomOpen(true)
    } else if (modal === 'addStuff' && roomId) {
      setSelectedRoomId(roomId)
      setIsModalAddStuffOpen(true)
    }
  }

  const handleAddRoom = async () => {
    try {
      const { roomName } = await formAddRoom.validateFields()
      const success = await createRoom(roomName)
      if (success) {
        formAddRoom.resetFields()
        setIsModalAddRoomOpen(false)
      }
    } catch (info) {
      console.error('Validate Failed:', info)
    }
  }

  const handleAddStuff = async () => {
    try {
      const { stuffName, itemQuantity } = await formAddStuff.validateFields()
      if (selectedRoomId) {
        const success = await createFurniture(selectedRoomId, {
          name: stuffName,
          quantity: itemQuantity
        })
        if (success) {
          formAddStuff.resetFields()
          setSelectedRoomId(null)
          setIsModalAddStuffOpen(false)
        }
      }
    } catch (info) {
      console.error('Validate Failed:', info)
    }
  }

  const handleCancel = (modal: string) => {
    if (modal === 'addRoom') {
      formAddRoom.resetFields()
      setIsModalAddRoomOpen(false)
    } else if (modal === 'addStuff') {
      formAddStuff.resetFields()
      setSelectedRoomId(null)
      setIsModalAddStuffOpen(false)
    }
  }

  const handleDeleteRoom = (roomId: string) => {
    Modal.confirm({
      title: t(
        'pages.real_property_details.tabs.inventory.remove_room_confirmation'
      ),
      okText: t('components.button.confirm'),
      cancelText: t('components.button.cancel'),
      okButtonProps: { danger: true },
      onOk: () => deleteRoom(roomId)
    })
  }

  const handleDeleteFurniture = (roomId: string, furnitureId: string) => {
    Modal.confirm({
      title: t(
        'pages.real_property_details.tabs.inventory.remove_stuff_confirmation'
      ),
      okText: t('components.button.confirm'),
      cancelText: t('components.button.cancel'),
      okButtonProps: { danger: true },
      onOk: () => deleteFurniture(roomId, furnitureId)
    })
  }

  if (error) {
    return <p>{t('pages.real_property.error.error_fetching_data')}</p>
  }

  return (
    <div className={style.tabContent}>
      <div className={style.buttonAddContainer}>
        <Button type="primary" onClick={() => showModal('addRoom')}>
          {t('components.button.add_room')}
        </Button>
      </div>

      {isLoading && <CardInventoryLoader cards={9} />}

      <div className={style.roomsContainer}>
        {inventory.map(room => (
          <div key={room.roomId} className={style.roomContainer}>
            <div className={style.roomHeader}>
              <span>{room.roomName}</span>
              <CloseOutlined
                onClick={() => handleDeleteRoom(room.roomId)}
                aria-label={t('component.button.close')}
              />
            </div>
            <div className={style.stuffsContainer}>
              {room.stuffs.map(stuff => (
                <div key={stuff.id} className={style.stuffCard}>
                  <span>{stuff.name}</span>
                  <CloseOutlined
                    onClick={() => handleDeleteFurniture(room.roomId, stuff.id)}
                    aria-label={t('component.button.close')}
                    className={style.removeStuffIcon}
                  />
                </div>
              ))}
              <div
                className={style.stuffCardAdd}
                onClick={() => showModal('addStuff', room.roomId)}
                onKeyDown={e => {
                  if (e.key === 'Enter') {
                    showModal('addStuff', room.roomName)
                  }
                }}
                role="button"
                tabIndex={0}
                aria-label={t('component.button.add')}
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

        {inventory.length === 0 && (
          <Empty
            description={
              <Typography.Text>
                {t('components.messages.no_rooms_in_inventory')}
              </Typography.Text>
            }
          />
        )}
      </div>

      {/* Add Room Modal */}
      <Modal
        title={t(
          'pages.real_property_details.tabs.inventory.add_room_modal_title'
        )}
        open={isModalAddRoomOpen}
        onOk={handleAddRoom}
        onCancel={() => handleCancel('addRoom')}
      >
        <Form form={formAddRoom} layout="vertical">
          <Form.Item
            name="roomName"
            label={t('components.input.room_name.label')}
            rules={[
              { required: true, message: t('components.input.room_name.error') }
            ]}
          >
            <Input maxLength={20} showCount />
          </Form.Item>
        </Form>
      </Modal>

      {/* Add Stuff Modal */}
      <Modal
        title={t(
          'pages.real_property_details.tabs.inventory.add_stuff_modal_title'
        )}
        open={isModalAddStuffOpen}
        onOk={handleAddStuff}
        onCancel={() => handleCancel('addStuff')}
      >
        <Form form={formAddStuff} layout="vertical">
          <Form.Item
            name="stuffName"
            label={t('components.input.stuff_name.label')}
            rules={[
              {
                required: true,
                message: t('components.input.stuff_name.error')
              }
            ]}
          >
            <Input maxLength={20} showCount />
          </Form.Item>
          <Form.Item
            name="itemQuantity"
            label={t('components.input.item_quantity.label')}
            rules={[
              {
                required: true,
                message: t('components.input.item_quantity.error')
              },
              { type: 'number', min: 1, max: 1000 }
            ]}
          >
            <InputNumber min={1} max={1000} style={{ width: '100%' }} />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  )
}

export default InventoryTab
