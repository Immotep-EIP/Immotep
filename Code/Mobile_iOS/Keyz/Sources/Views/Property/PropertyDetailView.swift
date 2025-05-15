//
//  PropertyDetailView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/11/2024.
//

import SwiftUI
import PDFKit

struct PropertyDetailView: View {
    @Binding var property: Property
    @ObservedObject var viewModel: PropertyViewModel
    @StateObject private var tenantViewModel = TenantViewModel()
    @State private var showInviteTenantSheet = false
    @State private var showEndLeasePopUp = false
    @State private var showCancelInvitePopUp = false
    @State private var showDeletePropertyPopUp = false
    @State private var showEditPropertyPopUp = false
    @Environment(\.dismiss) var dismiss
    @State private var selectedTab: String = "Details".localized()
    @State private var isLoading = false

    private let tabs = ["Details".localized(), "Documents".localized(), "Damages".localized()]

    var body: some View {
        ZStack {
            VStack(spacing: 0) {
                TopBar(title: "Keyz".localized())

                ScrollView {
                    VStack(alignment: .leading, spacing: 16) {
                        ZStack(alignment: .topLeading) {
                            if isLoading {
                                ProgressView()
                                    .progressViewStyle(CircularProgressViewStyle())
                                    .frame(height: 200)
                            } else if let uiImage = property.photo {
                                Image(uiImage: uiImage)
                                    .resizable()
                                    .scaledToFill()
                                    .frame(height: 200)
                                    .clipped()
                            } else {
                                Image("DefaultImageProperty")
                                    .resizable()
                                    .scaledToFill()
                                    .frame(height: 200)
                                    .clipped()
                                    .accessibilityLabel("image_property")
                            }

                            Button(action: {
                                dismiss()
                            }) {
                                Image(systemName: "chevron.left")
                                    .font(.title3)
                                    .foregroundColor(.white)
                                    .frame(width: 40, height: 40)
                                    .background(Color.black.opacity(0.6))
                                    .clipShape(Circle())
                            }
                            .padding(16)
                            .accessibilityLabel("back_button")

                            Menu {
                                Button(action: {
                                    showInviteTenantSheet = true
                                }) {
                                    Label("Invite Tenant".localized(), systemImage: "person.crop.circle.badge.plus")
                                }
                                
                                Button(action: {
                                    showEndLeasePopUp = true
                                }) {
                                    Label("End Lease".localized(), systemImage: "xmark.circle")
                                }
                                
                                Button(action: {
                                    showCancelInvitePopUp = true
                                }) {
                                    Label("Cancel Invite".localized(), systemImage: "person.crop.circle.badge.xmark")
                                }
                                
                                Button(action: {
                                    showEditPropertyPopUp = true
                                }) {
                                    Label("Edit Property".localized(), systemImage: "pencil")
                                }
                                
                                Button(action: {
                                    showDeletePropertyPopUp = true
                                }) {
                                    Label("Delete Property".localized(), systemImage: "trash")
                                }
                            } label: {
                                Image(systemName: "ellipsis")
                                    .font(.title3)
                                    .foregroundColor(.white)
                                    .frame(width: 40, height: 40)
                                    .background(Color.black.opacity(0.6))
                                    .clipShape(Circle())
                            }
                            .frame(maxWidth: .infinity, alignment: .topTrailing)
                            .padding(16)
                            .accessibilityLabel("options_button")
                        }

                        VStack(alignment: .leading, spacing: 4) {
                            HStack {
                                Text(property.name)
                                    .font(.title2)
                                    .fontWeight(.bold)
                                    .foregroundColor(Color("textColor"))

                                Text(property.isAvailable == "available" ? "Available".localized() : "Unavailable".localized())
                                    .font(.caption)
                                    .fontWeight(.medium)
                                    .foregroundColor(.white)
                                    .padding(.horizontal, 8)
                                    .padding(.vertical, 4)
                                    .background(
                                        RoundedRectangle(cornerRadius: 8)
                                            .fill(property.isAvailable == "available" ? Color("GreenAlert") : Color("RedAlert"))
                                    )
                                    .accessibilityLabel(property.isAvailable == "available" ? "text_available" : "text_unavailable")
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
                        .padding(.horizontal)

                        HStack(spacing: 10) {
                            ForEach(tabs, id: \.self) { tab in
                                Button(action: {
                                    withAnimation(.easeInOut(duration: 0.3)) {
                                        selectedTab = tab
                                    }
                                }) {
                                    Text(tab)
                                        .font(.system(size: 16, weight: .medium))
                                        .foregroundColor(selectedTab == tab ? .white : .gray)
                                        .padding(.vertical, 10)
                                        .padding(.horizontal, 16)
                                        .background(
                                            selectedTab == tab
                                                ? Color("LightBlue")
                                                : Color.gray.opacity(0.1)
                                        )
                                        .clipShape(RoundedRectangle(cornerRadius: 12))
                                        .scaleEffect(selectedTab == tab ? 1.05 : 1.0)
                                        .shadow(
                                            color: selectedTab == tab ? Color.black.opacity(0.2) : Color.clear,
                                            radius: 4, x: 0, y: 2
                                        )
                                }
                                .accessibilityLabel("tab_\(tab.lowercased())")
                            }
                        }
                        .frame(maxWidth: .infinity)
                        .padding(.vertical, 8)
                        .background(Color("basicWhiteBlack"))
                        .clipShape(RoundedRectangle(cornerRadius: 14))
                        .shadow(color: Color.black.opacity(0.1), radius: 6, x: 0, y: 8)

                        switch selectedTab {
                        case "Details".localized():
                            VStack(alignment: .leading, spacing: 16) {
                                LazyVGrid(
                                    columns: Array(repeating: GridItem(.flexible(), spacing: 15), count: 3),
                                    spacing: 15
                                ) {
                                    DetailItem(
                                        icon: "square.split.bottomrightquarter",
                                        value: "\(formattedValue(property.surface))m2",
                                        label: "Area".localized()
                                    )
                                    DetailItem(
                                        icon: "coloncurrencysign.arrow.trianglehead.counterclockwise.rotate.90",
                                        value: "\(property.monthlyRent)$",
                                        label: "Rent /month".localized()
                                    )
                                    DetailItem(
                                        icon: "eurosign.bank.building",
                                        value: "\(property.deposit)$",
                                        label: "Deposit".localized()
                                    )
                                }
                                .padding(.horizontal)

                                VStack(alignment: .leading, spacing: 8) {
                                    Text("Tenant(s)".localized())
                                        .font(.headline)
                                        .foregroundColor(Color("textColor"))
                                    Text(property.tenantName ?? "No tenant assigned".localized())
                                        .foregroundColor(.gray)
                                }
                                .padding(.horizontal)

                                VStack(alignment: .leading, spacing: 8) {
                                    Text("Dates".localized())
                                        .font(.headline)
                                        .foregroundColor(Color("textColor"))
                                    Text(formatLeaseDates())
                                        .foregroundColor(.gray)
                                }
                                .padding(.horizontal)
                            }

                        case "Documents".localized():
                            DocumentsGrid(documents: $property.documents)
                                .padding(.horizontal)

                        case "Damages".localized():
                            VStack {
                                Text("no_damages_reported".localized())
                                    .foregroundColor(.gray)
                                    .padding()
                            }

                        default:
                            EmptyView()
                        }
                    }
                    .padding(.bottom, 20)
                }

                NavigationLink {
                    InventoryTypeView(property: $property)
                } label: {
                    Text("Start the entry/exit inventory".localized())
                        .frame(maxWidth: .infinity)
                        .padding(.vertical, 15)
                        .background(Color("LightBlue"))
                        .foregroundColor(.white)
                        .cornerRadius(10)
                        .padding(.horizontal)
                        .padding(.bottom, 10)
                }
                .accessibilityLabel("inventory_btn_start")
            }
            .navigationBarBackButtonHidden(true)
            .onAppear {
                Task {
                    if !CommandLine.arguments.contains("-skipLogin") {
                        do {
                            isLoading = true
                            await viewModel.fetchProperties()
                            if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                                property = updatedProperty
                            }
                            try await viewModel.fetchPropertyDocuments(propertyId: property.id)
                            if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                                property = updatedProperty
                            }
                        } catch {
                            print("Error fetching property data: \(error.localizedDescription)")
                        }
                        isLoading = false
                    }
                }
            }
            .sheet(isPresented: $showInviteTenantSheet) {
                InviteTenantView(tenantViewModel: tenantViewModel, property: property)
            }

            if showCancelInvitePopUp {
                CustomAlertTwoButtons(
                    isActive: $showCancelInvitePopUp,
                    title: "Cancel Invite".localized(),
                    message: "Are you sure you want to cancel the pending invite?".localized(),
                    buttonTitle: "Confirm".localized(),
                    secondaryButtonTitle: "Cancel".localized(),
                    action: {},
                    secondaryAction: {}
                )
                .accessibilityIdentifier("InviteTenantAlert")
            }
            if showEndLeasePopUp {
                CustomAlertTwoButtons(
                    isActive: $showEndLeasePopUp,
                    title: "End Lease".localized(),
                    message: "Are you sure you want to end the current lease?".localized(),
                    buttonTitle: "Confirm".localized(),
                    secondaryButtonTitle: "Cancel".localized(),
                    action: {},
                    secondaryAction: {}
                )
                .accessibilityIdentifier("EndLeaseAlert")
            }
            if showDeletePropertyPopUp {
                CustomAlertTwoButtons(
                    isActive: $showDeletePropertyPopUp,
                    title: "Delete Property".localized(),
                    message: "Are you sure you want to delete this property?".localized(),
                    buttonTitle: "Confirm".localized(),
                    secondaryButtonTitle: "Cancel".localized(),
                    action: {
                        Task {
                            do {
                                try await viewModel.deleteProperty(propertyId: property.id)
                                dismiss()
                            } catch {
                                print("Error deleting property: \(error.localizedDescription)")
                            }
                        }
                    },
                    secondaryAction: {}
                )
                .accessibilityIdentifier("DeletePropertyAlert")
            }
        }
    }

