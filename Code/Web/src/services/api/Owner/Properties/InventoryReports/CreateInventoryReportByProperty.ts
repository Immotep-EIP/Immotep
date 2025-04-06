import callApi from '@/services/api/apiCaller'
import { CreateInventoryReports } from '@/interfaces/Property/InventoryReports/InventoryReports'
import endpoints from '@/enums/EndPointEnum'

const CreateInventoryReportByProperty = async (
  PropertyId: string,
  data: CreateInventoryReports
) => {
  try {
    return await callApi({
      method: 'POST',
      endpoint: endpoints.owner.properties.inventoryReports.create(PropertyId),
      data
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default CreateInventoryReportByProperty
