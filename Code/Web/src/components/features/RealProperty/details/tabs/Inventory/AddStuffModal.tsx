import React from 'react'
import { Modal, Form, Input, InputNumber } from 'antd'
import { useTranslation } from 'react-i18next'

import { AddFurnitureModalProps } from '@/interfaces/Property/Inventory/Room/Furniture/Furniture'

const AddStuffModal: React.FC<AddFurnitureModalProps> = ({
  isOpen,
  onOk,
  onCancel,
  form
}) => {
  const { t } = useTranslation()

  return (
    <Modal
      title={t(
        'pages.real_property_details.tabs.inventory.add_stuff_modal_title'
      )}
      open={isOpen}
      onOk={onOk}
      onCancel={onCancel}
    >
      <Form form={form} layout="vertical">
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
  )
}

export default AddStuffModal
