## TopText Component Documentation

The `TopText` component in the Immotep application is a customizable text display for a title and subtitle, designed to offer flexible top padding and centered alignment. It is primarily used to create a heading section with variable spacing based on different UI contexts.

### Overview

`TopText` presents a title and subtitle in a vertically centered column, with options to adjust the top margin. The component's styling aligns with the app's primary color scheme and is designed to ensure consistent text sizing and spacing.

### Properties

* **Parameters**:
  - **`title`**: (String) The main title text to display, styled with a larger font size and semi-bold weight.
  - **`subtitle`**: (String) The subtitle text that appears below the title, styled with a smaller font size.
  - **`limitMarginTop`**: (Boolean) If true, applies a top padding of 15.dp.
  - **`noMarginTop`**: (Boolean) If true, removes top padding entirely.

* **Modifiers**:
  - **`fillMaxWidth`**: Ensures the component spans the full width of the screen.
  - **`padding`**:
    - Controls the top margin based on `limitMarginTop` and `noMarginTop` flags.
    - Default top padding is 90.dp if both flags are false.

### Functionality

* **Top Padding Control**:
  - The component’s top padding can vary:
    - **`noMarginTop = true`**: No top margin applied.
    - **`limitMarginTop = true`**: Applies a limited top margin of 15.dp.
    - **Both flags false**: Applies a default top margin of 90.dp.
* **Text Styling**:
  - **Title**: Displayed with a font size of 30.sp, semi-bold weight, and the app's primary color.
  - **Subtitle**: Displayed below the title with a font size of 15.sp, using the same primary color.

### Usage

To use `TopText`, pass the required `title` and `subtitle` strings, and set the margin control flags as needed:

```kotlin
TopText(
    title = "Welcome",
    subtitle = "Your personalized dashboard",
    limitMarginTop = true
)
```

### Layout

1. **Centered Column Layout**:
   - **Title**: Displayed at the top, styled for emphasis.
   - **Subtitle**: Positioned directly below the title, providing additional context.
   - Centered horizontally to create a balanced and prominent header.

### Visual Customization

* **Text Colors**:
  - Both title and subtitle use the primary color of the app's theme, allowing for consistency across different themes.

### Accessibility

* **Considerations**:
  - **Descriptive Text**: The title and subtitle should be self-explanatory and relevant to the screen's purpose to enhance accessibility.

### Example

```kotlin
TopText(
    title = "Account Settings",
    subtitle = "Manage your personal information and preferences",
    noMarginTop = false
)
```

This example would render a top heading with the default 90.dp top margin and app-styled title and subtitle texts.