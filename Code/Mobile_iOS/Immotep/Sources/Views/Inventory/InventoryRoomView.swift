//
//  InventoryRoomView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI

struct InventoryRoomView: View {
    @EnvironmentObject var inventoryViewModel: InventoryViewModel
    @AppStorage("theme") private var selectedTheme: String = ThemeOption.system.rawValue

    @State private var newRoomName: String = ""
    @State private var showAddRoomAlert: Bool = false
    @State private var showDeleteConfirmationAlert: Bool = false
    @State private var roomToDelete: PropertyRooms?

    var body: some View {
        NavigationView {
            ZStack {
                VStack(spacing: 0) {
                    TopBar(title: inventoryViewModel.isEntryInventory ? "Entry Inventory" : "Exit Inventory")
                    VStack {
                        Spacer()
                        List {
                            ForEach(inventoryViewModel.property.rooms) { room in
                                ZStack(alignment: .leading) {
                                    RoomCard(room: room, isEntryInventory: inventoryViewModel.isEntryInventory)
                                        .onTapGesture {
                                            inventoryViewModel.selectRoom(room)
                                        }

                                    NavigationLink(destination:
                                                    InventoryStuffView(selectedRoom: room)
                                                        .environmentObject(inventoryViewModel)
                                    ) {
                                        EmptyView()
                                    }
                                    .opacity(0.0)
                                }
                                .swipeActions(edge: .trailing, allowsFullSwipe: true) {
                                    Button(role: .destructive) {
                                        roomToDelete = room
                                        showDeleteConfirmationAlert = true
                                    } label: {
                                        Label("Delete", systemImage: "trash")
                                    }
                                }
                                .listRowInsets(EdgeInsets())
                                .listRowSeparator(.hidden)
                            }
                        }
                        .listStyle(.plain)

                        Button {
                            showAddRoomAlert = true
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
                    Spacer()

                    TaskBar()
                }
                .navigationTransition(
                    .fade(.in).animation(.easeInOut(duration: 0))
                )
                .onAppear {
                    Task {
                        await inventoryViewModel.fetchRooms()
                    }
                }

                if showAddRoomAlert {
                    CustomAlert(
                        isActive: $showAddRoomAlert,
                        title: "Add a Room",
                        message: "Choose a name for your new room:",
                        buttonTitle: "Add",
                        secondaryButtonTitle: "Cancel",
                        showTextField: true,
                        action: {
                            Task {
                                do {
                                    try await inventoryViewModel.addRoom(name: newRoomName)
                                    newRoomName = ""
                                } catch {
                                    print("Error adding room: \(error.localizedDescription)")
                                }
                            }
                        },
                        secondaryAction: {
                            newRoomName = ""
                        }
                    )
                }

                if showDeleteConfirmationAlert {
                    CustomAlert(
                        isActive: $showDeleteConfirmationAlert,
                        title: "Delete Room",
                        message: roomToDelete != nil ? "Are you sure you want to delete the room \(roomToDelete!.name)?" : "",
                        buttonTitle: "Delete",
                        secondaryButtonTitle: "Cancel",
                        showTextField: false,
                        action: {
                            if let roomToDelete = roomToDelete {
                                Task {
                                    await inventoryViewModel.deleteRoom(roomToDelete)
                                }
                            }
                        },
                        secondaryAction: {
                            roomToDelete = nil
                        }
                    )
                }
            }
        }
        .navigationBarBackButtonHidden(true)
    }
}

struct RoomCard: View {
    let room: PropertyRooms
    let isEntryInventory: Bool

    var body: some View {
        HStack {
            if room.checked {
                Image(systemName: "checkmark")
                    .foregroundStyle(Color.green)
            }
            Text(room.name)
                .foregroundStyle(Color("textColor"))
            Spacer()
            Image(systemName: "chevron.right")
                .font(.title2)
                .foregroundStyle(Color("textColor"))
        }
        .frame(maxWidth: .infinity)
        .padding()
        .overlay(
            RoundedRectangle(cornerRadius: 10)
                .stroke(Color.gray.opacity(0.5), lineWidth: 1)
        )
        .padding(.horizontal)
        .padding(.vertical, 5)
    }
}

struct InventoryRoomView_Previews: PreviewProvider {
    static var previews: some View {
        let fakeProperty = exampleDataProperty
        let viewModel = InventoryViewModel(property: fakeProperty)
        InventoryRoomView()
            .environmentObject(viewModel)
    }
}

extension View {
    func customAlert(
        isActive: Binding<Bool>,
        title: String,
        message: String,
        buttonTitle: String,
        secondaryButtonTitle: String?,
        showTextField: Bool,
        action: @escaping () -> Void,
        secondaryAction: (() -> Void)?
    ) -> some View {
        self.background(
            CustomAlert(
                isActive: isActive,
                title: title,
                message: message,
                buttonTitle: buttonTitle,
                secondaryButtonTitle: secondaryButtonTitle,
                showTextField: showTextField,
                action: action,
                secondaryAction: secondaryAction
            )
        )
    }
}
