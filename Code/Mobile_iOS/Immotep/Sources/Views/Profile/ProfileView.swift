//
//  ProfileView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import SwiftUI

struct ProfileView: View {
    @StateObject private var viewModel = ProfileViewModel()
    @StateObject private var keyboardObserver = KeyboardObserver()
    @AppStorage("isLoggedIn") private var isLoggedIn: Bool = false
    @State private var navigateToLogin = false

    @State private var editableEmail: String = ""
    @State private var editableFirstname: String = ""
    @State private var editableLastname: String = ""
    @State private var editablePhone: String = "+33 123456789"
    @State private var editablePassword: String = "Password123"

    @State private var isEditing: Bool = false

    var body: some View {
        NavigationStack {
            VStack(spacing: 0) {
                TopBar(title: "Profile".localized())

                ScrollView {
                    VStack(spacing: 20) {
                        Image(systemName: "person.crop.circle.fill")
                            .font(.system(size: 50))
                            .foregroundStyle(Color("textColor"))
                            .padding(.top, 20)

                        VStack(spacing: 20) {
                            CustomTextInput(title: "Email", placeholder: "", text: $editableEmail, isSecure: false)
                                .disabled(!isEditing)
                            CustomTextInput(title: "Password", placeholder: "", text: $editablePassword, isSecure: true)
                                .disabled(!isEditing)
                            CustomTextInput(title: "First name", placeholder: "", text: $editableFirstname, isSecure: false)
                                .disabled(!isEditing)
                            CustomTextInput(title: "Name", placeholder: "", text: $editableLastname, isSecure: false)
                                .disabled(!isEditing)
                            CustomTextInput(title: "Phone number", placeholder: "", text: $editablePhone, isSecure: false)
                                .disabled(!isEditing)
                        }
                        .padding([.bottom, .leading, .trailing], 20)

                        HStack {
                            Button(isEditing ? "Cancel".localized() : "Edit".localized()) {
                                isEditing.toggle()
                                if !isEditing {
                                    resetFields()
                                }
                            }
                            .padding()

                            if isEditing {
                                Button("Confirm".localized()) {
                                    saveChanges()
                                    isEditing = false
                                }
                                .padding()
                            }
                        }
                    }
                    .overlay(
                        RoundedRectangle(cornerRadius: 10)
                            .stroke(Color.gray, lineWidth: 0.5)
                    )
                    .padding(10)

                    Button("Logout".localized()) {
                        signOut()
                    }
                    .frame(maxWidth: .infinity)
                    .padding()
                    .background(Color("btnColor"))
                    .foregroundColor(.white)
                    .font(.headline)
                    .cornerRadius(10)
                    .padding([.bottom, .leading, .trailing], 20)
                }

                if !keyboardObserver.isKeyboardVisible {
                    TaskBar()
                }
            }
        }
        .navigationBarBackButtonHidden(true)
        .onChange(of: viewModel.user?.email) {
            loadUserData()
        }
    }

    private func signOut() {
        TokenStorage.clearTokens()
        isLoggedIn = false
        viewModel.user = nil
        navigateToLogin = true
    }

    private func resetFields() {
        loadUserData()
    }

    private func saveChanges() {
        print("Saving changes: \(editableEmail), \(editableFirstname), \(editableLastname), \(editablePhone)") // Update when API route done
    }

    private func loadUserData() {
        editableEmail = viewModel.user?.email ?? ""
        editableFirstname = viewModel.user?.firstname ?? ""
        editableLastname = viewModel.user?.lastname ?? ""
    }
}

struct ProfileView_Previews: PreviewProvider {
    static var previews: some View {
        ProfileView()
    }
}
