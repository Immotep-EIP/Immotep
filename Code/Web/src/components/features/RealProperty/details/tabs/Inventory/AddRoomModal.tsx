import React, { useState, useEffect } from 'react'
import { useTranslation } from 'react-i18next'

import { Modal, Form, Select, Checkbox, InputNumber, Space } from 'antd'
import { PlusOutlined } from '@ant-design/icons'

import { Button, Input } from '@/components/common'

import { ROOM_TEMPLATES } from '@/utils/types/roomTypes'

import { AddRoomModalProps } from '@/interfaces/Property/Inventory/Room/Room'
import { CustomFurniture } from '@/interfaces/Property/Inventory/Room/Furniture/Furniture'

const AddRoomModal: React.FC<AddRoomModalProps> = ({
  isOpen,
  onOk,
  onCancel,
  form,
  roomTypes
}) => {
  const { t } = useTranslation()
  const [selectedType, setSelectedType] = useState<string | null>(null)
  const [templateItems, setTemplateItems] = useState<
    { name: string; quantity: number; checked: boolean }[]
  >([])
  const [customItems, setCustomItems] = useState<CustomFurniture[]>([])

  useEffect(() => {
    if (selectedType && ROOM_TEMPLATES[selectedType]) {
      setTemplateItems(
        ROOM_TEMPLATES[selectedType].map(item => ({ ...item, checked: true }))
      )
    } else {
      setTemplateItems([])
    }
  }, [selectedType])

  const handleOk = async () => {
    await form.validateFields()
    onOk([
      ...templateItems
        .filter(item => item.checked)
        .map(({ name, quantity }) => ({ name, quantity })),
      ...customItems.map(({ name, quantity }) => ({ name, quantity }))
    ])
    setSelectedType(null)
    setTemplateItems([])
    setCustomItems([])
  }

  const handleCancel = () => {
    setSelectedType(null)
    setTemplateItems([])
    setCustomItems([])
    onCancel()
  }

  const addCustomItem = () => {
    setCustomItems([
      ...customItems,
      { id: crypto.randomUUID(), name: '', quantity: 1 }
    ])
  }

  const removeCustomItem = (id: string) => {
    setCustomItems(customItems.filter(item => item.id !== id))
  }

  const updateCustomItem = (
    id: string,
    field: 'name' | 'quantity',
    value: string | number
  ) => {
    setCustomItems(
      customItems.map(item =>
        item.id === id ? { ...item, [field]: value } : item
      )
    )
  }

  return (
    <Modal
      title={t(
        'pages.real_property_details.tabs.inventory.add_room_modal_title'
      )}
      open={isOpen}
      onOk={handleOk}
      onCancel={handleCancel}
      okText={t('components.button.add')}
      cancelText={t('components.button.cancel')}
      width={600}
    >
      <Form form={form} layout="vertical">
        <Form.Item
          name="roomName"
          label={t('components.input.room_name.label')}
          rules={[
            { required: true, message: t('components.input.room_name.error') }
          ]}
        >
          <Input maxLength={20} showCount />
        </Form.Item>
        <Form.Item
          name="roomType"
          label={t('components.select.room_type.placeholder')}
          rules={[
            { required: true, message: t('components.input.room_type.error') }
          ]}
        >
          <Select
            options={roomTypes.filter(type => type.value !== 'all')}
            onChange={value => setSelectedType(value)}
          />
        </Form.Item>
      </Form>

      {templateItems.length > 0 && (
        <div style={{ marginTop: 16 }}>
          <p>{t('components.inventory.furniture.suggested_items')}</p>
          {templateItems.map((item, index) => (
            <div
              key={item.name}
              style={{ display: 'flex', alignItems: 'center', marginBottom: 8 }}
            >
              <Checkbox
                checked={item.checked}
                onChange={e => {
                  const newItems = [...templateItems]
                  newItems[index].checked = e.target.checked
                  setTemplateItems(newItems)
                }}
              >
                {t(`components.inventory.furniture.${item.name}`)}
              </Checkbox>
              <InputNumber
                min={1}
                value={item.quantity}
                onChange={val => {
                  const newItems = [...templateItems]
                  newItems[index].quantity = val || 1
                  setTemplateItems(newItems)
                }}
                disabled={!item.checked}
                style={{ marginLeft: 8 }}
              />
            </div>
          ))}
        </div>
      )}

      <div style={{ marginTop: 24 }}>
        <p>{t('components.inventory.furniture.custom_items')}</p>
        {customItems.map(item => (
          <Space key={item.id} style={{ marginBottom: 8 }}>
            <Input
              placeholder={t('components.inventory.furniture.name')}
              value={item.name}
              onChange={e => updateCustomItem(item.id, 'name', e)}
              style={{ width: 200 }}
            />
            <InputNumber
              min={1}
              value={item.quantity}
              onChange={val => updateCustomItem(item.id, 'quantity', val || 1)}
            />
            <Button ghost danger onClick={() => removeCustomItem(item.id)}>
              {t('components.button.remove')}
            </Button>
          </Space>
        ))}
        <Button
          type="dashed"
          onClick={addCustomItem}
          icon={<PlusOutlined />}
          style={{ marginTop: 8 }}
        >
          {t('components.inventory.furniture.add_custom_item')}
        </Button>
      </div>
    </Modal>
  )
}

export default AddRoomModal
