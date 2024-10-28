//
//  ContentView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 22/09/2024.
//

import SwiftUI

struct ContentView: View {
    @AppStorage("isLoggedIn") var isLoggedIn: Bool = false
    @StateObject private var profileViewModel = ProfileViewModel()

    var body: some View {
        if isLoggedIn {
            OverviewView()
                .environmentObject(profileViewModel)
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
