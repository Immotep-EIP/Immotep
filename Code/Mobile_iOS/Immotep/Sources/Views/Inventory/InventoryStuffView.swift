//
//  InventoryStuffView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI
struct InventoryStuffView: View {
    @EnvironmentObject var inventoryViewModel: InventoryViewModel
    @Binding var selectedRoom: LocalRoom // Utiliser un Binding

    @State private var showAddStuffAlert: Bool = false
    @State private var showDeleteConfirmationAlert: Bool = false
    @State private var stuffToDelete: LocalInventory?

//    init(selectedRoom: LocalRoom) {
//        self._selectedRoom = State(initialValue: selectedRoom)
//        print("selectedRoom initialized: \(selectedRoom)")
//    }

    var body: some View {
        NavigationView {
            ZStack {
                VStack(spacing: 0) {
                    TopBar(title: "Inventory")
                    VStack {
                        Spacer()
                        List {
                            ForEach(selectedRoom.inventory) { stuff in
                                NavigationLink(
                                    destination: {
                                        if inventoryViewModel.isEntryInventory {
                                            InventoryEntryEvaluationView(selectedStuff: stuff)
                                                .environmentObject(inventoryViewModel)
                                        } else {
                                            InventoryExitEvaluationView()
                                                .environmentObject(inventoryViewModel)
                                        }
                                    },
                                    label: {
                                        StuffCard(stuff: stuff)
                                    }
                                )
                                .swipeActions(edge: .trailing, allowsFullSwipe: true) {
                                    Button(role: .destructive) {
                                        stuffToDelete = stuff
                                        showDeleteConfirmationAlert = true
                                    } label: {
                                        Label("Delete", systemImage: "trash")
                                    }
                                }
                                .listRowInsets(EdgeInsets())
                                .listRowSeparator(.hidden)
                                .padding()
                                .overlay(
                                    RoundedRectangle(cornerRadius: 10)
                                        .stroke(Color.gray.opacity(0.5), lineWidth: 1)
                                )
                                .padding(.horizontal)
                                .padding(.vertical, 5)
                            }
                        }
                        .listStyle(.plain)

                        Button {
                            showAddStuffAlert = true
                        } label: {
                            HStack {
                                Image(systemName: "plus.circle")
                                    .font(.title)
                            }
                            .frame(maxWidth: .infinity)
                            .foregroundStyle(Color("textColor"))
                            .padding()
                            .overlay(
                                RoundedRectangle(cornerRadius: 10)
                                    .stroke(Color.gray.opacity(0.5), lineWidth: 1)
                            )
                            .padding(.horizontal)
                            .padding(.vertical, 5)
                        }
                    }
                    .onAppear {
                        Task {
                            await inventoryViewModel.fetchStuff(selectedRoom)
                            if inventoryViewModel.isRoomCompleted(selectedRoom) {
                                await inventoryViewModel.markRoomAsChecked(selectedRoom)
                            }
                        }
//                        if !selectedRoom.inventory.isEmpty {
//                            print("selected room inventory 0: \(selectedRoom.inventory[0])")
//                        } else {
//                            print("selected room inventory is empty")
//                        }
                    }
                    Spacer()
                    TaskBar()
                }
                .navigationTransition(
                    .fade(.in).animation(.easeInOut(duration: 0))
                )

                if showAddStuffAlert {
                    CustomAlertWithTwoTextFields(
                        isActive: $showAddStuffAlert,
                        title: "Add an element",
                        message: "Please enter details:",
                        buttonTitle: "Add",
                        secondaryButtonTitle: "Cancel",
                        action: { name, quantity in
                            Task {
                                do {
                                    try await inventoryViewModel.addStuff(name: name, quantity: quantity, to: selectedRoom)
                                } catch {
                                    print("Error adding stuff: \(error.localizedDescription)")
                                }
                            }
                        },
                        secondaryAction: {
                        }
                    )
                }

                if showDeleteConfirmationAlert {
                    CustomAlertTwoButtons(
                        isActive: $showDeleteConfirmationAlert,
                        title: "Delete Stuff",
                        message: stuffToDelete != nil ? "Are you sure you want to delete the stuff \(stuffToDelete!.name)?" : "",
                        buttonTitle: "Delete",
                        secondaryButtonTitle: "Cancel",
                        action: {
                            if let stuffToDelete = stuffToDelete {
                                Task {
                                    await inventoryViewModel.deleteStuff(stuffToDelete, from: selectedRoom)
                                }
                            }
                        },
                        secondaryAction: {
                            stuffToDelete = nil
                        }
                    )
                }
            }
        }
        .navigationBarBackButtonHidden(true)
        .onAppear {
            inventoryViewModel.selectRoom(selectedRoom)
        }
    }
}

struct StuffCard: View {
    let stuff: LocalInventory
    @EnvironmentObject var inventoryViewModel: InventoryViewModel

    var body: some View {
        HStack {
            if inventoryViewModel.checkedStuffStatus[stuff.id] == true {
                Image(systemName: "checkmark")
                    .foregroundStyle(Color.green)
            }
            Text(stuff.name)
                .foregroundStyle(Color("textColor"))

            if !stuff.images.isEmpty {
                Image(systemName: "photo")
                    .foregroundStyle(Color.blue)
            }

            if !stuff.comment.isEmpty {
                Image(systemName: "text.bubble")
                    .foregroundStyle(Color.orange)
            }
        }
    }
}

struct InventoryStuffView_Previews: PreviewProvider {
    static var previews: some View {
        let fakeProperty = exampleDataProperty
        _ = InventoryViewModel(property: fakeProperty)

        // Créez une instance de LocalRoom pour l'aperçu
        let exampleLocalRoom = LocalRoom(
            id: fakeProperty.rooms[0].id,
            name: fakeProperty.rooms[0].name,
            checked: fakeProperty.rooms[0].checked,
            inventory: fakeProperty.rooms[0].inventory.map { inventory in
                LocalInventory(
                    id: inventory.id,
                    propertyId: inventory.propertyId,
                    roomId: inventory.roomId,
                    name: inventory.name,
                    quantity: inventory.quantity,
                    checked: inventory.checked,
                    images: inventory.images,
                    status: inventory.status,
                    comment: inventory.comment
                )
            }
        )

        // Utilisez @State pour créer un Binding
        struct PreviewWrapper: View {
            @State private var selectedRoom: LocalRoom

            init(selectedRoom: LocalRoom) {
                self._selectedRoom = State(initialValue: selectedRoom)
            }

            var body: some View {
                InventoryStuffView(selectedRoom: $selectedRoom)
                    .environmentObject(InventoryViewModel(property: exampleDataProperty))
            }
        }

        return PreviewWrapper(selectedRoom: exampleLocalRoom)
    }
}

// selectedRoom.inventory est vide, voir le remplissage des stuffs dans chaque room
