import callApi from '@/services/api/apiCaller'

import { Damage } from '@/interfaces/Property/Damage/Damage'
import endpoints from '@/enums/EndPointEnum'

const GetPropertyDamageById = async (
  propertyId: string,
  damageId: string
): Promise<Damage> => {
  try {
    return await callApi<Damage>({
      method: 'GET',
      endpoint: endpoints.owner.properties.damages.byId(propertyId, damageId)
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default GetPropertyDamageById
