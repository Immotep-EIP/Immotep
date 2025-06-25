# InviteTenantCallerService

## Overview

`InviteTenantCallerService` extends [`ApiCallerService`](./ApiCallerService.md) to manage tenant invitation-related API calls, including sending and canceling tenant invites for a property.

---

## Data Classes

### `InviteInput`

```kotlin
data class InviteInput(
    val tenant_email: String,
    val start_date: String,
    val end_date: String,
)
```

Represents the input data required to send a tenant invitation.

| Field         | Type     | Description                    |
| ------------- | -------- | ------------------------------ |
| tenant\_email | `String` | Email address of the tenant    |
| start\_date   | `String` | Start date of the lease/invite |
| end\_date     | `String` | End date of the lease/invite   |

---

### `InviteOutput`

```kotlin
data class InviteOutput(
    val id: String,
    val property_id: String,
    val tenant_email: String,
    val start_date: String,
    val end_date: String,
    val created_at: String
)
```

Represents the data returned from the API when querying an invite.

| Field         | Type     | Description                                   |
| ------------- | -------- | --------------------------------------------- |
| id            | `String` | Unique identifier of the invite               |
| property\_id  | `String` | ID of the property associated with the invite |
| tenant\_email | `String` | Email of the invited tenant                   |
| start\_date   | `String` | Start date of the invite/lease                |
| end\_date     | `String` | End date of the invite/lease                  |
| created\_at   | `String` | Timestamp when the invite was created         |

---

## Functions

### `suspend fun invite(propertyId: String, inviteInput: InviteInput): CreateOrUpdateResponse`

Sends an invitation to a tenant for a specific property.

---

### `suspend fun cancelInvite(propertyId: String)`

Cancels a previously sent tenant invitation for the given property.

---

## Dependencies

* [`ApiService`](../ApiClient/ApiClientAndService.md)
* [`ApiCallerService`](./ApiCallerService.md)
* `CreateOrUpdateResponse` (generic response for create or update API calls)

