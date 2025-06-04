import React from 'react'
import { useTranslation } from 'react-i18next'

import { Empty, Typography } from 'antd'
import { FilePdfOutlined } from '@ant-design/icons'

import { Document } from '@/interfaces/Property/Document'

import style from './DocumentList.module.css'

const DocumentList: React.FC<{
  documents: Document[]
  onDocumentClick: (data: string) => void
}> = ({ documents, onDocumentClick }) => {
  const { t } = useTranslation()

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
  )
}

export default DocumentList
