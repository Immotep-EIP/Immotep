//
//  ImmotepApp.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 22/09/2024.
//

import SwiftUI
import Foundation

// only for local test purpose
// let baseURL = URL(string: "http://localhost:3001/api/v1")!
// only for online test purpose
// let baseURL = URL(string: "https://dev.backend.immotep.mazettt.fr/api/v1")!

@main
struct ImmotepApp: App {
    @UIApplicationDelegateAdaptor(AppDelegate.self) private var appdelegate
    @AppStorage("theme") private var selectedTheme: String = ThemeOption.system.rawValue

    var body: some Scene {
        let isUITestMode = CommandLine.arguments.contains("-skipLogin")
        WindowGroup {
            if isUITestMode {
                TestImmotepView()
                    .onAppear {
                        Task { @MainActor in
                            ThemeManager.applyTheme(theme: selectedTheme)
                        }
                    }
            } else {
                ContentView()
                    .onAppear {
                        Task { @MainActor in
                            ThemeManager.applyTheme(theme: selectedTheme)
                        }
                    }
            }
        }
    }
}
