import { useEffect, useState } from 'react'
import { Damage, UseDamagesReturn } from '@/interfaces/Property/Damage/Damage'
import GetPropertyDamages from '@/services/api/Owner/Properties/GetPropertyDamages'

const useDamages = (propertyId: string, status: string): UseDamagesReturn => {
  const [damages, setDamages] = useState<Damage[] | null>(null)
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

  useEffect(() => {
    if (propertyId && status !== 'available') {
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
