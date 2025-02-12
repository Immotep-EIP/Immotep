//
//  User.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import Foundation

struct User: Decodable, Encodable {
    var id: String
    var email: String
    var firstname: String
    var lastname: String
    var role: String
//    var createdAt: Date
//    var updatedAt: Date

    enum CodingKeys: String, CodingKey {
        case id
        case email
        case firstname
        case lastname
        case role
//        case createdAt = "created_at"
//        case updatedAt = "updated_at"
    }
}

actor UserService: UserServiceProtocol {
    static let shared = UserService()

    private var currentUser: User?

    func getCurrentUser() async throws -> User {
        if TokenStorage.keepMeSignedIn() && TokenStorage.isTokenExpired() {
            do {
                let newAccessToken = try await refreshAccessTokenIfNeeded()
                return try await fetchUserProfile(with: newAccessToken)
            } catch {
                throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Failed to refresh access token."])
            }
        }

        guard let accessToken = await TokenStorage.getAccessToken() else {
            throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "No access token found."])
        }
        let user = try await fetchUserProfile(with: accessToken)
        currentUser = user
        return user
    }

    private func refreshAccessTokenIfNeeded() async throws -> String {
        guard let refreshToken = await TokenStorage.getRefreshToken() else {
            throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "No refresh token found. Please log in again."])
        }
        do {
            let (newAccessToken, _) =
            try await AuthService.shared.requestToken(grantType: "refresh_token", refreshToken: refreshToken, keepMeSignedIn: true)
            TokenStorage.storeAccessToken(newAccessToken)
            return newAccessToken
        } catch {
            print("Error while refreshing token: \(error)")
            throw error
        }

    }

    func fetchUserProfile(with token: String) async throws -> User {
        let url = URL(string: "\(baseURL)/profile/")!
        var request = URLRequest(url: url)
        request.httpMethod = "GET"
        request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")

        let (data, response) = try await URLSession.shared.data(for: request)

        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server."])
        }

        guard httpResponse.statusCode == 200 else {
            throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Failed to fetch user profile."])
        }

        do {
            let decoder = JSONDecoder()
            decoder.dateDecodingStrategy = .iso8601
            let userProfile = try decoder.decode(User.self, from: data)
            return userProfile
        } catch {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to decode user profile."])
        }
    }

    func logout() {
        currentUser = nil
    }
}

protocol UserServiceProtocol {
    func getCurrentUser() async throws -> User
    func fetchUserProfile(with token: String) async throws -> User
    func logout() async
}
