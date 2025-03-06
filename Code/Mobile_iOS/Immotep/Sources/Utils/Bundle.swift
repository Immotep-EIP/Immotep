//
//  Bundle.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import Foundation
import SwiftUI

extension String {
    func localized() -> String {
        let lang = UserDefaults.standard.string(forKey: "lang") ?? "en"
        if let path = Bundle.main.path(forResource: lang, ofType: "lproj"),
           let bundle = Bundle(path: path) {
            return NSLocalizedString(self, tableName: nil, bundle: bundle, value: "", comment: "")
        }
        return self
    }
}

@MainActor
extension Bundle {
    static func setLanguage(_ language: String) async {
        UserDefaults.standard.set(language, forKey: "lang")
        UserDefaults.standard.synchronize()
    }
}
