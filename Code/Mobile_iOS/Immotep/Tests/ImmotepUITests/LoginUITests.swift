//
//  LoginUITests.swift
//  ImmotepTests
//
//  Created by Liebenguth Alessio on 12/10/2024.
//

import XCTest
import SwiftUI

final class LoginUITests: XCTestCase {
    let app = XCUIApplication()

    override func setUpWithError() throws {
        continueAfterFailure = false
        XCUIDevice.shared.orientation = .portrait
        app.launch()
    }

    override func tearDownWithError() throws {
        // Put teardown code here. This method is called after the invocation of each test method in the class.
    }

    func navigateToLoginView() {
        if app.buttons["person.crop.circle.fill"].exists {
            app.buttons["person.crop.circle.fill"].tap()

            let logoutButton = app.buttons["Logout"].exists || app.buttons["Se déconnecter"].exists
            XCTAssertTrue(logoutButton)

            if app.buttons["Logout"].exists {
                app.buttons["Logout"].tap()
            } else if app.buttons["Se déconnecter"].exists {
                app.buttons["Se déconnecter"].tap()
            }
        } else {
            return
        }
    }

    func testWelcomeTextExists() throws {
        navigateToLoginView()
        let welcomeText = app.staticTexts["Welcome back"].exists || app.staticTexts["Bienvenue !"].exists
        XCTAssertTrue(welcomeText)
    }

    func testEmailTextFieldExists() throws {
        navigateToLoginView()
        let emailTextField = app.textFields["Enter your email"].exists || app.textFields["Entrez votre email"].exists
        XCTAssertTrue(emailTextField)
    }

    func testPasswordSecureFieldExists() throws {
        navigateToLoginView()
        let passwordSecureField = app.secureTextFields["Enter your password"].exists || app.secureTextFields["Entrez votre mot de passe"].exists
        XCTAssertTrue(passwordSecureField)
    }

    func testSignInButtonExists() throws {
        navigateToLoginView()
        let signInButton = app.buttons["Sign In"].exists || app.buttons["Se connecter"].exists
        XCTAssertTrue(signInButton)
    }

    func testSignInWithValidCredentials() throws {
        navigateToLoginView()
        let emailTextField: XCUIElement
        let passwordSecureField: XCUIElement
        let signInButton: XCUIElement

        if app.textFields["Enter your email"].exists {
            emailTextField = app.textFields["Enter your email"]
        } else {
            emailTextField = app.textFields["Entrez votre email"]
        }

        if app.secureTextFields["Enter your password"].exists {
            passwordSecureField = app.secureTextFields["Enter your password"]
        } else {
            passwordSecureField = app.secureTextFields["Entrez votre mot de passe"]
        }

        if app.buttons["Sign In"].exists {
            signInButton = app.buttons["Sign In"]
        } else {
            signInButton = app.buttons["Se connecter"]
        }

        XCTAssertTrue(emailTextField.exists, "The email field does not exist.")
        XCTAssertTrue(passwordSecureField.exists, "The password field does not exist.")
        XCTAssertTrue(signInButton.exists, "The sign-in button does not exist.")

        emailTextField.tap()
        emailTextField.typeText("test@example.com")

        passwordSecureField.tap()
        passwordSecureField.typeText("testpassword")

        XCTAssertEqual(emailTextField.value as? String, "test@example.com", "The email is not filled in correctly.")
        XCTAssertEqual(passwordSecureField.value as? String, "••••••••••••", "The password is not filled in correctly.")
    }

        func testDontHaveAnAccountTextExists() throws {
            navigateToLoginView()
            let dontHaveAnAccount = app.staticTexts["Don't have an account ?"].exists || app.staticTexts["Vous n’avez pas de compte ?"].exists
            XCTAssertTrue(dontHaveAnAccount)
        }

        func testSignUpLinkExists() throws {
            navigateToLoginView()
            let linkExist = app.buttons["Sign Up"].exists || app.buttons["Se connecter"].exists
            XCTAssertTrue(linkExist)

            let signUpLink = app.buttons["signUpLink"]
            XCTAssertTrue(signUpLink.exists)
            signUpLink.tap()

            let registerTitle = app.staticTexts["Create your account"].exists || app.staticTexts["Créer un compte"].exists
            XCTAssertTrue(registerTitle)
        }
}
