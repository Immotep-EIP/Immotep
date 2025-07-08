# Keyz Utilities Overview

---

This page provides a high-level overview of the core utility building blocks in the Keyz app. It covers **Application Lifecycle Management**, **Localization**, **Caching**, **Input Handling**, **Image Processing**, **Date Formatting**, and **Theming** utilities designed for consistent reuse across the app.

---

## Application Lifecycle Management

The app includes utilities to manage lifecycle events and authentication state, ensuring proper session handling.

* **AppDelegate**
  * `AppDelegate.md`
  * Manages iOS application lifecycle events (e.g., launch, active, terminate) and integrates with `UserDefaults` and `TokenStorage` to handle the `isLoggedIn` state. Used to control login persistence when the app becomes active or terminates.

---

## Localization

Localization utilities enable multilingual support by managing language settings and string localization.

* **Bundle Extensions**
  * `Bundle.md`
  * Extends `String` for localization using `.lproj` files and `Bundle` to set the app’s language. Integrates with `UserDefaults` to persist language preferences and is used in UI components for displaying localized text.

---

## Caching

Caching utilities optimize performance by storing frequently accessed data (images and documents) in memory.

* **DocumentCache**
  * `DocumentCache.md`
  * A thread-safe singleton for caching `PDFDocument` objects, used for lease agreements or inventory reports. Reduces redundant loading from disk or network, likely integrated with `PropertyManagementService` for document fetching.

* **ImageCache**
  * `ImageCache.md`
  * A thread-safe singleton for caching `UIImage` objects, used for property pictures or damage report images. Optimizes performance for UI rendering, integrated with API calls like `fetchPropertiesPicture`.

---

## Input Handling

Input handling utilities manage user interactions to prevent redundant or rapid executions.

* **Debouncer**
  * `Debouncer.md`
  * A utility class for debouncing actions (e.g., search queries or button taps) by delaying execution. Used in view models or UI components to throttle rapid inputs, improving performance and reducing API load.

---

## Image Processing

Image processing utilities handle image selection and conversion for API compatibility and UI display.

* **ImagePicker**
  * `ImagePicker.md`
  * A SwiftUI wrapper for `UIImagePickerController` to select images from the camera or photo library, with utility functions for converting `UIImage` to/from base64 strings. Used in views like `PropertyDetailView` for uploading property or damage images.

---

## Date Formatting

Date formatting utilities handle date conversions for consistent display and processing.

* **Format**
  * `FormatDate.md`
  * Provides functions to convert date strings to `Date` objects and reformat ISO 8601 dates to `dd/MM/yyyy`. Used in views or view models (e.g., `PropertyDetailView` or `InventoryViewModel`) for handling lease or inventory timestamps.

---

## Theming

Theming utilities manage the app’s appearance, supporting light, dark, and system themes.

* **ThemeManager**
  * `ThemeManager.md`
  * Defines a `ThemeOption` enum and `ThemeManager` struct to apply light, dark, or system themes by updating the UI style. Integrated with settings views to allow users to select themes, with localized theme names.

---

## Utility Integration

* **Dependencies**:
  * Most utilities interact with `UserDefaults` for state persistence (e.g., `isLoggedIn`, `lang`, theme settings).
  * `DocumentCache` and `ImageCache` are likely used with `PropertyManagementService` for efficient data handling.
  * `ImagePicker` and `Format` support API payloads and UI display (e.g., base64 images, formatted dates).
* **Context**:
  * These utilities are designed for reuse across SwiftUI views and view models, ensuring consistent behavior for localization, caching, input handling, and theming.
  * `AppDelegate` ties into the app’s lifecycle, while others support specific features like property management or user settings.
