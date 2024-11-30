//
//  PropertyModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import Foundation
import SwiftUI

struct Property: Identifiable {
    let id: UUID
    var address: String
    var postalCode: String
    var country: String
    var photo: UIImage?
    var monthlyRent: Double
    var deposit: Double
    var surface: Double
    var isAvailable: Bool
    var tenantName: String?
    var leaseStartDate: Date?
    var leaseEndDate: Date?
    var documents: [PropertyDocument]
}

struct PropertyDocument: Identifiable {
    let id: UUID
    var title: String
    var fileName: String
}
