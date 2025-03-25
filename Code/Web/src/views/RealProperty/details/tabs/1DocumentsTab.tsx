import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Button, Modal, Form, Input, Upload, message, Spin } from 'antd'
import { UploadOutlined, FilePdfOutlined } from '@ant-design/icons'
import { usePropertyId } from '@/context/propertyIdContext'
import useDocument from '@/hooks/useEffect/useDocument'
import style from './1DocumentsTab.module.css'

const DocumentsTab: React.FC = () => {
  const { t } = useTranslation()
  const propertyId = usePropertyId()
  const { documents, loading, error, refreshDocuments } = useDocument(
    propertyId || ''
  )
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [form] = Form.useForm()

  const showModal = () => {
    setIsModalOpen(true)
  }

  const handleOk = () => {
    form
      .validateFields()
      .then(values => {
        console.log('values', values)
        message.success(t('components.documents.success_add'))
        form.resetFields()
        setIsModalOpen(false)
        if (propertyId) {
          refreshDocuments(propertyId)
        }
      })
      .catch(errorInfo => {
        console.error('Validation Failed:', errorInfo)
      })
  }

  const handleCancel = () => {
    form.resetFields()
    setIsModalOpen(false)
  }

  const handleDocumentClick = (documentData: string) => {
    const newWindow = window.open()
    if (newWindow) {
      const iframe = newWindow.document.createElement('iframe')
      iframe.src = documentData
      iframe.style.width = '100%'
      iframe.style.height = '100vh'
      iframe.style.border = 'none'
      newWindow.document.body.appendChild(iframe)
    }
  }

  if (loading) {
    return (
      <div className={style.loadingContainer}>
        <Spin size="large" />
      </div>
    )
  }

  if (error) {
    return (
      <div className={style.errorContainer}>
        <p>{error}</p>
      </div>
    )
  }

  return (
    <div className={style.tabContent}>
      <div className={style.buttonAddContainer}>
        <Button type="primary" onClick={showModal}>
          {t('components.button.add_document')}
        </Button>
      </div>
      <Modal
        title={t('pages.real_property_details.tabs.documents.modal_title')}
        open={isModalOpen}
        onCancel={handleCancel}
        footer={[
          <Button key="back" onClick={handleCancel}>
            {t('components.button.cancel')}
          </Button>,
          <Button key="submit" type="primary" onClick={handleOk}>
            {t('components.button.add')}
          </Button>
        ]}
      >
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
            <Upload name="file" listType="text" beforeUpload={() => false}>
              <Button icon={<UploadOutlined />}>
                {t('components.input.document.placeholder')}
              </Button>
            </Upload>
          </Form.Item>
        </Form>
      </Modal>
      <div className={style.documentsContainer}>
        {(!documents || documents === null || documents.length === 0) && (
          <div className={style.noDocuments}>
            <p>
              {t('pages.real_property_details.tabs.documents.no_documents')}
            </p>
          </div>
        )}
        {documents?.map(document => (
          <div
            key={document.id}
            className={style.documentContainer}
            onClick={() => handleDocumentClick(document.data)}
            role="button"
            tabIndex={0}
            onKeyDown={e => {
              if (e.key === 'Enter') {
                handleDocumentClick(document.data)
              }
            }}
          >
            <div className={style.documentDateContainer}>
              <span>{new Date(document.created_at).toLocaleDateString()}</span>
            </div>
            <div className={style.documentPreviewContainer}>
              <FilePdfOutlined className={style.pdfIcon} />
            </div>
            <div className={style.documentName}>
              <span>{document.name}</span>
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default DocumentsTab
