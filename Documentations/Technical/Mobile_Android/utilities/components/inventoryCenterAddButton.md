# InventoryCenterAddButton

---

## UI Component

### `InventoryCenterAddButton`

* A full-width button with a centered add icon, styled for inventory screens.

---

## Parameters

* `onClick: () -> Unit`
  Callback triggered when the button is clicked.

* `testTag: String`
  Tag for UI testing purposes.

---

## Behavior & Styling

* Uses a `Button` with:

  * A 1 dp border using the theme’s primary color.
  * Rounded corners with a 5 dp radius.
  * Background color from the theme’s background.
  * Modifier to add vertical padding (10 dp top and bottom) and fill the maximum width.
  * Test tag set for UI testing.

* Inside the button, a `Row` centers its content horizontally and vertically.

* Displays the outlined add circle icon (`Icons.Outlined.AddCircleOutline`), tinted with `onPrimaryContainer` color from the theme.

---

## Usage Example

```kotlin
InventoryCenterAddButton(
    onClick = { /* handle add */ },
    testTag = "addInventoryButton"
)
```

---

## Notes

* Suitable as a prominent add action button in inventory-related UI.
* Combines Material theming colors with Material icons for consistent design.
