import { useTranslation } from 'react-i18next'
import { Form, InputNumber, Upload, UploadProps } from 'antd'
import { InboxOutlined } from '@ant-design/icons'
import { Input } from '@/components/common'
import { PropertyFormFieldsType } from '@/utils/types/propertyType'

const { Dragger } = Upload

const PropertyFormFields = ({ uploadProps }: { uploadProps: UploadProps }) => {
  const { t } = useTranslation()

  return (
    <fieldset style={{ border: 'none', padding: 0, margin: 0 }}>
      <legend className="sr-only">
        {t('components.form.property_details.legend')}
      </legend>

      <Form.Item<PropertyFormFieldsType>
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
          id="property-name"
          aria-required="true"
          aria-describedby="property-name-help"
        />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.apartment_number.label')}
        name="apartment_number"
      >
        <Input
          placeholder={t('components.input.apartment_number.placeholder')}
          id="apartment-number"
          aria-describedby="apartment-number-help"
        />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.address.label')}
        name="address"
        rules={[
          { required: true, message: t('components.input.address.error') }
        ]}
      >
        <Input
          placeholder={t('components.input.address.placeholder')}
          id="property-address"
          aria-required="true"
          aria-describedby="property-address-help"
        />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.zip_code.label')}
        name="postal_code"
        rules={[
          { required: true, message: t('components.input.zip_code.error') }
        ]}
      >
        <Input
          placeholder={t('components.input.zip_code.placeholder')}
          id="postal-code"
          aria-required="true"
          aria-describedby="postal-code-help"
        />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.city.label')}
        name="city"
        rules={[{ required: true, message: t('components.input.city.error') }]}
      >
        <Input
          placeholder={t('components.input.city.placeholder')}
          id="city"
          aria-required="true"
          aria-describedby="city-help"
        />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.country.label')}
        name="country"
        rules={[
          { required: true, message: t('components.input.country.error') }
        ]}
      >
        <Input
          placeholder={t('components.input.country.placeholder')}
          id="country"
          aria-required="true"
          aria-describedby="country-help"
        />
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
          id="area-sqm"
          aria-required="true"
          aria-describedby="area-sqm-help"
          aria-label={t('components.input.area.aria_label')}
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
          id="rental-price"
          aria-required="true"
          aria-describedby="rental-price-help"
          aria-label={t('components.input.rental.aria_label')}
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
          id="deposit-price"
          aria-required="true"
          aria-describedby="deposit-price-help"
          aria-label={t('components.input.deposit.aria_label')}
        />
      </Form.Item>

      <Form.Item<PropertyFormFieldsType>
        label={t('components.input.picture.label')}
        name="picture"
        help={t('components.input.picture.help')}
      >
        <Dragger
          {...uploadProps}
          aria-label={t('components.input.picture.aria_label')}
        >
          <p className="ant-upload-drag-icon">
            <InboxOutlined aria-hidden="true" />
          </p>
          <p className="ant-upload-text">
            {t('components.input.picture.placeholder')}
          </p>
          <p className="ant-upload-hint" id="upload-hint">
            {t('components.input.picture.hint')}
          </p>
        </Dragger>
      </Form.Item>
    </fieldset>
  )
}

export default PropertyFormFields
