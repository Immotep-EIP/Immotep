import React from 'react'
import { Form, Button, Modal, message } from 'antd'
import { useTranslation } from 'react-i18next'

import useProperties from '@/hooks/Property/useProperties'
import PropertyFormFields from '@/components/RealProperty/PropertyForm/PropertyFormFields'
import useImageUpload from '@/hooks/Image/useImageUpload'
import { PropertyFormFieldsType } from '@/utils/types/propertyType'
import style from './RealPropertyCreate.module.css'

interface RealPropertyCreateProps {
  showModalCreate: boolean
  setShowModalCreate: (show: boolean) => void
  setIsPropertyCreated: (isCreated: boolean) => void
}

const RealPropertyCreate: React.FC<RealPropertyCreateProps> = ({
  showModalCreate,
  setShowModalCreate,
  setIsPropertyCreated
}) => {
  const { t } = useTranslation()
  const { loading, createProperty } = useProperties()
  const { uploadProps, imageBase64 } = useImageUpload()
  const [form] = Form.useForm()

  const onFinish = async (values: PropertyFormFieldsType) => {
    try {
      await createProperty(values, imageBase64)
      setShowModalCreate(false)
      message.success(
        t('pages.real_property.add_real_property.property_created')
      )
      setIsPropertyCreated(true)
    } catch (err) {
      message.error(
        t('pages.real_property.add_real_property.error_property_created')
      )
    }
  }

  const onFinishFailed = () => {
    message.error(t('pages.real_property.add_real_property.fill_all_fields'))
  }

  return (
    <Modal
      title={t('pages.real_property.add_real_property.document_title')}
      open={showModalCreate}
      onCancel={() => setShowModalCreate(false)}
      footer={[
        <Button key="back" onClick={() => setShowModalCreate(false)}>
          {t('components.button.cancel')}
        </Button>,
        <Button
          key="submit"
          type="primary"
          loading={loading}
          onClick={() => form.submit()}
        >
          {t('components.button.add')}
        </Button>
      ]}
      style={{
        top: '10%',
        overflow: 'hidden'
      }}
      styles={{
        body: {
          maxHeight: 'calc(70vh - 55px)',
          overflowY: 'auto'
        }
      }}
    >
      <div className={style.pageContainer}>
        <Form
          form={form}
          name="propertyForm"
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          autoComplete="off"
          layout="vertical"
          style={{ width: '90%', maxWidth: '500px', margin: '20px' }}
        >
          <PropertyFormFields uploadProps={uploadProps} />
        </Form>
      </div>
    </Modal>
  )
}

export default RealPropertyCreate
