//
//  RegisterView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 05/10/2024.
//

import SwiftUI

struct RegisterView: View {
    @StateObject private var viewModel = RegisterViewModel()
    @State private var showAlert = false
    @State private var alertMessage: String?
    @State private var isRegistrationSuccessful = false
    @State private var successAlert = false
    @State private var isLoading: Bool = false
    @Environment(\.dismiss) var dismiss

    var body: some View {
        NavigationStack {
            ZStack {
                VStack {
                    Spacer()
                    VStack {
                        Text("Create your account".localized())
                            .font(.system(size: 25))
                            .bold()
                            .padding(.bottom, 5)
                        Text("Join Keyz for your peace of mind!".localized())
                            .font(.system(size: 14))
                            .padding(.bottom, 50)

                        VStack(alignment: .leading, spacing: 5) {
                            CustomTextInput(title: "First name*", placeholder: "Enter your first name", text: $viewModel.model.firstName, isSecure: false)
                            CustomTextInput(title: "Name*", placeholder: "Enter your name", text: $viewModel.model.name, isSecure: false)
                            CustomTextInput(title: "Email*", placeholder: "Enter your email", text: $viewModel.model.email, isSecure: false)
                            CustomTextInput(title: "Password*", placeholder: "Enter your password", text: $viewModel.model.password, isSecure: true)
                            CustomTextInput(title: "Password confirmation*", placeholder: "Enter your password confirmation", text: $viewModel.model.passwordConfirmation, isSecure: true)

                            HStack {
                                Button(action: {
                                    viewModel.model.agreement.toggle()
                                }, label: {
                                    HStack {
                                        Image(systemName: viewModel.model.agreement ? "checkmark.circle.fill" : "circle")
                                        Text("I agree to all Term, Privacy Policy and Fees".localized())
                                            .font(.system(size: 12))
                                            .multilineTextAlignment(.leading)
                                    }
                                    .foregroundStyle((Color("textColor")))
                                })
                                .accessibilityIdentifier("AgreementButton")
                                .padding(.top, 10)

                                Spacer()
                            }

                            Button(action: {
                                showAlert = false
                                alertMessage = nil
                                isLoading = true
                                Task {
                                    await viewModel.signIn()
                                    isLoading = false
                                }
                            }, label: {
                                ZStack {
                                    if isLoading {
                                        ProgressView()
                                            .progressViewStyle(CircularProgressViewStyle())
                                            .tint(.white)
                                    } else {
                                        Text("Sign Up".localized())
                                    }
                                }
                                .frame(maxWidth: .infinity)
                                .padding()
                                .background(viewModel.model.name.isEmpty || viewModel.model.firstName.isEmpty || viewModel.model.email.isEmpty || viewModel.model.password.isEmpty || viewModel.model.passwordConfirmation.isEmpty || !viewModel.model.agreement ? Color.gray : Color("btnColor"))
                                .foregroundColor(.white)
                                .font(.headline)
                                .cornerRadius(20)
                                .padding(.top, 30)
                                .padding(.bottom, 10)
                                .scaleEffect(isLoading ? 0.95 : 1.0)
                                .animation(.easeInOut(duration: 0.2), value: isLoading)
                            })
                            .accessibilityIdentifier("SignUpButton")
                            .disabled(isLoading || viewModel.model.name.isEmpty || viewModel.model.firstName.isEmpty || viewModel.model.email.isEmpty || viewModel.model.password.isEmpty || viewModel.model.passwordConfirmation.isEmpty || !viewModel.model.agreement)

                            HStack {
                                Text("Already have an account ?".localized())
                                    .font(.subheadline)
                                NavigationLink(destination: LoginView()) {
                                    Text("Log In".localized())
                                        .font(.subheadline)
                                        .foregroundColor(.blue)
                                }
                            }
                        }
                        .padding(.horizontal, 40)
                    }
                    Spacer()
                }
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

                if showAlert, let message = alertMessage {
                    ErrorNotificationView(message: message, type: successAlert ? .success : .error)
                        .onDisappear {
                            showAlert = false
                            alertMessage = nil
                        }
                }
            }
            .navigationBarBackButtonHidden(true)
            .navigationDestination(isPresented: $isRegistrationSuccessful) {
                LoginView()
            }
            .onChange(of: viewModel.registerStatus) {
                if viewModel.registerStatus.lowercased() == "registration successful!" {
                    successAlert = true
                    alertMessage = viewModel.registerStatus.localized()
                    showAlert = true
                    DispatchQueue.main.asyncAfter(deadline: .now() + 2.0) {
                        dismiss()
                    }
                } else if !viewModel.registerStatus.isEmpty {
                    alertMessage = viewModel.registerStatus.localized()
                    showAlert = true
                    isRegistrationSuccessful = false
                }
            }
        }
    }
}

struct RegisterView_Previews: PreviewProvider {
    static var previews: some View {
        RegisterView()
    }
}
