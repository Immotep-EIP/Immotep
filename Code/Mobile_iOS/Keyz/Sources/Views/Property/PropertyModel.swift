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
    var leaseId: String?
    var leaseStartDate: String?
    var leaseEndDate: String?
    var documents: [PropertyDocument]
    var createdAt: String?
    var rooms: [PropertyRooms]
    var damages: [DamageResponse]

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

struct PropertyDocument: Identifiable {
    let id: String
    let title: String
    let fileName: String
    let data: String
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
    let createdAt: String
    let isAvailable: String
    let lease: LeaseInfo?

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
        case createdAt = "created_at"
        case isAvailable = "status"
        case lease
    }
}

struct LeaseResponse: Codable {
    let id: String
    let propertyId: String
    let propertyName: String
    let ownerId: String
    let ownerName: String
    let ownerEmail: String
    let tenantId: String
    let tenantName: String
    let tenantEmail: String
    let active: Bool
    let startDate: Date
    let endDate: Date?
    let createdAt: Date
    
    enum CodingKeys: String, CodingKey {
        case id, active
        case propertyId = "property_id"
        case propertyName = "property_name"
        case ownerId = "owner_id"
        case ownerName = "owner_name"
        case ownerEmail = "owner_email"
        case tenantId = "tenant_id"
        case tenantName = "tenant_name"
        case tenantEmail = "tenant_email"
        case startDate = "start_date"
        case endDate = "end_date"
        case createdAt = "created_at"
    }
}

struct DamageResponse: Codable, Equatable {
    let id: String
    let comment: String
    let priority: String
    let roomName: String
    let fixStatus: String
    let pictures: [String]?
    let createdAt: String
    let updatedAt: String?
    let fixPlannedAt: String?
    let fixedAt: String?
    let leaseId: String
    let propertyId: String
    let propertyName: String
    let tenantName: String?
    let read: Bool

    enum CodingKeys: String, CodingKey {
        case id
        case comment
        case priority
        case roomName = "room_name"
        case fixStatus = "fix_status"
        case pictures
        case createdAt = "created_at"
        case updatedAt = "updated_at"
        case fixPlannedAt = "fix_planned_at"
        case fixedAt = "fixed_at"
        case leaseId = "lease_id"
        case propertyId = "property_id"
        case propertyName = "property_name"
        case tenantName = "tenant_name"
        case read
    }
}

struct DamageRequest: Codable {
    let comment: String
    let priority: String
    let roomId: String
    let pictures: [String]?
    
    enum CodingKeys: String, CodingKey {
        case comment
        case priority
        case roomId = "room_id"
        case pictures
    }
}


struct LeaseInfo: Codable {
    let id: String
    let tenantName: String
    let tenantEmail: String
    let active: Bool
    let startDate: String?
    let endDate: String?

    enum CodingKeys: String, CodingKey {
        case id
        case tenantName = "tenant_name"
        case tenantEmail = "tenant_email"
        case active
        case startDate = "start_date"
        case endDate = "end_date"
    }
}


extension Property {
    func copyWith(
        id: String? = nil,
        ownerID: String? = nil,
        name: String? = nil,
        address: String? = nil,
        city: String? = nil,
        postalCode: String? = nil,
        country: String? = nil,
        photo: UIImage?? = nil,
        monthlyRent: Int? = nil,
        deposit: Int? = nil,
        surface: Double? = nil,
        isAvailable: String? = nil,
        tenantName: String?? = nil,
        leaseId: String?? = nil,
        leaseStartDate: String?? = nil,
        leaseEndDate: String?? = nil,
        documents: [PropertyDocument]? = nil,
        createdAt: String?? = nil,
        rooms: [PropertyRooms]? = nil,
        damages: [DamageResponse]? = nil
    ) -> Property {
        return Property(
            id: id ?? self.id,
            ownerID: ownerID ?? self.ownerID,
            name: name ?? self.name,
            address: address ?? self.address,
            city: city ?? self.city,
            postalCode: postalCode ?? self.postalCode,
            country: country ?? self.country,
            photo: photo ?? self.photo,
            monthlyRent: monthlyRent ?? self.monthlyRent,
            deposit: deposit ?? self.deposit,
            surface: surface ?? self.surface,
            isAvailable: isAvailable ?? self.isAvailable,
            tenantName: tenantName ?? self.tenantName,
            leaseId: leaseId ?? self.leaseId,
            leaseStartDate: leaseStartDate ?? self.leaseStartDate,
            leaseEndDate: leaseEndDate ?? self.leaseEndDate,
            documents: documents ?? self.documents,
            createdAt: createdAt ?? self.createdAt,
            rooms: rooms ?? self.rooms,
            damages: damages ?? self.damages
        )
    }
}
