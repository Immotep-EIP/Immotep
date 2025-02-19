import { openDB } from 'idb'
import { PropertyDetails } from '@/interfaces/Property/Property'

const DB_NAME = 'immotep-properties-db'
const STORE_NAME = 'properties'

export const getDB = async () =>
  openDB(DB_NAME, 1, {
    upgrade(db) {
      if (!db.objectStoreNames.contains(STORE_NAME)) {
        db.createObjectStore(STORE_NAME, { keyPath: 'id' })
      }
    }
  })

export const savePropertiesToDB = async (properties: PropertyDetails[]) => {
  const db = await getDB()
  const tx = db.transaction(STORE_NAME, 'readwrite')
  const store = tx.objectStore(STORE_NAME)
  properties.forEach(property => store.put(property))
  await tx.done
}

export const getPropertiesFromDB = async (): Promise<PropertyDetails[]> => {
  const db = await getDB()
  const tx = db.transaction(STORE_NAME, 'readonly')
  const store = tx.objectStore(STORE_NAME)
  return store.getAll()
}

export const removePropertyFromDB = async (propertyId: string) => {
  const db = await getDB()
  const tx = db.transaction(STORE_NAME, 'readwrite')
  const store = tx.objectStore(STORE_NAME)
  store.delete(propertyId)
  await tx.done
}

export const updatePropertyInDB = async (property: PropertyDetails) => {
  const db = await getDB()
  const tx = db.transaction(STORE_NAME, 'readwrite')
  const store = tx.objectStore(STORE_NAME)
  const existingProperty = await store.get(property.id)
  if (!existingProperty) {
    throw new Error(`Property with id ${property.id} not found`)
  }
  const updatedProperty = { ...existingProperty, ...property }
  await store.put(updatedProperty)
  await tx.done
}
