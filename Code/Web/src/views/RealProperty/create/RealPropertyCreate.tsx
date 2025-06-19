import React from 'react'
import { CloseCircleOutlined, PlusCircleOutlined } from '@ant-design/icons'
import { useTranslation } from 'react-i18next'

import { Drawer, Form, message } from 'antd'

import useProperties from '@/hooks/Property/useProperties'
import useImageUpload from '@/hooks/Image/useImageUpload'
import { Button } from '@/components/common'
import PropertyFormFields from '@/components/features/RealProperty/PropertyForm/PropertyFormFields'
import { PropertyFormFieldsType } from '@/utils/types/propertyType'

import { RealPropertyCreateProps } from '@/interfaces/Property/Property'

import style from './RealPropertyCreate.module.css'

const RealPropertyCreate: React.FC<RealPropertyCreateProps> = ({
  showModalCreate,
  setShowModalCreate,
  setIsPropertyCreated
}) => {
  const { t } = useTranslation()
  const { loading, createProperty } = useProperties()
  const { uploadProps, imageBase64, resetImage } = useImageUpload()
  const [form] = Form.useForm()

  const onFinish = async (values: PropertyFormFieldsType) => {
    try {
      await createProperty(values, imageBase64)
      setShowModalCreate(false)
      message.success(
        t('pages.real_property.add_real_property.property_created')
      )
      setIsPropertyCreated(true)
      form.resetFields()
      resetImage()
    } catch (err) {
      message.error(
        t('pages.real_property.add_real_property.error_property_created')
      )
    }
  }

  const onFinishFailed = () => {
    message.error(t('pages.real_property.add_real_property.fill_all_fields'))
  }

  const handleCancel = () => {
    setShowModalCreate(false)
    form.resetFields()
    resetImage()
  }

  return (
    <Drawer
      title={
        <div className={style.drawerTitle}>
          {t('pages.real_property.add_real_property.document_title')}
          <div className={style.buttonsContainer}>
            <Button
              type="default"
              key="back"
              onClick={handleCancel}
              icon={<CloseCircleOutlined />}
            >
              {t('components.button.cancel')}
            </Button>
            <Button
              key="submit"
              loading={loading}
              onClick={() => form.submit()}
              icon={<PlusCircleOutlined />}
            >
              {t('components.button.add')}
            </Button>
          </div>
        </div>
      }
      open={showModalCreate}
      onClose={handleCancel}
      width={550}
    >
      <div className={style.pageContainer}>
        <Form
          form={form}
          name="propertyForm"
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          autoComplete="off"
          layout="vertical"
          style={{ width: '100%' }}
        >
          <PropertyFormFields uploadProps={uploadProps} />
        </Form>
      </div>
    </Drawer>
  )
}

export default RealPropertyCreate
