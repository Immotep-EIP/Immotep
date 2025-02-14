import { useEffect, useState } from 'react';
import GetProperties from '@/services/api/Owner/Properties/GetProperties.ts';
import { PropertyDetails } from '@/interfaces/Property/Property.tsx';
import CreatePropertyFunction from '@/services/api/Owner/Properties/CreateProperty';
import UpdatePropertyPicture from '@/services/api/Owner/Properties/UpdatePropertyPicture';
import GetPropertyDetails from '@/services/api/Owner/Properties/GetPropertyDetails';
import {
  savePropertiesToDB,
  getPropertiesFromDB,
} from '@/utils/cache/property/indexedDB';

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
>;

const useProperties = (propertyId: string | null = null) => {
  const [properties, setProperties] = useState<PropertyDetails[]>([]);
  const [propertyDetails, setPropertyDetails] = useState<PropertyDetails | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const createProperty = async (
    propertyData: CreatePropertyData,
    imageBase64: string | null
  ) => {
    setLoading(true);
    setError(null);
    try {
      const createdProperty = await CreatePropertyFunction(propertyData);
      if (createdProperty) {
        if (imageBase64) {
          await UpdatePropertyPicture(createdProperty.id, imageBase64.split(',')[1]);
        }
        await savePropertiesToDB([createdProperty]);
        setProperties((prevProperties) => [...prevProperties, createdProperty]);
      } else {
        throw new Error('Property creation failed.');
      }
    } catch (err: any) {
      setError(err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  const fetchProperties = async () => {
    try {
      setLoading(true);
      const cachedProperties = await getPropertiesFromDB();
      if (cachedProperties.length > 0) {
        setProperties(cachedProperties);
      } else {
        const res = await GetProperties();
        setProperties(res);
        await savePropertiesToDB(res);
      }
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const refreshProperties = async () => {
    await fetchProperties();
  };

  const getPropertyDetails = async (propertyId: string) => {
    try {
      setLoading(true);
      const cachedProperties = await getPropertiesFromDB();
      const cachedProperty = cachedProperties.find(
        (property) => property.id === propertyId
      );
      if (cachedProperty) {
        setPropertyDetails(cachedProperty);
      } else {
        const res = await GetPropertyDetails(propertyId);
        setPropertyDetails(res);
        await savePropertiesToDB([res]);
      }
    } catch (err: any) {
      console.error('Error fetching property details:', err.message);
      setError(err.message);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        await fetchProperties();
        if (propertyId) {
          await getPropertyDetails(propertyId);
        }
      } catch (err) {
        console.error(err);
      }
    };

    fetchData();
  }, [propertyId]);

  return {
    properties,
    propertyDetails,
    loading,
    error,
    createProperty,
    getPropertyDetails,
    refreshProperties,
  };
};

export default useProperties;