import { useState, useEffect } from 'react'
import base64ToFile from '@/utils/base64/baseToFile'

const useImageCache = (
  id: string | undefined,
  fetchImage: (id: string) => Promise<any>
) => {
  const [data, setData] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState<boolean>(true)

  const updateCache = async (newImageData: string) => {
    if (!id) return

    const file = base64ToFile(
      newImageData.split(',')[1],
      'image.jpg',
      'image/jpeg'
    )

    const cache = await caches.open('immotep-cache-v1')
    await cache.put(`/images/${id}`, new Response(file))

    const url = URL.createObjectURL(file)
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

  return { data, isLoading, updateCache }
}

export default useImageCache
