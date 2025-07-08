import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'

import { Form, message, Spin, Modal } from 'antd'

import { usePropertyContext } from '@/context/propertyContext'
import useDocument from '@/hooks/Property/useDocument'
import useLeasePermissions from '@/hooks/Property/useLeasePermissions'
import { Button, Empty } from '@/components/common'
import DocumentList from '@/components/features/RealProperty/details/tabs/Documents/DocumentList'
import UploadForm from '@/components/features/RealProperty/details/tabs/Documents/UploadForm'

import PropertyStatusEnum from '@/enums/PropertyEnum'

import style from './1DocumentsTab.module.css'

interface DocumentsTabProps {
  status?: string
}

const DocumentsTab: React.FC<DocumentsTabProps> = ({ status }) => {
  const { t } = useTranslation()
  const { property, selectedLeaseId, selectedLease } = usePropertyContext()
  const { canModify } = useLeasePermissions()

  let leaseIdToUse: string | undefined
  if (selectedLeaseId) {
    leaseIdToUse = selectedLeaseId
  } else if (property?.status === PropertyStatusEnum.UNAVAILABLE) {
    leaseIdToUse = 'current'
  } else {
    leaseIdToUse = undefined
  }

  const {
    documents,
    loading,
    error,
    refreshDocuments,
    uploadDocument,
    deleteDocument
  } = useDocument(property?.id || '', leaseIdToUse)

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
      <section
        className={style.loadingContainer}
        role="status"
        aria-live="polite"
        aria-labelledby="documents-loading-title"
      >
        <h2 id="documents-loading-title" className="sr-only">
          {t('pages.real_property_details.tabs.documents.loading_title')}
        </h2>
        <Spin size="large" />
      </section>
    )
  }

  if (
    (status === 'available' && !selectedLease) ||
    (status === 'invite sent' && !selectedLease)
  ) {
    return (
      <section
        className={style.tabContentEmpty}
        role="status"
        aria-live="polite"
        aria-labelledby="documents-empty-title"
      >
        <h2 id="documents-empty-title" className="sr-only">
          {t('pages.real_property_details.tabs.documents.empty_title')}
        </h2>
        <Empty description={t('pages.real_property.error.no_tenant_linked')} />
      </section>
    )
  }

  if (error) {
    return (
      <section
        className={style.errorContainer}
        role="alert"
        aria-live="assertive"
        aria-labelledby="documents-error-title"
      >
        <h2 id="documents-error-title" className="sr-only">
          {t('pages.real_property_details.tabs.documents.error_title')}
        </h2>
        <p>{error}</p>
      </section>
    )
  }

  return (
    <div className={style.tabContent} aria-labelledby="documents-tab-title">
      <h2 id="documents-tab-title" className="sr-only">
        {t('pages.real_property_details.tabs.documents.tab_title')}
      </h2>
      {canModify && (
        <div
          className={style.buttonAddContainer}
          aria-labelledby="documents-actions-title"
        >
          <h3 id="documents-actions-title" className="sr-only">
            {t('pages.real_property_details.tabs.documents.actions_title')}
          </h3>
          <Button
            onClick={showModal}
            aria-label={t(
              'pages.real_property_details.tabs.documents.add_document_aria'
            )}
          >
            {t('components.button.add_document')}
          </Button>
        </div>
      )}
      <Modal
        title={t('pages.real_property_details.tabs.documents.modal_title')}
        open={isModalOpen}
        onCancel={handleCancel}
        aria-labelledby="add-document-modal-title"
        aria-describedby="add-document-modal-description"
        footer={[
          <Button
            key="back"
            type="default"
            onClick={handleCancel}
            aria-label={t('components.button.cancel_add_document')}
          >
            {t('components.button.cancel')}
          </Button>,
          <Button
            key="submit"
            onClick={handleOk}
            aria-label={t('components.button.add_document_submit')}
          >
            {t('components.button.add')}
          </Button>
        ]}
      >
        <div id="add-document-modal-title" className="sr-only">
          {t('pages.real_property_details.tabs.documents.modal_title')}
        </div>
        <div id="add-document-modal-description" className="sr-only">
          {t('pages.real_property_details.tabs.documents.modal_description')}
        </div>
        <UploadForm form={form} />
      </Modal>
      <h3 id="documents-list-title" className="sr-only">
        {t('pages.real_property_details.tabs.documents.list_title')}
      </h3>
      <DocumentList
        documents={documents || []}
        onDocumentClick={handleDocumentClick}
        onDeleteDocument={deleteDocument}
      />
    </div>
  )
}

export default DocumentsTab
