import React, { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Image, Spin, Button, Modal, DatePicker, Space } from 'antd'
import { EyeOutlined, CalendarOutlined } from '@ant-design/icons'
import { useLocation } from 'react-router-dom'
import dayjs from 'dayjs'
import DamageHeader from './DamageHeader'
import style from './DetailsPart.module.css'
import SubtitledElement from '@/components/ui/SubtitledElement/SubtitledElement'
import base64ToFileAsString from '@/utils/base64/baseToFileAsString'
import useDamages from '@/hooks/Property/useDamages'
import useProperties from '@/hooks/Property/useProperties'
import PriorityTag from '@/components/common/PriorityTag'
import DamageStatusTag from '@/components/common/DamageStatusTag'
import UpdateDamage from '@/services/api/Owner/Properties/UpdateDamage'

const DetailsPart: React.FC = () => {
  const { t } = useTranslation()
  const location = useLocation()
  let { id, damageId } = location.state || {}
  if (!id || !damageId) {
    const pathParts = location.pathname.split('/')
    const lastPart = pathParts[pathParts.length - 1]
    const idFromUrl = pathParts[pathParts.length - 3]
    const damageIdFromUrl = lastPart.split('?')[0]
    if (!idFromUrl || !damageIdFromUrl) {
      throw new Error(
        'Property ID or Damage ID not found in location state or URL'
      )
    }
    id = idFromUrl
    damageId = damageIdFromUrl

    console.log('id', id)
    console.log('damageId', damageId)
  }
  const [refreshTrigger, setRefreshTrigger] = useState(0)
  const { propertyDetails: propertyData } = useProperties(id)
  const { damage, loading, error, updateDamage } = useDamages(
    id || '',
    propertyData?.status || '',
    damageId || '',
    refreshTrigger
  )

  const [isModalOpen, setIsModalOpen] = useState<boolean>(false)
  const [selectedDate, setSelectedDate] = useState<Date | null>(null)

  const handleRefresh = () => {
    setRefreshTrigger(prev => prev + 1)
  }

  const handleOpenModal = () => {
    setIsModalOpen(true)
    if (damage?.fix_planned_at) {
      setSelectedDate(new Date(damage.fix_planned_at))
    }
  }

  const handleCancel = () => {
    setIsModalOpen(false)
    setSelectedDate(null)
  }

  const handleOk = async () => {
    try {
      if (selectedDate) {
        const isoDate = selectedDate.toISOString()
        await updateDamage(id, damageId, {
          fix_planned_at: isoDate
        })
        handleRefresh()
        handleCancel()
      }
    } catch (error) {
      console.error('Error while setting intervention date:', error)
    }
  }

  const handleDateChange = (date: any) => {
    if (date) {
      const jsDate = date.toDate()
      setSelectedDate(jsDate)
    } else {
      setSelectedDate(null)
    }
  }

  useEffect(() => {
    const setDamageAsRead = async () => {
      await UpdateDamage({ read: true }, id, damageId)
    }
    setDamageAsRead()
  }, [damage, id, damageId])

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
          <DamageHeader />
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
                  <div className={style.dateWithButtonContainer}>
                    <span>
                      {damage.fix_planned_at
                        ? (() => {
                            const date = new Date(damage.fix_planned_at)
                            return `${date.toLocaleDateString()} ${t('pages.damage_details.at')} ${date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`
                          })()
                        : '-'}
                    </span>
                    <Button
                      type="primary"
                      icon={<CalendarOutlined />}
                      size="small"
                      onClick={handleOpenModal}
                      style={{ marginLeft: '0.5rem' }}
                    />
                  </div>
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

          <Modal
            title={
              damage.fix_planned_at
                ? t('components.button.modify_intervention_date')
                : t('components.button.add_intervention_date')
            }
            open={isModalOpen}
            onOk={handleOk}
            onCancel={handleCancel}
            okText={t('components.button.confirm')}
            cancelText={t('components.button.cancel')}
          >
            <Space
              direction="vertical"
              style={{ width: '100%', marginTop: 16 }}
            >
              <DatePicker
                onChange={handleDateChange}
                style={{ width: '100%' }}
                showTime
                value={selectedDate ? dayjs(selectedDate) : null}
              />
            </Space>
          </Modal>
        </>
      )}
    </div>
  )
}

export default DetailsPart
