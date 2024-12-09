type ApiMethod = 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'

export interface ApiCallerParams<TData> {
  method: ApiMethod
  endpoint: string
  data?: TData | string
  headers?: Record<string, string>
}
