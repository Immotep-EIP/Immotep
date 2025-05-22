import React from 'react'
import { useNavigate } from 'react-router-dom'

import { Badge } from 'antd'

import NavigationEnum from '@/enums/NavigationEnum'
import { WidgetProps } from '@/interfaces/Widgets/Widgets.ts'

import ArrowRight from '@/assets/icons/arrowRight.png'
import style from './LastMessages.module.css'

const LastMessages: React.FC<WidgetProps> = ({ height }) => {
  const rowHeight = 120
  const pixelHeight = height * rowHeight
  const navigate = useNavigate()

  const messages = [
    {
      expeditor: 'Expeditor tenant 1',
      lastMessage:
        'This is the last message sent by the expeditor, with a very long text to test the overflow',
      date: '2023-10-01',
      read: false
    },
    {
      expeditor: 'Expeditor tenant 2',
      lastMessage: 'This is the last message sent by the expeditor',
      date: '2023-10-01',
      read: true
    },
    {
      expeditor: 'Expeditor very long tenant 3',
      lastMessage: 'This is the last message sent by the expeditor',
      date: '2023-10-01',
      read: false
    }
  ]

  return (
    <div
      className={style.layoutContainer}
      style={{ height: `${pixelHeight}px` }}
    >
      <div className={style.contentContainer}>
        {messages.map(message => (
          <div
            key={`${message.expeditor}-${message}`}
            className={style.messageContainer}
            onClick={() => {
              navigate(NavigationEnum.MESSAGES)
            }}
            onKeyDown={(e: React.KeyboardEvent) => {
              if (e.key === 'Enter') {
                navigate(NavigationEnum.MESSAGES)
              }
            }}
            role="button"
            tabIndex={0}
          >
            <div className={style.messageInfosContainer}>
              <span className={style.expeditorText}>{message.expeditor}</span>
              <span className={style.messageDate}>{message.date}</span>
            </div>
            <div className={style.messageContentContainer}>
              {!message.read ? (
                <span className={style.messageText}>{message.lastMessage}</span>
              ) : (
                <Badge
                  className={style.messageText}
                  key={`blue${message.lastMessage}`}
                  color="blue"
                  text={message.lastMessage}
                  style={{
                    fontWeight: 700
                  }}
                />
              )}
              <div className={style.informationsIconContainer}>
                <img
                  src={ArrowRight}
                  alt="messages"
                  style={{
                    width: '15px',
                    opacity: 0.5
                  }}
                />
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default LastMessages
