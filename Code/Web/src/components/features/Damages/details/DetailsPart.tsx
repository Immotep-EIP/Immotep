import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Image, Spin } from 'antd'
import { EyeOutlined } from '@ant-design/icons'
import { useLocation } from 'react-router-dom'
import DamageHeader from './DamageHeader'
import style from './DetailsPart.module.css'
import SubtitledElement from '@/components/ui/SubtitledElement/SubtitledElement'
import base64ToFileAsString from '@/utils/base64/baseToFileAsString'
import useDamages from '@/hooks/Property/useDamages'
import useProperties from '@/hooks/Property/useProperties'
import PriorityTag from '@/components/common/PriorityTag'
import DamageStatusTag from '@/components/common/DamageStatusTag'
// import UpdateDamage from '@/services/api/Owner/Properties/UpdateDamage'

const DetailsPart: React.FC = () => {
  const { t } = useTranslation()
  const location = useLocation()
  const { id, damageId } = location.state || {}
  const [refreshTrigger, setRefreshTrigger] = useState(0)
  const { propertyDetails: propertyData } = useProperties(id)
  const { damage, loading, error } = useDamages(
    id || '',
    propertyData?.status || '',
    damageId || '',
    refreshTrigger
  )

  const handleRefresh = () => {
    setRefreshTrigger(prev => prev + 1)
  }

  // useEffect(() => {
  //   const setDamageAsRead = async () => {
  //     if (damage && damage.read === 'false') {
  //       await UpdateDamage({ read: 'true' }, id, damageId)
  //     }
  //   }
  //   setDamageAsRead()
  // }, [damage, id, damageId])

  return (
    <div className={style.mainContainer}>
      {loading ||
        (!damage && (
          <div className={style.loadingContainer}>
            <Spin size="large" />
          </div>
        ))}
      {error && <div>{t('components.error', { message: error })}</div>}
      {damage && (
        <>
          <DamageHeader
            propertyId={id}
            propertyStatus={propertyData?.status || ''}
            damageId={damageId}
            onDataUpdated={handleRefresh}
          />
          <div className={style.headerInformationContainer}>
            <div className={style.damageInfosContainer}>
              <div className={style.rowContainer}>
                <SubtitledElement
                  subtitleKey={t('pages.damage_details.priority')}
                  subTitleStyle={{ marginBottom: '0.5rem' }}
                >
                  <PriorityTag priority={damage?.priority} />
                </SubtitledElement>
                <SubtitledElement
                  subtitleKey={t('pages.damage_details.fix_status')}
                  subTitleStyle={{ marginBottom: '0.5rem' }}
                >
                  <DamageStatusTag status={damage.fix_status} />
                </SubtitledElement>
              </div>

              <div className={style.rowContainer}>
                <SubtitledElement
                  subtitleKey={t('pages.damage_details.tenant_name')}
                  subTitleStyle={{ marginBottom: '0.5rem' }}
                >
                  {damage.tenant_name}
                </SubtitledElement>
                <SubtitledElement
                  subtitleKey={t('pages.damage_details.created_at')}
                  subTitleStyle={{ marginBottom: '0.5rem' }}
                >
                  {damage?.created_at
                    ? new Date(damage.created_at).toLocaleDateString()
                    : ''}
                </SubtitledElement>
                <SubtitledElement
                  subtitleKey={t('pages.damage_details.fix_planned_at')}
                  subTitleStyle={{ marginBottom: '0.5rem' }}
                >
                  {damage.fix_planned_at
                    ? (() => {
                        const date = new Date(damage.fix_planned_at)
                        return `${date.toLocaleDateString()} ${t('pages.damage_details.at')} ${date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`
                      })()
                    : '-'}
                </SubtitledElement>
              </div>
              <div className={style.rowContainer}>
                <SubtitledElement
                  subtitleKey={t('pages.damage_details.room_name')}
                  subTitleStyle={{ marginBottom: '0.5rem' }}
                >
                  {damage?.room_name ||
                    t('pages.real_property_details.tabs.damage.unknown_room')}
                </SubtitledElement>
                <SubtitledElement
                  subtitleKey={t('pages.damage_details.comment')}
                  subTitleStyle={{ marginBottom: '0.5rem' }}
                >
                  {damage?.comment}
                </SubtitledElement>
              </div>
            </div>
            <SubtitledElement
              subtitleKey={t('pages.damage_details.pictures')}
              subTitleStyle={{ marginBottom: '0.5rem' }}
            >
              <div className={style.pictureContainer}>
                {(!damage?.pictures || damage.pictures.length === 0) &&
                  t('pages.real_property_details.tabs.damage.no_pictures')}
                {damage?.pictures?.map(picture => (
                  <div key={picture} className={style.pictureWrapper}>
                    <Image
                      height="100%"
                      width={150}
                      src={base64ToFileAsString(picture)}
                      className={style.picture}
                      preview={{
                        mask: <EyeOutlined />,
                        style: { borderRadius: '15px' }
                      }}
                    />
                  </div>
                ))}
              </div>
            </SubtitledElement>
          </div>
        </>
      )}
    </div>
  )
}

export default DetailsPart
