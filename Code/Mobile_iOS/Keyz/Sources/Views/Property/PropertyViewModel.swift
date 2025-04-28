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
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)"])
            }
        }

        let propertyID = try JSONDecoder().decode(PropertyID.self, from: data)

        if let photo = request.photo {
            do {
                let result = try await updatePropertyPicture(token: token, propertyPicture: photo, propertyID: propertyID.id)
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
                    rooms: []
                )
            }

            await withTaskGroup(of: (index: Int, image: UIImage?).self) { group in
                for index in properties.indices {
                    group.addTask {
                        let propertyId = await self.properties[index].id
                        do {
                            let image = try await self.fetchPropertiesPicture(propertyId: propertyId)
                            return (index: index, image: image)
                        } catch {
                            return (index: index, image: nil)
                        }
                    }
                }

                for await (index, image) in group {
                    self.properties[index].photo = image
                }
            }
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
                    if let jsonString = String(data: data, encoding: .utf8) {
                        print("Raw JSON response for property \(propertyId): \(jsonString)")
                    } else {
                        print("Failed to decode raw JSON response for property \(propertyId)")
                    }

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
                        print("Failed to decode JSON for property \(propertyId): \(error.localizedDescription)")
                        throw error
                    }
                case 204:
                    print("No picture associated with property \(propertyId)")
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
                    print("Property \(propertyId) not found")
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
            print("Request body: \(String(data: jsonData, encoding: .utf8) ?? "Invalid JSON")")
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
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(errorBody)"])
            }

            let decoder = JSONDecoder()
            let documentsData = try decoder.decode([PropertyDocumentResponse].self, from: data)

            if let index = properties.firstIndex(where: { $0.id == propertyId }) {
                var updatedProperty = properties[index]
                updatedProperty.documents = documentsData.map { doc in
                    PropertyDocument(id: doc.id, title: doc.name, fileName: doc.name, data: doc.data)
                }
                properties[index] = updatedProperty
                objectWillChange.send()
            }
        } catch {
            throw error
        }
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
