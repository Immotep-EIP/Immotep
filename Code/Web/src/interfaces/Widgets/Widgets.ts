import React from 'react'

export interface Widget {
  i: string
  children: React.ReactNode
  name: string
  logo?: React.ReactElement
  x: number
  y: number
  w: number
  h: number
}

export interface Layout {
  i: string
  x: number
  y: number
  w: number
  h: number
}

export interface MaintenanceTask {
  id: number
  description: string
  priority: 'high' | 'medium' | 'low'
  completed: boolean
}

export type addWidgetType = {
  name: string
  width: number
  height: number
  types: string
}

export interface AddWidgetModalProps {
  isOpen: boolean
  onClose: () => void
  onAddWidget: (widget: addWidgetType) => void
}

export interface WidgetProps {
  height: number
}
