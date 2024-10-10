//
//  LoginView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 05/10/2024.
//
import SwiftUI

struct LoginView: View {
    @StateObject private var viewModel = LoginViewModel()

    var body: some View {
        NavigationStack {
            VStack {
                Spacer()
                VStack {
                    Text("Welcome back")
                        .font(.system(size: 25))
                        .bold()
                        .padding(.bottom, 5)
                    Text("Please enter your details to sign in.")
                        .font(.system(size: 14))
                        .padding(.bottom, 50)

                    VStack(alignment: .leading, spacing: 5) {
                        Text("Email*")
                            .font(.system(size: 14))
                            .frame(alignment: .leading)

                        TextField("Enter your email", text: $viewModel.model.email)
                            .textFieldStyle(RoundedBorderTextFieldStyle())
                            .font(.system(size: 14))
                            .padding(.bottom, 20)
                            .keyboardType(.emailAddress)

                        Text("Password*")
                            .font(.system(size: 14))
                            .frame(alignment: .leading)

                        SecureField("Enter your password", text: $viewModel.model.password)
                            .textFieldStyle(RoundedBorderTextFieldStyle())
                            .font(.system(size: 14))
                            .padding(.bottom, 20)

                        HStack {
                            Button(action: {
                                viewModel.model.keepMeSignedIn.toggle()
                            }, label: {
                                HStack {
                                    Image(systemName: viewModel.model.keepMeSignedIn ? "checkmark.circle.fill" : "circle")
                                    Text("Keep me signed in")
                                        .font(.system(size: 14))
                                }
                                .foregroundStyle(.black)
                            })
                            Spacer()

                            Button(action: {
                                // create the forgot password action when possible
                            }, label: {
                                Text("Forgot password?")
                                    .font(.system(size: 14))
                            })
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
                                .padding(.top, 50)
                                .padding(.bottom, 20)
                        })

                        HStack {
                            Text("Don’t have an account ?")
                                .font(.subheadline)
                            NavigationLink(destination: RegisterView()) {
                                Text("Sign Up")
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
            /*VStack {
                Spacer()
                // Fix this, not working now
                Text(viewModel.loginStatus)
                    .foregroundColor(viewModel.loginStatus == "Login successful!" ? .green : .red)
                    .padding(.top, 10)
                .padding(.top, 20)
                Spacer()
            }
            .padding()*/
        }
        .navigationBarBackButtonHidden(true)
    }
}

struct LoginView_Previews: PreviewProvider {
    static var previews: some View {
        LoginView()
    }
}
