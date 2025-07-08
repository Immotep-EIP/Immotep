# Input Component Documentation

---

## Overview
The Input component is a wrapper around Ant Design's Input component that provides enhanced functionality with built-in label support, error handling, and consistent form field behavior throughout the Keyz application.

---

## Component Interface

```typescript
import React from 'react'
import { Input as AntInput } from 'antd'
import type { InputProps as AntInputProps } from 'antd'

interface InputProps extends Omit<AntInputProps, 'onChange'> {
  label?: string
  error?: string
  value?: string
  onChange?: (value: string) => void
  required?: boolean
}

const Input: React.FC<InputProps>
```

---

## Basic Usage

### Simple Input
```tsx
import { Input } from '@/components/common'

// Basic input
<Input 
  placeholder="Enter text"
  value={value}
  onChange={setValue}
/>

// With label
<Input 
  label="Full Name"
  placeholder="Enter your full name"
  value={name}
  onChange={setName}
/>
```

### Required Input
```tsx
// Required input with asterisk
<Input 
  label="Email Address"
  placeholder="Enter your email"
  value={email}
  onChange={setEmail}
  required
/>
```

### Input with Error
```tsx
// Input with error state
<Input 
  label="Password"
  type="password"
  placeholder="Enter your password"
  value={password}
  onChange={setPassword}
  error="Password must be at least 8 characters"
  required
/>
```

### Different Input Types
```tsx
// Email input
<Input 
  label="Email"
  type="email"
  placeholder="user@example.com"
  value={email}
  onChange={setEmail}
/>

// Number input
<Input 
  label="Age"
  type="number"
  placeholder="Enter your age"
  value={age}
  onChange={setAge}
/>

// Password input
<Input 
  label="Password"
  type="password"
  placeholder="Enter password"
  value={password}
  onChange={setPassword}
/>
```

---

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `label` | `string` | - | Label text displayed above input |
| `error` | `string` | - | Error message displayed below input |
| `value` | `string` | - | Input value |
| `onChange` | `(value: string) => void` | - | Value change handler |
| `required` | `boolean` | `false` | Show required asterisk in label |
| `type` | `'text' \| 'email' \| 'password' \| 'number' \| ...` | `'text'` | HTML input type |
| `placeholder` | `string` | - | Placeholder text |
| `disabled` | `boolean` | `false` | Disable input |
| `size` | `'large' \| 'middle' \| 'small'` | `'middle'` | Input size |
| `maxLength` | `number` | - | Maximum character length |
| `prefix` | `ReactNode` | - | Prefix icon or element |
| `suffix` | `ReactNode` | - | Suffix icon or element |
| `id` | `string` | - | HTML id attribute |
| `name` | `string` | - | HTML name attribute |

---

## Real-world Examples

### Login Form
```tsx
const LoginForm = () => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [errors, setErrors] = useState<{email?: string, password?: string}>({})

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    // Validation logic
    const newErrors: typeof errors = {}
    
    if (!email) newErrors.email = 'Email is required'
    if (!password) newErrors.password = 'Password is required'
    if (password.length < 8) newErrors.password = 'Password must be at least 8 characters'
    
    setErrors(newErrors)
    
    if (Object.keys(newErrors).length === 0) {
      // Submit form
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <Input
        label="Email Address"
        type="email"
        placeholder="Enter your email"
        value={email}
        onChange={setEmail}
        error={errors.email}
        required
      />
      
      <Input
        label="Password"
        type="password"
        placeholder="Enter your password"
        value={password}
        onChange={setPassword}
        error={errors.password}
        required
      />
      
      <Button htmlType="submit">Sign In</Button>
    </form>
  )
}
```

### Property Form
```tsx
const PropertyForm = ({ property, onSave }: PropertyFormProps) => {
  const [formData, setFormData] = useState({
    title: property?.title || '',
    address: property?.address || '',
    price: property?.price || '',
    description: property?.description || ''
  })
  const [errors, setErrors] = useState<Record<string, string>>({})

  const updateField = (field: string) => (value: string) => {
    setFormData(prev => ({ ...prev, [field]: value }))
    // Clear error when user starts typing
    if (errors[field]) {
      setErrors(prev => ({ ...prev, [field]: '' }))
    }
  }

  return (
    <form>
      <Input
        label="Property Title"
        placeholder="Enter property title"
        value={formData.title}
        onChange={updateField('title')}
        error={errors.title}
        maxLength={100}
        required
      />
      
      <Input
        label="Address"
        placeholder="Enter property address"
        value={formData.address}
        onChange={updateField('address')}
        error={errors.address}
        required
      />
      
      <Input
        label="Monthly Rent"
        type="number"
        placeholder="Enter monthly rent"
        value={formData.price}
        onChange={updateField('price')}
        error={errors.price}
        prefix="$"
        suffix="/month"
        required
      />
      
      <Input
        label="Description"
        placeholder="Enter property description"
        value={formData.description}
        onChange={updateField('description')}
        error={errors.description}
        maxLength={500}
      />
    </form>
  )
}
```

---

## Accessibility

### Best Practices
- Always associate labels with inputs using `htmlFor` and `id`
- Provide meaningful error messages
- Use appropriate input types for better mobile experience
- Support keyboard navigation
- Include proper ARIA attributes for screen readers
