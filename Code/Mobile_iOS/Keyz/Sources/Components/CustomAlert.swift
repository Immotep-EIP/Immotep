//
//  CustomAlert.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 19/01/2025.
//

import SwiftUI
import Combine

struct CustomAlert: View {
    @Binding var isActive: Bool
    @Binding var textFieldInput: String
    @State private var isClosing: Bool = false
    @State private var keyboardHeight: CGFloat = 0

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
                    .font(.title3)
                    .bold()
                    .padding(.bottom, 10)
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
                                .foregroundColor(Color("LightBlue"))

                            Text(buttonTitle)
                                .font(.system(size: 16, weight: .bold))
                                .foregroundColor(.white)
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
            .shadow(radius: 20)
            .padding(.horizontal, 30)
            .padding(.bottom, keyboardHeight)
            .offset(x: 0, y: offset)
            .onAppear {
                if !isClosing {
                    withAnimation(.spring) {
                        offset = 0
                    }
                }
            }
            .onReceive(keyboardPublisher) { height in
                withAnimation(.easeInOut(duration: 0.2)) {
                    keyboardHeight = height / 2
                }
            }
        }
        .ignoresSafeArea()
    }

    func close() {
        if isClosing { return }
        isClosing = true
        withAnimation(.easeInOut(duration: 0.2)) {
            offset = 1000
        }
        DispatchQueue.main.asyncAfter(deadline: .now() + 0.2) {
            isActive = false
            isClosing = false
            offset = 1000
        }
    }

    private var keyboardPublisher: AnyPublisher<CGFloat, Never> {
        Publishers.Merge(
            NotificationCenter.default
                .publisher(for: UIResponder.keyboardWillShowNotification)
                .map { notification in
                    (notification.userInfo?[UIResponder.keyboardFrameEndUserInfoKey] as? CGRect)?.height ?? 0
                },
            NotificationCenter.default
                .publisher(for: UIResponder.keyboardWillHideNotification)
                .map { _ in 0 }
        )
        .eraseToAnyPublisher()
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
    @State private var isClosing: Bool = false
    @State private var keyboardHeight: CGFloat = 0

    var body: some View {
        ZStack {
            Color(.black)
                .opacity(0.5)
                .onTapGesture {
                    close()
                }

            VStack {
                Text(title)
                    .font(.title3)
                    .bold()
                    .padding(.bottom, 10)
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
                                .foregroundColor(Color("LightBlue"))

                            Text(buttonTitle)
                                .font(.system(size: 16, weight: .bold))
                                .foregroundColor(.white)
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
            .shadow(radius: 20)
            .padding(.horizontal, 30)
            .padding(.bottom, keyboardHeight)
            .offset(x: 0, y: offset)
            .onAppear {
                if !isClosing {
                    withAnimation(.spring) {
                        offset = 0
                    }
                }
            }
            .onReceive(keyboardPublisher) { height in
                withAnimation(.easeInOut(duration: 0.2)) {
                    keyboardHeight = height / 2
                }
            }
        }
        .ignoresSafeArea()
    }

    func close() {
        if isClosing { return }
        isClosing = true
        withAnimation(.easeInOut(duration: 0.2)) {
            offset = 1000
        }
        DispatchQueue.main.asyncAfter(deadline: .now() + 0.2) {
            isActive = false
            isClosing = false
            offset = 1000
        }
    }

    private var keyboardPublisher: AnyPublisher<CGFloat, Never> {
        Publishers.Merge(
            NotificationCenter.default
                .publisher(for: UIResponder.keyboardWillShowNotification)
                .map { notification in
                    (notification.userInfo?[UIResponder.keyboardFrameEndUserInfoKey] as? CGRect)?.height ?? 0
                },
            NotificationCenter.default
                .publisher(for: UIResponder.keyboardWillHideNotification)
                .map { _ in 0 }
        )
        .eraseToAnyPublisher()
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
    @State private var isClosing: Bool = false
    @State private var keyboardHeight: CGFloat = 0

    var body: some View {
        ZStack {
            Color(.black)
                .opacity(0.5)
                .onTapGesture {
                    close()
                }

            VStack {
                Text(title)
                    .font(.title3)
                    .bold()
                    .padding(.bottom, 10)
                    .foregroundStyle(Color("textColor"))

                Text(message)
                    .font(.body)
                    .multilineTextAlignment(.center)
                    .padding(.bottom, 20)
                    .foregroundStyle(Color("textColor"))

                TextField("Enter a new furniture name".localized(), text: $textFieldInputString)
                    .frame(height: 35)
                    .padding(.leading, 10)
                    .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
                    .padding(.horizontal)

                TextField("Enter the furniture quantity".localized(), text: $textFieldInputInt)
                    .frame(height: 35)
                    .padding(.leading, 10)
                    .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
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
                                .foregroundColor(Color("LightBlue"))

                            Text(buttonTitle)
                                .font(.system(size: 16, weight: .bold))
                                .foregroundColor(.white)
                        }
                        .frame(height: 50)
                    }
                    .disabled(textFieldInputString.isEmpty || textFieldInputInt.isEmpty)
                }
                .padding()
            }
            .fixedSize(horizontal: false, vertical: true)
            .padding()
            .background(Color("basicWhiteBlack"))
            .clipShape(RoundedRectangle(cornerRadius: 20))
            .shadow(radius: 20)
            .padding(.horizontal, 30)
            .padding(.bottom, keyboardHeight)
            .offset(x: 0, y: offset)
            .onAppear {
                if !isClosing {
                    withAnimation(.spring) {
                        offset = 0
                    }
                }
            }
            .onReceive(keyboardPublisher) { height in
                withAnimation(.easeInOut(duration: 0.2)) {
                    keyboardHeight = height / 2
                }
            }
        }
        .ignoresSafeArea()
    }

    func close() {
        if isClosing { return }
        isClosing = true
        withAnimation(.easeInOut(duration: 0.2)) {
            offset = 1000
        }
        DispatchQueue.main.asyncAfter(deadline: .now() + 0.2) {
            isActive = false
            isClosing = false
            offset = 1000
        }
    }

    private var keyboardPublisher: AnyPublisher<CGFloat, Never> {
        Publishers.Merge(
            NotificationCenter.default
                .publisher(for: UIResponder.keyboardWillShowNotification)
                .map { notification in
                    (notification.userInfo?[UIResponder.keyboardFrameEndUserInfoKey] as? CGRect)?.height ?? 0
                },
            NotificationCenter.default
                .publisher(for: UIResponder.keyboardWillHideNotification)
                .map { _ in 0 }
        )
        .eraseToAnyPublisher()
    }
}

