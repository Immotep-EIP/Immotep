# RoomManager

## Overview

`RoomManager` is a class designed to handle API interactions related to rooms within a property in the Keyz iOS application. It manages fetching rooms, adding new rooms, archiving rooms, and handling room selection and completion states. It interacts with the `InventoryViewModel` to update the UI and maintain local state consistency. The class supports role-based API endpoints (owner-focused in this implementation).

---

## Enum Definitions

### `RoomType`

```swift
enum RoomType: String, CaseIterable {
    case dressing
    case laundryroom
    case bedroom
    case playroom
    case bathroom
    case toilet
    case livingroom
    case diningroom
    case kitchen
    case hallway
    case balcony
    case cellar
    case garage
    case storage
    case office
    case other
}
```

Enumerates all possible types of rooms that can be associated with a property.

---

## Data Structures

### `RoomResponse`

```swift
struct RoomResponse: Codable {
    let id: String
    let name: String
    let property_id: String
    let type: String
    let archived: Bool
}
```

Represents the API response for a room, including its ID, name, associated property ID, type, and archival status.

### `AddRoomInput`

```swift
struct AddRoomInput: Codable {
    let name: String
    let type: String
}
```

Input structure for adding a new room, specifying its name and type.

### `IdResponse`

```swift
struct IdResponse: Codable {
    let id: String
}
```

Represents the API response when creating a new room, returning the created room's ID.

### `ErrorResponse`

```swift
struct ErrorResponse: Codable {
    let error: String
}
```

Represents the API error response, containing an error message.

---

## RoomManager Methods

### `init(viewModel: InventoryViewModel)`

Initializes the `RoomManager` with a reference to an `InventoryViewModel` for updating UI state and accessing authentication tokens.

- **Parameter**: `viewModel` - The `InventoryViewModel` instance for managing UI and state.

### `func fetchRooms() async`

Fetches all rooms for the current property from the API (`/owner/properties/{propertyId}/rooms/`). Updates the `viewModel`'s `property.rooms` and `localRooms` with the fetched data, preserving local state (e.g., `checked`, `inventory`, `images`, `status`, `comment`) for existing rooms.

- **Behavior**:
  - Validates the property ID and constructs the API URL.
  - Retrieves an authentication token from the view model.
  - Sends a GET request with Bearer token authorization.
  - Decodes the response into an array of `RoomResponse`.
  - Maps the response to `PropertyRooms` and updates `localRooms` with merged state.
  - Sets `viewModel.errorMessage` on failure.

### `func addRoom(name: String, type: String) async throws`

Adds a new room to the specified property using the API (`/owner/properties/{propertyId}/rooms/`).

- **Parameters**:
  - `name`: The name of the room.
  - `type`: The type of the room (must be a valid `RoomType` raw value).
- **Behavior**:
  - Validates the property ID, URL, token, and room type.
  - Sends a POST request with a JSON body containing the room details.
  - Refreshes the room list by calling `fetchRooms()` on success.
  - Throws an `NSError` with a descriptive message on failure.

### `func deleteRoom(_ room: LocalRoom) async`

Archives a room by ID using the API (`/owner/properties/{propertyId}/rooms/{roomId}/archive/`).

- **Parameter**: `room` - The `LocalRoom` to archive.
- **Behavior**:
  - Validates the property ID, URL, and token.
  - Sends a PUT request with a JSON body setting `archive: true`.
  - Refreshes the room list and removes the archived room from `localRooms` on success.
  - Sets `viewModel.errorMessage` on failure.

### `func selectRoom(_ room: LocalRoom)`

Selects a room, updating the `viewModel`'s `selectedRoom`, `selectedInventory`, and `roomStatus`.

- **Parameter**: `room` - The `LocalRoom` to select.
- **Behavior**:
  - Updates the `viewModel` with the selected room's details.
  - Sets the inventory and status based on the room's local state.

### `func isRoomCompleted(_ room: LocalRoom) -> Bool`

Checks if a room is completed by verifying that all its inventory items are checked.

- **Parameter**: `room` - The `LocalRoom` to check.
- **Returns**: `true` if all inventory items are checked, `false` otherwise.

### `func areAllRoomsCompleted() -> Bool`

Checks if all rooms in the `viewModel`'s `localRooms` are completed.

- **Returns**: `true` if all rooms are checked, `false` otherwise or if the view model is unavailable.

### `func markRoomAsChecked(_ room: LocalRoom) async`

Marks a room as checked in the `viewModel`'s `localRooms`.

- **Parameter**: `room` - The `LocalRoom` to mark as checked.
- **Behavior**: Updates the `checked` property of the specified room in `localRooms`.

---

## Dependencies

- **InventoryViewModel**: Provides property details, authentication tokens, and state management for rooms and inventory.
- **APIConfig**: Supplies the base URL for API requests.
- **URLSession**: Used for making HTTP requests to the API.
- **JSONDecoder/JSONSerialization**: Handles JSON encoding and decoding for API requests and responses.

---

## Notes

- All API calls use Bearer token authentication, retrieved via `viewModel.getToken()`.
- The class assumes an owner role, using `/owner/` endpoints. Tenant-specific endpoints are not implemented in this version.
- Error handling updates `viewModel.errorMessage` for UI feedback or throws `NSError` for `addRoom`.
- The `fetchRooms` and `deleteRoom` methods refresh the room list to ensure state consistency.
- Furniture details are not fetched in this implementation, unlike the Android counterpart which uses `FurnitureCallerService`.
