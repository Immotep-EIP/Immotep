## LoggedBottomBar Component Documentation

This section provides an overview of the `LoggedBottomBar` and `LoggedBottomBarElement` components in the Immotep mobile application.

### Overview

The `LoggedBottomBar` component serves as a customizable bottom navigation bar for logged-in users. It allows navigation between different sections of the app using labeled icons. Each icon is highlighted when its associated page is active, helping users quickly identify their current location within the app.

### Components

* **`LoggedBottomBarElement`**: A single item in the bottom bar, representing a navigation option.
* **`LoggedBottomBar`**: The container for multiple `LoggedBottomBarElement` items, displaying them in a row with separation and a bottom line for clarity.

### Properties

#### `LoggedBottomBarElement`

* **Parameters**:
  - **`navController`**: The navigation controller responsible for handling page changes.
  - **`name`** (`String`): The displayed name of the navigation option (localized).
  - **`icon`** (`ImageVector`): The icon image representing the page.
  - **`iconDescription`** (`String`): The description of the icon for accessibility.
  - **`pageName`** (`String`): The route name of the associated page.

* **Selected State Styling**:
  - The component checks if the current route matches `pageName`.
  - If selected, a magenta line is drawn below the item to indicate the active page.

* **Visuals**:
  - **`Icon`**: Displays the `icon` with `contentDescription` for accessibility.
  - **`Text`**: Displays the `name` below the icon.

#### `LoggedBottomBar`

* **Parameters**:
  - **`navController`**: Passed down to each `LoggedBottomBarElement` for navigation control.

* **Divider Styling**:
  - A light gray line is drawn across the width of the bar as a visual separator from the content above.

* **Row Layout**:
  - Arranges each `LoggedBottomBarElement` horizontally, with spacing between elements and centered alignment.
  
* **Icons and Labels**:
  - Uses icons for home, real property, messages, and settings with corresponding page names and descriptions for accessibility.

### Usage

The `LoggedBottomBar` component is meant to be placed at the bottom of a screen, providing users with a way to switch between core app sections.

### Example Usage

```kotlin
LoggedBottomBar(navController)
```

### Visual Layout

1. **Spacer Divider**: A thin horizontal line separates the bar from the content above.
2. **Row Layout**: `LoggedBottomBarElement` items are arranged in a row, spaced equally.
3. **Selected Indicator**: A magenta line below the active item shows the current page visually.

### Accessibility

* **`contentDescription`** for each icon provides information for screen readers.
  
### Tags

* **`Modifier.testTag("header")`**: Test tag is not defined here but can be added if needed for testing individual elements.