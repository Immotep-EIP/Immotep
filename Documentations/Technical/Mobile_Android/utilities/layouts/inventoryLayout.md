# InventoryLayout

---

## UI Layout

### `InventoryLayout`

* Fullscreen layout scaffold specifically designed for inventory-related pages.
* Composed of:

  * A custom **top bar** (`InventoryTopBar`) with a title and close icon.
  * A main **content** section that fills remaining space.

---

## Parameters

* `testTag: String`
  Identifier for Compose UI tests on the outer layout container.

* `onExit: () -> Unit`
  Callback triggered when the user taps the close icon in the top bar.

* `content: @Composable () -> Unit`
  Dynamic composable displayed as the inventory screen's main content.

---

## Components

### `InventoryTopBar(onExit: () -> Unit)`

* Top header bar with:

  * **Logo** — switches based on system theme using `ThemeUtils.getIcon(...)`.
  * **Title** — uses `R.string.inventory_title` and Material theme color.
  * **Close Button** — triggers `onExit()` when pressed.
* Includes a subtle bottom line separator via `drawBehind`.

---

## Layout Behavior

* Wraps all content in a vertical `Column`.
* Top area:

  * Fixed height bar via `InventoryTopBar`.
* Content area:

  * Uses `Modifier.weight(1f)` to fill vertical space.
  * Adds `10.dp` padding and internal `testTag("inventoryLayout")`.

---

## Tags for Testing

* `"inventoryTopBar"` — wrapper for the top bar.
* `"inventoryTopBarImage"` — for the app logo inside the bar.
* `"inventoryTopBarText"` — for the screen title text.
* `"inventoryTopBarCloseIcon"` — for the close button.
* `"inventoryLayout"` — for the inner content section.
* `"testTag"` — outermost layout identifier (parameterized).

---

## Notes

* Designed for full-screen inventory flows such as check-in/check-out.
* Easily reusable in various contexts with custom composable `content()`.
* Offers consistent branding and user interaction via logo and exit icon.
