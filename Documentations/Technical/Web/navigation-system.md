# Navigation System

## Overview

The navigation system in our application is built using **React Router**. To simplify and centralize navigation logic, we use a custom hook called `useNavigation`. This hook provides pre-defined functions for navigating between pages, ensuring consistency and maintainability.

---

## Components and Structure

### **`useNavigation` Hook**

The `useNavigation` hook is a custom React hook that wraps the `useNavigate` function from React Router. It provides a set of reusable navigation methods for different pages in the application.

### **`NavigationEnum`**

We use an enumeration (`NavigationEnum`) to store all route paths as constants. This prevents hardcoding paths throughout the codebase and makes the system easier to maintain.

---

## Usage

### Importing and Using the Hook

To use the navigation functionality, import the `useNavigation` hook into your component.

#### Example: Navigating to Login Page

```typescript
import useNavigation from "@/hooks/useNavigation";

function ExampleComponent() {
  const { goToLogin } = useNavigation();

  return <button onClick={goToLogin}>Go to Login</button>;
}
```
