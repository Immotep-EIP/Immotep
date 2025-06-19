import React from 'react'
import { Card as AntCard } from 'antd'
import type { CardProps as AntCardProps } from 'antd'

interface CardProps extends Omit<AntCardProps, 'title' | 'variant'> {
  title?: React.ReactNode
  children: React.ReactNode
  customVariant?: 'default' | 'elevated' | 'outlined'
  padding?: 'none' | 'small' | 'medium' | 'large'
}

const Card: React.FC<CardProps> & { Grid: typeof AntCard.Grid } = ({
  title,
  children,
  customVariant = 'default',
  padding = 'medium',
  ...props
}) => {
  const getPadding = () => {
    switch (padding) {
      case 'none':
        return 0
      case 'small':
        return 12
      case 'large':
        return 24
      default:
        return 16
    }
  }

  const getStyle = () => {
    const baseStyle = {
      padding: getPadding()
    }

    switch (customVariant) {
      case 'elevated':
        return {
          ...baseStyle,
          boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)'
        }
      case 'outlined':
        return {
          ...baseStyle,
          border: '1px solid #d9d9d9'
        }
      default:
        return baseStyle
    }
  }

  return (
    <AntCard title={title} styles={{ body: getStyle() }} {...props}>
      {children}
    </AntCard>
  )
}

Card.Grid = AntCard.Grid

export default Card
