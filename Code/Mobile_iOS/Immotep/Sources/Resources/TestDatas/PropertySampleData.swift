//
//  PropertySampleData.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 14/12/2024.
//

import Foundation
import SwiftUI

func stringToDate(_ string: String) -> Date? {
    let formatter = DateFormatter()
    formatter.dateFormat = "dd/MM/yyyy"
    return formatter.date(from: string)
}

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
//        PropertyDocument(id: UUID(), title: "Lease Agreement", fileName: "lease_agreement.pdf"),
//        PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
//        PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
//        PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
//        PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
//        PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf")
    ],
    createdAt: "2024-12-10",
    rooms: [
        PropertyRooms(
            id: "1",
            name: "Entrance",
            checked: true,
            inventory: [
                RoomInventory(
                    id: "1.1",
                    propertyId: "1",
                    roomId: "1",
                    name: "Right Wall",
                    quantity: 1,
                    checked: true,
                    images: [],
                    status: "Good",
                    comment: "No visible damage"
                ),
                RoomInventory(
                    id: "1.2",
                    propertyId: "1",
                    roomId: "1",
                    name: "Left Wall",
                    quantity: 1,
                    checked: true,
                    images: [],
                    status: "Used",
                    comment: "Minor scratches"
                ),
                RoomInventory(
                    id: "1.3",
                    propertyId: "1",
                    roomId: "1",
                    name: "Back Wall",
                    quantity: 1,
                    checked: true,
                    images: [],
                    status: "Used",
                    comment: "Faded paint"
                ),
                RoomInventory(
                    id: "1.4",
                    propertyId: "1",
                    roomId: "1",
                    name: "Door Wall",
                    quantity: 1,
                    checked: true,
                    images: [],
                    status: "Good",
                    comment: "Recently painted"
                ),
                RoomInventory(
                    id: "1.5",
                    propertyId: "1",
                    roomId: "1",
                    name: "Shutters",
                    quantity: 2,
                    checked: false,
                    images: [],
                    status: "Broken",
                    comment: "One shutter jammed"
                ),
                RoomInventory(
                    id: "1.6",
                    propertyId: "1",
                    roomId: "1",
                    name: "Ground",
                    quantity: nil,
                    checked: false,
                    images: [],
                    status: "Good",
                    comment: "Clean tiles"
                )
            ]
        ),
        PropertyRooms(
            id: "2",
            name: "Kitchen",
            checked: false,
            inventory: [
                RoomInventory(
                    id: "2.1",
                    propertyId: "1",
                    roomId: "2",
                    name: "Right Wall",
                    quantity: 1,
                    checked: true,
                    images: [],
                    status: "Good",
                    comment: "No issues"
                ),
                RoomInventory(
                    id: "2.2",
                    propertyId: "1",
                    roomId: "2",
                    name: "Left Wall",
                    quantity: 1,
                    checked: false,
                    images: [],
                    status: "Used",
                    comment: "Some stains"
                ),
                RoomInventory(
                    id: "2.3",
                    propertyId: "1",
                    roomId: "2",
                    name: "Back Wall",
                    quantity: 1,
                    checked: true,
                    images: [],
                    status: "Used",
                    comment: "Minor wear"
                ),
                RoomInventory(
                    id: "2.4",
                    propertyId: "1",
                    roomId: "2",
                    name: "Door Wall",
                    quantity: 1,
                    checked: false,
                    images: [],
                    status: "Good",
                    comment: "Clean"
                ),
                RoomInventory(
                    id: "2.5",
                    propertyId: "1",
                    roomId: "2",
                    name: "Shutters",
                    quantity: 2,
                    checked: true,
                    images: [],
                    status: "Broken",
                    comment: "Needs repair"
                ),
                RoomInventory(
                    id: "2.6",
                    propertyId: "1",
                    roomId: "2",
                    name: "Ground",
                    quantity: nil,
                    checked: true,
                    images: [],
                    status: "Good",
                    comment: "New tiles"
                )
            ]
        ),
        PropertyRooms(
            id: "3",
            name: "Corridor",
            checked: false,
            inventory: [
                RoomInventory(
                    id: "3.1",
                    propertyId: "1",
                    roomId: "3",
                    name: "Right Wall",
                    quantity: 1,
                    checked: false,
                    images: [],
                    status: "Good",
                    comment: "No damage"
                ),
                RoomInventory(
                    id: "3.2",
                    propertyId: "1",
                    roomId: "3",
                    name: "Left Wall",
                    quantity: 1,
                    checked: false,
                    images: [],
                    status: "Used",
                    comment: "Slight wear"
                ),
                RoomInventory(
                    id: "3.3",
                    propertyId: "1",
                    roomId: "3",
                    name: "Back Wall",
                    quantity: 1,
                    checked: true,
                    images: [],
                    status: "Used",
                    comment: "Faded color"
                ),
                RoomInventory(
                    id: "3.4",
                    propertyId: "1",
                    roomId: "3",
                    name: "Door Wall",
                    quantity: 1,
                    checked: true,
                    images: [],
                    status: "Good",
                    comment: "Good condition"
                ),
                RoomInventory(
                    id: "3.5",
                    propertyId: "1",
                    roomId: "3",
                    name: "Shutters",
                    quantity: 2,
                    checked: true,
                    images: [],
                    status: "Broken",
                    comment: "One shutter missing"
                ),
                RoomInventory(
                    id: "3.6",
                    propertyId: "1",
                    roomId: "3",
                    name: "Ground",
                    quantity: nil,
                    checked: true,
                    images: [],
                    status: "Good",
                    comment: "Polished wood"
                )
            ]
        )
    ]
)
