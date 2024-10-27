//
//  ProfileViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import SwiftUI

@MainActor
class ProfileViewModel: ObservableObject {
    @Published var user: User?
    @AppStorage("isLoggedIn") private var isLoggedIn: Bool = false
    private let userService = UserService()

    init() {
        if isLoggedIn {
            loadUser()
        }
    }

    func loadUser() {
        Task {
            do {
                let currentUser = try await userService.getCurrentUser()
                self.user = currentUser
            } catch {
                print("Failed to load user: \(error)")
            }
        }
    }
}
