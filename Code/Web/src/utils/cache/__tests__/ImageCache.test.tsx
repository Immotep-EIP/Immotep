import { openDB } from 'idb'

import imageCache from '../ImageCache'

jest.mock('idb', () => ({
  openDB: jest.fn()
}))

describe('ImageCache', () => {
  const mockDB = {
    get: jest.fn(),
    put: jest.fn(),
    clear: jest.fn(),
    delete: jest.fn(),
    objectStoreNames: {
      contains: jest.fn()
    },
    createObjectStore: jest.fn()
  }

  beforeEach(() => {
    jest.clearAllMocks()
    ;(openDB as jest.Mock).mockResolvedValue(mockDB)
  })

  describe('init', () => {
    it('initializes the database with correct configuration', async () => {
      await imageCache.init()

      expect(openDB).toHaveBeenCalledWith('imageCache', 1, {
        upgrade: expect.any(Function)
      })
    })

    it('creates object store if it does not exist', async () => {
      const upgradeDB = {
        objectStoreNames: {
          contains: jest.fn().mockReturnValue(false)
        },
        createObjectStore: jest.fn()
      }

      await imageCache.init()

      const upgradeFn = (openDB as jest.Mock).mock.calls[0][2].upgrade
      upgradeFn(upgradeDB)

      expect(upgradeDB.objectStoreNames.contains).toHaveBeenCalledWith('images')
      expect(upgradeDB.createObjectStore).toHaveBeenCalledWith('images', {
        keyPath: 'id'
      })
    })
  })

  describe('getImage', () => {
    it('returns null if image is not found', async () => {
      mockDB.get.mockResolvedValue(undefined)

      const result = await imageCache.getImage('test-id')

      expect(result).toBeNull()
      expect(mockDB.get).toHaveBeenCalledWith('images', 'test-id')
    })

    it('returns blob if image is found', async () => {
      const mockBlob = new Blob(['test'])
      mockDB.get.mockResolvedValue({ blob: mockBlob })

      const result = await imageCache.getImage('test-id')

      expect(result).toBe(mockBlob)
      expect(mockDB.get).toHaveBeenCalledWith('images', 'test-id')
    })
  })

  describe('setImage', () => {
    it('stores image blob with correct data', async () => {
      const mockBlob = new Blob(['test'])
      jest.spyOn(Date, 'now').mockReturnValue(123456)

      await imageCache.setImage('test-id', mockBlob)

      expect(mockDB.put).toHaveBeenCalledWith('images', {
        id: 'test-id',
        blob: mockBlob,
        timestamp: 123456
      })
    })

    it('initializes db if not initialized', async () => {
      ;(imageCache as any).db = null
      ;(openDB as jest.Mock).mockClear()
      const mockBlob = new Blob(['test'])

      await imageCache.setImage('test-id', mockBlob)

      expect(openDB).toHaveBeenCalledTimes(1)
    })

    it('handles errors during image storage', async () => {
      const mockBlob = new Blob(['test'])
      const error = new Error('Storage error')
      mockDB.put.mockRejectedValue(error)

      await expect(imageCache.setImage('test-id', mockBlob)).rejects.toThrow(
        'Storage error'
      )
    })
  })

  describe('clearCache', () => {
    it('clears all images from cache', async () => {
      await imageCache.clearCache()

      expect(mockDB.clear).toHaveBeenCalledWith('images')
    })

    it('initializes db if not initialized', async () => {
      ;(imageCache as any).db = null
      ;(openDB as jest.Mock).mockClear()

      await imageCache.clearCache()

      expect(openDB).toHaveBeenCalledTimes(1)
    })

    it('handles errors during cache clearing', async () => {
      const error = new Error('Clear error')
      mockDB.clear.mockRejectedValue(error)

      await expect(imageCache.clearCache()).rejects.toThrow('Clear error')
    })
  })

  describe('removeImage', () => {
    it('removes specific image from cache', async () => {
      await imageCache.removeImage('test-id')

      expect(mockDB.delete).toHaveBeenCalledWith('images', 'test-id')
    })

    it('initializes db if not initialized', async () => {
      ;(imageCache as any).db = null
      ;(openDB as jest.Mock).mockClear()

      await imageCache.removeImage('test-id')

      expect(openDB).toHaveBeenCalledTimes(1)
    })

    it('handles errors during image removal', async () => {
      const error = new Error('Delete error')
      mockDB.delete.mockRejectedValue(error)

      await expect(imageCache.removeImage('test-id')).rejects.toThrow(
        'Delete error'
      )
    })

    it('handles non-existent image removal gracefully', async () => {
      mockDB.delete.mockResolvedValue(undefined)

      await expect(
        imageCache.removeImage('non-existent-id')
      ).resolves.not.toThrow()
    })
  })

  describe('error handling', () => {
    it('initializes db if not initialized when calling methods', async () => {
      ;(openDB as jest.Mock).mockClear()
      ;(imageCache as any).db = null

      await imageCache.getImage('test-id')

      expect(openDB).toHaveBeenCalledTimes(1)
      expect(openDB).toHaveBeenCalledWith('imageCache', 1, {
        upgrade: expect.any(Function)
      })
    })

    it('handles database operation errors gracefully', async () => {
      const error = new Error('Database error')
      mockDB.get.mockRejectedValue(error)

      await expect(imageCache.getImage('test-id')).rejects.toThrow(
        'Database error'
      )
    })
  })
})
