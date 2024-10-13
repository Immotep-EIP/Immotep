//
//  LoginView.swift
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
                    Text("Create your account")
                        .font(.system(size: 25))
                        .bold()
                        .padding(.bottom, 5)
                    Text("Join Immotep for your peace of mind!")
                        .font(.system(size: 14))
                        .padding(.bottom, 50)

                    VStack(alignment: .leading, spacing: 5) {

                        CustomTextInput(title: "Name*", placeholder: "Enter your name", text: $viewModel.model.name, isSecure: false)

                        CustomTextInput(title: "First name*", placeholder: "Enter your first name", text: $viewModel.model.firstName, isSecure: false)

                        CustomTextInput(title: "Email*", placeholder: "Enter your email", text: $viewModel.model.email, isSecure: false)

                        CustomTextInput(title: "Password*", placeholder: "Enter your password", text: $viewModel.model.password, isSecure: true)

                        CustomTextInput(title: "Password confirmation*", placeholder: "Enter your password confirmation", text: $viewModel.model.passwordConfirmation, isSecure: true)

                        HStack {
                            Button(action: {
                                viewModel.model.agreement.toggle()
                            }, label: {
                                HStack {
                                    Image(systemName: viewModel.model.agreement ? "checkmark.circle.fill" : "circle")
                                    Text("I agree to all Term, Privacy Policy and Fees")
                                        .font(.system(size: 14))
                                }
                                .foregroundStyle(.black)
                            })
                            Spacer()
                        }

                        Button(action: {
                            viewModel.signIn() // Not doing the job for now
                        }, label: {
                            Text("Sign In")
                                .frame(maxWidth: .infinity)
                                .padding()
                                .background(Color.black)
                                .foregroundColor(.white)
                                .font(.headline)
                                .cornerRadius(20)
                                .padding(.top, 30)
                                .padding(.bottom, 10)
                        })

                        HStack {
                            Text("Already have an account ?")
                                .font(.subheadline)
                            NavigationLink(destination: LoginView()) {
                                Text("Log In")
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
                    Image("immotepLogo")
                        .resizable()
                        .frame(width: 50, height: 50)
                        .padding(.bottom, 40)
                    Text("Immotep")
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
