//
//  LoginViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 05/10/2024.
//

import Foundation
import Combine
import SwiftUI

@MainActor
class LoginViewModel: ObservableObject {
    @Published var model = LoginModel()
    @Published var loginStatus: String = ""
    @AppStorage("isLoggedIn") var isLoggedIn: Bool = false
    @Published var user: User?

    @AppStorage("user") private var storedUserData: String = ""
    @Published var profileViewModel: ProfileViewModel?

    public var cancellables = Set<AnyCancellable>()
    public let userService: UserServiceProtocol
    public let authService: AuthServiceProtocol

    init(userService: UserServiceProtocol = UserService(), authService: AuthServiceProtocol = AuthService.shared) {
        self.userService = userService
        self.authService = authService
    }

    func signIn() async {
        loginStatus = ""

        if let errorMessage = validateFields() {
            loginStatus = errorMessage
            return
        }

        let userServiceCopy = userService
        let authServiceCopy = authService

        Task {
            do {
                let (accessToken, refreshToken) = try await authServiceCopy.loginUser(email: model.email, password: model.password, keepMeSignedIn: model.keepMeSignedIn)
                TokenStorage.storeTokens(accessToken: accessToken, refreshToken: refreshToken, expiresIn: nil, keepMeSignedIn: model.keepMeSignedIn )
                user = try await userServiceCopy.fetchUserProfile(with: accessToken)
                loginStatus = "Login successful!"
//                DispatchQueue.main.asyncAfter(deadline: .now() + 1) {

                if let user = user {
                    saveUser(user)
                    profileViewModel = ProfileViewModel()
                }
                    self.isLoggedIn = true
//                }
            } catch {
                loginStatus = "Error: \(error.localizedDescription)"
            }
        }
    }

    func validateFields() -> String? {
        guard !model.email.isEmpty, !model.password.isEmpty else {
            return "Please enter both email and password."
        }
        return nil
    }

    private func saveUser(_ user: User) {
        let encoder = JSONEncoder()
        if let encodedData = try? encoder.encode(user) {
            if let jsonString = String(data: encodedData, encoding: .utf8) {
                storedUserData = jsonString
            }
        }
    }
}
