//
//  LoginViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 05/10/2024.
//

import Foundation
import Combine

@MainActor
class LoginViewModel: ObservableObject {
    @Published var model = LoginModel()
    @Published var loginStatus: String = ""
    @Published var isLoggedIn: Bool = false
    @Published var user: User?

    private var cancellables = Set<AnyCancellable>()
    private let userService = UserService()

    func signIn() {
        loginStatus = ""

        if let errorMessage = validateFields() {
            loginStatus = errorMessage
            return
        }

        Task {
            do {
                let (accessToken, refreshToken) = try await AuthService.shared.loginUser(email: model.email, password: model.password)
                TokenStorage.storeTokens(accessToken: accessToken, refreshToken: refreshToken)
                                user = try await userService.fetchUserProfile(with: accessToken)
                loginStatus = "Login successful!"
                DispatchQueue.main.asyncAfter(deadline: .now() + 2) {
                    self.isLoggedIn = true
                }
            } catch {
                loginStatus = "Error: \(error.localizedDescription)"
            }
        }
    }

    private func validateFields() -> String? {
        guard !model.email.isEmpty, !model.password.isEmpty else {
            return "Please enter both email and password."
        }
        return nil
    }
}
