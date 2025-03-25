//
//  CustomAlert.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 19/01/2025.
//

import SwiftUI

struct CustomAlert: View {
    @Binding var isActive: Bool
    @Binding var textFieldInput: String

    let title: String
    let message: String
    let buttonTitle: String
    let secondaryButtonTitle: String?
    let action: () -> Void
    let secondaryAction: (() -> Void)?

    @State private var offset: CGFloat = 1000

    var body: some View {
        ZStack {
            Color(.black)
                .opacity(0.5)
                .onTapGesture {
                    close()
                }

            VStack {
                Text(title)
                    .font(.title2)
                    .bold()
                    .padding()
                    .foregroundStyle(Color("textColor"))

                Text(message)
                    .font(.body)
                    .multilineTextAlignment(.center)
                    .padding(.bottom, 10)
                    .foregroundStyle(Color("textColor"))

                TextField("".localized(), text: $textFieldInput)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .padding(.horizontal)

                HStack {
                    if let secondaryButtonTitle = secondaryButtonTitle {
                        Button {
                            secondaryAction?()
                            close()
                        } label: {
                            ZStack {
                                RoundedRectangle(cornerRadius: 20)
                                    .foregroundColor(.gray)

                                Text(secondaryButtonTitle)
                                    .font(.system(size: 16, weight: .bold))
                                    .foregroundColor(.white)
                                    .padding()
                            }
                            .frame(height: 50)
                        }
                    }

                    Button {
                        action()
                        close()
                    } label: {
                        ZStack {
                            RoundedRectangle(cornerRadius: 20)
                                .foregroundColor(.red)

                            Text(buttonTitle)
                                .font(.system(size: 16, weight: .bold))
                                .foregroundColor(.white)
                                .padding()
                        }
                        .frame(height: 50)
                    }
                }
                .padding()
            }
            .fixedSize(horizontal: false, vertical: true)
            .padding()
            .background(Color("primaryBackgroundColor"))
            .clipShape(RoundedRectangle(cornerRadius: 20))
            .overlay(alignment: .topTrailing) {
                Button {
                    close()
                } label: {
                    Image(systemName: "xmark")
                        .font(.title2)
                        .fontWeight(.medium)
                }
                .tint(Color("textColor"))
                .padding()
            }
            .shadow(radius: 20)
            .padding(30)
            .offset(x: 0, y: offset)
            .onAppear {
                withAnimation(.spring()) {
                    offset = 0
                }
            }
        }
        .ignoresSafeArea()
    }

    func close() {
        withAnimation(.spring()) {
            offset = 1000
            isActive = false
        }
    }
}

struct CustomAlertTwoButtons: View {
    @Binding var isActive: Bool

    let title: String
    let message: String
    let buttonTitle: String
    let secondaryButtonTitle: String?
    let action: () -> Void
    let secondaryAction: (() -> Void)?

    @State private var offset: CGFloat = 1000

    var body: some View {
        ZStack {
            Color(.black)
                .opacity(0.5)
                .onTapGesture {
                    close()
                }

            VStack {
                Text(title)
                    .font(.title2)
                    .bold()
                    .padding()
                    .foregroundStyle(Color("textColor"))

                Text(message)
                    .font(.body)
                    .multilineTextAlignment(.center)
                    .padding(.bottom, 20)
                    .foregroundStyle(Color("textColor"))

                HStack {
                    if let secondaryButtonTitle = secondaryButtonTitle {
                        Button {
                            secondaryAction?()
                            close()
                        } label: {
                            ZStack {
                                RoundedRectangle(cornerRadius: 20)
                                    .foregroundColor(.gray)

                                Text(secondaryButtonTitle)
                                    .font(.system(size: 16, weight: .bold))
                                    .foregroundColor(.white)
                                    .padding()
                            }
                            .frame(height: 50)
                        }
                    }

                    Button {
                        action()
                        close()
                    } label: {
                        ZStack {
                            RoundedRectangle(cornerRadius: 20)
                                .foregroundColor(.red)

                            Text(buttonTitle)
                                .font(.system(size: 16, weight: .bold))
                                .foregroundColor(.white)
                                .padding()
                        }
                        .frame(height: 50)
                    }
                }
                .padding()
            }
            .fixedSize(horizontal: false, vertical: true)
            .padding()
            .background(Color("primaryBackgroundColor"))
            .clipShape(RoundedRectangle(cornerRadius: 20))
            .overlay(alignment: .topTrailing) {
                Button {
                    close()
                } label: {
                    Image(systemName: "xmark")
                        .font(.title2)
                        .fontWeight(.medium)
                }
                .tint(Color("textColor"))
                .padding()
            }
            .shadow(radius: 20)
            .padding(30)
            .offset(x: 0, y: offset)
            .onAppear {
                withAnimation(.spring()) {
                    offset = 0
                }
            }
        }
        .ignoresSafeArea()
    }

