# InitialFadeIn

---

## UI Component

### `InitialFadeIn`

* Provides a fade-in animation for its content when first composed.

---

## Parameters

* `durationMs: Int = 1000`
  Duration of the fade-in animation in milliseconds (default 1000ms).

* `content: @Composable AnimatedVisibilityScope.() -> Unit`
  Composable content to be shown with the fade-in effect.

---

## Behavior

* Uses `rememberSaveable` to track visibility state, initially false.
* On first composition (`LaunchedEffect` with `Unit` key), sets visibility to true, triggering the fade-in.
* Wraps the content inside `AnimatedVisibility` with a fade-in animation using a tween easing.

---

## Usage Example

```kotlin
InitialFadeIn(durationMs = 1500) {
    Text("Hello with fade-in!")
}
```

---

## Notes

* Animation runs only once when the composable first enters composition.
* `AnimatedVisibilityScope` receiver allows usage of scoped animation APIs inside content if needed.
