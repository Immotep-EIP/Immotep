import callApi from '@/services/api/apiCaller'
import { PropertyDetails } from '@/interfaces/Property/Property'
import endpoints from '@/enums/EndPointEnum'

const GetDashboard = async (): Promise<any> => {
  try {
    return await callApi<PropertyDetails[]>({
      method: 'GET',
      endpoint: endpoints.owner.dashboard.list()
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetDashboard
