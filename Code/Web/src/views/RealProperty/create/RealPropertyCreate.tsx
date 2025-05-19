import React, { useState, useEffect } from 'react'
import { Form, Button, Modal, message } from 'antd'
import { DownOutlined } from '@ant-design/icons'
import { useTranslation } from 'react-i18next'

import useProperties from '@/hooks/Property/useProperties'
import PropertyFormFields from '@/components/features/RealProperty/PropertyForm/PropertyFormFields'
import useImageUpload from '@/hooks/Image/useImageUpload'
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
  const [showScrollIndicator, setShowScrollIndicator] = useState(true)
  const [canScroll, setCanScroll] = useState(false)

  useEffect(() => {
    if (showModalCreate) {
      setShowScrollIndicator(true)

      const checkIfScrollable = () => {
        const modalBody = document.querySelector('.ant-modal-body')
        if (modalBody) {
          const canScrollContent =
            modalBody.scrollHeight > modalBody.clientHeight
          setCanScroll(canScrollContent)

          const handleScroll = () => {
            if (modalBody.scrollTop > 30) {
              setShowScrollIndicator(false)
            }
          }

          modalBody.addEventListener('scroll', handleScroll)
          return () => modalBody.removeEventListener('scroll', handleScroll)
        }

        return undefined
      }

      setTimeout(checkIfScrollable, 300)
    }
  }, [showModalCreate])

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
    <Modal
      title={
        <div className={style.modalTitleContainer}>
          {t('pages.real_property.add_real_property.document_title')}
        </div>
      }
      open={showModalCreate}
      onCancel={handleCancel}
      footer={[
        <Button key="back" onClick={handleCancel}>
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
          overflowY: 'auto',
          position: 'relative'
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
        {canScroll && showScrollIndicator && (
          <div className={style.scrollIndicator}>
            <DownOutlined className={style.scrollIcon} />
          </div>
        )}
      </div>
    </Modal>
  )
}

export default RealPropertyCreate
