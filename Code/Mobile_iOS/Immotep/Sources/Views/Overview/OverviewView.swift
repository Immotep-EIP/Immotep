//
//  OverviewView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 21/10/2024.
//

import SwiftUI

struct OverviewView: View {
    var body: some View {
        NavigationStack {
            VStack {
                Spacer()
                    .safeAreaInset(edge: .top) {
                        HStack(spacing: 20) {
                            Image("immotepLogo")
                                .resizable()
                                .frame(width: 50, height: 50)
                            Text("Immotep")
                                .font(.title)
                                .bold()
                            Spacer()

                            Button(action: {
                            }, label: {
                                Image(systemName: "person.crop.circle.fill")
                                    .frame(width: 70, height: 50)
                                    .font(.system(size: 34))
                                    .foregroundStyle(Color.black)
                            })
                        }
                        .padding(.leading, 20)
                        .padding(.bottom, 40)

                        Divider()
                            .background(Color.black)
                            .padding(.top, 30)
                    }

                TaskBar()
            }
            .navigationBarBackButtonHidden(true)
        }
    }
}

struct OverviewView_Previews: PreviewProvider {
    static var previews: some View {
        OverviewView()
    }
}
