# InternalLoading

---

## UI Component

### `InternalLoading`

* Displays a full-screen centered loading indicator.

---

## Behavior

* Uses a `Column` that fills the entire available size.
* Centers its content vertically and horizontally.
* Shows a `CircularProgressIndicator` sized 100x100 dp.
* The progress indicator uses the theme’s secondary color with a track color from `onSurfaceVariant`.

---

## Styling

* `CircularProgressIndicator` is customized with:

  * Width and height set to 100 dp.
  * `color` from `MaterialTheme.colorScheme.secondary`.
  * `trackColor` from `MaterialTheme.colorScheme.onSurfaceVariant`.

---

## Usage Example

```kotlin
InternalLoading()
```

---

## Notes

* Suitable as a simple loading screen indicator while waiting for data or processes.
* Uses Material 3 theming colors for consistent UI appearance.
