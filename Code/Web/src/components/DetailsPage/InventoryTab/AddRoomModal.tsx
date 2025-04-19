import React from 'react'
import { Modal, Form, Input, Select } from 'antd'
import { useTranslation } from 'react-i18next'

import { AddRoomModalProps } from '@/interfaces/Property/Inventory/Room/Room'

const AddRoomModal: React.FC<AddRoomModalProps> = ({
  isOpen,
  onOk,
  onCancel,
  form,
  roomTypes
}) => {
  const { t } = useTranslation()

  return (
    <Modal
      title={t(
        'pages.real_property_details.tabs.inventory.add_room_modal_title'
      )}
      open={isOpen}
      onOk={onOk}
      onCancel={onCancel}
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
          <Select options={roomTypes.filter(type => type.value !== 'all')} />
        </Form.Item>
      </Form>
    </Modal>
  )
}

export default AddRoomModal
