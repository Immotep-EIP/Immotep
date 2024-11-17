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
let baseURL = URL(string: "https://test1.icytree-5b429d30.eastus.azurecontainerapps.io/")!

@main
struct ImmotepApp: App {
    @UIApplicationDelegateAdaptor(AppDelegate.self) private var appdelegate

    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
}
