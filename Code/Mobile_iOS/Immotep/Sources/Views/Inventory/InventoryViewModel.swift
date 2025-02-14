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

    @Published var selectedRoom: LocalRoom?
    @Published var selectedInventory: [LocalInventory] = []

    @Published var selectedStuff: LocalInventory?
    @Published var selectedImages: [UIImage] = []
    @Published var comment: String = ""
    @Published var selectedStatus: String = "Select your equipment status"

    @Published var roomToDelete: LocalRoom?
    @Published var showDeleteConfirmation: Bool = false

    @Published var checkedStuffStatus: [String: Bool] = [:]
    @Published var localRooms: [LocalRoom] = []

    init(property: Property) {
        self.property = property
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

    // ROOM API CALLS

    func fetchRooms() async {
        guard localRooms.isEmpty else {
//            print("localRooms is already populated, skipping fetch.")
            return
        }

        guard let url = URL(string: "\(baseURL)/owner/properties/\(property.id)/rooms/") else {
            errorMessage = "Invalid URL"
            return
        }

        guard let token = await getToken() else {
            errorMessage = "Failed to retrieve token"
            return
        }

        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")

        do {
            let (data, response) = try await URLSession.shared.data(for: urlRequest)

            guard let httpResponse = response as? HTTPURLResponse else {
                errorMessage = "Invalid response"
                return
            }

            guard (200...299).contains(httpResponse.statusCode) else {
                errorMessage = "Error: Status code \(httpResponse.statusCode)"
                if let responseBody = String(data: data, encoding: .utf8) {
                    print("Response Body: \(responseBody)")
                }
                return
            }

            let decoder = JSONDecoder()
            let roomsData = try decoder.decode([RoomResponse].self, from: data)

            property.rooms = roomsData.map { roomResponse in
                PropertyRooms(
                    id: roomResponse.id,
                    name: roomResponse.name,
                    checked: false,
                    inventory: []
                )
            }

            if localRooms.isEmpty {
                localRooms = property.rooms.map { room in
                    LocalRoom(
                        id: room.id,
                        name: room.name,
                        checked: room.checked,
                        inventory: room.inventory.map { inventory in
                            LocalInventory(
                                id: inventory.id,
                                propertyId: inventory.propertyId,
                                roomId: inventory.roomId,
                                name: inventory.name,
                                quantity: inventory.quantity,
                                checked: inventory.checked,
                                images: inventory.images,
                                status: inventory.status,
                                comment: inventory.comment
                            )
                        }
                    )
                }
            }
            dump(localRooms)
        } catch {
            errorMessage = "Error fetching rooms: \(error.localizedDescription)"
        }
    }

    func addRoom(name: String) async throws {
        guard let url = URL(string: "\(baseURL)/owner/properties/\(property.id)/rooms/") else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid URL"])
        }

        guard let token = await getToken() else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to retrieve token"])
        }

        let body: [String: Any] = [
            "name": name
        ]

        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "POST"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")

        do {
            let jsonData = try JSONSerialization.data(withJSONObject: body)
            urlRequest.httpBody = jsonData

            let (data, response) = try await URLSession.shared.data(for: urlRequest)

            guard let httpResponse = response as? HTTPURLResponse else {
                throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server"])
            }

            guard (200...299).contains(httpResponse.statusCode) else {
                if let responseBody = String(data: data, encoding: .utf8) {
                    print("Response Body: \(responseBody)")
                }
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode)"])
            }

            await fetchRooms()

            if let newRoom = property.rooms.last {
                localRooms.append(LocalRoom(
                    id: newRoom.id,
                    name: newRoom.name,
                    checked: newRoom.checked,
                    inventory: []
                ))
            }
        } catch {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Error creating room: \(error.localizedDescription)"])
        }
    }

    func deleteRoom(_ room: LocalRoom) async {
        guard let url = URL(string: "\(baseURL)/owner/properties/\(property.id)/rooms/\(room.id)/") else {
            errorMessage = "Invalid URL"
            return
        }

        guard let token = await getToken() else {
            errorMessage = "Failed to retrieve token"
            return
        }

        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "DELETE"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")

        do {
            let (_, response) = try await URLSession.shared.data(for: urlRequest)

            guard let httpResponse = response as? HTTPURLResponse else {
                errorMessage = "Invalid response"
                return
            }

            guard (200...299).contains(httpResponse.statusCode) else {
                errorMessage = "Error: Status code \(httpResponse.statusCode)"
                return
            }

            await fetchRooms()

            localRooms.removeAll { $0.id == room.id }
        } catch {
            errorMessage = "Error deleting room: \(error.localizedDescription)"
        }
    }

    func selectRoom(_ room: LocalRoom) {
//        print("room selected: \(room)")
        selectedRoom = room
        if let roomIndex = localRooms.firstIndex(where: { $0.id == room.id }) {
            selectedInventory = localRooms[roomIndex].inventory
        }
    }

    func isRoomCompleted(_ room: LocalRoom) -> Bool {
        return room.inventory.allSatisfy { $0.checked }
    }

    func areAllRoomsCompleted() -> Bool {
        return localRooms.allSatisfy { $0.checked }
    }

    func markRoomAsChecked(_ room: LocalRoom) async {
        guard let index = localRooms.firstIndex(where: { $0.id == room.id }) else { return }
        localRooms[index].checked = true

        await MainActor.run {
            self.localRooms[index].checked = true
        }
    }

    func updateRoomCheckedStatus() {
        guard let selectedRoom = selectedRoom else { return }

        let allStuffChecked = selectedInventory.allSatisfy { $0.checked }
//        for stuff in selectedInventory {
//            print("\(stuff.name) checked ? : ", stuff.checked)
//        }
//        print("all checked ? : ", allStuffChecked)

        if let roomIndex = localRooms.firstIndex(where: { $0.id == selectedRoom.id }) {
            localRooms[roomIndex].checked = allStuffChecked
        }
    }

    // STUFF (FURNITURE) API CALLS

    func markStuffAsChecked(_ stuff: LocalInventory) async throws {
        guard let index = selectedInventory.firstIndex(where: { $0.id == stuff.id }) else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Stuff not found in inventory"])
        }

        selectedInventory[index].checked = true
        checkedStuffStatus[stuff.id] = true

        if let roomIndex = localRooms.firstIndex(where: { $0.id == selectedRoom?.id }),
           let stuffIndex = localRooms[roomIndex].inventory.firstIndex(where: { $0.id == stuff.id }) {
            localRooms[roomIndex].inventory[stuffIndex].checked = true
        }

        await MainActor.run {
            self.selectedInventory[index].checked = true
            self.updateRoomCheckedStatus()
        }
    }

    func fetchStuff(_ room: LocalRoom) async {
        guard let roomIndex = localRooms.firstIndex(where: { $0.id == room.id }) else {
            errorMessage = "Room not found in localRooms"
            return
        }

        guard localRooms[roomIndex].inventory.isEmpty else {
//            print("Inventory for room \(room.name) is already populated, skipping fetch.")
            return
        }

        guard let url = URL(string: "\(baseURL)/owner/properties/\(property.id)/rooms/\(room.id)/furnitures/") else {
            errorMessage = "Invalid URL"
            return
        }
        guard let token = await getToken() else {
            errorMessage = "Failed to retrieve token"
            return
        }

        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")

        do {
            let (data, response) = try await URLSession.shared.data(for: urlRequest)

            guard let httpResponse = response as? HTTPURLResponse else {
                errorMessage = "Invalid response"
                return
            }

            guard (200...299).contains(httpResponse.statusCode) else {
                errorMessage = "Error: Status code \(httpResponse.statusCode)"
                if let responseBody = String(data: data, encoding: .utf8) {
                    print("Response Body: \(responseBody)")
                }
                return
            }

            let decoder = JSONDecoder()
            let furnitures = try decoder.decode([FurnitureResponse].self, from: data)

            if let index = localRooms.firstIndex(where: { $0.id == room.id }) {
                localRooms[index].inventory = furnitures.map { furniture in
                    LocalInventory(id: furniture.id, propertyId: furniture.propertyId,
                                  roomId: furniture.roomId, name: furniture.name, quantity: furniture.quantity)
                }
                selectedInventory = localRooms[index].inventory

                selectedRoom = localRooms[index]
            }
//            print("fetch stuff: ")
//            dump(localRooms)
        } catch {
            errorMessage = "Error fetching furnitures: \(error.localizedDescription)"
        }
    }

    func addStuff(name: String, quantity: Int, to room: LocalRoom) async throws {
        guard let url = URL(string: "\(baseURL)/owner/properties/\(property.id)/rooms/\(room.id)/furnitures/") else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid URL"])
        }

        guard let token = await getToken() else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to retrieve token"])
        }

        let body: [String: Any] = [
            "name": name,
            "quantity": quantity
        ]

        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "POST"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")

        let jsonData = try JSONSerialization.data(withJSONObject: body)
        urlRequest.httpBody = jsonData

        let (_, response) = try await URLSession.shared.data(for: urlRequest)

        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server"])
        }

        guard (200...299).contains(httpResponse.statusCode) else {
            if httpResponse.statusCode == 400 {
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid furniture data"])
            } else if httpResponse.statusCode == 403 {
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property not yours"])
            } else if httpResponse.statusCode == 404 {
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "Room not found"])
            } else {
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode)"])
            }
        }

        await fetchStuff(room)
    }

    func deleteStuff(_ stuff: LocalInventory, from room: LocalRoom) async {
        isLoading = true
        defer { isLoading = false }

        guard let url = URL(string: "\(baseURL)/owner/properties/\(property.id)/rooms/\(room.id)/furnitures/\(stuff.id)/") else {
            errorMessage = "Invalid URL"
            return
        }

        guard let token = await getToken() else {
            errorMessage = "Failed to retrieve token"
            return
        }

        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "DELETE"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")

        do {
            let (_, response) = try await URLSession.shared.data(for: urlRequest)

            guard let httpResponse = response as? HTTPURLResponse else {
                errorMessage = "Invalid response"
                return
            }

            guard (200...299).contains(httpResponse.statusCode) else {
                errorMessage = "Error: Status code \(httpResponse.statusCode)"
                return
            }

            await fetchStuff(room)
        } catch {
            errorMessage = "Error deleting furniture: \(error.localizedDescription)"
        }
    }

    func selectStuff(_ stuff: LocalInventory) {
//        print("stuff selected: \(stuff)")
        selectedStuff = stuff
        selectedImages = stuff.images
        comment = stuff.comment
        selectedStatus = stuff.status
    }

    func sendStuffReport() async throws {
        guard let url = URL(string: "\(baseURL)/owner/properties/\(property.id)/inventory-reports/summarize/") else {
            throw URLError(.badURL)
        }
        guard let token = await getToken() else {
            throw URLError(.userAuthenticationRequired)
        }
        let base64Images = convertUIImagesToBase64(selectedImages)

        guard let stuffID = selectedStuff?.id else {
            throw URLError(.badServerResponse)
        }

        let body = SummarizeRequest(
            id: stuffID,
            pictures: base64Images,
            type: "furniture"
        )

        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")

        let encoder = JSONEncoder()
        request.httpBody = try encoder.encode(body)

        let (data, response) = try await URLSession.shared.data(for: request)

        guard let httpResponse = response as? HTTPURLResponse, (200...299).contains(httpResponse.statusCode) else {
            _ = String(data: data, encoding: .utf8) ?? "No response body"
            throw URLError(.badServerResponse)
        }

        let decoder = JSONDecoder()
        let summarizeResponse = try decoder.decode(SummarizeResponse.self, from: data)

        let stateMapping: [String: String] = [
            "not_set": "Select your equipment status",
            "broken": "Broken",
            "bad": "Bad",
            "good": "Good",
            "new": "New"
        ]
        let uiStatus = stateMapping[summarizeResponse.state] ?? "Select your equipment status"

        if let index = selectedInventory.firstIndex(where: { $0.id == stuffID }) {
            selectedInventory[index].images = selectedImages
            selectedInventory[index].status = uiStatus
            selectedInventory[index].comment = summarizeResponse.note
        }

        if let roomIndex = localRooms.firstIndex(where: { $0.id == selectedRoom?.id }),
           let stuffIndex = localRooms[roomIndex].inventory.firstIndex(where: { $0.id == stuffID }) {
            localRooms[roomIndex].inventory[stuffIndex].images = selectedImages
            localRooms[roomIndex].inventory[stuffIndex].status = uiStatus
            localRooms[roomIndex].inventory[stuffIndex].comment = summarizeResponse.note
        }

        await MainActor.run {
            self.comment = summarizeResponse.note
            self.selectedStatus = uiStatus
        }

        print("Report sent successfully: \(summarizeResponse)")
    }

    func finalizeInventory() async throws {
        guard areAllRoomsCompleted() else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Not all rooms and stuff are checked"])
        }

        let roomsData = localRooms.map { room in
            RoomStateRequest(
                id: room.id,
                cleanliness: "clean",
                state: "good",
                note: "Inventory completed",
                pictures: room.inventory.flatMap { $0.images.map { convertUIImageToBase64($0) } }, // Replace all datas by real datas when it will be implemented. Ask to the team.
                furnitures: room.inventory.map { stuff in
                    FurnitureStateRequest(
                        id: stuff.id,
                        cleanliness: "clean", // add cleanliness to local
                        note: stuff.comment,
                        pictures: stuff.images.map { convertUIImageToBase64($0) },
                        state: stuff.status.lowercased()
                    )
                }
            )
        }

        let requestBody = InventoryReportRequest(
            type: "start",
            rooms: roomsData
        )

        guard let url = URL(string: "\(baseURL)/owner/properties/\(property.id)/inventory-reports/") else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid URL"])
        }

        guard let token = await getToken() else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to retrieve token"])
        }

        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "POST"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")

        let encoder = JSONEncoder()
        urlRequest.httpBody = try encoder.encode(requestBody)

        let (data, response) = try await URLSession.shared.data(for: urlRequest)

        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server"])
        }

        guard (200...299).contains(httpResponse.statusCode) else {
            if let responseBody = String(data: data, encoding: .utf8) {
                print("Response Body: \(responseBody)")
            }
            throw NSError(domain: "", code: httpResponse.statusCode,
                          userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode)"])
        }

//        let decoder = JSONDecoder()
//        let responseData = try decoder.decode(InventoryReportResponse.self, from: data)
//        print("Inventory finalized successfully: \(responseData)")
        print("report inventory successfully created")
    }
}