struct CustomAlertWithTextAndDropdown: View {
    @Binding var isActive: Bool
    @Binding var textFieldInput: String
    @State private var selectedRoomType: String
    @State private var isClosing: Bool = false
    @State private var keyboardHeight: CGFloat = 0

    private let roomTypeMapping: [(apiValue: String, displayName: String)] = [
        ("dressing", "dressing".localized()),
        ("laundryroom", "laundryroom".localized()),
        ("bedroom", "bedroom".localized()),
        ("playroom", "playroom".localized()),
        ("bathroom", "bathroom".localized()),
        ("toilet", "toilet".localized()),
        ("livingroom", "livingroom".localized()),
        ("diningroom", "diningroom".localized()),
        ("kitchen", "kitchen".localized()),
        ("hallway", "hallway".localized()),
        ("balcony", "balcony".localized()),
        ("cellar", "cellar".localized()),
        ("garage", "garage".localized()),
        ("storage", "storage".localized()),
        ("office", "office".localized()),
        ("other", "other".localized())
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
        self._selectedRoomType = State(initialValue: "dressing")
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
                    .font(.title3)
                    .bold()
                    .padding(.bottom, 10)
                    .foregroundStyle(Color("textColor"))

                Text(message)
                    .font(.body)
                    .multilineTextAlignment(.center)
                    .padding(.bottom, 10)
                    .foregroundStyle(Color("textColor"))

                TextField("Enter room name".localized(), text: $textFieldInput)
                    .frame(height: 35)
                    .padding(.leading, 10)
                    .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
                    .padding(.horizontal)

                Picker("Room Type".localized(), selection: $selectedRoomType) {
                    ForEach(roomTypeMapping, id: \.apiValue) { mapping in
                        Text(mapping.displayName.capitalized)
                            .tag(mapping.apiValue)
                    }
                }
                .pickerStyle(MenuPickerStyle())
                .frame(maxWidth: .infinity)
                .padding(.horizontal)
                .overlay(
                    RoundedRectangle(cornerRadius: 8)
                        .stroke(Color.gray.opacity(0.5), lineWidth: 1)
                )
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
                                .foregroundColor(Color("LightBlue"))

                            Text(buttonTitle)
                                .font(.system(size: 16, weight: .bold))
                                .foregroundColor(.white)
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
            .shadow(radius: 20)
            .padding(.horizontal, 30)
            .padding(.bottom, keyboardHeight)
            .offset(x: 0, y: offset)
            .onAppear {
                if !isClosing {
                    withAnimation(.spring) {
                        offset = 0
                    }
                }
            }
            .onReceive(keyboardPublisher) { height in
                withAnimation(.easeInOut(duration: 0.2)) {
                    keyboardHeight = height / 2
                }
            }
        }
        .ignoresSafeArea()
    }

    func close() {
        if isClosing { return }
        isClosing = true
        withAnimation(.easeInOut(duration: 0.2)) {
            offset = 1000
        }
        DispatchQueue.main.asyncAfter(deadline: .now() + 0.2) {
            isActive = false
            isClosing = false
            offset = 1000
        }
    }

    private var keyboardPublisher: AnyPublisher<CGFloat, Never> {
        Publishers.Merge(
            NotificationCenter.default
                .publisher(for: UIResponder.keyboardWillShowNotification)
                .map { notification in
                    (notification.userInfo?[UIResponder.keyboardFrameEndUserInfoKey] as? CGRect)?.height ?? 0
                },
            NotificationCenter.default
                .publisher(for: UIResponder.keyboardWillHideNotification)
                .map { _ in 0 }
        )
        .eraseToAnyPublisher()
    }
}