    func close() {
        withAnimation(.spring()) {
            offset = 1000
            isActive = false
        }
    }
}

struct CustomAlertWithTwoTextFields: View {
    @Binding var isActive: Bool

    let title: String
    let message: String
    let buttonTitle: String
    let secondaryButtonTitle: String?
    let action: (String, Int) -> Void
    let secondaryAction: (() -> Void)?

    @State private var offset: CGFloat = 1000
    @State private var textFieldInputString: String = ""
    @State private var textFieldInputInt: String = ""

    var body: some View {
        ZStack {
            Color(.black)
                .opacity(0.5)
                .onTapGesture {
                    close()
                }

            VStack {
                Text(title)
                    .font(.title2)
                    .bold()
                    .padding()
                    .foregroundStyle(Color("textColor"))

                Text(message)
                    .font(.body)
                    .multilineTextAlignment(.center)
                    .padding(.bottom, 20)
                    .foregroundStyle(Color("textColor"))

                TextField("Enter a new furniture name", text: $textFieldInputString)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .padding(.horizontal)

                TextField("Enter the furniture quantity", text: $textFieldInputInt)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .padding(.horizontal)
                    .keyboardType(.numberPad)
                    .onChange(of: textFieldInputInt) {
                        let filtered = textFieldInputInt.filter { "0123456789".contains($0) }
                        if filtered != textFieldInputInt {
                            textFieldInputInt = filtered
                        }
                    }

                HStack {
                    if let secondaryButtonTitle = secondaryButtonTitle {
                        Button {
                            secondaryAction?()
                            close()
                        } label: {
                            ZStack {
                                RoundedRectangle(cornerRadius: 20)
                                    .foregroundColor(.gray)

                                Text(secondaryButtonTitle)
                                    .font(.system(size: 16, weight: .bold))
                                    .foregroundColor(.white)
                                    .padding()
                            }
                            .frame(height: 50)
                        }
                    }

                    Button {
                        let intValue = Int(textFieldInputInt) ?? 0
                        action(textFieldInputString, intValue)
                        close()
                    } label: {
                        ZStack {
                            RoundedRectangle(cornerRadius: 20)
                                .foregroundColor(.red)

                            Text(buttonTitle)
                                .font(.system(size: 16, weight: .bold))
                                .foregroundColor(.white)
                                .padding()
                        }
                        .frame(height: 50)
                    }
                    .disabled(textFieldInputString.isEmpty || textFieldInputInt.isEmpty)
                }
                .padding()
            }
            .fixedSize(horizontal: false, vertical: true)
            .padding()
            .background(Color("primaryBackgroundColor"))
            .clipShape(RoundedRectangle(cornerRadius: 20))
            .overlay(alignment: .topTrailing) {
                Button {
                    close()
                } label: {
                    Image(systemName: "xmark")
                        .font(.title2)
                        .fontWeight(.medium)
                }
                .tint(Color("textColor"))
                .padding()
            }
            .shadow(radius: 20)
            .padding(30)
            .offset(x: 0, y: offset)
            .onAppear {
                withAnimation(.spring()) {
                    offset = 0
                }
            }
        }
        .ignoresSafeArea()
    }

    func close() {
        withAnimation(.spring()) {
            offset = 1000
            isActive = false
        }
    }
}

struct CustomAlert_Previews: PreviewProvider {
    static var previews: some View {
        VStack {
            CustomAlert(
                isActive: .constant(true),
                textFieldInput: .constant(""),
                title: "Alerte avec Saisie",
                message: "Veuillez entrer quelque chose :",
                buttonTitle: "Valider",
                secondaryButtonTitle: "Annuler",
                action: {
                    print("Validation action")
                },
                secondaryAction: {
                    print("Annulation action")
                }
            )

            CustomAlertTwoButtons(
                isActive: .constant(true),
                title: "Alerte Double",
                message: "Ceci est une alerte avec deux boutons.",
                buttonTitle: "Confirmer",
                secondaryButtonTitle: "Annuler",
                action: {
                    print("Confirmation action")
                },
                secondaryAction: {
                    print("Annulation action")
                }
            )
        }
    }
}
