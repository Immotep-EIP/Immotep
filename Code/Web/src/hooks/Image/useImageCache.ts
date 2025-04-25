import { useState, useEffect, useRef } from 'react'
import base64ToFile from '@/utils/base64/baseToFile'

const useImageCache = (
  id: string | undefined,
  fetchImage: (id: string) => Promise<any>
) => {
  const [data, setData] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState<boolean>(true)
  const objectUrlRef = useRef<string | null>(null)

  const updateCache = async (newImageData: string) => {
    if (!id) return

    const file = base64ToFile(
      newImageData.split(',')[1],
      'image.jpg',
      'image/jpeg'
    )

    const cache = await caches.open('keyz-cache-v1')
    await cache.put(`/images/${id}`, new Response(file))

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

    const cachedResponse = await caches.match(`/images/${id}`)
    if (cachedResponse) {
      const blob = await cachedResponse.blob()
      const url = URL.createObjectURL(blob)
      objectUrlRef.current = url
      setData(url)
      setIsLoading(false)
      return
    }

    try {
      const response = await fetchImage(id)
      if (response) {
        const file = base64ToFile(response.data, 'image.jpg', 'image/jpeg')

        const cache = await caches.open('keyz-cache-v1')
        await cache.put(`/images/${id}`, new Response(file))

        const fileUrl = URL.createObjectURL(file)
        objectUrlRef.current = fileUrl
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
      if (objectUrlRef.current) {
        URL.revokeObjectURL(objectUrlRef.current)
      }
    }
  }, [id, fetchImage])

  return { data, isLoading, updateCache }
}

export default useImageCache
