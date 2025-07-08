//
//  OwnerPropertyViewModel.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/06/2025.
//

import SwiftUI
import Foundation

class OwnerPropertyViewModel: ObservableObject {
    @Published var properties: [Property] = []
    @Published var damages: [DamageResponse] = []
    @Published var isFetchingDamages = false
    @Published var damagesError: String?
    weak var propertyViewModel: PropertyViewModel?
    
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
        
        let propertyID = try await Task.detached(priority: .background) {
            try JSONDecoder().decode(PropertyID.self, from: data)
        }.value
        
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
        guard let imageString = try await convertUIImageToBase64(propertyPicture), !imageString.isEmpty else {
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
//        let startTime = Date()

        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/")!
        let token = try await TokenStorage.getValidAccessToken()
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")

        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        guard let httpResponse = response as? HTTPURLResponse, (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            throw NSError(domain: "", code: (response as? HTTPURLResponse)?.statusCode ?? 0, userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \((response as? HTTPURLResponse)?.statusCode ?? 0) - \(errorBody)".localized()])
        }

//        let decodeTime = Date()
//        print("Network request took: \(decodeTime.timeIntervalSince(startTime)) seconds")

        let propertiesData = try await Task.detached(priority: .background) {
            let decoder = JSONDecoder()
            return try decoder.decode([PropertyResponse].self, from: data)
        }.value

//        print("Decoding took: \(Date().timeIntervalSince(decodeTime)) seconds")

        await MainActor.run {
            self.properties = propertiesData.map { propertyResponse in
                let existingProperty = self.properties.first(where: { $0.id == propertyResponse.id })
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
        }

        await withTaskGroup(of: (index: Int, image: UIImage?).self) { group in
            for index in self.properties.indices where self.properties[index].photo == nil {
                group.addTask {
                    let propertyId = self.properties[index].id
                    do {
                        let image = try await self.fetchPropertiesPicture(propertyId: propertyId)
                        return (index: index, image: image)
                    } catch {
                        return (index: index, image: nil)
                    }
                }
            }

            for await (index, image) in group {
                await MainActor.run {
                    self.properties[index].photo = image
                }
            }
        }

        await MainActor.run {
            self.objectWillChange.send()
//            print("UI update took: \(Date().timeIntervalSince(decodeTime)) seconds")
        }
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
        
        let idResponse = try await Task.detached(priority: .background) {
            try JSONDecoder().decode(IdResponse.self, from: data)
        }.value
        
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
        await MainActor.run {
            self.propertyViewModel?.isFetchingDocuments = true
        }
        defer {
            Task { @MainActor in
                self.propertyViewModel?.isFetchingDocuments = false
            }
        }

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
        
        await MainActor.run {
            if let index = self.properties.firstIndex(where: { $0.id == propertyId }) {
                var updatedProperty = self.properties[index]
                updatedProperty.documents = documents
                self.properties[index] = updatedProperty
                self.objectWillChange.send()
            }
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
                
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
                
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
        
        let leases = try await Task.detached(priority: .background) {
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
            return try decoder.decode([LeaseResponse].self, from: data)
        }.value
        
        if let activeLease = leases.first(where: { $0.active && ($0.endDate == nil || $0.endDate ?? Date() > Date()) }) {
            return activeLease.id
        }
        return nil
    }
    
    @MainActor
    func fetchPropertyDamages(propertyId: String, fixed: Bool? = nil) async throws -> [DamageResponse] {
        isFetchingDamages = true
        damagesError = nil

        var urlComponents = URLComponents(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/leases/current/damages/")!
        
        if let fixed = fixed {
            urlComponents.queryItems = [URLQueryItem(name: "fixed", value: String(fixed))]
        }
        
        guard let url = urlComponents.url else {
            damagesError = "Invalid URL".localized()
            isFetchingDamages = false
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid URL".localized()])
        }
        
        let token = try await TokenStorage.getValidAccessToken()
        var request = URLRequest(url: url)
        request.httpMethod = "GET"
        request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        request.setValue("application/json", forHTTPHeaderField: "Accept")
        
//        let startTime = Date()
        
        let (data, response) = try await URLSession.shared.data(for: request)
        
        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid server response".localized()])
        }

//        let networkTime = Date()
//        print("Network request took: \(networkTime.timeIntervalSince(startTime)) seconds")

        if httpResponse.statusCode == 404 {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            if errorBody.contains("no-active-lease") {
                if let index = properties.firstIndex(where: { $0.id == propertyId }) {
                    var updatedProperty = properties[index]
                    updatedProperty.damages = []
                    properties[index] = updatedProperty
                    objectWillChange.send()
                }
                isFetchingDamages = false
                self.damages = []
                return []
            }
            throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "Failed with status code: 404 - \(errorBody)".localized()])
        }

        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)".localized()])
        }

        let damages = try await Task.detached(priority: .background) {
            let decoder = JSONDecoder()
            decoder.dateDecodingStrategy = .iso8601
            return try decoder.decode([DamageResponse].self, from: data)
        }.value
        
//        print("Decoding took: \(Date().timeIntervalSince(networkTime)) seconds")

        if let index = properties.firstIndex(where: { $0.id == propertyId }) {
            var updatedProperty = properties[index]
            updatedProperty.damages = damages
            properties[index] = updatedProperty
        }
        
        self.damages = damages
        objectWillChange.send()
        isFetchingDamages = false
//        print("UI update took: \(Date().timeIntervalSince(networkTime)) seconds")
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
        
        let report = try await Task.detached(priority: .background) {
            let decoder = JSONDecoder()
            return try decoder.decode(InventoryReportResponse.self, from: data)
        }.value
        return report
    }

    func fetchDamageByID(propertyId: String, damageId: String, token: String) async throws -> DamageResponse {
//        let startTime = Date()

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
        
//        let decodeTime = Date()
//        print("Network request took: \(decodeTime.timeIntervalSince(startTime)) seconds")

        let damage = try await Task.detached(priority: .background) {
            let decoder = JSONDecoder()
            return try decoder.decode(DamageResponse.self, from: data)
        }.value

//        print("Decoding took: \(Date().timeIntervalSince(decodeTime)) seconds")
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

    func uploadOwnerDocument(propertyId: String, fileName: String, base64Data: String) async throws -> String {
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
        
        let idResponse = try await Task.detached(priority: .background) {
            try JSONDecoder().decode(IdResponse.self, from: data)
        }.value
        return idResponse.id
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

        await MainActor.run {
            if let index = self.properties.firstIndex(where: { $0.id == propertyId }) {
                var updatedProperty = self.properties[index]
                updatedProperty.documents.removeAll { $0.id == documentId }
                self.properties[index] = updatedProperty
                self.objectWillChange.send()
            }
        }
    }
    
    private func convertUIImageToBase64(_ image: UIImage) async throws -> String? {
        await Task.detached(priority: .background) {
            guard let imageData = image.jpegData(compressionQuality: 0.8) else {
                return nil
            }
            let base64String = "data:image/jpeg;base64,\(imageData.base64EncodedString())"
            return base64String
        }.value
    }
}
