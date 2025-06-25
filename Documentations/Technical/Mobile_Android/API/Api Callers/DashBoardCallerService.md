# DashBoardCallerService

## Overview

`DashBoardCallerService` is a subclass of [`ApiCallerService`](./ApiCallerService.md) that retrieves all dashboard-related data for the logged-in owner, including open damages, reminders, and property statistics.

---

## Data Classes

### `DashBoardReminder`

```kotlin
data class DashBoardReminder(
    val advice: String,
    val id: String,
    val link: String,
    val priority: String,
    val title: String
)
```

Represents a dashboard tip or reminder shown to the user.

| Field    | Type     | Description                          |
| -------- | -------- | ------------------------------------ |
| advice   | `String` | Guidance or instructions.            |
| id       | `String` | Unique identifier of the reminder.   |
| link     | `String` | Actionable link (e.g., to a screen). |
| priority | `String` | Priority level of the reminder.      |
| title    | `String` | Title or topic of the reminder.      |

---

### `DashBoardPropertyOutput`

```kotlin
data class DashBoardPropertyOutput(
    val id: String,
    val apartment_number: String?,
    val archived: Boolean,
    val owner_id: String,
    val name: String,
    val address: String,
    val city: String,
    val postal_code: String,
    val country: String,
    val area_sqm: Double,
    val rental_price_per_month: Int,
    val deposit_price: Int,
    val created_at: String,
)
```

API response structure for a property in the dashboard.

#### Method

```kotlin
fun toDetailedProperty(): DetailedProperty
```

* Converts this to an internal `DetailedProperty` object.

---

### `DashBoardPropertiesOutput`

```kotlin
data class DashBoardPropertiesOutput(
    val list_recently_added: Array<DashBoardPropertyOutput>?,
    val nbr_archived: Int,
    val nbr_available: Int,
    val nbr_occupied: Int,
    val nbr_pending_invites: Int,
    val nbr_total: Int
)
```

Summary of property statistics for the dashboard.

#### Method

```kotlin
fun toDashBoardProperties(): DashBoardProperties
```

---

### `DashBoardProperties`

```kotlin
data class DashBoardProperties(
    val listRecentlyAdded: Array<DetailedProperty> = arrayOf(),
    val nbrArchived: Int = 0,
    val nbrAvailable: Int = 0,
    val nbrOccupied: Int = 0,
    val nbrPendingInvites: Int = 0,
    val nbrTotal: Int = 0
)
```

Typed internal model of property statistics used in UI logic.

---

### `DashBoardOpenDamageOutput`

```kotlin
data class DashBoardOpenDamageOutput(
    val list_to_fix: Array<DamageOutput>?,
    val nbr_high: Int,
    val nbr_low: Int,
    val nbr_medium: Int,
    val nbr_planned_to_fix_this_week: Int,
    val nbr_total: Int,
    val nbr_urgent: Int
)
```

Raw response from the API for open damages.

#### Method

```kotlin
fun toDashBoardOpenDamage(): DashBoardOpenDamage
```

---

### `DashBoardOpenDamage`

```kotlin
data class DashBoardOpenDamage(
    val listToFix: Array<Damage> = arrayOf(),
    val nbrHigh: Int = 0,
    val nbrLow: Int = 0,
    val nbrMedium: Int = 0,
    val nbrPlannedToFixThisWeek: Int = 0,
    val nbrTotal: Int = 0,
    val nbrUrgent: Int = 0
)
```

Internal model representing grouped damage statistics and pending actions.

---

### `GetDashBoardOutput`

```kotlin
data class GetDashBoardOutput(
    val open_damages: DashBoardOpenDamageOutput,
    val properties: DashBoardPropertiesOutput,
    val reminders: Array<DashBoardReminder>
)
```

The complete dashboard data structure returned by the API.

#### Method

```kotlin
fun toDashBoard(): DashBoard
```

---

### `DashBoard`

```kotlin
data class DashBoard(
    val reminders: Array<DashBoardReminder> = arrayOf(),
    val openDamages: DashBoardOpenDamage = DashBoardOpenDamage(),
    val properties: DashBoardProperties = DashBoardProperties()
)
```

Final app-ready model combining open damages, property stats, and reminders.

---

## Functions

### `suspend fun getDashBoard(): DashBoard`

Fetches and transforms the dashboard content into a `DashBoard` model.

* Handles localization with `Locale.getDefault().language`.
* Uses built-in exception management from `ApiCallerService`.

---

## Dependencies

* [`ApiCallerService`](./ApiCallerService.md)
* [`DetailedProperty`](./RealPropertyCallerService.md)
* [`Damage, DamageOutput, DamageStatus`](./DamageCallerService.md)

