import base64ToFile from '@/utils/base64/baseToFile'

describe('base64ToFile', () => {
  it('should convert a base64 string to a File object', () => {
    const base64 = 'aGVsbG8gd29ybGQ='
    const filename = 'test.txt'
    const mimeType = 'text/plain'

    const file = base64ToFile(base64, filename, mimeType)

    expect(file).toBeInstanceOf(File)
    expect(file.name).toBe(filename)
    expect(file.type).toBe(mimeType)

    const reader = new FileReader()
    reader.onload = event => {
      if (event.target?.result) {
        expect(event.target.result).toEqual(
          new TextEncoder().encode('hello world').buffer
        )
      }
    }
    reader.readAsArrayBuffer(file)
  })

  it('should handle empty base64 strings', () => {
    const base64 = ''
    const filename = 'empty.txt'
    const mimeType = 'text/plain'

    const file = base64ToFile(base64, filename, mimeType)

    expect(file).toBeInstanceOf(File)
    expect(file.name).toBe(filename)
    expect(file.type).toBe(mimeType)

    const reader = new FileReader()
    reader.onload = event => {
      if (event.target?.result) {
        expect(event.target.result).toEqual(new ArrayBuffer(0))
      }
    }
    reader.readAsArrayBuffer(file)
  })

  it('should handle non-ASCII characters', () => {
    const base64 = '4pyTIMOgIGxhIG1vZGU='
    const filename = 'unicode.txt'
    const mimeType = 'text/plain'

    const file = base64ToFile(base64, filename, mimeType)

    expect(file).toBeInstanceOf(File)
    expect(file.name).toBe(filename)
    expect(file.type).toBe(mimeType)

    const reader = new FileReader()
    reader.onload = event => {
      if (event.target?.result) {
        expect(event.target.result).toEqual(
          new TextEncoder().encode('✓ à la mode').buffer
        )
      }
    }
    reader.readAsArrayBuffer(file)
  })
})
