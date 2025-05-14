const PropertyStatusEnum = {
  AVAILABLE: 'available',
  UNAVAILABLE: 'unavailable',
  INVITATION_SENT: 'invite sent'
}

const TenantStatusEnum = {
  [PropertyStatusEnum.AVAILABLE]: {
    text: 'pages.real_property.status.available',
    color: 'green'
  },
  [PropertyStatusEnum.UNAVAILABLE]: {
    text: 'pages.real_property.status.unavailable',
    color: 'red'
  },
  [PropertyStatusEnum.INVITATION_SENT]: {
    text: 'pages.real_property.status.invitation_sent',
    color: 'orange'
  }
}

export default PropertyStatusEnum
export { TenantStatusEnum }
