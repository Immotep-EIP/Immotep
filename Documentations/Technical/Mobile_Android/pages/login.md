## Login Page Documentation

This section provides a detailed overview of the Login page functionality within the Immotep mobile application.

**Components:**

* **Header:** Displays the application logo or name. (Implemented using `Header` composable)
* **Top Text:** Shows a welcome message and login instruction. (Implemented using `TopText` composable)
* **Email Input:** A text field for users to enter their email address. (Implemented using `OutlinedTextField` composable with `KeyboardOptions.keyboardType = KeyboardType.Email`)
* **Password Input:** A password field for users to enter their login password. (Implemented using `PasswordInput` composable)
* **Keep Signed Checkbox:** An option for users to stay signed in after login. (Implemented using `CheckBoxWithLabel` composable)
* **Forgot Password Text:** A clickable text leading to the forgot password functionality. (Implemented using a clickable `Text` composable)
* **Login Button:** Triggers the login process when clicked. (Implemented using a `Button` composable)
* **Sign Up Text:** A clickable text leading to the registration page. (Implemented using a clickable `Text` composable)

**Functionalities:**

* Users can enter their email address and password in the respective input fields.
* The `Email Input` enforces email format validation.
* Users can choose to stay signed in after login using the `Keep Signed Checkbox`.
* Clicking the `Forgot Password Text` navigates to the forgot password functionality.
* Clicking the `Login Button` triggers the login process, performing the following actions:
    * Validates email format and password length.
    * Calls the `AuthService` (injected through ViewModel) to perform login with the provided credentials.
    * Upon successful login, navigates to the dashboard screen.
    * In case of login failure, displays relevant error messages based on the received error code.

**Data Flow:**

* User interaction with input fields updates the login state (`LoginState`) in the ViewModel.
* The login button triggers the `login` function in the ViewModel.
* The `login` function validates user input and initiates the login process through `AuthService`.
* `AuthService` handles login logic and persists user data using DataStore (if `keepSigned` is selected).
* Based on the login outcome, the ViewModel updates either the login state or the error state.
* The UI layer displays relevant information based on the updated state (login success, error messages).

**Navigation:**

* This page utilizes navigation components to navigate between screens:
    * Clicking `Forgot Password Text` triggers navigation to the forgot password functionality.
    * Successful login navigates to the dashboard screen.
    * Clicking `Sign Up Text` navigates to the registration page.

**ViewModel:**

* The `LoginViewModel` manages the login state (`LoginState`), including email, password, and keep signed information.
* It also handles the error state (`LoginErrorState`), indicating validation errors or API error codes.
* The `updateEmailAndPassword` function allows updating the login state based on user input.
* The `login` function performs input validation, calls `AuthService` for login, and handles success/failure scenarios.

