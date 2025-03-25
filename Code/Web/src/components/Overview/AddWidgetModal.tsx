import React from 'react'
import { Modal, Form, Input, InputNumber, Button, message, Select } from 'antd'
import { useTranslation } from 'react-i18next'
import {
  addWidgetType,
  AddWidgetModalProps
} from '@/interfaces/Widgets/Widgets.ts'

const AddWidgetModal: React.FC<AddWidgetModalProps> = ({
  isOpen,
  onClose,
  onAddWidget
}) => {
  const { t } = useTranslation()

  const widgetTypes = [
    { label: 'User Info', value: 'UserInfoWidget' },
    { label: 'Maintenance', value: 'MaintenanceWidget' }
  ]

  const onFinish = (values: addWidgetType) => {
    onAddWidget(values)
    message.success(t('pages.overview.widget_created'))
    onClose()
  }

  const onFinishFailed = () => {
    message.error(t('pages.overview.fill_fields'))
  }

  return (
    <Modal
      title={t('pages.overview.widget_creation.title')}
      open={isOpen}
      onCancel={onClose}
      footer={null}
    >
      <Form
        name="add_widget"
        onFinish={onFinish}
        onFinishFailed={onFinishFailed}
        layout="vertical"
      >
        <Form.Item
          label={t('components.input.widget_name.label')}
          name="name"
          rules={[
            { required: true, message: t('components.input.widget_name.error') }
          ]}
        >
          <Input
            placeholder={t('components.input.widget_name.placeholder')}
            aria-label={t('components.input.widget_name.placeholder')}
          />
        </Form.Item>

        <Form.Item
          label={t('components.input.widget_type.label')}
          name="types"
          rules={[
            { required: true, message: t('components.input.widget_type.error') }
          ]}
        >
          <Select
            placeholder={t('components.input.widget_type.placeholder')}
            aria-label={t('components.input.widget_type.placeholder')}
            options={widgetTypes}
          />
        </Form.Item>

        <Form.Item
          label={t('components.input.widget_width.label')}
          name="width"
          rules={[
            {
              required: true,
              message: t('components.input.widget_width.error')
            },
            { type: 'number', min: 1, message: t('form.error.widgetWidthMin') }
          ]}
        >
          <InputNumber
            min={1}
            placeholder={t('components.input.widget_width.placeholder')}
            aria-label={t('components.input.widget_width.placeholder')}
            style={{ width: '100%' }}
          />
        </Form.Item>

        <Form.Item
          label={t('components.input.widget_height.label')}
          name="height"
          rules={[
            {
              required: true,
              message: t('components.input.widget_height.error')
            },
            { type: 'number', min: 1, message: t('form.error.widgetHeightMin') }
          ]}
        >
          <InputNumber
            min={1}
            placeholder={t('components.input.widget_height.placeholder')}
            aria-label={t('components.input.widget_height.placeholder')}
            style={{ width: '100%' }}
          />
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit">
            {t('components.button.add_widget')}
          </Button>
        </Form.Item>
      </Form>
    </Modal>
  )
}

export default AddWidgetModal
