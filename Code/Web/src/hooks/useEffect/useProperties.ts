import { useEffect, useState } from 'react'
import GetProperties from '@/services/api/Owner/Properties/GetProperties.ts'
import { PropertyDetails } from '@/interfaces/Property/Property.tsx'
import CreatePropertyFunction from '@/services/api/Owner/Properties/CreateProperty'
import UpdatePropertyPicture from '@/services/api/Owner/Properties/UpdatePropertyPicture'

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

const useProperties = () => {
  const [properties, setProperties] = useState<PropertyDetails[]>([])
  const [loading, setLoading] = useState<boolean>(true)
  const [error, setError] = useState<string | null>(null)

  const createProperty = async (
    propertyData: CreatePropertyData,
    imageBase64: string | null
  ) => {
    setLoading(true)
    setError(null)
    try {
      const createdProperty = await CreatePropertyFunction(propertyData)
      if (createdProperty) {
        if (imageBase64) {
          await UpdatePropertyPicture(
            createdProperty.id,
            imageBase64.split(',')[1]
          )
        }
        setProperties(prevProperties => [...prevProperties, createdProperty])
      } else {
        throw new Error('Property creation failed.')
      }
    } catch (err: any) {
      setError(err.message)
      throw err
    } finally {
      setLoading(false)
    }
  }

  const fetchData = async () => {
    try {
      setLoading(true)
      const res = await GetProperties()
      if (res) {
        setProperties(res)
      } else {
        throw new Error('No data received')
      }
    } catch (err: any) {
      console.error('Error fetching properties:', err.message)
      setError(err.message)
      throw err
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchData()
  }, [])

  return { properties, loading, error, createProperty }
}

export default useProperties
