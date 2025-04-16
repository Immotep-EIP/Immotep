//
//  ContentView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 22/09/2024.
//

import SwiftUI

struct ContentView: View {
    @AppStorage("isLoggedIn") var isLoggedIn: Bool = false
    @StateObject private var loginViewModel = LoginViewModel()
    @StateObject private var propertyViewModel = PropertyViewModel()
    @AppStorage("lang") var lang: String = "en"
    @State private var selectedTab: Int = 0

    var body: some View {
        if isLoggedIn {
            TabView(selection: $selectedTab) {
                OverviewView()
                    .environmentObject(loginViewModel)
                    .tabItem {
                        Image(systemName: "house")
                        Text("Overview".localized())
                    }
                    .tag(0)

                PropertyView()
                    .environmentObject(propertyViewModel)
                    .tabItem {
                        Image(systemName: "building.2")
                        Text("Real Property".localized())
                    }
                    .tag(1)

                MessagesView()
                    .tabItem {
                        Image(systemName: "envelope")
                        Text("Messages".localized())
                    }
                    .tag(2)

                SettingsView()
                    .environmentObject(loginViewModel)
                    .tabItem {
                        Image(systemName: "gearshape")
                        Text("Settings".localized())
                    }
                    .tag(3)
            }
            .id(lang)
        } else {
            LoginView()
                .environmentObject(loginViewModel)
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
