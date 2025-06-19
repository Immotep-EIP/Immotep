import React from 'react'
import { Badge as AntBadge } from 'antd'
import type { BadgeProps as AntBadgeProps } from 'antd'

export interface BadgeProps extends AntBadgeProps {}

const Badge: React.FC<BadgeProps> & { Ribbon: typeof AntBadge.Ribbon } = ({
  children,
  ...props
}) => <AntBadge {...props}>{children}</AntBadge>

Badge.Ribbon = AntBadge.Ribbon

export default Badge
