# DeleteOrSeePictureModal & AddingPicturesCarousel

---

## UI Components

### `DeleteOrSeePictureModal`

* Shows a modal bottom sheet when a picture is selected.
* Displays two buttons:

  * **See Picture**: Opens the selected image in fullscreen.
  * **Delete Picture** (optional): Deletes the selected picture.
* Closes the modal after an action is taken.

### `AddingPicturesCarousel`

* Displays a horizontal carousel of pictures.
* Supports two modes for pictures:

  * List of URIs (`uriPictures`).
  * List of Base64-encoded strings (`stringPictures`).
* Allows adding pictures via:

  * Gallery picker.
  * Camera capture (`TakePhotoButton`).
* Shows an error message below the carousel if `error` is not null.
* Displays a modal to choose between gallery or camera when adding a picture.
* Shows a modal (`DeleteOrSeePictureModal`) when tapping on a picture to delete or view fullscreen.
* Opens fullscreen image viewer (`FullscreenZoomableImage`) when requested.

---

## State & Callbacks

* **Internal state:**

  * `chooseOpen`: Whether the add-picture modal is open.
  * `pictureSelected`: Currently selected picture (URI and index) for deletion/view.
  * `pictureFullScreen`: URI of the picture shown in fullscreen.

* **Parameters:**

  * `uriPictures: List<Uri>?` — pictures as URIs.
  * `addPicture: ((Uri) -> Unit)?` — callback to add a picture.
  * `removePicture: ((Int) -> Unit)?` — callback to remove a picture by index.
  * `stringPictures: List<String>?` — pictures as Base64 strings.
  * `maxPictures: Int` — maximum pictures allowed (default 10).
  * `error: String?` — optional error message.

---

## Behavior Overview

1. Shows carousel of pictures (URI or Base64).
2. If room to add more pictures, shows an "Add picture" button.
3. Clicking "Add picture" opens modal to pick from gallery or take photo.
4. Picking or taking photo calls `addPicture`.
5. Clicking a picture opens `DeleteOrSeePictureModal`.
6. From modal, user can view fullscreen or delete picture.
7. Shows error message below carousel if present.

---

## Usage Example

```kotlin
AddingPicturesCarousel(
    uriPictures = myUriList,
    addPicture = { uri -> /* add picture logic */ },
    removePicture = { index -> /* remove picture logic */ },
    maxPictures = 5,
    error = if (someError) "Too many pictures" else null
)
```

---

## Notes

* Uses Jetpack Compose Material3 components and Carousel API.
* Uses Coil `AsyncImage` for URI images.
* Decodes Base64 strings to bitmaps for display.
* Uses `rememberLauncherForActivityResult` to launch system photo picker.
* `TakePhotoButton` component handles camera capture (not shown here).
* The carousel adapts layout and spacing for smooth UI experience.
