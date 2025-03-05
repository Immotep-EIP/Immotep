//
//  TestImmotepApp.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 09/12/2024.
//

import SwiftUI

struct TestImmotepView: View {
    @AppStorage("isLoggedIn") var isLoggedIn: Bool = false
    @StateObject private var profileViewModel = ProfileViewModel()
    @State private var propertyExample: Property = exampleDataProperty

    var body: some View {
        let isUITestMode = CommandLine.arguments.contains("-skipLogin")
        Group {
            if isLoggedIn || isUITestMode {
                if CommandLine.arguments.contains("-propertyList") {
                    propertyListView()
                } else if CommandLine.arguments.contains("-inventoryTypeView") {
                    InventoryTypeView(property: $propertyExample)
                } else if CommandLine.arguments.contains("-inventoryRoomView") {
                    inventoryRoomView()
                } else if CommandLine.arguments.contains("-inventoryStuffView") {
                    inventoryStuffView()
                } else if CommandLine.arguments.contains("-inventoryEntryEvaluationView") {
                    inventoryEntryEvaluationView()
                } else if CommandLine.arguments.contains("-inventoryExitEvaluationView") {
                    inventoryExitEvaluationView()
                } else {
                    OverviewView()
                        .environmentObject(profileViewModel)
                }
            } else {
                if isLoggedIn {
                    OverviewView()
                        .environmentObject(profileViewModel)
                } else {
                    LoginView()
                }
            }
        }
    }

    private func propertyListView() -> some View {
        let viewModel = PropertyViewModel()
        viewModel.properties = [
            Property(
                id: "cm7gijdee000ly7i82uq0qf35",
                ownerID: "owner123",
                name: "Maison de Campagne",
                address: "123 Rue des Champs",
                city: "Paris",
                postalCode: "75001",
                country: "France",
                photo: UIImage(named: "DefaultImageProperty"),
                monthlyRent: 1200,
                deposit: 2400,
                surface: 85.5,
                isAvailable: true,
                tenantName: nil,
                leaseStartDate: nil,
                leaseEndDate: nil,
                documents: [
//                    PropertyDocument(id: UUID(), title: "Lease Agreement", fileName: "lease_agreement.pdf"),
//                    PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf")
                ],
                createdAt: "2023-10-26T10:00:00Z",
                rooms: [
                    PropertyRooms(id: "room1", name: "Salon", checked: false, inventory: [])
                ]
            ),
            Property(
                id: "cm7gijdee000ly7i82uq0qf36",
                ownerID: "owner124",
                name: "Appartement Moderne",
                address: "456 Avenue des Lumières",
                city: "Lyon",
                postalCode: "69002",
                country: "France",
                photo: UIImage(named: "DefaultImageProperty"),
                monthlyRent: 1500,
                deposit: 3000,
                surface: 65.0,
                isAvailable: false,
                tenantName: "Jean Dupont",
                leaseStartDate: Date(),
                leaseEndDate: nil,
                documents: [
//                    PropertyDocument(id: UUID(), title: "Lease Agreement", fileName: "lease_agreement.pdf"),
//                    PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf")
                ],
                createdAt: "2023-11-15T14:30:00Z",
                rooms: [
                    PropertyRooms(id: "room2", name: "Chambre", checked: true, inventory: [])
                ]
            )
        ]
        return PropertyView()
            .environmentObject(viewModel)
    }

    private func inventoryRoomView() -> some View {
        let viewModel = InventoryViewModel(
            property: propertyExample,
            localRooms: [
                LocalRoom(id: "1", name: "Living Room", checked: true, inventory: [
                    LocalInventory(id: "1.1", propertyId: propertyExample.id, roomId: "1", name: "Sofa", quantity: 1)
                ]),
                LocalRoom(id: "2", name: "Kitchen", checked: true, inventory: [
                    LocalInventory(id: "2.1", propertyId: propertyExample.id, roomId: "2", name: "Table", quantity: 1)
                ])
            ]
        )
        return InventoryRoomView()
            .environmentObject(viewModel)
    }

    private func inventoryStuffView() -> some View {
        let viewModel = InventoryViewModel(
            property: propertyExample,
            localRooms: [
                LocalRoom(id: "1", name: "Living Room", checked: false, inventory: [
                    LocalInventory(id: "1.1", propertyId: propertyExample.id, roomId: "1", name: "Sofa", quantity: 1),
                    LocalInventory(id: "1.2", propertyId: propertyExample.id, roomId: "1", name: "Chair", quantity: 2)
                ]),
                LocalRoom(id: "2", name: "Kitchen", checked: true, inventory: [
                    LocalInventory(id: "2.1", propertyId: propertyExample.id, roomId: "2", name: "Table", quantity: 1)
                ])
            ]
        )
        return InventoryStuffView(roomId: "1")
            .environmentObject(viewModel)
    }

    private func inventoryEntryEvaluationView() -> some View {
        let viewModel = InventoryViewModel(
            property: propertyExample,
            isEntryInventory: true,
            localRooms: [
                LocalRoom(id: "1", name: "Living Room", checked: false, inventory: [
                    LocalInventory(id: "1.1", propertyId: propertyExample.id, roomId: "1", name: "Sofa", quantity: 1),
                    LocalInventory(id: "1.2", propertyId: propertyExample.id, roomId: "1", name: "Chair", quantity: 2)
                ])
            ]
        )
        return InventoryEntryEvaluationView(selectedStuff: LocalInventory(
            id: "1.1", propertyId: propertyExample.id, roomId: "1", name: "Sofa", quantity: 1
        ))
            .environmentObject(viewModel)
    }

    private func inventoryExitEvaluationView() -> some View {
        let viewModel = InventoryViewModel(
            property: propertyExample,
            isEntryInventory: false,
            localRooms: [
                LocalRoom(id: "1", name: "Living Room", checked: false, inventory: [
                    LocalInventory(id: "1.1", propertyId: propertyExample.id, roomId: "1", name: "Sofa", quantity: 1),
                    LocalInventory(id: "1.2", propertyId: propertyExample.id, roomId: "1", name: "Chair", quantity: 2)
                ])
            ]
        )
        return InventoryExitEvaluationView(selectedStuff: LocalInventory(
            id: "1.1", propertyId: propertyExample.id, roomId: "1", name: "Sofa", quantity: 1
        ))
            .environmentObject(viewModel)
    }
}
