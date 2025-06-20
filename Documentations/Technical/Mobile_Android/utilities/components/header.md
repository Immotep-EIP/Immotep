# Header

---

## UI Component

### `Header`

* Displays the app’s header with a logo icon and app name text in a horizontal row.

---

## Behavior

* Detects the current system theme (dark or light) and selects the appropriate logo icon.
* Shows the logo on the left, sized 50dp with right padding of 10dp.
* Displays the app name text next to the logo with font size 30sp.
* Text color uses the theme’s primary color for consistency.
* The entire row is vertically centered.

---

## Styling

* Uses `MaterialTheme.colorScheme.primary` for the app name color.
* Image and text aligned vertically centered.
* Applies test tag `"header"` for UI testing identification.

---

## Usage Example

```kotlin
Header()
```

---

## Notes

* The logo icon is resolved dynamically by `ThemeUtils.getIcon()` based on the system dark mode.
* The content description for accessibility is localized from resources (`R.string.immotep_logo_desc`).
