import React from 'react'
import { useTranslation } from 'react-i18next'
import PageMeta from '@/components/ui/PageMeta/PageMeta'
import style from './DamageDetails.module.css'
import DetailsPart from '@/components/features/Damages/details/DetailsPart'

const DamageDetails: React.FC = () => {
  const { t } = useTranslation()

  return (
    <>
      <PageMeta
        title={t('pages.damage_details.document_title')}
        description={t('pages.damage_details.document_description')}
        keywords="damage details, damage info, Keyz"
      />
      <div className={style.pageContainer}>
        <DetailsPart />
      </div>
    </>
  )
}

export default DamageDetails
