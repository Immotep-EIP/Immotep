# InventoryCallerService

## Overview

`InventoryCallerService` is a subclass of [`ApiCallerService`](./ApiCallerService.md). It provides methods to interact with inventory report-related endpoints, such as creating and retrieving inventory reports for a specific property.

---

## Data Classes

### `InventoryReportInput`

```kotlin
data class InventoryReportInput(
    val type: String,
    val rooms: Vector<InventoryReportRoom>
)
```

Represents the input payload required to create an inventory report.

| Field | Type                          | Description                                  |
| ----- | ----------------------------- | -------------------------------------------- |
| type  | `String`                      | Type of the report (e.g. entry, exit, visit) |
| rooms | `Vector<InventoryReportRoom>` | List of rooms to be included in the report   |

---

### `CreatedInventoryReport`

```kotlin
data class CreatedInventoryReport(
    val date: String,
    val errors : Array<String>,
    val id: String,
    val lease_id: String,
    val pdf_data: String,
    val pdf_name: String,
    val property_id: String,
    val type: String
)
```

Represents the response returned after a successful inventory report creation.

| Field        | Type            | Description                           |
| ------------ | --------------- | ------------------------------------- |
| date         | `String`        | Date of the report creation           |
| errors       | `Array<String>` | List of errors that may have occurred |
| id           | `String`        | ID of the created report              |
| lease\_id    | `String`        | Related lease ID                      |
| pdf\_data    | `String`        | Base64 encoded PDF data               |
| pdf\_name    | `String`        | Name of the PDF file                  |
| property\_id | `String`        | ID of the related property            |
| type         | `String`        | Type of the inventory report          |

---

## Functions

### `suspend fun createInventoryReport(propertyId: String, inventoryReportInput: InventoryReportInput, leaseId: String): CreatedInventoryReport`

Creates a new inventory report.

* Wraps the `apiService.inventoryReport(...)` call.
* If an exception occurs, it logs and throws a generic `ApiCallerServiceException`.

---

### `suspend fun getAllInventoryReports(propertyId: String): Array<InventoryReportOutput>`

Fetches all inventory reports associated with a property.

---

### `suspend fun getLastInventoryReport(propertyId: String): InventoryReportOutput`

Retrieves the most recent inventory report for the given property.

---

### `suspend fun getInventoryReportById(propertyId: String, reportId: String): InventoryReportOutput`

Fetches a specific inventory report by its ID.

---

## Dependencies

* `InventoryReportRoom`, `InventoryReportOutput` from the `inventory` package
* [`ApiService`](../ApiClient/ApiClientAndService.md)
* [`ApiCallerService`](./ApiCallerService.md)

