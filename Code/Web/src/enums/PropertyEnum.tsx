const PropertyStatusEnum = {
  AVAILABLE: 'available',
  UNAVAILABLE: 'unavailable',
  INVITATION_SENT: 'invite sent'
}

const TenantStatusEnum = {
  available: {
    text: 'pages.real_property.status.available',
    color: 'green'
  },
  unavailable: {
    text: 'pages.real_property.status.unavailable',
    color: 'red'
  },
  'invite sent': {
    text: 'pages.real_property.status.invitation_sent',
    color: 'orange'
  }
}

export default PropertyStatusEnum
export { TenantStatusEnum }
