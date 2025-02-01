import { useState, useEffect } from 'react'
import base64ToFile from '@/utils/base64/baseToFile'

const useImageCache = (
  id: string | undefined,
  fetchImage: (id: string) => Promise<any>
) => {
  const [data, setData] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState<boolean>(true)

  const fetchData = async () => {
    if (!id) {
      setIsLoading(false)
      setData(null)
      return
    }

    setIsLoading(true)

    const cachedResponse = await caches.match(`/images/${id}`)
    if (cachedResponse) {
      const blob = await cachedResponse.blob()
      const url = URL.createObjectURL(blob)
      setData(url)
      setIsLoading(false)
      return
    }

    try {
      const response = await fetchImage(id)
      if (response) {
        const file = base64ToFile(response.data, 'image.jpg', 'image/jpeg')
        const fileUrl = URL.createObjectURL(file)
        setData(fileUrl)
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
      if (data) URL.revokeObjectURL(data)
    }
  }, [id, fetchImage])

  return { data, isLoading }
}

export default useImageCache
