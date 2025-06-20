# LoggedBottomBar

---

## UI Component

### `LoggedBottomBar`

* Bottom navigation bar for logged-in users, providing quick access to main app sections.

---

## Data Model

### `NavItem`

* Represents each navigation tab with:

  * `name`: Display label (localized string).
  * `icon`: Icon vector from Material icons.
  * `iconDescription`: Accessibility description for the icon.
  * `pageName`: Navigation route string.

---

## Behavior & Styling

* Defines a list of 4 `NavItem`s:

  1. Dashboard (Home icon)
  2. Real Property (HomeWork icon)
  3. Messages (MailOutline icon)
  4. Profile/Settings (Settings icon)

* Uses `NavigationBar` composable with:

  * Background color from the current Material theme.
  * Test tag `"loggedBottomBar"` for UI testing.

* For each navigation item, renders a `NavigationBarItem`:

  * Selected state determined by matching current navigation route.
  * Clicking triggers navigation to the associated page.
  * Shows the icon with accessibility description.
  * Displays label text, colored:

    * Theme’s `onPrimaryContainer` if selected.
    * Gray otherwise.
  * Uses themed colors for selected/unselected icons and indicator.
  * Each item has its own test tag with page name.

---

## Usage Example

```kotlin
LoggedBottomBar(navController = navController)
```

---

## Notes

* Integrates with `NavController` to handle navigation state and actions.
* Accessibility-conscious by providing icon descriptions.
* Uses consistent Material 3 theming for colors and styles.
* Test tags help UI automation identify components and elements.
