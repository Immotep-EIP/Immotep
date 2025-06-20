# RetrofitClient (API Client)

## Overview

The `RetrofitClient` object configures the base networking client for all API calls in the application using Retrofit and OkHttp.

## Configuration

- **Base URL**:  
  Currently set to `https://dev.space.keyz-app.fr/`.

- **Timeouts**:  
  All network timeouts (connect, read, write) are set to 600 seconds to support longer operations.

- **Logging**:  
  HTTP requests and responses are logged in full using `HttpLoggingInterceptor.Level.BODY`.

- **Serialization**:  
  JSON conversion is handled via `GsonConverterFactory`.

## Components

- `RetrofitClient.retrofit`: Lazy-initialized `Retrofit` instance used to create the API interface.
- `ApiClient.apiService`: Lazily provides a concrete implementation of the `ApiService` interface.

## Example Usage

```kotlin
val apiService = ApiClient.apiService
```

This `apiService` can then be injected into repositories or view models to perform API calls.


# ApiService (API Interface)

## Overview

The `ApiService` interface defines the contract between the client application and the Keyz backend API. It uses Retrofit annotations to describe REST endpoints, headers, and data serialization.

## General Info

- **Base Endpoint**: `/api/v1`
- **Authentication**: Most endpoints require a Bearer token via the `Authorization` header.
- **Serialization**: Uses `@Body`, `@Query`, `@Field`, and `@Path` annotations to bind parameters to HTTP requests.
- **Error Handling**: API errors should be parsed using appropriate HTTP status codes and payloads.

## Areas Covered

This interface includes methods for interacting with the following backend resources:

- **Authentication**
  - Login
  - Refresh token
  - Registration

- **User Profile**
  - Get/update profile information

- **Property Management**
  - Create, update, archive properties
  - Upload and retrieve property pictures and documents

- **Tenant Interactions**
  - Invitations
  - Property and document retrieval per lease

- **Room and Furniture Management**
  - Add/Archive rooms and furnitures
  - Get data for both owner and tenant views

- **Inventory Reports**
  - Create and retrieve inventory reports
  - AI-powered summary and comparison

- **Damage Reports**
  - Submit and retrieve damage information

- **Dashboard**
  - Owner dashboard summary endpoint

## Usage

Each method is marked `suspend` and must be called from a coroutine scope (e.g., using `viewModelScope.launch {}`).

## Example

```kotlin
val response = apiService.login(username = "user@example.com", password = "securepass")
```

## See Also

* [`AuthService`](../AuthService.md) – Service layer using this interface to implement application-level authentication logic.
