import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'

import { Empty, Typography, Modal } from 'antd'
import { FilePdfOutlined, DeleteOutlined } from '@ant-design/icons'

import { Document } from '@/interfaces/Property/Document'

import style from './DocumentList.module.css'

const DocumentList: React.FC<{
  documents: Document[]
  onDocumentClick: (data: string) => void
  onDeleteDocument: (documentId: string) => void
}> = ({ documents, onDocumentClick, onDeleteDocument }) => {
  const { t } = useTranslation()
  const [documentToDelete, setDocumentToDelete] = useState<string | null>(null)

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

  return (
    <div className={style.documentsContainer}>
      {(!documents || documents.length === 0) && (
        <div className={style.noDocuments}>
          <Empty
            description={
              <Typography.Text>
                {t('pages.real_property_details.tabs.documents.no_documents')}
              </Typography.Text>
            }
          />
        </div>
      )}
      {documents?.map(document => (
        <div
          key={document.id}
          className={style.documentContainer}
          onClick={() => onDocumentClick(document.data)}
          role="button"
          tabIndex={0}
          onKeyDown={e => {
            if (e.key === 'Enter') {
              onDocumentClick(document.data)
            }
          }}
        >
          <div className={style.documentHeader}>
            <div className={style.documentDateContainer}>
              <span>{new Date(document.created_at).toLocaleDateString()}</span>
            </div>
            <DeleteOutlined
              className={style.deleteIcon}
              onClick={e => handleDelete(e, document.id)}
            />
          </div>
          <div className={style.documentPreviewContainer}>
            <FilePdfOutlined className={style.pdfIcon} />
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
      >
        <p>
          {t(
            'pages.real_property_details.tabs.documents.delete_confirmation.message'
          )}
        </p>
      </Modal>
    </div>
  )
}

export default DocumentList
