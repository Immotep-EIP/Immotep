import SwiftUI

struct InventoryRoomView: View {
    @EnvironmentObject var inventoryViewModel: InventoryViewModel
    @AppStorage("theme") private var selectedTheme: String = ThemeOption.system.rawValue

    @State private var newRoomName: String = ""
    @State private var showAddRoomAlert: Bool = false
    @State private var showDeleteConfirmationAlert: Bool = false
    @State private var roomToDelete: LocalRoom?
    @State private var showCompletionMessage: Bool = false
    @State private var showErrorAlert: Bool = false

    var body: some View {
        ZStack {
            VStack(spacing: 0) {
                TopBar(title: "Keyz")
                VStack {
                    Spacer()
                    RoomListView(showDeleteConfirmationAlert: $showDeleteConfirmationAlert, roomToDelete: $roomToDelete)
                        .environmentObject(inventoryViewModel)

                    AddRoomButton(showAddRoomAlert: $showAddRoomAlert)

                    if inventoryViewModel.areAllRoomsCompleted() {
                        Button(action: {
                            Task {
                                do {
                                    try await inventoryViewModel.finalizeInventory()
                                    showCompletionMessage = true
                                } catch {
                                    showCompletionMessage = true
                                }
                            }
                        }, label: {
                            Text("Finalize Inventory".localized())
                                .padding()
                                .frame(maxWidth: .infinity)
                                .background(Color.blue)
                                .foregroundColor(.white)
                                .cornerRadius(10)
                        })
                        .padding()
                    }
                }
                Spacer()
            }
            .onAppear {
                Task {
                    await inventoryViewModel.fetchRooms()
                }
            }

            if showAddRoomAlert {
                CustomAlertWithTextAndDropdown(
                    isActive: $showAddRoomAlert,
                    textFieldInput: $newRoomName,
                    title: "Add a Room".localized(),
                    message: "Choose a name and type for your new room:".localized(),
                    buttonTitle: "Add".localized(),
                    secondaryButtonTitle: "Cancel".localized(),
                    action: { roomName, roomType in
                        Task {
                            do {
                                try await inventoryViewModel.addRoom(name: roomName, type: roomType)
                                newRoomName = ""
                            } catch {
                                inventoryViewModel.errorMessage = "Error adding room: \(error.localizedDescription)"
                                showErrorAlert = true
                            }
                        }
                    },
                    secondaryAction: {
                        newRoomName = ""
                    }
                )
                .accessibilityIdentifier("AddRoomAlert")
            }

            if showDeleteConfirmationAlert {
                CustomAlertTwoButtons(
                    isActive: $showDeleteConfirmationAlert,
                    title: "Delete Room".localized(),
                    message: roomToDelete != nil ? "Are you sure you want to delete the room \(roomToDelete!.name)?".localized() : "",
                    buttonTitle: "Delete".localized(),
                    secondaryButtonTitle: "Cancel".localized(),
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
                .accessibilityIdentifier("DeleteRoomAlert")
            }

            if showCompletionMessage, let message = inventoryViewModel.completionMessage {
                CustomAlertTwoButtons(
                    isActive: $showCompletionMessage,
                    title: inventoryViewModel.isEntryInventory ? "Entry Inventory".localized() : "Exit Inventory".localized(),
                    message: message,
                    buttonTitle: "OK",
                    secondaryButtonTitle: nil,
                    action: {},
                    secondaryAction: nil
                )
            }

            if showErrorAlert, let errorMessage = inventoryViewModel.errorMessage {
                CustomAlertTwoButtons(
                    isActive: $showErrorAlert,
                    title: "Error".localized(),
                    message: errorMessage,
                    buttonTitle: "OK",
                    secondaryButtonTitle: nil,
                    action: {
                        inventoryViewModel.errorMessage = nil
                    },
                    secondaryAction: nil
                )
                .accessibilityIdentifier("ErrorAlert")
            }
        }
        .navigationBarBackButtonHidden(true)
    }
}

struct RoomListView: View {
    @EnvironmentObject var inventoryViewModel: InventoryViewModel
    @Binding var showDeleteConfirmationAlert: Bool
    @Binding var roomToDelete: LocalRoom?

    var body: some View {
        List {
            ForEach(inventoryViewModel.localRooms) { room in
                NavigationLink(destination: {
                    InventoryStuffView(roomId: room.id)
                        .environmentObject(inventoryViewModel)
                }, label: {
                    RoomCard(room: room, isEntryInventory: inventoryViewModel.isEntryInventory)
                })
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
}

struct AddRoomButton: View {
    @Binding var showAddRoomAlert: Bool

    var body: some View {
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
}

struct RoomCard: View {
    let room: LocalRoom
    let isEntryInventory: Bool

    var body: some View {
        HStack {
            if room.checked {
                Image(systemName: "checkmark")
                    .foregroundStyle(Color.green)
            }
            Text(room.name)
                .foregroundStyle(Color("textColor"))
        }
    }
}

struct InventoryRoomView_Previews: PreviewProvider {
    static var previews: some View {
        let fakeProperty = exampleDataProperty
        let viewModel = InventoryViewModel(property: fakeProperty, isEntryInventory: false)
        InventoryRoomView()
            .environmentObject(viewModel)
    }
}
