import React, { useState } from 'react'
import {
  Form,
  Input,
  Button,
  Upload,
  UploadProps,
  message,
  UploadFile,
  Modal
} from 'antd'
import { useTranslation } from 'react-i18next'
import { UploadOutlined } from '@ant-design/icons'
import fileToBase64 from '@/utils/base64/fileToBase'
import useProperties from '@/hooks/useEffect/useProperties'
import style from './RealPropertyCreate.module.css'

type FieldType = {
  name: string
  apartment_number: string
  address: string
  zipCode: string
  city: string
  country: string
  area: string
  rental: string
  deposit: string
  picture: string
}

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
  const [fileList, setFileList] = useState<UploadFile[]>([])
  const [imageBase64, setImageBase64] = useState<string | null>(null)
  const [form] = Form.useForm()

  const uploadProps: UploadProps = {
    name: 'propertyPicture',
    maxCount: 1,
    fileList,
    accept: '.png, .jpg, .jpeg',
    beforeUpload: async file => {
      const base64 = await fileToBase64(file)
      setImageBase64(base64)
      return false
    },
    onChange(info) {
      setFileList(info.fileList)
      if (info.file.status === 'done') {
        message.success(`${info.file.name} file uploaded successfully`)
      } else if (info.file.status === 'error') {
        message.error(`${info.file.name} file upload failed.`)
      }
    }
  }

  const onFinish = async (values: FieldType) => {
    const propertyData = {
      name: values.name,
      apartment_number: values.apartment_number,
      address: values.address,
      city: values.city,
      postal_code: values.zipCode,
      country: values.country,
      area_sqm: parseFloat(values.area || '0'),
      rental_price_per_month: parseFloat(values.rental || '0'),
      deposit_price: parseFloat(values.deposit || '0'),
      picture: ''
    }
    try {
      await createProperty(propertyData, imageBase64)
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
          <Form.Item<FieldType>
            label={t('components.input.property_name.label')}
            name="name"
            rules={[
              {
                required: true,
                message: t('components.input.property_name.error')
              }
            ]}
          >
            <Input
              placeholder={t('components.input.property_name.placeholder')}
            />
          </Form.Item>

          <Form.Item<FieldType>
            label={t('components.input.apartment_number.label')}
            name="apartment_number"
            rules={[
              {
                required: false,
                message: t('components.input.apartment_number.error')
              }
            ]}
          >
            <Input
              placeholder={t('components.input.apartment_number.placeholder')}
            />
          </Form.Item>

          <Form.Item<FieldType>
            label={t('components.input.address.label')}
            name="address"
            rules={[
              { required: true, message: t('components.input.address.error') }
            ]}
          >
            <Input placeholder={t('components.input.address.placeholder')} />
          </Form.Item>

          <Form.Item<FieldType>
            label={t('components.input.zip_code.label')}
            name="zipCode"
            rules={[
              { required: true, message: t('components.input.zip_code.error') }
            ]}
          >
            <Input placeholder={t('components.input.zip_code.placeholder')} />
          </Form.Item>

          <Form.Item<FieldType>
            label={t('components.input.city.label')}
            name="city"
            rules={[
              { required: true, message: t('components.input.city.error') }
            ]}
          >
            <Input placeholder={t('components.input.city.placeholder')} />
          </Form.Item>

          <Form.Item<FieldType>
            label={t('components.input.country.label')}
            name="country"
            rules={[
              { required: true, message: t('components.input.country.error') }
            ]}
          >
            <Input placeholder={t('components.input.country.placeholder')} />
          </Form.Item>

          <Form.Item<FieldType>
            label={t('components.input.area.label')}
            name="area"
            rules={[
              { required: true, message: t('components.input.area.error') }
            ]}
          >
            <Input placeholder={t('components.input.area.placeholder')} />
          </Form.Item>

          <Form.Item<FieldType>
            label={t('components.input.rental.label')}
            name="rental"
            rules={[
              { required: true, message: t('components.input.rental.error') }
            ]}
          >
            <Input placeholder={t('components.input.rental.placeholder')} />
          </Form.Item>

          <Form.Item<FieldType>
            label={t('components.input.deposit.label')}
            name="deposit"
            rules={[
              { required: true, message: t('components.input.deposit.error') }
            ]}
          >
            <Input placeholder={t('components.input.deposit.placeholder')} />
          </Form.Item>

          <Form.Item<FieldType>
            label={t('components.input.picture.label')}
            name="picture"
          >
            <Upload {...uploadProps}>
              <Button icon={<UploadOutlined />}>
                {t('components.input.picture.placeholder')}
              </Button>
            </Upload>
          </Form.Item>
        </Form>
      </div>
    </Modal>
  )
}

export default RealPropertyCreate
