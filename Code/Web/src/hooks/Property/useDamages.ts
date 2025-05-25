import { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'

import { message } from 'antd'

import GetPropertyDamages from '@/services/api/Owner/Properties/GetPropertyDamages'
import GetPropertyDamageById from '@/services/api/Owner/Properties/GetPropertyDamageById'
import UpdateDamage from '@/services/api/Owner/Properties/UpdateDamage'

import {
  Damage,
  DamageDetails,
  UseDamagesReturn
} from '@/interfaces/Property/Damage/Damage'
import PropertyStatusEnum from '@/enums/PropertyEnum'

const useDamages = (
  propertyId: string,
  status: string,
  damageId?: string,
  refreshTrigger: number = 0
): UseDamagesReturn => {
  const { t } = useTranslation()
  const [damages, setDamages] = useState<Damage[] | null>(null)
  const [damage, setDamage] = useState<Damage | null>(null)
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)

  const fetchDamages = async (propertyId: string) => {
    try {
      setLoading(true)
      setError(null)
      const response = await GetPropertyDamages(propertyId)
      setDamages(response)
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : 'An error occurred while fetching the documents'
      )
      setDamages(null)
    } finally {
      setLoading(false)
    }
  }

  const getDamageByID = async (propertyId: string, damageId: string) => {
    try {
      setLoading(true)
      setError(null)
      const response = await GetPropertyDamageById(propertyId, damageId)
      setDamage(response)
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : 'An error occurred while fetching the documents'
      )
      setDamage(null)
    } finally {
      setLoading(false)
    }
  }

  const updateDamage = async (
    propertyId: string,
    damageId: string,
    data: DamageDetails
  ) => {
    try {
      setLoading(true)
      setError(null)
      const response = await UpdateDamage(data, propertyId, damageId)
      if (response) {
        await getDamageByID(propertyId, damageId)
        message.success(t('pages.damage_details.update_success'))
      } else {
        message.error(t('pages.damage_details.update_error'))
      }
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : 'An error occurred while updating the damage'
      )
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    if (propertyId && status === PropertyStatusEnum.UNAVAILABLE) {
      fetchDamages(propertyId)
    }
  }, [propertyId, status])

  useEffect(() => {
    if (damageId && propertyId) {
      getDamageByID(propertyId, damageId)
    }
  }, [damageId, propertyId, refreshTrigger])

  return {
    damages,
    damage,
    loading,
    error,
    refreshDamages: fetchDamages,
    getDamageByID,
    updateDamage
  }
}

export default useDamages
