# Card Component Documentation

## Overview
The Card component is a wrapper around Ant Design's Card component that provides enhanced functionality with custom variants, padding options, and consistent styling throughout the Keyz application.

## Component Interface

```typescript
import React from 'react'
import { Card as AntCard } from 'antd'
import type { CardProps as AntCardProps } from 'antd'

interface CardProps extends Omit<AntCardProps, 'title' | 'variant'> {
  title?: React.ReactNode
  children: React.ReactNode
  customVariant?: 'default' | 'elevated' | 'outlined'
  padding?: 'none' | 'small' | 'medium' | 'large'
}

const Card: React.FC<CardProps> & { Grid: typeof AntCard.Grid }
```

## Basic Usage

### Simple Card
```tsx
import { Card } from '@/components/common'

// Basic card with content
<Card>
  <p>This is card content</p>
</Card>

// Card with title
<Card title="Card Title">
  <p>Card content goes here</p>
</Card>
```

### Card Variants
```tsx
// Default variant
<Card customVariant="default">
  <p>Default card</p>
</Card>

// Elevated variant with shadow
<Card customVariant="elevated">
  <p>Elevated card with shadow</p>
</Card>

// Outlined variant with border
<Card customVariant="outlined">
  <p>Outlined card with border</p>
</Card>
```

### Padding Options
```tsx
// No padding
<Card padding="none">
  <p>No padding content</p>
</Card>

// Small padding
<Card padding="small">
  <p>Small padding content</p>
</Card>

// Medium padding (default)
<Card padding="medium">
  <p>Medium padding content</p>
</Card>

// Large padding
<Card padding="large">
  <p>Large padding content</p>
</Card>
```

### Card Grid
```tsx
// Using Card.Grid for layout
<Card title="Grid Layout">
  <Card.Grid style={{ width: '25%' }}>
    Content 1
  </Card.Grid>
  <Card.Grid style={{ width: '25%' }}>
    Content 2
  </Card.Grid>
  <Card.Grid style={{ width: '50%' }}>
    Content 3
  </Card.Grid>
</Card>
```

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `title` | `ReactNode` | - | Card title |
| `children` | `ReactNode` | - | Card content (required) |
| `customVariant` | `'default' \| 'elevated' \| 'outlined'` | `'default'` | Visual variant of the card |
| `padding` | `'none' \| 'small' \| 'medium' \| 'large'` | `'medium'` | Internal padding size |
| `bordered` | `boolean` | `true` | Show card border |
| `hoverable` | `boolean` | `false` | Lift up on hover |
| `loading` | `boolean` | `false` | Show loading skeleton |
| `size` | `'default' \| 'small'` | `'default'` | Card size |
| `actions` | `ReactNode[]` | - | Action buttons at bottom |
| `extra` | `ReactNode` | - | Extra content in title area |
| `cover` | `ReactNode` | - | Cover image/content |

### Padding Values
- `none`: 0px
- `small`: 12px  
- `medium`: 16px (default)
- `large`: 24px

## Real-world Examples

### Property Card
```tsx
const PropertyCard = ({ property }: { property: Property }) => {
  return (
    <Card
      customVariant="elevated"
      hoverable
      cover={
        <img 
          src={property.image} 
          alt={property.title}
          style={{ height: 200, objectFit: 'cover' }}
        />
      }
      actions={[
        <Button key="view" type="text">View</Button>,
        <Button key="edit" type="text">Edit</Button>,
        <Button key="delete" type="text" danger>Delete</Button>
      ]}
    >
      <Card.Meta 
        title={property.title}
        description={`${property.price} • ${property.location}`}
      />
    </Card>
  )
}
```

### Content Grid
```tsx
const ServiceGrid = ({ services }: { services: Service[] }) => {
  return (
    <Card title="Our Services" customVariant="outlined">
      {services.map(service => (
        <Card.Grid 
          key={service.id}
          style={{ width: '33.33%', textAlign: 'center' }}
        >
          <div className="service-item">
            <Icon component={service.icon} style={{ fontSize: 24 }} />
            <h4>{service.title}</h4>
            <p>{service.description}</p>
          </div>
        </Card.Grid>
      ))}
    </Card>
  )
}
```

## Accessibility

### Best Practices
- Provide meaningful titles for screen readers
- Use proper heading hierarchy in card content
- Ensure interactive elements have proper focus states
- Add appropriate ARIA labels for complex cards
