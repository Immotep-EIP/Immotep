import React from 'react'
import { Empty as AntEmpty, Typography } from 'antd'

import empty from '@/assets/images/EmptyImage.svg'

interface EmptyStateProps {
  description: React.ReactNode
  className?: string
}

const Empty: React.FC<EmptyStateProps> = ({ description, className }) => (
  <AntEmpty
    image={empty}
    description={<Typography.Text>{description}</Typography.Text>}
    className={className}
  />
)

export default Empty
