//
//  PropertyViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import SwiftUI
import Foundation

@MainActor
class PropertyViewModel: ObservableObject {
    @Published var properties: [Property] = []
    @Published var damages: [DamageResponse] = []
    @Published var isFetchingDamages = false
    @Published var damagesError: String?
    @Published var rooms: [PropertyRoomsTenant] = []
    @Published var activeLeaseId: String?
    @Published var isFetchingDocuments: Bool = false
    
    private let ownerViewModel = OwnerPropertyViewModel()
    public let tenantViewModel = TenantPropertyViewModel()
    private let loginViewModel: LoginViewModel
    
    init(loginViewModel: LoginViewModel) {
        self.loginViewModel = loginViewModel
        self.ownerViewModel.propertyViewModel = self
        self.tenantViewModel.propertyViewModel = self
    }
    
    @AppStorage("userRole") var storedUserRole: String?
    
    func createProperty(request: Property, token: String) async throws -> String {
        guard storedUserRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can create properties.".localized()])
        }
        return try await ownerViewModel.createProperty(request: request, token: token)
    }
    
    func updatePropertyPicture(token: String, propertyPicture: UIImage, propertyID: String) async throws -> String {
        guard storedUserRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can update property pictures.".localized()])
        }
        return try await ownerViewModel.updatePropertyPicture(token: token, propertyPicture: propertyPicture, propertyID: propertyID)
    }
    
    func fetchProperties() async throws {
        guard let userRole = storedUserRole else {
            throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "User role not found.".localized()])
        }
        
