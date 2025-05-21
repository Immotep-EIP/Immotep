import { Form, Input, InputNumber, Upload, Button, UploadProps } from 'antd'
import { UploadOutlined } from '@ant-design/icons'
import { useTranslation } from 'react-i18next'
import { PropertyFormFieldsType } from '@/utils/types/propertyType'

const PropertyFormFields = ({ uploadProps }: { uploadProps: UploadProps }) => {
  const { t } = useTranslation()

  return (
    <>
      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.property_name.label')}
        name="name"
        rules={[
          { required: true, message: t('components.input.property_name.error') }
        ]}
      >
        <Input placeholder={t('components.input.property_name.placeholder')} />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.apartment_number.label')}
        name="apartment_number"
      >
        <Input
          placeholder={t('components.input.apartment_number.placeholder')}
        />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.address.label')}
        name="address"
        rules={[
          { required: true, message: t('components.input.address.error') }
        ]}
      >
        <Input placeholder={t('components.input.address.placeholder')} />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.zip_code.label')}
        name="postal_code"
        rules={[
          { required: true, message: t('components.input.zip_code.error') }
        ]}
      >
        <Input placeholder={t('components.input.zip_code.placeholder')} />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.city.label')}
        name="city"
        rules={[{ required: true, message: t('components.input.city.error') }]}
      >
        <Input placeholder={t('components.input.city.placeholder')} />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.country.label')}
        name="country"
        rules={[
          { required: true, message: t('components.input.country.error') }
        ]}
      >
        <Input placeholder={t('components.input.country.placeholder')} />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.area.label')}
        name="area_sqm"
        rules={[{ required: true, message: t('components.input.area.error') }]}
      >
        <InputNumber
          placeholder={t('components.input.area.placeholder')}
          min={0}
          style={{ width: '100%' }}
        />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.rental.label')}
        name="rental_price_per_month"
        rules={[
          { required: true, message: t('components.input.rental.error') }
        ]}
      >
        <InputNumber
          placeholder={t('components.input.rental.placeholder')}
          min={1}
          style={{ width: '100%' }}
        />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.deposit.label')}
        name="deposit_price"
        rules={[
          { required: true, message: t('components.input.deposit.error') }
        ]}
      >
        <InputNumber
          placeholder={t('components.input.deposit.placeholder')}
          min={1}
          style={{ width: '100%' }}
        />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.picture.label')}
        name="picture"
      >
        <Upload {...uploadProps}>
          <Button icon={<UploadOutlined />}>
            {t('components.input.picture.placeholder')}
          </Button>
        </Upload>
      </Form.Item>
    </>
  )
}

export default PropertyFormFields
