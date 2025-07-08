# ImagePicker Documentation

## Overview

The `ImagePicker.swift` file in the Keyz app provides a SwiftUI `UIViewControllerRepresentable` wrapper for `UIImagePickerController`, enabling image selection from the camera or photo library, along with utility functions for base64 image conversion.

---

## Functionality

### `ImagePicker`
* **Purpose**: A SwiftUI view that wraps `UIImagePickerController` for selecting images.
* **Properties**:
  * `@Binding sourceType: UIImagePickerController.SourceType`: The source (e.g., `.camera`, `.photoLibrary`).
  * `@Binding selectedImage: UIImage?`: The selected image, updated after picking.
* **Features**:
  * Configures `UIImagePickerController` with the specified source type and disables editing.
  * Uses a `Coordinator` to handle delegate callbacks.
  * Presents modally in full-screen mode.
* **Usage Example**:
  ```swift
  @State private var sourceType: UIImagePickerController.SourceType = .photoLibrary
  @State private var selectedImage: UIImage?
  var body: some View {
      ImagePicker(sourceType: $sourceType, selectedImage: $selectedImage)
  }
  ```

### `convertUIImagesToBase64(_:) -> [String]`
* **Purpose**: Converts an array of `UIImage` objects to base64-encoded strings.
* **Parameters**:
  * `images: [UIImage]`: The array of images to convert.
* **Features**:
  * Uses JPEG compression with quality 0.8.
  * Returns an array of base64-encoded strings, excluding any images that fail conversion.
* **Usage Example**:
  ```swift
  let base64Images = convertUIImagesToBase64([uiImage]) // Returns ["data:image/jpeg;base64,..."]
  ```

### `convertUIImageToBase64(_:) -> String`
* **Purpose**: Converts a single `UIImage` to a base64-encoded string with MIME type.
* **Parameters**:
  * `image: UIImage`: The image to convert.
* **Features**:
  * Uses JPEG compression with quality 0.8.
  * Returns a base64 string prefixed with `data:image/jpeg;base64,`.
  * Returns an empty string if conversion fails.
* **Usage Example**:
  ```swift
  let base64String = convertUIImageToBase64(uiImage) // Returns "data:image/jpeg;base64,..."
  ```

### `convertBase64ToUIImage(_:) -> UIImage?`
* **Purpose**: Converts a base64-encoded string to a `UIImage`.
* **Parameters**:
  * `base64: String`: The base64-encoded image data (with or without MIME prefix).
* **Features**:
  * Decodes the base64 string to `Data` and creates a `UIImage`.
  * Returns `nil` if decoding fails.
* **Usage Example**:
  ```swift
  let image = convertBase64ToUIImage("data:image/jpeg;base64,...") // Returns UIImage or nil
  ```

---

## Data Flow

```mermaid
graph TD
View -->|sourceType, selectedImage| ImagePicker
ImagePicker -->|UIImagePickerController| Coordinator
Coordinator -->|selectedImage| View
View -->|UIImage| convertUIImagesToBase64
View -->|UIImage| convertUIImageToBase64
View -->|base64| convertBase64ToUIImage
convertUIImagesToBase64 -->|[String]| View
convertUIImageToBase64 -->|String| View
convertBase64ToUIImage -->|UIImage?| View
```

---

## Integration

* **Usage**: 
  * `ImagePicker` is used in SwiftUI views to allow users to select images (e.g., for property pictures or damage reports).
  * The conversion functions are used to prepare images for API calls (e.g., `updatePropertyPicture` or `createDamage` in `PropertyManagementService`) or to display API-fetched images.
* **Context**:
  * `ImagePicker` integrates with the camera or photo library, updating a bound `UIImage` variable for use in views like `PropertyDetailView`.
  * Conversion functions handle base64 encoding/decoding for API payloads and UI display.
* **Assumptions**:
  * The app requires camera/photo library permissions, configured in `Info.plist`.
  * Base64 conversion is used for API compatibility (e.g., `PropertyPictureResponse` or `DamageRequest`).
  * The `selectedImage` binding is used to trigger UI updates or API calls after image selection.

---

## Helper Features

* **Thread Safety**:
  * Image picking and conversion occur on the main thread, as they update UI or interact with `UIImagePickerController`.
* **Error Handling**:
  * Conversion functions handle failures gracefully (`nil` or empty string).
  * No explicit error feedback for failed image picking.
* **Performance**:
  * JPEG compression (quality 0.8) balances file size and quality for API uploads.
  * No caching in conversion functions; relies on `ImageCache` for performance.
