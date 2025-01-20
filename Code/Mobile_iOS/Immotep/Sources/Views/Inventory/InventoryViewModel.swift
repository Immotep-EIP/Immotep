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

    @Published var selectedRoom: PropertyRooms?
    @Published var selectedInventory: [RoomInventory] = []

    @Published var selectedStuff: RoomInventory?
    @Published var selectedImages: [UIImage] = []
    @Published var comment: String = ""
    @Published var selectedStatus: String = "Select your equipment status"

    @Published var roomToDelete: PropertyRooms?
    @Published var showDeleteConfirmation: Bool = false

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

    func fetchRooms() async {

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

        let jsonData = try JSONSerialization.data(withJSONObject: body)
        urlRequest.httpBody = jsonData

        let (_, response) = try await URLSession.shared.data(for: urlRequest)

        guard let httpResponse = response as? HTTPURLResponse else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server"])
        }

        guard (200...299).contains(httpResponse.statusCode) else {
            if httpResponse.statusCode == 400 {
                throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid room data"])
            } else if httpResponse.statusCode == 403 {
                throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Property not yours"])
            } else if httpResponse.statusCode == 404 {
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "Property not found"])
            } else {
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode)"])
            }
        }

        await fetchRooms()
    }

    func deleteRoom(_ room: PropertyRooms) async {
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
        } catch {
            errorMessage = "Error deleting room: \(error.localizedDescription)"
        }
    }

    func fetchStuff(_ room: PropertyRooms) async {

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

            if let index = property.rooms.firstIndex(where: { $0.id == room.id }) {
                property.rooms[index].inventory = furnitures.map { furniture in
                    RoomInventory(id: furniture.id, propertyId: furniture.propertyId,
                                  roomId: furniture.roomId, name: furniture.name, quantity: furniture.quantity)
                    }
                selectedInventory = property.rooms[index].inventory
                }
        } catch {
            errorMessage = "Error fetching furnitures: \(error.localizedDescription)"
        }
    }

    func addStuff(name: String, quantity: Int, to room: PropertyRooms) async throws {

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

    func deleteStuff(_ stuff: RoomInventory, from room: PropertyRooms) async {
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
}
