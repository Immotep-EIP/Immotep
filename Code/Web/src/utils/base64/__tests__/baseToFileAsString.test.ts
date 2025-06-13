import base64ToFileAsString from '../baseToFileAsString'

describe('base64ToFileAsString', () => {
  const mockURL = {
    createObjectURL: jest.fn()
  }
  const originalURL = global.URL

  beforeEach(() => {
    // Mock URL.createObjectURL
    global.URL = mockURL as any
    mockURL.createObjectURL.mockReturnValue('mock-file-url')
  })

  afterEach(() => {
    // Restore original URL
    global.URL = originalURL
    jest.clearAllMocks()
  })

  it('converts base64 string with data URL format to file URL', () => {
    const base64String = [
      'data:image/png;base64,',
      'iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg=='
    ].join('')
    const result = base64ToFileAsString(base64String)

    expect(mockURL.createObjectURL).toHaveBeenCalledWith(expect.any(Blob))
    expect(result).toBe('mock-file-url')

    // Verify Blob properties
    const blob = mockURL.createObjectURL.mock.calls[0][0]
    expect(blob.type).toBe('image/png')
  })

  it('handles base64 string without data URL format', () => {
    const base64String =
      'iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg=='
    const result = base64ToFileAsString(base64String)

    expect(mockURL.createObjectURL).toHaveBeenCalledWith(expect.any(Blob))
    expect(result).toBe('mock-file-url')

    // Verify Blob properties
    const blob = mockURL.createObjectURL.mock.calls[0][0]
    expect(blob.type).toBe('application/octet-stream')
  })

  it('handles empty base64 string', () => {
    const base64String = ''
    const result = base64ToFileAsString(base64String)

    expect(mockURL.createObjectURL).toHaveBeenCalledWith(expect.any(Blob))
    expect(result).toBe('mock-file-url')

    // Verify Blob properties
    const blob = mockURL.createObjectURL.mock.calls[0][0]
    expect(blob.type).toBe('application/octet-stream')
    expect(blob.size).toBe(0)
  })

  it('handles invalid base64 string', () => {
    const base64String = 'invalid-base64!@#$%'
    expect(() => base64ToFileAsString(base64String)).toThrow()
  })

  it('handles large base64 strings', () => {
    // Create a large base64 string (1MB)
    const largeBase64 = 'A'.repeat(1024 * 1024)
    const result = base64ToFileAsString(largeBase64)

    expect(mockURL.createObjectURL).toHaveBeenCalledWith(expect.any(Blob))
    expect(result).toBe('mock-file-url')

    // Verify Blob properties
    const blob = mockURL.createObjectURL.mock.calls[0][0]
    expect(blob.type).toBe('application/octet-stream')
  })
})
