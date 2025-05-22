import callApi from '@/services/api/apiCaller'

import { InventoryReports } from '@/interfaces/Property/InventoryReports/InventoryReports'
import endpoints from '@/enums/EndPointEnum'

const GetInventoryReportById = async (
  propertyId: string,
  reportId: string
): Promise<InventoryReports> => {
  try {
    return await callApi<InventoryReports>({
      method: 'GET',
      endpoint: endpoints.owner.properties.inventoryReports.byId(
        propertyId,
        reportId
      )
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetInventoryReportById
