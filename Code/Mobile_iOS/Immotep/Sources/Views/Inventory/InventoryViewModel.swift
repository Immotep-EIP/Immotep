//
//  InventoryViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI

class InventoryViewModel: ObservableObject {
    @Published var property: Property
    @Published var isEntryInventory: Bool = true
    @Published var isLoading: Bool = false
    @Published var errorMessage: String?

    @Published var selectedRoom: PropertyRooms?

    @Published var selectedInventory: [RoomInventory] = []

    @Published var selectedStuff: RoomInventory?
    @Published var selectedImages: [UIImage] = []
    @Published var comment: String = ""
    @Published var selectedStatus: String = "Select your equipment status"

    init(property: Property) {
        self.property = property
    }

    func fetchRooms() {
    }

    func addRoom(name: String) {
    }

    func selectRoom(_ room: PropertyRooms) {
        selectedRoom = room
        selectedInventory = room.inventory
    }

    func selectStuff(_ stuff: RoomInventory) {
        selectedStuff = stuff
        selectedImages = []
        comment = ""
        selectedStatus = "Select your equipment status"
    }

    func addStuff(name: String) {
    }
}