//        let startTime = Date()
        
        if userRole == "owner" {
            try await ownerViewModel.fetchProperties()
            await MainActor.run {
                self.properties = ownerViewModel.properties
                self.objectWillChange.send()
//                print("UI update took: \(Date().timeIntervalSince(startTime)) seconds")
            }
        } else if userRole == "tenant" {
            let property = try await tenantViewModel.fetchTenantProperty()
            await MainActor.run {
                self.properties = [property]
                self.rooms = tenantViewModel.rooms
                self.damages = tenantViewModel.damages
                self.activeLeaseId = tenantViewModel.activeLeaseId
                self.objectWillChange.send()
//                print("UI update took: \(Date().timeIntervalSince(startTime)) seconds")
            }
        } else {
            throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid user role.".localized()])
        }
    }
    
    func fetchPropertiesPicture(propertyId: String) async throws -> UIImage? {
        return try await ownerViewModel.fetchPropertiesPicture(propertyId: propertyId)
    }
    
    func updateProperty(request: Property, token: String) async throws -> String {
        guard storedUserRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can update properties.".localized()])
        }
        let result = try await ownerViewModel.updateProperty(request: request, token: token)
        try await fetchProperties()
        return result
    }
    
    func deleteProperty(propertyId: String) async throws {
        guard storedUserRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can delete properties.".localized()])
        }
        try await ownerViewModel.deleteProperty(propertyId: propertyId)
        try await fetchProperties()
    }
    
    func fetchPropertyDocuments(propertyId: String, forceRefresh: Bool = false) async throws {
        if !forceRefresh, let property = properties.first(where: { $0.id == propertyId }), !property.documents.isEmpty {
            return
        }
        
        await MainActor.run {
            self.isFetchingDocuments = true
        }
        defer {
            Task { @MainActor in
                self.isFetchingDocuments = false
            }
        }
        
        let documents: [PropertyDocument]
        if storedUserRole == "tenant" {
            if let leaseId = try await fetchActiveLeaseIdForProperty(propertyId: propertyId, token: try await TokenStorage.getValidAccessToken()) {
                documents = try await tenantViewModel.fetchTenantPropertyDocuments(leaseId: leaseId, propertyId: propertyId)
            } else {
                documents = []
            }
        } else {
            documents = try await ownerViewModel.fetchPropertyDocuments(propertyId: propertyId)
        }
        
        await MainActor.run {
            if let index = self.properties.firstIndex(where: { $0.id == propertyId }) {
                var updatedProperty = self.properties[index]
                updatedProperty.documents = documents
                self.properties[index] = updatedProperty
                self.objectWillChange.send()
            }
        }
    }
    
    func uploadDocument(propertyId: String, fileName: String, base64Data: String) async throws -> String {
        guard let userRole = storedUserRole else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "User role not defined.".localized()])
        }
        
        let documentId: String
        if userRole == "tenant" {
            guard let leaseId = try await fetchActiveLeaseIdForProperty(propertyId: propertyId, token: try await TokenStorage.getValidAccessToken()) else {
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "No active lease found.".localized()])
            }
            documentId = try await tenantViewModel.uploadTenantDocument(leaseId: leaseId, propertyId: propertyId, fileName: fileName, base64Data: base64Data)
        } else {
            documentId = try await ownerViewModel.uploadOwnerDocument(propertyId: propertyId, fileName: fileName, base64Data: base64Data)
        }
        
        let dateFormatter = DateFormatter()
        dateFormatter.dateFormat = "yyyy-MM-dd"
        let outputFormatter = DateFormatter()
        outputFormatter.dateFormat = "dd-MM-yyyy"
        
        let components = fileName.split(separator: "_")
        var title = fileName
        if let dateString = components.last?.prefix(10),
           dateFormatter.date(from: String(dateString)) != nil {
            title = outputFormatter.string(from: dateFormatter.date(from: String(dateString))!)
        }
        
        let newDocument = PropertyDocument(
            id: documentId,
            title: title,
            fileName: fileName,
            data: base64Data
        )
        
        await MainActor.run {
            if let index = self.properties.firstIndex(where: { $0.id == propertyId }) {
                var updatedProperty = self.properties[index]
                updatedProperty.documents.append(newDocument)
                self.properties[index] = updatedProperty
                self.objectWillChange.send()
            }
        }
        return documentId
    }
    
    func cancelInvite(propertyId: String, token: String) async throws {
        guard storedUserRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can cancel invites.".localized()])
        }
        try await ownerViewModel.cancelInvite(propertyId: propertyId, token: token)
        try await fetchProperties()
    }
    
    func endLease(propertyId: String, leaseId: String, token: String) async throws {
        guard storedUserRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can end leases.".localized()])
        }
        try await ownerViewModel.endLease(propertyId: propertyId, leaseId: leaseId, token: token)
        try await fetchProperties()
    }
    
    func fetchActiveLease(propertyId: String, token: String) async throws -> String? {
        if storedUserRole == "tenant" {
            return try await tenantViewModel.fetchActiveLeaseIdForProperty(propertyId: propertyId, token: token)
        } else {
            return try await ownerViewModel.fetchActiveLease(propertyId: propertyId, token: token)
        }
    }
    
    func fetchPropertyDamages(propertyId: String, fixed: Bool? = nil) async throws {
        guard let userRole = storedUserRole else {
            throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "User role not found.".localized()])
        }
        
