//
//  MessagesView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import SwiftUI

struct MessagesView: View {
    var body: some View {
        NavigationStack {
            VStack {
                TopBar(title: "Messages".localized())
                Spacer()

                TaskBar()
            }
            .navigationBarBackButtonHidden(true)
            .navigationTransition(
                .fade(.in).animation(.easeInOut(duration: 0))
            )
        }
    }
}

struct MessagesView_Previews: PreviewProvider {
    static var previews: some View {
        MessagesView()
    }
}
