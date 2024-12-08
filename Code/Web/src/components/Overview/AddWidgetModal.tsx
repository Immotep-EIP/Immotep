import React from 'react';
import {Modal, Form, Input, InputNumber, Button, message, Select} from 'antd';
import {useTranslation} from 'react-i18next';
import {addWidgetType, AddWidgetModalProps} from "@/interfaces/Widgets/Widgets.ts";

const AddWidgetModal: React.FC<AddWidgetModalProps> = ({isOpen, onClose, onAddWidget}) => {
    const {t} = useTranslation();

    const widgetTypes = [
        {label: 'User Info', value: 'UserInfoWidget'},
        {label: 'Maintenance', value: 'MaintenanceWidget'},
    ];

    const onFinish = (values: addWidgetType) => {
        onAddWidget(values);
        message.success(t('pages.overview.widgetCreation.widgetCreated'));
        onClose();
    };

    const onFinishFailed = () => {
        message.error(t('pages.overview.widgetCreation.fillFields'));
    };

    return (
        <Modal
            title={t('pages.overview.widgetCreation.title')}
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
                    label={t('components.input.widgetName.label')}
                    name="name"
                    rules={[
                        {required: true, message: t('components.input.widgetName.error')},
                    ]}
                >
                    <Input placeholder={t('components.input.widgetName.placeholder')}/>
                </Form.Item>

                <Form.Item
                    label={t('components.input.widgetType.label')}
                    name="types"
                    rules={[
                        {required: true, message: t('components.input.widgetType.error')},
                    ]}
                >
                    <Select
                        placeholder={t('components.input.widgetType.placeholder')}
                        options={widgetTypes}
                    />
                </Form.Item>

                <Form.Item
                    label={t('components.input.widgetWidth.label')}
                    name="width"
                    rules={[
                        {required: true, message: t('components.input.widgetWidth.error')},
                        {type: 'number', min: 1, message: t('form.error.widgetWidthMin')},
                    ]}
                >
                    <InputNumber min={1} placeholder={t('components.input.widgetWidth.placeholder')}
                                 style={{width: '100%'}}/>
                </Form.Item>

                <Form.Item
                    label={t('components.input.widgetHeight.label')}
                    name="height"
                    rules={[
                        {required: true, message: t('components.input.widgetHeight.error')},
                        {type: 'number', min: 1, message: t('form.error.widgetHeightMin')},
                    ]}
                >
                    <InputNumber min={1} placeholder={t('components.input.widgetHeight.placeholder')}
                                 style={{width: '100%'}}/>
                </Form.Item>

                <Form.Item>
                    <Button type="primary" htmlType="submit">
                        {t('components.button.addWidget')}
                    </Button>
                </Form.Item>
            </Form>
        </Modal>
    );
};

export default AddWidgetModal;