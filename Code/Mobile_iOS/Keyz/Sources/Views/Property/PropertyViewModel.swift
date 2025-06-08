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
    
    private let ownerViewModel = OwnerPropertyViewModel()
    public let tenantViewModel = TenantPropertyViewModel()
    private let loginViewModel: LoginViewModel
    
    init(loginViewModel: LoginViewModel) {
        self.loginViewModel = loginViewModel
    }
    
    func createProperty(request: Property, token: String) async throws -> String {
        guard loginViewModel.userRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can create properties.".localized()])
        }
        return try await ownerViewModel.createProperty(request: request, token: token)
    }
    
    func updatePropertyPicture(token: String, propertyPicture: UIImage, propertyID: String) async throws -> String {
        guard loginViewModel.userRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can update property pictures.".localized()])
        }
        return try await ownerViewModel.updatePropertyPicture(token: token, propertyPicture: propertyPicture, propertyID: propertyID)
    }
    
    func fetchProperties() async {
        if loginViewModel.userRole == "tenant" {
            do {
                let property = try await tenantViewModel.fetchTenantProperty()
                self.properties = [property]
            } catch {
                print("Error fetching tenant property: \(error.localizedDescription)")
            }
        } else {
            do {
                try await ownerViewModel.fetchProperties()
                self.properties = ownerViewModel.properties
            } catch {
                print("Error fetching owner properties: \(error.localizedDescription)")
            }
        }
        objectWillChange.send()
    }
    
    func fetchPropertiesPicture(propertyId: String) async throws -> UIImage? {
        return try await ownerViewModel.fetchPropertiesPicture(propertyId: propertyId)
    }
    
    func updateProperty(request: Property, token: String) async throws -> String {
        guard loginViewModel.userRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can update properties.".localized()])
        }
        let result = try await ownerViewModel.updateProperty(request: request, token: token)
        await fetchProperties()
        return result
    }
    
    func deleteProperty(propertyId: String) async throws {
        guard loginViewModel.userRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can delete properties.".localized()])
        }
        try await ownerViewModel.deleteProperty(propertyId: propertyId)
        await fetchProperties()
    }
    
    func fetchPropertyDocuments(propertyId: String) async throws -> [PropertyDocument] {
        let documents: [PropertyDocument]
        if loginViewModel.userRole == "tenant" {
            if let leaseId = try await fetchActiveLeaseIdForProperty(propertyId: propertyId, token: try await TokenStorage.getValidAccessToken()) {
                documents = try await tenantViewModel.fetchTenantPropertyDocuments(leaseId: leaseId, propertyId: propertyId)
            } else {
                throw NSError(domain: "", code: 404, userInfo: [NSLocalizedDescriptionKey: "No active lease found.".localized()])
            }
        } else {
            documents = try await ownerViewModel.fetchPropertyDocuments(propertyId: propertyId)
        }
        
        if let index = properties.firstIndex(where: { $0.id == propertyId }) {
            var updatedProperty = properties[index]
            updatedProperty.documents = documents
            properties[index] = updatedProperty
            objectWillChange.send()
        }
        return documents
    }
    
    func cancelInvite(propertyId: String, token: String) async throws {
        guard loginViewModel.userRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can cancel invites.".localized()])
        }
        try await ownerViewModel.cancelInvite(propertyId: propertyId, token: token)
        await fetchProperties()
    }
    
    func endLease(propertyId: String, leaseId: String, token: String) async throws {
        guard loginViewModel.userRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can end leases.".localized()])
        }
        try await ownerViewModel.endLease(propertyId: propertyId, leaseId: leaseId, token: token)
        await fetchProperties()
    }
    
    func fetchActiveLease(propertyId: String, token: String) async throws -> String? {
        if loginViewModel.userRole == "tenant" {
            return try await tenantViewModel.fetchActiveLeaseIdForProperty(propertyId: propertyId, token: token)
        } else {
            return try await ownerViewModel.fetchActiveLease(propertyId: propertyId, token: token)
        }
    }
    
    func fetchPropertyDamages(propertyId: String) async throws {
        if loginViewModel.userRole == "tenant" {
            if let leaseId = try await fetchActiveLeaseIdForProperty(propertyId: propertyId, token: try await TokenStorage.getValidAccessToken()) {
                let fetchedDamages = try await tenantViewModel.fetchTenantDamages(leaseId: leaseId)
                if let index = properties.firstIndex(where: { $0.id == propertyId }) {
                    var updatedProperty = properties[index]
                    updatedProperty.damages = fetchedDamages
                    properties[index] = updatedProperty
                    objectWillChange.send()
                }
            }
        } else {
            let fetchedDamages = try await ownerViewModel.fetchPropertyDamages(propertyId: propertyId)
            if let index = properties.firstIndex(where: { $0.id == propertyId }) {
                var updatedProperty = properties[index]
                updatedProperty.damages = fetchedDamages
                properties[index] = updatedProperty
                objectWillChange.send()
            }
        }
    }
    
    func fetchActiveLeaseIdForProperty(propertyId: String, token: String) async throws -> String? {
        if let leaseId = activeLeaseId {
            return leaseId
        }
        let leaseId = try await tenantViewModel.fetchActiveLeaseIdForProperty(propertyId: propertyId, token: token)
        self.activeLeaseId = leaseId
        return leaseId
    }
    
    func fetchPropertyRooms(propertyId: String, token: String) async throws -> [PropertyRoomsTenant] {
        if !rooms.isEmpty {
            return rooms
        }
        let fetchedRooms = try await tenantViewModel.fetchPropertyRooms(token: token)
        self.rooms = fetchedRooms
        return fetchedRooms
    }
    
    func createDamage(propertyId: String, leaseId: String, damage: DamageRequest, token: String) async throws -> String {
        guard loginViewModel.userRole == "tenant" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only tenants can report damages.".localized()])
        }
        return try await tenantViewModel.createDamage(propertyId: propertyId, leaseId: leaseId, damage: damage, token: token)
    }
    
    func fetchLastInventoryReport(propertyId: String, leaseId: String) async throws -> InventoryReportResponse? {
        guard loginViewModel.userRole == "owner" else {
            throw NSError(domain: "", code: 403, userInfo: [NSLocalizedDescriptionKey: "Only owners can fetch inventory reports.".localized()])
        }
        return try await ownerViewModel.fetchLastInventoryReport(propertyId: propertyId, leaseId: leaseId)
    }
}

struct PropertyID: Decodable {
    let id: String
}

struct PropertyImageBase64: Decodable {
    let data: String
}

private func convertUIImageToBase64(_ image: UIImage) -> String? {
    guard let imageData = image.jpegData(compressionQuality: 0.8) else {
        return nil
    }
    return imageData.base64EncodedString()
}
