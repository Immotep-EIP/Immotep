# ThemeUtils

## Overview

`ThemeUtils` is a utility object that provides theme-related helper functions, primarily to select appropriate resources based on the current app theme.

---

## Functions

### `getIcon(isDark: Boolean): Int`

* Returns the drawable resource ID of the app logo depending on the dark mode status.
* Returns a white logo resource for dark themes.
* Returns a blue logo resource for light themes.

---

## Summary

* Helps in dynamically selecting icons or resources based on the app's light or dark theme.
* Simplifies theme-aware resource management within the app.
