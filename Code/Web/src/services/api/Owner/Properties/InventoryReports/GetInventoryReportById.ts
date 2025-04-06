import callApi from '@/services/api/apiCaller'
import { InventoryReports } from '@/interfaces/Property/InventoryReports/InventoryReports'
import endpoints from '@/enums/EndPointEnum'

const GetInventoryReportById = async (PropertyId: string, ReportId: string) => {
  try {
    return await callApi<InventoryReports>({
      method: 'GET',
      endpoint: endpoints.owner.properties.inventoryReports.byId(
        PropertyId,
        ReportId
      )
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetInventoryReportById
