import React, { createContext, useContext } from 'react'

const PropertyIdContext = createContext<string | null>(null)

export const PropertyIdProvider: React.FC<{
  id: string
  children: React.ReactNode
}> = ({ id, children }) => (
  <PropertyIdContext.Provider value={id}>{children}</PropertyIdContext.Provider>
)

export const usePropertyId = () => useContext(PropertyIdContext)
