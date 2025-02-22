//
//  InventoryModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI
import Foundation

struct RoomResponse: Codable {
    let id: String
    let name: String
}

struct FurnitureRequest: Codable {
    let name: String
}

struct FurnitureResponse: Codable, Identifiable {
    let id: String
    let propertyId: String
    let roomId: String
    let name: String
    let quantity: Int

    enum CodingKeys: String, CodingKey {
        case id
        case propertyId = "property_id"
        case roomId = "room_id"
        case name
        case quantity
    }
}

struct SummarizeRequest: Codable {
    let id: String
    let pictures: [String]
    let type: String
}

struct SummarizeResponse: Codable {
    let cleanliness: String
    let note: String
    let state: String
}

struct LocalRoom: Identifiable, Equatable {
    var id: String
    var name: String
    var checked: Bool
    var inventory: [LocalInventory]

    static func == (lhs: LocalRoom, rhs: LocalRoom) -> Bool {
        return lhs.id == rhs.id &&
               lhs.name == rhs.name &&
               lhs.checked == rhs.checked &&
               lhs.inventory == rhs.inventory
    }
}

struct LocalInventory: Identifiable, Equatable {
    var id: String
    let propertyId: String
    let roomId: String
    var name: String
    var quantity: Int?
    var checked: Bool = false
    var images: [UIImage] = []
    var status: String = "Select your equipment status"
    var comment: String = ""

    static func == (lhs: LocalInventory, rhs: LocalInventory) -> Bool {
        return lhs.id == rhs.id &&
               lhs.propertyId == rhs.propertyId &&
               lhs.roomId == rhs.roomId &&
               lhs.name == rhs.name &&
               lhs.quantity == rhs.quantity &&
               lhs.checked == rhs.checked &&
               lhs.status == rhs.status &&
               lhs.comment == rhs.comment
    }
}

struct InventoryReportRequest: Codable {
    let type: String
    let rooms: [RoomStateRequest]
}

struct RoomStateRequest: Codable {
    let id: String
    let cleanliness: String
    let state: String
    let note: String
    let pictures: [String]
    let furnitures: [FurnitureStateRequest]
}

struct InventoryReportResponse: Codable {
    let date: String
    let id: String
    let propertyId: String
    let rooms: [RoomStateResponse]
    let type: String
}

struct RoomStateResponse: Codable {
    let id: String
    let cleanliness: String
    let state: String
    let note: String
    let pictures: [String]
}

struct FurnitureStateRequest: Codable {
    let id: String
    let cleanliness: String
    let note: String
    let pictures: [String]
    let state: String
}

struct LastInventoryReportResponse: Codable {
    let id: String
    let propertyId: String
    let date: String
    let type: String
    let rooms: [LastRoomStateResponse]

    enum CodingKeys: String, CodingKey {
        case id
        case propertyId = "property_id"
        case date
        case type
        case rooms
    }
}

struct LastRoomStateResponse: Codable {
    let id: String
    let name: String
    let cleanliness: String
    let state: String
    let note: String
    let pictures: [String]
    let furnitures: [FurnitureStateResponse]
}

struct FurnitureStateResponse: Codable {
    let id: String
    let name: String
    let quantity: Int
    let state: String
    let cleanliness: String
    let note: String
    let pictures: [String]
}
