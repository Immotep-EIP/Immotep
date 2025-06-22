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
    @State private var showError: Bool = false
    @State private var errorMessage: String?
    @State private var navigateToReportDamage: Bool = false
    @State private var navigateToInventory: Bool = false
    @State private var selectedPropertyId: String?
    @State private var rooms: [PropertyRoomsTenant] = []
    @State private var activeLeaseId: String?

    var body: some View {
        ZStack {
            NavigationStack {
                VStack(spacing: 0) {
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
                        if isLoading {
                            VStack {
                                Spacer()
                                ProgressView()
                                    .progressViewStyle(CircularProgressViewStyle())
                                Spacer()
                            }
                        } else {
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
                    }

                    if loginViewModel.userRole != "tenant" {
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
                       viewModel.properties.first(where: { $0.id == propertyId }) != nil {
                        ReportDamageView(
                            viewModel: viewModel,
                            propertyId: propertyId,
                            rooms: rooms,
                            leaseId: activeLeaseId,
                            onDamageCreated: {
                                Task {
                                    do {
                                        try await viewModel.fetchPropertyDamages(propertyId: propertyId)
                                    } catch {
                                        errorMessage = "Error refreshing damages: \(error.localizedDescription)".localized()
                                        showError = true
                                    }
                                }
                            }
                        )
                    } else {
                        Text("No property selected".localized())
                            .foregroundColor(.red)
                            .padding()
                    }
                }
            }

            if showError, let message = errorMessage {
                ErrorNotificationView(message: message)
                    .onDisappear {
                        showError = false
                        errorMessage = nil
                    }
            }
        }
        .onAppear {
            Task {
                isLoading = true
                do {
                    if loginViewModel.userRole == "tenant" {
                        await viewModel.fetchProperties()
                        tenantProperty = viewModel.properties.first
                        if let propertyId = tenantProperty?.id {
                            let token = try await TokenStorage.getValidAccessToken()
                            rooms = try await viewModel.fetchPropertyRooms(propertyId: propertyId, token: token)
                            activeLeaseId = try await viewModel.fetchActiveLeaseIdForProperty(propertyId: propertyId, token: token)
                        }
                    } else {
                        await viewModel.fetchProperties()
                    }
                } catch {
                    errorMessage = "Error fetching properties: \(error.localizedDescription)".localized()
                    showError = true
                }
                isLoading = false
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
