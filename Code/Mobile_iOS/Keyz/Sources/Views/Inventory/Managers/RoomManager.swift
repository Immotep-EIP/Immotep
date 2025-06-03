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
        guard let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(viewModel.property.id)/rooms/") else {
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

            var updatedLocalRooms: [LocalRoom] = []
            for room in viewModel.property.rooms {
                if let existingRoom = viewModel.localRooms.first(where: { $0.id == room.id }) {
                    updatedLocalRooms.append(LocalRoom(
                        id: room.id,
                        name: room.name,
                        checked: existingRoom.checked,
                        inventory: existingRoom.inventory
                    ))
                } else {
                    updatedLocalRooms.append(LocalRoom(
                        id: room.id,
                        name: room.name,
                        checked: false,
                        inventory: []
                    ))
                }
            }

            viewModel.localRooms = updatedLocalRooms

        } catch {
            viewModel.errorMessage = "Error fetching rooms: \(error.localizedDescription)"
        }
    }

    func addRoom(name: String, type: String) async throws {
        guard let viewModel = viewModel else { return }
        guard let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(viewModel.property.id)/rooms/") else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid URL"])
        }

        guard let token = await viewModel.getToken() else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to retrieve token"])
        }

        let validRoomTypes = [
            "dressing", "laundryroom", "bedroom", "playroom", "bathroom", "toilet",
            "livingroom", "diningroom", "kitchen", "hallway", "balcony", "cellar",
            "garage", "storage", "office", "other"
        ]
        guard validRoomTypes.contains(type) else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid room type: \(type)"])
        }

        let body: [String: Any] = [
            "name": name,
            "type": type
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
                // Decode error response for more details
                if let errorResponse = try? JSONDecoder().decode(ErrorResponse.self, from: data) {
                    throw NSError(domain: "", code: httpResponse.statusCode,
                                  userInfo: [NSLocalizedDescriptionKey: "API error: \(errorResponse.error)"])
                }
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode)"])
            }

            let decoder = JSONDecoder()
            let idResponse = try decoder.decode(IdResponse.self, from: data)

            await fetchRooms()
        } catch {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Error creating room: \(error.localizedDescription)"])
        }
    }

    func deleteRoom(_ room: LocalRoom) async {
        guard let viewModel = viewModel else { return }
        guard let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(viewModel.property.id)/rooms/\(room.id)/archive/") else {
            viewModel.errorMessage = "Invalid URL"
            return
        }

        guard let token = await viewModel.getToken() else {
            viewModel.errorMessage = "Failed to retrieve token"
            return
        }

        let body: [String: Any] = [
                "archive": true
            ]

        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "PUT"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")

        do {
            let jsonData = try JSONSerialization.data(withJSONObject: body)
            urlRequest.httpBody = jsonData
        } catch {
            print("error: \(error.localizedDescription)")
//            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to serialize request body: \(error.localizedDescription)".localized()])
        }

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
