# Layout Components Documentation

---

## Overview
Layout components are structural components that define the overall page structure and navigation patterns in the Keyz application. These components handle responsive behavior, layout-specific state, and provide consistent user interface patterns across different views.

---

## Available Components

### 1. MainLayout
Main application layout with navigation sidebar and responsive mobile menu.
- **Location**: `./MainLayout/MainLayout.tsx`
- **Documentation**: [MainLayout Component](./MainLayout/main-layout-component.md)
- **Key Features**: Responsive navigation, mobile hamburger menu, route-based active states, internationalization

### 2. DividedPage
Split-screen layout component for authentication and onboarding flows.
- **Location**: `./DividedPage/DividedPage.tsx`
- **Documentation**: [DividedPage Component](./DividedPage/divided-page-component.md)
- **Key Features**: 50/50 split layout, responsive design, branded header, authentication flows

---

## Design Principles

### 1. Responsive Design
- Mobile-first approach with progressive enhancement
- Breakpoint-based layout adjustments
- Touch-friendly navigation on mobile devices
- Proper viewport handling

### 2. Navigation Patterns
- Consistent navigation structure across the application
- Clear visual indicators for active routes
- Accessible keyboard navigation
- Internationalized menu labels

### 3. Layout Consistency
- Standardized spacing and proportions
- Consistent header and branding placement
- Proper content area management

### 4. Accessibility
- ARIA attributes for navigation elements
- Keyboard navigation support
- Screen reader compatibility

---

## Layout Structure

### MainLayout Structure
```
MainLayout
├── Header (Fixed)
│   ├── Logo & Brand
│   └── Mobile Menu Toggle
├── Sidebar Navigation (Desktop)
│   ├── Overview
│   ├── Properties
│   ├── Messages
│   └── Settings
├── Mobile Dropdown Menu
└── Content Area (Outlet)
```

### DividedPage Structure
```
DividedPage
├── Left Panel (50% width)
│   └── Custom Content
└── Right Panel (50% width)
    ├── Header (Logo & Brand)
    └── Content Area
```

---

## Navigation Configuration

### Menu Items Structure
```typescript
const items = [
  {
    label: 'components.button.overview',
    key: NavigationEnum.OVERVIEW,
    icon: <img src={Overview} alt="Overview" />
  },
  {
    label: 'components.button.real_property', 
    key: NavigationEnum.REAL_PROPERTY,
    icon: <img src={Property} alt="Real Property" />
  },
  {
    label: 'components.button.messages',
    key: NavigationEnum.MESSAGES,
    icon: <img src={Messages} alt="Messages" />
  },
  {
    label: 'components.button.settings',
    key: NavigationEnum.SETTINGS,
    icon: <img src={Settings} alt="Settings" />
  }
]
```

---

## Internationalization

### Translation Keys
Layout components use i18n for menu labels:

```json
{
  "components": {
    "button": {
      "overview": "Overview",
      "real_property": "Properties", 
      "messages": "Messages",
      "settings": "Settings"
    }
  }
}
```

---

## Accessibility Features

### Keyboard Navigation
- Tab navigation through menu items
- Enter/Space key activation
- ARIA labels

### Screen Reader Support
- Descriptive alt text for icons
- Proper heading hierarchy

---

## Contributing

When modifying layout components:

1. **Maintain responsive behavior** across all breakpoints
2. **Test keyboard navigation** thoroughly
3. **Verify internationalization** works for all text
4. **Check accessibility** with screen readers
5. **Test on mobile devices** for touch interactions

### Layout Component Structure
```
src/components/layout/NewLayout/
├── NewLayout.tsx              # Main component file
├── NewLayout.module.css       # Scoped styles
├── __tests__/
│   └── NewLayout.test.tsx     # Test file
```

---

## Conclusion
Layout components are the structural foundation of the Keyz application, providing consistent navigation patterns and responsive design across all views. Following these guidelines ensures proper layout management, accessibility compliance, and maintainable code structure. Always refer to this documentation when creating or modifying layout components to maintain consistency with the overall application architecture.