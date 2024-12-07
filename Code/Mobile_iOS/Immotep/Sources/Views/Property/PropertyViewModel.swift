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

        let url = URL(string: "\(baseURL)/owner/properties")!

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

        return "Property successfully created!"
    }

    func fetchProperties() async {
        let url = URL(string: "\(baseURL)/owner/properties")!

        guard let token = await TokenStorage.getAccessToken() else {
            print("Token is nil")
            return
        }

        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")

        do {
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
                    isAvailable: true,
                    tenantName: nil,
                    leaseStartDate: nil,
                    leaseEndDate: nil,
                    documents: []
                )
            }
        } catch {
            print("Error fetching properties: \(error.localizedDescription)")
        }
    }

}
