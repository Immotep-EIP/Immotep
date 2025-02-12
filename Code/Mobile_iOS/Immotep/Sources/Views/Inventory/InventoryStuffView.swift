//
//  InventoryStuffView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI

struct InventoryStuffView: View {
    @EnvironmentObject var inventoryViewModel: InventoryViewModel
    @State private var selectedRoom: PropertyRooms

    @State private var showAddStuffAlert: Bool = false
    @State private var showDeleteConfirmationAlert: Bool = false
    @State private var stuffToDelete: RoomInventory?

    init(selectedRoom: PropertyRooms) {
        self._selectedRoom = State(initialValue: selectedRoom)
    }

    var body: some View {
        NavigationView {
            ZStack {
                VStack(spacing: 0) {
                    TopBar(title: "Inventory")
                    VStack {
                        Spacer()
                        List {
                            ForEach(inventoryViewModel.selectedInventory) { stuff in
                                ZStack(alignment: .leading) {
                                    StuffCard(stuff: stuff)
                                        .onTapGesture {
                                            inventoryViewModel.selectStuff(stuff)
                                        }

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
                                            EmptyView()
                                        }
                                    )
                                    .opacity(0.0)
                                }
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
                                // Revenir à InventoryRoomView
                                // Exemple : presentationMode.wrappedValue.dismiss()
                            }
                        }
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
    }
}

struct StuffCard: View {
    let stuff: RoomInventory
    var body: some View {
        HStack {
            if stuff.checked {
                Image(systemName: "checkmark")
                    .foregroundStyle(Color.green)
            }
            Text(stuff.name)
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

struct InventoryStuffView_Previews: PreviewProvider {
    static var previews: some View {
        let fakeProperty = exampleDataProperty
        let viewModel = InventoryViewModel(property: fakeProperty)
        viewModel.selectRoom(fakeProperty.rooms[0])
        return InventoryStuffView(selectedRoom: fakeProperty.rooms[0])
            .environmentObject(viewModel)
    }
}
