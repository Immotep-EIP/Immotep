import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'

import { Modal, Form, message, Spin, Empty, Typography } from 'antd'

import { usePropertyContext } from '@/context/propertyContext'
import useDocument from '@/hooks/Property/useDocument'
import useLeasePermissions from '@/hooks/Property/useLeasePermissions'
import { Button } from '@/components/common'
import DocumentList from '@/components/features/RealProperty/details/tabs/Documents/DocumentList'
import UploadForm from '@/components/features/RealProperty/details/tabs/Documents/UploadForm'

import style from './1DocumentsTab.module.css'

interface DocumentsTabProps {
  status?: string
}

const DocumentsTab: React.FC<DocumentsTabProps> = ({ status }) => {
  const { t } = useTranslation()
  const { property, selectedLeaseId, selectedLease } = usePropertyContext()
  const { canModify } = useLeasePermissions()

  const {
    documents,
    loading,
    error,
    refreshDocuments,
    uploadDocument,
    deleteDocument
  } = useDocument(property?.id || '', selectedLeaseId || 'current')
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
        uploadDocument(file, values.documentName, property?.id || '')
          .then(() => {
            message.success(t('components.documents.success_add'))
            form.resetFields()
            setIsModalOpen(false)
            if (property?.id) {
              refreshDocuments(property.id)
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

  if (
    (status === 'available' && !selectedLease) ||
    (status === 'invite sent' && !selectedLease)
  ) {
    return (
      <div className={style.tabContentEmpty}>
        <Empty
          image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
          imageStyle={{
            height: 60
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
      {canModify && (
        <div className={style.buttonAddContainer}>
          <Button onClick={showModal}>
            {t('components.button.add_document')}
          </Button>
        </div>
      )}
      <Modal
        title={t('pages.real_property_details.tabs.documents.modal_title')}
        open={isModalOpen}
        onCancel={handleCancel}
        footer={[
          <Button key="back" type="default" onClick={handleCancel}>
            {t('components.button.cancel')}
          </Button>,
          <Button key="submit" onClick={handleOk}>
            {t('components.button.add')}
          </Button>
        ]}
      >
        <UploadForm form={form} />
      </Modal>
      <DocumentList
        documents={documents || []}
        onDocumentClick={handleDocumentClick}
        onDeleteDocument={deleteDocument}
      />
    </div>
  )
}

export default DocumentsTab
