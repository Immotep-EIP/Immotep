//
//  PropertyViewModelProtocol.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/06/2025.
//

import Foundation
import SwiftUI

@MainActor
protocol PropertyViewModelProtocol: ObservableObject {
    var properties: [Property] { get set }
    var damages: [DamageResponse] { get set }
    var isFetchingDamages: Bool { get set }
    var damagesError: String? { get set }

    func fetchProperties() async
    func fetchPropertyDocuments(propertyId: String) async throws -> [PropertyDocument]
    func fetchPropertyDamages(propertyId: String) async throws
    func fetchActiveLeaseIdForProperty(propertyId: String, token: String) async throws -> String?
    func fetchLastInventoryReport(propertyId: String, leaseId: String) async throws -> InventoryReportResponse?
    func fetchTenantDamages(leaseId: String) async throws
}
