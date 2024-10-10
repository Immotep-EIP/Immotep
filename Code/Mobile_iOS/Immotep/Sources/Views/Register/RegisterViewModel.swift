//
//  LoginViewModel.swift
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
        guard !model.email.isEmpty && !model.password.isEmpty else {
            registerStatus = "Please enter both email and password."
            return
        }
        if model.email == "test@example.com" && model.password == "password" {
            registerStatus = "Login successful!"
        } else {
            registerStatus = "Invalid email or password."
        }
    }
}
