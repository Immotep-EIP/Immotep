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
                TopBar(title: "Keyz".localized())
                Spacer()
                Text("Messaging features will be available soon!".localized())
                Spacer()
            }
            .navigationBarBackButtonHidden(true)
        }
    }
}

struct MessagesView_Previews: PreviewProvider {
    static var previews: some View {
        MessagesView()
    }
}
