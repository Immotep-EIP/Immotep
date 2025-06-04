import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'

import { Input, Button } from 'antd'
import { SendOutlined, PlusCircleOutlined } from '@ant-design/icons'

import { Contact } from '@/interfaces/Messages/Contact'
import { Message } from '@/interfaces/Messages/Message'

import defaultAvatar from '@/assets/images/DefaultProfile.png'
import style from './ChatInterface.module.css'

const ChatHeader: React.FC<{ contact: Contact }> = ({ contact }) => (
  <div className={style.chatHeader}>
    <img
      src={contact.avatar || defaultAvatar}
      alt={contact.name}
      className={style.chatHeaderAvatar}
    />
    <div className={style.chatHeaderInfo}>
      <div className={style.chatHeaderName}>{contact.name}</div>
      <div className={style.chatHeaderProperty}>{contact.propertyName}</div>
    </div>
  </div>
)

const ChatInterface: React.FC<{ contact: Contact }> = ({ contact }) => {
  const [message, setMessage] = useState('')
  const { t } = useTranslation()

  // Mock messages
  const messages: Message[] = [
    {
      id: '1',
      sender: 'contact',
      content: 'Hello! How can I help you today?',
      timestamp: new Date('2024-01-15T10:00:00')
    },
    {
      id: '2',
      sender: 'me',
      content: 'I have a question about my lease.',
      timestamp: new Date('2024-01-15T10:05:00')
    },
    {
      id: '3',
      sender: 'contact',
      content: 'Sure, what would you like to know?',
      timestamp: new Date('2024-01-15T10:10:00')
    }
  ]

  const handleSend = () => {
    if (message.trim()) {
      setMessage('')
    }
  }

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault()
      handleSend()
    }
  }

  return (
    <>
      <ChatHeader contact={contact} />
      <div className={style.messagesContainer}>
        {messages.map(msg => (
          <div
            key={msg.id}
            className={`${style.messageBubble} ${
              msg.sender === 'me' ? style.myMessage : style.contactMessage
            }`}
          >
            <div className={style.messageContent}>{msg.content}</div>
            <div className={style.messageTimestamp}>
              {msg.timestamp.toLocaleTimeString([], {
                hour: '2-digit',
                minute: '2-digit'
              })}
            </div>
          </div>
        ))}
      </div>
      <div className={style.inputContainer}>
        <Button icon={<PlusCircleOutlined />} onClick={() => {}} />
        <Input.TextArea
          className={style.messageInput}
          value={message}
          onChange={e => setMessage(e.target.value)}
          onKeyDown={handleKeyDown}
          placeholder={t('pages.messages.type_message')}
          autoSize={{ minRows: 1, maxRows: 4 }}
        />
        <Button
          type="primary"
          icon={<SendOutlined />}
          onClick={handleSend}
          disabled={!message.trim()}
        />
      </div>
    </>
  )
}

export default ChatInterface
