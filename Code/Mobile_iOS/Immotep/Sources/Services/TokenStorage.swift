//
//  TokenStorage.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 20/10/2024.
//

import Foundation

struct TokenStorage {
    private static let accessTokenKey = "access_token"
    private static let refreshTokenKey = "refresh_token"
    private static let tokenDateKey = "token_date"
    private static let tokenExpiryKey = "token_expiry"

    static func storeTokens(accessToken: String, refreshToken: String, expiresIn: TimeInterval? = nil) {
        let currentDate = Date()
        UserDefaults.standard.set(accessToken, forKey: accessTokenKey)
        UserDefaults.standard.set(refreshToken, forKey: refreshTokenKey)
        UserDefaults.standard.set(currentDate, forKey: tokenDateKey)

        if let expiresIn = expiresIn {
            UserDefaults.standard.set(expiresIn, forKey: tokenExpiryKey)
        }
    }

    static func storeAccessToken(_ accessToken: String) {
        UserDefaults.standard.set(accessToken, forKey: accessTokenKey)
        UserDefaults.standard.set(Date(), forKey: tokenDateKey)
    }

    static func getAccessToken() -> String? {
        return UserDefaults.standard.string(forKey: accessTokenKey)
    }

    static func getRefreshToken() -> String? {
        return UserDefaults.standard.string(forKey: refreshTokenKey)
    }

    static func isTokenExpired() -> Bool {
        guard let tokenDate = UserDefaults.standard.object(forKey: tokenDateKey) as? Date,
              let expiresIn = UserDefaults.standard.object(forKey: tokenExpiryKey) as? TimeInterval else {
            return true
        }
        return Date().timeIntervalSince(tokenDate) >= expiresIn
    }

    static func clearTokens() {
        UserDefaults.standard.removeObject(forKey: accessTokenKey)
        UserDefaults.standard.removeObject(forKey: refreshTokenKey)
        UserDefaults.standard.removeObject(forKey: tokenDateKey)
        UserDefaults.standard.removeObject(forKey: tokenExpiryKey)
    }

    static func extractTokens(from data: Data) throws -> (String, String) {
        if let json = try JSONSerialization.jsonObject(with: data, options: []) as? [String: Any] {
            guard let accessToken = json["access_token"] as? String,
                  let refreshToken = json["refresh_token"] as? String else {
                throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Missing access or refresh token in response."])
            }
            return (accessToken, refreshToken)
        } else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to parse response for tokens."])
        }
    }

    static func extractExpiryDate(from data: Data) throws -> TimeInterval {
        if let json = try JSONSerialization.jsonObject(with: data, options: []) as? [String: Any],
           let expiresIn = json["expires_in"] as? TimeInterval {
            return expiresIn
        } else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to parse expiry date from response."])
        }
    }
}
