import React from 'react'
import { Typography } from 'antd'
import { useTranslation } from 'react-i18next'

import useNavigation from '@/hooks/Navigation/useNavigation'
import { Button } from '@/components/common'
import PageMeta from '@/components/ui/PageMeta/PageMeta'

import styles from './Lost.module.css'

const { Title, Text } = Typography

const Lost: React.FC = () => {
  const { t } = useTranslation()
  const { goToOverview } = useNavigation()

  return (
    <>
      <PageMeta
        title={t('pages.lost.document_title')}
        description={t('pages.lost.document_description')}
        keywords="404, not found, Keyz"
      />
      <main
        className={styles.pageContainer}
        role="main"
        aria-labelledby="lost-page-title"
      >
        <h1 id="lost-page-title" className="sr-only">
          {t('pages.lost.page_title')}
        </h1>
        <Title
          level={1}
          style={{ fontSize: '4rem', marginBottom: '1rem' }}
          aria-hidden="true"
        >
          404
        </Title>
        <Text
          type="secondary"
          style={{ fontSize: '1.5rem', marginBottom: '2rem' }}
          role="alert"
          aria-live="polite"
        >
          {t('pages.lost.page_not_found')}
        </Text>
        <Button
          type="primary"
          onClick={goToOverview}
          aria-label={t('pages.lost.back_home_aria')}
        >
          {t('pages.lost.back_home')}
        </Button>
      </main>
    </>
  )
}

export default Lost
