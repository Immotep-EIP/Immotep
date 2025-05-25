//
//  PropertyViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//
import SwiftUI
import Foundation

@MainActor
class PropertyViewModel: ObservableObject {
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
            print("Create Property failed with status \(httpResponse.statusCode): \(errorBody)")
            switch httpResponse.statusCode {
            case 400:
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid property data: \(errorBody)".localized()])
            case 401:
                throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized. Please check your token.".localized()])
            default:
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)"])
            }
        }
        
        let propertyID = try JSONDecoder().decode(PropertyID.self, from: data)
        print("Property created with ID: \(propertyID.id)")
        
        if let photo = request.photo {
            do {
                print("Uploading photo for property \(propertyID.id)")
                let result = try await updatePropertyPicture(token: token, propertyPicture: photo, propertyID: propertyID.id)
                //                print("Photo upload result: \(result)")
            } catch {
                print("Error uploading property picture: \(error.localizedDescription)")
            }
        }
        
        await fetchProperties()
        return "Property successfully created!"
    }
    
    func updatePropertyPicture(token: String, propertyPicture: UIImage, propertyID: String) async throws -> String {
        guard let imageString = convertUIImageToBase64(propertyPicture) else {
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
            print("Update Picture failed with status \(httpResponse.statusCode): \(errorBody)")
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
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)"])
            }
        }
        
        return "Successfully updated property picture."
    }
    
    func fetchProperties() async {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/")!
        
        do {
            let token = try await TokenStorage.getValidAccessToken()
            
            print("token: \(token)")
            
            var urlRequest = URLRequest(url: url)
            urlRequest.httpMethod = "GET"
            urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
            urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
            
            let (data, response) = try await URLSession.shared.data(for: urlRequest)
            
            guard let httpResponse = response as? HTTPURLResponse else {
                print("Invalid response from server")
                return
            }
            
            guard (200...299).contains(httpResponse.statusCode) else {
                let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
                print("Fetch Properties failed with status \(httpResponse.statusCode): \(errorBody)")
                return
            }
            
            let decoder = JSONDecoder()
            let propertiesData = try decoder.decode([PropertyResponse].self, from: data)
            
            self.properties = propertiesData.map { propertyResponse in
                Property(
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
                    tenantName: propertyResponse.tenant,
                    leaseStartDate: propertyResponse.startDate,
                    leaseEndDate: propertyResponse.endDate,
                    documents: [],
                    createdAt: propertyResponse.createdAt,
                    rooms: [],
                    damages: []
                )
            }
            
            await withTaskGroup(of: (index: Int, image: UIImage?).self) { group in
                for index in properties.indices {
                    group.addTask {
                        let propertyId = await self.properties[index].id
                        do {
                            let image = try await self.fetchPropertiesPicture(propertyId: propertyId)
                            print("Image fetch result for property \(propertyId): \(image != nil ? "Success" : "Nil")")
                            return (index: index, image: image)
                        } catch {
                            print("Failed to fetch picture for Property ID: \(propertyId), error: \(error.localizedDescription)")
                            return (index: index, image: nil)
                        }
                    }
                }
                
                for await (index, image) in group {
                    self.properties[index].photo = image
                    print("Image assigned to property \(self.properties[index].id): \(image != nil ? "Loaded" : "Nil")")
                }
            }
            //            print("Properties updated with images: \(properties.map { ($0.id, $0.photo != nil) })")
            objectWillChange.send()
        } catch {
            print("Error fetching properties: \(error.localizedDescription)")
        }
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
                    //                    if let jsonString = String(data: data, encoding: .utf8) {
                    //                        print("Raw JSON response for property \(propertyId): \(jsonString)")
                    //                    } else {
                    //                        print("Failed to decode raw JSON response for property \(propertyId)")
                    //                    }
                    
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
                    print("Property \(propertyId) does not belong to the user")
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
    
    func updateProperty(request: Property, token: String) async throws -> PropertyResponse {
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
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid URL"])
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
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)"])
            }
        }
        
        let decoder = JSONDecoder()
        let updatedProperty = try decoder.decode(PropertyResponse.self, from: data)
        await fetchProperties()
        return updatedProperty
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
    
    func fetchPropertyDocuments(propertyId: String) async throws {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/documents/")!
        
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
                print("Fetch Documents failed with status \(httpResponse.statusCode): \(errorBody)")
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)"])
            }
            
            let decoder = JSONDecoder()
            let documentsData = try decoder.decode([PropertyDocumentResponse].self, from: data)
            print("Documents fetched for property \(propertyId): \(documentsData.count)")
            
            if let index = properties.firstIndex(where: { $0.id == propertyId }) {
                var updatedProperty = properties[index]
                updatedProperty.documents = documentsData.map { doc in
                    PropertyDocument(id: doc.id, title: doc.name, fileName: doc.name, data: doc.data)
                }
                properties[index] = updatedProperty
                objectWillChange.send()
            }
        } catch {
            print("Fetch documents error for property \(propertyId): \(error.localizedDescription)")
            throw error
        }
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
            print("Cancel Invite failed with status \(httpResponse.statusCode): \(errorBody)")
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
        print("Invite cancelled for property \(propertyId)")
    }
    
    func endLease(propertyId: String, leaseId: String, token: String) async throws {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/leases/\(leaseId)/end/")!
        
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
            print("End Lease failed with status \(httpResponse.statusCode): \(errorBody)")
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
        print("Lease ended for property \(propertyId), lease \(leaseId)")
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
            print("Fetch Active Lease failed with status \(httpResponse.statusCode): \(errorBody)")
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
        let leases = try decoder.decode([LeaseResponse].self, from: data)
        if let activeLease = leases.first(where: { $0.endDate == nil }) {
            print("Active lease found for property \(propertyId): \(activeLease.id)")
            return activeLease.id
        }
        print("No active lease found for property \(propertyId)")
        print("leaseIDs: \(leases.map(\.id))")
        return nil
    }
    
    
    func fetchPropertyDamages(propertyId: String) async throws {
        let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(propertyId)/damages/")!
        
        let token = try await TokenStorage.getValidAccessToken()
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
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
            print("Fetch Damages failed with status \(httpResponse.statusCode): \(errorBody)")
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
        let damagesData = try decoder.decode([DamageResponse].self, from: data)
        print("Damages fetched for property \(propertyId): \(damagesData.count)")
        
        if let index = properties.firstIndex(where: { $0.id == propertyId }) {
            properties[index].damages = damagesData
        } else {
            print("Property with ID \(propertyId) not found in properties array")
        }
        
        self.damages = damagesData
        objectWillChange.send()
    }
    
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

        let decoder = JSONDecoder()
        let propertyResponse = try decoder.decode(PropertyResponse.self, from: data)

        var property = Property(
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
            tenantName: propertyResponse.tenant,
            leaseStartDate: propertyResponse.startDate,
            leaseEndDate: propertyResponse.endDate,
            documents: [],
            createdAt: propertyResponse.createdAt,
            rooms: [],
            damages: []
        )

        do {
            property.photo = try await fetchPropertiesPicture(propertyId: property.id)
        } catch {
            print("Failed to fetch picture for tenant property \(property.id), error: \(error.localizedDescription)")
        }

        do {
            try await fetchPropertyDocuments(propertyId: property.id)
            if let updatedProperty = properties.first(where: { $0.id == property.id }) {
                property.documents = updatedProperty.documents
            }
        } catch {
            print("Failed to fetch documents for tenant property \(property.id), error: \(error.localizedDescription)")
        }

        do {
            try await fetchPropertyDamages(propertyId: property.id)
            if let updatedProperty = properties.first(where: { $0.id == property.id }) {
                property.damages = updatedProperty.damages
            }
        } catch {
            print("Failed to fetch damages for tenant property \(property.id), error: \(error.localizedDescription)")
        }

        return property
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
            print("Fetch Tenant Damages failed with status \(httpResponse.statusCode): \(errorBody)")
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
        print("Damages fetched for lease \(leaseId): \(damagesData.count)")

        if let propertyIndex = properties.firstIndex(where: { $0.id == damagesData.first?.propertyId }) {
            properties[propertyIndex].damages = damagesData
            objectWillChange.send()
        }
        self.damages = damagesData
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
            print("Fetch Active Lease ID failed with status \(httpResponse.statusCode): \(errorBody)")
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
            print("Create Damage failed with status \(httpResponse.statusCode): \(errorBody)")
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
        print("Damage created with ID: \(idResponse.id)")
        return idResponse.id
    }

    struct IdResponse: Codable {
        let id: String
    }
}

struct PropertyID: Decodable {
    let id: String
}

struct PropertyImageBase64: Decodable {
    let data: String
}

private func convertUIImageToBase64(_ image: UIImage) -> String? {
    guard let imageData = image.jpegData(compressionQuality: 0.8) else {
        print("Failed to convert UIImage to JPEG data")
        return nil
    }
    return imageData.base64EncodedString()
}
