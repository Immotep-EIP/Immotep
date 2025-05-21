import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'

import { Empty, Typography } from 'antd'
import { MessageOutlined } from '@ant-design/icons'

import useProperties from '@/hooks/Property/useProperties'
import PageMeta from '@/components/ui/PageMeta/PageMeta'
import PageTitle from '@/components/ui/PageText/Title'
import ContactList from '@/components/features/Messages/ContactList'
import ChatInterface from '@/components/features/Messages/ChatInterface'

import { Contact } from '@/interfaces/Messages/Contact'

import style from './Messages.module.css'

const Messages: React.FC = () => {
  const { t } = useTranslation()
  const { loading } = useProperties()
  const [selectedContact, setSelectedContact] = useState<Contact | null>(null)

  const mockTenants = [
    {
      id: '1',
      name: 'John Doe',
      propertyName: 'Sunset Apartments 3B',
      lastMessage: 'The heating system is not working properly.',
      lastMessageDate: new Date('2024-01-15T10:30:00'),
      avatar: 'https://randomuser.me/api/portraits/men/1.jpg'
    },
    {
      id: '2',
      name: 'Sarah Wilson',
      propertyName: 'Ocean View Condo 12A',
      lastMessage: 'When will the maintenance team arrive?',
      lastMessageDate: new Date('2024-01-14T16:45:00'),
      avatar: 'https://randomuser.me/api/portraits/women/2.jpg'
    },
    {
      id: '3',
      name: 'Michael Chen',
      propertyName: 'Pine Street House',
      lastMessage: 'Thank you for fixing the issue so quickly!',
      lastMessageDate: new Date('2024-01-13T09:15:00'),
      avatar: 'https://randomuser.me/api/portraits/men/3.jpg'
    },
    {
      id: '4',
      name: 'Emma Thompson',
      propertyName: 'Downtown Loft 7C',
      lastMessage: 'Is it possible to renew my lease early?',
      lastMessageDate: new Date('2024-01-12T14:20:00'),
      avatar: 'https://randomuser.me/api/portraits/women/4.jpg'
    }
  ]

  const contacts: Contact[] = mockTenants.map(tenant => ({
    id: tenant.id,
    name: tenant.name,
    propertyName: tenant.propertyName,
    lastMessage: tenant.lastMessage,
    lastMessageDate: tenant.lastMessageDate,
    avatar: tenant.avatar
  }))

  if (loading) {
    return <div>{t('components.loading')}</div>
  }

  if (contacts.length === 0) {
    return (
      <div className={style.emptyContainer}>
        <Empty
          image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
          description={
            <Typography.Text>
              {t('components.messages.no_properties_so_no_messages')}
            </Typography.Text>
          }
        />
      </div>
    )
  }

  return (
    <>
      <PageMeta
        title={t('pages.messages.document_title')}
        description={t('pages.messages.document_description')}
        keywords="messages, communication,
         Keyz"
      />
      <div className={style.pageContainer}>
        <div className={style.pageHeader}>
          <PageTitle title={t('pages.messages.title')} size="title" />
        </div>
        <div className={style.contentContainer}>
          <ContactList
            contacts={contacts}
            onSelectContact={setSelectedContact}
            selectedContact={selectedContact}
          />
          <div className={style.chatContainer}>
            {selectedContact ? (
              <ChatInterface contact={selectedContact} />
            ) : (
              <div className={style.noChatSelected}>
                <MessageOutlined style={{ fontSize: '48px' }} />
                <Typography.Text>
                  {t('pages.messages.select_contact')}
                </Typography.Text>
              </div>
            )}
          </div>
        </div>
      </div>
    </>
  )
}

export default Messages
