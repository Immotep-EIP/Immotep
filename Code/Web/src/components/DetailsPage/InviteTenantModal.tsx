import React from 'react'
import { Modal, Form, Input, DatePicker, Button, FormProps, message } from 'antd'
import { useTranslation } from 'react-i18next'

import { InviteTenant, InviteTenantModalProps } from '@/interfaces/Tenant/InviteTenant.ts'
import InviteTenants from '@/services/api/Tenant/InviteTenant.ts'

const InviteTenantModal: React.FC<InviteTenantModalProps> = ({ isOpen, onClose }) => {
    const { t } = useTranslation()

    const onFinish: FormProps<InviteTenant>['onFinish'] = async (tenantInfo) => {
        try {
            const formattedTenantInfo = {
                ...tenantInfo,

                // waiting real ID's from the detailPage
                propertyId: "cm3xcwfrs000512r6oh1dt0nm",
            }
            await InviteTenants(formattedTenantInfo)
            message.success(t('pages.realPropertyDetails.inviteTenant'))
            onClose()
        } catch (error: any) {
            if (error.response.status === 409)
                message.error(t('pages.realPropertyDetails.409Error'))
        }
    }

    const onFinishFailed: FormProps<InviteTenant>['onFinishFailed'] = () => {
        message.error(t('pages.realPropertyDetails.fillFields'));
    }

    return (
        <Modal
            title={t('components.button.addTenant')}
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
                        { required: true, message: t('form.error.email') },
                        { type: 'email', message: t('form.error.valid_email') },
                    ]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label={t('components.input.startDate.label')}
                    name="start_date"
                    rules={[{ required: true, message: t('form.error.start_date') }]}
                >
                    <DatePicker style={{ width: '100%' }} />
                </Form.Item>

                <Form.Item
                    label={t('components.input.endDate.label')}
                    name="end_date"
                    rules={[{ required: true, message: t('form.error.end_date') }]}
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