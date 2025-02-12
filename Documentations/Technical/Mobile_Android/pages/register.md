## Register Page Documentation

This section provides a detailed overview of the Register page functionality within the Immotep mobile application.

### Components

* **Header:** Displays the application logo or name. (Implemented using `Header` composable)
* **Top Text:** Shows the "Create Account" header and additional instructions. (Implemented using `TopText` composable)
* **Error Alert:** Displays any error messages related to API responses or input validation. (Implemented using `ErrorAlert` composable)
* **Last Name Input:** A text field for users to enter their last name. (Implemented using `OutlinedTextField` composable with validation feedback if incorrect)
* **First Name Input:** A text field for users to enter their first name. (Implemented using `OutlinedTextField` composable with validation feedback if incorrect)
* **Email Input:** A text field for users to enter their email address. (Implemented using `OutlinedTextField` composable with `KeyboardOptions.keyboardType = KeyboardType.Email` and validation feedback if incorrect)
* **Password Input:** A password field for users to enter their chosen password. (Implemented using `PasswordInput` composable with validation feedback if incorrect)
* **Confirm Password Input:** A password field for users to re-enter and confirm their password. (Implemented using `PasswordInput` composable with validation feedback if passwords do not match)
* **Agree to Terms Checkbox:** An option for users to agree to the terms and conditions before registration. (Implemented using `CheckBoxWithLabel` composable)
* **Sign Up Button:** Triggers the registration process when clicked. (Implemented using a `Button` composable)
* **Already Have an Account Text:** A clickable text leading to the login page for users who already have an account. (Implemented using a clickable `Text` composable)

### Functionalities

* Users can input their last name, first name, email address, password, and confirm password fields.
* The `Last Name Input` and `First Name Input` enforce minimum and maximum length requirements.
* The `Email Input` validates the email format.
* The `Password Input` enforces a minimum password length.
* The `Confirm Password Input` checks if the confirmed password matches the initially entered password.
* The `Agree to Terms Checkbox` must be selected to enable registration.
* Clicking the `Sign Up Button` triggers the registration process, performing the following actions:
    * Validates all input fields for required formats and conditions.
    * Calls the `ApiClient` service in the ViewModel to register the user.
    * Upon successful registration, navigates to the login page.
    * In case of registration failure, displays relevant error messages based on the API error codes.

### Data Flow

* User interaction with input fields updates the registration form state (`RegisterForm` and `RegisterConfirm`) in the `RegisterViewModel`.
* The sign-up button triggers the `onSubmit` function in the ViewModel.
* The `onSubmit` function validates the user input, checks if all fields are correctly filled, and initiates registration through the API.
* The `ApiClient` service handles the registration logic and responds with success or error codes.
* Based on the outcome, the ViewModel updates either the registration state or the error state.
* The UI layer displays relevant information based on the updated state (registration success or error messages).

### Navigation

* This page uses navigation components to transition between screens:
    * Successful registration navigates to the login page.
    * Clicking "Already Have an Account" navigates to the login page.

### ViewModel

* The `RegisterViewModel` manages the registration state (`RegisterForm`), including last name, first name, email, password, and confirm password information.
* It also manages the confirmation state (`RegisterConfirm`) for fields like the password confirmation and terms agreement.
* The `RegisterFormError` state holds the error flags for each input field and the API error code.
* The `onSubmit` function:
    * Validates each input field.
    * Checks if the confirmed password matches and if the terms agreement checkbox is selected.
    * Calls the `registerToApi` function, which interacts with `ApiClient` to register the user and navigates to the login page upon success.
* In case of failure, `registerToApi` catches exceptions and sets appropriate error messages based on API responses.