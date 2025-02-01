import callApi from '@/services/api/apiCaller'
import { InventoryReports } from '@/interfaces/Property/InventoryReports/InventoryReports'

const GetInventoryReportsByProperty = async (PropertyId: string) => {
  try {
    return await callApi<InventoryReports[]>({
      method: 'GET',
      endpoint: `owner/properties/${PropertyId}/inventory-reports/`
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetInventoryReportsByProperty
