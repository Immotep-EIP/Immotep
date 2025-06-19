import React from 'react'
import { Modal as AntModal } from 'antd'
import type { ModalProps as AntModalProps } from 'antd'

interface ModalProps extends Omit<AntModalProps, 'onOk' | 'onCancel'> {
  isOpen: boolean
  title: string
  children: React.ReactNode
  onOk?: () => void
  onCancel?: () => void
  size?: 'small' | 'medium' | 'large'
}

const Modal: React.FC<ModalProps> = ({
  isOpen,
  title,
  children,
  onOk,
  onCancel,
  size = 'medium',
  ...props
}) => {
  const getWidth = () => {
    switch (size) {
      case 'small':
        return 400
      case 'large':
        return 800
      default:
        return 600
    }
  }

  return (
    <AntModal
      title={title}
      open={isOpen}
      onOk={onOk}
      onCancel={onCancel}
      width={getWidth()}
      {...props}
    >
      {children}
    </AntModal>
  )
}

export default Modal
