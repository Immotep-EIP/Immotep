# FurnitureCallerService

## Overview

`FurnitureCallerService` is a subclass of [`ApiCallerService`](./ApiCallerService.md) used to handle API calls related to furniture data within a specific room and property.

---

## Data Classes

### `FurnitureOutput`

```kotlin
data class FurnitureOutput(
    val id: String,
    val property_id: String,
    val room_id: String,
    val name: String,
    val quantity: Int
)
```

Represents a piece of furniture returned by the API.

| Field        | Type     | Description                            |
| ------------ | -------- | -------------------------------------- |
| id           | `String` | Unique ID of the furniture             |
| property\_id | `String` | ID of the related property             |
| room\_id     | `String` | ID of the room containing this item    |
| name         | `String` | Name of the furniture item             |
| quantity     | `Int`    | Quantity of this furniture in the room |

#### Method

```kotlin
fun toRoomDetail(): RoomDetail
```

Converts the furniture to a `RoomDetail` placeholder used in inventory reports. Default values are used for untracked fields.

---

### `FurnitureInput`

```kotlin
data class FurnitureInput(
    val name: String,
    val quantity: Int
)
```

Represents the input payload for adding a new piece of furniture.

| Field    | Type     | Description           |
| -------- | -------- | --------------------- |
| name     | `String` | Name of the furniture |
| quantity | `Int`    | Quantity to be added  |

---

## Functions

### `suspend fun getFurnituresByRoomId(propertyId: String, roomId: String): Array<FurnitureOutput>`

Fetches all furnitures in a given room belonging to a specific property.

* Requires a valid bearer token.
* Handles exceptions and logs out if unauthorized.

---

### `suspend fun addFurniture(propertyId: String, roomId: String, furniture: FurnitureInput): CreateOrUpdateResponse`

Adds a furniture entry in a specific room.

* Sends `FurnitureInput` as the request body.
* Returns the created or updated resource ID.

---

## Dependencies

* [`ApiService`](../ApiClient/ApiClientAndService.md)
* [`ApiCallerService`](./ApiCallerService.md)
* `RoomDetail`, `State`, `Cleanliness` from the `inventory` package
