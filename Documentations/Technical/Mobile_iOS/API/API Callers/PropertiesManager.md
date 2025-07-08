# TenantPropertyViewModel

## Overview

`TenantPropertyViewModel` is an `ObservableObject` that manages tenant-specific property-related data and API interactions in a SwiftUI application. It handles fetching property details, rooms, damages, documents, and pictures, as well as creating damages and uploading documents. It supports tenant-specific endpoints and maintains state for UI updates.

---

## Data Structures

### `PropertyRoomsTenant`

```swift
struct PropertyRoomsTenant: Identifiable, Equatable {
    let id: String
    let name: String
}
```

Represents a room in a property for tenant view, with an identifier and name.

---

### `FurnitureResponseTenant`

```swift
struct FurnitureResponseTenant: Codable {
    let id: String
    let name: String
    let quantity: Int
    let archived: Bool
}
```

Represents furniture details in a room for tenant view.

---

### `RoomResponseTenant`

```swift
struct RoomResponseTenant: Codable {
    let id: String
    let name: String
    let archived: Bool
    let furnitures: [FurnitureResponseTenant]
}
```

Represents a room with its furniture details for tenant view.

---

### `PropertyInventoryResponse`

```swift
struct PropertyInventoryResponse: Codable {
    let id: String
    let ownerId: String
    let name: String
    let address: String
    let city: String
    let postalCode: String
    let country: String
    let areaSqm: Double
    let rentalPricePerMonth: Int
    let depositPrice: Int
    let createdAt: String
    let archived: Bool
    let nbDamage: Int
    let status: String
    let lease: LeaseInfo?
    let rooms: [RoomResponseTenant]
}
```

Represents the full inventory response for a property, including rooms and lease information.

---

### `PropertyDocument`

```swift
struct PropertyDocument {
    let id: String
    let title: String
    let fileName: String
    let data: String
}
```

Represents a document associated with a property or lease, including base64 data.

---

### `DamageRequest`

```swift
struct DamageRequest {
    let description: String
    let roomId: String
    let priority: String
    let photos: [String]
}
```

Represents input data for creating a damage report.

---

### `DamageResponse`

```swift
struct DamageResponse {
    let id: String
    let description: String
    let roomId: String
    let priority: String
    let status: String
    let photos: [String]
    let fixed: Bool
}
```

Represents a damage report returned by the API.

---

### `PropertyImageBase64`

```swift
struct PropertyImageBase64 {
    let data: String
}
```

Represents a property picture in base64 format.

---

### `IdResponse`

```swift
struct IdResponse: Codable {
    let id: String
}
```

Represents an API response containing an identifier for created resources.

---

### `PropertyResponse`

```swift
struct PropertyResponse {
    let id: String
    let ownerId: String
    let name: String
    let address: String
    let city: String
    let postalCode: String
    let country: String
    let areaSqm: Double
    let rentalPricePerMonth: Int
    let depositPrice: Int
    let isAvailable: Bool
    let createdAt: String
    let lease: LeaseInfo?
}
```

Represents a property object fetched from the API.

---

### `LeaseResponse`

```swift
struct LeaseResponse {
    let id: String
    let active: Bool
    let propertyId: String
    let startDate: Date
    let endDate: Date?
    let tenantName: String
}
```

Represents lease details associated with a property.

---

### `LeaseInfo`

```swift
struct LeaseInfo {
    let tenantName: String
    let startDate: Date
    let endDate: Date?
}
```

Represents basic lease information included in property responses.

---

## TenantPropertyViewModel Properties

* `@Published var damages: [DamageResponse]`
  Array of damage reports for the current property.

* `@Published var isFetching166Damages: Bool`
  Indicates if damages are being fetched.

* `@Published var damagesError: String?`
  Stores error messages for damage fetching failures.

* `@Published var rooms: [PropertyRoomsTenant]`
  Array of rooms in the current property.

* `private var tenantProperty: Property?`
  Cached property details.

* `public var activeLeaseId: String?`
  Identifier of the active lease for the current property.

* `private var isFetchingProperty: Bool`
  Indicates if property details are being fetched.

* `private var isFetchingRooms: Bool`
  Indicates if rooms are being fetched.

* `private var isFetchingLease: Bool`
  Indicates if lease details are being fetched.

* `weak var propertyViewModel: PropertyViewModel?`
  Reference to a parent view model for coordination.

---

## TenantPropertyViewModel Functions

* `func fetchTenantProperty() async throws -> Property`
  Fetches and caches tenant property details, including lease, documents, damages, rooms, and picture.

* `func fetchPropertyRooms(token: String) async throws -> [PropertyRoomsTenant]`
  Fetches and caches room details for the current property.

* `func fetchPropertiesPicture(propertyId: String) async throws -> UIImage?`
  Fetches and caches the property picture, returning a `UIImage` or `nil`.

* `func fetchTenantPropertyDocuments(leaseId: String, propertyId: String) async throws -> [PropertyDocument]`
  Fetches documents associated with the property or lease.

* `@MainActor func fetchTenantDamages(leaseId: String, fixed: Bool? = nil) async throws -> [DamageResponse]`
  Fetches damage reports for the current lease, optionally filtered by fixed status.

* `func fetchActiveLeaseIdForProperty(propertyId: String, token: String) async throws -> String?`
  Fetches and caches the active lease ID for the property.

* `func createDamage(propertyId: String, leaseId: String, damage: DamageRequest, token: String) async throws -> String`
  Creates a new damage report and returns its ID.

* `private func performFetchActiveLeaseId(propertyId: String, token: String) async`
  Internal function to fetch and set the active lease ID.

* `func fetchDamageByID(damageId: String, token: String) async throws -> DamageResponse`
  Fetches details for a specific damage report by ID.

* `func fixDamage(damageId: String, token: String) async throws`
  Marks a damage report as fixed.

* `func uploadTenantDocument(leaseId: String, propertyId: String, fileName: String, base64Data: String) async throws -> String`
  Uploads a document for the property or lease and returns its ID.

---

## Dependencies

* `URLSession` for network requests.
* `JSONEncoder` and `JSONDecoder` for data serialization.
* `TokenStorage` for managing access tokens.
* `ImageCache` for caching property images.
* `APIConfig` for base URL configuration.
