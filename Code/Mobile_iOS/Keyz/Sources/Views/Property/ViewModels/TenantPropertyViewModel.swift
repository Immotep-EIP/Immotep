//
//  TenantPropertyViewModel.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/06/2025.
//

import SwiftUI
import Foundation

@MainActor
class TenantPropertyViewModel: ObservableObject {
    @Published var damages: [DamageResponse] = []
    @Published var isFetchingDamages = false
    @Published var damagesError: String?
    
    func fetchTenantProperty() async throws -> Property {
        let url = URL(string: "\(APIConfig.baseURL)/tenant/leases/current/property/")!
        let token = try await TokenStorage.getValidAccessToken()
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 401:
                throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized. Please check your token.".localized()])
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "Property not found.".localized()])
            case 500:
                throw NSError(domain: "", code: 500, userInfo: [NSLocalizedDescriptionKey: "Internal server error: \(errorBody)".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
        
        var property: Property
        do {
            let decoder = JSONDecoder()
            let propertyResponse = try decoder.decode(PropertyResponse.self, from: data)
            
            property = Property(
                id: propertyResponse.id,
                ownerID: propertyResponse.ownerId,
                name: propertyResponse.name,
                address: propertyResponse.address,
                city: propertyResponse.city,
                postalCode: propertyResponse.postalCode,
                country: propertyResponse.country,
                photo: nil,
                monthlyRent: propertyResponse.rentalPricePerMonth,
                deposit: propertyResponse.depositPrice,
                surface: propertyResponse.areaSqm,
                isAvailable: propertyResponse.isAvailable,
                tenantName: propertyResponse.lease?.tenantName,
                leaseId: propertyResponse.lease?.id,
                leaseStartDate: propertyResponse.lease?.startDate,
                leaseEndDate: propertyResponse.lease?.endDate,
                documents: [],
                createdAt: propertyResponse.createdAt,
                rooms: [],
                damages: []
            )
        } catch {
            throw error
        }
        
        do {
            property.photo = try await fetchPropertiesPicture(propertyId: property.id)
        } catch {
            print("Error fetching property picture: \(error.localizedDescription)")
        }
        
        do {
            if let leaseId = try await fetchActiveLeaseIdForProperty(propertyId: property.id, token: token) {
                try await fetchTenantPropertyDocuments(leaseId: leaseId, propertyId: property.id)
                try await fetchTenantDamages(leaseId: leaseId)
                property.damages = damages
            }
        } catch {
            print("Error fetching tenant data: \(error.localizedDescription)")
        }
        
        return property
    }
    
    func fetchPropertiesPicture(propertyId: String) async throws -> UIImage? {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/picture/")!
        var attemptCount = 0
        let maxAttempts = 2
        
        while attemptCount < maxAttempts {
            attemptCount += 1
            do {
                let token = try await TokenStorage.getValidAccessToken()
                
                var urlRequest = URLRequest(url: url)
                urlRequest.httpMethod = "GET"
                urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
                urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
                
                let (data, response) = try await URLSession.shared.data(for: urlRequest)
                
                guard let httpResponse = response as? HTTPURLResponse else {
                    throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
                }
                
                switch httpResponse.statusCode {
                case 200, 201:
                    do {
                        let propertyImage = try JSONDecoder().decode(PropertyImageBase64.self, from: data)
                        
                        var base64String = propertyImage.data
                        if base64String.contains(",") {
                            base64String = base64String.components(separatedBy: ",").last ?? base64String
                        }
                        
                        guard let imageData = Data(base64Encoded: base64String, options: [.ignoreUnknownCharacters]) else {
                            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to decode base64 data.".localized()])
                        }
                        guard let image = UIImage(data: imageData) else {
                            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to create image from data.".localized()])
                        }
                        return image
                    } catch {
                        throw error
                    }
                case 204:
                    return nil
                case 401:
                    if attemptCount == maxAttempts {
                        throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized after \(maxAttempts) attempts.".localized()])
                    }
                    continue
                case 403:
                    return nil
                case 404:
                    return nil
                default:
                    let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
                    throw NSError(domain: "", code: httpResponse.statusCode,
                                  userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)"])
                }
            } catch {
                if attemptCount == maxAttempts {
                    throw error
                }
            }
        }
        return nil
    }
    
    func fetchTenantPropertyDocuments(leaseId: String, propertyId: String) async throws -> [PropertyDocument] {
        let url = URL(string: "\(APIConfig.baseURL)/tenant/leases/\(leaseId)/documents/")!
        let token = try await TokenStorage.getValidAccessToken()
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let config = URLSessionConfiguration.default
        config.waitsForConnectivity = true
        config.timeoutIntervalForRequest = 120
        config.timeoutIntervalForResource = 300
        config.httpMaximumConnectionsPerHost = 10
        let session = URLSession(configuration: config)
        
        do {
            let (data, response) = try await session.data(for: urlRequest)
            
            guard let httpResponse = response as? HTTPURLResponse else {
                throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server".localized()])
            }
            
            guard (200...299).contains(httpResponse.statusCode) else {
                let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)"])
            }
            
            let decoder = JSONDecoder()
            let documentsData = try decoder.decode([PropertyDocumentResponse].self, from: data)
            
            return documentsData.map { doc in
                let cleanBase64 = doc.data.components(separatedBy: ",").last ?? doc.data
                return PropertyDocument(id: doc.id, title: doc.name, fileName: doc.name, data: cleanBase64)
            }
        } catch {
            throw error
        }
    }
    
    func fetchTenantDamages(leaseId: String, fixed: Bool? = nil) async throws {
        let urlComponents = URLComponents(string: "\(APIConfig.baseURL)/tenant/leases/\(leaseId)/damages/")!
        var urlRequest = URLRequest(url: urlComponents.url!)
        
        if let fixed = fixed {
            var components = URLComponents(url: urlComponents.url!, resolvingAgainstBaseURL: false)!
            components.queryItems = [URLQueryItem(name: "fixed", value: String(fixed))]
            urlRequest.url = components.url
        }
        
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(try await TokenStorage.getValidAccessToken())", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        isFetchingDamages = true
        damagesError = nil
        defer { isFetchingDamages = false }
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Lease not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "No active lease.".localized()])
            case 500:
                throw NSError(domain: "", code: 500, userInfo: [NSLocalizedDescriptionKey: "Internal server error: \(errorBody)".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
        
        let decoder = JSONDecoder()
        let damagesData = try decoder.decode([DamageResponse].self, from: data)
        
        self.damages = damagesData
        objectWillChange.send()
    }
    
    func fetchActiveLeaseIdForProperty(propertyId: String, token: String) async throws -> String? {
        let url = URL(string: "\(APIConfig.baseURL)/tenant/properties/\(propertyId)/leases/current/")!
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "No active lease.".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
        
        let decoder = JSONDecoder()
        let leaseResponse = try decoder.decode(LeaseResponse.self, from: data)
        return leaseResponse.id
    }
    
    func createDamage(propertyId: String, leaseId: String, damage: DamageRequest, token: String) async throws -> String {
        let url = URL(string: "\(APIConfig.baseURL)/tenant/properties/\(propertyId)/leases/\(leaseId)/damages/")!
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "POST"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let jsonData = try JSONEncoder().encode(damage)
        urlRequest.httpBody = jsonData
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard httpResponse.statusCode == 201 else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 400:
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Missing fields or bad base64 string: \(errorBody)".localized()])
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "No active lease.".localized()])
            case 500:
                throw NSError(domain: "", code: 500, userInfo: [NSLocalizedDescriptionKey: "Internal server error: \(errorBody)".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
        
        let decoder = JSONDecoder()
        let idResponse = try decoder.decode(IdResponse.self, from: data)
        return idResponse.id
    }
}
