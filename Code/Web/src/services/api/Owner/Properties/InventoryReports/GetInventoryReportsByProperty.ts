import callApi from '@/services/api/apiCaller'
import { InventoryReports } from '@/interfaces/Property/InventoryReports/InventoryReports'
import endpoints from '@/enums/EndPointEnum'

const GetInventoryReportsByProperty = async (PropertyId: string) => {
  try {
    return await callApi<InventoryReports[]>({
      method: 'GET',
      endpoint: endpoints.owner.properties.inventoryReports.list(PropertyId)
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetInventoryReportsByProperty
