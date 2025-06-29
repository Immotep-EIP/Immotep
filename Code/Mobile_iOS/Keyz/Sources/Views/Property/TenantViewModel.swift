//
//  TenantViewModel.swift
//  Immotep
//
//  Created by Alessio Liebenguth on 08/04/2025.
//

import Foundation

class TenantViewModel: ObservableObject {
    private let authService = AuthService.shared
    
    struct InviteRequest: Codable {
        let tenantEmail: String
        let startDate: String
        let endDate: String?
        
        enum CodingKeys: String, CodingKey {
            case tenantEmail = "tenant_email"
            case startDate = "start_date"
            case endDate = "end_date"
        }
    }
    
    struct InviteIDResponse: Codable {
        let id: String
    }
    
    func inviteTenant(propertyId: String, email: String, startDate: Date, endDate: Date? = nil) async throws {
        let endpoint = "owner/properties/\(propertyId)/send-invite/"
        let url = URL(string: "\(APIConfig.baseURL)/\(endpoint)")!
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        request.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let token = try await TokenStorage.getValidAccessToken()
        request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        
        let dateFormatter = ISO8601DateFormatter()
        let startDateString = dateFormatter.string(from: startDate)
        let endDateString = endDate != nil ? dateFormatter.string(from: endDate!) : nil
        
        let inviteRequest = InviteRequest(
            tenantEmail: email,
            startDate: startDateString,
            endDate: endDateString
        )
        let jsonData = try JSONEncoder().encode(inviteRequest)
        request.httpBody = jsonData
        
        let (data, response) = try await URLSession.shared.data(for: request)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server"])
        }
        
        switch httpResponse.statusCode {
        case 201:
            let inviteResponse = try JSONDecoder().decode(InviteIDResponse.self, from: data)
        case 400:
            throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Missing fields"])
        case 403:
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property is not yours"])
        case 404:
            throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "Property not found"])
        case 409:
            throw NSError(domain: "", code: 409, userInfo: [NSLocalizedDescriptionKey: "Invite already exists for this email"])
        default:
            throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Unexpected error: \(httpResponse.statusCode)"])
        }
    }
}
