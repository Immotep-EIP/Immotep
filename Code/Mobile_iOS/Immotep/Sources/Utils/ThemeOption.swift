//
//  Appearance.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 17/11/2024.
//

import SwiftUI

enum ThemeOption: String, CaseIterable {
    case system = "System"
    case light = "Light"
    case dark = "Dark"
}

struct ThemeManager {
    @MainActor
    static func applyTheme(theme: String) {
        guard let windowScene = UIApplication.shared.connectedScenes.first as? UIWindowScene else { return }
        let rootViewController = windowScene.windows.first?.rootViewController

        switch theme {
        case ThemeOption.light.rawValue:
            rootViewController?.overrideUserInterfaceStyle = .light
        case ThemeOption.dark.rawValue:
            rootViewController?.overrideUserInterfaceStyle = .dark
        default:
            rootViewController?.overrideUserInterfaceStyle = .unspecified
        }
    }
}
