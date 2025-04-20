type ApiMethod = 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'

export interface ApiCallerParams<TResponse = any, TBody = any> {
  method: ApiMethod
  endpoint: string
  body?: TBody | string
  headers?: Record<string, string>
  responseType?: TResponse
  params?: Record<string, any>
}

export interface GetRoomsByPropertyParams {
  propertyId: string
  archive?: boolean
}

export interface GetFurnituresByRoomParams {
  propertyId: string
  roomId: string
  archive?: boolean
}
