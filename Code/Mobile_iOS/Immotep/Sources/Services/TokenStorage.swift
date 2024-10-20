//
//  TokenStorage.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 20/10/2024.
//

import Foundation

struct TokenStorage {
    static func storeTokens(accessToken: String, refreshToken: String) {
        UserDefaults.standard.set(accessToken, forKey: "access_token")
        UserDefaults.standard.set(refreshToken, forKey: "refresh_token")
    }

    static func storeAccessToken(_ accessToken: String) {
        UserDefaults.standard.set(accessToken, forKey: "access_token")
    }

    static func getAccessToken() -> String? {
        return UserDefaults.standard.string(forKey: "access_token")
    }

    static func getRefreshToken() -> String? {
        return UserDefaults.standard.string(forKey: "refresh_token")
    }

    static func clearTokens() {
        UserDefaults.standard.removeObject(forKey: "access_token")
        UserDefaults.standard.removeObject(forKey: "refresh_token")
    }

    static func extractTokens(from data: Data) throws -> (String, String) {
        if let json = try JSONSerialization.jsonObject(with: data, options: []) as? [String: Any] {
            let accessToken = json["access_token"] as? String ?? "Unknown Access Token"
            let refreshToken = json["refresh_token"] as? String ?? "Unknown Refresh Token"
            return (accessToken, refreshToken)
        } else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to parse response for tokens."])
        }
    }
}
