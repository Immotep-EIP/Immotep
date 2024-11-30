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

    var body: some View {
        VStack(alignment: .leading, spacing: 5) {
            Text(title.localized())

            if isSecure {
                SecureField(placeholder.localized(), text: $text)
                    .padding(8)
                    .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
                    .foregroundStyle(Color("placeholderColor"))

            } else {
                TextField(placeholder.localized(), text: $text)
                    .padding(8)
                    .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
                    .foregroundStyle(Color("placeholderColor"))

            }
        }
        .font(.system(size: 14))
        .autocapitalization(.none)
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
                        return value.map { String(format: "%.2f", $0.doubleValue) } ?? ""
                    },
                    set: { newValue in
                        if let doubleValue = Double(newValue) {
                            value = NSNumber(value: doubleValue)
                        } else {
                            value = nil
                        }
                    }
                ))
                .padding(8)
                .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
                .foregroundStyle(Color("placeholderColor"))
            } else {
                TextField(placeholder.localized(), text: Binding(
                    get: {
                        return value.map { String(format: "%.2f", $0.doubleValue) } ?? ""
                    },
                    set: { newValue in
                        if let doubleValue = Double(newValue) {
                            value = NSNumber(value: doubleValue)
                        } else {
                            value = nil
                        }
                    }
                ))
                .padding(8)
                .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
                .foregroundStyle(Color("placeholderColor"))
            }
        }
        .font(.system(size: 14))
        .autocapitalization(.none)
        .keyboardType(.decimalPad)
    }
}
