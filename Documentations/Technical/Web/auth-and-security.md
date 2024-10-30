# Authentication and Token Management

## 1. User Registration and Login

- **Registration Page**: A form is provided for users to sign up with the following fields:

  - First Name
  - Last Name
  - Email
  - Password
  - Confirm Password

- **Login Page**: Users log in using:
  - Email
  - Password
  - Remember Me checkbox

The **Remember Me** option stores tokens in `localStorage` if checked, ensuring they persist across sessions. If unchecked, tokens are stored in `sessionStorage` for the current session only.

## 2. User Logout

The **Logout Button** allows users to:

- Clear all tokens from both `localStorage` and `sessionStorage`.
- Be redirected to the login page.

## 3. Token Types and Expiration

We use two types of tokens:

- **Access Token**: Expires after 24 hours and is included in every API request for secure communication.
- **Refresh Token**: Used to obtain a new access token when the current one expires.

When the `access_token` expires, the app will attempt to use the `refresh_token` to request a new token from the API.

## 4. Axios Interceptors for Token Management

- **Request Interceptor**: Each outgoing API request is intercepted to add the `access_token` to the `Authorization` header.
- **Response Interceptor**: If a 401 `Unauthorized` error is received, the interceptor checks if a valid `refresh_token` exists and requests a new `access_token` from the API.

If successful, the new token is saved, and the original request is retried with the updated credentials.

## 5. Token Expiry and User Session

- **Access Token**: Expires after 24 hours.
- **15-Minute Grace Period**: If the `access_token` expires, users have a 15-minute window where the `refresh_token` can still be used to obtain a new `access_token`. After this grace period, users are automatically logged out.
