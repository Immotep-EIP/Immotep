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

  useEffect(() => {
    if (!id) {
      setIsLoading(false)
      return
    }

    const fetchData = async () => {
      setIsLoading(true)

      const cachedResponse = await caches.match(`/images/${id}`)
      if (cachedResponse) {
        const blob = await cachedResponse.blob()
        const url = URL.createObjectURL(blob)
        setData(url)
        setIsLoading(false)
        return
      }

      const response = await fetchImage(id)
      const file = base64ToFile(response.data, 'image.jpg', 'image/jpeg')

      const cache = await caches.open('immotep-cache-v1')
      await cache.put(`/images/${id}`, new Response(file))

      const url = URL.createObjectURL(file)
      setData(url)
      setIsLoading(false)
    }

    fetchData()

    // eslint-disable-next-line consistent-return
    return () => {
      if (data) URL.revokeObjectURL(data)
    }
  }, [id, fetchImage])

  return { data, isLoading, updateCache }
}

export default useImageCache
