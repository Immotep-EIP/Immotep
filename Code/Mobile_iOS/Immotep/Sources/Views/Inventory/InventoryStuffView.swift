//
//  InventoryStuffView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI

struct InventoryStuffView: View {
    @ObservedObject var inventoryViewModel: InventoryViewModel

    var body: some View {
        NavigationView {
            VStack(spacing: 0) {
                TopBar(title: "Inventory")
                VStack {
                    ScrollView {
                        ForEach(inventoryViewModel.selectedInventory) { stuff in
                            NavigationLink(destination: InventoryEvaluationView(inventoryViewModel: inventoryViewModel)) {
                                StuffCard(stuff: stuff)
                                    .onTapGesture {
                                        inventoryViewModel.selectStuff(stuff)
                                    }
                            }
                        }
                        Button {
                            inventoryViewModel.addStuff(name: "New Stuff")
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
                    .padding(.top, 10)
                }
                TaskBar()
            }
            .navigationTransition(
                .fade(.in).animation(.easeInOut(duration: 0))
            )
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

struct InventoryStuffView_Previews: PreviewProvider {
    static var previews: some View {
        let fakeProperty = exampleDataProperty
        let viewModel = InventoryViewModel(property: fakeProperty)
        viewModel.selectRoom(fakeProperty.rooms[0])
        return InventoryStuffView(inventoryViewModel: viewModel)
    }
}
