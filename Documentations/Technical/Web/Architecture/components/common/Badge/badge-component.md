# Badge Component Documentation

## Overview
The Badge component is a wrapper around Ant Design's Badge component that provides a consistent interface for displaying badges, counts, status indicators, and ribbons throughout the Keyz application.

## Component Interface

```typescript
import React from 'react'
import { Badge as AntBadge } from 'antd'
import type { BadgeProps as AntBadgeProps } from 'antd'

export interface BadgeProps extends AntBadgeProps {}

const Badge: React.FC<BadgeProps> & { Ribbon: typeof AntBadge.Ribbon }
```

## Basic Usage

### Count Badge
```tsx
import { Badge } from '@/components/common'

// Simple count badge
<Badge count={5}>
  <Avatar />
</Badge>

// With overflow count
<Badge count={100} overflowCount={99}>
  <Avatar />
</Badge>

// Show zero count
<Badge count={0} showZero>
  <Avatar />
</Badge>
```

### Dot Badge
```tsx
// Simple dot indicator
<Badge dot>
  <NotificationIcon />
</Badge>

// Dot with color
<Badge dot color="red">
  <NotificationIcon />
</Badge>
```

### Status Badge
```tsx
// Status indicators
<Badge status="success" text="Success" />
<Badge status="error" text="Error" />
<Badge status="warning" text="Warning" />
<Badge status="processing" text="Processing" />
<Badge status="default" text="Default" />
```

### Custom Colors
```tsx
// Predefined colors
<Badge count={5} color="blue">
  <Avatar />
</Badge>

// Custom hex color
<Badge count={5} color="#87d068">
  <Avatar />
</Badge>
```

## Ribbon Usage

The Badge component exposes `Badge.Ribbon` for creating ribbon badges:

```tsx
// Basic ribbon
<Badge.Ribbon text="New">
  <Card>
    <h3>Product Card</h3>
    <p>Product description</p>
  </Card>
</Badge.Ribbon>

// Ribbon with custom color
<Badge.Ribbon text="Hot" color="red">
  <Card>
    <h3>Featured Product</h3>
  </Card>
</Badge.Ribbon>

// Ribbon with different placement
<Badge.Ribbon text="Popular" placement="start">
  <Card>
    <h3>Popular Item</h3>
  </Card>
</Badge.Ribbon>
```

## Props

The Badge component accepts all Ant Design Badge props:

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `count` | `number \| ReactNode` | - | Number to show in badge |
| `dot` | `boolean` | `false` | Show a red dot without count |
| `color` | `string` | - | Custom color for the badge |
| `size` | `'default' \| 'small'` | `'default'` | Size of the badge |
| `status` | `'success' \| 'processing' \| 'default' \| 'error' \| 'warning'` | - | Badge status |
| `text` | `ReactNode` | - | Text to display next to status badge |
| `showZero` | `boolean` | `false` | Show badge when count is zero |
| `overflowCount` | `number` | `99` | Max count to show |
| `offset` | `[number, number]` | - | Offset of the badge |
| `title` | `string` | - | Title attribute for accessibility |
| `children` | `ReactNode` | - | Content to wrap with badge |

## Badge.Ribbon Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `text` | `ReactNode` | - | Text content of the ribbon |
| `color` | `string` | - | Color of the ribbon |
| `placement` | `'start' \| 'end'` | `'end'` | Position of the ribbon |
| `children` | `ReactNode` | - | Content to wrap with ribbon |

## Real-world Examples

### Product Card with Ribbon
```tsx
const ProductCard = ({ product }: { product: Product }) => {
  const ribbonProps = {
    text: product.isNew ? 'New' : product.isHot ? 'Hot' : null,
    color: product.isNew ? 'blue' : 'red'
  }

  const content = (
    <Card>
      <img src={product.image} alt={product.name} />
      <h3>{product.name}</h3>
      <p>{product.price}</p>
    </Card>
  )

  return ribbonProps.text ? (
    <Badge.Ribbon {...ribbonProps}>
      {content}
    </Badge.Ribbon>
  ) : content
}
```

### Status Indicator
```tsx
const ServiceStatus = ({ status }: { status: ServiceStatusType }) => {
  const statusConfig = {
    online: { status: 'success' as const, text: 'Online' },
    offline: { status: 'error' as const, text: 'Offline' },
    maintenance: { status: 'warning' as const, text: 'Maintenance' },
    loading: { status: 'processing' as const, text: 'Starting...' }
  }

  return <Badge {...statusConfig[status]} />
}
```

## Accessibility

### Best Practices
- Always provide meaningful `title` attributes for screen readers
- Use appropriate ARIA labels when needed
- Ensure sufficient color contrast for custom colors
- Test with keyboard navigation

### Example with Accessibility
```tsx
<Badge 
  count={unreadCount} 
  title={`${unreadCount} unread notifications`}
  aria-label={`${unreadCount} unread notifications`}
>
  <Button icon={<BellOutlined />} />
</Badge>
```