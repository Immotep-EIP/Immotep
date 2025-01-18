//
//  InventoryTypeView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI

struct InventoryTypeView: View {
    @Binding var property: Property
    @StateObject private var inventoryViewModel: InventoryViewModel

    init(property: Binding<Property>) {
        self._property = property
        self._inventoryViewModel = StateObject(wrappedValue: InventoryViewModel(property: property.wrappedValue))
    }

    var body: some View {
        NavigationView {
            VStack(spacing: 0) {
                TopBar(title: "Inventory")
                VStack {
                    NavigationLink {
                        InventoryRoomView(inventoryViewModel: inventoryViewModel)
                    } label: {
                        HStack {
                            Image(systemName: "figure.walk.arrival")
                                .foregroundStyle(Color("textColor"))
                                .fontWeight(.bold)
                                .font(.title2)
                            Text("Entry Inventory")
                                .foregroundStyle(Color("textColor"))
                            Spacer()
                            Image(systemName: "arrowshape.right.circle.fill")
                                .foregroundStyle(Color("textColor"))
                                .fontWeight(.bold)
                                .font(.title)
                        }
                    }
                    .frame(maxWidth: .infinity)
                    .padding()
                    .overlay(
                        RoundedRectangle(cornerRadius: 10)
                            .stroke(Color.gray.opacity(0.5), lineWidth: 1)
                    )
                    .padding(.horizontal)
                    .onTapGesture {
                        inventoryViewModel.isEntryInventory = true
                    }

                    NavigationLink {
                        InventoryRoomView(inventoryViewModel: inventoryViewModel)
                    } label: {
                        HStack {
                            Image(systemName: "figure.walk.departure")
                                .foregroundStyle(Color("textColor"))
                                .fontWeight(.bold)
                                .font(.title2)
                            Text("Exit Inventory")
                                .foregroundStyle(Color("textColor"))
                            Spacer()
                            Image(systemName: "arrowshape.right.circle.fill")
                                .foregroundStyle(Color("textColor"))
                                .fontWeight(.bold)
                                .font(.title)
                        }
                    }
                    .frame(maxWidth: .infinity)
                    .padding()
                    .overlay(
                        RoundedRectangle(cornerRadius: 10)
                            .stroke(Color.gray.opacity(0.5), lineWidth: 1)
                    )
                    .padding()
                    .onTapGesture {
                        inventoryViewModel.isEntryInventory = false
                    }
                }
                .padding(.top, 20)

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

struct InventoryTypeView_Previews: PreviewProvider {
    static var previews: some View {
        let fakeProperty = exampleDataProperty
        InventoryTypeView(property: .constant(fakeProperty))
    }
}
