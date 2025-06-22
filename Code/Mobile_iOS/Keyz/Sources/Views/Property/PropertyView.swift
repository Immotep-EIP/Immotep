//
//  PropertyView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import SwiftUI

struct PropertyView: View {
    @EnvironmentObject var viewModel: PropertyViewModel
    @EnvironmentObject var loginViewModel: LoginViewModel
    @State private var tenantProperty: Property?
    @State private var isLoading = false
    @State private var errorMessage: String?
    @State private var navigateToReportDamage: Bool = false
    @State private var navigateToInventory: Bool = false
    @State private var selectedPropertyId: String?

    var body: some View {
        NavigationStack {
            if loginViewModel.userRole == "tenant" {
                if isLoading {
                    VStack {
                        Spacer()
                        ProgressView()
                            .progressViewStyle(CircularProgressViewStyle())
                        Spacer()
                    }
                } else if let property = tenantProperty {
                    PropertyDetailView(
                        propertyId: property.id,
                        viewModel: viewModel,
                        navigateToReportDamage: $navigateToReportDamage,
                        navigateToInventory: $navigateToInventory
                    )
                } else {
                    VStack {
                        Spacer()
                        Text("No property associated.".localized())
                            .foregroundColor(.gray)
                            .padding()
                        Spacer()
                    }
                }
            } else {
                ZStack {
                    VStack(spacing: 0) {
                        TopBar(title: "Keyz".localized())

                        Text("Real Property".localized())
                            .font(.title2)
                            .fontWeight(.bold)
                            .frame(maxWidth: .infinity, alignment: .leading)
                            .padding(.horizontal, 20)
                            .padding(.top, 10)
                            .padding(.bottom, 5)

                        if viewModel.properties.isEmpty {
                            VStack {
                                Spacer()
                                Text("No properties available".localized())
                                    .foregroundColor(.gray)
                                Spacer()
                            }
                        } else {
                            ScrollView {
                                LazyVGrid(
                                    columns: [GridItem(.flexible())],
                                    spacing: 15
                                ) {
                                    ForEach(viewModel.properties) { property in
                                        NavigationLink(
                                            destination: PropertyDetailView(
                                                propertyId: property.id,
                                                viewModel: viewModel,
                                                navigateToReportDamage: $navigateToReportDamage,
                                                navigateToInventory: $navigateToInventory
                                            )
                                            .environmentObject(loginViewModel)
                                        ) {
                                            PropertyCard(property: property)
                                        }
                                        .simultaneousGesture(TapGesture().onEnded {
                                            selectedPropertyId = property.id
                                        })
                                    }
                                }
                                .padding(.horizontal)
                                .padding(.vertical, 10)
                            }
                        }
                    }

                    VStack {
                        Spacer()
                        HStack {
                            Spacer()
                            NavigationLink {
                                CreatePropertyView(viewModel: viewModel)
                            } label: {
                                Image(systemName: "plus.circle.fill")
                                    .resizable()
                                    .scaledToFit()
                                    .frame(width: 50, height: 50)
                                    .foregroundColor(Color("LightBlue"))
                                    .background(
                                        Circle()
                                            .fill(Color.white)
                                            .shadow(radius: 4)
                                    )
                            }
                            .padding(.bottom, 30)
                            .padding(.trailing, 20)
                            .accessibilityLabel("add_property_btn")
                        }
                    }
                }
                .navigationDestination(isPresented: $navigateToReportDamage) {
                    if let propertyId = selectedPropertyId,
                       let property = viewModel.properties.first(where: { $0.id == propertyId }) {
                        ReportDamageView(
                            viewModel: viewModel,
                            propertyId: propertyId,
                            rooms: [],
                            leaseId: viewModel.activeLeaseId,
                            onDamageCreated: {
                                Task {
                                    do {
                                        try await viewModel.fetchPropertyDamages(propertyId: propertyId)
                                    } catch {
                                        errorMessage = "Error refreshing damages: \(error.localizedDescription)".localized()
                                    }
                                }
                            }
                        )
                    }
                }
                .onAppear {
                    Task {
                        await viewModel.fetchProperties()
                    }
                }
            }
        }
        .onAppear {
            if loginViewModel.userRole == "tenant" {
                Task {
                    isLoading = true
                    await viewModel.fetchProperties()
                    tenantProperty = viewModel.properties.first
                    isLoading = false
                }
            }
        }
    }
}

struct PropertyCard: View {
    let property: Property

    var body: some View {
        ZStack(alignment: .topTrailing) {
            VStack(alignment: .leading, spacing: 8) {
                if let uiImage = property.photo {
                    Image(uiImage: uiImage)
                        .resizable()
                        .scaledToFill()
                        .frame(height: 150)
                        .clipped()
                        .cornerRadius(10)
                } else {
                    Image("DefaultImageProperty")
                        .resizable()
                        .scaledToFill()
                        .frame(height: 150)
                        .clipped()
                        .cornerRadius(10)
                        .accessibilityLabel("image_property")
                }

                HStack {
                    Text(property.name)
                        .font(.headline)
                        .foregroundColor(Color("textColor"))
                    Spacer()
                    Text(statusText)
                        .font(.caption)
                        .fontWeight(.medium)
                        .foregroundColor(.white)
                        .padding(.horizontal, 8)
                        .padding(.vertical, 4)
                        .background(
                            RoundedRectangle(cornerRadius: 8)
                                .fill(statusColor)
                        )
                        .accessibilityLabel(statusAccessibilityLabel)
                }

                HStack(spacing: 4) {
                    Image(systemName: "mappin.and.ellipse.circle")
                        .font(.caption)
                        .foregroundColor(Color("LightBlue"))
                    Text("\(property.address), \(property.city), \(property.country)")
                        .font(.subheadline)
                        .foregroundColor(.gray)
                        .lineLimit(1)
                        .accessibilityLabel("text_address")
                }
            }
            .padding()
            .background(Color("basicWhiteBlack"))
            .cornerRadius(10)
            .shadow(color: Color.black.opacity(0.1), radius: 5, x: 0, y: 2)
        }
    }

    private var statusText: String {
        switch property.isAvailable {
        case "available": return "Available".localized()
        case "pending": return "Pending".localized()
        case "unavailable": return "Unavailable".localized()
        default: return "Unknown".localized()
        }
    }

    private var statusColor: Color {
        switch property.isAvailable {
        case "available": return Color("GreenAlert")
        case "pending": return .orange
        case "unavailable": return Color("RedAlert")
        default: return .gray
        }
    }

    private var statusAccessibilityLabel: String {
        switch property.isAvailable {
        case "available": return "text_available"
        case "pending": return "text_pending"
        case "unavailable": return "text_unavailable"
        default: return "text_unknown"
        }
    }
}

struct PropertyView_Previews: PreviewProvider {
    static var previews: some View {
        PropertyView()
            .environmentObject(PropertyViewModel(loginViewModel: LoginViewModel()))
            .environmentObject(LoginViewModel())
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
        leaseId: "id",
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
        ],
        damages: []
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
        leaseId: "id",
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
        ],
        damages: []
    )
]
