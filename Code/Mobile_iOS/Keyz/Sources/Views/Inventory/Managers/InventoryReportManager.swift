//
//  InventoryReportManager.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 02/03/2025.
//

import Foundation
import SwiftUI

@MainActor
class InventoryReportManager {
    private weak var viewModel: InventoryViewModel?

    init(viewModel: InventoryViewModel) {
        self.viewModel = viewModel
    }

    func sendStuffReport() async throws {
        guard let viewModel = viewModel else {
            throw URLError(.cannotFindHost)
        }
        
        guard let leaseId = viewModel.property.leaseId, !leaseId.isEmpty else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "No active lease found for property \(viewModel.property.id)"])
        }
        
        guard let token = await viewModel.getToken() else {
            throw URLError(.userAuthenticationRequired)
        }
        
        guard let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(viewModel.property.id)/leases/current/inventory-reports/summarize/") else {
            throw URLError(.badURL)
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

        guard let httpResponse = response as? HTTPURLResponse else {
            throw URLError(.badServerResponse)
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No response body"
            print("API Error: Status code \(httpResponse.statusCode) - \(errorBody)")
            if httpResponse.statusCode == 404 {
                throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Property or lease not found"])
            } else if httpResponse.statusCode == 403 {
                throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Property not yours"])
            } else if httpResponse.statusCode == 400 {
                throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Invalid request: \(errorBody)"])
            } else {
                throw URLError(.badServerResponse)
            }
        }

        let decoder = JSONDecoder()
        let summarizeResponse = try decoder.decode(SummarizeResponse.self, from: data)

        let validStates = ["not_set", "broken", "needsRepair", "bad", "medium", "good", "new"]
        let apiStatus = validStates.contains(summarizeResponse.state) ? summarizeResponse.state : "not_set"

        if let index = viewModel.selectedInventory.firstIndex(where: { $0.id == stuffID }) {
            viewModel.selectedInventory[index].images = viewModel.selectedImages
            viewModel.selectedInventory[index].status = apiStatus
            viewModel.selectedInventory[index].comment = summarizeResponse.note
        }

        if let roomIndex = viewModel.localRooms.firstIndex(where: { $0.id == viewModel.selectedRoom?.id }),
           let stuffIndex = viewModel.localRooms[roomIndex].inventory.firstIndex(where: { $0.id == stuffID }) {
            viewModel.localRooms[roomIndex].inventory[stuffIndex].images = viewModel.selectedImages
            viewModel.localRooms[roomIndex].inventory[stuffIndex].status = apiStatus
            viewModel.localRooms[roomIndex].inventory[stuffIndex].comment = summarizeResponse.note
        }

        viewModel.comment = summarizeResponse.note
        viewModel.selectedStatus = apiStatus
    }

    func finalizeInventory() async throws {
        guard let viewModel = viewModel else { return }
        guard viewModel.areAllRoomsCompleted() else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Not all rooms and stuff are checked"])
        }
        guard let leaseId = viewModel.property.leaseId, !leaseId.isEmpty else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "No active lease found for property \(viewModel.property.id)"])
        }

        let placeholderImage = UIImage(named: "DefaultImageProperty") ?? createBlankImage()
        let placeholderBase64 = convertUIImageToBase64(placeholderImage)

        let roomsData = viewModel.localRooms.map { room in
            let roomPictures = room.images.isEmpty ? [placeholderBase64] : room.images.map { convertUIImageToBase64($0) }
            let roomState = room.status.lowercased() == "select room status" ? "good" : room.status.lowercased()
            let roomNote = room.comment.isEmpty ? "No comment provided" : room.comment

            return RoomStateRequest(
                id: room.id,
                cleanliness: "clean",
                state: roomState,
                note: roomNote,
                pictures: roomPictures,
                furnitures: room.inventory.map { stuff in
                    let validStates = ["broken", "bad", "good", "new"]
                    let stuffState = validStates.contains(stuff.status.lowercased()) ? stuff.status.lowercased() : "good"
                    let stuffNote = stuff.comment.isEmpty ? "No comment provided" : stuff.comment
                    let stuffPictures = stuff.images.map { convertUIImageToBase64($0) }
                    let finalStuffPictures = stuffPictures.isEmpty ? [placeholderBase64] : stuffPictures

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

        guard let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(viewModel.property.id)/leases/current/inventory-reports/") else {
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
            viewModel.onDocumentsRefreshNeeded?()
        } catch {
            viewModel.completionMessage = viewModel.isEntryInventory
            ? "Failed to finalize entry inventory: \(error.localizedDescription)"
            : "Failed to finalize exit inventory: \(error.localizedDescription)"
            throw error
        }
    }

    func createBlankImage() -> UIImage {
        let size = CGSize(width: 1, height: 1)
        UIGraphicsBeginImageContext(size)
        let context = UIGraphicsGetCurrentContext()
        context?.setFillColor(UIColor.white.cgColor)
        context?.fill(CGRect(origin: .zero, size: size))
        let image = UIGraphicsGetImageFromCurrentImageContext()
        UIGraphicsEndImageContext()
        return image ?? UIImage()
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
        guard let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(viewModel.property.id)/leases/current/inventory-reports/compare/\(oldReportId)/") else {
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

        let validStates = ["not_set", "broken", "needsRepair", "bad", "medium", "good", "new"]
        let apiStatus = validStates.contains(summarizeResponse.state) ? summarizeResponse.state : "not_set"

        if let index = viewModel.selectedInventory.firstIndex(where: { $0.id == stuffID }) {
            viewModel.selectedInventory[index].images = viewModel.selectedImages
            viewModel.selectedInventory[index].status = apiStatus
            viewModel.selectedInventory[index].comment = summarizeResponse.note
        }

        if let roomIndex = viewModel.localRooms.firstIndex(where: { $0.id == viewModel.selectedRoom?.id }),
           let stuffIndex = viewModel.localRooms[roomIndex].inventory.firstIndex(where: { $0.id == stuffID }) {
            viewModel.localRooms[roomIndex].inventory[stuffIndex].images = viewModel.selectedImages
            viewModel.localRooms[roomIndex].inventory[stuffIndex].status = apiStatus
            viewModel.localRooms[roomIndex].inventory[stuffIndex].comment = summarizeResponse.note
        }

        viewModel.comment = summarizeResponse.note
        viewModel.selectedStatus = apiStatus
    }
    
    func sendRoomReport() async throws {
        guard let viewModel = viewModel else {
            throw URLError(.cannotFindHost)
        }
        
        guard let leaseId = viewModel.property.leaseId, !leaseId.isEmpty else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "No active lease found for property \(viewModel.property.id)"])
        }
        
        guard let token = await viewModel.getToken() else {
            throw URLError(.userAuthenticationRequired)
        }
        
        guard let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(viewModel.property.id)/leases/current/inventory-reports/summarize/") else {
            throw URLError(.badURL)
        }
        
        let base64Images = convertUIImagesToBase64(viewModel.selectedImages)
        guard let roomId = viewModel.selectedRoom?.id else {
            throw URLError(.badServerResponse)
        }

        let body = SummarizeRequest(
            id: roomId,
            pictures: base64Images,
            type: "room"
        )

        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")

        let encoder = JSONEncoder()
        request.httpBody = try encoder.encode(body)

        let (data, response) = try await URLSession.shared.data(for: request)

        guard let httpResponse = response as? HTTPURLResponse else {
            throw URLError(.badServerResponse)
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No response body"
            print("API Error: Status code \(httpResponse.statusCode) - \(errorBody)")
            if httpResponse.statusCode == 404 {
                throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Property or lease not found"])
            } else if httpResponse.statusCode == 403 {
                throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Property not yours"])
            } else if httpResponse.statusCode == 400 {
                throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Invalid request: \(errorBody)"])
            } else {
                throw URLError(.badServerResponse)
            }
        }

        let decoder = JSONDecoder()
        let summarizeResponse = try decoder.decode(SummarizeResponse.self, from: data)

        let validStates = ["not_set", "broken", "needsRepair", "bad", "medium", "good", "new"]
        let apiStatus = validStates.contains(summarizeResponse.state) ? summarizeResponse.state : "not_set"

        if let roomIndex = viewModel.localRooms.firstIndex(where: { $0.id == roomId }) {
            viewModel.localRooms[roomIndex].images = viewModel.selectedImages
            viewModel.localRooms[roomIndex].status = apiStatus
            viewModel.localRooms[roomIndex].comment = summarizeResponse.note
            viewModel.selectedRoom = viewModel.localRooms[roomIndex]
        } else {
            print("Error: Room with ID \(roomId) not found in localRooms")
        }

        viewModel.comment = summarizeResponse.note
        viewModel.selectedStatus = apiStatus
    }

    func compareRoomReport(oldReportId: String) async throws {
        guard let viewModel = viewModel else {
            throw URLError(.cannotFindHost)
        }
        
        guard let url = URL(string: "\(APIConfig.baseURL)/owner/properties/\(viewModel.property.id)/leases/current/inventory-reports/compare/\(oldReportId)/") else {
            throw URLError(.badURL)
        }
        
        guard let token = await viewModel.getToken() else {
            throw URLError(.userAuthenticationRequired)
        }
        
        let base64Images = convertUIImagesToBase64(viewModel.selectedImages)
        guard let roomId = viewModel.selectedRoom?.id else {
            throw URLError(.badServerResponse)
        }

        let body = SummarizeRequest(
            id: roomId,
            pictures: base64Images,
            type: "room"
        )

        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")

        let encoder = JSONEncoder()
        request.httpBody = try encoder.encode(body)

        let (data, response) = try await URLSession.shared.data(for: request)

        guard let httpResponse = response as? HTTPURLResponse else {
            throw URLError(.badServerResponse)
        }
        
        guard (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No response body"
            print("API Error: Status code \(httpResponse.statusCode) - \(errorBody)")
            if httpResponse.statusCode == 404 {
                throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Property or lease not found"])
            } else if httpResponse.statusCode == 403 {
                throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Property not yours"])
            } else if httpResponse.statusCode == 400 {
                throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Invalid request: \(errorBody)"])
            } else {
                throw URLError(.badServerResponse)
            }
        }

        let decoder = JSONDecoder()
        let summarizeResponse = try decoder.decode(SummarizeResponse.self, from: data)

        let validStates = ["not_set", "broken", "needsRepair", "bad", "medium", "good", "new"]
        let apiStatus = validStates.contains(summarizeResponse.state) ? summarizeResponse.state : "not_set"

        if let roomIndex = viewModel.localRooms.firstIndex(where: { $0.id == roomId }) {
            viewModel.localRooms[roomIndex].images = viewModel.selectedImages
            viewModel.localRooms[roomIndex].status = apiStatus
            viewModel.localRooms[roomIndex].comment = summarizeResponse.note
        }

        viewModel.comment = summarizeResponse.note
        viewModel.selectedStatus = apiStatus
    }

    private func convertUIImagesToBase64(_ images: [UIImage]) -> [String] {
        return images.compactMap { image in
            guard let imageData = image.jpegData(compressionQuality: 0.8) else {
                print("Failed to convert UIImage to JPEG data")
                return nil
            }
            let base64String = imageData.base64EncodedString()
            let fullString = "data:image/jpeg;base64,\(base64String)"
            if let decodedData = Data(base64Encoded: base64String), let _ = UIImage(data: decodedData) {
                return fullString
            } else {
                print("Invalid Base64 string for image")
                return nil
            }
        }
    }

    private func convertUIImageToBase64(_ image: UIImage) -> String {
        guard let imageData = image.jpegData(compressionQuality: 0.8) else {
            print("Failed to convert UIImage to JPEG data")
            return ""
        }
        let base64String = imageData.base64EncodedString()
        return "data:image/jpeg;base64,\(base64String)"
    }
}
