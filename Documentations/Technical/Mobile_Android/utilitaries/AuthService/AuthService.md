Here’s a summary of the `AuthService` class:

---

### Overview
The `AuthService` class manages user authentication within the app, providing methods for logging in, logging out, refreshing tokens, and storing/retrieving tokens from the data store. It communicates with a backend API via the `ApiClient` to authenticate users and manage session tokens.

### Key Responsibilities
- **Login**: The `onLogin` method sends a request to the backend API to authenticate a user using their username and password. If successful, it stores the returned `access_token` and `refresh_token` in the local data store.
  
- **Token Refresh**: The `refreshToken` method retrieves the stored `refresh_token` and uses it to request a new `access_token` from the backend. It stores the new tokens if successful.

- **Token Retrieval**: The `getToken` method fetches the stored `access_token` from the data store, while `getBearerToken` adds the "Bearer" prefix for use in authorization headers.

- **Logout**: The `onLogout` method removes the stored tokens from the data store and navigates the user to the login screen.

### Data Storage
- The `access_token` and `refresh_token` are stored in a `DataStore` instance, which persists these values across app sessions.

### Error Handling
- Errors during network requests (login or refresh) are caught and decoded using the `decodeRetroFitMessagesToHttpCodes` utility function. If an error occurs, an exception is thrown with the HTTP error code.

### Token Storage Keys
- `ACCESS_TOKEN`: A key to store the `access_token` in the data store.
- `REFRESH_TOKEN`: A key to store the `refresh_token` in the data store.

---

### Code Breakdown:
- **Methods**:
  - `onLogin`: Logs the user in with the provided credentials and stores the tokens.
  - `refreshToken`: Refreshes the `access_token` using the `refresh_token`.
  - `getToken`: Retrieves the stored `access_token`.
  - `getBearerToken`: Retrieves the stored `access_token` prefixed with "Bearer" for API authorization.
  - `onLogout`: Logs the user out by clearing the tokens from storage.

- **Companion Object**: 
  - `ACCESS_TOKEN` and `REFRESH_TOKEN` are keys used to store/retrieve tokens from the data store.

---

This class plays a critical role in managing user authentication and session states within the app, ensuring secure API communication by handling token storage and renewal.