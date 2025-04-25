import { useEffect, useState } from 'react'
import { Damage, UseDamagesReturn } from '@/interfaces/Property/Damage/Damage'
import GetPropertyDamages from '@/services/api/Owner/Properties/GetPropertyDamages'
import useProperties from './useProperties'

const useDamages = (propertyId: string): UseDamagesReturn => {
  const [damages, setDamages] = useState<Damage[] | null>(null)
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)
  const { propertyDetails: propertyData } = useProperties(propertyId)

  const fetchDamages = async (propertyId: string) => {
    try {
      if (
        !propertyId ||
        propertyId === '' ||
        propertyData?.status !== 'available'
      ) {
        setError('No tenant assigned to this property')
        return
      }
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

  useEffect(() => {
    if (propertyId) {
      fetchDamages(propertyId)
    }
  }, [propertyId])

  return {
    damages,
    loading,
    error,
    refreshDamages: fetchDamages
  }
}

export default useDamages
