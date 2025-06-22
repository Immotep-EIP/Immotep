//
//  SettingsView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 21/10/2024.
//

import SwiftUI

struct SettingsView: View {
    @EnvironmentObject private var loginViewModel: LoginViewModel
    @Binding var selectedTab: Int

    @AppStorage("lang") private var lang: String = "en"
    @AppStorage("theme") private var selectedTheme: String = ThemeOption.system.rawValue
    @AppStorage("isLoggedIn") private var isLoggedIn: Bool = false
    
    @State private var navigateToLogin = false
    @State private var editableEmail: String = ""
    @State private var editableFirstname: String = ""
    @State private var editableLastname: String = ""
    @State private var isEditing: Bool = false
    
    var body: some View {
        NavigationStack {
            VStack(spacing: 0) {
                TopBar(title: "Keyz".localized())
                
                VStack(spacing: 0) {
                    Text("Settings".localized())
                        .font(.title2)
                        .fontWeight(.bold)
                        .frame(maxWidth: .infinity, alignment: .leading)
                        .padding(.horizontal, 20)
                        .padding(.top, 10)
                        .padding(.bottom, 5)
                    
                    Form {
                        VStack(alignment: .center, spacing: 10) {
                            Image(systemName: "person.crop.circle.fill")
                                .font(.system(size: 50))
                                .foregroundStyle(Color("textColor"))
                                .padding(.vertical, 10)
                            
                            Text("User information".localized())
                                .font(.headline)
                                .frame(maxWidth: .infinity, alignment: .leading)
                            
                            HStack {
                                Text("ID".localized())
                                    .foregroundColor(.gray)
                                Spacer()
                                Text(loginViewModel.user?.id ?? "N/A")
                                    .foregroundColor(.gray)
                            }
                            .padding(.vertical, 5)
                            
                            HStack {
                                Text("Role".localized())
                                    .foregroundColor(.gray)
                                Spacer()
                                Text(loginViewModel.user?.role == "tenant" ? "Tenant" : "Owner")
                                    .foregroundColor(.gray)
                            }
                            .padding(.vertical, 5)
                            
                            CustomTextInput(title: "First Name".localized(), placeholder: "", text: $editableFirstname, isSecure: false)
                                .disabled(!isEditing)
//                                .onTapGesture {
//                                    if isEditing {
//                                    }
//                                }
                            CustomTextInput(title: "Name".localized(), placeholder: "", text: $editableLastname, isSecure: false)
                                .disabled(!isEditing)
//                                .onTapGesture {
//                                    if isEditing {
//                                    }
//                                }
                            CustomTextInput(title: "Email", placeholder: "", text: $editableEmail, isSecure: false)
                                .disabled(!isEditing)
//                                .onTapGesture {
//                                    if isEditing {
//                                    }
//                                }
                            
//                            HStack {
//                                Button(isEditing ? "Cancel".localized() : "Edit".localized()) {
//                                    isEditing.toggle()
//                                    if !isEditing {
//                                        resetFields()
//                                    }
//                                }
//                                .padding()
//                                
//                                if isEditing {
//                                    Button("Confirm".localized()) {
//                                        saveChanges()
//                                        isEditing = false
//                                    }
//                                    .padding()
//                                }
//                            }
                        }
                        
                        VStack {
                            Text("Language".localized())
                                .font(.headline)
                                .frame(maxWidth: .infinity, alignment: .leading)
                            
                            Picker(selection: $lang, label: Text("Language")) {
                                Text("Français").tag("fr")
                                Text("English").tag("en")
                            }
                            .pickerStyle(SegmentedPickerStyle())
                            .onChange(of: lang) {
                                Task {
                                    await Bundle.setLanguage(lang)
                                }
                            }
                        }

                        VStack {
                            Text("Theme".localized())
                                .font(.headline)
                                .frame(maxWidth: .infinity, alignment: .leading)
                            
                            Picker("Theme", selection: $selectedTheme) {
                                ForEach(ThemeOption.allCases, id: \.self) { theme in
                                    Text(theme.rawValue.localized())
                                        .tag(theme.rawValue)
                                }
                            }
                            .pickerStyle(SegmentedPickerStyle())
                            .onChange(of: selectedTheme) {
                                Task { @MainActor in
                                    ThemeManager.applyTheme(theme: selectedTheme)
                                }
                            }
                        }

                        Button("Logout".localized()) {
                            Task {
                                await signOut()
                            }
                        }
                        .frame(maxWidth: .infinity)
                        .padding()
                        .background(Color.red)
                        .foregroundColor(.white)
                        .font(.headline)
                        .cornerRadius(10)
                        .padding(.vertical, 20)
                    }
                }
                .background(Color(UIColor.systemGroupedBackground))
            }
            .navigationBarBackButtonHidden(true)
            .onChange(of: loginViewModel.user?.email) {
                loadUserData()
            }
            .onAppear {
                loginViewModel.loadUser()
                loadUserData()
            }
        }
    }
    
    private func signOut() async {
        await UserService.shared.logout()
        TokenStorage.clearTokens()
        isLoggedIn = false
        loginViewModel.user = nil
        loginViewModel.userId = nil
        loginViewModel.userRole = nil
        selectedTab = 0
        navigateToLogin = true
    }
    
    private func resetFields() {
        loadUserData()
    }
    
    private func saveChanges() {
        if var updatedUser = loginViewModel.user {
            updatedUser.email = editableEmail
            updatedUser.firstname = editableFirstname
            updatedUser.lastname = editableLastname
            loginViewModel.user = updatedUser
            loginViewModel.saveUser(updatedUser)
        }
    }
    
    private func loadUserData() {
        editableEmail = loginViewModel.user?.email ?? ""
        editableFirstname = loginViewModel.user?.firstname ?? ""
        editableLastname = loginViewModel.user?.lastname ?? ""
        loginViewModel.userId = loginViewModel.user?.id
        loginViewModel.userRole = loginViewModel.user?.role
    }
}

struct SettingsView_Previews: PreviewProvider {
    static var previews: some View {
        @State var selectedTab: Int = 0
        SettingsView(selectedTab: $selectedTab)
            .environmentObject(LoginViewModel())
    }
}
