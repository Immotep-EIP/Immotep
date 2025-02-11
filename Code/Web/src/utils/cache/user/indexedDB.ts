import { openDB } from 'idb'
import { User } from '@/interfaces/User/User'

const DB_NAME = 'immotep-user-db'
const STORE_NAME = 'user'

export const getDB = async () =>
  openDB(DB_NAME, 1, {
    upgrade(db) {
      if (!db.objectStoreNames.contains(STORE_NAME)) {
        db.createObjectStore(STORE_NAME, { keyPath: 'id' })
      }
    }
  })

export const saveUserToDB = async (user: User) => {
  const db = await getDB()
  const tx = db.transaction(STORE_NAME, 'readwrite')
  const store = tx.objectStore(STORE_NAME)
  store.put(user)
  await tx.done
}

export const getUserFromDB = async (): Promise<User> => {
  const db = await getDB()
  return db.get(STORE_NAME, 1)
}
