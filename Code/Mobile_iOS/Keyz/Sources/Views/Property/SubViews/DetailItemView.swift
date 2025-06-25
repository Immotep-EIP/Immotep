//
//  DetailItemView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/06/2025.
//

import SwiftUI

struct DetailItemView: View {
    let icon: String
    let value: String
    let label: String

    var body: some View {
        VStack(spacing: 5) {
            Image(systemName: icon)
                .font(.title3)
                .foregroundColor(Color("LightBlue"))
            Text(value)
                .font(.headline)
                .foregroundColor(Color("textColor"))
            Text(label)
                .font(.caption)
                .foregroundColor(.gray)
                .multilineTextAlignment(.center)
        }
        .padding()
        .background(Color("basicWhiteBlack"))
        .cornerRadius(10)
        .shadow(color: Color.black.opacity(0.1), radius: 10, x: 0, y: 2)
    }
}

struct DetailItemView_Previews: PreviewProvider {
    static var previews: some View {
        DetailItemView(icon: "square.split.bottomrightquarter", value: "80m2", label: "Area".localized())
    }
}
