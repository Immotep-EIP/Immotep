# Global Overview of the iOS Mobile Application

## Technologies Used

This iOS application is developed natively using [SwiftUI 6](https://developer.apple.com/documentation/swiftui/) to provide a reactive and declarative user interface. We use the [MVVM (Model-View-ViewModel)](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93viewmodel) architecture to ensure a clear separation between the UI and data management, as well as application logic, promoting a clean and maintainable codebase. Development and testing of the application are performed with [Xcode](https://developer.apple.com/xcode/), while external dependencies are managed through the [Swift Package Manager](https://swift.org/package-manager/).

### Versions of Technologies Used

- [Swift](https://swift.org/): 6.0
- [SwiftUI](https://developer.apple.com/documentation/swiftui/): 6.0
- iOS SDK minimum version: 18.0
- iOS SDK target version: 18.0
- [Xcode](https://developer.apple.com/xcode/): 16.0

## Pages Hierarchy

All the views are located in the **`Sources/`** folder. Each view is organized into separate folders based on the feature it represents. In accordance with the [MVVM](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93viewmodel) architecture, each feature folder contains at least two files: a UI file and a ViewModel file (ending with `ViewModel.swift`).

- The UI file contains the SwiftUI code that manages user interactions and visual elements.
- The ViewModel file is responsible for the logic, state management, and data fetching related to the UI. API calls and state variables are defined within this file.


### Folder Structure

The project's codebase is organized into a clear and modular structure, facilitating maintainability and scalability. Here's an overview of the key directories:

- **Immotep/**: The root directory for the entire project.
  - **.swiftlint/**: Configuration for SwiftLint.
  - **ImmotepTests/**: Contains unit tests for the application.
  - **ImmotepUITests/**: Contains UI tests for the application.
  - **Resources/**: 
    - **Assets/**: Contains various asset files.
  - **Views/**: Contains different views of the application, organized by features.

This structure promotes a distinct separation between UI and application logic, enhancing the clarity and modularity of the codebase.


## Testing

#### Code Quality
To maintain code quality, we use [SwiftLint](https://github.com/realm/SwiftLint), a tool that enforces Swift style and conventions. SwiftLint helps catch common programming errors and ensures consistency across the codebase.

#### Unit Testing
We use [XCTest](https://developer.apple.com/documentation/xctest) for unit testing, focusing on validating individual functions and methods to ensure their accuracy. The unit tests are organized within the **`ImmotepTests/`** folder, which follows the same structure as the main application.


#### UI Testing
For UI testing, we also use XCTest to simulate user interactions and validate that the UI behaves as expected. UI tests are located in the **`ImmotepUITests/`** folder.


You can run the unit tests through Xcode's built-in test navigator or execute them directly from the terminal using the appropriate command :

```
xcodebuild test -scheme Immotep -destination 'platform=iOS Simulator,name=iPhone 16' [-only-testing:TestName]
```

