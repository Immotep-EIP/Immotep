# UI Components

Base reusable components styled to match the visual identity of the Keyz app.
They wrap common Material3 elements like `TextField`, `Button`, and `IconButton` to ensure design consistency and reduce repetitive boilerplate.

---

## `BackButton`

* A minimal back navigation button using the `ChevronLeft` icon.
* Applies app theme colors and exposes a test tag for instrumentation.

### Parameters

* `onClick: () -> Unit`
  Callback invoked when the button is pressed.

---

## `DateRangeInput`

* An input field to select a date using a modal date picker.
* The field is read-only and formatted as `MM/dd/yyyy`.

### Parameters

* `currentDate: Long`
  Pre-selected timestamp (in millis) to display.

* `onDateSelected: (Long?) -> Unit`
  Callback triggered when a date is picked.

* `label: String`
  Optional label for the input field (default: `"Date"`).

* `errorMessage: String?`
  Displays an error below the field if not null.

* `globalTestTag: String`
  Base tag used to identify UI elements in tests.

---

## `DropDown<T>`

* Custom dropdown menu using a clickable box and `DropdownMenu`.
* Styled with borders, corner radius, and error display.

### Parameters

* `items: List<DropDownItem<T>>`
  Available options with label and value.

* `selectedItem: T`
  Currently selected item (matches `value`).

* `onItemSelected: (T) -> Unit`
  Callback invoked with the selected value.

* `error: String?`
  Displays an error below if present.

* `testTag: String`
  Identifier for UI testing (default: `"dropDown"`).

---

## `DropDownItem<T>`

A simple data class used in dropdowns.

```kotlin
data class DropDownItem<T>(
    val label: String,
    val value: T
)
```

---

## `PasswordInput`

* Styled password input with toggle visibility.
* Uses lock icon in a trailing icon button.

### Parameters

* `value: String`
  Input text value.

* `onValueChange: (String) -> Unit`
  Callback for changes in text field.

* `label: String`
  Field label.

* `errorMessage: String?`
  Displays an error if present.

* `iconButtonTestId: String`
  Test ID for the trailing icon (toggle button).

* (Other common `TextField` parameters are also supported, including `placeholder`, `leadingIcon`, `keyboardOptions`, etc.)

---

## `StyledButton`

* App-wide button with customizable error and loading states.
* Primary styling uses `colorScheme.secondary`.

### Parameters

* `onClick: () -> Unit`
  Button action (disabled if loading).

* `text: String`
  Text displayed inside the button.

* `error: Boolean`
  Uses `error` color theme if true.

* `isLoading: Boolean`
  Shows a spinner next to the text if true.

* `testTag: String`
  Identifier for UI testing (default: `"StyledButton"`).

---

## `OutlinedTextField`

* Reusable custom `OutlinedTextField` with:

  * Styled label
  * Optional icons
  * Error and helper messages
  * Visual transformations

### Parameters

* `value: String`
  Current input text.

* `onValueChange: (String) -> Unit`
  Callback for text updates.

* `label: String`
  Label displayed above the input.

* `errorMessage: String?`
  Shows error text if provided.

* `helperMessage: String?`
  Fallback message shown if no error.

* (Supports other `TextField` features like `placeholder`, `leadingIcon`, `trailingIcon`, `keyboardOptions`, `visualTransformation`, etc.)

---

## Helper Function

### `convertMillisToDate(millis: Long): String`

* Converts a `Long` timestamp into a formatted string (`MM/dd/yyyy`).
* Locale-aware using the system default.

---

## Notes

* These components help enforce consistency in spacing, colors, and theming.
* Fully testable with consistent `testTag` usage.
* Built for reuse across authentication, dashboard, and form-related flows.
