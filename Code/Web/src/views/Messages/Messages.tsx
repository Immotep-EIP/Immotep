import React from 'react'
import { useTranslation } from 'react-i18next'
import PageMeta from '@/components/PageMeta/PageMeta'
import style from './Messages.module.css'

const Messages: React.FC = () => {
  const { t } = useTranslation()

  return (
    <>
      <PageMeta
        title={t('pages.message.document_title')}
        description={t('pages.message.document_description')}
        keywords="messages, communication, Immotep"
      />
      <div className={style.layoutContainer}>Messages</div>
    </>
  )
}

export default Messages
