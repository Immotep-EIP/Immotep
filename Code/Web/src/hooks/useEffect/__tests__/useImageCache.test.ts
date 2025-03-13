import { act, renderHook, waitFor } from '@testing-library/react'
import base64ToFile from '@/utils/base64/baseToFile'
import useImageCache from '../useImageCache'

jest.mock('@/utils/base64/baseToFile', () => jest.fn())

global.Response = jest.fn((body, _init) => ({
  blob: jest.fn().mockResolvedValue(body)
})) as any

describe('useImageCache', () => {
  const id = '123'
  const mockFetchImage = jest.fn()
  const base64Data = 'data:image/jpeg;base64,...'
  const mockUrl = 'blob:http://localhost/image123'
  const mockCachePut = jest.fn()

  beforeEach(() => {
    jest.clearAllMocks()

    jest.spyOn(console, 'error').mockImplementation(() => {})

    global.URL.createObjectURL = jest.fn(() => mockUrl)

    global.URL.revokeObjectURL = jest.fn()

    global.caches = {
      open: jest.fn().mockResolvedValue({
        put: mockCachePut
      }),
      match: jest.fn().mockResolvedValue(null),
      delete: jest.fn(),
      has: jest.fn(),
      keys: jest.fn()
    }
    ;(base64ToFile as jest.Mock).mockImplementation((data, filename, type) => {
      const base64 = data.startsWith('data:') ? data.split(',')[1] : data
      return new File([base64], filename, { type })
    })
  })

  it('should do nothing if id is undefined', async () => {
    const { result } = renderHook(() =>
      useImageCache(undefined, mockFetchImage)
    )

    await act(async () => {
      result.current.updateCache(base64Data)
    })

    expect(base64ToFile).not.toHaveBeenCalled()
    expect(global.caches.open).not.toHaveBeenCalled()
    expect(global.URL.createObjectURL).not.toHaveBeenCalled()
  })

  it('should update the cache and set the new image URL if id is defined', async () => {
    const { result } = renderHook(() => useImageCache(id, mockFetchImage))

    await act(async () => {
      await result.current.updateCache(base64Data)
    })

    expect(base64ToFile).toHaveBeenCalledWith(
      base64Data.split(',')[1],
      'image.jpg',
      'image/jpeg'
    )

    expect(global.caches.open).toHaveBeenCalledWith('keyz-cache-v1')

    expect(mockCachePut).toHaveBeenCalledWith(
      `/images/${id}`,
      expect.objectContaining({
        blob: expect.any(Function)
      })
    )

    expect(global.URL.createObjectURL).toHaveBeenCalled()

    expect(result.current.data).toBeTruthy()
  })

  it('should fetch and cache the image if not in cache', async () => {
    mockFetchImage.mockResolvedValueOnce({ data: base64Data })

    const { result } = renderHook(() => useImageCache(id, mockFetchImage))

    await waitFor(() => {
      expect(result.current.isLoading).toBe(false)
    })

    expect(mockFetchImage).toHaveBeenCalledWith(id)

    expect(base64ToFile).toHaveBeenCalledWith(
      base64Data,
      'image.jpg',
      'image/jpeg'
    )

    expect(global.caches.open).toHaveBeenCalledWith('keyz-cache-v1')

    expect(result.current.data).toEqual(mockUrl)

    expect(result.current.isLoading).toBe(false)
  })

  it('should use cached image if available', async () => {
    const mockFile = new File(['mock data'], 'image.jpg', {
      type: 'image/jpeg'
    })
    global.caches.match = jest
      .fn()
      .mockResolvedValueOnce(new Response(mockFile))

    const { result } = renderHook(() => useImageCache(id, mockFetchImage))

    await waitFor(() => {
      expect(result.current.isLoading).toBe(false)
    })

    expect(mockFetchImage).not.toHaveBeenCalled()

    expect(global.caches.match).toHaveBeenCalledWith(`/images/${id}`)

    expect(result.current.data).toEqual(mockUrl)

    expect(result.current.isLoading).toBe(false)
  })

  it('should handle errors gracefully', async () => {
    mockFetchImage.mockRejectedValueOnce(new Error('Fetch failed'))

    const { result } = renderHook(() => useImageCache(id, mockFetchImage))

    await waitFor(() => {
      expect(result.current.isLoading).toBe(false)
    })

    expect(mockFetchImage).toHaveBeenCalledWith(id)
    expect(result.current.data).toBeNull()
    expect(result.current.isLoading).toBe(false)
  })

  it('should not fetch or cache if id is undefined', async () => {
    const { result } = renderHook(() =>
      useImageCache(undefined, mockFetchImage)
    )

    expect(result.current.data).toBeNull()
    expect(result.current.isLoading).toBe(false)
    expect(mockFetchImage).not.toHaveBeenCalled()
    expect(global.caches.match).not.toHaveBeenCalled()
  })

  it('should revoke the object URL on unmount', async () => {
    const revokeObjectURLMock = jest.fn()
    global.URL.revokeObjectURL = revokeObjectURLMock

    const { result, unmount } = renderHook(() =>
      useImageCache(id, mockFetchImage)
    )

    await act(async () => {
      result.current.updateCache(base64Data)
    })

    unmount()

    expect(revokeObjectURLMock).toHaveBeenCalledWith(mockUrl)
  })
})
