//
//  RegisterView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 05/10/2024.
//

import SwiftUI

struct RegisterView: View {
    @StateObject private var viewModel = RegisterViewModel()

    var body: some View {
        NavigationStack {
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

                        Text(viewModel.registerStatus.localized())
                            .font(.subheadline)
                            .foregroundColor(viewModel.registerStatus == "Registration successful!" ? .green : .red)
                            .padding(.top, 10)

                        Button(action: {
                            Task {
                                await viewModel.signIn()
                            }
                        }, label: {
                            Text("Sign Up".localized())
                                .frame(maxWidth: .infinity)
                                .padding()
                                .background(Color("btnColor"))
                                .foregroundColor(.white)
                                .font(.headline)
                                .cornerRadius(20)
                                .padding(.top, 30)
                                .padding(.bottom, 10)
                        })
                        .accessibilityIdentifier("SignUpButton")

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
        }
        .navigationBarBackButtonHidden(true)
    }
}

struct RegisterView_Previews: PreviewProvider {
    static var previews: some View {
        RegisterView()
    }
}