    private func formattedValue(_ value: Double) -> String {
        value == Double(Int(value)) ? String(format: "%.0f", value) : String(format: "%.2f", value)
    }

    private func formatLeaseDates() -> String {
        if property.leaseStartDate == nil && property.leaseEndDate == nil {
            return "No active lease".localized()
        }

        let startDateText: String
        if let startDateString = property.leaseStartDate {
            startDateText = formatDateString(startDateString)
        } else {
            startDateText = "No start date assigned".localized()
        }

        let endDateText: String
        if let endDateString = property.leaseEndDate {
            endDateText = formatDateString(endDateString)
        } else {
            endDateText = "No end date assigned".localized()
        }

        if property.leaseStartDate == nil {
            return startDateText
        }

        return "\(startDateText) - \(endDateText)"
    }

    private func formatDateString(_ dateString: String) -> String {
        let formatter = ISO8601DateFormatter()
        if let date = formatter.date(from: dateString) {
            let displayFormatter = DateFormatter()
            displayFormatter.dateStyle = .medium
            displayFormatter.timeStyle = .none
            return displayFormatter.string(from: date)
        }
        return dateString
    }
}

struct DetailItem: View {
    let icon: String
    let value: String
    let label: String

    var body: some View {
        VStack(spacing: 5) {
            Image(systemName: icon)
                .font(.title3)
                .foregroundColor(Color("LightBlue"))
            Text(value)
                .font(.headline)
                .foregroundColor(Color("textColor"))
            Text(label)
                .font(.caption)
                .foregroundColor(.gray)
                .multilineTextAlignment(.center)
        }
        .padding()
        .background(Color("basicWhiteBlack"))
        .cornerRadius(10)
        .shadow(color: Color.black.opacity(0.1), radius: 10, x: 0, y: 2)
    }
}

