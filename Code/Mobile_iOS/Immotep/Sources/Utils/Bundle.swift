//
//  Bundle.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import Foundation
import SwiftUI

extension String {
  func localizeString(string: String) -> String {

      let path = Bundle.main.path(forResource: string, ofType: "lproj")
      let bundle = Bundle(path: path!)
      return NSLocalizedString(self, tableName: nil, bundle: bundle!,
      value: "", comment: "")
  }
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
    private static var customBundle: Bundle?

    static func setLanguage(_ language: String) async {
        UserDefaults.standard.set([language], forKey: "AppleLanguages")
        UserDefaults.standard.synchronize()

        if let path = Bundle.main.path(forResource: language, ofType: "lproj") {
            customBundle = Bundle(path: path)
        } else {
            customBundle = Bundle.main
        }
    }

    static func localizedString(forKey key: String) -> String {
        return (customBundle ?? Bundle.main).localizedString(forKey: key, value: nil, table: nil)
    }
}
