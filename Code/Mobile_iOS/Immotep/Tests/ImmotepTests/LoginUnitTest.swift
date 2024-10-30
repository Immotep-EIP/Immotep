//
//  LoginUnitTest.swift
//  ImmotepTests
//
//  Created by Liebenguth Alessio on 13/10/2024.
//

import XCTest
import Foundation
import Combine

@testable import Immotep

enum UserServiceError: Error {
    case tokenExpired
    case noAccessToken
    case noRefreshToken
    case fetchProfileFailed
    case decodingFailed
}

extension UserServiceError: LocalizedError {
    var errorDescription: String? {
        switch self {
        case .tokenExpired:
            return "Token has expired."
        case .noAccessToken:
            return "No access token found."
        case .noRefreshToken:
            return "No refresh token found. Please log in again."
        case .fetchProfileFailed:
            return "Failed to fetch user profile."
        case .decodingFailed:
            return "Failed to decode user profile."
        }
    }
}

class MockAuthService: AuthServiceProtocol {
    var shouldReturnError: Bool = false

    func loginUser(email: String, password: String) async throws -> (String, String) {
        if shouldReturnError {
            throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Invalid credentials."])
        }
        return ("mockAccessToken", "mockRefreshToken")
    }

    func requestToken(grantType: String, email: String? = nil, password: String? = nil, refreshToken: String? = nil, keepMeSignedIn: Bool) async throws -> (String, String) {
        throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Mock requestToken not implemented."])
    }

    func authorizedRequest(for endpoint: String) async throws -> Data {
        throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Mock authorizedRequest not implemented."])
    }
}

class MockUserService: UserServiceProtocol {
    var shouldReturnError: Bool = false
    var userProfile: User = User(id: "1", email: "john.doe@example.com", firstname: "John", lastname: "Doe", role: "User")

    func getCurrentUser() async throws -> User {
        if shouldReturnError {
            throw UserServiceError.fetchProfileFailed
        }
        return userProfile
    }

    func fetchUserProfile(with token: String) async throws -> User {
        if shouldReturnError {
            throw UserServiceError.fetchProfileFailed
        }
        return userProfile
    }

    func logout() {
        // Implement logout logic if needed
    }
}

@MainActor
class LoginViewModelTests: XCTestCase {
    var viewModel: LoginViewModel!
    var authService: MockAuthService!
    var userService: MockUserService!

    override func setUp() {
        super.setUp()
        authService = MockAuthService()
        userService = MockUserService()
        viewModel = LoginViewModel(userService: userService, authService: authService)
    }

    func testSuccessfulLogin() async {
        viewModel.model.email = "john.doe@example.com"
        viewModel.model.password = "password123456789"
        authService.shouldReturnError = false
        userService.shouldReturnError = false

        await viewModel.signIn()

        try? await Task.sleep(nanoseconds: 5_000_000_000)

        XCTAssertEqual(viewModel.loginStatus, "Login successful!")
        XCTAssertTrue(viewModel.isLoggedIn)
        XCTAssertNotNil(viewModel.user)
        XCTAssertEqual(viewModel.user?.email, "john.doe@example.com")
    }

    func testLoginWithEmptyFields() async {
        await viewModel.signIn()

        XCTAssertEqual(viewModel.loginStatus, "Please enter both email and password.")
    }

    func testLoginWithMockedError() async {
        viewModel.model.email = "john.doe@example.com"
        viewModel.model.password = "password123456789"
        authService.shouldReturnError = true

        await viewModel.signIn()
        try? await Task.sleep(nanoseconds: 5_000_000_000)

        XCTAssertEqual(viewModel.loginStatus, "Error: Invalid credentials.")
    }

    func testLoginWithApiError() async {
        viewModel.model.email = "john.doe@example.com"
        viewModel.model.password = "wrongpassword"
        authService.shouldReturnError = true

        await viewModel.signIn()
        try? await Task.sleep(nanoseconds: 5_000_000_000)

        XCTAssertEqual(viewModel.loginStatus, "Error: Invalid credentials.")
    }

    func testFetchUserProfileError() async {
        viewModel.model.email = "john.doe@example.com"
        viewModel.model.password = "password123456789"
        authService.shouldReturnError = false
        userService.shouldReturnError = true

        await viewModel.signIn()
        try? await Task.sleep(nanoseconds: 5_000_000_000)

        XCTAssertEqual(viewModel.loginStatus, "Error: Failed to fetch user profile.")
    }
}
