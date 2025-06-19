# RoomCallerService

## Overview

`RoomCallerService` extends the base [`ApiCallerService`](./ApiCallerService.md) and manages API calls related to rooms within a property. It supports fetching rooms (with or without furniture details), adding new rooms, and archiving rooms, with role-based access (owner vs tenant).

---

## Enum Classes

### `RoomType`

```kotlin
enum class RoomType {
    dressing,
    laundryroom,
    bedroom,
    playroom,
    bathroom,
    toilet,
    livingroom,
    diningroom,
    kitchen,
    hallway,
    balcony,
    cellar,
    garage,
    storage,
    office,
    other
}
```

Enumerates all possible types of rooms.

---

## Data Classes

### `AddRoomInput`

```kotlin
data class AddRoomInput(
    val name: String,
    val type: RoomType
)
```

Input data class for adding a new room, specifying its name and type.

---

### `RoomOutput`

```kotlin
data class RoomOutput(
    val id: String,
    val name: String,
    val property_id: String,
    val type: String,
    val archived: Boolean
) {
    fun toRoom(details: Array<RoomDetail>?): Room { ... }
}
```

Represents the API output for a room.
Includes a helper function to convert itself to a domain model `Room`, optionally with `RoomDetail` array.

---

## RoomCallerService Functions

* `suspend fun getAllRooms(propertyId: String): Array<RoomOutput>`
  Retrieves all rooms for a given property, switching endpoint based on whether the current user is owner or tenant.

* `suspend fun getAllRoomsWithFurniture(propertyId: String, onErrorRoomFurniture: (String) -> Unit): Array<Room>`
  Fetches all rooms including their furniture details. Calls `FurnitureCallerService` internally to retrieve furniture by room.
  Errors during furniture retrieval for individual rooms call `onErrorRoomFurniture` with the room name.

* `suspend fun addRoom(propertyId: String, room: AddRoomInput): CreateOrUpdateResponse`
  Adds a new room to the specified property.

* `suspend fun archiveRoom(propertyId: String, roomId: String)`
  Archives a room by ID.

---

## Dependencies

* Uses `FurnitureCallerService` internally to get furniture details by room.

