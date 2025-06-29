//
//  LoginView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 05/10/2024.
//

import SwiftUI

struct LoginView: View {
    @StateObject private var viewModel = LoginViewModel()
    @AppStorage("isLoggedIn") var isLoggedIn: Bool = false
    @AppStorage("keepMeSignedIn") var keepMeSignedIn: Bool = false
    @State private var showError = false
    @State private var errorMessage: String?

    var body: some View {
        NavigationStack {
            ZStack {
                VStack {
                    Spacer()
                    VStack {
                        Text("Welcome back!".localized())
                            .font(.system(size: 25))
                            .bold()
                            .padding(.bottom, 5)

                        Text("Please enter your details to sign in.".localized())
                            .font(.system(size: 14))
                            .padding(.bottom, 50)

                        VStack(alignment: .leading, spacing: 15) {
                            CustomTextInput(title: "Email*", placeholder: "Enter your email", text: $viewModel.model.email, isSecure: false)
                            CustomTextInput(title: "Password*", placeholder: "Enter your password", text: $viewModel.model.password, isSecure: true)

                            HStack {
                                Button(action: {
                                    viewModel.model.keepMeSignedIn.toggle()
                                }, label: {
                                    HStack {
                                        Image(systemName: viewModel.model.keepMeSignedIn ? "checkmark.circle.fill" : "circle")
                                        Text("Keep me signed in".localized())
                                            .font(.system(size: 14))
                                    }
                                    .foregroundStyle((Color("btnColor")))
                                })
                                Spacer()
                                Button(action: {}, label: {
                                    Text("Forgot password?".localized())
                                        .font(.system(size: 14))
                                })
                            }

                            Button(action: {
                                showError = false
                                errorMessage = nil
                                Task {
                                    if let validationError = viewModel.validateFields() {
                                        errorMessage = validationError
                                        showError = true
                                    } else {
                                        await viewModel.signIn()
                                    }
                                }
                            }, label: {
                                Text("Sign In".localized())
                                    .frame(maxWidth: .infinity)
                                    .padding()
                                    .background(Color("btnColor"))
                                    .foregroundColor(.white)
                                    .font(.headline)
                                    .cornerRadius(20)
                                    .padding(.top, 50)
                                    .padding(.bottom, 20)
                            })

                            HStack {
                                Text("Don't have an account ?".localized())
                                    .font(.subheadline)
                                NavigationLink(destination: RegisterView()) {
                                    Text("Sign Up".localized())
                                        .font(.subheadline)
                                        .foregroundColor(.blue)
                                        .accessibilityIdentifier("signUpLink")
                                }
                            }
                        }
                        .padding(.horizontal, 40)
                    }
                    Spacer()
                }
                .frame(maxHeight: .infinity)
                .safeAreaInset(edge: .top) {
                    HStack(spacing: 20) {
                        Image("KeyzLogo")
                            .resizable()
                            .frame(width: 50, height: 50)
                            .padding(.bottom, 40)
                        Text("Keyz")
                            .font(.title)
                            .bold()
                            .padding(.bottom, 40)
                        Spacer()
                    }
                    .padding(.leading, 20)
                }

                if showError, let message = errorMessage {
                    ErrorNotificationView(message: message)
                        .onDisappear {
                            showError = false
                            errorMessage = nil
                        }
                }
            }
            .navigationBarBackButtonHidden(true)
            .navigationDestination(isPresented: $viewModel.isLoggedIn) {
                OverviewView()
            }
            .onChange(of: viewModel.loginStatus) {
                if viewModel.loginStatus.contains("Error") {
                    errorMessage = viewModel.loginStatus.replacingOccurrences(of: "Error: ", with: "")
                    showError = true
                } else if viewModel.loginStatus == "Login successful!" {
                    showError = false
                    errorMessage = nil
                }
            }
        }
    }
}

struct LoginView_Previews: PreviewProvider {
    static var previews: some View {
        LoginView()
    }
}
