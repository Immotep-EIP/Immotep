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
                .font(.system(size: 14))
                .frame(alignment: .leading)

            if isSecure {
                SecureField(placeholder.localized(), text: $text)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .font(.system(size: 14))
                    .autocapitalization(.none)
            } else {
                TextField(placeholder.localized(), text: $text)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .font(.system(size: 14))
                    .autocapitalization(.none)
            }
        }
    }
}
