import { useState } from 'react'
import { UploadProps, message, UploadFile } from 'antd'
import fileToBase64 from '@/utils/base64/fileToBase'

const useImageUpload = () => {
  const [fileList, setFileList] = useState<UploadFile[]>([])
  const [imageBase64, setImageBase64] = useState<string | null>(null)

  const resetImage = () => {
    setImageBase64(null)
    setFileList([])
  }

  const uploadProps: UploadProps = {
    name: 'propertyPicture',
    maxCount: 1,
    fileList,
    accept: '.png, .jpg, .jpeg',
    beforeUpload: async file => {
      const base64 = await fileToBase64(file)
      setImageBase64(base64)
      return false
    },
    onChange(info) {
      setFileList(info.fileList)
      if (info.file.status === 'done') {
        message.success(`${info.file.name} file uploaded successfully`)
      } else if (info.file.status === 'error') {
        message.error(`${info.file.name} file upload failed.`)
      }
    }
  }

  return {
    uploadProps,
    imageBase64,
    setImageBase64,
    fileList,
    setFileList,
    resetImage
  }
}

export default useImageUpload
