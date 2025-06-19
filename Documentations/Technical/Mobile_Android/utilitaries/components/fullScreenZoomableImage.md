# FullscreenZoomableImage

---

## UI Component

### `FullscreenZoomableImage`

* Displays an image in a fullscreen modal with zoom, pan, and rotation gestures.
* Allows user to interactively scale, move, and rotate the displayed image.
* Tapping outside the image or on the background closes the fullscreen modal.

---

## Parameters

* `uri: Uri?`
  The image URI to display fullscreen. If `null`, the modal is closed.

* `onClose: () -> Unit`
  Callback to close the fullscreen modal.

---

## Behavior

* Uses `BigModalLayout` to present a modal covering 90% of the screen height with a black background.
* Applies gesture detection via `detectTransformGestures` to handle:

  * Zoom (pinch to zoom, constrained between 1x and 5x).
  * Pan (drag to move).
  * Rotation.
* Image transforms (scale, translation, rotation) are applied with `graphicsLayer`.
* Clicking the background or image triggers `onClose`.

---

## Styling

* Background color is black to emphasize the image.
* Image content scaled with `ContentScale.Fit` to maintain aspect ratio.
* The modal and image fill the maximum available size.

---

## Usage Example

```kotlin
var fullscreenImageUri by remember { mutableStateOf<Uri?>(null) }

if (fullscreenImageUri != null) {
    FullscreenZoomableImage(uri = fullscreenImageUri, onClose = { fullscreenImageUri = null })
}
```

---

## Notes

* The component supports smooth multitouch gestures on the image.
* State for scale, offset, and rotation is remembered and updated dynamically.
* Useful for photo galleries or any image preview functionality needing rich user interaction.
