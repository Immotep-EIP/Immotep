# LoadingDialog

---

## UI Component

### `LoadingDialog`

* Displays a modal dialog containing a centered circular loading indicator.

---

## Parameters

* `isOpen: Boolean`
  Controls the visibility of the dialog. When `false`, the dialog is not shown.

---

## Behavior & Styling

* If `isOpen` is `false`, the composable returns early without rendering anything.
* When open, shows a `Dialog` that cannot be dismissed by the user (`onDismissRequest` is empty).
* Inside the dialog, displays a `CircularProgressIndicator`:

  * Sized 100x100 dp.
  * Uses the Material 3 theme’s secondary color for the indicator.
  * Uses `onSurfaceVariant` for the track color.

---

## Usage Example

```kotlin
LoadingDialog(isOpen = isLoading)
```

---

## Notes

* Useful to block user interaction while a background process is running.
* The dialog prevents dismissal to ensure process completion before interaction.
