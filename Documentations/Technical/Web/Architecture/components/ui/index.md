# UI Components Documentation

## Overview
UI components are basic building blocks that provide fundamental user interface elements for the Keyz application. These components handle typography, loading states, meta information, and specialized UI patterns that enhance user experience and maintain visual consistency.

## Available Components

### 1. PageText (Title)
Typography component for consistent page titles and headings.
- **Location**: `./PageText/Title.tsx`

```typescript
// PageTitle for consistent headings
<PageTitle 
  title="Dashboard Overview"
  size="title"
  margin={true}
/>

// Subtitle variant
<PageTitle 
  title="Property Details"
  size="subtitle"
  margin={false}
/>
```

### 2. PageMeta
SEO and meta information management component using React Helmet.
- **Location**: `./PageMeta/PageMeta.tsx`

```typescript
// PageMeta for SEO
<PageMeta
  title="Property Management Dashboard"
  description="Manage your properties efficiently with Keyz"
  keywords="property, management, real estate, keyz"
/>
```

### 3. Loader Components
Skeleton loading components for different content types.
- **Location**: `./Loader/`

```typescript
// Property cards loading state
<CardPropertyLoader cards={6} />

// Inventory loading state
<CardInventoryLoader cards={4} />
```

## Design Principles

### 1. Typography Consistency
- Standardized font families (Jost)
- Consistent sizing hierarchy
- Proper color contrast
- Responsive text scaling

### 2. Loading States
- Skeleton patterns that match content structure
- Consistent animation timing
- Proper accessibility for screen readers
- Performance-optimized rendering

### 3. Semantic Structure
- Proper HTML semantics
- Accessibility attributes
- Screen reader compatibility
- Keyboard navigation support

### 4. Internationalization
- Translation key integration
- RTL layout support
- Dynamic content adaptation
- Consistent translation patterns

## Accessibility Features

### Typography Accessibility
- Proper heading hierarchy (h1, h2, h3)
- Sufficient color contrast ratios
- Scalable font sizes
- Screen reader compatible

### Loading State Accessibility
- ARIA live regions for dynamic content
- Skeleton components with proper labels
- Screen reader announcements
- Keyboard navigation preservation

### Success Page Accessibility
- Proper focus management
- Success announcements for screen readers
- Semantic HTML structure
- Clear visual hierarchy

## Internationalization

### Translation Keys Structure
```json
{
  "property": {
    "details": {
      "address": "Property Address",
      "description": "Description"
    },
    "form": {
      "basic_info": "Basic Information",
      "details": "Property Details"
    }
  },
  "pages": {
    "login_tenant": {
      "login_success": "Login Successful",
      "login_success_description": "Welcome back to Keyz"
    },
    "register_tenant": {
      "register_success": "Registration Successful",
      "register_success_description": "Your account has been created"
    }
  }
}
```

### Translation Usage
```typescript
// SubtitledElement automatically translates
<SubtitledElement subtitleKey="property.details.address">
  {/* Content */}
</SubtitledElement>

// Success pages use useTranslation hook
const { t } = useTranslation()
<Result
  title={t('pages.login_tenant.login_success')}
  subTitle={t('pages.login_tenant.login_success_description')}
/>
```

## Best Practices

### Typography Usage
```typescript
// ✅ Use PageTitle for consistent headings
<PageTitle title="Dashboard" size="title" />

// ✅ Use appropriate sizes
<PageTitle title="Section Header" size="subtitle" />

// ❌ Avoid inline styles for typography
<h1 style={{ fontSize: '1.4rem' }}>Title</h1>
```

### Loading States
```typescript
// ✅ Match skeleton structure to actual content
{loading ? (
  <CardPropertyLoader cards={properties.length || 6} />
) : (
  <PropertyCards properties={properties} />
)}

// ✅ Provide meaningful loading counts
<CardInventoryLoader cards={expectedRoomCount} />

// ❌ Don't use generic loading spinners for structured content
{loading ? <Spin /> : <ComplexCardLayout />}
```

### Meta Information
```typescript
// ✅ Provide descriptive titles and descriptions
<PageMeta
  title="Property Details - 123 Main St"
  description="View and manage details for this property"
  keywords="property, management, details"
/>

// Always include PageMeta even for dynamic content
```

## Contributing

When adding new UI components:

1. **Follow naming conventions** (PascalCase for components)
2. **Include TypeScript interfaces** for all props
3. **Add CSS Modules** for styling
4. **Support internationalization** where applicable
5. **Include accessibility features** by default
6. **Write comprehensive tests** with good coverage
7. **Document usage patterns** and examples

### UI Component Structure
```
src/components/ui/NewComponent/
├── NewComponent.tsx           # Main component file
├── NewComponent.module.css    # Scoped styles
├── __tests__/
│   └── NewComponent.test.tsx  # Test file
```

## Conclusion
UI components form the foundational layer of the Keyz application's user interface, providing consistent typography, loading states, and specialized UI patterns. Following these guidelines ensures proper component usage, accessibility compliance, and maintainable code structure. Always refer to this documentation when creating or modifying UI components to maintain consistency with the overall design system.
