//
//  InventoryModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI

struct PropertyRooms: Identifiable {
    var id: String
    var name: String
    var checked: Bool
    var inventory: [RoomInventory]
}

struct RoomInventory: Identifiable {
    var id: String
    var name: String
    var number: Int?
    var state: String?
    var image: String?
    var description: String?
    var checked: Bool
}
