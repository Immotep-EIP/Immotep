//
//  CustomTextInput.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 13/10/2024.
//

import SwiftUI

struct CustomTextInput: View {
    var title: String
    var placeholder: String
    @Binding var text: String
    var isSecure: Bool = false

    @State private var isSecured: Bool = true

    var body: some View {
        VStack(alignment: .leading, spacing: 5) {
            Text(title.localized())

            ZStack(alignment: .trailing) {
                Group {
                    if isSecure && isSecured {
                        SecureField(placeholder.localized(), text: $text)
                            .accessibilityIdentifier("\(title)_textfield")
                    } else {
                        TextField(placeholder.localized(), text: $text)
                            .accessibilityIdentifier("\(title)_textfield")
                    }
                }
                .padding(8)
                .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
                .foregroundStyle(Color("placeholderColor"))
                .autocapitalization(.none)

                if isSecure {
                    Button(action: {
                        isSecured.toggle()
                    }, label: {
                        Image(systemName: isSecured ? "eye.slash" : "eye")
                            .accentColor(.gray)
                            .padding(.trailing, 8)
                    })
                }
            }
        }
        .font(.system(size: 14))
    }
}

struct CustomTextInputNB: View {
    var title: String
    var placeholder: String
    @Binding var value: NSNumber?
    var isSecure: Bool = false

    var body: some View {
        VStack(alignment: .leading, spacing: 5) {
            Text(title.localized())

            if isSecure {
                SecureField(placeholder.localized(), text: Binding(
                    get: {
                        guard let value = value else { return "" }
                        if value.doubleValue == floor(value.doubleValue) {
                            return "\(value.intValue)"
                        } else {
                            return String(format: "%.2f", value.doubleValue)
                        }
                    },
                    set: { newValue in
                        if let intValue = Int(newValue) {
                            value = NSNumber(value: intValue)
                        } else if let doubleValue = Double(newValue) {
                            value = NSNumber(value: doubleValue)
                        } else {
                            value = nil
                        }
                    }
                ))
                .padding(8)
                .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
                .foregroundStyle(Color("placeholderColor"))
                .accessibilityIdentifier("\(title)_textfield")
            } else {
                TextField(placeholder.localized(), text: Binding(
                    get: {
                        guard let value = value else { return "" }
                        if value.doubleValue == floor(value.doubleValue) {
                            return "\(value.intValue)"
                        } else {
                            return String(format: "%.2f", value.doubleValue)
                        }
                    },
                    set: { newValue in
                        if let intValue = Int(newValue) {
                            value = NSNumber(value: intValue)
                        } else if let doubleValue = Double(newValue) {
                            value = NSNumber(value: doubleValue)
                        } else {
                            value = nil
                        }
                    }
                ))
                .padding(8)
                .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
                .foregroundStyle(Color("placeholderColor"))
                .accessibilityIdentifier("\(title)_textfield")

            }
        }
        .font(.system(size: 14))
        .autocapitalization(.none)
        .keyboardType(.decimalPad)
    }
}
