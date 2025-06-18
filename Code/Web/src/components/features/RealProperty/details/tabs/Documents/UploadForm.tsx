import React from 'react'
import { UploadOutlined } from '@ant-design/icons'
import { useTranslation } from 'react-i18next'

import { Form, Upload, FormInstance } from 'antd'

import { Button, Input } from '@/components/common'

const UploadForm: React.FC<{ form: FormInstance }> = ({ form }) => {
  const { t } = useTranslation()

  return (
    <Form
      form={form}
      layout="vertical"
      name="add_document_form"
      initialValues={{ remember: true }}
    >
      <Form.Item
        label={t('components.input.document_name.label')}
        name="documentName"
        rules={[
          {
            required: true,
            message: t('components.input.document_name.error')
          }
        ]}
      >
        <Input
          placeholder={t('components.input.document_name.placeholder')}
          aria-label={t('components.input.document_name.placeholder')}
        />
      </Form.Item>

      <Form.Item
        label={t('components.input.document.label')}
        name="documentFile"
        valuePropName="fileList"
        getValueFromEvent={e => (Array.isArray(e) ? e : e?.fileList)}
        rules={[
          { required: true, message: t('components.input.document.error') }
        ]}
      >
        <Upload
          name="file"
          accept=".pdf,.docx,.xlsx"
          listType="text"
          beforeUpload={() => false}
          maxCount={1}
        >
          <Button type="default" icon={<UploadOutlined />}>
            {t('components.input.document.placeholder')}
          </Button>
        </Upload>
      </Form.Item>
    </Form>
  )
}

export default UploadForm
