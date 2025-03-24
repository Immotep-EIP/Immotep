//
//  InventoryReportManager.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 02/03/2025.
//

import Foundation

@MainActor
class InventoryReportManager {
    private weak var viewModel: InventoryViewModel?

    init(viewModel: InventoryViewModel) {
        self.viewModel = viewModel
    }

    func sendStuffReport() async throws {
        guard let viewModel = viewModel else { return }
        guard let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(viewModel.property.id)/inventory-reports/summarize/") else {
            throw URLError(.badURL)
        }
        guard let token = await viewModel.getToken() else {
            throw URLError(.userAuthenticationRequired)
        }
        let base64Images = convertUIImagesToBase64(viewModel.selectedImages)

        guard let stuffID = viewModel.selectedStuff?.id else {
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

        if let index = viewModel.selectedInventory.firstIndex(where: { $0.id == stuffID }) {
            viewModel.selectedInventory[index].images = viewModel.selectedImages
            viewModel.selectedInventory[index].status = uiStatus
            viewModel.selectedInventory[index].comment = summarizeResponse.note
        }

        if let roomIndex = viewModel.localRooms.firstIndex(where: { $0.id == viewModel.selectedRoom?.id }),
           let stuffIndex = viewModel.localRooms[roomIndex].inventory.firstIndex(where: { $0.id == stuffID }) {
            viewModel.localRooms[roomIndex].inventory[stuffIndex].images = viewModel.selectedImages
            viewModel.localRooms[roomIndex].inventory[stuffIndex].status = uiStatus
            viewModel.localRooms[roomIndex].inventory[stuffIndex].comment = summarizeResponse.note
        }

        viewModel.comment = summarizeResponse.note
        viewModel.selectedStatus = uiStatus
//        print("Report sent successfully: \(summarizeResponse)")
    }

    func finalizeInventory() async throws {
        guard let viewModel = viewModel else { return }
        guard viewModel.areAllRoomsCompleted() else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Not all rooms and stuff are checked"])
        }

        let roomsData = viewModel.localRooms.map { room in
            let roomPictures = room.inventory.flatMap { $0.images.map { convertUIImageToBase64($0) } }
            let finalRoomPictures = roomPictures.isEmpty ? [""] : roomPictures

            return RoomStateRequest(
                id: room.id,
                cleanliness: "clean",
                state: "good",
                note: "Inventory completed",
                pictures: finalRoomPictures,
                furnitures: room.inventory.map { stuff in
                    let validStates = ["broken", "bad", "good", "new"]
                    let stuffState = validStates.contains(stuff.status.lowercased()) ? stuff.status.lowercased() : "good"
                    let stuffNote = stuff.comment.isEmpty ? "No comment provided" : stuff.comment
                    let stuffPictures = stuff.images.map { convertUIImageToBase64($0) }
                    let finalStuffPictures = stuffPictures.isEmpty ? [""] : stuffPictures

                    return FurnitureStateRequest(
                        id: stuff.id,
                        cleanliness: "clean",
                        note: stuffNote,
                        pictures: finalStuffPictures,
                        state: stuffState
                    )
                }
            )
        }

        let requestBody = InventoryReportRequest(type: viewModel.isEntryInventory ? "start" : "end", rooms: roomsData)

        guard let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(viewModel.property.id)/inventory-reports/") else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid URL"])
        }

        guard let token = await viewModel.getToken() else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Failed to retrieve token"])
        }

        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "POST"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Content-Type")

        let encoder = JSONEncoder()
        urlRequest.httpBody = try encoder.encode(requestBody)

        do {
            let (_, response) = try await URLSession.shared.data(for: urlRequest)

            guard let httpResponse = response as? HTTPURLResponse else {
                throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid response from server"])
            }

            guard (200...299).contains(httpResponse.statusCode) else {
                throw NSError(domain: "", code: httpResponse.statusCode,
                              userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode)"])
            }
            viewModel.completionMessage = viewModel.isEntryInventory
            ? "Entry inventory finalized successfully!" : "Exit inventory finalized successfully!"
//            print("Report inventory successfully created")
        } catch {
            viewModel.completionMessage = viewModel.isEntryInventory
            ? "Failed to finalize entry inventory: \(error.localizedDescription)"
            : "Failed to finalize exit inventory: \(error.localizedDescription)"
            throw error
        }
    }

    func fetchLastInventoryReport() async {
        guard let viewModel = viewModel else { return }
        guard let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(viewModel.property.id)/inventory-reports/latest/") else {
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
                viewModel.errorMessage = "Invalid server response"
                return
            }

            _ = String(data: data, encoding: .utf8) ?? "Unable to decode response"

            guard (200...299).contains(httpResponse.statusCode) else {
                viewModel.errorMessage = "Failed to fetch last report: \(httpResponse.statusCode)"
                viewModel.lastReportId = nil
                return
            }

            if data.isEmpty {
                viewModel.lastReportId = nil
                return
            }

            let decoder = JSONDecoder()
            let report = try decoder.decode(LastInventoryReportResponse.self, from: data)
            viewModel.lastReportId = report.id
        } catch {
            viewModel.lastReportId = nil
        }
    }

    func compareStuffReport(oldReportId: String) async throws {
        guard let viewModel = viewModel else { return }
        guard let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(viewModel.property.id)/inventory-reports/compare/\(oldReportId)/") else {
            throw URLError(.badURL)
        }
        guard let token = await viewModel.getToken() else {
            throw URLError(.userAuthenticationRequired)
        }
        let base64Images = convertUIImagesToBase64(viewModel.selectedImages)

        guard let stuffID = viewModel.selectedStuff?.id else {
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
            throw URLError(.badServerResponse)
        }

        let decoder = JSONDecoder()
        let summarizeResponse = try decoder.decode(SummarizeResponse.self, from: data)

        let stateMapping: [String: String] = [
            "broken": "Broken",
            "bad": "Bad",
            "good": "Good",
            "new": "New",
            "needsRepair": "Needs Repair",
            "medium": "Medium"
        ]
        let uiStatus = stateMapping[summarizeResponse.state] ?? "Select your equipment status"

        if let index = viewModel.selectedInventory.firstIndex(where: { $0.id == stuffID }) {
            viewModel.selectedInventory[index].images = viewModel.selectedImages
            viewModel.selectedInventory[index].status = uiStatus
            viewModel.selectedInventory[index].comment = summarizeResponse.note
        }

        if let roomIndex = viewModel.localRooms.firstIndex(where: { $0.id == viewModel.selectedRoom?.id }),
           let stuffIndex = viewModel.localRooms[roomIndex].inventory.firstIndex(where: { $0.id == stuffID }) {
            viewModel.localRooms[roomIndex].inventory[stuffIndex].images = viewModel.selectedImages
            viewModel.localRooms[roomIndex].inventory[stuffIndex].status = uiStatus
            viewModel.localRooms[roomIndex].inventory[stuffIndex].comment = summarizeResponse.note
        }

        viewModel.comment = summarizeResponse.note
        viewModel.selectedStatus = uiStatus
//        print("Comparison report sent successfully: \(summarizeResponse)")
    }
}
