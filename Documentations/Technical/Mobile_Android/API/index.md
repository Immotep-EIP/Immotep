# API Overview

This document provides a high-level overview of the core API-related services and components used in the application. It covers authentication, API caller service hierarchy, property and room management, and mock API implementations for testing purposes.

---

## Contents

* [Authentication Service (`AuthService`)](#authentication-service-authservice)
* [API Caller Service (`ApiCallerService`) and Subclasses](#api-caller-service-apicallerservice-and-subclasses)
* [Mock API Service (`MockedApiService`)](#mock-api-service-mockedapiservice)

---

## Authentication Service (`AuthService`)

`AuthService` handles user authentication, token management, and user registration by interacting with the backend API via `ApiService`. It manages secure storage of access and refresh tokens with Android's DataStore and refreshes tokens when expired.

* **Features:**

  * User login and token storage
  * Access token refreshing
  * User registration
  * Role detection (owner or tenant)
  * Logout and token deletion
* **Key classes:**
  `LoginResponse`, `RegistrationInput`, `RegistrationResponse`
* **Token Management:** Uses DataStore with keys for `ACCESS_TOKEN`, `REFRESH_TOKEN`, `EXPIRES_IN`, and `IS_OWNER`.
* **Error Handling:** Exceptions include detailed HTTP error codes for API failures.

For detailed technical documentation, see the [AuthService Docs](#).

---

## API Caller Service (`ApiCallerService`) and Subclasses

`ApiCallerService` is the **base abstract service** that provides common functionality and error handling for all API caller subclasses responsible for specific backend domains.

* **Core Responsibilities:**

  * Centralized management of bearer token retrieval and refresh via embedded `AuthService`
  * Unified Retrofit exception handling, wrapping errors in `ApiCallerServiceException`
  * Automated logout flow on unauthorized (401) errors, triggering navigation to login screen
  * Determining user role (owner or tenant) for conditional logic in subclasses

* **Exception Class:**
  `ApiCallerServiceException` encapsulates HTTP error codes and general exceptions.

* **Subclasses Overview:**
  Specialized service classes extend `ApiCallerService` to handle domain-specific API operations, such as:

  * **`RoomCallerService`:** Manages room-related API calls including fetching rooms, adding, archiving, and retrieving furniture details.
  * **`FurnitureCallerService`:** Handles furniture CRUD operations for rooms.
  * **`InventoryCallerService`:** Manages inventory report creation and retrieval.
  * **(Others):** Additional callers for property, tenant, damage, and dashboard API endpoints.

* **Integration:**
  Each subclass uses `changeRetrofitExceptionByApiCallerException` to safely execute API calls and handle errors uniformly.

This layered architecture promotes code reuse, consistent error handling, and clean separation of concerns between different API domains.

For detailed technical documentation, see the [ApiCallerService Docs](#) and the respective subclasses.

---

## Mock API Service (`MockedApiService`)

`MockedApiService` simulates the behavior of `ApiService` for testing and development without backend dependency. It provides predefined responses for authentication, property, room, furniture, and inventory-related API calls.

* **Features:**

  * Simulate login, registration, token refresh
  * Mock property and room retrieval and modifications
  * Simulate furniture and inventory report APIs
  * Supports tenant invitation and profile updates
* **Purpose:** Enables unit testing and development in isolation.
* **Implementation:** Returns fixed data objects for each API method, mimicking real API responses.

For detailed technical documentation, see the [MockedApiService Docs](#).

---

## Additional Notes

* All API calls depend on Retrofit interfaces implemented by `ApiService`.
* Authentication token lifecycle is managed centrally in `AuthService` and exposed through `ApiCallerService`.
* Navigation handling is integrated in `ApiCallerService` to ensure secure session management.
* Common error decoding and mapping is centralized using `decodeRetroFitMessagesToHttpCodes`.