struct DocumentsGrid: View {
    @Binding var documents: [PropertyDocument]

    var body: some View {
        LazyVGrid(
            columns: Array(repeating: GridItem(.flexible(), spacing: 15), count: 3),
            spacing: 15
        ) {
            ForEach(documents) { document in
                NavigationLink(destination: PDFViewer(base64String: document.data)) {
                    VStack {
                        Image(systemName: "text.document")
                            .resizable()
                            .scaledToFit()
                            .frame(width: 50, height: 50)

                        Text(document.title)
                            .font(.caption)
                            .multilineTextAlignment(.center)
                            .frame(maxWidth: .infinity)
                    }
                    .padding()
                    .frame(maxWidth: .infinity)
                    .background(Color.gray.opacity(0.1))
                    .cornerRadius(8)
                }
            }
        }
        .padding()
    }
}

struct PropertyDetailView_Previews: PreviewProvider {
    static var previews: some View {
        let property = Property(
            id: "",
            ownerID: "",
            name: "Condo",
            address: "4391 Hedge Street",
            city: "New Jersey",
            postalCode: "07102",
            country: "USA",
            photo: nil,
            monthlyRent: 1200,
            deposit: 2400,
            surface: 80.0,
            isAvailable: "Busy",
            tenantName: "John & Mary Doe",
            leaseStartDate: "2025-04-08T22:00:00Z",
            leaseEndDate: nil,
            documents: [],
            rooms: []
        )
        
        let viewModel = PropertyViewModel()
        
        PropertyDetailView(property: .constant(property), viewModel: viewModel)
            .environmentObject(viewModel)
    }
}
