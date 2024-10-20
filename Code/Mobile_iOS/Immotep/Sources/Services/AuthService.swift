//
//  AuthService.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 20/10/2024.
//

import Foundation

actor AuthService: Sendable {
    static let shared = AuthService()

    let apiUrl = "http://localhost:3001/api/v1"

    func loginUser(email: String, password: String) async throws -> (String, String) {
        let url = URL(string: "\(apiUrl)/auth/token")!
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("application/x-www-form-urlencoded", forHTTPHeaderField: "Content-Type")

        let body = "grant_type=password&username=\(email)&password=\(password)"
        request.httpBody = body.data(using: .utf8)

        let (data, response) = try await URLSession.shared.data(for: request)

        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server."])
        }

        guard httpResponse.statusCode == 200 else {
            if httpResponse.statusCode == 401 {
                throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Invalid credentials."])
            } else {
                throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode)"])
            }
        }

        let (accessToken, refreshToken) = try TokenStorage.extractTokens(from: data)
        TokenStorage.storeTokens(accessToken: accessToken, refreshToken: refreshToken)
        return (accessToken, refreshToken)
    }

    func refreshToken(refreshToken: String) async throws -> String {
        let url = URL(string: "\(apiUrl)/auth/token")!
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("application/x-www-form-urlencoded", forHTTPHeaderField: "Content-Type")

        let body = "refresh_token=\(refreshToken)"
        request.httpBody = body.data(using: .utf8)

        let (data, response) = try await URLSession.shared.data(for: request)

        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server."])
        }

        guard httpResponse.statusCode == 200 else {
            throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Failed to refresh token."])
        }

        let (accessToken, refreshToken) = try TokenStorage.extractTokens(from: data)
        TokenStorage.storeTokens(accessToken: accessToken, refreshToken: refreshToken)
        return accessToken
    }
}
