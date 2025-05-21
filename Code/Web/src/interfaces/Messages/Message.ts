export interface Message {
  id: string
  sender: 'me' | 'contact'
  content: string
  timestamp: Date
}
