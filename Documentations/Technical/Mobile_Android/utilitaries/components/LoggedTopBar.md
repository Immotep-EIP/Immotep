## LoggedTopBar Component Documentation

The `LoggedTopBar` component in the Immotep application is a customizable top bar for logged-in users, offering navigation and logout functionality. It provides an icon-based navigation to the user profile and a logo that triggers a logout on click.

### Overview

The `LoggedTopBar` is designed to be placed at the top of the screen, providing a consistent header with branding and user account navigation. The component also includes a logout function triggered by clicking the app logo, which clears user data and navigates to the login screen.

### Components

* **`LoggedTopBarViewModel`**: Handles user logout functionality by interfacing with `AuthService`.
* **`LoggedTopBar`**: The composable UI component displaying the app logo, title, and profile icon.

### Properties

#### `LoggedTopBarViewModel`

* **`logout` function**: Initiates a logout sequence using `AuthService` and navigates the user to the login screen.

#### `LoggedTopBar`

* **Parameters**:
  - **`navController`**: Manages navigation for logout and profile routes.

* **Visuals**:
  - **Logo Image**: Displays the app logo, clickable to initiate logout.
  - **App Title**: Shows the app’s name with a primary color and font weight styling.
  - **Profile Icon**: Leads to the user’s profile, allowing users to navigate to or from their account page.

* **Divider Styling**:
  - A light gray line is drawn at the bottom edge of the top bar for separation.

### Functionality

* **Logo Click**: The logo functions as a logout button. When clicked, it calls the `logout` function in `LoggedTopBarViewModel` to clear user data and navigate back to the login screen.
* **Profile Icon Button**:
  - Navigates to the profile page if the current page is not "profile."
  - If on the profile page, pressing the icon pops the back stack, returning to the previous screen.

### Usage

To integrate the `LoggedTopBar` in a composable layout, pass the navigation controller as a parameter:

```kotlin
LoggedTopBar(navController)
```

### Visual Layout

1. **Row Layout**: The layout arranges the logo, app title, and profile icon horizontally, with spacing for visual clarity.
2. **Logo with Logout**: Positioned to the left, the clickable logo triggers the logout function.
3. **App Title**: Center-aligned, displaying the app’s name.
4. **Profile Icon**: Positioned to the right, allowing easy access to the profile.

### Accessibility

* **Content Descriptions**:
  - **Logo**: A description based on `R.string.immotep_logo_desc`.
  - **Profile Icon**: Described as "Account circle, go back to the login page."

### Tags

* **`Modifier.testTag("header")`**: Assigns a test tag for UI testing purposes.