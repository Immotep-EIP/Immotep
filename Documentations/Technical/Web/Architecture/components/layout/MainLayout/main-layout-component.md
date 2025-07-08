# MainLayout Component Documentation

## Overview
The MainLayout component is the primary application layout that provides navigation structure, responsive behavior, and consistent UI patterns across all authenticated pages in the Keyz application.

## Component Interface

```typescript
import React, { useState } from 'react'
import { useTranslation } from 'react-i18next'
import { useNavigate, Outlet, useLocation } from 'react-router-dom'
import { Layout, Menu } from 'antd'

const MainLayout: React.FC = () => {
  // Component logic
}
```

## Key Features

### 1. Responsive Navigation
- **Desktop**: Fixed sidebar navigation with menu items
- **Mobile**: Hamburger menu with dropdown overlay

### 2. Route-based Active States
- Automatically highlights current page in navigation
- Uses React Router location for state management
- Visual feedback for user orientation

### 3. Internationalization Support
- All menu labels are translatable
- Dynamic language switching
- RTL layout compatibility

### 4. Accessibility Features
- Keyboard navigation support
- ARIA attributes for screen readers
- Focus management

## Navigation Structure

### Menu Items Configuration
```typescript
const items = [
  {
    label: 'components.button.overview',
    key: NavigationEnum.OVERVIEW,
    icon: <img src={Overview} alt="Overview" className={style.menuIcon} />
  },
  {
    label: 'components.button.real_property',
    key: NavigationEnum.REAL_PROPERTY,
    icon: <img src={Property} alt="Real Property" className={style.menuIcon} />
  },
  {
    label: 'components.button.messages',
    key: NavigationEnum.MESSAGES,
    icon: <img src={Messages} alt="Messages" className={style.menuIcon} />
  },
  {
    label: 'components.button.settings',
    key: NavigationEnum.SETTINGS,
    icon: <img src={Settings} alt="Settings" className={style.menuIcon} />
  }
]
```

### Navigation Enum Values
```typescript
enum NavigationEnum {
  OVERVIEW = '/overview',
  REAL_PROPERTY = '/properties',
  MESSAGES = '/messages',
  SETTINGS = '/settings'
}
```

## Layout Structure

### Desktop Layout (>= 900px)
```
┌─────────────────────────────────────┐
│ Header (Fixed)                      │
│ ┌─────┐ Keyz                        │
│ │Logo │                             │
│ └─────┘                             │
├─────────────┬───────────────────────┤
│ Sidebar     │ Content Area          │
│ ┌─────────┐ │ ┌─────────────────┐   │
│ │Overview │ │ │                 │   │
│ │Properties │ │    <Outlet />   │   │
│ │Messages │ │ │                 │   │
│ │Settings │ │ │                 │   │
│ └─────────┘ │ └─────────────────┘   │
└─────────────┴───────────────────────┘
```

### Mobile Layout (< 900px)
```
┌─────────────────────────────────────┐
│ Header (Fixed)                   ☰  │
│ ┌─────┐ Keyz                        │
│ │Logo │                             │
│ └─────┘                             │
├─────────────────────────────────────┤
│ Content Area (Full Width)           │
│ ┌─────────────────────────────────┐ │
│ │                                 │ │
│ │           <Outlet />            │ │
│ │                                 │ │
│ └─────────────────────────────────┘ │
└─────────────────────────────────────┘
```

## State Management

### Mobile Menu State
```typescript
const [menuOpen, setMenuOpen] = useState(false)

const toggleMenu = () => {
  setMenuOpen(!menuOpen)
}
```

### Active Route Detection
```typescript
const location = useLocation()
const currentLocation = `/${location.pathname.split('/')[1] || ''}`
```

### Navigation Handler
```typescript
const onClick: MenuProps['onClick'] = e => {
  navigate(e.key)
}
```

## Real-world Usage

### Router Configuration
```typescript
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import MainLayout from '@/components/layout/MainLayout/MainLayout'

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        {/* Public routes */}
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        
        {/* Protected routes with MainLayout */}
        <Route path="/" element={<MainLayout />}>
          <Route index element={<Navigate to="/overview" replace />} />
          <Route path="/overview" element={<OverviewPage />} />
          <Route path="/properties" element={<PropertiesPage />} />
          <Route path="/properties/:id" element={<PropertyDetailPage />} />
          <Route path="/messages" element={<MessagesPage />} />
          <Route path="/settings" element={<SettingsPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}
```
