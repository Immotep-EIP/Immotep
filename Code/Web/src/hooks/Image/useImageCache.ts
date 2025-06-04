import { useState, useEffect, useRef } from 'react'

import base64ToFileAsString from '@/utils/base64/baseToFileAsString'
import imageCache from '@/utils/cache/ImageCache'

const useImageCache = (
  id: string | undefined,
  fetchImage: (id: string) => Promise<any>
) => {
  const [data, setData] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState<boolean>(true)
  const objectUrlRef = useRef<string | null>(null)

  const updateCache = async (newImageData: string) => {
    if (!id) return

    const fileUrl = base64ToFileAsString(newImageData)
    const response = await fetch(fileUrl)
    const blob = await response.blob()
    const file = new File([blob], 'image.jpg', { type: blob.type })

    await imageCache.setImage(id, file)
    const url = URL.createObjectURL(file)
    objectUrlRef.current = url
    setData(url)
  }

  const fetchData = async () => {
    if (!id) {
      setIsLoading(false)
      setData(null)
      return
    }

    setIsLoading(true)

    try {
      const cachedBlob = await imageCache.getImage(id)

      if (cachedBlob) {
        const url = URL.createObjectURL(cachedBlob)
        objectUrlRef.current = url
        setData(url)
        setIsLoading(false)
        return
      }

      const response = await fetchImage(id)
      if (response && response.data) {
        const fileUrl = base64ToFileAsString(response.data)
        const fetchResponse = await fetch(fileUrl)
        const blob = await fetchResponse.blob()
        const file = new File([blob], 'image.jpg', { type: blob.type })

        await imageCache.setImage(id, file)

        const url = URL.createObjectURL(file)
        objectUrlRef.current = url
        setData(url)
      } else {
        setData(null)
      }
    } catch (error) {
      console.error('Error fetching image:', error)
      setData(null)
    } finally {
      setIsLoading(false)
    }
  }

  useEffect(() => {
    fetchData()

    return () => {
      if (objectUrlRef.current) {
        URL.revokeObjectURL(objectUrlRef.current)
      }
    }
  }, [id, fetchImage])

  return { data, isLoading, updateCache }
}

export default useImageCache
