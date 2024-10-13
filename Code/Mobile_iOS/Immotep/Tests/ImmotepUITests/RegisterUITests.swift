//
//  RegisterUITests.swift
//  ImmotepTests
//
//  Created by Liebenguth Alessio on 12/10/2024.
//

import XCTest

final class RegisterUITests: XCTestCase {
    let app = XCUIApplication()

    override func setUpWithError() throws {
        continueAfterFailure = false
        app.launch()

        let buttonSignUp = app.buttons["Sign Up"]
        if buttonSignUp.exists {
            buttonSignUp.tap()
        }
    }

    override func tearDownWithError() throws {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
    }

    func testWelcomeTextExists() throws {
        let welcomeText = app.staticTexts["Create your account"]
        XCTAssertTrue(welcomeText.exists)
    }

    func testSecondaryWelcomeTextExists() throws {
        let secWelcomeText = app.staticTexts["Join Immotep for your peace of mind!"]
        XCTAssertTrue(secWelcomeText.exists)
    }

    func testRequiredFieldsExist() throws {
        let nameText = app.staticTexts["Name*"]
        let firstNameText = app.staticTexts["First name*"]
        let emailText = app.staticTexts["Email*"]
        let passwordSecure = app.staticTexts["Password*"]
        let confirmPasswordSecure = app.staticTexts["Password confirmation*"]

        XCTAssertTrue(nameText.exists)
        XCTAssertTrue(firstNameText.exists)
        XCTAssertTrue(emailText.exists)
        XCTAssertTrue(passwordSecure.exists)
        XCTAssertTrue(confirmPasswordSecure.exists)
    }

    func testTextFieldsExist() throws {
        let nameTextField = app.textFields["Enter your name"]
        let firstNameTextField = app.textFields["Enter your first name"]
        let emailTextField = app.textFields["Enter your email"]
        let passwordSecureField = app.secureTextFields["Enter your password"]
        let confirmPasswordSecureField = app.secureTextFields["Enter your password confirmation"]

        XCTAssertTrue(nameTextField.exists)
        XCTAssertTrue(firstNameTextField.exists)
        XCTAssertTrue(emailTextField.exists)
        XCTAssertTrue(passwordSecureField.exists)
        XCTAssertTrue(confirmPasswordSecureField.exists)
    }

    func testFillingInFields() throws {
        let nameTextField = app.textFields["Enter your name"]
        let firstNameTextField = app.textFields["Enter your first name"]
        let emailTextField = app.textFields["Enter your email"]
        let passwordSecureField = app.secureTextFields["Enter your password"]
        let confirmPasswordSecureField = app.secureTextFields["Enter your password confirmation"]

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

        // Check if the two passwords are equal
    }

    func testAgreementToggle() throws {
        let agreementButton = app.buttons["I agree to all Term, Privacy Policy and Fees"]

        XCTAssertTrue(agreementButton.exists)

        agreementButton.tap()

        let imageAgreementButtonChecked = app.images["checkmark.circle.fill"]
        XCTAssertTrue(imageAgreementButtonChecked.exists)

        agreementButton.tap()

        let imageAgreementButtonUnchecked = app.images["circle"]
        XCTAssertTrue(imageAgreementButtonUnchecked.exists)
    }

    func testSuccessfulRegistration() throws {
        let nameTextField = app.textFields["Enter your name"]
        let firstNameTextField = app.textFields["Enter your first name"]
        let emailTextField = app.textFields["Enter your email"]
        let passwordSecureField = app.secureTextFields["Enter your password"]
        let confirmPasswordSecureField = app.secureTextFields["Enter your password confirmation"]
        let signInButton = app.buttons["Sign In"]

        nameTextField.tap()
        nameTextField.typeText("testName")

        firstNameTextField.tap()
        firstNameTextField.typeText("testFirstName")

        emailTextField.tap()
        emailTextField.typeText("Test@example.com")

        passwordSecureField.tap()
        passwordSecureField.typeText("Password")

        confirmPasswordSecureField.tap()
        confirmPasswordSecureField.typeText("Password")

        app.buttons["I agree to all Term, Privacy Policy and Fees"].tap()

        signInButton.tap()

        let successMessage = app.staticTexts["Registration successful!"]
        XCTAssertTrue(successMessage.waitForExistence(timeout: 5))
    }

    func testLaunchPerformance() throws {
        if #available(macOS 10.15, iOS 13.0, tvOS 13.0, watchOS 7.0, *) {
            measure(metrics: [XCTApplicationLaunchMetric()]) {
                XCUIApplication().launch()
            }
        }
    }
}
