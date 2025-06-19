# ApiCallerService (Base Class)

## Overview

`ApiCallerService` is a sealed abstract class that provides the base for all API caller services in the application. These services are responsible for executing authenticated network operations and managing error handling consistently across the app.

## Purpose

The base class abstracts common behavior needed by API caller services, including:

- Fetching and refreshing the bearer token using [`AuthService`](#authservice).
- Handling unauthorized errors (`401`) and triggering logout via navigation if needed.
- Catching and converting Retrofit exceptions into a domain-specific `ApiCallerServiceException`.
- Providing a consistent way to determine user roles (e.g., owner vs. tenant).

## Exception Handling

```kotlin
class ApiCallerServiceException(message: String) : Exception(message) {
    fun getCode(): Int = message?.toIntOrNull() ?: 400
}
```

* All API exceptions are wrapped into `ApiCallerServiceException`, with the HTTP status code extracted from the original exception.

## Key Methods

### `getBearerToken(): String`

Retrieves and returns a fresh access token. If token retrieval fails, triggers logout and rethrows the error.

### `changeRetrofitExceptionByApiCallerException(...)`

Generic function that executes a given suspend function and wraps any `HttpException` or general exceptions into `ApiCallerServiceException`. Handles auto-logout if the error is 401 and `logoutOnUnauthorized` is true.

### `isOwner(): Boolean`

Returns whether the authenticated user is an owner by delegating to `AuthService`.

## Usage

This class should be extended by specific domain services (e.g., `ProfileApiCallerService`, `PropertyApiCallerService`, etc.) that use `apiService` to perform business-specific operations. These subclasses inherit the token and error management behavior from `ApiCallerService`.

## Example

```kotlin
class ProfileApiCallerService(
    apiService: ApiService,
    navController: NavController
) : ApiCallerService(apiService, navController) {
    suspend fun getProfile(): ProfileResponse = changeRetrofitExceptionByApiCallerException {
        apiService.getProfile("Bearer ${getBearerToken()}")
    }
}
```
