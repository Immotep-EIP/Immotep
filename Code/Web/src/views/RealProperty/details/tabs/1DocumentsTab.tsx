import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Button, Modal, Form, message, Spin, Empty, Typography } from 'antd'
import { usePropertyId } from '@/context/propertyIdContext'
import useDocument from '@/hooks/Property/useDocument'
import UploadForm from '@/components/features/RealProperty/details/tabs/Documents/UploadForm'
import DocumentList from '@/components/features/RealProperty/details/tabs/Documents/DocumentList'
import style from './1DocumentsTab.module.css'

interface DocumentsTabProps {
  status?: string
}

const DocumentsTab: React.FC<DocumentsTabProps> = ({ status }) => {
  const { t } = useTranslation()
  const propertyId = usePropertyId()
  const { documents, loading, error, refreshDocuments, uploadDocument } =
    useDocument(propertyId || '', status || '')
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [form] = Form.useForm()

  const showModal = () => {
    setIsModalOpen(true)
  }

  const handleOk = () => {
    form
      .validateFields()
      .then(values => {
        const file = values.documentFile[0].originFileObj
        uploadDocument(file, values.documentName, propertyId || '', 'current')
          .then(() => {
            message.success(t('components.documents.success_add'))
            form.resetFields()
            setIsModalOpen(false)
            if (propertyId) {
              refreshDocuments(propertyId)
            }
          })
          .catch(error => {
            console.error('Upload failed:', error)
            message.error(t('components.documents.error_add'))
          })
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

  if (status === 'available') {
    return (
      <div className={style.tabContentEmpty}>
        <Empty
          image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
          styles={{
            image: {
              height: 60
            }
          }}
          description={
            <Typography.Text>
              {t('pages.real_property.error.no_tenant_linked')}
            </Typography.Text>
          }
        />
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
        <UploadForm form={form} />
      </Modal>
      <DocumentList
        documents={documents || []}
        onDocumentClick={handleDocumentClick}
      />
    </div>
  )
}

export default DocumentsTab
