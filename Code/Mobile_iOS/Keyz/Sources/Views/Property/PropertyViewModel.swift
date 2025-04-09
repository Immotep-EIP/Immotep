//
//  PropertyViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

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

        let jsonData = try JSONSerialization.data(withJSONObject: body)
        urlRequest.httpBody = jsonData

        let (_, response) = try await URLSession.shared.data(for: urlRequest)

        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
        }

        guard (200...299).contains(httpResponse.statusCode) else {
            if httpResponse.statusCode == 400 {
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid property data.".localized()])
            } else if httpResponse.statusCode == 401 {
                throw NSError(domain: "", code: 401, userInfo: [NSLocalizedDescriptionKey: "Unauthorized. Please check your token.".localized()])
            } else {
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode)"])
            }
        }
        await fetchProperties()
        return "Property successfully created!"
    }
    func fetchProperties() async {
            let url = URL(string: "\(APIConfig.baseURL)/owner/properties/")!

            do {
                let token = try await TokenStorage.getValidAccessToken()

                var urlRequest = URLRequest(url: url)
                urlRequest.httpMethod = "GET"
                urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")

                let (data, response) = try await URLSession.shared.data(for: urlRequest)

                guard let httpResponse = response as? HTTPURLResponse else {
                    print("Invalid response")
                    return
                }

                guard (200...299).contains(httpResponse.statusCode) else {
                    print("Error: Status code \(httpResponse.statusCode)")
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
                objectWillChange.send()
            } catch {
                print("Error fetching properties: \(error.localizedDescription)")
            }
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

            let jsonData = try JSONSerialization.data(withJSONObject: body)
            urlRequest.httpBody = jsonData

            let (data, response) = try await URLSession.shared.data(for: urlRequest)

            guard let httpResponse = response as? HTTPURLResponse else {
                throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server.".localized()])
            }

            guard (200...299).contains(httpResponse.statusCode) else {
                let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
                print("Update Property failed with status \(httpResponse.statusCode): \(errorBody)")
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
            print("Error response: \(errorBody)")
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
        // print("Token used: \(token)")

        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")
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
                let responseBody = String(data: data, encoding: .utf8) ?? "No response body"
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(responseBody)"])
            }

            let decoder = JSONDecoder()
            let documentsData = try decoder.decode([PropertyDocumentResponse].self, from: data)
//            print("Documents fetched: \(documentsData.count)")

            if let index = properties.firstIndex(where: { $0.id == propertyId }) {
                var updatedProperty = properties[index]
                updatedProperty.documents = documentsData.map { doc in
                    PropertyDocument(id: doc.id, title: doc.name, fileName: doc.name, data: doc.data)
                }
                properties[index] = updatedProperty
                objectWillChange.send()
            }
        } catch {
            print("Fetch error details: \(error.localizedDescription)")
            throw error
        }
    }
}
