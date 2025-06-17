//
//  OverviewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 21/10/2024.
//

import Foundation

struct DashboardResponse: Decodable, Equatable {
    let reminders: [Reminder]
    let properties: PropertyStats
    let openDamages: OpenDamageStats

    enum CodingKeys: String, CodingKey {
        case reminders
        case properties
        case openDamages = "open_damages"
    }
}

struct Reminder: Decodable, Identifiable, Equatable {
    let id: String
    let priority: String
    let title: String
    let advice: String
    let link: String
}

struct PropertyStats: Decodable, Equatable {
    let total: Int
    let archived: Int
    let occupied: Int
    let available: Int
    let pendingInvites: Int
    let recentlyAdded: [PropertySummary]

    enum CodingKeys: String, CodingKey {
        case total = "nbr_total"
        case archived = "nbr_archived"
        case occupied = "nbr_occupied"
        case available = "nbr_available"
        case pendingInvites = "nbr_pending_invites"
        case recentlyAdded = "list_recently_added"
    }
}

struct PropertySummary: Decodable, Identifiable, Equatable {
    let id: String
    let name: String
    let address: String
    let city: String
    let postalCode: String
    let country: String
    let areaSqm: Double
    let rentalPricePerMonth: Int
    let depositPrice: Int
    let createdAt: String
    let archived: Bool
    let ownerId: String

    enum CodingKeys: String, CodingKey {
        case id, name, address, city, country, archived
        case postalCode = "postal_code"
        case areaSqm = "area_sqm"
        case rentalPricePerMonth = "rental_price_per_month"
        case depositPrice = "deposit_price"
        case createdAt = "created_at"
        case ownerId = "owner_id"
    }
}

struct OpenDamageStats: Decodable, Equatable {
    let total: Int
    let urgent: Int
    let high: Int
    let medium: Int
    let low: Int
    let plannedThisWeek: Int
    let toFix: [DamageSummary]

    enum CodingKeys: String, CodingKey {
        case total = "nbr_total"
        case urgent = "nbr_urgent"
        case high = "nbr_high"
        case medium = "nbr_medium"
        case low = "nbr_low"
        case plannedThisWeek = "nbr_planned_to_fix_this_week"
        case toFix = "list_to_fix"
    }
}

struct DamageSummary: Decodable, Identifiable, Equatable {
    let id: String
    let leaseId: String
    let tenantName: String
    let propertyId: String
    let propertyName: String
    let roomId: String
    let roomName: String
    let comment: String
    let priority: String
    let read: Bool
    let createdAt: String
    let updatedAt: String
    let fixStatus: String
    let fixPlannedAt: String?

    enum CodingKeys: String, CodingKey {
        case id, comment, priority, read
        case leaseId = "lease_id"
        case tenantName = "tenant_name"
        case propertyId = "property_id"
        case propertyName = "property_name"
        case roomId = "room_id"
        case roomName = "room_name"
        case createdAt = "created_at"
        case updatedAt = "updated_at"
        case fixStatus = "fix_status"
        case fixPlannedAt = "fix_planned_at"
    }
}
