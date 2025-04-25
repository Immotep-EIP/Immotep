import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { Input } from 'antd'
import { SearchOutlined } from '@ant-design/icons'

import { ContactListProps } from '@/interfaces/Messages/Contact'
import defaultAvatar from '@/assets/images/DefaultProfile.png'
import style from './ContactList.module.css'

const ContactList: React.FC<ContactListProps> = ({
  contacts,
  onSelectContact,
  selectedContact
}) => {
  const [searchTerm, setSearchTerm] = useState('')
  const { t } = useTranslation()

  const filteredContacts = contacts.filter(
    contact =>
      contact.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      contact.propertyName.toLowerCase().includes(searchTerm.toLowerCase())
  )

  return (
    <div className={style.sidebarContainer}>
      <div className={style.searchContainer}>
        <Input
          prefix={<SearchOutlined />}
          placeholder={t('pages.messages.search_contacts')}
          onChange={e => setSearchTerm(e.target.value)}
          className={style.searchInput}
        />
      </div>
      <div className={style.contactsList}>
        {filteredContacts.map(contact => (
          <div
            key={contact.id}
            className={`${style.contactItem} ${
              selectedContact?.id === contact.id ? style.active : ''
            }`}
            onClick={() => onSelectContact(contact)}
            role="button"
            tabIndex={0}
            onKeyDown={e => {
              if (e.key === 'Enter' || e.key === ' ') {
                onSelectContact(contact)
              }
            }}
          >
            <img
              src={contact.avatar || defaultAvatar}
              alt={contact.name}
              className={style.contactAvatar}
            />
            <div className={style.contactInfo}>
              <div className={style.contactName}>{contact.name}</div>
              <div className={style.propertyName}>{contact.propertyName}</div>
              {contact.lastMessage && (
                <div className={style.lastMessage}>{contact.lastMessage}</div>
              )}
            </div>
            {contact.lastMessageDate && (
              <div className={style.messageDate}>
                {new Date(contact.lastMessageDate).toLocaleDateString()}
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  )
}

export default ContactList
