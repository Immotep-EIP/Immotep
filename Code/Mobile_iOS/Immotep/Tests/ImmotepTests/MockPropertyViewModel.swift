//
//  MockPropertyViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 14/12/2024.
//

import Foundation
@testable import Immotep

@MainActor
class MockPropertyViewModel: ObservableObject {
    @Published var properties: [Property] = []

    func createProperty(request: Property, token: String) async throws -> String {
        try await Task.sleep(nanoseconds: 500_000_000)

        self.properties.append(request)

        return "Property successfully created!"
    }

    func fetchProperties() async {
        try? await Task.sleep(nanoseconds: 500_000_000)

        let mockProperties = [
            Property(
                id: "1",
                ownerID: "123",
                name: "Mock Property 1",
                address: "1234 Mock St",
                city: "Mock City",
                postalCode: "56789",
                country: "Mock Country",
                photo: nil,
                monthlyRent: 1000,
                deposit: 2000,
                surface: 50.0,
                isAvailable: true,
                tenantName: nil,
                leaseStartDate: nil,
                leaseEndDate: nil,
                documents: []
            ),
            Property(
                id: "2",
                ownerID: "123",
                name: "Mock Property 2",
                address: "5678 Fake Ave",
                city: "Fake City",
                postalCode: "98765",
                country: "Fake Country",
                photo: nil,
                monthlyRent: 1200,
                deposit: 2400,
                surface: 60.0,
                isAvailable: false,
                tenantName: "John Doe",
                leaseStartDate: Date(),
                leaseEndDate: nil,
                documents: []
            )
        ]

        self.properties = mockProperties
    }
}
