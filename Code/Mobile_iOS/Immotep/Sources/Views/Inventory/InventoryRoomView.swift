//
//  InventoryRoomView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI

struct InventoryRoomView: View {
    @ObservedObject var inventoryViewModel: InventoryViewModel

    var body: some View {
        NavigationView {
            VStack(spacing: 0) {
                TopBar(title: inventoryViewModel.isEntryInventory ? "Entry Inventory" : "Exit Inventory")
                VStack {
                    HStack {
                        Spacer()
                        Button {
                            // Edit Existing Rooms
                        } label: {
                            Text("Edit")
                                .font(.headline)
                                .foregroundStyle(Color("textColor"))
                                .padding(.vertical, 10)
                                .padding(.horizontal, 15)
                                .background(Color("btnColor"))
                                .cornerRadius(10)
                                .padding()
                        }
                    }
                    ScrollView {
                        ForEach(inventoryViewModel.property.rooms) { room in
                            NavigationLink(destination: InventoryStuffView(inventoryViewModel: inventoryViewModel)) {
                                RoomCard(room: room, isEntryInventory: inventoryViewModel.isEntryInventory)
                                    .onTapGesture {
                                        inventoryViewModel.selectRoom(room)
                                    }
                            }
                        }
                        Button {
                            inventoryViewModel.addRoom(name: "New Room")
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
                Spacer()

                TaskBar()
            }
            .navigationTransition(
                .fade(.in).animation(.easeInOut(duration: 0))
            )
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
            if isEntryInventory {
                Image(systemName: "figure.walk.arrival")
                    .font(.title)
                    .foregroundStyle(Color("textColor"))
            } else {
                Image(systemName: "figure.walk.departure")
                    .font(.title)
                    .foregroundStyle(Color("textColor"))
            }
            Image(systemName: "arrowshape.right.circle.fill")
                .font(.title)
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
        InventoryRoomView(inventoryViewModel: viewModel)
    }
}
