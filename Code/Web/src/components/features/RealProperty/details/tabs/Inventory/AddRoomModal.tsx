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
  const [selectedType, setSelectedType] = useState<string | null>('other')
  const [templateItems, setTemplateItems] = useState<
    { name: string; quantity: number; checked: boolean }[]
  >([])
  const [customItems, setCustomItems] = useState<CustomFurniture[]>([])

  useEffect(() => {
    if (isOpen) {
      form.resetFields()
    }
  }, [isOpen, form])

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
    setSelectedType('other')
    setTemplateItems([])
    setCustomItems([])
  }

  const handleCancel = () => {
    setSelectedType('other')
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

  const handleRoomTypeChange = (value: string) => {
    setSelectedType(value)

    const selectedRoomType = roomTypes.find(
      roomType => roomType.value === value
    )
    if (selectedRoomType) {
      form.setFieldsValue({ roomName: selectedRoomType.label })
    }
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
      aria-labelledby="add-room-modal-title"
      aria-describedby="add-room-modal-description"
    >
      <div id="add-room-modal-description" className="sr-only">
        {t('pages.real_property_details.tabs.inventory.add_room_modal_title')}
      </div>

      <Form
        form={form}
        layout="vertical"
        initialValues={{ roomType: 'other' }}
        aria-labelledby="add-room-modal-title"
      >
        <Form.Item
          name="roomType"
          label={t('components.select.room_type.placeholder')}
          rules={[
            { required: true, message: t('components.input.room_type.error') }
          ]}
        >
          <Select
            options={roomTypes.filter(type => type.value !== 'all')}
            onChange={handleRoomTypeChange}
            aria-label={t('components.select.room_type.placeholder')}
            aria-required="true"
          />
        </Form.Item>
        <Form.Item
          name="roomName"
          label={t('components.input.room_name.label')}
          rules={[
            { required: true, message: t('components.input.room_name.error') }
          ]}
        >
          <Input
            maxLength={20}
            showCount
            id="room-name-input"
            aria-label={t('components.input.room_name.label')}
            aria-required="true"
          />
        </Form.Item>
      </Form>

      {templateItems.length > 0 && (
        <section
          style={{ marginTop: 16 }}
          aria-labelledby="suggested-items-title"
        >
          <h3 id="suggested-items-title" className="sr-only">
            {t('components.inventory.furniture.suggested_items')}
          </h3>
          <p>{t('components.inventory.furniture.suggested_items')}</p>
          <div role="group" aria-labelledby="suggested-items-title">
            {templateItems.map((item, index) => (
              <div
                key={item.name}
                style={{
                  display: 'flex',
                  alignItems: 'center',
                  marginBottom: 8
                }}
              >
                <Checkbox
                  checked={item.checked}
                  onChange={e => {
                    const newItems = [...templateItems]
                    newItems[index].checked = e.target.checked
                    setTemplateItems(newItems)
                  }}
                  aria-describedby={`quantity-${item.name}`}
                >
                  {t(`components.inventory.furniture.${item.name}`)}
                </Checkbox>
                <InputNumber
                  id={`quantity-${item.name}`}
                  min={1}
                  value={item.quantity}
                  onChange={val => {
                    const newItems = [...templateItems]
                    newItems[index].quantity = val || 1
                    setTemplateItems(newItems)
                  }}
                  disabled={!item.checked}
                  style={{ marginLeft: 8 }}
                  aria-label={`${t('components.inventory.furniture.quantity')} ${t(`components.inventory.furniture.${item.name}`)}`}
                />
              </div>
            ))}
          </div>
        </section>
      )}

      <section style={{ marginTop: 24 }} aria-labelledby="custom-items-title">
        <h3 id="custom-items-title" className="sr-only">
          {t('components.inventory.furniture.custom_items')}
        </h3>
        <p>{t('components.inventory.furniture.custom_items')}</p>
        <div role="group" aria-labelledby="custom-items-title">
          {customItems.map(item => (
            <Space key={item.id} style={{ marginBottom: 8 }}>
              <Input
                placeholder={t('components.inventory.furniture.name')}
                value={item.name}
                onChange={e => updateCustomItem(item.id, 'name', e)}
                style={{ width: 200 }}
                aria-label={`${t('components.inventory.furniture.name')} ${customItems.indexOf(item) + 1}`}
              />
              <InputNumber
                min={1}
                value={item.quantity}
                onChange={val =>
                  updateCustomItem(item.id, 'quantity', val || 1)
                }
                aria-label={`${t('components.inventory.furniture.quantity')} ${customItems.indexOf(item) + 1}`}
              />
              <Button
                ghost
                danger
                onClick={() => removeCustomItem(item.id)}
                aria-label={`${t('components.button.remove')} ${item.name || t('components.inventory.furniture.Custom Item')}`}
              >
                {t('components.button.remove')}
              </Button>
            </Space>
          ))}
        </div>
        <Button
          type="dashed"
          onClick={addCustomItem}
          icon={<PlusOutlined aria-hidden="true" />}
          style={{ marginTop: 8 }}
          aria-label={t('components.inventory.furniture.add_custom_item')}
        >
          {t('components.inventory.furniture.add_custom_item')}
        </Button>
      </section>
    </Modal>
  )
}

export default AddRoomModal
