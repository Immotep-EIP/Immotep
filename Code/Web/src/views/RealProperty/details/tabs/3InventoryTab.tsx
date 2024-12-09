import React, { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { CloseOutlined } from '@ant-design/icons'
import { Button, Form, Input, Modal } from 'antd'
import addIcon from '@/assets/icons/plus.png'
import style from './3InventoryTab.module.css'

const inventoryFromApi = [
  {
    roomName: 'Room 1',
    stuffs: ['Stuff 1', 'Stuff 2', 'Stuff 3']
  },
  {
    roomName: 'Room 2',
    stuffs: ['Stuff 1', 'Stuff 2', 'Stuff 3']
  },
  {
    roomName: 'Room 3',
    stuffs: ['Stuff 1', 'Stuff 2', 'Stuff 3']
  }
]

const InventoryTab: React.FC = () => {
  const { t } = useTranslation()
  const [formAddRoom] = Form.useForm();
  const [formAddStuff] = Form.useForm();
  const [isModalAddRoomOpen, setIsModalAddRoomOpen] = useState(false)
  const [isModalAddStuffOpen, setIsModalAddStuffOpen] = useState(false)

  const [inventory, setInventory] = useState(inventoryFromApi)
  const [selectedRoomName, setSelectedRoomName] = useState<string | null>(null)

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
      .then(() => {
        const values = formAddRoom.getFieldsValue()
        setInventory([...inventory, { ...values, stuffs: [] }])
        formAddRoom.resetFields()
        setIsModalAddRoomOpen(false)
      })
      .catch(info => {
        console.error('Validate Failed:', info)
      })
  }

  const handleAddStuff = () => {
    formAddStuff
    .validateFields()
    .then(() => {
        const { stuffName } = formAddStuff.getFieldsValue()
        if (selectedRoomName) {
          setInventory(
            inventory.map(room =>
              room.roomName === selectedRoomName
                ? { ...room, stuffs: [...room.stuffs, stuffName] }
                : room
            )
          )
        }
        formAddStuff.resetFields()
        setSelectedRoomName(null)
        setIsModalAddStuffOpen(false)
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

  useEffect(() => {
    console.log('Inventory:', inventory)
  }, [inventory])

  const removeRoom = (roomName: string) => {
    setInventory(inventory.filter(r => r.roomName !== roomName))
  }

  const removeStuff = (roomName: string, stuff: string) => {
    setInventory(
      inventory.map(r =>
        r.roomName === roomName
          ? {
              ...r,
              stuffs: r.stuffs.filter(s => s !== stuff)
            }
          : r
      )
    )
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
          'pages.realPropertyDetails.tabs.inventory.add_room_modal_title'
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
              <CloseOutlined onClick={() => removeRoom(room.roomName)} />
            </div>
            <div className={style.stuffsContainer}>
              {room?.stuffs?.map(stuff => (
                <div key={stuff} className={style.stuffCard}>
                  <span>{stuff}</span>
                  <CloseOutlined
                    onClick={() => removeStuff(room.roomName, stuff)}
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
            'pages.realPropertyDetails.tabs.inventory.add_stuff_modal_title'
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
          </Form>
        </Modal>
      </div>
    </div>
  )
}

export default InventoryTab
