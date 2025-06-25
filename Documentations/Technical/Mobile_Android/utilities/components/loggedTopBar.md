# LoggedTopBar

---

## UI Component

### `LoggedTopBar`

* Top bar displayed for logged-in users, containing the app logo and name, with a built-in logout trigger.

---

## Behavior & Styling

* **Layout**:

  * Uses a horizontal `Row` with:

    * Vertical alignment: `CenterVertically`
    * Height: `35.dp`
    * Horizontal padding: `10.dp`
    * Bottom border drawn manually with a light gray line using `drawBehind`.

* **App Logo**:

  * Displays the app logo using `Image`.
  * Image resource selected dynamically depending on current theme (`dark` or `light`) via `ThemeUtils.getIcon(...)`.
  * Icon size: `35.dp`
  * Has a test tag: `"loggedTopBarImage"`
  * Click behavior:

    * Triggers logout using `AuthService.onLogout(...)` inside a coroutine launched from the `ViewModel`’s scope.
    * Requires `NavController` and current `apiService` (from `LocalApiService`).

* **App Name**:

  * Displayed next to the icon using `Text`.
  * Font size: `20.sp`, color: `MaterialTheme.colorScheme.primary`, weight: `Medium`.
  * Has a test tag: `"loggedTopBarText"`

* **Spacer**:

  * Fills remaining horizontal space to push logo and text to the left.
  * Ensures top bar stretches full width of the container.

---

## ViewModel

### `LoggedTopBarViewModel`

* Empty `ViewModel`, used to provide a coroutine scope (`viewModelScope`) for logout operations.

---

## Usage Example

```kotlin
LoggedTopBar(navController = navController)
```

---

## Notes

* Integrates logout directly into the logo click for quick access.
* Custom theme-aware logo support.
* Lightweight and testable (with proper test tags for UI testing).
