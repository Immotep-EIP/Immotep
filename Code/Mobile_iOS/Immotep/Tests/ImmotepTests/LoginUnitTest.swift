//
//  LoginUnitTest.swift
//  ImmotepTests
//
//  Created by Liebenguth Alessio on 13/10/2024.
//

import XCTest
@testable import Immotep

final class LoginUnitTest: XCTestCase {
    var viewModel: LoginViewModel!

    override func setUpWithError() throws {
        viewModel = LoginViewModel()
    }

    override func tearDownWithError() throws {
        viewModel = nil
    }

    func testSignInWithEmptyCredentials() {
        viewModel.model.email = ""
        viewModel.model.password = ""

        viewModel.signIn()

        XCTAssertEqual(viewModel.loginStatus, "Please enter both email and password.")
    }

    func testSignInWithValidCredentials() {
        viewModel.model.email = "test@example.com"
        viewModel.model.password = "password"

        viewModel.signIn()

        XCTAssertEqual(viewModel.loginStatus, "Login successful!")
    }

    func testSignInWithInvalidEmail() {
        viewModel.model.email = "invalid@example.com"
        viewModel.model.password = "password"

        viewModel.signIn()

        XCTAssertEqual(viewModel.loginStatus, "Invalid email or password.")
    }

    func testSignInWithInvalidPassword() {
        viewModel.model.email = "test@example.com"
        viewModel.model.password = "wrongpassword"

        viewModel.signIn()

        XCTAssertEqual(viewModel.loginStatus, "Invalid email or password.")
    }

    func testSignInWithEmailOnly() {
        viewModel.model.email = "test@example.com"
        viewModel.model.password = ""

        viewModel.signIn()

        XCTAssertEqual(viewModel.loginStatus, "Please enter both email and password.")
    }

    func testSignInWithPasswordOnly() {
        viewModel.model.email = ""
        viewModel.model.password = "password"

        viewModel.signIn()

        XCTAssertEqual(viewModel.loginStatus, "Please enter both email and password.")
    }

//    func testPerformanceExample() throws {
//        // This is an example of a performance test case.
//        self.measure {
//            // Put the code you want to measure the time of here.
//        }
//    }
}
