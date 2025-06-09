//
//  TestImmotepApp.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 09/12/2024.
//

import SwiftUI

struct TestImmotepView: View {
    @AppStorage("isLoggedIn") var isLoggedIn: Bool = false
    @StateObject private var loginViewModel = LoginViewModel()
    @State private var propertyExample: Property = exampleDataProperty
    @State private var mockProperty = Property(
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
        isAvailable: "available",
        tenantName: nil,
        leaseId: "leaseID",
        leaseStartDate: nil,
        leaseEndDate: nil,
        documents: [],
        createdAt: "2023-10-26T10:00:00Z",
        rooms: [
            PropertyRooms(id: "room1", name: "Salon", checked: false, inventory: [])
        ],
        damages: []
    )

    var body: some View {
        let isUITestMode = CommandLine.arguments.contains("-skipLogin")
        Group {
            if isLoggedIn || isUITestMode {
                if CommandLine.arguments.contains("-propertyList") {
                    propertyListView()
//                } else if CommandLine.arguments.contains("-inventoryTypeView") {
//                    InventoryTypeView(property: $propertyExample)
                } else if CommandLine.arguments.contains("-inventoryRoomView") {
                    inventoryRoomView()
                } else if CommandLine.arguments.contains("-inventoryStuffView") {
                    inventoryStuffView()
                } else if CommandLine.arguments.contains("-inventoryEntryEvaluationView") {
                    inventoryEntryEvaluationView()
                } else if CommandLine.arguments.contains("-inventoryExitEvaluationView") {
                    inventoryExitEvaluationView()
                } else if CommandLine.arguments.contains("-createPropertyView") {
                    propertyCreateView()
                } else if CommandLine.arguments.contains("-editPropertyView") {
                    propertyEditView()
                } else {
                    OverviewView()
                        .environmentObject(loginViewModel)
                }
            } else {
                if isLoggedIn {
                    OverviewView()
                        .environmentObject(loginViewModel)
                } else {
                    LoginView()
                        .environmentObject(loginViewModel)
                }
            }
        }
    }

    private func propertyListView() -> some View {
        let viewModel = PropertyViewModel(loginViewModel: LoginViewModel())
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
                isAvailable: "available",
                tenantName: nil,
                leaseId: "leaseId",
                leaseStartDate: nil,
                leaseEndDate: nil,
                documents: [],
                createdAt: "2023-10-26T10:00:00Z",
                rooms: [
                    PropertyRooms(id: "room1", name: "Salon", checked: false, inventory: [])
                ],
                damages: []
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
                isAvailable: "unavailable",
                tenantName: "Jean Dupont",
                leaseId: "leaseId",
                leaseStartDate: "2023-10-26T10:00:00Z",
                leaseEndDate: nil,
                documents: [],
                createdAt: "2023-11-15T14:30:00Z",
                rooms: [
                    PropertyRooms(id: "room2", name: "Chambre", checked: true, inventory: [])
                ],
                damages: []

            )
        ]
        return PropertyView()
            .environmentObject(viewModel)
    }

    private func inventoryRoomView() -> some View {
        @State var navigateToInventory: Bool = false
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

    private func propertyCreateView() -> some View {
        let viewModel = PropertyViewModel(loginViewModel: LoginViewModel())
        return NavigationStack {
            PropertyView()
                .environmentObject(viewModel)
                .navigationDestination(isPresented: .constant(true)) {
                    CreatePropertyView(viewModel: viewModel)
                }
        }
    }

    private func propertyEditView() -> some View {
        let viewModel = PropertyViewModel(loginViewModel: LoginViewModel())
        return NavigationStack {
            PropertyView()
                .environmentObject(viewModel)
                .navigationDestination(isPresented: .constant(true)) {
                    EditPropertyView(viewModel: viewModel, property: $mockProperty)
                }
        }
    }
}

struct TestImmotepView_Previews: PreviewProvider {
    static var previews: some View {
        TestImmotepView()
    }
}
