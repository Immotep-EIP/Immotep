# ProfileCallerService

## Overview

`ProfileCallerService` extends [`ApiCallerService`](./ApiCallerService.md) to manage API calls related to user profile retrieval and updates.

---

## Data Classes

### `ProfileResponse`

```kotlin
data class ProfileResponse(
    val id: String,
    val email: String,
    val firstname: String,
    val lastname: String,
    val role: String,
    val created_at: String,
    val updated_at: String,
)
```

Represents the user profile data received from the API.

| Field       | Type     | Description                      |
| ----------- | -------- | -------------------------------- |
| id          | `String` | Unique user identifier           |
| email       | `String` | User email address               |
| firstname   | `String` | User first name                  |
| lastname    | `String` | User last name                   |
| role        | `String` | User role (e.g., owner, tenant)  |
| created\_at | `String` | Timestamp of profile creation    |
| updated\_at | `String` | Timestamp of last profile update |

---

### `ProfileUpdateInput`

```kotlin
data class ProfileUpdateInput(
    val email: String,
    val firstname: String,
    val lastname: String
)
```

Represents the input data for updating a user profile.

| Field     | Type     | Description       |
| --------- | -------- | ----------------- |
| email     | `String` | New email address |
| firstname | `String` | New first name    |
| lastname  | `String` | New last name     |

---

## Functions

### `suspend fun getProfile(): ProfileResponse`

Fetches the current user's profile data from the API.

---

### `suspend fun updateProfile(profileUpdateInput: ProfileUpdateInput)`

Sends updated profile data to the API to modify the user's profile.

---

## Dependencies

* [`ApiService`](../apiClient/ApiService.md)
* [`ApiCallerService`](./ApiCallerService.md)

