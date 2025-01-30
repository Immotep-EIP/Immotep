import callApi from '@/services/api/apiCaller'
import { InventoryReports } from '@/interfaces/Property/InventoryReports/InventoryReports'

const GetInventoryReportById = async (PropertyId: string, ReportId: string) => {
  try {
    return await callApi<InventoryReports>({
      method: 'GET',
      endpoint: `owner/properties/${PropertyId}/inventory-reports/${ReportId}/`
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetInventoryReportById
