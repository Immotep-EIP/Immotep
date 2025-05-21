import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import {
  Button,
  Dropdown,
  MenuProps,
  Modal,
  DatePicker,
  Space,
  DatePickerProps
} from 'antd'
import { MoreOutlined } from '@ant-design/icons'
import dayjs from 'dayjs'
import style from './DetailsPart.module.css'
import PageTitle from '@/components/ui/PageText/Title'
import returnIcon from '@/assets/icons/retour.svg'
import useDamages from '@/hooks/Property/useDamages'

interface DamageHeaderProps {
  propertyId: string
  propertyStatus: string
  damageId: string
  onDataUpdated?: () => void
}

const DamageHeader: React.FC<DamageHeaderProps> = ({
  propertyId,
  propertyStatus,
  damageId,
  onDataUpdated
}) => {
  const { t } = useTranslation()
  const { updateDamage } = useDamages(propertyId, propertyStatus, damageId)
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false)
  const [selectedDate, setSelectedDate] = useState<Date | null>(null)

  const handleOpenModal = () => {
    setIsModalOpen(true)
  }

  const handleCancel = () => {
    setIsModalOpen(false)
  }

  const handleOk = async () => {
    try {
      if (selectedDate) {
        const isoDate = selectedDate.toISOString()
        await updateDamage(propertyId, damageId, {
          fix_planned_at: isoDate
        })
        if (onDataUpdated) {
          onDataUpdated()
        }

        handleCancel()
      }
    } catch (error) {
      console.error('Error while formatting date:', error)
    }
  }

  const handleDateChange: DatePickerProps['onChange'] = date => {
    if (date) {
      const jsDate = date.toDate()
      setSelectedDate(jsDate)
    } else {
      setSelectedDate(null)
    }
  }

  const items: MenuProps['items'] = [
    {
      key: '1',
      label: t('components.button.add_intervention_date'),
      onClick: handleOpenModal
    }
  ]

  return (
    <div className={style.moreInfosContainer}>
      <div className={style.titleContainer}>
        <div
          className={style.returnButtonContainer}
          onClick={() => window.history.back()}
          tabIndex={0}
          role="button"
          onKeyDown={e => {
            if (e.key === 'Enter') {
              window.history.back()
            }
          }}
        >
          <img src={returnIcon} alt="Return" className={style.returnIcon} />
        </div>
        <PageTitle
          title={t('pages.damage_details.title')}
          size="title"
          margin={false}
        />
      </div>
      <Dropdown menu={{ items }} trigger={['click']} placement="bottomRight">
        <Button
          type="text"
          icon={<MoreOutlined />}
          className={style.actionButton}
        />
      </Dropdown>

      <Modal
        title={t('components.button.add_intervention_date')}
        open={isModalOpen}
        onOk={handleOk}
        onCancel={handleCancel}
        okText={t('components.button.confirm')}
        cancelText={t('components.button.cancel')}
      >
        <Space direction="vertical" style={{ width: '100%', marginTop: 16 }}>
          <DatePicker
            onChange={handleDateChange}
            style={{ width: '100%' }}
            showTime
            value={selectedDate ? dayjs(selectedDate) : null}
          />
        </Space>
      </Modal>
    </div>
  )
}

export default DamageHeader
