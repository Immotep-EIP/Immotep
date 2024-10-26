//
//  ContentView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 22/09/2024.
//

import SwiftUI

struct ContentView: View {
    @AppStorage("isLoggedIn") private var isLoggedIn: Bool = false

    var body: some View {
        if isLoggedIn {
            OverviewView()
        } else {
            LoginView()
                .onAppear {
                    if TokenStorage.getAccessToken() != nil {
                        isLoggedIn = true
                    }
                }
        }
    }
}
