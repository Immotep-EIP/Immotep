//
//  LoginViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 05/10/2024.
//
import Foundation
import Combine

class LoginViewModel: ObservableObject {
    @Published var model = LoginModel()
    @Published var loginStatus: String = ""

    private var cancellables = Set<AnyCancellable>()

    // Simulating login for now (replace with real login logic)
    func signIn() {
        guard !model.email.isEmpty && !model.password.isEmpty else {
            loginStatus = "Please enter both email and password."
            return
        }
        if model.email == "test@example.com" && model.password == "password" {
            loginStatus = "Login successful!"
        } else {
            loginStatus = "Invalid email or password."
        }
    }
}
