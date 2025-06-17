//
//  InventoryViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI

@MainActor
class InventoryViewModel: ObservableObject {
    @Published var property: Property
    @Published var isEntryInventory: Bool = true
    @Published var isLoading: Bool = false
    @Published var errorMessage: String?
    @Published var completionMessage: String?

    @Published var selectedRoom: LocalRoom?
    @Published var selectedInventory: [LocalInventory] = []
    @Published var selectedStuff: LocalInventory?
    @Published var selectedImages: [UIImage] = []
    @Published var comment: String = ""
    @Published var selectedStatus: String = "Select your equipment status"

    @Published var roomToDelete: LocalRoom?
    @Published var showDeleteConfirmation: Bool = false

    @Published var checkedStuffStatus: [String: Bool] = [:]
    @Published var localRooms: [LocalRoom]
    @Published var lastReportId: String?

    private var roomManager: RoomManager?
    private var furnitureManager: FurnitureManager?
    private var reportManager: InventoryReportManager?
    
    var onInventoryFinalized: (() -> Void)?

    init(property: Property, isEntryInventory: Bool = true, localRooms: [LocalRoom]? = nil) {
        self.property = property
        self.isEntryInventory = isEntryInventory
        self.localRooms = localRooms ?? []

        self.roomManager = RoomManager(viewModel: self)
        self.furnitureManager = FurnitureManager(viewModel: self)
        self.reportManager = InventoryReportManager(viewModel: self)
    }

    func getToken() async -> String? {
        do {
            let token = try await TokenStorage.getValidAccessToken()
            if token.isEmpty {
                print("Token is empty")
            }
            return token
        } catch {
            print("Error fetching token: \(error.localizedDescription)")
            return nil
        }
    }

    func fetchRooms() async {
        await roomManager?.fetchRooms()
    }

    func addRoom(name: String, type: String) async throws {
        try await roomManager?.addRoom(name: name, type: type)
    }

    func deleteRoom(_ room: LocalRoom) async {
        await roomManager?.deleteRoom(room)
    }

    func selectRoom(_ room: LocalRoom) {
        roomManager?.selectRoom(room)
    }

    func isRoomCompleted(_ room: LocalRoom) -> Bool {
        roomManager?.isRoomCompleted(room) ?? false
    }

    func areAllRoomsCompleted() -> Bool {
        roomManager?.areAllRoomsCompleted() ?? false
    }

    func markRoomAsChecked(_ room: LocalRoom) async {
        await roomManager?.markRoomAsChecked(room)
    }

    func markStuffAsChecked(_ stuff: LocalInventory) async throws {
        try await furnitureManager?.markStuffAsChecked(stuff)
    }

    func fetchStuff(_ room: LocalRoom) async {
        await furnitureManager?.fetchStuff(room)
    }

    func addStuff(name: String, quantity: Int, to room: LocalRoom) async throws {
        try await furnitureManager?.addStuff(name: name, quantity: quantity, to: room)
    }

    func deleteStuff(_ stuff: LocalInventory, from room: LocalRoom) async {
        await furnitureManager?.deleteStuff(stuff, from: room)
    }

    func selectStuff(_ stuff: LocalInventory) {
        furnitureManager?.selectStuff(stuff)
    }

    func sendStuffReport() async throws {
        try await reportManager?.sendStuffReport()
    }

    func finalizeInventory() async throws {
        try await reportManager?.finalizeInventory()
        onInventoryFinalized?()
    }

    func fetchLastInventoryReport() async {
        await reportManager?.fetchLastInventoryReport()
    }

    func compareStuffReport(oldReportId: String) async throws {
        try await reportManager?.compareStuffReport(oldReportId: oldReportId)
    }
    
    func sendRoomReport() async throws {
        try await reportManager?.sendRoomReport()
    }

    func compareRoomReport(oldReportId: String) async throws {
        try await reportManager?.compareRoomReport(oldReportId: oldReportId)
    }
}
