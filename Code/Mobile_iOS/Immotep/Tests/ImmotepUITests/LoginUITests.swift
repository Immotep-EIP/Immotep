//
//  LoginUITests.swift
//  ImmotepTests
//
//  Created by Liebenguth Alessio on 12/10/2024.
//

import XCTest

final class LoginUITests: XCTestCase {
    let app = XCUIApplication()

    override func setUpWithError() throws {
        continueAfterFailure = false
        app.launch()
    }

    override func tearDownWithError() throws {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
    }

    func testWelcomeTextExists() throws {
        let welcomeText = app.staticTexts["Welcome back"]
        XCTAssertTrue(welcomeText.exists)
    }

    func testEmailTextFieldExists() throws {
        let emailTextField = app.textFields["Enter your email"]
        XCTAssertTrue(emailTextField.exists)
    }

    func testPasswordSecureFieldExists() throws {
        let passwordSecureField = app.secureTextFields["Enter your password"]
        XCTAssertTrue(passwordSecureField.exists)
    }

    func testSignInButtonExists() throws {
        let signInButton = app.buttons["Sign In"]
        XCTAssertTrue(signInButton.exists)
    }

    func testSignInWithValidCredentials() throws {
        let emailTextField = app.textFields["Enter your email"]
        let passwordSecureField = app.secureTextFields["Enter your password"]
        let signInButton = app.buttons["Sign In"]

        emailTextField.tap()
        emailTextField.typeText("test@example.com")

        passwordSecureField.tap()
        passwordSecureField.typeText("testpassword")

        signInButton.tap()

        // search for a way to test the toggle

        // let loginStatusMessage = app.staticTexts["Login successful!"]
        // XCTAssertTrue(loginStatusMessage.waitForExistence(timeout: 5))
    }

    func testDontHaveAnAccountTextExists() throws {
        let dontHaveAnAccount = app.staticTexts["Don't have an account ?"]
        XCTAssertTrue(dontHaveAnAccount.exists)
    }

    func testSignUpLinkExists() throws {
        let signUpLink = app.buttons["signUpLink"]
        XCTAssertTrue(signUpLink.exists)
        signUpLink.tap()

        let registerTitle = app.staticTexts["Create your account"]
        XCTAssertTrue(registerTitle.waitForExistence(timeout: 5))
    }

    func testLaunchPerformance() throws {
        if #available(macOS 10.15, iOS 13.0, tvOS 13.0, watchOS 7.0, *) {
            measure(metrics: [XCTApplicationLaunchMetric()]) {
                XCUIApplication().launch()
            }
        }
    }
}
