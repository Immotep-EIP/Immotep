export interface Contact {
  id: string
  name: string
  propertyName: string
  lastMessage?: string
  lastMessageDate?: Date
  avatar?: string
}

export interface ContactListProps {
  contacts: Contact[]
  onSelectContact: (contact: Contact) => void
  selectedContact: Contact | null
}
