import { openDB, IDBPDatabase } from 'idb'
import { ImageDBSchema } from '@/interfaces/Property/Property'

class ImageCache {
  private db: IDBPDatabase<ImageDBSchema> | null = null

  async init() {
    this.db = await openDB<ImageDBSchema>('imageCache', 1, {
      upgrade(db) {
        if (!db.objectStoreNames.contains('images')) {
          db.createObjectStore('images', { keyPath: 'id' })
        }
      }
    })
  }

  async getImage(id: string): Promise<Blob | null> {
    if (!this.db) await this.init()
    const image = await this.db!.get('images', id)
    return image?.blob || null
  }

  async setImage(id: string, blob: Blob): Promise<void> {
    if (!this.db) await this.init()
    await this.db!.put('images', {
      id,
      blob,
      timestamp: Date.now()
    })
  }

  async clearCache(): Promise<void> {
    if (!this.db) await this.init()
    await this.db!.clear('images')
  }

  async removeImage(id: string): Promise<void> {
    if (!this.db) await this.init()
    await this.db!.delete('images', id)
  }
}

const imageCache = new ImageCache()
export default imageCache
