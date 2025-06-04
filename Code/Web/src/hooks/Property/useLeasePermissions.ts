import { usePropertyContext } from '@/context/propertyContext'

const useLeasePermissions = () => {
  const { property, selectedLeaseId } = usePropertyContext()

  const selectedLease = property?.leases.find(
    lease => lease.id === selectedLeaseId
  )
  const isLeaseActive = selectedLease?.active ?? false

  const canModify = !selectedLeaseId || isLeaseActive

  return {
    canModify,
    isLeaseActive,
    selectedLease
  }
}

export default useLeasePermissions
