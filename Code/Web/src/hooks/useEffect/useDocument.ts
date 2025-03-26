import { useEffect, useState } from 'react'
import { Document } from '@/interfaces/Property/Document'
import GetPropertyDocuments from '@/services/api/Owner/Properties/GetPropertyDocuments'

interface UseDocumentReturn {
  documents: Document[] | null
  loading: boolean
  error: string | null
  refreshDocuments: (propertyId: string) => Promise<void>
}

const useDocument = (propertyId: string): UseDocumentReturn => {
  const [documents, setDocuments] = useState<Document[] | null>(null)
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)

  const fetchDocuments = async (propertyId: string) => {
    try {
      setLoading(true)
      setError(null)
      const response = await GetPropertyDocuments(propertyId)
      setDocuments(response)
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : 'An error occurred while fetching the documents'
      )
      setDocuments(null)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    if (propertyId) {
      fetchDocuments(propertyId)
    }
  }, [propertyId])

  return {
    documents,
    loading,
    error,
    refreshDocuments: fetchDocuments
  }
}

export default useDocument