//        let startTime = Date()
        
        let damages: [DamageResponse]
        if userRole == "tenant" {
            guard let leaseId = try await fetchActiveLeaseIdForProperty(propertyId: propertyId, token: try await TokenStorage.getValidAccessToken()) else {
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "No active lease found.".localized()])
            }
            damages = try await tenantViewModel.fetchTenantDamages(leaseId: leaseId, fixed: fixed)
        } else {
            damages = try await ownerViewModel.fetchPropertyDamages(propertyId: propertyId, fixed: fixed)
        }
        
        await MainActor.run {
            if let index = self.properties.firstIndex(where: { $0.id == propertyId }) {
                var updatedProperty = self.properties[index]
                updatedProperty.damages = damages
                self.properties[index] = updatedProperty
            }
            self.damages = damages
            self.isFetchingDamages = false
            self.objectWillChange.send()
//            print("UI update took: \(Date().timeIntervalSince(startTime)) seconds")
        }
    }
    
    func fetchActiveLeaseIdForProperty(propertyId: String, token: String) async throws -> String? {
        if let leaseId = activeLeaseId {
            return leaseId
        }
        
        guard let userRole = storedUserRole else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "User role not defined.".localized()])
        }
        
        let leaseId: String?
        if userRole == "tenant" {
            leaseId = try await tenantViewModel.fetchActiveLeaseIdForProperty(propertyId: propertyId, token: token)
        } else {
            leaseId = try await ownerViewModel.fetchActiveLease(propertyId: propertyId, token: token)
        }
        
        await MainActor.run {
            self.activeLeaseId = leaseId
            self.objectWillChange.send()
        }
        return leaseId
    }
    
    func fetchPropertyRooms(propertyId: String, token: String) async throws -> [PropertyRoomsTenant] {
        if !rooms.isEmpty {
            return rooms
        }
        let fetchedRooms = try await tenantViewModel.fetchPropertyRooms(token: token)
        await MainActor.run {
            self.rooms = fetchedRooms
            self.objectWillChange.send()
        }
        return fetchedRooms
    }
    
    func createDamage(propertyId: String, leaseId: String, damage: DamageRequest, token: String) async throws -> String {
        guard storedUserRole == "tenant" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only tenants can report damages.".localized()])
        }
        return try await tenantViewModel.createDamage(propertyId: propertyId, leaseId: leaseId, damage: damage, token: token)
    }
    
    func fetchLastInventoryReport(propertyId: String, leaseId: String) async throws -> InventoryReportResponse? {
        guard storedUserRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can fetch inventory reports.".localized()])
        }
        return try await ownerViewModel.fetchLastInventoryReport(propertyId: propertyId, leaseId: leaseId)
    }
    
    func fetchDamageByID(propertyId: String, damageId: String, token: String) async throws -> DamageResponse {
        guard let userRole = storedUserRole else {
            throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "User role not found.".localized()])
        }
        
        if userRole == "tenant" {
            return try await tenantViewModel.fetchDamageByID(damageId: damageId, token: token)
        } else {
            return try await ownerViewModel.fetchDamageByID(propertyId: propertyId, damageId: damageId, token: token)
        }
    }
    
    func updateDamageStatus(propertyId: String, damageId: String, fixPlannedAt: String?, read: Bool, token: String) async throws {
        guard storedUserRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can update damage status.".localized()])
        }
        try await ownerViewModel.updateDamageStatus(propertyId: propertyId, damageId: damageId, fixPlannedAt: fixPlannedAt, read: read, token: token)
    }
    
    func fixDamage(propertyId: String, damageId: String, token: String) async throws {
        guard let userRole = storedUserRole else {
            throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "User role not found.".localized()])
        }
        
        if userRole == "tenant" {
            try await tenantViewModel.fixDamage(damageId: damageId, token: token)
        } else {
            try await ownerViewModel.fixDamage(propertyId: propertyId, damageId: damageId, token: token)
        }
    }
    
    func uploadOwnerDocument(propertyId: String, fileName: String, base64Data: String) async throws -> String {
        guard storedUserRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can upload owner documents.".localized()])
        }
        return try await ownerViewModel.uploadOwnerDocument(propertyId: propertyId, fileName: fileName, base64Data: base64Data)
    }
    
    func uploadTenantDocument(leaseId: String, propertyId: String, fileName: String, base64Data: String) async throws -> String {
        guard storedUserRole == "tenant" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only tenants can upload tenant documents.".localized()])
        }
        return try await tenantViewModel.uploadTenantDocument(leaseId: leaseId, propertyId: propertyId, fileName: fileName, base64Data: base64Data)
    }
    
    func deleteDocument(propertyId: String, documentId: String) async throws {
        guard storedUserRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can delete documents.".localized()])
        }
        try await ownerViewModel.deleteDocument(propertyId: propertyId, documentId: documentId)
        try await fetchPropertyDocuments(propertyId: propertyId, forceRefresh: true)
    }
    
    func fetchPropertyById(_ propertyId: String, leaseId: String? = nil) async throws -> Property? {
        if let existingProperty = properties.first(where: { $0.id == propertyId }) {
            return existingProperty
        }
        
        guard let userRole = storedUserRole else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "User role not defined.".localized()])
        }
        
        let token = try await TokenStorage.getValidAccessToken()
        var urlString = "\(APIConfig.baseURL)/\(userRole == "tenant" ? "tenant/leases/current/property" : "owner/properties/\(propertyId)")/"
        if let leaseId = leaseId {
            urlString += "?lease_id=\(leaseId)"
        }
        guard let url = URL(string: urlString) else {
            throw NSError(domain: "", code: 400, userInfo: [NSLocalizedDescriptionKey: "Invalid URL".localized()])
        }
        
        var urlRequest = URLRequest(url: url)
        urlRequest.httpMethod = "GET"
        urlRequest.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        urlRequest.setValue("application/json", forHTTPHeaderField: "Accept")
        
        let (data, response) = try await URLSession.shared.data(for: urlRequest)
        
        guard let httpResponse = response as? HTTPURLResponse, (200...299).contains(httpResponse.statusCode) else {
            let errorBody = String(data: data, encoding: .utf8) ?? "No error details"
            throw NSError(domain: "", code: (response as? HTTPURLResponse)?.statusCode ?? 0, userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \((response as? HTTPURLResponse)?.statusCode ?? 0) - \(errorBody)".localized()])
        }
        
        let propertyResponse = try await Task.detached(priority: .background) {
            let decoder = JSONDecoder()
            let dateFormatter = DateFormatter()
            dateFormatter.dateFormat = "yyyy-MM-dd'T'HH:mm:ss.SSSZ"
            let fallbackFormatter = DateFormatter()
            fallbackFormatter.dateFormat = "yyyy-MM-dd'T'HH:mm:ssZ"
            decoder.dateDecodingStrategy = .custom { decoder in
                let container = try decoder.singleValueContainer()
                let dateString = try container.decode(String.self)
                if let date = dateFormatter.date(from: dateString) {
                    return date
                } else if let date = fallbackFormatter.date(from: dateString) {
                    return date
                }
                throw DecodingError.dataCorruptedError(in: container, debugDescription: "Invalid date format: \(dateString)")
            }
            return try decoder.decode(PropertyResponse.self, from: data)
        }.value
        
        let photo = try await fetchPropertiesPicture(propertyId: propertyId)
        
        let property = Property(
            id: propertyResponse.id,
            ownerID: propertyResponse.ownerId,
            name: propertyResponse.name,
            address: propertyResponse.address,
            city: propertyResponse.city,
            postalCode: propertyResponse.postalCode,
            country: propertyResponse.country,
            photo: photo,
            monthlyRent: propertyResponse.rentalPricePerMonth,
            deposit: propertyResponse.depositPrice,
            surface: propertyResponse.areaSqm,
            isAvailable: propertyResponse.isAvailable == "invite sent" ? "pending" : propertyResponse.isAvailable,
            tenantName: propertyResponse.lease?.tenantName,
            leaseId: propertyResponse.lease?.id,
            leaseStartDate: propertyResponse.lease?.startDate,
            leaseEndDate: propertyResponse.lease?.endDate,
            documents: [],
            createdAt: propertyResponse.createdAt,
            rooms: [],
            damages: []
        )
        
        await MainActor.run {
            if let index = self.properties.firstIndex(where: { $0.id == propertyId }) {
                self.properties[index] = property
            } else {
                self.properties.append(property)
            }
            self.objectWillChange.send()
        }
        return property
    }
}

struct PropertyID: Decodable {
    let id: String
}

struct PropertyImageBase64: Decodable {
    let data: String
}
