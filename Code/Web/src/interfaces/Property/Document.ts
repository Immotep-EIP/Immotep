export interface Document {
  created_at: string
  data: string
  id: string
  name: string
}

export interface UseDocumentReturn {
  documents: Document[] | null
  loading: boolean
  error: string | null
  refreshDocuments: (propertyId: string) => Promise<void>
  uploadDocument: (
    file: File,
    documentName: string,
    propertyId: string,
    leaseId: string
  ) => Promise<void>
  deleteDocument: (documentId: string) => Promise<void>
}
