import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'

import { Modal } from 'antd'
import { FilePdfOutlined, DeleteOutlined } from '@ant-design/icons'

import { Button, Empty } from '@/components/common'
import useLeasePermissions from '@/hooks/Property/useLeasePermissions'

import { Document } from '@/interfaces/Property/Document'

import style from './DocumentList.module.css'

const DocumentList: React.FC<{
  documents: Document[]
  onDocumentClick: (data: string) => void
  onDeleteDocument: (documentId: string) => void
}> = ({ documents, onDocumentClick, onDeleteDocument }) => {
  const { t } = useTranslation()
  const [documentToDelete, setDocumentToDelete] = useState<string | null>(null)
  const { canModify } = useLeasePermissions()

  const handleDelete = (e: React.MouseEvent, documentId: string) => {
    e.stopPropagation()
    setDocumentToDelete(documentId)
  }

  const handleConfirmDelete = () => {
    if (documentToDelete) {
      onDeleteDocument(documentToDelete)
      setDocumentToDelete(null)
    }
  }

  const handleCancelDelete = () => {
    setDocumentToDelete(null)
  }

  if (!documents || documents.length === 0) {
    return (
      <div className={style.noDocuments} role="status" aria-live="polite">
        <Empty
          description={t(
            'pages.real_property_details.tabs.documents.no_documents'
          )}
        />
      </div>
    )
  }

  return (
    <div className={style.documentsContainer}>
      {documents?.map(document => (
        <div
          key={document.id}
          className={style.documentContainer}
          onClick={() => onDocumentClick(document.data)}
          role="button"
          tabIndex={0}
          onKeyDown={e => {
            if (e.key === 'Enter') {
              e.preventDefault()
              onDocumentClick(document.data)
            }
          }}
          aria-label={`${t('components.button.download')} ${document.name} - ${new Date(document.created_at).toLocaleDateString()}`}
        >
          <div className={style.documentHeader}>
            <div className={style.documentDateContainer}>
              <time dateTime={document.created_at}>
                {new Date(document.created_at).toLocaleDateString()}
              </time>
            </div>
            {canModify && (
              <Button
                type="text"
                className={style.deleteIcon}
                onClick={e => handleDelete(e, document.id)}
                aria-label={`${t('components.button.delete')} ${document.name}`}
                title={`${t('components.button.delete')} ${document.name}`}
                style={{
                  border: 'none',
                  background: 'transparent',
                  cursor: 'pointer'
                }}
              >
                <DeleteOutlined aria-hidden="true" />
              </Button>
            )}
          </div>
          <div className={style.documentPreviewContainer} aria-hidden="true">
            <FilePdfOutlined className={style.pdfIcon} aria-hidden="true" />
          </div>
          <div className={style.documentName}>
            <span>{document.name}</span>
          </div>
        </div>
      ))}

      <Modal
        title={t(
          'pages.real_property_details.tabs.documents.delete_confirmation.title'
        )}
        open={!!documentToDelete}
        onOk={handleConfirmDelete}
        onCancel={handleCancelDelete}
        okText={t('components.button.confirm')}
        cancelText={t('components.button.cancel')}
        okButtonProps={{ danger: true }}
        aria-labelledby="delete-document-modal"
        aria-describedby="delete-document-description"
      >
        <p id="delete-document-description">
          {t(
            'pages.real_property_details.tabs.documents.delete_confirmation.message'
          )}
        </p>
      </Modal>
    </div>
  )
}

export default DocumentList
