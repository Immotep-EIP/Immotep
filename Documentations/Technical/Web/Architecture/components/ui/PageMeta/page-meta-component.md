# PageMeta Component

## Overview
The PageMeta component manages SEO and meta information for pages using React Helmet. It provides dynamic page titles, meta descriptions, and keywords to improve search engine optimization and social media sharing.

## Component Interface

```typescript
interface PageMetaProps {
  title: string
  description?: string
  keywords?: string
}

const PageMeta: React.FC<PageMetaProps>
```

## Basic Usage

### Simple Page Title
```typescript
import PageMeta from '@/components/ui/PageMeta/PageMeta'

const DashboardPage = () => {
  return (
    <>
      <PageMeta title="Dashboard" />
      <div>
        {/* Page content */}
      </div>
    </>
  )
}
```

### Complete Meta Information
```typescript
const PropertyDetailsPage = ({ property }) => {
  return (
    <>
      <PageMeta
        title={`${property.name} - Property Details`}
        description={`View details for ${property.name} located at ${property.address}`}
        keywords="property, real estate, rental, management, details"
      />
      <div>
        {/* Page content */}
      </div>
    </>
  )
}
```

## Props Documentation

| Prop | Type | Default | Required | Description |
|------|------|---------|----------|-------------|
| `title` | `string` | - | ✅ | Page title (will be prefixed with "Keyz") |
| `description` | `string` | - | ❌ | Meta description for SEO |
| `keywords` | `string` | - | ❌ | Meta keywords for SEO |
