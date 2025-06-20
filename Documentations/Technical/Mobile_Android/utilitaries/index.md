# Project Utilities Overview

---

This page presents a high-level overview of the core utility building blocks in the project. It covers the **UI Components**, **Layouts & Themes**, and **Utility Classes** that are designed and styled for consistent reuse across the app.

---

## Components

The project contains a rich set of **feature-specific components** organized in modules and files to manage distinct UI parts and functionalities. These components complement the base UI components and layouts, handling complex UI logic and interactions.

### Main Components Tree:

* **addDamageModal**

  * `AddDamageModal.kt`
  * `AddDamageModalViewModel.kt`
* **AddingPicturesCarousel.kt**
* **addOrEditPropertyModal**

  * `AddOrEditPropertyModal.kt`
  * `AddOrEditPropertyViewModel.kt`
* **CheckBoxWithLabel.kt**
* **DeletePopUp.kt**
* **ErrorAlert.kt**
* **FullScreenZoomableImage.kt**
* **Header.kt**
* **InitialFadeIn.kt**
* **InternalLoading.kt**
* **inventory**

  * `AddRoomOrDetailModal.kt`
  * `EditRoomOrDetailModal.kt`
  * `NextInventoryButton.kt`
* **InventoryCenterAddButton.kt**
* **inviteTenantModal**

  * `InviteTenantModal.kt`
  * `InviteTenantViewModel.kt`
* **LoadingDialog.kt**
* **LoggedBottomBar.kt**
* **LoggedTopBar.kt**
* **TakePhotoButton.kt**
* **TopText.kt**

These components encapsulate complex UI behaviors like modals, dialogs, image carousels, loading states, and top/bottom navigation bars. Each may be paired with its own ViewModel to handle state and business logic separately.

---

## UI Components

The project contains a set of **Base UI Components** that have been customized to fit the app's branding and style guidelines. These components simplify UI development by providing consistent behavior, styling, and accessibility support.

### Key Components Include:

* **BackButton**: A styled icon button to navigate backward.
* **DateRangeInput**: A date picker input field with error handling and label support.
* **DropDown**: A customizable dropdown menu with error display and item selection logic.
* **PasswordInput**: Secure input field with toggleable password visibility.
* **StyledButton**: Buttons with loading states, error colors, and consistent padding and shapes.
* **OutlinedTextField**: Text fields with helper messages, error states, and icon support.

Each component uses **Material3** theming and is designed to integrate seamlessly with Compose’s declarative UI system.

---

## Layouts & Themes

Layouts manage common screen structures, promoting a uniform look and feel:

* **BigModalLayout**: A full-screen modal sheet with customizable height and smooth open/close animations.
* **DashBoardLayout**: A top-bar, content area, and bottom-bar arrangement tailored for logged-in dashboard screens.
* **InventoryLayout** and **InventoryTopBar**: Structured layouts with custom headers and exit buttons for inventory-related screens.

### Theme System

The app’s **Theme** is a core utility:

* Supports **light**, **dark**, and **contrast-adjusted** color schemes.
* Defines comprehensive color palettes with semantic color roles.
* Applies consistent typography styles.
* Automatically adapts to system theme settings for better user experience.

This theming system ensures visual coherence and accessibility throughout the app.

---

## Utility Classes

Utility classes provide essential non-UI functionalities widely used across the project:

* **DateFormatter**: Handles date formatting and conversions, supporting modern `OffsetDateTime` and legacy formats with fallback handling.
* **PdfsUtils**: Simplifies opening PDF files using Android’s native intent system with proper URI permissions.
* **RegexUtils**: Provides validation logic, such as email format validation, ensuring user input correctness.
* **ThemeUtils**: Helps select theme-dependent resources like icons based on the current app theme.

These utilities enhance code reuse, maintainability, and keep core logic centralized.

---
