//
//  RegisterViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 05/10/2024.
//

import Foundation
import Combine

@MainActor
class RegisterViewModel: ObservableObject {
    @Published var model = RegisterModel()
    @Published var registerStatus: String = ""

    private var cancellables = Set<AnyCancellable>()
    private var apiService: ApiServiceProtocol

    init(apiService: ApiServiceProtocol = ApiService.shared) {
        self.apiService = apiService
    }

    func signIn() async {
        registerStatus = ""

        if let errorMessage = validateFields() {
            registerStatus = errorMessage
            return
        }
        let apiServiceCopy = apiService

        Task {
            do {
                let response = try await apiServiceCopy.registerUser(with: model)
                registerStatus = "Registration successful!"
            } catch {
                registerStatus = "Error: \(error.localizedDescription)"
            }
        }
    }

    func validateFields() -> String? {
        guard !model.name.isEmpty,
              !model.firstName.isEmpty,
              !model.email.isEmpty,
              !model.password.isEmpty,
              !model.passwordConfirmation.isEmpty else {
            return "Please fill in all fields."
        }

        guard model.password == model.passwordConfirmation else {
            return "Passwords do not match."
        }

        guard model.agreement else {
            return "You must agree to the terms and conditions."
        }

        guard isValidEmail(model.email) else {
            return "Please enter a valid email address."
        }

        return nil
    }

    private func isValidEmail(_ email: String) -> Bool {
        return email.contains("@")
    }
}
