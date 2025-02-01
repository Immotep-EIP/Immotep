## Header Component Documentation

This section provides an overview of the `Header` component in the Immotep mobile application.

### Overview

The `Header` composable component is a simple, styled header that displays the application logo and name side-by-side. It is typically used at the top of screens to reinforce the app’s branding.

### Components

* **Logo Image**: An `Image` composable displaying the Immotep logo (`immotep_png_logo`).
* **App Name Text**: A `Text` composable displaying the application name, styled with a large font size and primary color.

### Properties

* **Logo (`Image`)**:
  - **`painterResource(id = R.drawable.immotep_png_logo)`**: Loads the app logo from resources.
  - **`contentDescription`**: Uses a localized description of the logo, provided by `stringResource(R.string.immotep_logo_desc)`.
  - **`Modifier.size(50.dp)`**: Sets the image size to 50x50 dp.
  - **`Modifier.padding(end = 10.dp)`**: Adds padding to the right of the image to separate it visually from the text.

* **App Name (`Text`)**:
  - **`stringResource(R.string.app_name)`**: Loads the application name from resources for localization.
  - **`fontSize = 30.sp`**: Sets the font size to 30 sp for emphasis.
  - **`color = MaterialTheme.colorScheme.primary`**: Uses the primary color from the theme to style the text.

### Usage

The `Header` component is designed for use at the top of app screens to provide consistent branding across the application.

### Example Usage

```kotlin
Header()
```

### Visual Layout

1. **Row Layout**: Arranges the logo and app name horizontally in a row.
2. **Spacing and Alignment**: The image and text are centered vertically for alignment, with padding separating the two elements for readability.

### Accessibility

* **`contentDescription`** for the logo image makes the component accessible by providing a description for screen readers.
  
### Tags

* **`Modifier.testTag("header")`**: Adds a test tag to the component for UI testing, allowing automated tests to locate the header component easily.