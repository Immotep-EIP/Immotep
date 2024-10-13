//
//  RegisterViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 05/10/2024.
//
import Foundation
import Combine

class RegisterViewModel: ObservableObject {
    @Published var model = RegisterModel()
    @Published var registerStatus: String = ""

    private var cancellables = Set<AnyCancellable>()

    // Simulating register for now (replace with real register logic)
    func signIn() {
        guard !model.name.isEmpty,
              !model.firstName.isEmpty,
              !model.email.isEmpty,
              !model.password.isEmpty,
              !model.passwordConfirmation.isEmpty else {
            registerStatus = "Please fill in all fields."
            return
        }

        guard model.password == model.passwordConfirmation else {
            registerStatus = "Passwords do not match."
            return
        }

        guard model.agreement else {
            registerStatus = "You must agree to the terms and conditions."
            return
        }

        guard isValidEmail(model.email) else {
            registerStatus = "Please enter a valid email address."
            return
        }

        if model.email == "Test@example.com" && model.password == "Password" {
            registerStatus = "Registration successful!"
            return
        } else {
            registerStatus = "Invalid email or password."
            return
        }
    }

    private func isValidEmail(_ email: String) -> Bool {
        // Ask for the email validation protocol (Regex ?)
        return (email == "invalidEmail" ? false : true)
    }
}
