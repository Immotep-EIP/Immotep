# API Documentation for Dashboard and Furniture Management

## Overview

This document provides an overview of the API calls related to Dashboard and Furniture Management in the application. It includes details on data models and functions used to interact with these features.

---

## Dashboard Management

### Overview

The `OverviewViewModel` class is responsible for fetching and managing dashboard-related data for a property owner. This includes retrieving open damages, reminders, and property statistics.

### Data Models

#### DashboardResponse

The complete dashboard data structure returned by the API.

| Field | Type | Description |
|-------|------|-------------|
| reminders | [DashboardReminder] | List of reminders for the user. |
| properties | DashboardProperties | Summary of property statistics. |
| openDamages | DashboardOpenDamage | Summary of open damages. |

#### DashboardReminder

Represents a dashboard tip or reminder shown to the user.

| Field | Type | Description |
|-------|------|-------------|
| advice | String | Guidance or instructions. |
| id | String | Unique identifier of the reminder. |
| link | String | Actionable link (e.g., to a screen). |
| priority | String | Priority level of the reminder. |
| title | String | Title or topic of the reminder. |

#### DashboardProperties

Summary of property statistics for the dashboard.

| Field | Type | Description |
|-------|------|-------------|
| listRecentlyAdded | [Property]? | List of recently added properties. |
| nbrArchived | Int | Number of archived properties. |
| nbrAvailable | Int | Number of available properties. |
| nbrOccupied | Int | Number of occupied properties. |
| nbrPendingInvites | Int | Number of pending invites. |
| nbrTotal | Int | Total number of properties. |

#### DashboardOpenDamage

Summary of open damages for the dashboard.

| Field | Type | Description |
|-------|------|-------------|
| listToFix | [DamageResponse]? | List of damages to fix. |
| nbrHigh | Int | Number of high-priority damages. |
| nbrLow | Int | Number of low-priority damages. |
| nbrMedium | Int | Number of medium-priority damages. |
| nbrPlannedToFixThisWeek | Int | Number of damages planned to fix this week. |
| nbrTotal | Int | Total number of damages. |
| nbrUrgent | Int | Number of urgent damages. |

### Functions

#### `fetchDashboardData()`

Fetches dashboard data for the logged-in owner.
