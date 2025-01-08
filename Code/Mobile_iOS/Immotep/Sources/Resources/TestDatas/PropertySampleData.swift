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
    createdAt: "2024-12-10",
    rooms: [
        PropertyRooms(
            id: "1",
            name: "Entrance",
            checked: true,
            inventory: [
                RoomInventory(id: "1.1", name: "Right Wall", number: nil, state: "Good", image: nil, description: nil, checked: true),
                RoomInventory(id: "1.2", name: "Left Wall", number: nil, state: "Used", image: nil, description: nil, checked: true),
                RoomInventory(id: "1.3", name: "Back Wall", number: nil, state: "Used", image: nil, description: nil, checked: true),
                RoomInventory(id: "1.4", name: "Door Wall", number: nil, state: "Good", image: nil, description: nil, checked: true),
                RoomInventory(id: "1.5", name: "Shutters", number: 2, state: "Broken", image: nil, description: nil, checked: false),
                RoomInventory(id: "1.6", name: "Ground", number: nil, state: "Good", image: nil, description: nil, checked: false)
            ]
        ),
        PropertyRooms(
            id: "2",
            name: "Kitchen",
            checked: false,
            inventory: [
                RoomInventory(id: "2.1", name: "Right Wall", number: nil, state: "Good", image: nil, description: nil, checked: true),
                RoomInventory(id: "2.2", name: "Left Wall", number: nil, state: "Used", image: nil, description: nil, checked: false),
                RoomInventory(id: "2.3", name: "Back Wall", number: nil, state: "Used", image: nil, description: nil, checked: true),
                RoomInventory(id: "2.4", name: "Door Wall", number: nil, state: "Good", image: nil, description: nil, checked: false),
                RoomInventory(id: "2.5", name: "Shutters", number: 2, state: "Broken", image: nil, description: nil, checked: true),
                RoomInventory(id: "2.6", name: "Ground", number: nil, state: "Good", image: nil, description: nil, checked: true)
            ]
        ),
        PropertyRooms(
            id: "3",
            name: "Corridor",
            checked: false,
            inventory: [
                RoomInventory(id: "3.1", name: "Right Wall", number: nil, state: "Good", image: nil, description: nil, checked: false),
                RoomInventory(id: "3.2", name: "Left Wall", number: nil, state: "Used", image: nil, description: nil, checked: false),
                RoomInventory(id: "3.3", name: "Back Wall", number: nil, state: "Used", image: nil, description: nil, checked: true),
                RoomInventory(id: "3.4", name: "Door Wall", number: nil, state: "Good", image: nil, description: nil, checked: true),
                RoomInventory(id: "3.5", name: "Shutters", number: 2, state: "Broken", image: nil, description: nil, checked: true),
                RoomInventory(id: "3.6", name: "Ground", number: nil, state: "Good", image: nil, description: nil, checked: true)
            ]
        )
    ]
)
