# ApiService and AuthService

## Overview

The `ApiService` and `AuthService` actors handle API interactions for the Keyz iOS application. `ApiService` manages user registration, while `AuthService` handles authentication-related tasks such as login, token refresh, and authorized requests. Both use `URLSession` for networking and are designed as actors to ensure thread-safe operations.

---

## Configuration

- **Base URL**: Configured via `APIConfig.baseURL` (e.g., `https://dev.space.keyz-app.fr/`).
- **Content Types**:
  - `application/json` for JSON payloads (e.g., registration).
  - `application/x-www-form-urlencoded` for token requests.
- **Error Handling**: Throws `NSError` with localized descriptions for invalid responses, status codes, or parsing errors.
- **Authentication**: Uses Bearer tokens for authorized requests, managed via `TokenStorage`.

---

## ApiService

### Overview

`ApiService` is an actor implementing the `ApiServiceProtocol` to handle user registration with the Keyz backend API.

### Protocol

```swift
protocol ApiServiceProtocol {
    func registerUser(with model: RegisterModel) async throws -> String
}
```

Defines the contract for user registration.

### Methods

#### `func registerUser(with model: RegisterModel) async throws -> String`

Registers a new user with the API (`/auth/register/`).

- **Parameter**: `model` - A `RegisterModel` containing email, password, first name, and last name.
- **Returns**: A success message ("Registration successful!") on completion.
- **Throws**:
  - `NSError` for invalid response, status code, or JSON serialization issues.
  - Specific errors for HTTP 409 (email already exists) or 400 (empty fields).
- **Behavior**:
  - Sends a POST request with a JSON body containing user details.
  - Expects a 201 status code and decodes an `IdResponse` from the response.
  - Throws descriptive errors for failures.

### Data Structures

#### `RegisterModel`

```swift
struct RegisterModel {
    let email: String
    let password: String
    let firstName: String
    let name: String
}
```

Input structure for user registration.

#### `IdResponse`

```swift
struct IdResponse: Codable {
    let id: String
}
```

Represents the API response for successful registration, containing the user ID.

#### `ErrorResponse`

```swift
struct ErrorResponse: Codable {
    let error: String
}
```

Represents the API error response with an error message.

---

## AuthService

### Overview

`AuthService` is an actor implementing the `AuthServiceProtocol` to manage authentication-related API calls, including login, token refresh, and authorized requests.

### Protocol

```swift
protocol AuthServiceProtocol {
    func loginUser(email: String, password: String, keepMeSignedIn: Bool) async throws -> (String, String, String, String)
    func requestToken(grantType: String, email: String?, password: String?, refreshToken: String?, keepMeSignedIn: Bool) async throws -> (String, String, String, String)
    func authorizedRequest(for endpoint: String) async throws -> Data
}
```

Defines the contract for authentication operations.

### Methods

#### `func loginUser(email: String, password: String, keepMeSignedIn: Bool) async throws -> (String, String, String, String)`

Logs in a user using email and password, calling `requestToken` with the `password` grant type.

- **Parameters**:
  - `email`: User’s email address.
  - `password`: User’s password.
  - `keepMeSignedIn`: Determines if tokens are stored persistently.
- **Returns**: A tuple of `(accessToken, refreshToken, userId, userRole)`.
- **Throws**: `NSError` for invalid responses, missing credentials, or API errors.

#### `func requestToken(grantType: String, email: String?, password: String?, refreshToken: String?, keepMeSignedIn: Bool) async throws -> (String, String, String, String)`

Requests an access token using either `password` or `refresh_token` grant type (`/auth/token/`).

- **Parameters**:
  - `grantType`: Either `"password"` or `"refresh_token"`.
  - `email`: Required for `password` grant type.
  - `password`: Required for `password` grant type.
  - `refreshToken`: Required for `refresh_token` grant type.
  - `keepMeSignedIn`: Persists tokens if `true`.
- **Returns**: A tuple of `(accessToken, refreshToken, userId, userRole)`.
- **Throws**: `NSError` for invalid grant type, missing parameters, or API errors.
- **Behavior**:
  - Sends a POST request with form-urlencoded body.
  - Decodes a `TokenResponse` and stores tokens via `TokenStorage`.
  - Encodes special characters in refresh tokens for safe URL usage.

#### `func authorizedRequest(for endpoint: String) async throws -> Data`

Performs an authorized GET request to the specified endpoint using a Bearer token.

- **Parameter**: `endpoint` - The API endpoint (e.g., `api/v1/some-resource`).
- **Returns**: The response `Data` from the API.
- **Throws**: `NSError` for invalid responses, expired tokens, or failed refresh attempts.
- **Behavior**:
  - Checks for a valid access token, refreshing it if expired or missing.
  - Sends a GET request with Bearer token authorization.
  - Returns raw response data for further processing.

#### `private func refreshAccessTokenIfNeeded() async throws -> String`

Refreshes the access token using a stored refresh token.

- **Returns**: A new access token.
- **Throws**: `NSError` if no refresh token is available or the refresh fails.
- **Behavior**: Calls `requestToken` with `refresh_token` grant type.

#### `func encodeSpecialCharacters(_ token: String) -> String`

Encodes special characters in tokens (e.g., `/`, `+`, `=`) for safe URL usage.

- **Parameter**: `token` - The token to encode.
- **Returns**: The encoded token string.

### Data Structures

#### `TokenResponse`

```swift
struct TokenResponse: Codable {
    let access_token: String
    let refresh_token: String
    let token_type: String
    let expires_in: Int
    let properties: UserProperties
}
```

Represents the API response for token requests, including access and refresh tokens, expiration time, and user properties.

#### `UserProperties`

```swift
struct UserProperties: Codable {
    let id: String
    let role: String
}
```

Contains user-specific information returned in token responses.

---

## Dependencies

- **APIConfig**: Provides the base URL for API requests.
- **URLSession**: Handles HTTP requests.
- **JSONSerialization/JSONDecoder**: Manages JSON encoding and decoding.
- **TokenStorage**: Stores and retrieves access/refresh tokens and expiration data.

---

## Notes

- Both services are implemented as actors to ensure thread safety for concurrent access.
- Unlike the Android `RetrofitClient`, this implementation uses `URLSession` directly, without a third-party networking library.
- `ApiService` currently supports only user registration, while `AuthService` covers authentication and token management. Other areas (e.g., property, room, furniture management) are handled by other components (e.g., `RoomManager`).
- Error handling is detailed, with specific messages for common HTTP status codes (e.g., 400, 409).
- The `authorizedRequest` method supports GET requests; other HTTP methods require additional implementation.
- Token refresh logic ensures seamless access for authorized requests, with fallback to re-login if no refresh token is available.
