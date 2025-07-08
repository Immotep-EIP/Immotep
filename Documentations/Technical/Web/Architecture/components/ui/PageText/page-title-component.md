# PageTitle Component

## Overview
The PageTitle component provides consistent typography for page titles and headings throughout the Keyz application. It supports flexible sizing, optional margins, and uses the standardized Jost font family for brand consistency.

## Component Interface

```typescript
interface PageTitleProps {
  title: string
  size?: 'title' | 'subtitle'
  margin?: boolean
}

const PageTitle: React.FC<PageTitleProps>
```

## Basic Usage

### Standard Title
```typescript
import PageTitle from '@/components/ui/PageText/Title'

const DashboardPage = () => {
  return (
    <div>
      <PageTitle title="Dashboard Overview" size="title" />
      {/* Page content */}
    </div>
  )
}
```

### Subtitle Variant
```typescript
const PropertySection = () => {
  return (
    <section>
      <PageTitle title="Recent Properties" size="subtitle" />
      {/* Section content */}
    </section>
  )
}
```

### Without Margin
```typescript
const CompactHeader = () => {
  return (
    <div className="compact-layout">
      <PageTitle 
        title="Settings" 
        size="title" 
        margin={false}
      />
      {/* Immediately following content */}
    </div>
  )
}
```

## Props Documentation

| Prop | Type | Default | Required | Description |
|------|------|---------|----------|-------------|
| `title` | `string` | - | ✅ | The text content to display as the title |
| `size` | `'title' \| 'subtitle'` | `'title'` | ❌ | Determines the font size and weight |
| `margin` | `boolean` | `true` | ❌ | Whether to apply bottom margin |

### Size Variants

| Size | Font Size | Font Weight | Use Case |
|------|-----------|-------------|----------|
| `title` | `1.4rem` | `500` | Main page headings, primary titles |
| `subtitle` | `1rem` | `400` | Section headers, secondary titles |

## Real-world Examples

### Dashboard Layout
```typescript
const Dashboard = () => {
  return (
    <div className="dashboard">
      <PageTitle title="Property Management Dashboard" size="title" />
      
      <div className="dashboard-sections">
        <section className="properties-section">
          <PageTitle title="My Properties" size="subtitle" />
          <PropertyGrid />
        </section>
        
        <section className="messages-section">
          <PageTitle title="Recent Messages" size="subtitle" />
          <MessageList />
        </section>
      </div>
    </div>
  )
}
```

### Property Details Page
```typescript
const PropertyDetails = ({ property }) => {
  return (
    <div className="property-details">
      <PageTitle 
        title={`${property.name} - Details`} 
        size="title" 
      />
      
      <div className="property-sections">
        <PageTitle 
          title="Basic Information" 
          size="subtitle" 
          margin={false}
        />
        <PropertyInfo property={property} />
        
        <PageTitle title="Inventory" size="subtitle" />
        <InventoryList propertyId={property.id} />
      </div>
    </div>
  )
}
```

## Typography Hierarchy

### Recommended Usage Patterns
```typescript
// Page level hierarchy
<PageTitle title="Property Management" size="title" />     // h1 level
<PageTitle title="My Properties" size="subtitle" />        // h2 level
<PageTitle title="Property Details" size="subtitle" />     // h3 level

// Content hierarchy
const ContentHierarchy = () => {
  return (
    <article>
      {/* Main article title */}
      <PageTitle title="Property Investment Guide" size="title" />
      
      {/* Section headings */}
      <PageTitle title="Getting Started" size="subtitle" />
      <PageTitle title="Advanced Strategies" size="subtitle" />
      <PageTitle title="Risk Management" size="subtitle" />
    </article>
  )
}
```
