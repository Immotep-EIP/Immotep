import { useEffect, useState } from 'react'
import {
  DashboardOpenDamages,
  DashboardProperties,
  DashboardReminders,
  UseDashboardReturn
} from '@/interfaces/Dashboard/Dashboard'
import GetDashboard from '@/services/api/Owner/Properties/GetDashboard'

const useDashboard = (): UseDashboardReturn => {
  const [reminders, setReminders] = useState<DashboardReminders[] | null>(null)
  const [properties, setProperties] = useState<DashboardProperties | null>(null)
  const [openDamages, setOpenDamages] = useState<DashboardOpenDamages | null>(
    null
  )
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)

  const fetchDashboardData = async () => {
    try {
      setLoading(true)
      setError(null)
      const response = await GetDashboard()
      if (!response) {
        throw new Error('No data received from the API')
      }
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
