//
//  TestImmotepApp.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 09/12/2024.
//

import SwiftUI

struct TestImmotepView: View {
    @AppStorage("isLoggedIn") var isLoggedIn: Bool = false
    @StateObject private var profileViewModel = ProfileViewModel()

    var body: some View {
        let isUITestMode = CommandLine.arguments.contains("-skipLogin")
        if isLoggedIn || isUITestMode {
            if CommandLine.arguments.contains("-propertyList") {
                PropertyView()
            } else {
                OverviewView()
                    .environmentObject(profileViewModel)
            }
        } else {
            if isLoggedIn {
                OverviewView()
                    .environmentObject(profileViewModel)
            } else {
                LoginView()
            }
        }
    }
}
