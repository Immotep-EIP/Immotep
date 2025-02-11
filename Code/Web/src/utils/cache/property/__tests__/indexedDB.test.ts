import { openDB } from 'idb'
import {
  savePropertiesToDB,
  getPropertiesFromDB,
  removePropertyFromDB,
  getDB
} from '@/utils/cache/property/indexedDB'
import { PropertyDetails } from '@/interfaces/Property/Property'

jest.mock('idb')

describe('indexDB Utilities', () => {
  const mockStore = {
    put: jest.fn(),
    getAll: jest.fn().mockResolvedValue([]),
    delete: jest.fn()
  }

  const mockTransaction = {
    objectStore: jest.fn().mockReturnValue(mockStore),
    done: Promise.resolve()
  }

  beforeEach(() => {
    ;(openDB as jest.Mock).mockResolvedValue({
      transaction: jest.fn().mockReturnValue(mockTransaction)
    })
    jest.clearAllMocks()
  })

  it('should create the database and object store if not already present', async () => {
    const mockCreateObjectStore = jest.fn()
    const mockDb = {
      objectStoreNames: {
        contains: jest.fn().mockReturnValue(false)
      },
      createObjectStore: mockCreateObjectStore
    }

    const openDBMock = openDB as jest.Mock
    openDBMock.mockImplementation((_name, _version, { upgrade }) => {
      upgrade(mockDb)
      return Promise.resolve(mockDb)
    })

    const db = await getDB()

    expect(openDBMock).toHaveBeenCalledWith(
      'immotep-properties-db',
      1,
      expect.any(Object)
    )

    expect(mockDb.objectStoreNames.contains).toHaveBeenCalledWith('properties')
    expect(mockCreateObjectStore).toHaveBeenCalledWith('properties', {
      keyPath: 'id'
    })

    expect(db).toBe(mockDb)
  })

  it('should save properties to the database', async () => {
    const properties: PropertyDetails[] = [
      {
        id: '1',
        owner_id: 'owner123',
        name: 'Property 1',
        address: '123 Main St',
        city: 'CityName',
        postal_code: '12345',
        country: 'CountryName',
        area_sqm: 100,
        rental_price_per_month: 1200,
        deposit_price: 2400,
        picture_id: 'pic123',
        created_at: '2023-02-09T12:00:00Z',
        nb_damage: 0,
        status: 'available',
        tenant: 'John Doe',
        start_date: '2023-02-10',
        end_date: '2023-12-10'
      }
    ]
    await savePropertiesToDB(properties)

    expect(mockStore.put).toHaveBeenCalledWith(properties[0])
    expect(mockTransaction.objectStore).toHaveBeenCalledWith('properties')
  })

  it('should get properties from the database', async () => {
    const result = await getPropertiesFromDB()

    expect(mockStore.getAll).toHaveBeenCalled()
    expect(result).toEqual([])
  })

  it('should remove a property from the database', async () => {
    const propertyId = '1'
    await removePropertyFromDB(propertyId)

    expect(mockStore.delete).toHaveBeenCalledWith(propertyId)
    expect(mockTransaction.objectStore).toHaveBeenCalledWith('properties')
  })
})
