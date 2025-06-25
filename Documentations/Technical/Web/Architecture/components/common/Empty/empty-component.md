# Empty Component Documentation

## Overview
The Empty component is a wrapper around Ant Design's Empty component that provides a consistent empty state UI with custom image and description throughout the Keyz application.

## Component Interface

```typescript
import React from 'react'
import { Empty as AntEmpty, Typography } from 'antd'

interface EmptyStateProps {
  description: React.ReactNode
  className?: string
}

const Empty: React.FC<EmptyStateProps>
```

## Basic Usage

### Simple Empty State
```tsx
import { Empty } from '@/components/common'

// Basic empty state
<Empty description="No data available" />

// With custom className
<Empty 
  description="No properties found" 
  className="custom-empty"
/>
```

### With Different Descriptions
```tsx
// Text description
<Empty description="No items to display" />

// Rich description with JSX
<Empty 
  description={
    <div>
      <p>No results found</p>
      <p>Try adjusting your search criteria</p>
    </div>
  } 
/>

// Description with action
<Empty 
  description={
    <div>
      <p>No properties in your list</p>
      <Button type="primary">Add Property</Button>
    </div>
  } 
/>
```

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `description` | `ReactNode` | - | Description text or element (required) |
| `className` | `string` | - | Additional CSS class names |

### Inherited Features
The component includes:
- Custom empty state image from `@/assets/images/EmptyImage.svg`
- Typography wrapper for consistent text styling
- All Ant Design Empty features through extension

## Real-world Examples

### Property List Empty State
```tsx
const PropertyListEmpty = ({ onAddProperty }: { onAddProperty: () => void }) => {
  return (
    <Empty 
      description={
        <div className="empty-property-list">
          <Typography.Title level={4}>No Properties Yet</Typography.Title>
          <Typography.Text type="secondary">
            Start by adding your first property to get started with Keyz
          </Typography.Text>
          <div style={{ marginTop: 16 }}>
            <Button type="primary" onClick={onAddProperty}>
              Add Your First Property
            </Button>
          </div>
        </div>
      }
    />
  )
}
```

### Messages Empty State
```tsx
const MessagesEmpty = () => {
  return (
    <Empty 
      description={
        <div className="messages-empty">
          <Typography.Title level={5}>No Messages</Typography.Title>
          <Typography.Text type="secondary">
            When you receive messages, they'll appear here
          </Typography.Text>
        </div>
      }
    />
  )
}
```

## Accessibility

### Best Practices
- Provide meaningful descriptions that explain the empty state
- Use proper semantic HTML in description content
- Ensure sufficient color contrast for text
- Include actionable guidance when appropriate

## Usage Patterns

### Conditional Rendering
```tsx
const DataList = ({ data, loading }: DataListProps) => {
  if (loading) {
    return <Spin size="large" />
  }

  if (!data || data.length === 0) {
    return <Empty description="No data available" />
  }

  return (
    <List
      dataSource={data}
      renderItem={item => <List.Item>{item.name}</List.Item>}
    />
  )
}
```