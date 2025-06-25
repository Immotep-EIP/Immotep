# Common Components Documentation

## Overview
Common components are reusable UI elements that form the foundation of the Keyz application. These components are designed to be highly reusable, configurable, and maintainable across different features.

## Centralized Exports
All common components are exported through a centralized `index.ts` file for easy importing:

```typescript
// Code/Web/src/components/common/index.ts
export { default as Button } from './Button/Button'
export { default as Input } from './Input/Input'
export { default as Card } from './Card/Card'
export { default as StatusTag } from './Tag/StatusTag'
export { default as Empty } from './Empty/Empty'
export { default as Badge } from './Badge/Badge'
```

## Available Components

### 1. Button
Enhanced button component with consistent defaults and accessibility features.
- **Location**: `./Button/Button.tsx`
- **Documentation**: [Button Component](./Button/button-component.md)
- **Key Features**: Custom defaults (primary type, middle size), loading states, accessibility

### 2. Input
Enhanced input component with built-in label, error handling, and form integration.
- **Location**: `./Input/Input.tsx`
- **Documentation**: [Input Component](./Input/input-component.md)
- **Key Features**: Label support, error display, required indicators, simplified onChange

### 3. Card
Enhanced card component with custom variants and padding options.
- **Location**: `./Card/Card.tsx`
- **Documentation**: [Card Component](./Card/card-component.md)
- **Key Features**: Custom variants (elevated, outlined), flexible padding, Card.Grid support

### 4. StatusTag
Specialized tag component for displaying internationalized status with color mapping.
- **Location**: `./Tag/StatusTag.tsx`
- **Documentation**: [StatusTag Component](./StatusTag/status-tag-component.md)
- **Key Features**: i18n support, color mapping, case-insensitive matching

### 5. Empty
Enhanced empty state component with custom image and consistent styling.
- **Location**: `./Empty/Empty.tsx`
- **Documentation**: [Empty Component](./Empty/empty-component.md)
- **Key Features**: Custom empty image, Typography wrapper, flexible descriptions

### 6. Badge
Enhanced badge component with status variants and ribbon support.
- **Location**: `./Badge/Badge.tsx`
- **Documentation**: [Badge Component](./Badge/badge-component.md)
- **Key Features**: Status badges, Ribbon support, custom colors, all Ant Design features

## Design Principles

### 1. Props Design
- Use TypeScript interfaces for all props
- Provide sensible defaults
- Make props explicit
- Use consistent naming conventions

### 2. Styling
- Use CSS Modules for component-specific styles
- Ensure responsive design

### 3. Accessibility
- Implement proper ARIA attributes
- Ensure keyboard navigation
- Support screen readers

## Implementation Guidelines

### Component Structure
```typescript
// 1. Import statements
import React from 'react';
import styles from './Component.module.css';

interface ComponentProps extends AntComponentProps {
  yourProps: props
}

// 2. Component implementation
export const Component: React.FC<ComponentProps> = ({
  // Props destructuring
}) => {
  return (
    <div className={styles.container}>
      {/* Component content */}
    </div>
  );
};
```

## Usage Examples

### Basic Import Pattern
```typescript
import { Button, Input, Card, StatusTag, Empty, Badge } from '@/components/common'
```

## Design Principles

### Consistency
- All components follow the same API patterns
- Consistent prop naming conventions
- Unified styling approach

### Accessibility
- ARIA attributes included by default
- Keyboard navigation support
- Screen reader compatibility
- Sufficient color contrast

### Internationalization
- Components support i18n where applicable
- Text content can be translated

### Customization
- Flexible prop interfaces
- Custom styling support
- Extensible through composition

## Best Practices

### Error Handling
```typescript
// ✅ Consistent error display
<Input
  label="Email"
  value={email}
  onChange={setEmail}
  error={validationErrors.email}
  required
/>
```

## Contributing

When adding new common components:

1. **Create the component** in a new subdirectory with proper TypeScript interfaces
2. **Add comprehensive tests** with full coverage
3. **Export the component** in the main `index.ts` file
4. **Create documentation** following the established template
5. **Update this overview** to include the new component

### Component Template Structure
```
src/components/common/NewComponent/
├── NewComponent.tsx          # Main component file
├── __tests__/
│   └── NewComponent.test.tsx # Test file
└── index.ts                 # Component export (optional)
```

## Conclusion
Common components are the building blocks of the Keyz application. Following these guidelines ensures consistency, maintainability, and reusability across the application. Always refer to this documentation when creating or modifying common components. 