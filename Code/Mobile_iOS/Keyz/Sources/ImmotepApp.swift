//
//  ImmotepApp.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 22/09/2024.
//

import SwiftUI
import Foundation
import UIKit

@main
struct ImmotepApp: App {
    @UIApplicationDelegateAdaptor(AppDelegate.self) private var appdelegate
    @AppStorage("theme") private var selectedTheme: String = ThemeOption.system.rawValue

    init() {
        if CommandLine.arguments.contains("--UITests") || CommandLine.arguments.contains("-skipLogin") {
            UIView.setAnimationsEnabled(false)
        }
    }

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
