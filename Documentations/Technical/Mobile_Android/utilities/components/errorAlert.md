# ErrorAlert

---

## UI Component

### `ErrorAlert`

* Displays a styled error message banner with an error icon.
* Maps HTTP error codes to localized user-friendly messages.
* Shows a custom message if provided or if the error code is unknown.
* Uses Material3 theming for background and text colors.

---

## Parameters

* `code: Int?`
  HTTP status code to determine the error message.
  If `null`, no alert is shown unless `customMessage` is provided.

* `login: Boolean?`
  Optional flag to customize the message for 401 errors depending on login context.

* `customMessage: String? = null`
  Optional custom error message to display instead of default code mappings.

---

## Behavior

* If both `code` and `customMessage` are `null`, nothing is displayed.
* Maps common HTTP codes (e.g., 400, 401, 403, 404, 500, etc.) to string resources.
* For 401 errors, the message depends on `login` flag (login error or unauthorized).
* Displays the message with a white text on a red error container background.
* Shows an error icon (`ReleaseAlert`) on the left side.

---

## Styling

* Background uses `errorContainer` color from Material3.
* Text color is white.
* Rounded corners with 10dp radius.
* Padding of 10dp around content.
* Horizontally aligned row with centered vertical alignment.
* Contains spacing between icon and text.

---

## Helper Function

### `decodeRetroFitMessagesToHttpCodes`

* Parses an exception’s message to extract HTTP status code if present.
* Returns the HTTP code as an `Int`.
* Returns `-1` if no valid code can be extracted.
* Specially treats messages starting with "Failed to connect" as HTTP 500 error.

---

## Usage Example

```kotlin
val errorCode = decodeRetroFitMessagesToHttpCodes(exception)
ErrorAlert(code = errorCode, login = false)
```

---

## Notes

* Designed to provide user-friendly error feedback for network/API operations.
* Relies on predefined string resources for localization and consistent messages.
