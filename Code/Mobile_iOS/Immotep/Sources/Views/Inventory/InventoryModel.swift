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
    var images: [UIImage] = [] // Stocker les images
    var status: String = "Select your equipment status" // Stocker le statut
    var comment: String = "" // Stocker le commentaire

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

// Local datas, used in UI, to be used later to make the final report

struct LocalRoom: Identifiable {
    var id: String
    var name: String
    var checked: Bool
    var inventory: [LocalInventory]
}

struct LocalInventory: Identifiable {
    var id: String
    let propertyId: String
    let roomId: String
    var name: String
    var quantity: Int?
    var checked: Bool = false
    var images: [UIImage] = []
    var status: String = "Select your equipment status"
    var comment: String = ""
}

// Modèle pour la requête
struct InventoryReportRequest: Codable {
    let type: String
    let rooms: [RoomStateRequest]
}

struct RoomStateRequest: Codable {
    let id: String
    let cleanliness: String
    let state: String
    let note: String
    let pictures: [String] // Images encodées en base64
    let furnitures: [FurnitureStateRequest]
}

// Modèle pour la réponse
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
