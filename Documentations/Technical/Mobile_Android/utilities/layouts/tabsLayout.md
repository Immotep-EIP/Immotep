# TabsLayout

---

## UI Layout

### `TabsLayout`

* Generic tabbed interface layout component.
* Displays a horizontal tab bar with dynamic content underneath.
* Highlights the selected tab and allows switching via callback.

---

## Parameters

* `tabIndex: Int`
  Index of the currently selected tab.

* `tabs: List<String>`
  List of tab titles (e.g., `["Info", "Photos", "History"]`).

* `setTabIndex: (Int) -> Unit`
  Callback that updates the selected tab index on user interaction.

* `content: @Composable () -> Unit`
  Dynamic composable content rendered below the tab row.
  Varies based on the selected tab.

---

## Behavior

* Uses `TabRow` from Material3 with:

  * Full-width layout (`Modifier.fillMaxWidth()`).
  * Custom tab indicator (`SecondaryIndicator`) using `colorScheme.secondary`.

* Tabs are created from the `tabs` list:

  * Each tab is selectable.
  * Active tab is styled using `selectedContentColor`.
  * Inactive tabs use `onBackground` color.

* `content()` is rendered beneath the tabs and updates as needed.

---

## Tags for Testing

* Each tab gets a unique test tag:

  * `"tab 0"` for the first,
  * `"tab 1"` for the second, etc.

---

## Styling

* Background color for the tab bar comes from `MaterialTheme.colorScheme.background`.
* The selected tab's indicator and text color use `colorScheme.secondary`.
* Unselected tabs use `colorScheme.onBackground`.

---

## Notes

* Reusable for any screen needing tabbed navigation.
* Ideal for breaking up content into logical sections.
* Clean integration with Compose’s state-driven UI model.
