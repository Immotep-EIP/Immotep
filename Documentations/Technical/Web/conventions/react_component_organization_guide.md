# React Component Organization Guide

## Directory Structure

```
src/
  components/
    common/                  # Generic reusable components
      Button/
      Input/
      ...
    layout/                  # Layout components
      MainLayout/
      DividedPage/
      ...
    features/                # Feature-specific components
      RealProperty/
        details/
          DetailsPart/
          PropertyHeader/
          PropertyInfo/
          PropertyImage/
        list/
        ...
      Messages/
      Authentication/
        AuthentificationPage/
        ProtectedRoute/
        ...
    ui/                      # UI components
      PageMeta/
      PageText/
      SubtitledElement/
      Loader/
      SuccesPage/
      ...
```

## Component Categories

### Common Components
Generic, reusable components that can be used across the entire application. These components are not tied to any specific feature and should be highly reusable.

Examples: `Button`, `Input`, `Card`, `Modal`, etc.

### Layout Components
Components that define the structure of your pages. These components typically handle the overall layout, navigation, and page structure.

Examples: `MainLayout`, `DividedPage`, `Sidebar`, `Header`, `Footer`, etc.

### Feature Components
Components specific to a particular feature or domain of your application. These components are typically not reusable outside their specific feature.

Examples: `PropertyDetails`, `AuthentificationPage`, `Messages`, etc.

### UI Components
UI elements that are specific to your application's design system but are not tied to a specific feature.

Examples: `PageMeta`, `PageText`, `SubtitledElement`, `Loader`, etc.

## Naming Conventions

### PascalCase for Component Folders
Use PascalCase for folders containing React components:

- ✅ `Button/`, `UserProfile/`, `PropertyDetails/`
- ❌ `button/`, `userProfile/`, `propertyDetails/`

### camelCase for Utility Folders
Use camelCase for folders that don't contain React components:

- ✅ `utils/`, `hooks/`, `services/`
- ❌ `Utils/`, `Hooks/`, `Services/`

## Import Conventions

Use absolute imports with aliases for better readability:

```typescript
// Before
import Button from '../../../components/common/Button';

// After
import Button from '@components/common/Button';
```

## Best Practices

1. **One component per file**: Each component should have its own file
2. **Co-location**: Keep related files close to each other
3. **Separation of concerns**: Split UI and logic into separate components
4. **Reusability**: Design components to be reusable when possible
5. **Consistency**: Follow consistent patterns across your codebase

## Benefits

- **Improved maintainability**: Easier to find and modify components
- **Better organization**: Clear structure for new developers
- **Enhanced reusability**: Components are easier to reuse
- **Scalability**: Structure that grows with your application
- **Collaboration**: Easier for teams to work together

This organization approach follows industry best practices and will help maintain a clean, scalable React codebase. 