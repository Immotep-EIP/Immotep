import React from 'react'
import { useTranslation } from 'react-i18next'
import PageMeta from '@/components/PageMeta/PageMeta'
import { Empty, Typography } from 'antd'
import style from './Messages.module.css'

const Messages: React.FC = () => {
  const { t } = useTranslation();

  return (
    <>
      <PageMeta
        title={t('pages.messages.document_title')}
        description={t('pages.messages.document_description')}
        keywords="messages, communication, Immotep"
      />
      <div className={style.layoutContainer}>
        <div className={style.emptyContainer}>
          <Empty
            image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
            description={
              <Typography.Text>
                {t('components.messages.no_properties_so_no_messages')}
              </Typography.Text>
            }
          />
        </div>
      </div>
    </>
  )
}

export default Messages
