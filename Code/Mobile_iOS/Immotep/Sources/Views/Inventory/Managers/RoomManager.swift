//
//  RoomManager.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 02/03/2025.
//


import Foundation

@MainActor
class RoomManager {
    private weak var viewModel: InventoryViewModel?

    init(viewModel: InventoryViewModel) {
        self.viewModel = viewModel
    }

    func fetchRooms() async {
        guard let viewModel = viewModel else { return }
        guard let url = URL(string: "\(baseURL)/owner/properties/\(viewModel.property.id)/rooms/") else {
            viewModel.errorMessage = "Invalid URL"
            return
        }

        guard let token = await viewModel.getToken() else {
            viewModel.errorMessage = "Failed to retrieve token"
            return
        }

        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")

        do {
            let (data, response) = try await URLSession.shared.data(for: urlRequest)

            guard let httpResponse = response as? HTTPURLResponse else {
                viewModel.errorMessage = "Invalid response"
                return
            }

            guard (200...299).contains(httpResponse.statusCode) else {
                viewModel.errorMessage = "Error: Status code \(httpResponse.statusCode)"
                return
            }

            let decoder = JSONDecoder()
            let roomsData = try decoder.decode([RoomResponse].self, from: data)

            viewModel.property.rooms = roomsData.map { roomResponse in
                PropertyRooms(
                    id: roomResponse.id,
                    name: roomResponse.name,
                    checked: false,
                    inventory: []
                )
            }

            if viewModel.localRooms.isEmpty {
                viewModel.localRooms = viewModel.property.rooms.map { room in
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
//            dump(viewModel.localRooms)
        } catch {
            viewModel.errorMessage = "Error fetching rooms: \(error.localizedDescription)"
        }
    }

    func addRoom(name: String) async throws {
        guard let viewModel = viewModel else { return }
        guard let url = URL(string: "\(baseURL)/owner/properties/\(viewModel.property.id)/rooms/") else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid URL"])
        }

        guard let token = await viewModel.getToken() else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to retrieve token"])
        }

        let body: [String: Any] = ["name": name]
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
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode)"])
            }

            let decoder = JSONDecoder()
            let newRoomResponse = try decoder.decode(RoomResponse.self, from: data)

            let newLocalRoom = LocalRoom(
                id: newRoomResponse.id,
                name: newRoomResponse.name,
                checked: false,
                inventory: []
            )
            viewModel.localRooms.append(newLocalRoom)
        } catch {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Error creating room: \(error.localizedDescription)"])
        }
    }

    func deleteRoom(_ room: LocalRoom) async {
        guard let viewModel = viewModel else { return }
        guard let url = URL(string: "\(baseURL)/owner/properties/\(viewModel.property.id)/rooms/\(room.id)/") else {
            viewModel.errorMessage = "Invalid URL"
            return
        }

        guard let token = await viewModel.getToken() else {
            viewModel.errorMessage = "Failed to retrieve token"
            return
        }

        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "DELETE"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")

        do {
            let (_, response) = try await URLSession.shared.data(for: urlRequest)

            guard let httpResponse = response as? HTTPURLResponse else {
                viewModel.errorMessage = "Invalid response"
                return
            }

            guard (200...299).contains(httpResponse.statusCode) else {
                viewModel.errorMessage = "Error: Status code \(httpResponse.statusCode)"
                return
            }

            await fetchRooms()
            viewModel.localRooms.removeAll { $0.id == room.id }
        } catch {
            viewModel.errorMessage = "Error deleting room: \(error.localizedDescription)"
        }
    }

    func selectRoom(_ room: LocalRoom) {
        guard let viewModel = viewModel else { return }
        viewModel.selectedRoom = room
        if let roomIndex = viewModel.localRooms.firstIndex(where: { $0.id == room.id }) {
            viewModel.selectedInventory = viewModel.localRooms[roomIndex].inventory
        }
    }

    func isRoomCompleted(_ room: LocalRoom) -> Bool {
        return room.inventory.allSatisfy { $0.checked }
    }

    func areAllRoomsCompleted() -> Bool {
        guard let viewModel = viewModel else { return false }
        return viewModel.localRooms.allSatisfy { $0.checked }
    }

    func markRoomAsChecked(_ room: LocalRoom) async {
        guard let viewModel = viewModel else { return }
        guard let index = viewModel.localRooms.firstIndex(where: { $0.id == room.id }) else { return }
        viewModel.localRooms[index].checked = true
    }

    func updateRoomCheckedStatus() {
        guard let viewModel = viewModel else { return }
        guard let selectedRoom = viewModel.selectedRoom else { return }
        let allStuffChecked = viewModel.selectedInventory.allSatisfy { $0.checked }
        if let roomIndex = viewModel.localRooms.firstIndex(where: { $0.id == selectedRoom.id }) {
            viewModel.localRooms[roomIndex].checked = allStuffChecked
        }
    }
}
