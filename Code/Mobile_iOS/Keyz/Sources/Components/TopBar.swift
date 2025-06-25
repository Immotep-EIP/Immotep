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
                Image("KeyzLogo")
                    .resizable()
                    .frame(width: 50, height: 50)

                Text(title)
                    .font(.title2)
                    .bold()
                    .lineLimit(1)
                    .frame(maxWidth: .infinity, alignment: .leading)
            }
            .padding(.leading, 20)

            Divider()
                .background(Color("textColor"))
                .padding(.top, 15)
        }
    }
}

extension UINavigationController {
    override open func viewDidLoad() {
        super.viewDidLoad()
        interactivePopGestureRecognizer?.delegate = nil
    }
}
