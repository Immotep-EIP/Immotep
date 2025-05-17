import { useEffect, useState } from 'react'
import { UseDashboardReturn } from '@/interfaces/Dashboard/Dashboard'
import GetDashboard from '@/services/api/Owner/Properties/GetDashboard'

const useDashboard = (): UseDashboardReturn => {
  const [reminders, setReminders] = useState<UseDashboardReturn['reminders']>(
    []
  )
  const [properties, setProperties] = useState<
    UseDashboardReturn['properties'] | null
  >(null)
  const [openDamages, setOpenDamages] = useState<
    UseDashboardReturn['open_damages'] | null
  >(null)
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)

  const fetchDashboardData = async () => {
    try {
      setLoading(true)
      setError(null)
      const response = await GetDashboard()
      setReminders(response.reminders)
      setProperties(response.properties)
      setOpenDamages(response.open_damages)
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : 'An error occurred while fetching the dashboard data'
      )
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchDashboardData()
  }, [])

  return {
    reminders,
    properties,
    open_damages: openDamages,
    loading,
    error,
    refreshDashboard: fetchDashboardData
  }
}

export default useDashboard
