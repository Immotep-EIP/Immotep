//
//  TopBar.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import SwiftUI

struct TopBar: View {
    var body: some View {
        VStack(spacing: 0) {
            HStack(spacing: 20) {
                Image("immotepLogo")
                    .resizable()
                    .frame(width: 50, height: 50)

                Text("Immotep")
                    .font(.title)
                    .bold()

                Spacer()

                NavigationLink(destination: ProfileView()) {
                    Image(systemName: "person.crop.circle.fill")
                        .frame(width: 70, height: 50)
                        .font(.system(size: 34))
                        .foregroundStyle(Color.black)
                }
                .buttonStyle(PlainButtonStyle())
            }
            .padding(.leading, 20)

            Divider()
                .background(Color.black)
                .padding(.top, 15)
        }
    }
}
