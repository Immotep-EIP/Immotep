import React from 'react'
import { Button as AntButton } from 'antd'
import type { ButtonProps as AntButtonProps } from 'antd'

interface ButtonProps extends AntButtonProps {
  children?: React.ReactNode
}

const Button: React.FC<ButtonProps> = ({
  children,
  type = 'primary',
  size = 'middle',
  loading = false,
  disabled = false,
  ...props
}) => (
  <AntButton
    type={type}
    size={size}
    loading={loading}
    disabled={disabled}
    {...props}
  >
    {children}
  </AntButton>
)

export default Button
