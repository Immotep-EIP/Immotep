//
//  InventoryStuffView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI

struct InventoryStuffView: View {
    @EnvironmentObject var inventoryViewModel: InventoryViewModel
    let roomId: String

    @State private var selectedRoom: LocalRoom?
    @State private var showAddStuffAlert: Bool = false
    @State private var showDeleteConfirmationAlert: Bool = false
    @State private var stuffToDelete: LocalInventory?

    var body: some View {
        NavigationView {
            ZStack {
                contentView
                alertViews
            }
        }
        .navigationBarBackButtonHidden(true)
        .onAppear {
            Task {
                selectedRoom = inventoryViewModel.localRooms.first { $0.id == roomId }
                if let room = selectedRoom {
                    if room.inventory.isEmpty {
                        await inventoryViewModel.fetchStuff(room)
                    }
                    if inventoryViewModel.isRoomCompleted(room) {
                        await inventoryViewModel.markRoomAsChecked(room)
                    }
                    inventoryViewModel.selectRoom(room)
                }
            }
        }
        .onChange(of: inventoryViewModel.localRooms) {
            if let updatedRoom = inventoryViewModel.localRooms.first(where: { $0.id == roomId }) {
                selectedRoom = updatedRoom
            }
        }
    }

    private var contentView: some View {
        VStack(spacing: 0) {
            TopBar(title: "Inventory")
            if let selectedRoom = selectedRoom {
                VStack {
                    Spacer()
                    stuffListView(room: selectedRoom)
                    addStuffButton
                }
                Spacer()
                TaskBar()
            } else {
                Text("Room not found")
            }
        }
        .navigationTransition(.fade(.in).animation(.easeInOut(duration: 0)))
    }

    private func stuffListView(room: LocalRoom) -> some View {
        List {
            ForEach(room.inventory) { stuff in
                NavigationLink(
                    destination: {
                        if inventoryViewModel.isEntryInventory {
                            InventoryEntryEvaluationView(selectedStuff: stuff)
                                .environmentObject(inventoryViewModel)
                        } else {
                            InventoryExitEvaluationView(selectedStuff: stuff)
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
    }

    private var addStuffButton: some View {
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

    private var alertViews: some View {
        Group {
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
                                if let selectedRoom = selectedRoom {
                                    try await inventoryViewModel.addStuff(name: name, quantity: quantity, to: selectedRoom)
                                    if let updatedRoom = inventoryViewModel.localRooms.first(where: { $0.id == roomId }) {
                                        self.selectedRoom = updatedRoom
                                    }
                                }
                            } catch {
                                print("Error adding stuff: \(error.localizedDescription)")
                            }
                        }
                    },
                    secondaryAction: {}
                )
                .accessibilityIdentifier("AddStuffAlert")
            }

            if showDeleteConfirmationAlert {
                CustomAlertTwoButtons(
                    isActive: $showDeleteConfirmationAlert,
                    title: "Delete Stuff",
                    message: stuffToDelete != nil ? "Are you sure you want to delete the stuff \(stuffToDelete!.name)?" : "",
                    buttonTitle: "Delete",
                    secondaryButtonTitle: "Cancel",
                    action: {
                        if let stuffToDelete = stuffToDelete, let selectedRoom = selectedRoom {
                            Task {
                                await inventoryViewModel.deleteStuff(stuffToDelete, from: selectedRoom)
                                if let updatedRoom = inventoryViewModel.localRooms.first(where: { $0.id == roomId }) {
                                    self.selectedRoom = updatedRoom
                                }
                            }
                        }
                    },
                    secondaryAction: {
                        stuffToDelete = nil
                    }
                )
                .accessibilityIdentifier("DeleteStuffAlert")
            }
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

// struct InventoryStuffView_Previews: PreviewProvider {
//    static var previews: some View {
//        let fakeProperty = exampleDataProperty
//        _ = InventoryViewModel(property: fakeProperty)
//
//        let exampleLocalRoom = LocalRoom(
//            id: fakeProperty.rooms[0].id,
//            name: fakeProperty.rooms[0].name,
//            checked: fakeProperty.rooms[0].checked,
//            inventory: fakeProperty.rooms[0].inventory.map { inventory in
//                LocalInventory(
//                    id: inventory.id,
//                    propertyId: inventory.propertyId,
//                    roomId: inventory.roomId,
//                    name: inventory.name,
//                    quantity: inventory.quantity,
//                    checked: inventory.checked,
//                    images: inventory.images,
//                    status: inventory.status,
//                    comment: inventory.comment
//                )
//            }
//        )
//
//        struct PreviewWrapper: View {
//            @State private var selectedRoom: LocalRoom
//
//            init(selectedRoom: LocalRoom) {
//                self._selectedRoom = State(initialValue: selectedRoom)
//            }
//
//            var body: some View {
//                InventoryStuffView(selectedRoom: $selectedRoom)
//                    .environmentObject(InventoryViewModel(property: exampleDataProperty))
//            }
//        }
//
//        return PreviewWrapper(selectedRoom: exampleLocalRoom)
//    }
// }
