//
//  OwnerPropertyViewModel.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/06/2025.
//

import SwiftUI
import Foundation

@MainActor
class OwnerPropertyViewModel: ObservableObject {
    @Published var properties: [Property] = []
    @Published var damages: [DamageResponse] = []
    @Published var isFetchingDamages = false
    @Published var damagesError: String?
    
    func createProperty(request: Property, token: String) async throws -> String {
        let body: [String: Any] = [
            "address": request.address,
            "area_sqm": request.surface,
            "city": request.city,
            "country": request.country,
            "deposit_price": request.deposit,
            "name": request.name,
            "postal_code": request.postalCode,
            "rental_price_per_month": request.monthlyRent
        ]
        
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/")!
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "POST"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let jsonData = try JSONSerialization.data(withJSONObject: body)
        urlRequest.httpBody = jsonData
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 400:
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid property data: \(errorBody)".localized()])
            case 401:
                throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized. Please check your token.".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
        
        let propertyID = try JSONDecoder().decode(PropertyID.self, from: data)
        
        if let photo = request.photo {
            do {
                _ = try await updatePropertyPicture(token: token, propertyPicture: photo, propertyID: propertyID.id)
            } catch {
                print("Error updating property picture: \(error.localizedDescription)")
            }
        }
        
        try await fetchProperties()
        return propertyID.id
    }
    
    func updatePropertyPicture(token: String, propertyPicture: UIImage, propertyID: String) async throws -> String {
        let imageString = convertUIImageToBase64(propertyPicture)
        guard !imageString.isEmpty else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to convert image to base64.".localized()])
        }
        
