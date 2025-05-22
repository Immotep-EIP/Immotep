import { act, renderHook } from '@testing-library/react'
import { message, UploadFile, UploadProps } from 'antd'
import type { RcFile } from 'antd/es/upload/interface'
import fileToBase64 from '@/utils/base64/fileToBase'
import useImageUpload from '../useImageUpload'

jest.mock('@/utils/base64/fileToBase', () => jest.fn())

jest.mock('antd', () => ({
  message: {
    success: jest.fn(),
    error: jest.fn()
  }
}))

describe('useImageUpload', () => {
  const mockFile = new File(['test'], 'test.jpg', {
    type: 'image/jpeg'
  }) as unknown as RcFile
  const mockBase64 = 'data:image/jpeg;base64,test'

  beforeEach(() => {
    jest.clearAllMocks()
    ;(fileToBase64 as jest.Mock).mockResolvedValue(mockBase64)
  })

  it('should initialize with empty state', () => {
    const { result } = renderHook(() => useImageUpload())

    expect(result.current.fileList).toEqual([])
    expect(result.current.imageBase64).toBeNull()
  })

  it('should handle file upload and convert to base64', async () => {
    const { result } = renderHook(() => useImageUpload())
    const { beforeUpload } = result.current.uploadProps

    await act(async () => {
      if (beforeUpload) {
        await beforeUpload(mockFile, [mockFile])
      }
    })

    expect(fileToBase64).toHaveBeenCalledWith(mockFile)
    expect(result.current.imageBase64).toBe(mockBase64)
  })

  it('should update fileList on change', () => {
    const { result } = renderHook(() => useImageUpload())
    const { onChange } = result.current.uploadProps
    const mockFileList: UploadFile[] = [
      {
        uid: '1',
        name: 'test.jpg',
        status: 'done' as const,
        url: 'test-url'
      }
    ]

    act(() => {
      if (onChange) {
        onChange({
          fileList: mockFileList,
          file: mockFileList[0]
        })
      }
    })

    expect(result.current.fileList).toEqual(mockFileList)
    expect(message.success).toHaveBeenCalledWith(
      'test.jpg file uploaded successfully'
    )
  })

  it('should handle upload error', () => {
    const { result } = renderHook(() => useImageUpload())
    const { onChange } = result.current.uploadProps
    const mockFileList: UploadFile[] = [
      {
        uid: '1',
        name: 'test.jpg',
        status: 'error' as const,
        url: 'test-url'
      }
    ]

    act(() => {
      if (onChange) {
        onChange({
          fileList: mockFileList,
          file: mockFileList[0]
        })
      }
    })

    expect(result.current.fileList).toEqual(mockFileList)
    expect(message.error).toHaveBeenCalledWith('test.jpg file upload failed.')
  })

  it('should reset image state', () => {
    const { result } = renderHook(() => useImageUpload())

    act(() => {
      result.current.setImageBase64(mockBase64)
      result.current.setFileList([{ uid: '1', name: 'test.jpg' } as UploadFile])
    })

    act(() => {
      result.current.resetImage()
    })

    expect(result.current.imageBase64).toBeNull()
    expect(result.current.fileList).toEqual([])
  })

  it('should have correct upload props configuration', () => {
    const { result } = renderHook(() => useImageUpload())

    expect(result.current.uploadProps).toEqual({
      name: 'propertyPicture',
      maxCount: 1,
      fileList: [],
      accept: '.png, .jpg, .jpeg',
      beforeUpload: expect.any(Function),
      onChange: expect.any(Function)
    })
  })
})
