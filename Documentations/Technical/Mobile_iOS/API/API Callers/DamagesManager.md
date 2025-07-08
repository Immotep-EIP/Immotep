# Damage Management API Documentation

## Overview

The `TenantPropertyViewModel` and `OwnerPropertyViewModel` classes handle interactions related to damage reports for tenants and owners, respectively. These classes provide methods to fetch, create, update, and manage damage reports associated with properties and leases.

## Data Models

### DamageResponse

Represents a damage report.

| Field | Type | Description |
|-------|------|-------------|
| id | String | Unique identifier for the damage report. |
| comment | String | Description of the damage. |
| createdAt | String | Creation timestamp of the damage report. |
| fixPlannedAt | String? | Planned date for fixing the damage. |
| fixStatus | String | Status of the fix (e.g., "PENDING", "FIXED"). |
| fixedAt | String? | Date when the damage was fixed. |
| leaseId | String | Identifier for the associated lease. |
| pictures | [String]? | List of image URLs or Base64-encoded strings related to the damage. |
| priority | String | Priority level of the damage. |
| read | Bool | Indicates if the damage report has been read. |
| roomId | String | Identifier for the room associated with the damage. |
| roomName | String | Name of the room associated with the damage. |
| tenantName | String | Name of the tenant who reported the damage. |
| updatedAt | String | Last update timestamp of the damage report. |

### DamageRequest

Used when submitting a new damage report.

| Field | Type | Description |
|-------|------|-------------|
| comment | String | Description of the damage. |
| pictures | [String] | List of Base64-encoded image strings or URLs. |
| priority | String | Priority level of the damage. |
| roomId | String? | Identifier for the room associated with the damage. |

## Functions

### `fetchTenantDamages(leaseId: String, fixed: Bool? = nil) -> [DamageResponse]`

Fetches damage reports for a tenant based on the lease ID and an optional filter for fixed status.

- **Parameters:**
  - `leaseId`: The ID of the lease.
  - `fixed`: Optional filter to fetch only fixed or unfixed damages.

- **Returns:** An array of `DamageResponse` objects.

### `createDamage(propertyId: String, leaseId: String, damage: DamageRequest, token: String) -> String`

Creates a new damage report for a property and lease.

- **Parameters:**
  - `propertyId`: The ID of the property.
  - `leaseId`: The ID of the lease.
  - `damage`: The damage report details.
  - `token`: Authentication token.

- **Returns:** The ID of the created damage report.

### `fetchDamageByID(damageId: String, token: String) -> DamageResponse`

Fetches a specific damage report by its ID.

- **Parameters:**
  - `damageId`: The ID of the damage report.
  - `token`: Authentication token.

- **Returns:** A `DamageResponse` object.

### `fixDamage(damageId: String, token: String)`

Marks a damage report as fixed.

- **Parameters:**
  - `damageId`: The ID of the damage report.
  - `token`: Authentication token.

### `fetchPropertyDamages(propertyId: String, fixed: Bool? = nil) -> [DamageResponse]`

Fetches damage reports for a property, with an optional filter for fixed status.

- **Parameters:**
  - `propertyId`: The ID of the property.
  - `fixed`: Optional filter to fetch only fixed or unfixed damages.

- **Returns:** An array of `DamageResponse` objects.

### `updateDamageStatus(propertyId: String, damageId: String, fixPlannedAt: String?, read: Bool, token: String)`

Updates the status of a damage report.

- **Parameters:**
  - `propertyId`: The ID of the property.
  - `damageId`: The ID of the damage report.
  - `fixPlannedAt`: Optional planned fix date.
  - `read`: Whether the damage report has been read.
  - `token`: Authentication token.

### `fixDamage(propertyId: String, damageId: String, token: String)`

Marks a damage report as fixed for a property.

- **Parameters:**
  - `propertyId`: The ID of the property.
  - `damageId`: The ID of the damage report.
  - `token`: Authentication token.

