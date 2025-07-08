# Button Component Documentation

---

## Overview
The Button component is a wrapper around Ant Design's Button component that provides consistent button styling and behavior throughout the Keyz application with customizable defaults.

---

## Component Interface

```typescript
import React from 'react'
import { Button as AntButton } from 'antd'
import type { ButtonProps as AntButtonProps } from 'antd'

interface ButtonProps extends AntButtonProps {
  children?: React.ReactNode
}

const Button: React.FC<ButtonProps>
```

---

## Basic Usage

### Primary Button
```tsx
import { Button } from '@/components/common'

// Basic primary button (default)
<Button>
  Click me
</Button>

// Explicit primary type
<Button type="primary">
  Primary Action
</Button>
```

### Button Types
```tsx
// Different button types
<Button type="primary">Primary</Button>
<Button type="default">Default</Button>
<Button type="dashed">Dashed</Button>
<Button type="text">Text</Button>
<Button type="link">Link</Button>
```

### Button Sizes
```tsx
// Different sizes
<Button size="large">Large Button</Button>
<Button size="middle">Middle Button</Button> {/* Default */}
<Button size="small">Small Button</Button>
```

### Button States
```tsx
// Loading state
<Button loading>
  Loading...
</Button>

// Disabled state
<Button disabled>
  Disabled
</Button>

// With icon
<Button icon={<SearchOutlined />}>
  Search
</Button>
```

---

## Props

The Button component accepts all Ant Design Button props with custom defaults:

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `type` | `'primary' \| 'default' \| 'dashed' \| 'text' \| 'link'` | `'primary'` | Button type |
| `size` | `'large' \| 'middle' \| 'small'` | `'middle'` | Button size |
| `loading` | `boolean` | `false` | Show loading spinner |
| `disabled` | `boolean` | `false` | Disable button |
| `children` | `ReactNode` | - | Button content |
| `icon` | `ReactNode` | - | Icon element |
| `shape` | `'default' \| 'circle' \| 'round'` | `'default'` | Button shape |
| `ghost` | `boolean` | `false` | Make button background transparent |
| `danger` | `boolean` | `false` | Set danger status |
| `htmlType` | `'button' \| 'submit' \| 'reset'` | `'button'` | HTML button type |
| `onClick` | `(event: MouseEvent) => void` | - | Click handler |
| `href` | `string` | - | Redirect URL (makes button a link) |
| `target` | `string` | - | Link target attribute |

---

## Real-world Examples

### Form Submit Button
```tsx
const LoginForm = () => {
  const [isLoading, setIsLoading] = useState(false)
  
  const handleSubmit = async () => {
    setIsLoading(true)
    try {
      await login()
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <form>
      {/* Form fields */}
      <Button 
        htmlType="submit" 
        loading={isLoading}
        onClick={handleSubmit}
      >
        Sign In
      </Button>
    </form>
  )
}
```

### Navigation Button
```tsx
const BackButton = ({ to }: { to: string }) => {
  return (
    <Button 
      type="text" 
      icon={<ArrowLeftOutlined />}
      href={to}
    >
      Back
    </Button>
  )
}
```

---

## Accessibility

### Best Practices
- Always provide meaningful button text or `aria-label` for icon-only buttons
- Ensure sufficient color contrast
- Support keyboard navigation
- Provide loading states for async operations