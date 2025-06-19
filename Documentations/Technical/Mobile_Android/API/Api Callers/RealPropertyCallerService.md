# RealPropertyCallerService

## Overview

`RealPropertyCallerService` extends the base [`ApiCallerService`](./ApiCallerService.md) to handle property-related API calls such as fetching, adding, updating properties, managing pictures, documents, and archives. It supports roles (owner vs tenant) by switching endpoints accordingly.

---

## Enum Classes

### `PropertyStatus`

```kotlin
enum class PropertyStatus {
    unavailable,
    available,
    invite_sent
}
```

Represents the status of a property.

---

### `DamageStatus`

```kotlin
enum class DamageStatus {
    PENDING,
    PLANNED,
    AWAITING_OWNER_CONFIRMATION,
    AWAITING_TENANT_CONFIRMATION,
    FIXED
}
```

Represents the current status of a damage report.

---

### `DamagePriority`

```kotlin
enum class DamagePriority {
    low,
    medium,
    high,
    urgent
}
```

Represents the priority level for a damage report.

---

## Enum conversion functions

* `stringToPropertyState(str: String): PropertyStatus`
* `stringToDamageStatus(str: String): DamageStatus`
* `stringToDamagePriority(str: String): DamagePriority`

These functions map string values from the API to corresponding enums.

---

## Data Classes

### `AddPropertyInput`

```kotlin
data class AddPropertyInput(
    val name: String = "",
    val address: String = "",
    val city: String = "",
    val postal_code: String ="",
    val country: String = "",
    val area_sqm: Double = 0.0,
    val rental_price_per_month: Int = 0,
    val deposit_price: Int = 0,
    val apartment_number: String = ""
) {
    fun toDetailedProperty(
        id: String,
        currentLease: LeaseDetailedProperty? = null,
        currentInvite: InviteDetailedProperty? = null
    ): DetailedProperty { ... }
}
```

Represents input data to add or update a property.
Includes a helper method to convert itself to a `DetailedProperty` by supplying an `id` and optional lease/invite details.

---

### `DocumentInput`

```kotlin
data class DocumentInput(
    val name: String = "",
    val data: String = ""
) {
    fun toDocument(id: String): Document { ... }
}
```

Represents input data for uploading a document attached to a property or lease.
Includes helper to convert to a `Document` with an assigned id.

---

### `UpdatePropertyPictureInput`

```kotlin
data class UpdatePropertyPictureInput(
    val data: String
)
```

Data class representing the image data (likely Base64 or similar) to update a property's picture.

---

### `InvitePropertyResponse`

```kotlin
data class InvitePropertyResponse(
    val end_date: String?,
    val start_date: String,
    val tenant_email: String
) {
    fun toInviteDetailedProperty(): InviteDetailedProperty { ... }
}
```

Represents invitation details for a tenant to a property.
Includes conversion to `InviteDetailedProperty`.

---

### `PropertyPictureResponse`

```kotlin
data class PropertyPictureResponse(
    val id: String,
    val created_at: String,
    val data: String
)
```

Represents a property picture returned by the API.

---

### `LeasePropertyResponse`

```kotlin
data class LeasePropertyResponse(
    val active: Boolean,
    val end_date: String?,
    val id: String,
    val start_date: String,
    val tenant_email: String,
    val tenant_name: String
) {
    fun toLeaseDetailedProperty(): LeaseDetailedProperty { ... }
}
```

Represents lease details associated with a property.
Includes conversion to `LeaseDetailedProperty`.

---

### `GetPropertyResponse`

```kotlin
data class GetPropertyResponse(
    val id: String,
    val apartment_number: String?,
    val archived: Boolean,
    val owner_id: String,
    val name: String,
    val address: String,
    val city: String,
    val postal_code: String,
    val country: String,
    val area_sqm: Double,
    val rental_price_per_month: Int,
    val deposit_price: Int,
    val created_at: String,
    val status: String,
    val nb_damage: Int,
    val picture_id: String?,
    val invite: InvitePropertyResponse?,
    val lease: LeasePropertyResponse?
) {
    fun toDetailedProperty(): DetailedProperty { ... }
}
```

Represents a property object fetched from the API.
Provides helper to convert itself to a domain model `DetailedProperty` including lease and invite info depending on status and ownership.

---

### `Document`

```kotlin
data class Document(
    val id: String,
    val name: String,
    val data: String,
    val created_at: String
)
```

Represents a document associated with a property or lease.

---

### `InviteDetailedProperty`

```kotlin
data class InviteDetailedProperty(
    val startDate: OffsetDateTime,
    val endDate: OffsetDateTime?,
    val tenantEmail: String
)
```

Represents detailed tenant invitation info with parsed date fields.

---

### `LeaseDetailedProperty`

```kotlin
data class LeaseDetailedProperty(
    val id: String,
    val startDate: OffsetDateTime,
    val endDate: OffsetDateTime?,
    val tenantEmail: String,
    val tenantName: String
)
```

Represents detailed lease info with parsed date fields.

---

### `DetailedProperty`

```kotlin
data class DetailedProperty(
    val id: String = "",
    val address: String = "",
    val status: PropertyStatus = PropertyStatus.unavailable,
    val appartementNumber: String? = "",
    val area: Int = 0,
    val rent: Int = 0,
    val deposit: Int = 0,
    val zipCode: String = "",
    val city: String = "",
    val country: String = "",
    val name: String = "",
    val picture: ImageBitmap? = null,
    val invite: InviteDetailedProperty? = null,
    val lease: LeaseDetailedProperty? = null,
) {
    fun toAddPropertyInput(): AddPropertyInput { ... }
}
```

Domain model representing detailed property info used throughout the app.
Contains a helper to convert back to `AddPropertyInput`.

---

### `ArchivePropertyInput`

```kotlin
data class ArchivePropertyInput(
    val archive: Boolean
)
```

Input data class to archive (or unarchive) a property by setting a boolean flag.

---

## RealPropertyCallerService Functions

* `suspend fun getPropertiesAsDetailedProperties(): Array<DetailedProperty>`
  Returns all properties as `DetailedProperty` array, handling ownership logic.

* `suspend fun getPropertyPicture(propertyId: String): String?`
  Retrieves the property picture data string or null if none.

* `suspend fun updatePropertyPicture(propertyId: String, propertyPicture: String): CreateOrUpdateResponse`
  Updates the property's picture with given data.

* `suspend fun addProperty(property: AddPropertyInput): CreateOrUpdateResponse`
  Adds a new property.

* `suspend fun archiveProperty(propertyId: String)`
  Archives a property.

* `suspend fun getPropertyWithDetails(propertyId: String? = null): DetailedProperty`
  Fetches a detailed property info depending on owner or tenant.

* `suspend fun updateProperty(property: AddPropertyInput, propertyId: String): CreateOrUpdateResponse`
  Updates an existing property.

* `suspend fun getPropertyDocuments(propertyId: String, leaseId: String): Array<Document>`
  Gets property documents with owner/tenant distinction.

* `suspend fun uploadDocument(propertyId: String, leaseId: String, document: DocumentInput): CreateOrUpdateResponse`
  Uploads a document for a property or lease with owner/tenant distinction.

