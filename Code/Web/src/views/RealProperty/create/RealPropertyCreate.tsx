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
        <header className={style.drawerTitle}>
          <h2 id="create-property-title">
            {t('pages.real_property.add_real_property.document_title')}
          </h2>
          <div
            className={style.buttonsContainer}
            role="toolbar"
            aria-label={t('pages.real_property.add_real_property.toolbar_aria')}
          >
            <Button
              type="default"
              key="back"
              onClick={handleCancel}
              icon={<CloseCircleOutlined aria-hidden="true" />}
              aria-label={t('components.button.cancel_create_property')}
            >
              {t('components.button.cancel')}
            </Button>
            <Button
              key="submit"
              loading={loading}
              disabled={loading}
              onClick={() => form.submit()}
              icon={<PlusCircleOutlined aria-hidden="true" />}
              aria-label={t('components.button.add_property')}
              aria-describedby="create-property-submit-help"
            >
              {loading
                ? t('components.button.adding')
                : t('components.button.add')}
            </Button>
          </div>
        </header>
      }
      open={showModalCreate}
      onClose={handleCancel}
      width={550}
      aria-labelledby="create-property-title"
      aria-describedby="create-property-description"
    >
      <main className={style.pageContainer}>
        <h3 id="create-property-form-title" className="sr-only">
          {t('pages.real_property.add_real_property.form_title')}
        </h3>
        <div id="create-property-description" className="sr-only">
          {t('pages.real_property.add_real_property.form_description')}
        </div>
        <div id="create-property-submit-help" className="sr-only">
          {t('pages.real_property.add_real_property.submit_help')}
        </div>
        <Form
          form={form}
          name="propertyForm"
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          autoComplete="on"
          layout="vertical"
          style={{ width: '100%' }}
          aria-labelledby="create-property-form-title"
          aria-describedby="create-property-description"
          noValidate
        >
          <PropertyFormFields uploadProps={uploadProps} />
        </Form>
      </main>
    </Drawer>
  )
}

export default RealPropertyCreate
