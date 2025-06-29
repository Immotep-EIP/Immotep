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
    @Environment(\.dismiss) var dismiss

    @State private var selectedRoom: LocalRoom?
    @State private var showAddStuffAlert: Bool = false
    @State private var showDeleteConfirmationAlert: Bool = false
    @State private var stuffToDelete: LocalInventory?
    @State private var showError: Bool = false
    @State private var errorMessage: String?

    var body: some View {
        ZStack {
            contentView

            if showAddStuffAlert {
                CustomAlertWithTwoTextFields(
                    isActive: $showAddStuffAlert,
                    title: "Add an element".localized(),
                    message: "Please enter details:".localized(),
                    buttonTitle: "Add".localized(),
                    secondaryButtonTitle: "Cancel".localized(),
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
                                errorMessage = "Error adding stuff: \(error.localizedDescription)".localized()
                                showError = true
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
                    title: "Delete Stuff".localized(),
                    message: stuffToDelete != nil ? "Are you sure you want to delete the stuff?".localized() : "",
                    buttonTitle: "Delete".localized(),
                    secondaryButtonTitle: "Cancel".localized(),
                    action: {
                        if let stuffToDelete, let selectedRoom = selectedRoom {
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

            if showError, let message = errorMessage {
                ErrorNotificationView(message: message)
                    .onDisappear {
                        showError = false
                        errorMessage = nil
                        inventoryViewModel.errorMessage = nil
                    }
            }
        }
        .navigationBarBackButtonHidden(true)
        .onAppear {
            Task {
                if !inventoryViewModel.isEntryInventory {
                    await inventoryViewModel.fetchLastInventoryReport()
                }
                selectedRoom = inventoryViewModel.localRooms.first { $0.id == roomId }
                if let room = selectedRoom {
                    if room.inventory.isEmpty {
                        await inventoryViewModel.fetchStuff(room)
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
            TopBar(title: "Keyz")
                .overlay(
                    HStack {
                        Button(action: {
                            dismiss()
                        }) {
                            Image(systemName: "chevron.left")
                                .font(.title3)
                                .foregroundColor(Color("textColor"))
                                .frame(width: 40, height: 40)
                                .background(Color.black.opacity(0.2))
                                .clipShape(Circle())
                        }
                        .padding(.trailing, 16)
                    },
                    alignment: .trailing
                )
            if let selectedRoom = selectedRoom {
                VStack {
                    Spacer()
                    stuffListView(room: selectedRoom)
                    addStuffButton
                    if selectedRoom.inventory.allSatisfy({ $0.checked }) || !selectedRoom.inventory.isEmpty {
                        if selectedRoom.checked {
                            confirmButton
                        } else {
                            analyzeRoomButton
                        }
                    }
                }
                Spacer()
            } else {
                Text("Room not found".localized())
            }
        }
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

    private var confirmButton: some View {
        Button {
            Task {
                if let room = selectedRoom {
                    await inventoryViewModel.markRoomAsChecked(room)
                    dismiss()
                }
            }
        } label: {
            Text("Confirm".localized())
                .padding()
                .frame(maxWidth: .infinity)
                .background(Color("LightBlue"))
                .foregroundColor(Color.white)
                .cornerRadius(10)
        }
        .padding()
    }

    private var analyzeRoomButton: some View {
        NavigationLink(
            destination: {
                if inventoryViewModel.isEntryInventory {
                    InventoryRoomEvaluationView(selectedRoom: selectedRoom!)
                        .environmentObject(inventoryViewModel)
                } else {
                    InventoryRoomExitEvaluationView(selectedRoom: selectedRoom!)
                        .environmentObject(inventoryViewModel)
                }
            },
            label: {
                Text("Analyze Room".localized())
                    .padding()
                    .frame(maxWidth: .infinity)
                    .background(Color("LightBlue"))
                    .foregroundColor(Color.white)
                    .cornerRadius(10)
            }
        )
        .padding()
    }
}

struct StuffCard: View {
    let stuff: LocalInventory
    @EnvironmentObject var viewModel: InventoryViewModel

    var body: some View {
        HStack {
            if viewModel.checkedStuffStatus[stuff.id] == true {
                Image(systemName: "checkmark")
                    .foregroundStyle(Color.green)
            }
            Text(stuff.name)
                .foregroundStyle(Color("textColor"))
            if !stuff.images.isEmpty {
                Image(systemName: "photo")
                    .foregroundStyle(.gray)
            }
            if !stuff.comment.isEmpty {
                Image(systemName: "text.bubble")
                    .foregroundStyle(Color.blue)
            }
        }
    }
}
