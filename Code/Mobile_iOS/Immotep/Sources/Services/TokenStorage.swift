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
    private static let keepMeSignedInKey = "keep_me_signed_in"

    static func storeTokens(accessToken: String, refreshToken: String, expiresIn: TimeInterval? = nil, keepMeSignedIn: Bool) {
        let currentDate = Date()
        if keepMeSignedIn {
            UserDefaults.standard.set(accessToken, forKey: accessTokenKey)
            UserDefaults.standard.set(refreshToken, forKey: refreshTokenKey)
            UserDefaults.standard.set(currentDate, forKey: tokenDateKey)
            UserDefaults.standard.set(keepMeSignedIn, forKey: keepMeSignedInKey)
            print("Keep me signed in: \(UserDefaults.standard.bool(forKey: keepMeSignedInKey))")
            if let expiresIn = expiresIn {
                UserDefaults.standard.set(expiresIn, forKey: tokenExpiryKey)
            }
        } else {
            Task {
                SessionStorage.setAccessToken(accessToken)
                SessionStorage.setRefreshToken(refreshToken)
            }
        }
    }

    static func storeAccessToken(_ accessToken: String) {
        UserDefaults.standard.set(accessToken, forKey: accessTokenKey)
        UserDefaults.standard.set(Date(), forKey: tokenDateKey)
    }

    static func keepMeSignedIn() -> Bool {
        return UserDefaults.standard.bool(forKey: keepMeSignedInKey)
    }

    static func getAccessTokenFromLocalStorage() -> String? {
        return UserDefaults.standard.string(forKey: accessTokenKey)
    }

    static func getAccessToken() async -> String? {
        if let token = UserDefaults.standard.string(forKey: accessTokenKey) {
            return token
        } else {
            return SessionStorage.getAccessToken()
        }
    }

    static func getRefreshToken() async -> String? {
        if let token = UserDefaults.standard.string(forKey: refreshTokenKey) {
            return token
        } else {
            return SessionStorage.getRefreshToken()
        }
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
        UserDefaults.standard.set(false, forKey: keepMeSignedInKey)
        print("Access token after clean: \(UserDefaults.standard.string(forKey: accessTokenKey) ?? "none")")
        print("Keep me signed in after clean: \(UserDefaults.standard.bool(forKey: keepMeSignedInKey))")

        Task {
            SessionStorage.setAccessToken(nil)
            SessionStorage.setRefreshToken(nil)
        }
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