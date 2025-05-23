//
//  AuthService.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 20/10/2024.
//

import Foundation

actor AuthService: AuthServiceProtocol {
    static let shared = AuthService()

    func loginUser(email: String, password: String, keepMeSignedIn: Bool) async throws -> (String, String, String, String) {
        return try await requestToken(grantType: "password", email: email, password: password, keepMeSignedIn: keepMeSignedIn)
    }

    func requestToken(grantType: String, email: String? = nil, password: String? = nil, refreshToken: String? = nil, keepMeSignedIn: Bool) async throws -> (String, String, String, String) {
        let url = URL(string: "\(APIConfig.baseURL)/auth/token/")!
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("application/x-www-form-urlencoded", forHTTPHeaderField: "Content-Type")
        request.setValue("application/json", forHTTPHeaderField: "Accept")

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

        let tokenResponse = try JSONDecoder().decode(TokenResponse.self, from: data)
        let accessToken = tokenResponse.access_token
        let refreshToken = tokenResponse.refresh_token
        let userId = tokenResponse.properties.id
        let userRole = tokenResponse.properties.role
        let expiresIn = TimeInterval(tokenResponse.expires_in)
        TokenStorage.storeTokens(accessToken: accessToken, refreshToken: refreshToken, expiresIn: expiresIn, keepMeSignedIn: keepMeSignedIn)

        return (accessToken, refreshToken, userId, userRole)
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
        guard let refreshToken = await TokenStorage.getRefreshToken() else {
            throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "No refresh token found. Please log in again."])
        }
        let (newAccessToken, _, _, _) = try await requestToken(grantType: "refresh_token", refreshToken: refreshToken, keepMeSignedIn: true)
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
    func loginUser(email: String, password: String, keepMeSignedIn: Bool) async throws -> (String, String, String, String)
    func requestToken(grantType: String, email: String?, password: String?, refreshToken: String?, keepMeSignedIn: Bool) async throws -> (String, String, String, String)
    func authorizedRequest(for endpoint: String) async throws -> Data
}

struct TokenResponse: Codable {
    let access_token: String
    let refresh_token: String
    let token_type: String
    let expires_in: Int
    let properties: UserProperties
}

struct UserProperties: Codable {
    let id: String
    let role: String
}
