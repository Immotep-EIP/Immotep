# API Overview

This document provides a high-level overview of the core API-related services and components used in the Keyz iOS application. It covers authentication, API caller managers, and their roles in managing interactions with the Keyz backend API.

---

## Contents

* [Authentication Service (`AuthService`)](#authentication-service-authservice)
* [API Service (`ApiService`)](#api-service-apiservice)
* [API Caller Managers](#api-caller-managers)

---

## Authentication Service (`AuthService`)

`AuthService` is an actor responsible for handling user authentication and token management. It interacts with the backend API to perform login, token refresh, and authorized requests, ensuring secure storage of access and refresh tokens.

* **Features:**
  * User login with email and password
  * Token refresh using refresh tokens
  * Authorized API requests with Bearer token
  * Special character encoding for safe token usage
* **Key Structures:**
  * `TokenResponse`: Contains access token, refresh token, expiration, and user properties.
  * `UserProperties`: Includes user ID and role.
* **Token Management:** Uses `TokenStorage` to store access tokens, refresh tokens, and expiration times, with support for persistent storage based on `keepMeSignedIn`.
* **Error Handling:** Throws `NSError` with descriptive messages for HTTP errors, invalid responses, or missing tokens.

For detailed technical documentation, see [AuthService Docs](APIService.md).

---

## API Service (`ApiService`)

`ApiService` is an actor implementing the `ApiServiceProtocol` to handle user registration with the backend API. It serves as the entry point for non-authentication-related API calls.

* **Features:**
  * User registration with email, password, first name, and last name
  * JSON-based request construction and response parsing
* **Key Structures:**
  * `RegisterModel`: Input structure for registration.
  * `IdResponse`: Response containing the created user’s ID.
  * `ErrorResponse`: Error details for failed requests.
* **Error Handling:** Throws specific `NSError` instances for HTTP 400 (empty fields) and 409 (email already exists) errors, with fallback for other status codes.

For detailed technical documentation, see [ApiService Docs](APIService.md).

---

## API Caller Managers

API caller managers are specialized classes that handle domain-specific API operations, interacting with the backend via `URLSession` and leveraging `AuthService` for authentication. Each manager is responsible for a specific resource area, ensuring clean separation of concerns.

* **Core Responsibilities:**
  * Construct and execute API requests with proper authentication
  * Handle JSON encoding/decoding and error responses
  * Update the application state via `InventoryViewModel` or similar view models
  * Refresh tokens when necessary using `AuthService`
* **Managers Overview:**
  * **`DamagesManager`:** Manages damage report submission and retrieval.
  * **`DashboardManager`:** Fetches owner dashboard summary data.
  * **`FurnituresManager`:** Handles furniture-related CRUD operations for rooms.
  * **`InventoryReportManager`:** Manages creation and retrieval of inventory reports.
  * **`InviteTenantManager`:** Handles tenant invitations for properties.
  * **`PropertiesManager`:** Manages property creation, updates, archiving, and document/picture uploads.
  * **`RoomsManager`:** Manages room-related operations, including fetching, adding, archiving, and state management.
* **Integration:**
  * Managers rely on `AuthService` for token management and authorized requests.
  * Each manager updates the application state through a shared view model (e.g., `InventoryViewModel`).
  * Error handling is consistent, with descriptive messages propagated to the UI via view models.

For detailed technical documentation, see the respective manager docs:
- [DamagesManager](API Callers/DamagesManager.md)
- [DashboardManager](API Callers/DashboardManager.md)
- [FurnituresManager](API Callers/FurnituresManager.md)
- [InventoryReportManager](API Callers/InventoryReportManager.md)
- [InviteTenantManager](API Callers/InviteTenantManager.md)
- [PropertiesManager](API Callers/PropertiesManager.md)
- [RoomsManager](API Callers/RoomsManager.md)

---

## Additional Notes

* All API interactions use `URLSession` for networking, with no dependency on third-party libraries like Retrofit (unlike the Android counterpart).
* Authentication tokens are managed centrally by `AuthService` and stored via `TokenStorage`.
* Managers are designed to work with `MainActor`-bound view models for UI updates, ensuring thread safety.
* The iOS implementation does not include a mock API service, unlike the Android `MockedApiService`. Testing relies on real API calls or manual mocks.
* Error handling is unified across managers, with errors surfaced to the UI via view model properties (e.g., `errorMessage`).
