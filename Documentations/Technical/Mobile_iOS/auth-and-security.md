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
  - "Keep me signed in" checkbox

The **"Keep me signed in"** option stores tokens in `localStorage` if checked, ensuring they persist across sessions. If unchecked, tokens are stored in `sessionStorage` for the current session only.

## 2. User Logout

The **Logout Button** allows users to:

- Clear all tokens from both `localStorage` and `sessionStorage`.
- Be redirected to the login page.

## 3. Token Types and Expiration

We use two types of tokens:

- **Access Token**: Expires after 24 hours and is included in every API request for secure communication.
- **Refresh Token**: Used to obtain a new access token when the current one expires.

When the `access_token` expires, the app will attempt to use the `refresh_token` to request a new token from the API.

