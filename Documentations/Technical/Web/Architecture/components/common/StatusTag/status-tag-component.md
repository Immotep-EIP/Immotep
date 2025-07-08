# StatusTag Component Documentation

---

## Overview
The StatusTag component is a specialized wrapper around Ant Design's Tag component that provides internationalized status display with customizable color mapping throughout the Keyz application.

---

## Component Interface

```typescript
import React from 'react'
import { Tag } from 'antd'
import { useTranslation } from 'react-i18next'

interface StatusTagProps {
  value: string
  colorMap: Record<string, string>
  i18nPrefix?: string
  defaultColor?: string
}

const StatusTag: React.FC<StatusTagProps>
```

---

## Basic Usage

### Simple Status Tag
```tsx
import { StatusTag } from '@/components/common'

// Basic status tag
const colorMap = {
  active: 'green',
  inactive: 'red',
  pending: 'orange'
}

<StatusTag 
  value="active" 
  colorMap={colorMap} 
/>
```

### With Custom Default Color
```tsx
// Status tag with custom default color
<StatusTag 
  value="unknown" 
  colorMap={colorMap}
  defaultColor="blue" 
/>
```

### With Internationalization
```tsx
// Status tag with translation support
<StatusTag 
  value="active" 
  colorMap={colorMap}
  i18nPrefix="status.property" 
/>
// Will attempt to translate using "status.property.active" key
```

---

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `value` | `string` | - | Status value to display (required) |
| `colorMap` | `Record<string, string>` | - | Mapping of status values to colors (required) |
| `i18nPrefix` | `string` | - | Translation key prefix for internationalization |
| `defaultColor` | `string` | `'default'` | Default color when value not found in colorMap |

### Color Mapping
- Keys in `colorMap` should be lowercase
- Values can be Ant Design color names or custom hex colors
- Component automatically converts `value` to lowercase for lookup

---

## Real-world Examples

### Property Status
```tsx
const PropertyStatus = ({ status }: { status: string }) => {
  const propertyColorMap = {
    available: 'green',
    occupied: 'blue',
    maintenance: 'orange',
    unavailable: 'red'
  }

  return (
    <StatusTag
      value={status}
      colorMap={propertyColorMap}
      i18nPrefix="property.status"
    />
  )
}

// Usage
<PropertyStatus status="available" />
// Displays: "Available" (translated) with green color
```

### Lease Status
```tsx
const LeaseStatus = ({ lease }: { lease: Lease }) => {
  const leaseColorMap = {
    active: 'green',
    pending: 'orange',
    expired: 'red',
    terminated: 'volcano',
    draft: 'blue'
  }

  return (
    <StatusTag
      value={lease.status}
      colorMap={leaseColorMap}
      i18nPrefix="lease.status"
      defaultColor="gray"
    />
  )
}
```

### Document Status
```tsx
const DocumentStatus = ({ document }: { document: Document }) => {
  const documentColorMap = {
    draft: 'gray',
    pending: 'orange',
    approved: 'green',
    rejected: 'red',
    expired: 'volcano'
  }

  return (
    <StatusTag
      value={document.status}
      colorMap={documentColorMap}
      i18nPrefix="document.status"
    />
  )
}
```

---

## Common Patterns

### Dynamic Color Mapping
```tsx
const getDynamicColorMap = (type: string) => {
  const colorMaps = {
    property: { available: 'green', occupied: 'blue' },
    lease: { active: 'green', pending: 'orange' },
    payment: { paid: 'green', overdue: 'red' }
  }
  
  return colorMaps[type] || {}
}

<StatusTag
  value={status}
  colorMap={getDynamicColorMap(entityType)}
  i18nPrefix={`${entityType}.status`}
/>
```

### Status with Icons
```tsx
const StatusWithIcon = ({ status, icon }: StatusWithIconProps) => {
  return (
    <div className="status-with-icon">
      {icon}
      <StatusTag
        value={status}
        colorMap={statusColorMap}
        i18nPrefix="status"
      />
    </div>
  )
}
``` 