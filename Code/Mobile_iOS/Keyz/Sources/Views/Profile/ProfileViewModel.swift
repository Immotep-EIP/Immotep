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
    @AppStorage("user") private var storedUserData: String = ""

    init() {
        if let loadedUser = loadUser() {
            self.user = loadedUser
        }
    }

    private func loadUser() -> User? {
        let decoder = JSONDecoder()
        if let data = storedUserData.data(using: .utf8),
           let decodedUser = try? decoder.decode(User.self, from: data) {
            return decodedUser
        }
        return nil
    }
}
