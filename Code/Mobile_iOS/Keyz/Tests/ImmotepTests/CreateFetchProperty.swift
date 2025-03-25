//
//  CreateFetchProperty.swift
//  ImmotepTests
//
//  Created by Liebenguth Alessio on 14/12/2024.
//

import XCTest
import SwiftUI
@testable import Immotep

final class MockPropertyViewModelTests: XCTestCase {
    private var viewModel: MockPropertyViewModel!

    override func setUp() async throws {
        viewModel = await MockPropertyViewModel()
    }

//    func testCreateProperty() async throws {
//        let testProperty = Property(
//            id: "3",
//            ownerID: "456",
//            name: "Test Property",
//            address: "789 Test Blvd",
//            city: "Test City",
//            postalCode: "12345",
//            country: "Test Country",
//            photo: nil,
//            monthlyRent: 1500,
//            deposit: 3000,
//            surface: 70.0,
//            isAvailable: true,
//            tenantName: nil,
//            leaseStartDate: nil,
//            leaseEndDate: nil,
//            documents: []
//        )
//
//        let result = try await viewModel.createProperty(request: testProperty, token: "testToken")
//
//        await MainActor.run {
//            XCTAssertEqual(result, "Property successfully created!")
//            XCTAssertEqual(viewModel.properties.count, 1)
//            XCTAssertEqual(viewModel.properties.first?.name, "Test Property")
//        }
//    }
//
//    func testFetchProperties() async {
//        await viewModel.fetchProperties()
//        await MainActor.run {
//            XCTAssertEqual(viewModel.properties.count, 2)
//            XCTAssertEqual(viewModel.properties.first?.name, "Mock Property 1")
//        }
//    }
}
