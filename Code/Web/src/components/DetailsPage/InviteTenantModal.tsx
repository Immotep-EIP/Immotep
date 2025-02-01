import React from 'react'
import {
  Modal,
  Form,
  Input,
  DatePicker,
  Button,
  FormProps,
  message
} from 'antd'
import { useTranslation } from 'react-i18next'
import dayjs from 'dayjs'

import {
  InviteTenant,
  InviteTenantModalProps
} from '@/interfaces/Tenant/InviteTenant.ts'
import InviteTenants from '@/services/api/Tenant/InviteTenant.ts'

const InviteTenantModal: React.FC<InviteTenantModalProps> = ({
  isOpen,
  onClose,
  propertyId
}) => {
  const { t } = useTranslation()
  const onFinish: FormProps<InviteTenant>['onFinish'] = async tenantInfo => {
    try {
      // eslint-disable-next-line camelcase
      const { start_date, end_date } = tenantInfo
      const formattedTenantInfo = {
        ...tenantInfo,
        propertyId
      }
      if (dayjs(start_date).isAfter(dayjs(end_date))) {
        message.error(t('pages.real_property_details.dateError'))
        return
      }
      await InviteTenants(formattedTenantInfo)
      message.success(t('pages.real_property_details.invite_tenant'))
      onClose()
    } catch (error: any) {
      if (error.response.status === 409)
        message.error(t('pages.real_property_details.409_error'))
    }
  }

  const onFinishFailed: FormProps<InviteTenant>['onFinishFailed'] = () => {
    message.error(t('pages.real_property_details.fill_fields'))
  }

  return (
    <Modal
      title={t('components.button.add_tenant')}
      open={isOpen}
      onCancel={onClose}
      footer={null}
    >
      <Form
        name="invite_tenant"
        onFinish={onFinish}
        onFinishFailed={onFinishFailed}
        layout="vertical"
        style={{ width: '90%', maxWidth: '500px', margin: '20px' }}
      >
        <Form.Item
          label={t('components.input.email.label')}
          name="tenant_email"
          rules={[
            { required: true, message: t('components.input.email.error') },
            { type: 'email', message: t('components.input.email.valid_email') }
          ]}
        >
          <Input placeholder={t('components.input.email.placeholder')} />
        </Form.Item>

        <Form.Item
          label={t('components.input.start_date.label')}
          name="start_date"
          rules={[{ required: true, message: t('form.error.start_date') }]}
        >
          <DatePicker style={{ width: '100%' }} />
        </Form.Item>

        <Form.Item
          label={t('components.input.end_date.label')}
          name="end_date"
          rules={[{ required: false, message: t('form.error.end_date') }]}
        >
          <DatePicker style={{ width: '100%' }} />
        </Form.Item>

        <Form.Item>
          <Button type="primary" htmlType="submit">
            {t('components.button.add')}
          </Button>
        </Form.Item>
      </Form>
    </Modal>
  )
}

export default InviteTenantModal
