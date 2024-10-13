//
//  RegisterUnitTest.swift
//  ImmotepTests
//
//  Created by Liebenguth Alessio on 13/10/2024.
//

import XCTest
@testable import Immotep

final class RegisterUnitTest: XCTestCase {
    var viewModel: RegisterViewModel!

    override func setUpWithError() throws {
        viewModel = RegisterViewModel()
    }

    override func tearDownWithError() throws {
        viewModel = nil
    }

    func testSignInWithEmptyFields() {
        viewModel.model.name = ""
        viewModel.model.firstName = ""
        viewModel.model.email = ""
        viewModel.model.password = ""
        viewModel.model.passwordConfirmation = ""

        viewModel.signIn()

        XCTAssertEqual(viewModel.registerStatus, "Please fill in all fields.")
    }

    func testSignInWithMismatchedPasswords() {
        viewModel.model.name = "testName"
        viewModel.model.firstName = "testFirstName"
        viewModel.model.email = "test@example.com"
        viewModel.model.password = "password"
        viewModel.model.passwordConfirmation = "differentpassword"

        viewModel.signIn()

        XCTAssertEqual(viewModel.registerStatus, "Passwords do not match.")
    }

    func testSignInWithoutAgreement() {
        viewModel.model.name = "testName"
        viewModel.model.firstName = "testFirstName"
        viewModel.model.email = "test@example.com"
        viewModel.model.password = "password"
        viewModel.model.passwordConfirmation = "password"
        viewModel.model.agreement = false

        viewModel.signIn()

        XCTAssertEqual(viewModel.registerStatus, "You must agree to the terms and conditions.")
    }

    func testSignInWithInvalidEmail() {
        viewModel.model.name = "testName"
        viewModel.model.firstName = "testFirstName"
        viewModel.model.email = "invalidEmail"
        viewModel.model.password = "password"
        viewModel.model.passwordConfirmation = "password"
        viewModel.model.agreement = true

        viewModel.signIn()

        XCTAssertEqual(viewModel.registerStatus, "Please enter a valid email address.")
    }

    func testSignInWithValidDetails() {
        viewModel.model.name = "testName"
        viewModel.model.firstName = "testFirstName"
        viewModel.model.email = "Test@example.com"
        viewModel.model.password = "Password"
        viewModel.model.passwordConfirmation = "Password"
        viewModel.model.agreement = true

        viewModel.signIn()

        XCTAssertEqual(viewModel.registerStatus, "Registration successful!")
    }

    func testSignInWithInvalidEmailAndPassword() {
        viewModel.model.name = "testName"
        viewModel.model.firstName = "testFirstName"
        viewModel.model.email = "invalid@example.com"
        viewModel.model.password = "wrongpassword"
        viewModel.model.passwordConfirmation = "wrongpassword"
        viewModel.model.agreement = true

        viewModel.signIn()

        XCTAssertEqual(viewModel.registerStatus, "Invalid email or password.")
    }

//    func testPerformanceExample() throws {
//        // This is an example of a performance test case.
//        self.measure {
//            // Put the code you want to measure the time of here.
//        }
//    }
}
