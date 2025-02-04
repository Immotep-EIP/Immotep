import React from 'react'
import { Button, Typography } from 'antd'
import { useTranslation } from 'react-i18next'

import useNavigation from '@/hooks/useNavigation/useNavigation'
import styles from './Lost.module.css'

const { Title, Text } = Typography

const Lost: React.FC = () => {
  const { t } = useTranslation()
  const { goToOverview } = useNavigation()

  return (
    <div className={styles.pageContainer}>
      <Title level={1} style={{ fontSize: '4rem', marginBottom: '1rem' }}>
        404
      </Title>
      <Text
        type="secondary"
        style={{ fontSize: '1.5rem', marginBottom: '2rem' }}
      >
        {t('pages.lost.page_not_found')}
      </Text>
      <Button type="primary" onClick={goToOverview}>
        {t('pages.lost.back_home')}
      </Button>
    </div>
  )
}

export default Lost
