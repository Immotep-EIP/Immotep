//
//  FurnitureManager.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 02/03/2025.
//

import Foundation

@MainActor
class FurnitureManager {
    private weak var viewModel: InventoryViewModel?

    init(viewModel: InventoryViewModel) {
        self.viewModel = viewModel
    }

    func markStuffAsChecked(_ stuff: LocalInventory) async throws {
        guard let viewModel = viewModel else { return }
        guard let index = viewModel.selectedInventory.firstIndex(where: { $0.id == stuff.id }) else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Stuff not found in inventory"])
        }

        viewModel.selectedInventory[index].checked = true
        viewModel.checkedStuffStatus[stuff.id] = true

        if let roomIndex = viewModel.localRooms.firstIndex(where: { $0.id == viewModel.selectedRoom?.id }),
           let stuffIndex = viewModel.localRooms[roomIndex].inventory.firstIndex(where: { $0.id == stuff.id }) {
            viewModel.localRooms[roomIndex].inventory[stuffIndex].checked = true
        }

        viewModel.selectedInventory[index].checked = true
//        updateRoomCheckedStatus()
    }

    func fetchStuff(_ room: LocalRoom) async {
        guard let viewModel = viewModel else { return }
        guard let roomIndex = viewModel.localRooms.firstIndex(where: { $0.id == room.id }) else {
            viewModel.errorMessage = "Room not found in localRooms"
            return
        }

        guard let url = URL(string: "\(baseURL)/owner/properties/\(viewModel.property.id)/rooms/\(room.id)/furnitures/") else {
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
            let furnitures = try decoder.decode([FurnitureResponse].self, from: data)

            var updatedInventory: [LocalInventory] = []
            let currentInventory = viewModel.localRooms[roomIndex].inventory

            for furniture in furnitures {
                if let existingIndex = currentInventory.firstIndex(where: { $0.id == furniture.id }) {
                    let existingStuff = currentInventory[existingIndex]
                    updatedInventory.append(LocalInventory(
                        id: furniture.id,
                        propertyId: furniture.propertyId,
                        roomId: furniture.roomId,
                        name: furniture.name,
                        quantity: furniture.quantity,
                        checked: existingStuff.checked,
                        images: existingStuff.images,
                        status: existingStuff.status,
                        comment: existingStuff.comment
                    ))
                } else {
                    updatedInventory.append(LocalInventory(
                        id: furniture.id,
                        propertyId: furniture.propertyId,
                        roomId: furniture.roomId,
                        name: furniture.name,
                        quantity: furniture.quantity
                    ))
                }
            }

            viewModel.localRooms[roomIndex].inventory = updatedInventory
            viewModel.selectedInventory = viewModel.localRooms[roomIndex].inventory
            if viewModel.selectedRoom?.id == room.id {
                viewModel.selectedRoom = viewModel.localRooms[roomIndex]
            }
        } catch {
            viewModel.errorMessage = "Error fetching furnitures: \(error.localizedDescription)"
        }
    }

    func addStuff(name: String, quantity: Int, to room: LocalRoom) async throws {
        guard let viewModel = viewModel else { return }
        guard let url = URL(string: "\(baseURL)/owner/properties/\(viewModel.property.id)/rooms/\(room.id)/furnitures/") else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid URL"])
        }

        guard let token = await viewModel.getToken() else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to retrieve token"])
        }

        let body: [String: Any] = ["name": name, "quantity": quantity]
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
        guard let viewModel = viewModel else { return }
        viewModel.isLoading = true
        defer { viewModel.isLoading = false }

        guard let url = URL(string: "\(baseURL)/owner/properties/\(viewModel.property.id)/rooms/\(room.id)/furnitures/\(stuff.id)/") else {
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

            await fetchStuff(room)
        } catch {
            viewModel.errorMessage = "Error deleting furniture: \(error.localizedDescription)"
        }
    }

    func selectStuff(_ stuff: LocalInventory) {
        guard let viewModel = viewModel else { return }
        if let roomIndex = viewModel.localRooms.firstIndex(where: { $0.id == viewModel.selectedRoom?.id }),
           let stuffIndex = viewModel.localRooms[roomIndex].inventory.firstIndex(where: { $0.id == stuff.id }) {
            viewModel.selectedStuff = viewModel.localRooms[roomIndex].inventory[stuffIndex]
            viewModel.selectedImages = viewModel.selectedStuff!.images
            viewModel.comment = viewModel.selectedStuff!.comment
            viewModel.selectedStatus = viewModel.selectedStuff!.status
        } else {
            viewModel.selectedStuff = stuff
            viewModel.selectedImages = stuff.images
            viewModel.comment = stuff.comment
            viewModel.selectedStatus = stuff.status
        }
    }

    private func updateRoomCheckedStatus() {
        guard let viewModel = viewModel else { return }
        guard let selectedRoom = viewModel.selectedRoom else { return }
        let allStuffChecked = viewModel.selectedInventory.allSatisfy { $0.checked }
        if let roomIndex = viewModel.localRooms.firstIndex(where: { $0.id == selectedRoom.id }) {
            viewModel.localRooms[roomIndex].checked = allStuffChecked
        }
    }
}
