## DashBoardLayout Component Documentation

The `DashBoardLayout` component is a flexible layout in the Immotep application designed to organize content with a consistent header and footer, creating a standard structure for dashboard-like screens. It combines a top bar, main content area, and a bottom bar for seamless navigation and interaction.

### Overview

The `DashBoardLayout` component structures a screen by displaying a fixed top bar, a flexible content area, and a bottom bar, enhancing user experience with consistent and easy navigation. This layout is particularly useful for screens where dashboard or main navigation elements need to be displayed in a uniform format.

### Properties

* **Parameters**:
  - **`navController`**: (`NavController`) The navigation controller used for navigating between screens.
  - **`testTag`**: (String) A tag to identify the layout in testing.
  - **`content`**: (`@Composable () -> Unit`) A composable function that represents the main content area. This content is displayed between the top and bottom bars.

* **Modifiers**:
  - **`testTag`**: Sets a test tag for the layout, helpful for UI testing.
  - **`padding`**: Adds 2.dp padding around the main content area to provide spacing within the layout.

### Structure

1. **Top Bar**:
   - Uses `LoggedTopBar`, a composable that serves as the top navigation bar.
   - This bar remains fixed at the top of the screen, providing consistent access to key actions (like logout or profile navigation).

2. **Content Area**:
   - A flexible area within the layout, determined by the `content` parameter.
   - Uses `weight(1f)` to make the area expandable, filling any remaining vertical space between the top and bottom bars.
   - **Padding**: 2.dp padding around the content area for spacing.

3. **Bottom Bar**:
   - Uses `LoggedBottomBar`, a composable navigation bar that remains fixed at the bottom of the screen.
   - Displays core navigation actions, enabling users to easily move between main sections.

### Functionality

* **Layout Control**:
  - The top and bottom bars remain fixed, while the main content area flexibly expands to fill the vertical space.
  - The layout makes it easy to display different content sections within a consistent structure.
* **Navigation**:
  - The `navController` parameter connects `LoggedTopBar` and `LoggedBottomBar` to facilitate navigation within the app.

### Usage

To use `DashBoardLayout`, provide the required `navController`, `testTag`, and any custom `content` for the main display area:

```kotlin
DashBoardLayout(
    navController = navController,
    testTag = "dashboardScreen",
    content = {
        Text("Welcome to the Dashboard")
        Button(onClick = { /* Action here */ }) {
            Text("Click Me")
        }
    }
)
```

### Visual Customization

* **Padding and Spacing**:
  - The layout includes 2.dp padding around the main content area, adding light spacing to improve readability.

### Testing

* **Test Tag**:
  - The `testTag` parameter allows for easy targeting in automated UI tests. By setting a unique tag for each screen, you can verify layout presence and behavior with testing frameworks.

### Example

```kotlin
DashBoardLayout(
    navController = navController,
    testTag = "profileScreen",
    content = {
        Text("Profile Information")
    }
)
```

This example renders a layout with the top and bottom navigation bars, and a centered main content area showing the text “Profile Information.” The screen can be identified in tests by the `testTag` "profileScreen".