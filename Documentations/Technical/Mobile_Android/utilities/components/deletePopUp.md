# DeletePopUp

---

## UI Component

### `DeletePopUp`

* Displays a modal bottom sheet with a delete button.
* When the delete button is clicked, shows a confirmation alert dialog before performing the delete action.
* Uses Material3 theming for colors and shapes.

---

## Parameters

* `open: Boolean`
  Controls whether the modal bottom sheet is visible.

* `delete: () -> Unit`
  Callback invoked to perform the deletion when confirmed.

* `close: () -> Unit`
  Callback to close the modal bottom sheet.

* `globalName: String`
  Name displayed in the delete button and confirmation dialog title.

* `detailedName: String`
  Name displayed in the confirmation dialog message for detail.

---

## Behavior

* Shows a `ModalBottomSheet` when `open` is true.
* Inside the sheet, a delete button labeled "Delete \[globalName]" is shown.
* Clicking the button opens a confirmation `AlertDialog`.
* The confirmation dialog asks:

  * Title: "Delete \[globalName] ?"
  * Text: "Are you sure to delete \[detailedName] ?"
* Confirming triggers the `delete()` callback, closes the dialog and the modal.
* Dismissing the dialog or sheet calls the `close()` callback.

---

## Styling

* Alert dialog and button use rounded corners (10dp and 5dp respectively).
* Colors for error actions come from Material3 `colorScheme.errorContainer` and `onError`.
* Text colors in dialog use `colorScheme.secondary`.

---

## Usage Example

```kotlin
DeletePopUp(
    open = showDeleteModal,
    delete = { performDelete() },
    close = { showDeleteModal = false },
    globalName = "Tenant",
    detailedName = "John Doe"
)
```

---

## Notes

* State for the confirmation dialog visibility is managed internally with `rememberSaveable`.
* Provides a safe user experience by requiring confirmation before delete.
