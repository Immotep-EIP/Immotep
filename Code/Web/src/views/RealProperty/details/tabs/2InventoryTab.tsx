import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'

import { Form, Modal, Space, message } from 'antd'

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
          const results = await Promise.all(
            templateItems.map(item =>
              createFurniture(result.roomId!, {
                name: item.name,
                quantity: item.quantity
              })
            )
          )

          const successCount = results.filter(r => r).length
          if (successCount > 0) {
            message.success(
              `${successCount} ${t('components.messages.items_added_successfully')}`
            )
          }
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
    return (
      <section
        role="alert"
        aria-live="assertive"
        aria-labelledby="inventory-error-title"
      >
        <h2 id="inventory-error-title" className="sr-only">
          {t('pages.real_property_details.tabs.inventory.error_title')}
        </h2>
        <p>{t('pages.real_property.error.error_fetching_data')}</p>
      </section>
    )
  }

  const renderContent = () => {
    if (isLoading) {
      return (
        <section
          role="status"
          aria-live="polite"
          aria-labelledby="inventory-loading-title"
        >
          <h3 id="inventory-loading-title" className="sr-only">
            {t('pages.real_property_details.tabs.inventory.loading_title')}
          </h3>
          <CardInventoryLoader cards={9} />
        </section>
      )
    }

    if (filteredInventory.rooms.length === 0) {
      return (
        <section
          role="status"
          aria-live="polite"
          aria-labelledby="inventory-empty-title"
        >
          <h3 id="inventory-empty-title" className="sr-only">
            {t('pages.real_property_details.tabs.inventory.empty_title')}
          </h3>
          <Empty description={t('components.messages.no_rooms_in_inventory')} />
        </section>
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
    <main className={style.tabContent} aria-labelledby="inventory-tab-title">
      <h2 id="inventory-tab-title" className="sr-only">
        {t('pages.real_property_details.tabs.inventory.tab_title')}
      </h2>
      <Space direction="vertical" size="large" style={{ width: '100%' }}>
        <section aria-labelledby="inventory-controls-title">
          <h3 id="inventory-controls-title" className="sr-only">
            {t('pages.real_property_details.tabs.inventory.controls_title')}
          </h3>
          <InventoryControls
            setSearchQuery={setSearchQuery}
            selectedRoomType={selectedRoomType}
            setSelectedRoomType={setSelectedRoomType}
            viewMode={viewMode}
            setViewMode={setViewMode}
            showModal={showModal}
            roomTypes={roomTypes}
          />
        </section>

        <section aria-labelledby="inventory-content-title">
          <h3 id="inventory-content-title" className="sr-only">
            {t('pages.real_property_details.tabs.inventory.content_title')}
          </h3>
          {renderContent()}
        </section>
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
    </main>
  )
}

export default InventoryTab
