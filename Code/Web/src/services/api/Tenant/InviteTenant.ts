import callApi from '@/services/api/apiCaller'
import {
  InviteTenantResponse,
  InviteTenant
} from '@/interfaces/Tenant/InviteTenant.ts'
import endpoints from '@/enums/EndPointEnum'

export const InviteTenants = async (tenantInfo: InviteTenant) => {
  try {
    return await callApi<InviteTenant, InviteTenantResponse>({
      method: 'POST',
      endpoint: endpoints.owner.properties.tenant.invite(tenantInfo.propertyId),
      data: tenantInfo
    })
  } catch (error) {
    console.error('Error fetching data:', error)
    throw error
  }
}

export default InviteTenants
