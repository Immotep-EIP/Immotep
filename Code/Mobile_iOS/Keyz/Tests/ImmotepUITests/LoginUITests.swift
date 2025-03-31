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
        app.launchArguments.append("-UIPreferredContentSizeCategoryName")
        app.launchArguments.append("-AppleLanguages")
        app.launchArguments.append("en")
        app.launchArguments += ["-UIViewAnimationDuration", "0"]
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

        let predicate = NSPredicate(format: "value == '••••••••••••'")
        let expectation = expectation(for: predicate, evaluatedWith: passwordSecureField, handler: nil)
        wait(for: [expectation], timeout: 5.0)

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
            XCTAssertTrue(linkExist, "Neither 'Sign Up' nor 'Se connecter' button exists.")

            let signUpLink = app.buttons["signUpLink"]
            let signUpLinkPredicate = NSPredicate(format: "exists == true")
            let signUpLinkExpectation = expectation(for: signUpLinkPredicate, evaluatedWith: signUpLink, handler: nil)
            wait(for: [signUpLinkExpectation], timeout: 5.0)
            XCTAssertTrue(signUpLink.exists, "The 'signUpLink' button does not exist.")
            signUpLink.tap()

            let registerTitlePredicate = NSPredicate { _, _ in
                return self.app.staticTexts["Create your account"].exists || self.app.staticTexts["Créer un compte"].exists
            }
            let registerTitleExpectation = expectation(for: registerTitlePredicate, evaluatedWith: nil, handler: nil)
            wait(for: [registerTitleExpectation], timeout: 10.0)

            let registerTitle = app.staticTexts["Create your account"].exists || app.staticTexts["Créer un compte"].exists
            XCTAssertTrue(registerTitle, "The register title ('Create your account' or 'Créer un compte') does not appear.")
        }
}
