# DividedPage Component Documentation

## Overview
The DividedPage component provides a split-screen layout primarily used for authentication flows, onboarding processes, and marketing pages. It features a 50/50 split layout with responsive behavior and branded header.

## Component Interface

```typescript
import React from 'react'
import logo from '@/assets/images/KeyzLogo.svg'
import style from './DividedPage.module.css'

interface DividedPageProps {
  childrenLeft: React.ReactNode
  childrenRight: React.ReactNode
}

const DividedPage: React.FC<DividedPageProps> = ({
  childrenLeft,
  childrenRight
}) => {
  // Component implementation
}
```

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `childrenLeft` | `ReactNode` | - | Content for the left panel (required) |
| `childrenRight` | `ReactNode` | - | Content for the right panel (required) |

## Key Features

### 1. Split Layout Design
- **Desktop**: 50/50 split between left and right panels
- **Mobile**: Left panel hidden, right panel takes full width

### 2. Branded Header
- Fixed header with Keyz logo and brand name
- Consistent branding across authentication flows
- Positioned in the right panel

### 3. Responsive Behavior
- Mobile-first responsive design
- Left panel completely hidden on mobile devices
- Right panel expands to full width on smaller screens

### 4. Flexible Content Areas
- Left panel: Typically used for marketing content, images, or branding
- Right panel: Used for forms, authentication, or interactive content

## Layout Structure

### Desktop Layout (>= 900px)
```
┌─────────────────────────────────────────────────────────┐
│                    DividedPage                          │
├─────────────────────────┬───────────────────────────────┤
│ Left Panel (50%)        │ Right Panel (50%)             │
│ ┌─────────────────────┐ │ ┌─────────────────────────────┐
│ │                     │ │ │ Header                      │ 
│ │                     │ │ │ ┌─────┐ Keyz                │
│ │   childrenLeft      │ │ │ │Logo │                     │
│ │                     │ │ │ └─────┘                     │ 
│ │                     │ │ ├─────────────────────────────┤
│ │                     │ │ │                             │
│ │                     │ │ │      childrenRight          │
│ │                     │ │ │                             │
│ └─────────────────────┘ │ └─────────────────────────────┘
└─────────────────────────┴───────────────────────────────┘
```

### Mobile Layout (< 900px)
```
┌─────────────────────────────────────┐
│           DividedPage               │
│ ┌─────────────────────────────────┐ │
│ │ Header                          │ │
│ │ ┌─────┐ Keyz                    │ │
│ │ │Logo │                         │ │
│ │ └─────┘                         │ │
│ ├─────────────────────────────────┤ │
│ │                                 │ │
│ │        childrenRight            │ │
│ │       (Full Width)              │ │
│ │                                 │ │
│ └─────────────────────────────────┘ │
└─────────────────────────────────────┘
```

## Real-world Usage Examples

### Login Page
```typescript
import DividedPage from '@/components/layout/DividedPage/DividedPage'
import LoginForm from '@/components/features/auth/LoginForm'
import WelcomeGraphic from '@/components/common/WelcomeGraphic'

const LoginPage = () => {
  return (
    <DividedPage
      childrenLeft={
        <div className="welcome-content">
          <WelcomeGraphic />
          <h2>Welcome to Keyz</h2>
          <p>Manage your properties with ease</p>
        </div>
      }
      childrenRight={<LoginForm />}
    />
  )
}
```

### Registration Page
```typescript
import DividedPage from '@/components/layout/DividedPage/DividedPage'
import RegisterForm from '@/components/features/auth/RegisterForm'
import FeaturesList from '@/components/common/FeaturesList'

const RegisterPage = () => {
  return (
    <DividedPage
      childrenLeft={
        <div className="features-showcase">
          <h2>Why Choose Keyz?</h2>
          <FeaturesList 
            features={[
              'Property Management',
              'Tenant Communication',
              'Maintenance Tracking',
              'Financial Reports'
            ]}
          />
        </div>
      }
      childrenRight={<RegisterForm />}
    />
  )
}
```

## Accessibility Features

### Semantic Structure
- Proper heading hierarchy
- Descriptive alt text for logo
- Proper focus management
