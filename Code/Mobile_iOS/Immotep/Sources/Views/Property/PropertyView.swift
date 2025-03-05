//
//  PropertyView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import SwiftUI

struct PropertyView: View {
    @StateObject private var viewModel = PropertyViewModel()
    @State private var isCreatingProperty = false
    @State private var showDeleteConfirmationAlert = false
    @State private var propertyToDelete: Property?
    @State private var navigateToEditId: String?
    @State private var listRefreshID = UUID()

    var body: some View {
        NavigationStack {
            ZStack {
                VStack(spacing: 0) {
                    TopBar(title: "Property".localized())
                    headerView
                    propertyListView
                    TaskBar()
                }
                .navigationTransition(.fade(.in).animation(.easeInOut(duration: 0)))

                if showDeleteConfirmationAlert {
                    CustomAlertTwoButtons(
                        isActive: $showDeleteConfirmationAlert,
                        title: "Delete Property".localized(),
                        message: propertyToDelete != nil ? "Are you sure you want to delete the property \(propertyToDelete!.name)?".localized() : "",
                        buttonTitle: "Delete".localized(),
                        secondaryButtonTitle: "Cancel".localized(),
                        action: {
                            if let propertyToDelete = propertyToDelete {
                                Task {
                                    await deleteProperty(propertyToDelete)
                                }
                            }
                        },
                        secondaryAction: {
                            self.propertyToDelete = nil
                        }
                    )
                }
            }
            .navigationDestination(isPresented: Binding(
                get: { navigateToEditId != nil },
                set: { if !$0 { navigateToEditId = nil } }
            )) {
                if let editId = navigateToEditId,
                   let propertyToEdit = viewModel.properties.first(where: { $0.id == editId }) {
                    EditPropertyView(viewModel: viewModel, property: Binding(
                        get: { viewModel.properties.first(where: { $0.id == editId }) ?? propertyToEdit },
                        set: { newValue in
                            if let index = viewModel.properties.firstIndex(where: { $0.id == newValue.id }) {
                                viewModel.properties[index] = newValue
                            }
                        }
                    ))
                }
            }
        }
        .navigationBarBackButtonHidden(true)
        .onAppear {
            Task {
                await viewModel.fetchProperties()
                listRefreshID = UUID()
            }
        }
        .onChange(of: viewModel.properties) {
            listRefreshID = UUID()
        }
        .onChange(of: navigateToEditId) {
            if navigateToEditId == nil {
                Task {
                    await viewModel.fetchProperties()
                }
            }
        }
    }

    private var headerView: some View {
        HStack {
            Spacer()
            NavigationLink(destination: CreatePropertyView(viewModel: viewModel)) {
                Text("Add a property".localized())
                    .font(.headline)
                    .foregroundColor(.white)
                    .padding(.horizontal)
                    .padding(.vertical, 8)
                    .background(Color.blue)
                    .cornerRadius(8)
            }
            .padding()
            .accessibilityLabel("add_property")
        }
    }

    private var propertyListView: some View {
        List {
            if !viewModel.properties.isEmpty {
                ForEach($viewModel.properties) { $property in
                    NavigationLink(destination: PropertyDetailView(property: $property, viewModel: viewModel)) {
                        PropertyCardView(property: property)
                    }
                    .swipeActions(edge: .trailing, allowsFullSwipe: true) {
                        Button(action: {
                            navigateToEditId = property.id
                        }, label: {
                            Label("Edit", systemImage: "pencil")
                        })
                        .tint(.blue)

                        Button(role: .destructive) {
                            propertyToDelete = property
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
                    .padding(.vertical, 15)
                }
            } else {
                Text("No properties available".localized())
                    .foregroundColor(.gray)
                    .padding()
            }
        }
        .listStyle(.plain)
        .id(listRefreshID)
    }

    private func deleteProperty(_ property: Property) async {
        do {
            try await viewModel.deleteProperty(propertyId: property.id)
            await viewModel.fetchProperties()
            propertyToDelete = nil
        } catch {
            print("Error deleting property: \(error)")
        }
    }
}
struct PropertyCardView: View {
    let property: Property

    var body: some View {
        ZStack(alignment: .topLeading) {
            VStack {
                HStack {
                    if let uiImage = property.photo {
                        Image(uiImage: uiImage)
                            .resizable()
                            .scaledToFill()
                            .frame(width: 50, height: 50)
                            .clipShape(Circle())
                            .overlay(Circle().stroke(Color.black, lineWidth: 1))
                    } else {
                        Image("DefaultImageProperty")
                            .resizable()
                            .scaledToFit()
                            .frame(width: 50, height: 50)
                            .clipShape(Circle())
                            .overlay(Circle().stroke(Color("textColor"), lineWidth: 1))
                            .accessibilityLabel("image_property")
                    }

                    VStack(alignment: .leading, spacing: 4) {
                        Text(property.name)
                            .font(.headline)
                            .padding(.trailing, 25)
                        Text(property.address)
                            .font(.subheadline)
                            .padding(.trailing, 25)
                            .lineLimit(2)
                            .accessibilityLabel("text_address")

                        if let tenant = property.tenantName {
                            Text(tenant)
                                .font(.subheadline)
                                .foregroundColor(.secondary)
                                .accessibilityLabel("text_tenant")
                        }

                        if let leaseStart = property.leaseStartDate {
                            Text(
                                String(
                                    format: "started_on".localized(),
                                    dateFormatter.string(from: leaseStart)
                                ))
                                .font(.subheadline)
                                .foregroundColor(.secondary)
                                .accessibilityLabel("text_started_on")
                        }
                    }
                }
                .padding(.trailing, 16)
            }

            if property.isAvailable {
                Text("Available".localized())
                    .font(.caption)
                    .foregroundColor(.green)
                    .padding(.horizontal, 8)
                    .padding(.vertical, 4)
                    .background(Capsule().fill(Color.green.opacity(0.2)))
                    .frame(maxWidth: .infinity, alignment: .topTrailing)
                    .accessibilityLabel("text_available")
            } else {
                Text("Busy".localized())
                    .font(.caption)
                    .foregroundColor(.red)
                    .padding(.horizontal, 8)
                    .padding(.vertical, 4)
                    .background(Capsule().fill(Color.red.opacity(0.2)))
                    .frame(maxWidth: .infinity, alignment: .topTrailing)
                    .accessibilityLabel("text_busy")
            }
        }
    }
}

private let dateFormatter: DateFormatter = {
    let formatter = DateFormatter()
    formatter.dateStyle = .medium
    return formatter
}()

struct PropertyView_Previews: PreviewProvider {
    static var previews: some View {
        let viewModel = PropertyViewModel()
        viewModel.properties = exampleDataProperty2
        print("Properties in preview: \(viewModel.properties.count)")
        return PropertyView()
            .environmentObject(viewModel)
            .onAppear {
                viewModel.properties = exampleDataProperty2
            }
    }
}

let exampleDataProperty2: [Property] = [
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
        isAvailable: true,
        tenantName: nil,
        leaseStartDate: nil,
        leaseEndDate: nil,
        documents: [],
        createdAt: "2023-10-26T10:00:00Z",
        rooms: [
            PropertyRooms(
                id: "room1",
                name: "Salon",
                checked: false,
                inventory: []
            )
        ]
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
        isAvailable: false,
        tenantName: "Jean Dupont",
        leaseStartDate: Date(),
        leaseEndDate: nil,
        documents: [],
        createdAt: "2023-11-15T14:30:00Z",
        rooms: [
            PropertyRooms(
                id: "room2",
                name: "Chambre",
                checked: true,
                inventory: []
            )
        ]
    )
]
