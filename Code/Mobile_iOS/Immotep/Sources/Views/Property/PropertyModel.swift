//
//  PropertyModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import Foundation
import SwiftUI

struct Property: Identifiable {
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
    var isAvailable: Bool
    var tenantName: String?
    var leaseStartDate: Date?
    var leaseEndDate: Date?
    var documents: [PropertyDocument]
    var createdAt: String?
}

struct PropertyDocument: Identifiable {
    let id: UUID
    var title: String
    var fileName: String
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
    }
}