struct CustomAlert_Previews: PreviewProvider {
    static var previews: some View {
        VStack {
            CustomAlert(
                isActive: .constant(true),
                textFieldInput: .constant(""),
                title: "Alerte avec Saisie".localized(),
                message: "Veuillez entrer quelque chose :".localized(),
                buttonTitle: "Valider".localized(),
                secondaryButtonTitle: "Annuler".localized(),
                action: {
                    print("Validation action")
                },
                secondaryAction: {
                    print("Annulation action")
                }
            )

            CustomAlertTwoButtons(
                isActive: .constant(true),
                title: "Alerte Double".localized(),
                message: "Ceci est une alerte avec deux boutons.".localized(),
                buttonTitle: "Confirmer".localized(),
                secondaryButtonTitle: "Annuler".localized(),
                action: {
                    print("Confirmation action")
                },
                secondaryAction: {
                    print("Annulation action")
                }
            )

            CustomAlertWithTwoTextFields(
                isActive: .constant(true),
                title: "Alerte avec Deux Champs".localized(),
                message: "Entrez un nom et une quantité :".localized(),
                buttonTitle: "Ajouter".localized(),
                secondaryButtonTitle: "Annuler".localized(),
                action: { name, quantity in
                    print("Nom: \(name), Quantité: \(quantity)")
                },
                secondaryAction: {
                    print("Annulation action")
                }
            )

            CustomAlertWithTextAndDropdown(
                isActive: .constant(true),
                textFieldInput: .constant(""),
                title: "Alerte avec Texte et Dropdown".localized(),
                message: "Entrez un nom et sélectionnez un type :".localized(),
                buttonTitle: "Ajouter".localized(),
                secondaryButtonTitle: "Annuler".localized(),
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
