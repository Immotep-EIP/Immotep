# Loader Components

---

## Overview
Loader components provide skeleton loading states for different content types in the Keyz application. They use Ant Design's Skeleton components to create placeholder content that matches the structure of actual data, preventing layout shifts and improving perceived performance.

---

## Available Components

### 1. CardPropertyLoader
Skeleton loader for property cards display.
- **Location**: `./CardPropertyLoader.tsx`
- **Use Case**: Property listings

### 2. CardInventoryLoader
Skeleton loader for inventory/room cards display.
- **Location**: `./CardInventoryLoader.tsx`
- **Use Case**: Room listings, furniture inventory, property details

---

## Component Interfaces

```typescript
interface CardLoaderProps {
  cards: number
}
```

---

## Basic Usage

### Property Cards Loading
```typescript
import CardPropertyLoader from '@/components/ui/Loader/CardPropertyLoader'

const PropertiesPage = () => {
  const { data: properties, isLoading } = useProperties()

  return (
    <div className="properties-grid">
      {isLoading ? (
        <CardPropertyLoader cards={6} />
      ) : (
        properties.map(property => (
          <PropertyCard key={property.id} property={property} />
        ))
      )}
    </div>
  )
}
```

### Inventory Loading
```typescript
import CardInventoryLoader from '@/components/ui/Loader/CardInventoryLoader'

const InventoryTab = ({ propertyId }) => {
  const { data: rooms, isLoading } = useRooms(propertyId)

  return (
    <div className="inventory-container">
      {isLoading ? (
        <CardInventoryLoader cards={4} />
      ) : (
        <RoomsList rooms={rooms} />
      )}
    </div>
  )
}
```

---

## Props Documentation

### CardLoader Props

| Prop | Type | Default | Required | Description |
|------|------|---------|----------|-------------|
| `cards` | `number` | - | ✅ | Number of skeleton property cards to render |

---

## Best Practices

### Do's ✅
- Match skeleton structure to actual content layout
- Use appropriate card counts based on expected data
- Provide loading states for all async operations
- Maintain consistent spacing and dimensions

### Don'ts ❌
- Don't use excessive card counts that impact performance
- Don't mix different loader types in the same context
- Don't forget to handle error states alongside loading states
- Don't use loaders for very fast operations (< 200ms)
