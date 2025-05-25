import callApi from '@/services/api/apiCaller'

import endpoints from '@/enums/EndPointEnum'
import { DamageDetails } from '@/interfaces/Property/Damage/Damage'

const UpdateDamage = async (
  data: DamageDetails,
  propertyId: string,
  damageId: string
): Promise<DamageDetails> => {
  try {
    return await callApi<DamageDetails>({
      method: 'PUT',
      endpoint: endpoints.owner.properties.damages.update(propertyId, damageId),
      body: data
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default UpdateDamage
