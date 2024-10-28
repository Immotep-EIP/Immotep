type ApiMethod = 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'

export interface ApiCallerParams<Type> {
  method: ApiMethod
  endpoint: string
  data?: Type | string
  headers?: Record<string, string>
}
