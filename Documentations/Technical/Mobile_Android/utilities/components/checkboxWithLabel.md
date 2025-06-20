# CheckBoxWithLabel

---

## UI Component

### `CheckBoxWithLabel`

* Displays a checkbox with a text label beside it.
* Shows an optional error message below the checkbox if provided.
* Uses Material3 theming for colors and typography.

---

## Parameters

* `modifier: Modifier = Modifier`
  Optional modifier to apply to the checkbox.

* `label: String`
  Text label shown next to the checkbox.

* `isChecked: Boolean`
  Current checked state of the checkbox.

* `onCheckedChange: (Boolean) -> Unit`
  Callback invoked when the checkbox is toggled.

* `errorMessage: String? = null`
  Optional error message shown below the checkbox in error color.

---

## Behavior

* Checkbox and label are horizontally aligned and vertically centered.
* Label uses primary color and font size 12sp.
* If `errorMessage` is not null, it is displayed below in error color at 10sp font size.

---

## Usage Example

```kotlin
CheckBoxWithLabel(
    label = "Accept Terms",
    isChecked = accepted,
    onCheckedChange = { accepted = it },
    errorMessage = if (!accepted) "You must accept the terms" else null
)
```

---

## Notes

* Checkbox uses primary color for unchecked state.
* The component ensures consistent spacing and alignment between checkbox, label, and error text.
