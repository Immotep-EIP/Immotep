//
//  RegisterUITests.swift
//  ImmotepTests
//
//  Created by Liebenguth Alessio on 12/10/2024.
//

import XCTest

final class RegisterUITests: XCTestCase {

    override func setUpWithError() throws {
        // Put setup code here. This method is called before the invocation of each test method in the class.

        // In UI tests it is usually best to stop immediately when a failure occurs.
        continueAfterFailure = false

        // In UI tests it’s important to set the initial state - such as interface orientation - required for your tests before they run. The setUp method is a good place to do this.
    }

    override func tearDownWithError() throws {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
    }

    func testRegisterViewFunctionality() throws {

        let app = XCUIApplication()
        app.launch()

        let buttonSignUp = app.buttons["Sign Up"]
        buttonSignUp.tap()

        let welcomeText = app.staticTexts["Create your account"]
        XCTAssertTrue(welcomeText.exists)

        let secWelcomeText = app.staticTexts["Join Immotep for your peace of mind!"]
        XCTAssertTrue(secWelcomeText.exists)

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

        nameTextField.tap()
        nameTextField.typeText("testName")

        firstNameTextField.tap()
        firstNameTextField.typeText("testFirstName")

        emailTextField.tap()
        emailTextField.typeText("testEmail")

        passwordSecureField.tap()
        passwordSecureField.typeText("testpassword")

        confirmPasswordSecureField.tap()
        confirmPasswordSecureField.typeText("testpassword")

        // check if the two passwords are equals

    }

    func testLaunchPerformance() throws {
        if #available(macOS 10.15, iOS 13.0, tvOS 13.0, watchOS 7.0, *) {
            // This measures how long it takes to launch your application.
            measure(metrics: [XCTApplicationLaunchMetric()]) {
                XCUIApplication().launch()
            }
        }
    }
}
