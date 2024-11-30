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

    init() {
        loadMockData()
    }

    func loadMockData() {
        properties = [
            Property(
                id: UUID(),
                address: "4391 Hedge Street, New Jersey",
                postalCode: "07102",
                country: "USA",
                photo: nil,
                monthlyRent: 1200.0,
                deposit: 2400.0,
                surface: 80.0,
                isAvailable: false,
                tenantName: "John & Mary Doe",
                leaseStartDate: Date(),
                leaseEndDate: Calendar.current.date(byAdding: .year, value: 1, to: Date()),
                documents: [
                    PropertyDocument(id: UUID(), title: "Lease Agreement", fileName: "lease_agreement.pdf"),
                    PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf")
                ]
            ),
            Property(
                id: UUID(),
                address: "4391 Hedge Street, New Jersey",
                postalCode: "07102",
                country: "USA",
                photo: nil,
                monthlyRent: 950.0,
                deposit: 1900.0,
                surface: 65.0,
                isAvailable: true,
                tenantName: nil,
                leaseStartDate: nil,
                leaseEndDate: nil,
                documents: []
            )
        ]
    }

    func addProperty(_ property: Property) {
        properties.append(property)
    }
}
