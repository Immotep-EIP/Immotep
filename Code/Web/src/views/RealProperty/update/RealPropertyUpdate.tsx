import React, { useEffect } from 'react'
import { Form, Modal, message } from 'antd'
import { useTranslation } from 'react-i18next'

import useProperties from '@/hooks/Property/useProperties'
import useImageUpload from '@/hooks/Image/useImageUpload'
import useImageCache from '@/hooks/Image/useImageCache'
import { Button } from '@/components/common'
import PropertyFormFields from '@/components/features/RealProperty/PropertyForm/PropertyFormFields'
import { PropertyFormFieldsType } from '@/utils/types/propertyType'

import { RealPropertyUpdateProps } from '@/interfaces/Property/Property'

import style from './RealPropertyUpdate.module.css'

const RealPropertyUpdate: React.FC<RealPropertyUpdateProps> = ({
  propertyData,
  isModalUpdateOpen,
  setIsModalUpdateOpen,
  setIsPropertyUpdated
}) => {
  const { t } = useTranslation()
  const { loading, updateProperty } = useProperties()
  const { uploadProps, imageBase64 } = useImageUpload()
  const { updateCache } = useImageCache(propertyData?.id, () =>
    Promise.resolve({ data: imageBase64 })
  )
  const [form] = Form.useForm()

  useEffect(() => {
    if (propertyData && isModalUpdateOpen) {
      form.setFieldsValue({
        name: propertyData.name,
        apartment_number: propertyData.apartment_number,
        address: propertyData.address,
        postal_code: propertyData.postal_code,
        city: propertyData.city,
        country: propertyData.country,
        area_sqm: propertyData.area_sqm,
        rental_price_per_month: propertyData.rental_price_per_month,
        deposit_price: propertyData.deposit_price
      })
    }
  }, [propertyData, form, isModalUpdateOpen])

  const onFinish = async (values: PropertyFormFieldsType) => {
    if (!propertyData || !values) return
    try {
      await updateProperty(values, imageBase64, propertyData.id)
      if (imageBase64) {
        await updateCache(imageBase64)
      }
      setIsModalUpdateOpen(false)
      message.success(
        t('pages.real_property.update_real_property.property_updated')
      )
      setIsPropertyUpdated(true)
    } catch (err) {
      message.error(
        t('pages.real_property.update_real_property.error_property_updated')
      )
    }
  }

  const onFinishFailed = () => {
    message.error(t('pages.real_property.update_real_property.fill_all_fields'))
  }

  return (
    <Modal
      title={t('pages.real_property.update_real_property.title')}
      open={isModalUpdateOpen}
      onCancel={() => setIsModalUpdateOpen(false)}
      footer={[
        <Button
          key="back"
          type="default"
          onClick={() => setIsModalUpdateOpen(false)}
        >
          {t('components.button.cancel')}
        </Button>,
        <Button key="submit" loading={loading} onClick={() => form.submit()}>
          {t('components.button.update')}
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

export default RealPropertyUpdate
