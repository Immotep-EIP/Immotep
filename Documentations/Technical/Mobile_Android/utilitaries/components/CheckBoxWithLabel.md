## CheckBoxWithLabel Component Documentation

This section provides an overview of the `CheckBoxWithLabel` component in the Immotep mobile application. 

### Overview

The `CheckBoxWithLabel` composable component provides a checkbox with a text label and optional error message. This component is designed to capture a user’s binary choice (e.g., agreeing to terms and conditions) and displays an error message when the checkbox selection is required but not selected.

### Parameters

* **label** (`String`): Text label displayed beside the checkbox, describing the option being selected.
* **isChecked** (`Boolean`): Current state of the checkbox, indicating whether it is selected (`true`) or not (`false`).
* **onCheckedChange** (`(Boolean) -> Unit`): Callback function invoked when the checkbox selection changes. Receives the new checkbox state as a parameter.
* **errorMessage** (`String?`): Optional error message displayed below the checkbox. Shown only when this parameter is non-null (e.g., validation feedback when the checkbox selection is required).
* **modifier** (`Modifier`): Modifier applied to the checkbox, allowing customization of layout properties such as padding and alignment.

### Components

* **Checkbox:** Displays the current selection state and invokes `onCheckedChange` when clicked. The `modifier` parameter can be used to adjust checkbox layout properties.
* **Label Text:** Displays the `label` parameter beside the checkbox. The text color is set to the primary color defined in `MaterialTheme.colorScheme` with a font size of `12.sp`.
* **Error Message Text:** Displays the `errorMessage` parameter below the checkbox and label. The text color is set to the error color defined in `MaterialTheme.colorScheme` with a font size of `10.sp`. This text is shown only if `errorMessage` is non-null.

### Usage

The `CheckBoxWithLabel` component is useful for scenarios where a binary option needs to be displayed with descriptive text, such as terms and conditions acceptance, or optional feature selection. This component is often used within forms and can be paired with error handling to provide validation feedback.

### Example Usage

```kotlin
CheckBoxWithLabel(
    label = "Agree to Terms and Conditions",
    isChecked = agreeToTerms,
    onCheckedChange = { newValue -> viewModel.setAgreeToTerms(newValue) },
    errorMessage = if (showError) "You must agree to the terms" else null,
    modifier = Modifier.padding(start = 8.dp)
)
```

### Visual Layout

1. **Row Layout**: Displays the checkbox and label text side-by-side, aligned vertically at the center.
2. **Error Message Layout**: The error message, if present, appears in a smaller font below the checkbox and label.

### Interactions

* **Checkbox Selection**: Users can select or deselect the checkbox, which updates the `isChecked` state and triggers the `onCheckedChange` callback.
* **Error Display**: If the component is used in a context where the selection is required, an error message can be passed through `errorMessage` to guide the user when validation fails.