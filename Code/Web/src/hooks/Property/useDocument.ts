import { useEffect, useState } from 'react'
import { Document, UseDocumentReturn } from '@/interfaces/Property/Document'
import GetPropertyDocuments from '@/services/api/Owner/Properties/GetPropertyDocuments'
import UploadDocument from '@/services/api/Owner/Properties/UploadDocument'
import fileToBase64 from '@/utils/base64/fileToBase'
import useProperties from './useProperties'

const useDocument = (propertyId: string): UseDocumentReturn => {
  const [documents, setDocuments] = useState<Document[] | null>(null)
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)
  const { propertyDetails: propertyData } = useProperties(propertyId)

  const fetchDocuments = async (propertyId: string) => {
    try {
      if (
        !propertyId ||
        propertyId === '' ||
        propertyData?.status !== 'available'
      ) {
        setError('No tenant assigned to this property')
        return
      }
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

  const uploadDocument = async (
    file: File,
    documentName: string,
    propertyId: string,
    leaseId: string = 'current'
  ) => {
    try {
      setLoading(true)
      setError(null)

      const base64Data = await fileToBase64(file)
      const payload = {
        name: documentName,
        data: base64Data.split(',')[1]
      }

      await UploadDocument(JSON.stringify(payload), propertyId, leaseId)
      await fetchDocuments(propertyId)
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : 'An error occurred while uploading the document'
      )
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
    refreshDocuments: fetchDocuments,
    uploadDocument
  }
}

export default useDocument
