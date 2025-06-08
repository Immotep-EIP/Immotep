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
    @Published var rooms: [PropertyRoomsTenant] = []
    
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
        
        guard let httpResponse = response as? HTTPURLResponse, (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
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
                let rooms = try await fetchPropertyRooms(token: token)
                self.rooms = rooms
                property.rooms = rooms.map { PropertyRooms(id: $0.id, name: $0.name, checked: false, inventory: []) }
            }
        } catch {
            print("Error fetching tenant data: \(error.localizedDescription)")
        }
        
        return property
    }
    
    func fetchPropertyRooms(token: String) async throws -> [PropertyRoomsTenant] {
        let url = URL(string: "\(APIConfig.baseURL)/tenant/leases/current/property/inventory/")!
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        do {
            let (data, response) = try await URLSession.shared.data(for: urlRequest)
            
            if let jsonString = String(data: data, encoding: .utf8) {
                print("API Response: \(jsonString)")
            } else {
                print("API Response: Unable to convert data to string")
            }
            
            guard let httpResponse = response as? HTTPURLResponse else {
                throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
            }
            
            guard (200...299).contains(httpResponse.statusCode) else {
                let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
                print("API Error: Status Code \(httpResponse.statusCode), Body: \(errorBody)")
                switch httpResponse.statusCode {
                case 403:
                    throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property not yours.".localized()])
                case 404:
                    throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "No property inventory found.".localized()])
                default:
                    throw NSError(domain: "", code: httpResponse.statusCode,
                                  userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
                }
            }
            
            do {
                let decoder = JSONDecoder()
                let inventoryResponse = try decoder.decode(PropertyInventoryResponse.self, from: data)
                print("Decoded Inventory Response: ID: \(inventoryResponse.id), Rooms: \(inventoryResponse.rooms.map { "\($0.name) (ID: \($0.id))" })")
                return inventoryResponse.rooms.map { room in
                    PropertyRoomsTenant(id: room.id, name: room.name)
                }
            } catch {
                print("Decoding Error: \(error)")
                if let decodingError = error as? DecodingError {
                    switch decodingError {
                    case .dataCorrupted(let context):
                        print("Data corrupted: \(context.debugDescription)")
                    case .keyNotFound(let key, let context):
                        print("Key '\(key)' not found: \(context.debugDescription)")
                    case .typeMismatch(let type, let context):
                        print("Type mismatch for type \(type): \(context.debugDescription)")
                    case .valueNotFound(let type, let context):
                        print("Value not found for type \(type): \(context.debugDescription)")
                    @unknown default:
                        print("Unknown decoding error")
                    }
                }
                throw error
            }
        } catch {
            print("Fetch Rooms Error: \(error.localizedDescription)")
            throw error
        }
    }
    
    func fetchPropertiesPicture(propertyId: String) async throws -> UIImage? {
        let url = URL(string: "\(APIConfig.baseURL)/tenant/leases/current/property/picture/")!
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
        let url = URL(string: "\(APIConfig.baseURL)/tenant/leases/current/docs/")!
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
        let urlComponents = URLComponents(string: "\(APIConfig.baseURL)/tenant/leases/current/damages/")!
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
        let url = URL(string: "\(APIConfig.baseURL)/tenant/leases/current/")!
        
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
        
        do {
            let decoder = JSONDecoder()
            let dateFormatter = DateFormatter()
            dateFormatter.dateFormat = "yyyy-MM-dd'T'HH:mm:ss.SSSZ"
            let fallbackFormatter = DateFormatter()
            fallbackFormatter.dateFormat = "yyyy-MM-dd'T'HH:mm:ssZ"
            decoder.dateDecodingStrategy = .custom { decoder in
                let container = try decoder.singleValueContainer()
                let dateString = try container.decode(String.self)
                if let date = dateFormatter.date(from: dateString) {
                    return date
                } else if let date = fallbackFormatter.date(from: dateString) {
                    return date
                }
                throw DecodingError.dataCorruptedError(in: container, debugDescription: "Format de date invalide: \(dateString)")
            }
            
            let leaseResponse = try decoder.decode(LeaseResponse.self, from: data)
            
            if leaseResponse.active && leaseResponse.propertyId == propertyId {
                return leaseResponse.id
            } else {
                return nil
            }
        } catch {
            throw error
        }
    }
    
    func createDamage(propertyId: String, leaseId: String, damage: DamageRequest, token: String) async throws -> String {
        let url = URL(string: "\(APIConfig.baseURL)/tenant/leases/current/damages/")!
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "POST"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let jsonData = try JSONEncoder().encode(damage)
        urlRequest.httpBody = jsonData
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse, httpResponse.statusCode == 201 else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            let code = (response as? HTTPURLResponse)?.statusCode ?? 0
            switch code {
            case 400:
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Missing fields or bad base64 string: \(errorBody)".localized()])
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Invalid data.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "No active lease.".localized()])
            case 500:
                throw NSError(domain: "", code: 500, userInfo: [NSLocalizedDescriptionKey: "Internal server error: \(errorBody)".localized()])
            default:
                throw NSError(domain: "", code: code, userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(code) - \(errorBody)".localized()])
            }
        }
        
        let decoder = JSONDecoder()
        let idResponse = try decoder.decode(IdResponse.self, from: data)
        return idResponse.id
    }
}

struct PropertyRoomsTenant: Identifiable, Equatable {
    let id: String
    let name: String
    
    static func == (lhs: PropertyRoomsTenant, rhs: PropertyRoomsTenant) -> Bool {
        return lhs.id == rhs.id
    }
}

struct FurnitureResponseTenant: Codable {
    let id: String
    let name: String
    let quantity: Int
    let archived: Bool
    
    enum CodingKeys: String, CodingKey {
        case id
        case name
        case quantity
        case archived
    }
}

struct RoomResponseTenant: Codable {
    let id: String
    let name: String
    let archived: Bool
    let furnitures: [FurnitureResponseTenant]
    
    enum CodingKeys: String, CodingKey {
        case id
        case name
        case archived
        case furnitures
    }
}

struct PropertyInventoryResponse: Codable {
    let id: String
    let ownerId: String
    let name: String
    let address: String
    let city: String
    let postalCode: String
    let country: String
    let areaSqm: Double
    let rentalPricePerMonth: Int
    let depositPrice: Int
    let createdAt: String
    let archived: Bool
    let nbDamage: Int
    let status: String
    let lease: LeaseInfo?
    let rooms: [RoomResponseTenant]
    
    enum CodingKeys: String, CodingKey {
        case id
        case ownerId = "owner_id"
        case name
        case address
        case city
        case postalCode = "postal_code"
        case country
        case areaSqm = "area_sqm"
        case rentalPricePerMonth = "rental_price_per_month"
        case depositPrice = "deposit_price"
        case createdAt = "created_at"
        case archived
        case nbDamage = "nb_damage"
        case status
        case lease
        case rooms
    }
}
