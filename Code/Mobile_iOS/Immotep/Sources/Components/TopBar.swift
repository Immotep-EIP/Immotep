//
//  TopBar.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import SwiftUI

struct TopBar: View {
    var title: String
    var body: some View {
        VStack(spacing: 0) {
            HStack {
                Image("immotepLogo")
                    .resizable()
                    .frame(width: 50, height: 50)

                Text(title)
                    .font(.title)
                    .bold()
                    .lineLimit(1)
                    .frame(maxWidth: .infinity, alignment: .leading)

                NavigationLink(destination: ProfileView()) {
                    Image(systemName: "person.crop.circle.fill")
                        .frame(width: 70, height: 50)
                        .font(.system(size: 34))
                        .foregroundStyle(Color("textColor"))
                }
                .buttonStyle(PlainButtonStyle())
            }
            .padding(.leading, 20)

            Divider()
                .background(Color("textColor"))
                .padding(.top, 15)
        }
    }
}
