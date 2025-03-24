//
//  AuthService.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 20/10/2024.
//

import Foundation

actor AuthService: AuthServiceProtocol {
    static let shared = AuthService()

    func loginUser(email: String, password: String, keepMeSignedIn: Bool) async throws -> (String, String) {
        return try await requestToken(grantType: "password", email: email, password: password, keepMeSignedIn: keepMeSignedIn)
    }

    func requestToken(grantType: String, email: String? = nil, password: String? = nil, refreshToken: String? = nil, keepMeSignedIn: Bool) async throws -> (String, String) {
        let url = URL(string: "\(APIConfig.baseURL)/auth/token")!
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("application/x-www-form-urlencoded", forHTTPHeaderField: "Content-Type")

        var body: String

        if grantType == "password" {
            guard let email = email, let password = password else {
                throw NSError(domain: "", code: 400,
                              userInfo: [NSLocalizedDescriptionKey: "Email and password are required for password grant type."])
            }
            body = "grant_type=password&username=\(email)&password=\(password)"
        } else if grantType == "refresh_token" {
            guard let refreshToken = refreshToken else {
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Refresh token is required for refresh token grant type."])
            }

            let formattedRefreshToken = encodeSpecialCharacters(refreshToken)
            body = "grant_type=refresh_token&refresh_token=\(formattedRefreshToken)"
        } else {
            throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid grant type."])
        }

        request.httpBody = body.data(using: .utf8)

        let (data, response) = try await URLSession.shared.data(for: request)

        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server."])
        }

        guard (200...299).contains(httpResponse.statusCode) else {
            let errorMessage = String(data: data, encoding: .utf8) ?? "Unknown error"
            throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Request failed: \(errorMessage)"])
        }

        let (accessToken, refreshToken) = try TokenStorage.extractTokens(from: data)
        let expiresIn = try TokenStorage.extractExpiryDate(from: data)
        TokenStorage.storeTokens(accessToken: accessToken, refreshToken: refreshToken, expiresIn: expiresIn, keepMeSignedIn: keepMeSignedIn)

        return (accessToken, refreshToken)
    }

    func authorizedRequest(for endpoint: String) async throws -> Data {
        var accessToken = await TokenStorage.getAccessToken()

        if accessToken == nil || TokenStorage.isTokenExpired() {
            accessToken = try await refreshAccessTokenIfNeeded()
        }

        let url = URL(string: "\(APIConfig.baseURL)/\(endpoint)")!
        var request = URLRequest(url: url)
        request.setValue("Bearer \(accessToken!)", forHTTPHeaderField: "Authorization")

        let (data, response) = try await URLSession.shared.data(for: request)

        guard let httpResponse = response as? HTTPURLResponse, httpResponse.statusCode == 200 else {
            throw NSError(domain: "", code: (response as? HTTPURLResponse)?.statusCode ?? 0, userInfo: [NSLocalizedDescriptionKey: "Request failed"])
        }

        return data
    }

    private func refreshAccessTokenIfNeeded() async throws -> String {
//        print("refresh access token")
        guard let refreshToken = await TokenStorage.getRefreshToken() else {
            throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "No refresh token found. Please log in again."])
        }
//        print("refresh pending...")
        let (newAccessToken, _) = try await requestToken(grantType: "refresh_token", refreshToken: refreshToken, keepMeSignedIn: true)
//        print("refresh complete.")
        return newAccessToken
    }

    func encodeSpecialCharacters(_ token: String) -> String {
        var encodedToken = token
        encodedToken = encodedToken.replacingOccurrences(of: "/", with: "%2F")
        encodedToken = encodedToken.replacingOccurrences(of: "+", with: "%2B")
        encodedToken = encodedToken.replacingOccurrences(of: "=", with: "%3D")
        return encodedToken
    }

}

protocol AuthServiceProtocol {
    func loginUser(email: String, password: String, keepMeSignedIn: Bool) async throws -> (String, String)
    func requestToken(grantType: String, email: String?, password: String?, refreshToken: String?, keepMeSignedIn: Bool) async throws -> (String, String)
    func authorizedRequest(for endpoint: String) async throws -> Data
}
