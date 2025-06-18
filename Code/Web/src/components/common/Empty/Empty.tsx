import React from 'react'
import { Empty as AntEmpty, Typography } from 'antd'

interface EmptyStateProps {
  description: React.ReactNode
  className?: string
}

const Empty: React.FC<EmptyStateProps> = ({ description, className }) => (
  <AntEmpty
    image="https://gw.alipayobjects.com/zos/antfincdn/ZHrcdLPrvN/empty.svg"
    description={<Typography.Text>{description}</Typography.Text>}
    className={className}
  />
)

export default Empty
