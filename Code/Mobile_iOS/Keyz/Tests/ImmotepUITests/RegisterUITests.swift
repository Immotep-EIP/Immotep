//
//  RegisterUITests.swift
//  ImmotepTests
//
//  Created by Liebenguth Alessio on 12/10/2024.
//

import XCTest

final class RegisterUITests: XCTestCase {
    let nameFr = "Entrez votre nom"
    let nameEn = "Enter your name"
    let firstNameFr = "Entrez votre prénom"
    let firstNameEn = "Enter your first name"
    let emailFr = "Entrez votre email"
    let emailEn = "Enter your email"
    let passwordFr = "Entrez votre mot de passe"
    let passwordEn = "Enter your password"
    let confirmPasswordFr = "Confirmez votre mot de passe"
    let confirmPasswordEn = "Enter your password confirmation"

    let app = XCUIApplication()

    override func setUpWithError() throws {
        continueAfterFailure = false
        XCUIDevice.shared.orientation = .portrait
        app.launch()
    }

    override func tearDownWithError() throws {
    }

    func navigateToRegisterView() {
        if app.buttons["person.crop.circle.fill"].exists {
            app.buttons["person.crop.circle.fill"].tap()

            let logoutButton = app.buttons["Logout"].exists || app.buttons["Se déconnecter"].exists
            XCTAssertTrue(logoutButton)

            if app.buttons["Logout"].exists {
                app.buttons["Logout"].tap()
            } else if app.buttons["Se déconnecter"].exists {
                app.buttons["Se déconnecter"].tap()
            }
            let signUpButton = app.buttons["signUpLink"]
            signUpButton.tap()
        } else {
            let signUpButton = app.buttons["signUpLink"]
            signUpButton.tap()
        }
    }

    func testWelcomeTextExists() throws {
        navigateToRegisterView()
        let welcomeText = app.staticTexts["Create your account"].exists || app.staticTexts["Créer un compte"].exists
        XCTAssertTrue(welcomeText)
    }

    func testSecondaryWelcomeTextExists() throws {
        navigateToRegisterView()
        let joinTextFr = "Rejoignez Immotep pour votre tranquillité d’esprit !"
        let joinTextEn = "Join Immotep for your peace of mind!"
        let secWelcomeText = app.staticTexts[joinTextEn].exists || app.staticTexts[joinTextFr].exists
        XCTAssertTrue(secWelcomeText)
    }

    func testRequiredFieldsExist() throws {
        navigateToRegisterView()
        let nameText = app.staticTexts["Name*"].exists || app.staticTexts["Nom*"].exists
        let firstNameText = app.staticTexts["First name*"].exists || app.staticTexts["Prénom*"].exists
        let emailText = app.staticTexts["Email*"].exists || app.staticTexts["Email*"].exists
        let passwordSecure = app.staticTexts["Password*"].exists || app.staticTexts["Mot de passe*"].exists
        let confirmPasswordSecure = app.staticTexts["Password confirmation*"].exists || app.staticTexts["Confirmation du mot de passe*"].exists

        XCTAssertTrue(nameText)
        XCTAssertTrue(firstNameText)
        XCTAssertTrue(emailText)
        XCTAssertTrue(passwordSecure)
        XCTAssertTrue(confirmPasswordSecure)
    }

    func testTextFieldsExist() throws {
        navigateToRegisterView()
        let nameTextField = app.textFields[nameEn].exists ? app.textFields[nameEn] : app.textFields[nameFr]
        let firstNameTextField = app.textFields[firstNameEn].exists ? app.textFields[firstNameEn] : app.textFields[firstNameFr]
        let emailTextField = app.textFields[emailEn].exists ? app.textFields[emailEn] : app.textFields[emailFr]
        let passwordSecureField = app.secureTextFields[passwordEn].exists ? app.secureTextFields[passwordEn] : app.secureTextFields[passwordFr]
        let confirmPasswordSecureField =
        app.secureTextFields[confirmPasswordEn].exists ? app.secureTextFields[confirmPasswordEn] : app.secureTextFields[confirmPasswordFr]

        XCTAssertTrue(nameTextField.exists)
        XCTAssertTrue(firstNameTextField.exists)
        XCTAssertTrue(emailTextField.exists)
        XCTAssertTrue(passwordSecureField.exists)
        XCTAssertTrue(confirmPasswordSecureField.exists)
    }

