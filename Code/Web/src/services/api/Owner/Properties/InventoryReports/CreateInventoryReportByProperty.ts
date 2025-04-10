import callApi from '@/services/api/apiCaller'
import {
  CreateInventoryReportsPayload,
  InventoryReportsResponse
} from '@/interfaces/Property/InventoryReports/InventoryReports'
import endpoints from '@/enums/EndPointEnum'

const CreateInventoryReportByProperty = async (
  PropertyId: string,
  data: CreateInventoryReportsPayload
): Promise<InventoryReportsResponse> => {
  try {
    return await callApi<
      InventoryReportsResponse,
      CreateInventoryReportsPayload
    >({
      method: 'POST',
      endpoint: endpoints.owner.properties.inventoryReports.create(PropertyId),
      body: data
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default CreateInventoryReportByProperty
