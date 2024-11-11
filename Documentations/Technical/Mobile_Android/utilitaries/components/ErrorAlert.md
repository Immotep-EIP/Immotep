## ErrorAlert Component Documentation

This section provides an overview of the `ErrorAlert` component in the Immotep mobile application.

### Overview

The `ErrorAlert` composable component displays an error message in a styled container based on an HTTP status code. It provides feedback to the user regarding specific error conditions (e.g., bad request, unauthorized access, or server error) with visual and textual cues.

### Parameters

* **code** (`Int?`): The HTTP error code used to determine the appropriate error message. If `null`, no error is displayed.
* **login** (`Boolean?`): Optional flag indicating if the error occurs in a login context, used to display a specific message for login-related errors when `code` is 401 (Unauthorized).

### Components

* **Error Text (`errorText`)**: Displays an error message based on the `code` value. The message is derived from the `stringResource` values for standard HTTP status codes, or a default "Unknown error" message if the code is unrecognized.

* **Error Icon**: An image icon representing an error state, displayed alongside the error text for visual emphasis (using `Release_alert` icon).

* **Styled Container**: A rounded container with background color from `MaterialTheme.colorScheme.errorContainer` and padding, which visually highlights the error message.

### Error Messages

The `ErrorAlert` component maps the `code` parameter to specific error messages:
- **400**: Bad Request
- **401**: Unauthorized (or Login Error if `login` is true)
- **403**: Forbidden
- **404**: Not Found
- **405**: Method Not Allowed
- **406**: Not Acceptable
- **409**: Conflict
- **410**: Gone
- **413**: Request Entity Too Large
- **415**: Unsupported Media Type
- **429**: Too Many Requests
- **500**: Internal Server Error
- **501**: Not Implemented
- **502**: Bad Gateway
- **503**: Service Unavailable
- **504**: Gateway Timeout
- **Unknown Code**: Displays a generic "Unknown error" message.

### Usage

The `ErrorAlert` component is suitable for displaying error messages based on HTTP responses. It is commonly used in form validation and API error handling, providing feedback directly on the screen when errors occur.

### Example Usage

```kotlin
ErrorAlert(
    code = errorCode,
    login = isLoginError
)
```

### Visual Layout

1. **Row Layout**: Displays the icon and text side-by-side, with the error text next to the icon.
2. **Container Styling**: The entire row is wrapped in a container with error-themed background color and rounded corners for emphasis.

### Interaction

The `ErrorAlert` component is non-interactive and displays error information based on the provided `code` and `login` parameters.

### Supporting Function

The `decodeRetroFitMessagesToHttpCodes` function parses HTTP error codes from exception messages and returns the error code for mapping within `ErrorAlert`.