//
//  LoginUITests.swift
//  ImmotepTests
//
//  Created by Liebenguth Alessio on 12/10/2024.
//

import XCTest

final class LoginUITests: XCTestCase {

    override func setUpWithError() throws {
        // Put setup code here. This method is called before the invocation of each test method in the class.

        // In UI tests it is usually best to stop immediately when a failure occurs.
        continueAfterFailure = false

        // In UI tests it’s important to set the initial state - such as interface orientation - required for your tests before they run. The setUp method is a good place to do this.
    }

    override func tearDownWithError() throws {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
    }

    @MainActor
    func testLoginViewFunctionality() throws {
        let app = XCUIApplication()
        app.launch()

        let welcomeText = app.staticTexts["Welcome back"]
        XCTAssertTrue(welcomeText.exists)

        let emailTextField = app.textFields["Enter your email"]
        let passwordSecureField = app.secureTextFields["Enter your password"]

        XCTAssertTrue(emailTextField.exists)
        XCTAssertTrue(passwordSecureField.exists)

        emailTextField.tap()
        emailTextField.typeText("test@example.com")

        passwordSecureField.tap()
        passwordSecureField.typeText("testpassword")

        let signInButton = app.buttons["Sign In"]
        XCTAssertTrue(signInButton.exists)
        signInButton.tap()

        // search for a way to test the toggle

        // let loginStatusMessage = app.staticTexts["Login successful!"]
        // XCTAssertTrue(loginStatusMessage.waitForExistence(timeout: 5))

        let dontHaveAnAccount = app.staticTexts["Don't have an account ?"]
        XCTAssertTrue(dontHaveAnAccount.exists)

        let signUpLink = app.buttons["signUpLink"]
        XCTAssertTrue(signUpLink.exists)
        signUpLink.tap()

        let registerTitle = app.staticTexts["Create your account"]
        XCTAssertTrue(registerTitle.waitForExistence(timeout: 5))
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
