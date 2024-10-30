//
//  ProfileView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import SwiftUI

struct ProfileView: View {
    @StateObject private var viewModel = ProfileViewModel()
    @AppStorage("isLoggedIn") private var isLoggedIn: Bool = false
    @State private var navigateToLogin = false

    var body: some View {
        NavigationStack {
            VStack(alignment: .leading) {
                TopBar()
                Text("Profile".localized())
                    .font(.title)
                    .fontWeight(.bold)
                    .padding(20)

                VStack(alignment: .leading, spacing: 20) {
                    if let user = viewModel.user {
                        Text("Name: \(user.firstname) \(user.lastname)".localized())
                        Text("Email: \(user.email)".localized())
                    } else {
                        Text("Loading user information...".localized())
                    }
                }
                .padding(.leading, 20)

                Button("Logout".localized()) {
                    signOut()
                }
                .padding()
                .accessibilityIdentifier("logoutButton")

                .navigationDestination(isPresented: $navigateToLogin) {
                    LoginView()
                }
                Spacer()
                TaskBar()
            }
        }
        .navigationBarBackButtonHidden(true)
    }

    private func signOut() {
        TokenStorage.clearTokens()
        isLoggedIn = false
        viewModel.user = nil
        navigateToLogin = true
    }
}

struct ProfileView_Previews: PreviewProvider {
    static var previews: some View {
        ProfileView()
    }
}
