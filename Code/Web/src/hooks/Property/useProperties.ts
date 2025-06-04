import { useEffect, useState } from 'react'

import GetProperties from '@/services/api/Owner/Properties/GetProperties.ts'
import CreatePropertyFunction from '@/services/api/Owner/Properties/CreateProperty'
import UpdatePropertyPicture from '@/services/api/Owner/Properties/UpdatePropertyPicture'
import GetPropertyDetails from '@/services/api/Owner/Properties/GetPropertyDetails'
import UpdatePropertyFunction from '@/services/api/Owner/Properties/UpdateProperty'
import GetLeasesByProperty from '@/services/api/Owner/Properties/Leases/GetLeasesByProperty'

import { Lease } from '@/interfaces/Property/Lease/Lease'

import {
  PropertyDetails,
  CreatePropertyPayload
} from '@/interfaces/Property/Property.tsx'

const useProperties = (
  propertyId: string | null = null,
  archive: boolean = false
) => {
  const [properties, setProperties] = useState<PropertyDetails[]>([])
  const [propertyDetails, setPropertyDetails] = useState<
    (PropertyDetails & { leases: Lease[] }) | null
  >(null)
  const [loading, setLoading] = useState<boolean>(true)
  const [error, setError] = useState<string | null>(null)

  const clearError = () => setError(null)

  // const extractBase64Content = (base64: string) => base64.split(',')[1]

  const createProperty = async (
    propertyData: CreatePropertyPayload,
    imageBase64: string | null
  ) => {
    setLoading(true)
    clearError()
    try {
      const createdProperty = await CreatePropertyFunction(propertyData)
      if (!createdProperty) throw new Error('Property creation failed.')

      if (imageBase64) {
        await UpdatePropertyPicture(createdProperty.id, imageBase64)
      }
      setProperties(prev => [...prev, { ...createdProperty, leases: [] }])
    } catch (err: unknown) {
      const errorMessage =
        err instanceof Error ? err.message : 'Unknown error occurred.'
      setError(errorMessage)
      throw err
    } finally {
      setLoading(false)
    }
  }

  const updateProperty = async (
    propertyData: CreatePropertyPayload,
    imageBase64: string | null,
    propertyId: string
  ) => {
    setLoading(true)
    clearError()
    try {
      const updatedProperty = await UpdatePropertyFunction(
        propertyData,
        propertyId
      )
      if (!updatedProperty) throw new Error('Property update failed.')
      if (imageBase64) {
        await UpdatePropertyPicture(updatedProperty.id, imageBase64)
      }
      setProperties(prev =>
        prev.map(prop =>
          prop.id === updatedProperty.id ? { ...updatedProperty } : prop
        )
      )
    } catch (err: unknown) {
      const errorMessage =
        err instanceof Error ? err.message : 'Unknown error occurred.'
      setError(errorMessage)
      throw err
    } finally {
      setLoading(false)
    }
  }

  const fetchProperties = async () => {
    try {
      setLoading(true)
      const res = await GetProperties(archive)
      setProperties(res)
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'Unknown error occurred.')
    } finally {
      setLoading(false)
    }
  }

  const refreshProperties = fetchProperties

  const getPropertyDetails = async (propertyId: string) => {
    try {
      setLoading(true)
      const res = await GetPropertyDetails(propertyId)
      const leases = await GetLeasesByProperty(res.id)
      setPropertyDetails({ ...res, leases })
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'Unknown error occurred.')
      throw err
    } finally {
      setLoading(false)
    }
  }

  const refreshPropertyDetails = getPropertyDetails

  useEffect(() => {
    if (!propertyId) {
      fetchProperties()
    } else {
      getPropertyDetails(propertyId)
    }
  }, [propertyId, archive])

  return {
    properties,
    propertyDetails,
    loading,
    error,
    createProperty,
    updateProperty,
    getPropertyDetails,
    refreshProperties,
    refreshPropertyDetails,
    clearError
  }
}

export default useProperties
