import callApi from '@/services/api/apiCaller'

import {
  InviteTenantResponse,
  InviteTenant
} from '@/interfaces/Tenant/InviteTenant.ts'
import endpoints from '@/enums/EndPointEnum'

export const InviteTenants = async (
  tenantInfo: InviteTenant
): Promise<InviteTenantResponse> => {
  try {
    if (!tenantInfo.propertyId) {
      throw new Error('Property ID is required')
    }
    return await callApi<InviteTenantResponse, InviteTenant>({
      method: 'POST',
      endpoint: endpoints.owner.properties.tenant.invite(tenantInfo.propertyId),
      body: tenantInfo
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default InviteTenants
