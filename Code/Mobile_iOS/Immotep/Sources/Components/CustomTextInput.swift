//
//  CustomTextInput.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 13/10/2024.
//

// Custom text input for now used in login and register views, see if works well with UI tests

import SwiftUI

struct CustomTextInput: View {
    var title: String
    var placeholder: String
    @Binding var text: String
    var isSecure: Bool = false

    var body: some View {
        VStack(alignment: .leading, spacing: 5) {
            Text(title)
                .font(.system(size: 14))
                .frame(alignment: .leading)

            if isSecure {
                SecureField(placeholder, text: $text)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .font(.system(size: 14))
                    .autocapitalization(.none)
            } else {
                TextField(placeholder, text: $text)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .font(.system(size: 14))
                    .autocapitalization(.none)
            }
        }
    }
}
