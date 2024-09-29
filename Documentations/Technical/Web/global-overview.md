# Global overview of the Web Application

## Technologies used

This web application is built with [React](https://react.dev/), using [TypeScript](https://www.typescriptlang.org/) to ensure static typing and improve code robustness. The project leverages [Vite](https://vitejs.dev/) as its bundling tool for faster development and optimized production builds.

To ensure high code quality, ESLint is configured with the Airbnb style guide to enforce best practices. Additionally, Husky is integrated to manage Git hooks, including a pre-commit that runs lint-staged. This ensures that ESLint only checks staged files, improving performance during commits. Prettier is also set up to work alongside ESLint, ensuring consistent code formatting across the entire project.

### Stack and Tools

- React version: 18.3.1
- TypeScript version: 5.5.4
- Vite version: 5.4.7

## Code Quality

- ESLint version: 8.57.1
- Husky version: 4.3.8
- Lint-Staged version: 15.2.10

## Code Structure

The project's codebase is organized into a clear and modular structure, facilitating maintainability and scalability. Here's an overview of the key directories and their contents:

- **src/**: The root directory for all source files.

  - **assets/**: Contains various asset files used throughout the application.

  - **components/**: This directory holds reusable components that can be utilized across the entire application.

  - **views/**: Contains different views of the application, organized in appropriate folders, with accompanying CSS files for styling.

  - **enums/**: This folder includes TypeScript enums used throughout the application, providing a way to define named constants.

  - **interfaces/**: Contains TypeScript interfaces that define the structure of various data types used in the application.

  - **hooks/**: Houses custom hooks used in the application, helping to avoid cluttering views with hook logic.

  - **services/**: Contains files that manage API calls and business logic, separating concerns from UI components.

  - **utils/**: A collection of utility functions that aid in the development of the application.

  - **context/**: This directory holds context files that can be shared across the application, enabling state management.

This organized structure promotes clarity, making it easier to navigate and maintain the codebase as the application grows.

### Unit testing
