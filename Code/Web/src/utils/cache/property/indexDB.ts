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
  return db.getAll(STORE_NAME)
}

export const removePropertyFromDB = async (propertyId: string) => {
  const db = await getDB()
  const tx = db.transaction(STORE_NAME, 'readwrite')
  const store = tx.objectStore(STORE_NAME)
  store.delete(propertyId)
  await tx.done
}