    func testFillingInFields() throws {
        navigateToRegisterView()
        let nameTextField = app.textFields[nameEn].exists ? app.textFields[nameEn] : app.textFields[nameFr]
        let firstNameTextField = app.textFields[firstNameEn].exists ? app.textFields[firstNameEn] : app.textFields[firstNameFr]
        let emailTextField = app.textFields[emailEn].exists ? app.textFields[emailEn] : app.textFields[emailFr]
        let passwordSecureField = app.secureTextFields[passwordEn].exists ? app.secureTextFields[passwordEn] : app.secureTextFields[passwordFr]
        let confirmPasswordSecureField =
        app.secureTextFields[confirmPasswordEn].exists ? app.secureTextFields[confirmPasswordEn] : app.secureTextFields[confirmPasswordFr]

        nameTextField.tap()
        nameTextField.typeText("testName")

        firstNameTextField.tap()
        firstNameTextField.typeText("testFirstName")

        emailTextField.tap()
        emailTextField.typeText("test@example.com")

        passwordSecureField.tap()
        passwordSecureField.typeText("testpassword")

        confirmPasswordSecureField.tap()
        confirmPasswordSecureField.typeText("testpassword")

        XCTAssertEqual(passwordSecureField.value as? String, confirmPasswordSecureField.value as? String, "The passwords do not match.")
    }

    func testAgreementToggle() throws {
        navigateToRegisterView()
        let agreeFr = "J'accepte les Termes, la Politique de confidentialité et les Frais"
        let agreeEn = "I agree to all Term, Privacy Policy and Fees"
        let agreementButton =
        app.buttons[agreeEn].exists ? app.buttons[agreeEn] : app.buttons[agreeFr]

        XCTAssertTrue(agreementButton.exists)

        agreementButton.tap()

        let imageAgreementButtonChecked = app.images["checkmark.circle.fill"]
        XCTAssertTrue(imageAgreementButtonChecked.exists)

        agreementButton.tap()

        let imageAgreementButtonUnchecked = app.images["circle"]
        XCTAssertTrue(imageAgreementButtonUnchecked.exists)
    }

    func testSuccessfulRegistration() throws {
        navigateToRegisterView()
        let nameTextField = app.textFields[nameEn].exists ? app.textFields[nameEn] : app.textFields[nameFr]
        let firstNameTextField = app.textFields[firstNameEn].exists ? app.textFields[firstNameEn] : app.textFields[firstNameFr]
        let emailTextField = app.textFields[emailEn].exists ? app.textFields[emailEn] : app.textFields[emailFr]
        let passwordSecureField = app.secureTextFields[passwordEn].exists ? app.secureTextFields[passwordEn] : app.secureTextFields[passwordFr]
        let confirmPasswordSecureField =
        app.secureTextFields[confirmPasswordEn].exists ? app.secureTextFields[confirmPasswordEn] : app.secureTextFields[confirmPasswordFr]
        let signUpButton = app.buttons["SignUpButton"]

        XCTAssertTrue(nameTextField.exists)
        XCTAssertTrue(firstNameTextField.exists)
        XCTAssertTrue(emailTextField.exists)
        XCTAssertTrue(passwordSecureField.exists)
        XCTAssertTrue(confirmPasswordSecureField.exists)
        XCTAssertTrue(signUpButton.exists)

        nameTextField.tap()
        nameTextField.typeText("testName")
        XCTAssertEqual(nameTextField.value as? String, "testName")

        firstNameTextField.tap()
        firstNameTextField.typeText("testFirstName")
        XCTAssertEqual(firstNameTextField.value as? String, "testFirstName")

        emailTextField.tap()
        emailTextField.typeText("test@example.com")
        XCTAssertEqual(emailTextField.value as? String, "test@example.com")

        passwordSecureField.tap()
        passwordSecureField.typeText("testpassword")

        confirmPasswordSecureField.tap()
        confirmPasswordSecureField.typeText("testpassword")

        let termsButton = app.buttons["AgreementButton"]
        termsButton.tap()
    }
}
