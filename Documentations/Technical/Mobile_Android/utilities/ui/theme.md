# Theme

Defines the color schemes and theming setup for the Keyz app using Material3.
Supports light, dark, and contrast variants for accessibility.

---

## Color Schemes

### Light and Dark Color Schemes

* `lightScheme` and `darkScheme` define the primary color palettes for light and dark modes respectively.
* Each includes comprehensive color roles such as `primary`, `onPrimary`, `primaryContainer`, `secondary`, `error`, `background`, `surface`, `outline`, and many more for fine-grained theming.

### Medium and High Contrast Variants

* `mediumContrastLightColorScheme` and `highContrastLightColorScheme` provide accessible color options for light mode with increased contrast.
* `mediumContrastDarkColorScheme` and `highContrastDarkColorScheme` provide similar accessible variants for dark mode.

These ensure the app meets accessibility standards by improving readability for users with visual impairments.

---

## `ColorFamily` Data Class

A convenient immutable container grouping related colors:

* `color: Color` — Primary color.
* `onColor: Color` — Color for content displayed on top of the primary color.
* `colorContainer: Color` — Container/background color variant.
* `onColorContainer: Color` — Content color used on top of the container color.

Example usage for managing related color sets.

---

## `unspecified_scheme`

A `ColorFamily` instance with all colors set to `Color.Unspecified`, useful as a placeholder or default.

---

## `AppTheme` Composable

Sets the Material3 theme for the app, including color scheme and typography.

### Parameters

* `darkTheme: Boolean`
  Controls whether to use dark or light theme. Defaults to system setting (`isSystemInDarkTheme()`).

* `content: @Composable () -> Unit`
  The composable content to wrap with the theme.

### Behavior

* Selects the appropriate color scheme based on `darkTheme`.
* Applies `AppTypography` (typography definitions imported from the theme package).
* Wraps content inside `MaterialTheme` with the selected colors and typography.

---

## Summary

* Centralized theming with detailed light/dark and contrast variants.
* Improves visual consistency and accessibility throughout the app.
* Easily customizable by changing or extending the defined color schemes.
