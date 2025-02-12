//
//  InventoryModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI

struct RoomResponse: Codable {
    let id: String
    let name: String
}

struct PropertyRooms: Identifiable {
    var id: String
    var name: String
    var checked: Bool
    var inventory: [RoomInventory]
}

struct RoomInventory: Identifiable {
    var id: String
    let propertyId: String
    let roomId: String
    var name: String
    var quantity: Int?
    var checked: Bool = false

    enum CodingKeys: String, CodingKey {
        case id
        case propertyId = "property_id"
        case roomId = "room_id"
        case name
        case quantity
    }
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
