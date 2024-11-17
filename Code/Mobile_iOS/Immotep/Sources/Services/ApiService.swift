//
//  ApiService.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 20/10/2024.
//

import Foundation

actor ApiService: Sendable, ApiServiceProtocol {
    static let shared = ApiService()

    func registerUser(with model: RegisterModel) async throws -> String {
        let url = URL(string: "\(baseURL)/auth/register")!
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")

        let body: [String: Any] = [
            "email": model.email,
            "password": model.password,
            "firstname": model.firstName,
            "lastname": model.name
        ]

        let jsonData = try JSONSerialization.data(withJSONObject: body)
        request.httpBody = jsonData

        let (_, response) = try await URLSession.shared.data(for: request)

        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server."])
        }
        guard httpResponse.statusCode == 201 else {
            if httpResponse.statusCode == 409 {
                throw NSError(domain: "", code: 409, userInfo: [NSLocalizedDescriptionKey: "Email already exists."])
            } else if httpResponse.statusCode == 400 {
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Empty fields."])
            } else {
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode)"])
            }
        }
        return "Registration successful!"
    }
}

protocol ApiServiceProtocol {
    func registerUser(with model: RegisterModel) async throws -> (String)
}
