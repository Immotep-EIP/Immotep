import React, { useState, useEffect } from 'react'
import { Modal, Form, Input, Select, Checkbox, InputNumber } from 'antd'
import { useTranslation } from 'react-i18next'

import { AddRoomModalProps } from '@/interfaces/Property/Inventory/Room/Room'
import { ROOM_TEMPLATES } from '@/utils/types/roomTypes'

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
    onOk(
      templateItems
        .filter(item => item.checked)
        .map(({ name, quantity }) => ({ name, quantity }))
    )
    setSelectedType(null)
    setTemplateItems([])
  }

  const handleCancel = () => {
    setSelectedType(null)
    setTemplateItems([])
    onCancel()
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
          label={t('components.input.room_type.label')}
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
    </Modal>
  )
}

export default AddRoomModal
