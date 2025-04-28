//
//  PropertyView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import SwiftUI

struct PropertyView: View {
    @EnvironmentObject var viewModel: PropertyViewModel
    @State private var isLoading = false

    var body: some View {
        NavigationStack {
            ZStack {
                VStack(spacing: 0) {
                    TopBar(title: "Keyz".localized())
                    headerView
                    propertyListView
                }
                VStack {
                    Spacer()
                    HStack {
                        Spacer()
                        NavigationLink(destination: CreatePropertyView(viewModel: viewModel)) {
                            Image(systemName: "plus")
                                .font(.title2)
                                .foregroundColor(.white)
                                .frame(width: 50, height: 50)
                                .background(Color("LightBlue"))
                                .clipShape(Circle())
                                .shadow(radius: 4)
                        }
                        .padding(.trailing, 20)
                        .padding(.bottom, 20)
                        .accessibilityLabel("add_property")
                    }
                }
            }
            .overlay(
                isLoading ? ProgressView("Loading properties...".localized())
                    .progressViewStyle(CircularProgressViewStyle())
                    .padding()
                    .background(Color.white.opacity(0.8))
                    .cornerRadius(10)
                    : nil
            )
        }
        .onAppear {
            if !CommandLine.arguments.contains("-skipLogin") {
                Task {
                    isLoading = true
                    await viewModel.fetchProperties()
                    isLoading = false
                }
            }
        }
    }

    private var headerView: some View {
        HStack {
            Text("Real property".localized())
                .font(.title2)
                .fontWeight(.bold)
                .padding(.horizontal)
                .padding(.vertical, 10)
            Spacer()
        }
    }

    private var propertyListView: some View {
        ScrollView {
            VStack(spacing: 20) {
                if viewModel.properties.isEmpty && !isLoading {
                    Text("No properties available".localized())
                        .foregroundColor(.gray)
                        .padding()
                } else {
                    ForEach($viewModel.properties) { $property in
                        NavigationLink(destination: PropertyDetailView(property: $property, viewModel: viewModel)) {
                            PropertyCardView(property: $property)
                                .background(Color.white)
                                .cornerRadius(15)
                                .shadow(radius: 2)
                                .padding(.horizontal)
                        }
                        .accessibilityIdentifier("property_card_\(property.id)")
                    }
                }
            }
            .padding(.vertical)
        }
    }
}

struct PropertyCardView: View {
    @Binding var property: Property

    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            ZStack(alignment: .topTrailing) {
                if let uiImage = property.photo {
                    Image(uiImage: uiImage)
                        .resizable()
                        .scaledToFill()
                        .frame(height: 150)
                        .clipShape(RoundedRectangle(cornerRadius: 10))
                        .overlay(
                            RoundedRectangle(cornerRadius: 10)
                                .stroke(Color.gray.opacity(0.3), lineWidth: 1)
                        )
                } else {
                    Image("DefaultImageProperty")
                        .resizable()
                        .scaledToFill()
                        .frame(height: 150)
                        .clipShape(RoundedRectangle(cornerRadius: 10))
                        .overlay(
                            RoundedRectangle(cornerRadius: 10)
                                .stroke(Color.gray.opacity(0.3), lineWidth: 1)
                        )
                        .accessibilityLabel("image_property")
                }

                Text(property.isAvailable == "available" ? "Available".localized() : "Unavailable".localized())
                    .font(.caption)
                    .fontWeight(.medium)
                    .foregroundStyle(.white)
                    .padding(.horizontal, 8)
                    .padding(.vertical, 4)
                    .background(
                        RoundedRectangle(cornerRadius: 8)
                            .fill(property.isAvailable == "available" ? Color("GreenAlert") : Color("RedAlert"))
                    )
                    .padding(8)
                    .accessibilityLabel(property.isAvailable == "available" ? "text_available" : "text_unavailable")
            }

            Text(property.name)
                .font(.headline)
                .foregroundColor(.black)

            HStack(spacing: 4) {
                Image(systemName: "mappin.and.ellipse.circle")
                    .font(.caption)
                    .foregroundColor(Color("LightBlue"))
                Text("\(property.address), \(property.city) \(property.postalCode)")
                    .font(.subheadline)
                    .foregroundColor(Color("LightBlue"))
                    .lineLimit(1)
                    .accessibilityLabel("text_address")
            }
        }
        .padding()
    }
}

struct PropertyView_Previews: PreviewProvider {
    static var previews: some View {
        let viewModel = PropertyViewModel()
        viewModel.properties = exampleDataProperty2
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
        isAvailable: "available",
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
        isAvailable: "unavailable",
        tenantName: "Jean Dupont",
        leaseStartDate: "2024-12-01T00:00:00Z",
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
