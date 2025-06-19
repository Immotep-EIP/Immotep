# DamageCallerService

## Overview

`DamageCallerService` is a subclass of [`ApiCallerService`](./ApiCallerService.md) that handles all interactions related to damage reports. It abstracts calls for both tenants and owners, automatically switching based on user role.

---

## Data Classes

### `DamageInput`

This class is used when submitting a new damage report.

```kotlin
data class DamageInput(
    val comment: String = "",
    val pictures: ArrayList<String> = ArrayList(),
    val priority: DamagePriority = DamagePriority.low,
    val room_id: String? = null
)
```

| Field    | Type                | Description                                          |
| -------- | ------------------- | ---------------------------------------------------- |
| comment  | `String`            | A textual description of the damage.                 |
| pictures | `ArrayList<String>` | A list of Base64-encoded image strings or URLs.      |
| priority | `DamagePriority`    | Priority level of the issue (e.g., low, high).       |
| room\_id | `String?`           | Optional room identifier associated with the damage. |

#### Method

```kotlin
fun toDamage(id: String, roomName: String, tenantName: String): Damage
```

* Converts the input into a `Damage` domain model using current timestamps and provided metadata.

---

### `DamageOutput`

This class models a damage object as returned by the API.

```kotlin
data class DamageOutput(
    val comment: String,
    val created_at: String,
    val fix_planned_at: String?,
    val fix_status: String,
    val fixed_at: String?,
    val id: String,
    val lease_id: String,
    val pictures: Array<String>?,
    val priority: String,
    val read: Boolean,
    val room_id: String,
    val room_name: String,
    val tenant_name: String,
    val updated_at: String
)
```

| Field            | Type             | Description                                  |
| ---------------- | ---------------- | -------------------------------------------- |
| comment          | `String`         | Damage description.                          |
| created\_at      | `String`         | Creation timestamp in ISO-8601 format.       |
| fix\_planned\_at | `String?`        | Optional planned date for fix.               |
| fix\_status      | `String`         | Status string (`PENDING`, `FIXED`, etc.).    |
| fixed\_at        | `String?`        | Optional actual fix timestamp.               |
| id               | `String`         | Unique identifier for the damage.            |
| lease\_id        | `String`         | Lease associated with this damage.           |
| pictures         | `Array<String>?` | Image URLs or Base64-encoded strings.        |
| priority         | `String`         | Priority as a string.                        |
| read             | `Boolean`        | Whether this damage report was acknowledged. |
| room\_id         | `String`         | ID of the damaged room.                      |
| room\_name       | `String`         | Name of the damaged room.                    |
| tenant\_name     | `String`         | Full name of the tenant who reported it.     |
| updated\_at      | `String`         | Last modification timestamp.                 |

#### Method

```kotlin
fun toDamage(): Damage
```

* Transforms this API data into a `Damage` domain model with `OffsetDateTime` and enums.

---

### `Damage`

This is the strongly-typed domain model used internally in the app logic.

```kotlin
data class Damage(
    val id: String,
    val comment: String,
    val createdAt: OffsetDateTime,
    val fixPlannedAt: OffsetDateTime?,
    val fixStatus: DamageStatus,
    val fixedAt: OffsetDateTime?,
    val leaseId: String,
    val pictures: Array<String>,
    val priority: DamagePriority,
    val read: Boolean,
    val roomId: String,
    val roomName: String,
    val tenantName: String,
    val updatedAt: OffsetDateTime
)
```

| Field        | Type              | Description                                     |
| ------------ | ----------------- | ----------------------------------------------- |
| id           | `String`          | Unique identifier for the damage.               |
| comment      | `String`          | Description of the problem.                     |
| createdAt    | `OffsetDateTime`  | When the damage was first logged.               |
| fixPlannedAt | `OffsetDateTime?` | Planned fix date, if any.                       |
| fixStatus    | `DamageStatus`    | Enum representing the fix status.               |
| fixedAt      | `OffsetDateTime?` | Actual date of fix.                             |
| leaseId      | `String`          | Associated lease ID.                            |
| pictures     | `Array<String>`   | Image URLs or Base64-encoded strings.           |
| priority     | `DamagePriority`  | Enum representing priority (LOW, MEDIUM, HIGH). |
| read         | `Boolean`         | Whether user has read/seen the damage.          |
| roomId       | `String`          | Room's ID.                                      |
| roomName     | `String`          | Room's name.                                    |
| tenantName   | `String`          | Reporting tenant's name.                        |
| updatedAt    | `OffsetDateTime`  | Last updated timestamp.                         |

---

## Functions

### `suspend fun getPropertyDamages(propertyId: String, leaseId: String): Array<Damage>`

Fetches damage reports for the specified lease and property:

* Calls tenant or owner-specific endpoint based on user role.
* Returns an array of domain `Damage` objects.

### `suspend fun addDamage(damageInput: DamageInput): CreateOrUpdateResponse`

Adds a new damage report under the current lease:

* Returns a `CreateOrUpdateResponse` containing the new damage ID.
* Handles exceptions and unauthorized cases automatically.

---

## Dependencies

* [`ApiService`](../apiClient/ApiService.md)
* [`ApiCallerService`](./ApiCallerService.md)
* `AuthService`
* `Retrofit`
* `OffsetDateTime`

