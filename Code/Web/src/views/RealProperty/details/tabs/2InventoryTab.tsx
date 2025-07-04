import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'

import { Form, Modal, Space } from 'antd'

import { Empty } from '@/components/common'
import { usePropertyContext } from '@/context/propertyContext'
import useInventory from '@/hooks/Property/useInventory'
import CardInventoryLoader from '@/components/ui/Loader/CardInventoryLoader'
import GridView from '@/components/features/RealProperty/details/tabs/Inventory/GridView'
import PlanView from '@/components/features/RealProperty/details/tabs/Inventory/PlanView'
import ListView from '@/components/features/RealProperty/details/tabs/Inventory/ListView'
import AddRoomModal from '@/components/features/RealProperty/details/tabs/Inventory/AddRoomModal'
import AddStuffModal from '@/components/features/RealProperty/details/tabs/Inventory/AddStuffModal'
import InventoryControls from '@/components/features/RealProperty/details/tabs/Inventory/InventoryControls'
import { getRoomTypeOptions } from '@/utils/types/roomTypes'

import style from './2InventoryTab.module.css'

const InventoryTab: React.FC = () => {
  const { t } = useTranslation()
  const [formAddRoom] = Form.useForm()
  const [formAddFurniture] = Form.useForm()
  const { property } = usePropertyContext()
  const [isModalAddRoomOpen, setIsModalAddRoomOpen] = useState(false)
  const [isModalAddStuffOpen, setIsModalAddStuffOpen] = useState(false)
  const [selectedRoomId, setSelectedRoomId] = useState<string | null>(null)
  const [searchQuery, setSearchQuery] = useState('')
  const [viewMode, setViewMode] = useState<'grid' | 'list' | 'plan'>('grid')
  const [selectedRoomType, setSelectedRoomType] = useState<string>('all')
  const [layouts, setLayouts] = useState<any>({})

  const {
    inventory,
    isLoading,
    error,
    createRoom,
    createFurniture,
    deleteRoom,
    deleteFurniture,
    refreshInventory
  } = useInventory(property?.id || '')

  const roomTypes = getRoomTypeOptions(t)

  const showModal = (modal: string, roomId?: string) => {
    if (modal === 'addRoom') {
      setIsModalAddRoomOpen(true)
    } else if (modal === 'addStuff' && roomId) {
      setSelectedRoomId(roomId)
      setIsModalAddStuffOpen(true)
    }
  }

  const handleAddRoom = async (
    templateItems?: { name: string; quantity: number }[]
  ) => {
    try {
      const { roomName, roomType } = await formAddRoom.validateFields()
      const result = await createRoom(roomName, roomType)

      if (result.success) {
        if (templateItems && templateItems.length > 0) {
          await Promise.all(
            templateItems.map(item =>
              createFurniture(result.roomId!, {
                name: item.name,
                quantity: item.quantity
              })
            )
          )
        }
        refreshInventory()
        formAddRoom.resetFields()
        setIsModalAddRoomOpen(false)
      }
    } catch (info) {
      console.error('Validate Failed:', info)
    }
  }

  const handleAddFurniture = async () => {
    try {
      const { stuffName, itemQuantity } =
        await formAddFurniture.validateFields()
      if (selectedRoomId) {
        const success = await createFurniture(selectedRoomId, {
          name: stuffName,
          quantity: itemQuantity
        })
        if (success) {
          formAddFurniture.resetFields()
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
      formAddFurniture.resetFields()
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
      onOk: () => {
        deleteRoom(roomId)
      }
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

  const filterInventory = () => ({
    rooms: inventory.rooms.filter(room => {
      const matchesSearch = room.name
        ? room.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
          room.furniture?.some(furniture =>
            furniture.name.toLowerCase().includes(searchQuery.toLowerCase())
          )
        : false
      const matchesType =
        selectedRoomType === 'all' ||
        (room.type &&
          room.type.toLowerCase() === selectedRoomType.toLowerCase())
      return matchesSearch && matchesType
    })
  })

  const filteredInventory = filterInventory()

  if (error) {
    return <p>{t('pages.real_property.error.error_fetching_data')}</p>
  }

  const renderContent = () => {
    if (isLoading) {
      return <CardInventoryLoader cards={9} />
    }

    if (filteredInventory.rooms.length === 0) {
      return (
        <Empty description={t('components.messages.no_rooms_in_inventory')} />
      )
    }

    switch (viewMode) {
      case 'grid':
        return (
          <GridView
            inventory={filteredInventory}
            showModal={showModal}
            handleDeleteRoom={handleDeleteRoom}
            handleDeleteFurniture={handleDeleteFurniture}
          />
        )
      case 'list':
        return (
          <ListView
            inventory={filteredInventory}
            showModal={showModal}
            handleDeleteRoom={handleDeleteRoom}
            handleDeleteFurniture={handleDeleteFurniture}
          />
        )
      case 'plan':
        return (
          <PlanView
            inventory={filteredInventory}
            layouts={layouts}
            setLayouts={setLayouts}
            showModal={showModal}
            handleDeleteFurniture={handleDeleteFurniture}
          />
        )
      default:
        return null
    }
  }

  return (
    <div className={style.tabContent}>
      <Space direction="vertical" size="large" style={{ width: '100%' }}>
        <InventoryControls
          setSearchQuery={setSearchQuery}
          selectedRoomType={selectedRoomType}
          setSelectedRoomType={setSelectedRoomType}
          viewMode={viewMode}
          setViewMode={setViewMode}
          showModal={showModal}
          roomTypes={roomTypes}
        />

        {renderContent()}
      </Space>

      <AddRoomModal
        isOpen={isModalAddRoomOpen}
        onOk={handleAddRoom}
        onCancel={() => handleCancel('addRoom')}
        form={formAddRoom}
        roomTypes={roomTypes}
      />

      <AddStuffModal
        isOpen={isModalAddStuffOpen}
        onOk={handleAddFurniture}
        onCancel={() => handleCancel('addStuff')}
        form={formAddFurniture}
      />
    </div>
  )
}

export default InventoryTab
