import React, { useEffect, useState } from 'react'
import { Button, Empty, Typography, Switch } from 'antd'
import { useTranslation } from 'react-i18next'

import useProperties from '@/hooks/Property/useProperties'
import PageTitle from '@/components/ui/PageText/Title'
import CardPropertyLoader from '@/components/ui/Loader/CardPropertyLoader'
import PageMeta from '@/components/ui/PageMeta/PageMeta'
import CardComponent from '@/components/features/RealProperty/PropertyCard'
import RealPropertyCreate from './create/RealPropertyCreate'
import style from './RealProperty.module.css'

const RealPropertyPage: React.FC = () => {
  const { t } = useTranslation()
  const [showArchived, setShowArchived] = useState(false)
  const { properties, loading, error, refreshProperties } = useProperties(
    null,
    showArchived
  )
  const [showModalCreate, setShowModalCreate] = useState(false)
  const [isPropertyCreated, setIsPropertyCreated] = useState(false)

  useEffect(() => {
    if (isPropertyCreated) {
      refreshProperties()
      setIsPropertyCreated(false)
    }
  }, [isPropertyCreated, refreshProperties])

  if (error) {
    return <p>{t('pages.real_property.error.error_fetching_data')}</p>
  }

  return (
    <>
      <PageMeta
        title={t('pages.real_property.document_title')}
        description={t('pages.real_property.document_description')}
        keywords="real property, Property info, Keyz"
      />
      <div className={style.pageContainer}>
        <div className={style.pageHeader}>
          <PageTitle title={t('pages.real_property.title')} size="title" />
          <div className={style.headerActions}>
            <div className={style.archiveFilter}>
              <Switch
                checked={showArchived}
                onChange={setShowArchived}
                checkedChildren={t('components.switch.show_archived')}
                unCheckedChildren={t('components.switch.show_active')}
              />
            </div>
            <Button type="primary" onClick={() => setShowModalCreate(true)}>
              {t('components.button.add_real_property')}
            </Button>
          </div>
        </div>

        <div className={style.cardsContainer}>
          {loading && <CardPropertyLoader cards={12} />}
          {!loading && properties.length === 0 && (
            <div className={style.emptyContainer}>
              <Empty
                image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
                description={
                  <Typography.Text>
                    {t('components.messages.no_properties')}
                  </Typography.Text>
                }
              />
            </div>
          )}
          {properties.map(realProperty => (
            <CardComponent
              key={realProperty.id}
              realProperty={realProperty}
              t={t}
            />
          ))}
        </div>
        <RealPropertyCreate
          showModalCreate={showModalCreate}
          setShowModalCreate={setShowModalCreate}
          setIsPropertyCreated={setIsPropertyCreated}
        />
      </div>
    </>
  )
}

export default RealPropertyPage
