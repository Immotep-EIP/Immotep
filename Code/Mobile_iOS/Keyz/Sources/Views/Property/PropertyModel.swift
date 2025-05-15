//
//  PropertyModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import Foundation
import SwiftUI

struct Property: Identifiable, Equatable {
    let id: String
    let ownerID: String
    var name: String
    var address: String
    var city: String
    var postalCode: String
    var country: String
    var photo: UIImage?
    var monthlyRent: Int
    var deposit: Int
    var surface: Double
    var isAvailable: String
    var tenantName: String?
    var leaseStartDate: String?
    var leaseEndDate: String?
    var documents: [PropertyDocument]
    var createdAt: String?
    var rooms: [PropertyRooms]

    static func == (lhs: Property, rhs: Property) -> Bool {
        return lhs.id == rhs.id
    }
}

struct PropertyDocumentResponse: Codable {
    let id: String
    let name: String
    let data: String
    let createdAt: String

    enum CodingKeys: String, CodingKey {
        case id
        case name
        case data
        case createdAt = "created_at"
    }
}

struct PropertyDocument: Identifiable, Equatable {
    let id: String
    var title: String
    var fileName: String
    let data: String

    static func == (lhs: PropertyDocument, rhs: PropertyDocument) -> Bool {
        return lhs.id == rhs.id
    }
}

struct PropertyRooms: Identifiable, Equatable {
    let id: String
    var name: String
    var checked: Bool
    var inventory: [RoomInventory]

    static func == (lhs: PropertyRooms, rhs: PropertyRooms) -> Bool {
        return lhs.id == rhs.id
    }
}

struct RoomInventory: Identifiable, Equatable {
    let id: String
    let propertyId: String
    let roomId: String
    var name: String
    var quantity: Int?
    var checked: Bool
    var images: [UIImage]
    var status: String
    var comment: String

    enum CodingKeys: String, CodingKey {
        case id
        case propertyId = "property_id"
        case roomId = "room_id"
        case name
        case quantity
    }

    static func == (lhs: RoomInventory, rhs: RoomInventory) -> Bool {
        return lhs.id == rhs.id
    }
}

struct PropertyResponse: Codable {
    let id: String
    let ownerId: String
    let name: String
    let address: String
    let city: String
    let postalCode: String
    let country: String
    let areaSqm: Double
    let rentalPricePerMonth: Int
    let depositPrice: Int
    let picture: String?
    let createdAt: String
    let isAvailable: String
    let tenant: String?
    let startDate: String?
    let endDate: String?

    enum CodingKeys: String, CodingKey {
        case id
        case ownerId = "owner_id"
        case name
        case address
        case city
        case postalCode = "postal_code"
        case country
        case areaSqm = "area_sqm"
        case rentalPricePerMonth = "rental_price_per_month"
        case depositPrice = "deposit_price"
        case picture
        case createdAt = "created_at"
        case isAvailable = "status"
        case tenant = "tenant"
        case startDate = "start_date"
        case endDate = "end_date"
    }
}

struct LeaseResponse: Codable {
    let id: String
    let propertyId: String
    let startDate: String?
    let endDate: String?

    enum CodingKeys: String, CodingKey {
        case id
        case propertyId = "property_id"
        case startDate = "start_date"
        case endDate = "end_date"
    }
}
