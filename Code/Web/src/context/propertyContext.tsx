import React, { createContext, useContext, useMemo, useState } from 'react'

import { PropertyDetails } from '@/interfaces/Property/Property'
import { Lease } from '@/interfaces/Property/Lease/Lease'

const PropertyContext = createContext<{
  property: (PropertyDetails & { leases: Lease[] }) | null
  selectedLeaseId: string
  setSelectedLeaseId: (id: string) => void
  selectedLease: Lease | null
}>({
  property: null,
  selectedLeaseId: '',
  setSelectedLeaseId: () => {},
  selectedLease: null
})

export const PropertyProvider: React.FC<{
  property: PropertyDetails & { leases: Lease[] }
  children: React.ReactNode
}> = ({ property, children }) => {
  const [selectedLeaseId, setSelectedLeaseId] = useState(property?.lease?.id)

  const selectedLease =
    property.leases.find(lease => lease.id === selectedLeaseId) ?? null

  const value = useMemo(
    () => ({
      property,
      selectedLeaseId,
      setSelectedLeaseId,
      selectedLease
    }),
    [property, selectedLeaseId, selectedLease]
  )

  return (
    <PropertyContext.Provider value={value}>
      {children}
    </PropertyContext.Provider>
  )
}

export const usePropertyContext = () => useContext(PropertyContext)
