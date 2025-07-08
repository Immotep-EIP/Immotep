# InventoryReportManager

## Overview

`InventoryReportManager` is a class designed to manage API calls for inventory reports within a SwiftUI application. It provides methods to send reports, finalize inventories, and compare inventory reports.

## Data Models

### `SummarizeRequest`

| Field | Type | Description |
|-------|------|-------------|
| `id` | String | Identifier for the object the report is being generated for. |
| `pictures` | [String] | List of base64 encoded images to be analyzed. |
| `type` | String | Type of object, such as "furniture" or "room". |

### `SummarizeResponse`

| Field | Type | Description |
|-------|------|-------------|
| `state` | String | Estimated state of the object. |
| `note` | String | Comment or summary provided by the AI. |

### `RoomStateRequest`

| Field | Type | Description |
|-------|------|-------------|
| `id` | String | Identifier for the room. |
| `cleanliness` | String | Cleanliness level of the room. |
| `state` | String | State of the room. |
| `note` | String | Comment about the room. |
| `pictures` | [String] | List of base64 encoded images of the room. |
| `furnitures` | [FurnitureStateRequest] | List of furniture items in the room. |

### `FurnitureStateRequest`

| Field | Type | Description |
|-------|------|-------------|
| `id` | String | Identifier for the furniture item. |
| `cleanliness` | String | Cleanliness level of the furniture item. |
| `note` | String | Comment about the furniture item. |
| `pictures` | [String] | List of base64 encoded images of the furniture item. |
| `state` | String | State of the furniture item. |

## Functions

### `sendStuffReport()`

Sends a report for a specific object. Converts selected images to base64 and sends a POST request to the API to get a summary of the object's state.

### `compareStuffReport(oldReportId: String)`

Compares an object report with an old inventory report identified by `oldReportId`.

### `sendRoomReport()`

Sends a report for a specific room. Converts selected images to base64 and sends a POST request to the API to get a summary of the room's state.

### `compareRoomReport(oldReportId: String)`

Compares a room report with an old inventory report identified by `oldReportId`.

### `finalizeInventory()`

Finalizes the inventory by sending a complete report for all rooms and objects. Ensures all rooms and objects have been checked before sending the report.

## Usage

To use `InventoryReportManager`, you first need to initialize an instance with an `InventoryViewModel`.

```swift
let viewModel = InventoryViewModel()
let manager = InventoryReportManager(viewModel: viewModel)
Task {
    do {
        try await manager.sendStuffReport()
        try await manager.finalizeInventory()
    } catch {
        print("Error: \(error)")
    }
}
```