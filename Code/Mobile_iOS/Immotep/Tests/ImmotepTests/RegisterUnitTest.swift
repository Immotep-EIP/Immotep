//
//  RegisterUnitTest.swift
//  ImmotepTests
//
//  Created by Liebenguth Alessio on 13/10/2024.
//

import XCTest
import Foundation
import Combine

@testable import Immotep

class MockRegisterViewModel: RegisterViewModel {
    var shouldReturnError: Bool = false

    override func signIn() async {
        registerStatus = ""

        if let errorMessage = validateFields() {
            registerStatus = errorMessage
            return
        }

        if shouldReturnError {
            registerStatus = "Error: Mocked error occurred"
            return
        }

        registerStatus = "Registration successful!"
    }
}

class MockApiService: ApiServiceProtocol {
    var shouldReturnError: Bool = false
    var mockResponseCode: Int = 201
    var mockResponseData: Data = Data()

    func registerUser(with model: RegisterModel) async throws -> (String) {
        if shouldReturnError {
            switch mockResponseCode {
            case 400:
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Empty fields."])
            case 409:
                throw NSError(domain: "", code: 409, userInfo: [NSLocalizedDescriptionKey: "Email already exists."])
            default:
                throw NSError(domain: "", code: mockResponseCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(mockResponseCode)"])
            }
        }
        print("############## OKAY ##############")
        return ("Registration successful!")
    }
}

@MainActor
final class RegisterViewModelTests: XCTestCase {
    var viewModel: RegisterViewModel!
    var apiService: MockApiService!

    override func setUp() {
        super.setUp()
        apiService = MockApiService()
        viewModel = RegisterViewModel(apiService: apiService)
    }

    func testSuccessfulRegistration() async {
        viewModel.model.name = "John"
        viewModel.model.firstName = "Doe"
        viewModel.model.email = "john.doe@example.com"
        viewModel.model.password = "password123456789"
        viewModel.model.passwordConfirmation = "password123456789"
        viewModel.model.agreement = true

        apiService.shouldReturnError = false

        await viewModel.signIn()

        try? await Task.sleep(nanoseconds: 5_000_000_000)
        XCTAssertEqual(viewModel.registerStatus, "Registration successful!")
    }

    func testRegistrationWithEmptyFields() async {
        await viewModel.signIn()

        let expectedMessages = [
            "Please fill in all fields.",
            "Veuillez remplir tous les champs."
        ]

        XCTAssertTrue(
            expectedMessages.contains(viewModel.registerStatus),
            "registerStatus is '\(viewModel.registerStatus)', but expected one of \(expectedMessages)"
        )
    }

    func testRegistrationWithMismatchedPasswords() async {
        viewModel.model.name = "John"
        viewModel.model.firstName = "Doe"
        viewModel.model.email = "john.doe@example.com"
        viewModel.model.password = "password123456789"
        viewModel.model.passwordConfirmation = "differentPassword123456789"
        viewModel.model.agreement = true

        await viewModel.signIn()

        let expectedMessages = [
            "Passwords do not match.",
            "Les mots de passes sont différents."
        ]

        XCTAssertTrue(
            expectedMessages.contains(viewModel.registerStatus),
            "registerStatus is '\(viewModel.registerStatus)', but expected one of \(expectedMessages)"
        )
    }

    func testRegistrationWithMockedError() async {
        apiService.shouldReturnError = true
        apiService.mockResponseCode = 500

        viewModel.model.name = "John"
        viewModel.model.firstName = "Doe"
        viewModel.model.email = "john.doe@example.com"
        viewModel.model.password = "password123456789"
        viewModel.model.passwordConfirmation = "password123456789"
        viewModel.model.agreement = true

        await viewModel.signIn()
        try? await Task.sleep(nanoseconds: 5_000_000_000)

        XCTAssertEqual(viewModel.registerStatus, "Error: Failed with status code: 500")
    }

    func testRegistrationWithApiError() async {
        apiService.shouldReturnError = true
        apiService.mockResponseCode = 400

        viewModel.model.name = "John"
        viewModel.model.firstName = "Doe"
        viewModel.model.email = "john.doe@example.com"
        viewModel.model.password = "password123456789"
        viewModel.model.passwordConfirmation = "password123456789"
        viewModel.model.agreement = true

        await viewModel.signIn()
        try? await Task.sleep(nanoseconds: 5_000_000_000)
        XCTAssertEqual(viewModel.registerStatus, "Error: Empty fields.")
    }
}
