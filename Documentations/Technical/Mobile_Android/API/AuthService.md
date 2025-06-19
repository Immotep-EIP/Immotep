# AuthService Technical Documentation

## Overview

`AuthService` is a Kotlin class responsible for handling user authentication, token management, and registration by interacting with an external API through the `ApiService`. It uses Android's `DataStore` to securely store tokens and user roles locally and manages token expiration and refresh logic.

---

## Data Classes

### `LoginResponse`

Represents the response from the login API call.

| Property        | Type               | Description                              |
| --------------- | ------------------ | ---------------------------------------- |
| `access_token`  | `String`           | The access token for API calls.          |
| `refresh_token` | `String`           | Token used to refresh access token.      |
| `token_type`    | `String`           | Type of the token (usually "Bearer").    |
| `expires_in`    | `Int`              | Lifetime of the access token in seconds. |
| `properties`    | `Map<String, Any>` | Additional properties returned.          |

---

### `RegistrationInput`

Input data required to register a new user.

| Property    | Type     | Description        |
| ----------- | -------- | ------------------ |
| `email`     | `String` | User's email.      |
| `password`  | `String` | User's password.   |
| `firstName` | `String` | User's first name. |
| `lastName`  | `String` | User's last name.  |

---

### `RegistrationResponse`

Response data returned after user registration.

| Property     | Type     | Description                 |
| ------------ | -------- | --------------------------- |
| `id`         | `String` | User ID.                    |
| `email`      | `String` | Registered email.           |
| `firstname`  | `String` | First name.                 |
| `lastname`   | `String` | Last name.                  |
| `role`       | `String` | User role.                  |
| `created_at` | `String` | Account creation timestamp. |
| `updated_at` | `String` | Account update timestamp.   |

---

## Class: `AuthService`

### Constructor Parameters

| Parameter    | Type                     | Description                          |
| ------------ | ------------------------ | ------------------------------------ |
| `dataStore`  | `DataStore<Preferences>` | Android DataStore for local storage. |
| `apiService` | `ApiService`             | Retrofit API service interface.      |

---

### Public Methods

#### `suspend fun onLogin(username: String, password: String)`

* Authenticates the user with the API.
* Stores the access and refresh tokens locally.
* Retrieves the user profile to determine role and stores if the user is an owner.
* Throws exceptions with HTTP code information on failure.

---

#### `suspend fun getToken(): String`

* Returns the current access token.
* Automatically refreshes the token if expired.
* Throws an exception if no token is stored.

---

#### `suspend fun getBearerToken(): String`

* Returns the access token formatted as a Bearer token for authorization headers.

---

#### `suspend fun deleteToken()`

* Deletes both access and refresh tokens from local storage.

---

#### `suspend fun onLogout(navController: NavController)`

* Clears tokens.
* Navigates the app to the login screen.

---

#### `suspend fun register(registrationInput: RegistrationInput): RegistrationResponse`

* Registers a new user via the API.
* Throws exceptions with HTTP code information on failure.

---

#### `suspend fun isUserOwner(): Boolean`

* Checks local storage to determine if the current user is an owner.
* Returns `true` if owner, `false` otherwise.

---

### Private Methods

#### `private suspend fun store(accessToken: String, refreshToken: String?, expiresIn: Int)`

* Stores the access token, refresh token, and expiration time in `DataStore`.
* Sets expiration time to 5 minutes before actual expiry for safety.

---

#### `private suspend fun refreshToken()`

* Uses the stored refresh token to get a new access token from the API.
* Stores the refreshed tokens.
* Throws exception on failure.

---

#### `private suspend fun isAccessTokenExpired(): Boolean`

* Checks if the access token has expired based on stored expiration time.

---

### Companion Object Keys

| Key             | Description                                                      |
| --------------- | ---------------------------------------------------------------- |
| `ACCESS_TOKEN`  | Preference key for access token.                                 |
| `REFRESH_TOKEN` | Preference key for refresh token.                                |
| `EXPIRES_IN`    | Preference key for token expiration.                             |
| `IS_OWNER`      | Preference key to store owner role flag (`"true"` or `"false"`). |

---

## Error Handling

* API call failures throw exceptions with decoded HTTP error codes via `decodeRetroFitMessagesToHttpCodes`.
* Login, token refresh, and registration errors are explicitly caught and rethrown with detailed messages.
* Profile fetching errors during login are caught but do not prevent login success.

---

## Notes

* Token expiration is handled proactively by subtracting 5 minutes from the expiry period.
* Navigation on logout assumes a screen named `"login"` exists.
* Uses Kotlin coroutines for all asynchronous API calls and DataStore operations.
