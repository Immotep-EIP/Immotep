import React from 'react'
import { Button as AntButton } from 'antd'
import type { ButtonProps as AntButtonProps } from 'antd'

interface ButtonProps extends AntButtonProps {
  children?: React.ReactNode
  type?: 'primary' | 'default' | 'dashed' | 'text' | 'link'
  size?: 'small' | 'middle' | 'large'
  loading?: boolean
  disabled?: boolean
  onClick?: () => void
}

const Button: React.FC<ButtonProps> = ({
  children,
  type = 'primary',
  size = 'middle',
  loading = false,
  disabled = false,
  onClick,
  ...props
}) => (
  <AntButton
    type={type}
    size={size}
    loading={loading}
    disabled={disabled}
    onClick={onClick}
    {...props}
  >
    {children}
  </AntButton>
)

export default Button
