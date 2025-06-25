import React from 'react'
import { Input as AntInput } from 'antd'
import type { InputProps as AntInputProps } from 'antd'

interface InputProps extends Omit<AntInputProps, 'onChange'> {
  label?: string
  error?: string
  value?: string
  onChange?: (value: string) => void
  required?: boolean
}

const Input: React.FC<InputProps> = ({
  label,
  error,
  value,
  onChange,
  required = false,
  ...props
}) => {
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    onChange?.(e.target.value)
  }

  return (
    <div>
      {label && (
        <label htmlFor={props.id || props.name}>
          {label}
          {required && <span style={{ color: 'red' }}> *</span>}
        </label>
      )}
      <AntInput
        id={props.id || props.name}
        value={value}
        onChange={handleChange}
        status={error ? 'error' : undefined}
        {...props}
      />
      {error && <div style={{ color: 'red', fontSize: '12px' }}>{error}</div>}
    </div>
  )
}

export default Input
