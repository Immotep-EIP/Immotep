//
//  PropertySampleData.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 14/12/2024.
//

import Foundation
import SwiftUI

let exampleDataProperty = Property(
    id: "1",
    ownerID: "123",
    name: "Sunny Apartment",
    address: "1234 Elm Street",
    city: "Paris",
    postalCode: "75001",
    country: "France",
    photo: nil,
    monthlyRent: 1500,
    deposit: 3000,
    surface: 60.5,
    isAvailable: true,
    tenantName: "John Doe",
    leaseStartDate: stringToDate("10/12/2024"),
    leaseEndDate: nil,
    documents: [
        PropertyDocument(id: UUID(), title: "Lease Agreement", fileName: "lease_agreement.pdf"),
        PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
        PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
        PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
        PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
        PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf")

    ],
    createdAt: "2024-12-10"
)
