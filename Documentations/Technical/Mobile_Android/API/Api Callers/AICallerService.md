# AICallerService

## Overview

`AICallerService` is a subclass of [`ApiCallerService`](./ApiCallerService.md) that provides methods to interact with AI-powered endpoints for inventory report summarization and comparison.

These endpoints use image data and inventory location types to automatically assess the state, cleanliness, and overall condition of properties.

## Data Models

### `AiCallInput`

| Field     | Type                  | Description                                    |
|-----------|-----------------------|------------------------------------------------|
| `id`      | `String`              | Identifier for the AI analysis context.        |
| `pictures`| `Vector<String>`      | List of picture URLs to be analyzed.           |
| `type`    | `InventoryLocationsTypes` | Enum representing the location type (e.g., room). |

### `AiCallOutput`

| Field         | Type          | Description                                       |
|---------------|---------------|---------------------------------------------------|
| `cleanliness` | `Cleanliness?`| Estimated cleanliness rating from the AI.         |
| `note`        | `String?`     | Optional comment or summary provided by the AI.  |
| `state`       | `State?`      | Estimated state/condition of the inventory item. |

## Functions

### `suspend fun summarize(input: AiCallInput, propertyId: String, leaseId: String): AiCallOutput`

Calls the backend to get an AI-generated summary of the provided pictures related to a given lease and property.

- Requires a valid bearer token.
- Handles API exceptions via `changeRetrofitExceptionByApiCallerException`.

### `suspend fun compare(input: AiCallInput, propertyId: String, oldReportId: String, leaseId: String): AiCallOutput`

Calls the backend to compare a new inventory report against a previous one (`oldReportId`), using AI to highlight differences.

- Automatically handles token retrieval and exception conversion.

## Dependencies

- [`ApiService`](../ApiClient/ApiClientAndService.md) — The Retrofit interface used to communicate with the backend.
- [`ApiCallerService`](./ApiCallerService.kt) — Provides authentication and error handling logic.

