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
                    .frame(height: 35)
                    .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
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
            .background(Color("basicWhiteBlack"))
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
            .background(Color("basicWhiteBlack"))
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

    struct CustomAlertWithTextAndDropdown: View {
        @Binding var isActive: Bool
        @Binding var textFieldInput: String
        @State private var selectedRoomType: String
        let roomTypes = [
            "dressing",
            "laundryroom",
            "bedroom",
            "playroom",
            "bathroom",
            "toilet",
            "livingroom",
            "diningroom",
            "kitchen",
            "hallway",
            "balcony",
            "cellar",
            "garage",
            "storage",
            "office",
            "other"
        ]

        let title: String
        let message: String
        let buttonTitle: String
        let secondaryButtonTitle: String?
        let action: (String, String) -> Void
        let secondaryAction: (() -> Void)?

        @State private var offset: CGFloat = 1000

        init(
            isActive: Binding<Bool>,
            textFieldInput: Binding<String>,
            title: String,
            message: String,
            buttonTitle: String,
            secondaryButtonTitle: String?,
            action: @escaping (String, String) -> Void,
            secondaryAction: (() -> Void)?
        ) {
            self._isActive = isActive
            self._textFieldInput = textFieldInput
            self.title = title
            self.message = message
            self.buttonTitle = buttonTitle
            self.secondaryButtonTitle = secondaryButtonTitle
            self.action = action
            self.secondaryAction = secondaryAction
            self._selectedRoomType = State(initialValue: roomTypes.first ?? "other")
        }

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

                    TextField("Enter room name", text: $textFieldInput)
                        .frame(height: 35)
                        .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
                        .padding(.horizontal)

                    Picker("Room Type", selection: $selectedRoomType) {
                        ForEach(roomTypes, id: \.self) { type in
                            Text(type.capitalized)
                                .tag(type)
                        }
                    }
                    .pickerStyle(MenuPickerStyle())
                    .padding(.horizontal)
                    .padding(.vertical, 5)
                    .overlay(
                        RoundedRectangle(cornerRadius: 8)
                            .stroke(Color.gray.opacity(0.5), lineWidth: 1)
                    )

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
                            action(textFieldInput, selectedRoomType)
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
                        .disabled(textFieldInput.isEmpty)
                    }
                    .padding()
                }
                .fixedSize(horizontal: false, vertical: true)
                .padding()
                .background(Color("basicWhiteBlack"))
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

            CustomAlertWithTextAndDropdown(
                isActive: .constant(true),
                textFieldInput: .constant(""),
                title: "Alerte avec Texte et Dropdown",
                message: "Entrez un nom et sélectionnez un type :",
                buttonTitle: "Ajouter",
                secondaryButtonTitle: "Annuler",
                action: { name, type in
                    print("Nom: \(name), Type: \(type)")
                },
                secondaryAction: {
                    print("Annulation action")
                }
            )
        }
    }
}
