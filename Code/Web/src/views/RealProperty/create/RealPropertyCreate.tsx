import React, { useState } from 'react'

import {
  FormProps,
  Form,
  Input,
  Button,
  Upload,
  UploadProps,
  message
} from 'antd'
import { useTranslation } from 'react-i18next'

import { UploadOutlined } from '@ant-design/icons'
import closeIcon from '@/assets/icons/close.png'

import CreatePropertyFunction from '@/services/api/Owner/Properties/CreateProperty'
import style from './RealPropertyCreate.module.css'

type FieldType = {
  name: string
  address: string
  zipCode: string
  city: string
  country: string
  area: string
  rental: string
  deposit: string
  picture: string
}

const props: UploadProps = {
  name: 'file',
  action: 'https://660d2bd96ddfa2943b33731c.mockapi.io/api/upload',
  maxCount: 1,
  headers: {
    authorization: 'authorization-text'
  },
  onChange(info) {
    if (info.file.status !== 'uploading') {
      console.log(info.file, info.fileList)
    }
    if (info.file.status === 'done') {
      message.success(`${info.file.name} file uploaded successfully`)
    } else if (info.file.status === 'error') {
      message.error(`${info.file.name} file upload failed.`)
    }
  }
}

const RealPropertyCreate: React.FC = () => {
  const { t } = useTranslation()
  const [loading, setLoading] = useState(false)

  const onFinish: FormProps<FieldType>['onFinish'] = (values: FieldType) => {
    const out = {
      name: values.name,
      address: values.address,
      city: values.city,
      postal_code: values.zipCode,
      country: values.country,
      area_sqm: parseFloat(values.area || '0'),
      rental_price_per_month: parseFloat(values.rental || '0'),
      deposit_price: parseFloat(values.deposit || '0'),
      picture: ''
    }
    const sendData = async () => {
      try {
        setLoading(true)
        const req = await CreatePropertyFunction(out)
        if (req) {
          setLoading(false)
          message.success(
            t('pages.real_property.add_real_property.property_created')
          )
        } else {
          setLoading(false)
          message.error(
            t('pages.real_property.add_real_property.error_property_created')
          )
        }
        window.history.back()
      } catch (error) {
        console.error('Error sending data:', error)
      }
    }
    sendData()
  }

  const onFinishFailed: FormProps<FieldType>['onFinishFailed'] = (
    errorInfo: any
  ) => {
    message.error(
      t('pages.real_property.add_real_property.fill_all_fields'),
      errorInfo
    )
  }

  return (
    <div className={style.pageContainer}>
      <div className={style.header}>
        <span className={style.title}>
          {t('pages.real_property.add_real_property.title')}
        </span>
        <Button
          shape="circle"
          style={{ margin: '20px', width: '40px', height: '40px' }}
          onClick={() => window.history.back()}
        >
          <img
            src={closeIcon}
            alt="close"
            style={{ width: '30px', height: '30px' }}
          />
        </Button>
      </div>
      <Form
        name="basic"
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
          rules={[{ required: false }]}
        >
          <Upload {...props}>
            <Button icon={<UploadOutlined />}>
              {t('components.input.picture.placeholder')}
            </Button>
          </Upload>
        </Form.Item>

        <div className={style.footer}>
          <Button
            type="primary"
            htmlType="submit"
            style={{ marginRight: '20px' }}
            loading={loading}
          >
            {t('components.button.add')}
          </Button>
        </div>
      </Form>
    </div>
  )
}

export default RealPropertyCreate
