//
//  ImmotepApp.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 22/09/2024.
//

import SwiftUI

@main
struct ImmotepApp: App {
    @UIApplicationDelegateAdaptor(AppDelegate.self) private var appdelegate
    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
}
