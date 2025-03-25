import callApi from '@/services/api/apiCaller'
import { CreateInventoryReports } from '@/interfaces/Property/InventoryReports/InventoryReports'

const CreateInventoryReportByProperty = async (
  PropertyId: string,
  data: CreateInventoryReports
) => {
  try {
    return await callApi({
      method: 'POST',
      endpoint: `owner/properties/${PropertyId}/inventory-reports/`,
      data
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default CreateInventoryReportByProperty