        let body: [String: Any] = ["data": imageString]
        
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyID)/picture/")!
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "PUT"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let jsonData = try JSONSerialization.data(withJSONObject: body)
        urlRequest.httpBody = jsonData
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 400:
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid image data: \(errorBody)".localized()])
            case 401:
                throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized. Please check your token.".localized()])
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "Property not found.".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
        
        return "Successfully updated property picture."
    }
    
    func fetchProperties() async throws {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/")!
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

        let decoder = JSONDecoder()
        let propertiesData = try decoder.decode([PropertyResponse].self, from: data)

        self.properties = propertiesData.map { propertyResponse in
            let existingProperty = properties.first(where: { $0.id == propertyResponse.id })
            let status = propertyResponse.isAvailable == "invite sent" ? "pending" : propertyResponse.isAvailable
            return Property(
                id: propertyResponse.id,
                ownerID: propertyResponse.ownerId,
                name: propertyResponse.name,
                address: propertyResponse.address,
                city: propertyResponse.city,
                postalCode: propertyResponse.postalCode,
                country: propertyResponse.country,
                photo: ImageCache.shared.getImage(forKey: propertyResponse.id),
                monthlyRent: propertyResponse.rentalPricePerMonth,
                deposit: propertyResponse.depositPrice,
                surface: propertyResponse.areaSqm,
                isAvailable: status,
                tenantName: propertyResponse.lease?.tenantName,
                leaseId: propertyResponse.lease?.id,
                leaseStartDate: propertyResponse.lease?.startDate,
                leaseEndDate: propertyResponse.lease?.endDate,
                documents: existingProperty?.documents ?? [],
                createdAt: propertyResponse.createdAt,
                rooms: [],
                damages: []
            )
        }

        await withTaskGroup(of: (index: Int, image: UIImage?).self) { group in
            for index in properties.indices {
                if properties[index].photo == nil {
                    group.addTask {
                        let propertyId = await self.properties[index].id
                        do {
                            let image = try await self.fetchPropertiesPicture(propertyId: propertyId)
                            return (index: index, image: image)
                        } catch {
                            print("Error fetching picture for property ID: \(propertyId), error: \(error.localizedDescription)")
                            return (index: index, image: nil)
                        }
                    }
                }
            }

            for await (index, image) in group {
                self.properties[index].photo = image
            }
        }
        objectWillChange.send()
    }
    
    func fetchPropertiesPicture(propertyId: String) async throws -> UIImage? {
        if let cachedImage = ImageCache.shared.getImage(forKey: propertyId) {
            return cachedImage
        }

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
                    return try await Task.detached(priority: .background) {
                        let image = try JSONDecoder().decode(PropertyImageBase64.self, from: data)
                        var base64String = image.data
                        if base64String.contains(",") {
                            base64String = image.data.components(separatedBy: ",").last ?? base64String
                        }
                        guard let data = Data(base64Encoded: base64String, options: [.ignoreUnknownCharacters]),
                              let image = UIImage(data: data) else {
                            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to decode image.".localized()])
                        }
                        ImageCache.shared.setImage(image, forKey: propertyId)
                        return image
                    }.value
                case 204, 403, 404:
                    ImageCache.shared.setImage(nil, forKey: propertyId)
                    return nil
                case 401:
                    if attemptCount == maxAttempts {
                        throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized after \(maxAttempts) attempts.".localized()])
                    }
                    continue
                default:
                    let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
                    throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
                }
            } catch {
                if attemptCount == maxAttempts {
                    throw error
                }
            }
        }
        return nil
    }
    
    func updateProperty(request: Property, token: String) async throws -> String {
        let body: [String: Any] = [
            "name": request.name,
            "address": request.address,
            "city": request.city,
            "postal_code": request.postalCode,
            "country": request.country,
            "area_sqm": request.surface,
            "rental_price_per_month": request.monthlyRent,
            "deposit_price": request.deposit
        ]
        
        guard let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(request.id)/") else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid URL".localized()])
        }
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "PUT"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let jsonData = try JSONSerialization.data(withJSONObject: body)
        urlRequest.httpBody = jsonData
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 400:
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid property data: \(errorBody)".localized()])
            case 401:
                throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized. Please check your token.".localized()])
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "Property not found.".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
        
        let decoder = JSONDecoder()
        let idResponse = try decoder.decode(IdResponse.self, from: data)
        
        return idResponse.id
    }
    
    func deleteProperty(propertyId: String) async throws {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/archive/")!
        
        let token = try await TokenStorage.getValidAccessToken()
        
        let body: [String: Any] = [
            "archive": true
        ]
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "PUT"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        do {
            let jsonData = try JSONSerialization.data(withJSONObject: body)
            urlRequest.httpBody = jsonData
        } catch {
            throw NSError(domain: "", code: 0,
                          userInfo: [NSLocalizedDescriptionKey: "Failed to serialize request body: \(error.localizedDescription)".localized()])
        }
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 400:
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Missing or invalid fields: \(errorBody)".localized()])
            case 401:
                throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized. Please check your token.".localized()])
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "Property not found.".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
    }
    
    func fetchPropertyDocuments(propertyId: String) async throws -> [PropertyDocument] {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/leases/current/docs/")!
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
        
        let (data, response) = try await session.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server".localized()])
        }
        
        let responseBody = String(data: data, encoding: .utf8) ?? "Unable to decode response"
        
        guard (200...299).contains(httpResponse.statusCode) else {
            if httpResponse.statusCode == 404 {
                return []
            }
            throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(responseBody)".localized()])
        }
        
        let documents = try await Task.detached(priority: .background) {
            let decoder = JSONDecoder()
            let documentsData = try decoder.decode([PropertyDocumentResponse].self, from: data)
            
            let dateFormatter = DateFormatter()
            dateFormatter.dateFormat = "yyyy-MM-dd"
            let outputFormatter = DateFormatter()
            outputFormatter.dateFormat = "dd-MM-yyyy"
            
            return documentsData.map { doc in
                let cleanBase64 = doc.data.contains(",") ? doc.data.components(separatedBy: ",").last ?? doc.data : doc.data
                let filename = doc.name
                
                let components = filename.split(separator: "_")
                var title = filename
                
                if let dateString = components.last?.prefix(10),
                   dateFormatter.date(from: String(dateString)) != nil {
                    title = outputFormatter.string(from: dateFormatter.date(from: String(dateString))!)
                }
                
                return PropertyDocument(id: doc.id, title: title, fileName: filename, data: cleanBase64)
            }
        }.value
        
        if let index = properties.firstIndex(where: { $0.id == propertyId }) {
            var updatedProperty = properties[index]
            updatedProperty.documents = documents
            properties[index] = updatedProperty
//             objectWillChange.send()
        }
        
        return documents
    }
    
    func cancelInvite(propertyId: String, token: String) async throws {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/cancel-invite/")!
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "DELETE"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard httpResponse.statusCode == 204 else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property is not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "No pending lease.".localized()])
            case 500:
                throw NSError(domain: "", code: 500, userInfo: [NSLocalizedDescriptionKey: "Internal server error: \(errorBody)".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
    }
    
    func endLease(propertyId: String, leaseId: String, token: String) async throws {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/leases/current/end/")!
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "PUT"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard httpResponse.statusCode == 204 else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 400:
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Cannot end non-current lease.".localized()])
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property is not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "No active lease.".localized()])
            case 500:
                throw NSError(domain: "", code: 500, userInfo: [NSLocalizedDescriptionKey: "Internal server error: \(errorBody)".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
    }
    
    func fetchActiveLease(propertyId: String, token: String) async throws -> String? {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/leases/")!
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        debugPrintAPIRequest(urlRequest)
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        debugPrintAPIResponse(data, response: response, error: nil)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property is not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "No active lease.".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
        
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
            throw DecodingError.dataCorruptedError(in: container, debugDescription: "Invalid date format: \(dateString)")
        }
        
        let leases = try decoder.decode([LeaseResponse].self, from: data)
        if let activeLease = leases.first(where: { $0.active && ($0.endDate == nil || $0.endDate ?? Date() > Date()) }) {
            return activeLease.id
        }
        return nil
    }
    
    func fetchPropertyDamages(propertyId: String, fixed: Bool? = nil) async throws -> [DamageResponse] {
        let token = try await TokenStorage.getValidAccessToken()
        var urlComponents = URLComponents(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/leases/current/damages/")!
        
        if let fixed = fixed {
            urlComponents.queryItems = [URLQueryItem(name: "fixed", value: String(fixed))]
        }
        
        guard let url = urlComponents.url else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid URL".localized()])
        }
        
        var request = URLRequest(url: url)
        request.httpMethod = "GET"
        request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        request.setValue("application/json", forHTTPHeaderField: "Accept")

        let (data, response) = try await URLSession.shared.data(for: request)
        guard let httpResponse = response as? HTTPURLResponse, (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            throw NSError(domain: "", code: (response as? HTTPURLResponse)?.statusCode ?? 0, userInfo: [NSLocalizedDescriptionKey: errorBody])
        }

        let decoder = JSONDecoder()
        let damages = try decoder.decode([DamageResponse].self, from: data)
        self.damages = damages
        objectWillChange.send()
        return damages
    }
    
    func fetchLastInventoryReport(propertyId: String, leaseId: String) async throws -> InventoryReportResponse? {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/leases/\(leaseId)/inventory-reports/latest/")!
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
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property not yours.".localized()])
            case 404:
                return nil
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
        
        do {
            let decoder = JSONDecoder()
            let report = try decoder.decode(InventoryReportResponse.self, from: data)
            return report
        } catch {
            throw error
        }
    }

    func fetchDamageByID(propertyId: String, damageId: String, token: String) async throws -> DamageResponse {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/leases/current/damages/\(damageId)/")!
        
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
            case 400:
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid request: \(errorBody)".localized()])
            case 401:
                throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized. Please check your token.".localized()])
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property or lease not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "Damage not found.".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
        
        let decoder = JSONDecoder()
        let damage = try decoder.decode(DamageResponse.self, from: data)
        return damage
    }

    func updateDamageStatus(propertyId: String, damageId: String, fixPlannedAt: String?, read: Bool, token: String) async throws {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/leases/current/damages/\(damageId)/")!
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "PUT"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        var body: [String: Any] = ["read": read]
        if let fixPlannedAt = fixPlannedAt {
            body["fix_planned_at"] = fixPlannedAt
        }
        
        let jsonData = try JSONSerialization.data(withJSONObject: body)
        urlRequest.httpBody = jsonData
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 400:
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid data: \(errorBody)".localized()])
            case 401:
                throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized. Please check your token.".localized()])
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property or lease not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "Damage not found.".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
    }

    func fixDamage(propertyId: String, damageId: String, token: String) async throws {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/leases/current/damages/\(damageId)/fix/")!
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "PUT"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 400:
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid request: \(errorBody)".localized()])
            case 401:
                throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized. Please check your token.".localized()])
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property or lease not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "Damage not found.".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
    }

    func uploadOwnerDocument(propertyId: String, fileName: String, base64Data: String) async throws {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/leases/current/docs/")!
        let token = try await TokenStorage.getValidAccessToken()
        
        let body: [String: Any] = [
            "name": fileName,
            "data": base64Data
        ]
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "POST"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let jsonData = try JSONSerialization.data(withJSONObject: body)
        urlRequest.httpBody = jsonData
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            switch httpResponse.statusCode {
            case 400:
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid document data: \(errorBody)".localized()])
            case 401:
                throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized. Please check your token.".localized()])
            case 403:
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property not yours.".localized()])
            case 404:
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "Property or lease not found.".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
            }
        }
        
        let decoder = JSONDecoder()
        let _ = try decoder.decode(IdResponse.self, from: data)
    }
    
    func deleteDocument(propertyId: String, documentId: String) async throws {
            let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/leases/current/docs/\(documentId)/")!
            let token = try await TokenStorage.getValidAccessToken()

            var urlRequest = URLRequest(url: url)
            urlRequest.httpMethod = "DELETE"
            urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
            urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")

            let (data, response) = try await URLSession.shared.data(for: urlRequest)

            guard let httpResponse = response as? HTTPURLResponse else {
                throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
            }

            guard httpResponse.statusCode == 204 else {
                let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
                switch httpResponse.statusCode {
                case 400:
                    throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid request: \(errorBody)".localized()])
                case 401:
                    throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized. Please check your token.".localized()])
                case 403:
                    throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property or document not yours.".localized()])
                case 404:
                    throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "Document or lease not found.".localized()])
                default:
                    throw NSError(domain: "", code: httpResponse.statusCode,
                                  userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
                }
            }
        }
}
