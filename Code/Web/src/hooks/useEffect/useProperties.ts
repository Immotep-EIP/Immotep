import { useEffect, useState } from 'react'
import GetProperties from '@/services/api/Owner/Properties/GetProperties.ts'
import { PropertyDetails } from '@/interfaces/Property/Property.tsx'
import CreatePropertyFunction from '@/services/api/Owner/Properties/CreateProperty'
import UpdatePropertyPicture from '@/services/api/Owner/Properties/UpdatePropertyPicture'
import GetPropertyDetails from '@/services/api/Owner/Properties/GetPropertyDetails'
import UpdatePropertyFunction from '@/services/api/Owner/Properties/UpdateProperty'

type CreatePropertyData = Omit<
  PropertyDetails,
  | 'id'
  | 'owner_id'
  | 'picture_id'
  | 'created_at'
  | 'nb_damage'
  | 'status'
  | 'tenant'
  | 'start_date'
  | 'end_date'
>

const useProperties = (propertyId: string | null = null) => {
  const [properties, setProperties] = useState<PropertyDetails[]>([])
  const [propertyDetails, setPropertyDetails] =
    useState<PropertyDetails | null>(null)
  const [loading, setLoading] = useState<boolean>(true)
  const [error, setError] = useState<string | null>(null)

  const clearError = () => setError(null)

  const extractBase64Content = (base64: string) => base64.split(',')[1]

  const createProperty = async (
    propertyData: CreatePropertyData,
    imageBase64: string | null
  ) => {
    setLoading(true)
    clearError()
    try {
      const createdProperty = await CreatePropertyFunction(propertyData)
      if (!createdProperty) throw new Error('Property creation failed.')

      if (imageBase64) {
        await UpdatePropertyPicture(
          createdProperty.id,
          extractBase64Content(imageBase64)
        )
      }
      setProperties(prev => [...prev, createdProperty])
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
    propertyData: CreatePropertyData,
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
        await UpdatePropertyPicture(
          updatedProperty.id,
          extractBase64Content(imageBase64)
        )
      }
      setProperties(prev =>
        prev.map(prop =>
          prop.id === updatedProperty.id ? updatedProperty : prop
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
      const res = await GetProperties()
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
      setPropertyDetails(res)
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
  }, [propertyId])

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
