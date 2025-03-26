import fileToBase64 from '../fileToBase'

describe('fileToBase64', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })

  it('should convert file to base64 string', async () => {
    const mockBase64 = 'data:image/jpeg;base64,mockBase64String'
    const mockFile = new File(['mock'], 'test.jpg', { type: 'image/jpeg' })

    const mockFileReader = {
      readAsDataURL: jest.fn(),
      onload: null as any,
      result: mockBase64
    }
    global.FileReader = jest.fn(() => mockFileReader) as any

    const promise = fileToBase64(mockFile)

    mockFileReader.onload()

    const result = await promise
    expect(result).toBe(mockBase64)
  })

  it('should handle errors during conversion', async () => {
    const mockFile = new File(['mock'], 'test.jpg', { type: 'image/jpeg' })
    const mockError = new Error('Conversion failed')

    const mockFileReader = {
      readAsDataURL: jest.fn(),
      onerror: null as any,
      error: mockError
    }
    global.FileReader = jest.fn(() => mockFileReader) as any

    const promise = fileToBase64(mockFile)

    mockFileReader.onerror(mockError)

    await expect(promise).rejects.toEqual(mockError)
  })

  afterEach(() => {
    jest.restoreAllMocks()
  })
})
