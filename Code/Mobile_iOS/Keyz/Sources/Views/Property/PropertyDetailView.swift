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
    @EnvironmentObject var loginViewModel: LoginViewModel
    @StateObject private var tenantViewModel = TenantViewModel()
    @StateObject private var inventoryViewModel = InventoryViewModel(property: Property(id: "", ownerID: "", name: "", address: "", city: "", postalCode: "", country: "", photo: nil, monthlyRent: 0, deposit: 0, surface: 0.0, isAvailable: "", tenantName: nil, leaseStartDate: nil, leaseEndDate: nil, documents: [], createdAt: nil, rooms: [], damages: []))
    @State private var showInviteTenantSheet = false
    @State private var showEndLeasePopUp = false
    @State private var showCancelInvitePopUp = false
    @State private var showDeletePropertyPopUp = false
    @State private var showEditPropertyPopUp = false
    @State private var showReportDamageView = false
    @State private var errorMessage: String?
    @State private var isLoading = false
    @State private var selectedTab: String = "Details".localized()
    @State private var isEntryInventory: Bool = true
    @State private var navigateToInventory: Bool = false
    @Environment(\.dismiss) var dismiss
    private let tabs = ["Details".localized(), "Documents".localized(), "Damages".localized()]

    var body: some View {
        ZStack {
            VStack(spacing: 0) {
                TopBar(title: "Keyz".localized())

                VStack(alignment: .leading, spacing: 16) {
                    ZStack(alignment: .topLeading) {
                        if isLoading {
                            VStack {
                                Spacer()
                                ProgressView()
                                    .progressViewStyle(CircularProgressViewStyle())
                                Spacer()
                            }
                            .frame(maxWidth: .infinity, maxHeight: 200)
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

                        if loginViewModel.userRole == "owner" {
                            Menu {
                                Button(action: { showInviteTenantSheet = true }) {
                                    Label("Invite Tenant".localized(), systemImage: "person.crop.circle.badge.plus")
                                }
                                Button(action: { showEndLeasePopUp = true }) {
                                    Label("End Lease".localized(), systemImage: "xmark.circle")
                                }
                                Button(action: { showCancelInvitePopUp = true }) {
                                    Label("Cancel Invite".localized(), systemImage: "person.crop.circle.badge.xmark")
                                }
                                Button(action: { showEditPropertyPopUp = true }) {
                                    Label("Edit Property".localized(), systemImage: "pencil")
                                }
                                Button(action: { showDeletePropertyPopUp = true }) {
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
                                        selectedTab == tab ? Color("LightBlue") : Color.gray.opacity(0.1)
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
                }

                ScrollView {
                    VStack(alignment: .leading, spacing: 16) {
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
                                if viewModel.isFetchingDamages {
                                    ProgressView()
                                        .progressViewStyle(CircularProgressViewStyle())
                                        .padding()
                                } else if let damagesError = viewModel.damagesError {
                                    Text(damagesError)
                                        .foregroundColor(.red)
                                        .padding()
                                } else if property.damages.isEmpty {
                                    Text("no_damages_reported".localized())
                                        .foregroundColor(.gray)
                                        .padding()
                                } else {
                                    LazyVGrid(
                                        columns: [GridItem(.flexible())],
                                        spacing: 10
                                    ) {
                                        ForEach(property.damages) { damage in
                                            DamageItem(damage: damage)
                                        }
                                    }
                                    .padding(.horizontal)
                                }

                                if loginViewModel.userRole == "tenant" {
                                    Button(action: { showReportDamageView = true }) {
                                        Text("Report Damage".localized())
                                            .frame(maxWidth: .infinity)
                                            .padding(.vertical, 15)
                                            .background(Color("LightBlue"))
                                            .foregroundColor(.white)
                                            .cornerRadius(10)
                                            .padding(.horizontal)
                                            .padding(.top, 10)
                                    }
                                    .accessibilityLabel("report_damage_btn")
                                }
                            }
                            .onAppear {
                                print("Tab switched to Damages, fetching damages...")
                                Task {
                                    do {
                                        if loginViewModel.userRole == "tenant" {
                                            if let leaseId = try await viewModel.fetchActiveLeaseIdForProperty(propertyId: property.id, token: try await TokenStorage.getValidAccessToken()) {
                                                try await viewModel.fetchTenantDamages(leaseId: leaseId)
                                                if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                                                    property.damages = updatedProperty.damages
                                                }
                                            } else {
                                                viewModel.damagesError = "No active lease found.".localized()
                                            }
                                        } else {
                                            try await viewModel.fetchPropertyDamages(propertyId: property.id)
                                            if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                                                property = updatedProperty
                                            }
                                        }
                                    } catch {
                                        viewModel.damagesError = "Error fetching damages: \(error.localizedDescription)".localized()
                                        print("Error fetching damages: \(error.localizedDescription)")
                                    }
                                }
                            }

                        default:
                            EmptyView()
                        }
                    }
                    .padding(.vertical, 20)
                }

                if loginViewModel.userRole == "owner" {
                    NavigationLink(
                        destination: InventoryRoomView()
                            .environmentObject(inventoryViewModel),
                        isActive: $navigateToInventory
                    ) {
                        Text(isEntryInventory ? "Start Entry Inventory".localized() : "Start Exit Inventory".localized())
                            .frame(maxWidth: .infinity)
                            .padding(.vertical, 15)
                            .background(Color("LightBlue"))
                            .foregroundColor(.white)
                            .cornerRadius(10)
                            .padding(.horizontal)
                            .padding(.bottom, 10)
                    }
                    .accessibilityLabel(isEntryInventory ? "inventory_btn_entry" : "inventory_btn_exit")
                    .onTapGesture {
                        inventoryViewModel.isEntryInventory = isEntryInventory
                        inventoryViewModel.property = property
                        navigateToInventory = true
                    }
                }

                if let errorMessage = errorMessage {
                    Text(errorMessage)
                        .foregroundColor(.red)
                        .padding()
                }
            }
            .navigationBarBackButtonHidden(true)
            .onAppear {
                Task {
                    do {
                        isLoading = true
                        if property.photo == nil {
                            await viewModel.fetchProperties()
                            if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                                property = updatedProperty
                            }
                        }
                        try await viewModel.fetchPropertyDocuments(propertyId: property.id)
                        if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                            property = updatedProperty
                        }

                        let token = try await TokenStorage.getValidAccessToken()
                        if let leaseId = try await viewModel.fetchActiveLease(propertyId: property.id, token: token) {
                            if let lastReport = try await viewModel.fetchLastInventoryReport(propertyId: property.id, leaseId: leaseId) {
                                isEntryInventory = lastReport.type == "end" || lastReport.type == "middle"
                            } else {
                                isEntryInventory = true
                            }
                        } else {
                            isEntryInventory = true
                        }
                    } catch {
                        errorMessage = "Error fetching property data: \(error.localizedDescription)".localized()
                        print("Error fetching property data: \(error.localizedDescription)")
                    }
                    isLoading = false
                }
            }
            .sheet(isPresented: $showEditPropertyPopUp) {
                EditPropertyView(viewModel: viewModel, property: $property)
            }
            .sheet(isPresented: $showInviteTenantSheet) {
                InviteTenantView(tenantViewModel: tenantViewModel, property: property)
            }
            .sheet(isPresented: $showReportDamageView) {
                ReportDamageView(propertyId: property.id)
            }
            .overlay(
                Group {
                    if showCancelInvitePopUp {
                        CustomAlertTwoButtons(
                            isActive: $showCancelInvitePopUp,
                            title: "Cancel Invite".localized(),
                            message: "Are you sure you want to cancel the pending invite?".localized(),
                            buttonTitle: "Confirm".localized(),
                            secondaryButtonTitle: "Cancel".localized(),
                            action: {
                                Task {
                                    do {
                                        let token = try await TokenStorage.getValidAccessToken()
                                        try await viewModel.cancelInvite(propertyId: property.id, token: token)
                                        await viewModel.fetchProperties()
                                        if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                                            property = updatedProperty
                                        }
                                    } catch {
                                        errorMessage = "Error cancelling invite: \(error.localizedDescription)".localized()
                                        print("Error cancelling invite: \(error.localizedDescription)")
                                    }
                                }
                            },
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
                            action: {
                                Task {
                                    do {
                                        let token = try await TokenStorage.getValidAccessToken()
                                        if let leaseId = try await viewModel.fetchActiveLease(propertyId: property.id, token: token) {
                                            try await viewModel.endLease(propertyId: property.id, leaseId: leaseId, token: token)
                                            await viewModel.fetchProperties()
                                            if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                                                property = updatedProperty
                                            }
                                        } else {
                                            errorMessage = "No active lease found.".localized()
                                            print("No active lease found for property \(property.id)")
                                        }
                                    } catch {
                                        errorMessage = "Error ending lease: \(error.localizedDescription)".localized()
                                        print("Error ending lease: \(error.localizedDescription)")
                                    }
                                }
                            },
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
                                        errorMessage = "Error deleting property: \(error.localizedDescription)".localized()
                                        print("Error deleting property: \(error.localizedDescription)")
                                    }
                                }
                            },
                            secondaryAction: {}
                        )
                        .accessibilityIdentifier("DeletePropertyAlert")
                    }
                }
            )
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

struct ReportDamageView: View {
    let propertyId: String
    @EnvironmentObject var viewModel: PropertyViewModel
    @Environment(\.dismiss) var dismiss
    @State private var comment = ""
    @State private var priority = "low"
    @State private var roomName = ""
    @State private var pictures: [String] = []
    @State private var errorMessage: String?

    var body: some View {
        NavigationView {
            VStack(spacing: 16) {
                Text("Report a Damage")
                    .font(.title)
                    .padding()

                TextField("Room Name", text: $roomName)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .padding(.horizontal)

                TextField("Comment", text: $comment)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .padding(.horizontal)

                Picker("Priority", selection: $priority) {
                    Text("Low").tag("low")
                    Text("Medium").tag("medium")
                    Text("High").tag("high")
                    Text("Urgent").tag("urgent")
                }
                .pickerStyle(MenuPickerStyle())
                .padding(.horizontal)

                Text("Add Photos (not implemented yet)")
                    .foregroundColor(.gray)
                    .padding(.horizontal)

                Button(action: {
                    Task {
                        do {
                            let token = try await TokenStorage.getValidAccessToken()
                            if let leaseId = try await viewModel.fetchActiveLeaseIdForProperty(propertyId: propertyId, token: token) {
                                let damageRequest = DamageRequest(comment: comment, priority: priority, roomName: roomName, pictures: pictures.isEmpty ? nil : pictures)
                                let damageId = try await viewModel.createDamage(propertyId: propertyId, leaseId: leaseId, damage: damageRequest, token: token)
                                print("Damage reported with ID: \(damageId)")
                                dismiss()
                            } else {
                                errorMessage = "No active lease found.".localized()
                            }
                        } catch {
                            errorMessage = "Error reporting damage: \(error.localizedDescription)".localized()
                            print("Error reporting damage: \(error.localizedDescription)")
                        }
                    }
                }) {
                    Text("Submit")
                        .frame(maxWidth: .infinity)
                        .padding()
                        .background(Color("LightBlue"))
                        .foregroundColor(.white)
                        .cornerRadius(10)
                }
                .padding(.horizontal)

                if let errorMessage = errorMessage {
                    Text(errorMessage)
                        .foregroundColor(.red)
                        .padding()
                }

                Spacer()
            }
            .navigationBarItems(trailing: Button("Cancel") {
                dismiss()
            })
        }
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

struct DamageItem: View {
    let damage: DamageResponse

    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            HStack {
                Text(damage.roomName)
                    .font(.headline)
                    .foregroundColor(Color("textColor"))
                Spacer()
                Text(damage.priority.capitalized)
                    .font(.caption)
                    .fontWeight(.medium)
                    .foregroundColor(.white)
                    .padding(.horizontal, 8)
                    .padding(.vertical, 4)
                    .background(
                        RoundedRectangle(cornerRadius: 8)
                            .fill(priorityColor(damage.priority))
                    )
            }
            Text(damage.comment)
                .font(.subheadline)
                .foregroundColor(.gray)
            HStack {
                Text("Status: \(damage.fixStatus.replacingOccurrences(of: "_", with: " ").capitalized)")
                    .font(.caption)
                    .foregroundColor(damage.fixStatus == "fixed" ? Color("GreenAlert") : Color("RedAlert"))
                Spacer()
                Text(formatDateString(damage.createdAt))
                    .font(.caption)
                    .foregroundColor(.gray)
            }
        }
        .padding()
        .background(Color("basicWhiteBlack"))
        .cornerRadius(10)
        .shadow(color: Color.black.opacity(0.1), radius: 5, x: 0, y: 2)
    }

    private func priorityColor(_ priority: String) -> Color {
        switch priority.lowercased() {
        case "low":
            return Color("GreenAlert")
        case "medium":
            return .yellow
        case "high":
            return .orange
        case "urgent":
            return Color("RedAlert")
        default:
            return .gray
        }
    }

    private func formatDateString(_ dateString: String) -> String {
        let formatter = ISO8601DateFormatter()
        formatter.formatOptions = [.withInternetDateTime, .withFractionalSeconds]
        if let date = formatter.date(from: dateString) {
            let displayFormatter = DateFormatter()
            displayFormatter.dateFormat = "dd/MM/yyyy"
            displayFormatter.locale = Locale(identifier: "en_GB")
            let formattedDate = displayFormatter.string(from: date)
            return formattedDate
        }
        return dateString
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
            rooms: [],
            damages: [
                DamageResponse(
                    id: "damage_001",
                    comment: "Cracked window in the living room",
                    priority: "high",
                    roomName: "Living Room",
                    fixStatus: "pending",
                    pictures: ["base64_image_1"],
                    createdAt: "2025-05-10T09:00:00Z",
                    updatedAt: nil,
                    fixPlannedAt: "2025-05-25T14:00:00Z",
                    fixedAt: nil,
                    leaseId: "lease_001",
                    propertyId: "",
                    propertyName: "Condo",
                    tenantName: "John & Mary Doe",
                    read: true
                ),
                DamageResponse(
                    id: "damage_002",
                    comment: "Leaking faucet in the kitchen",
                    priority: "medium",
                    roomName: "Kitchen",
                    fixStatus: "fixed",
                    pictures: ["base64_image_2"],
                    createdAt: "2025-05-12T11:00:00Z",
                    updatedAt: "2025-05-18T15:00:00Z",
                    fixPlannedAt: nil,
                    fixedAt: "2025-05-18T15:00:00Z",
                    leaseId: "lease_001",
                    propertyId: "",
                    propertyName: "Condo",
                    tenantName: "John & Mary Doe",
                    read: false
                ),
                DamageResponse(
                    id: "damage_003",
                    comment: "Scratched floor in bedroom",
                    priority: "low",
                    roomName: "Bedroom",
                    fixStatus: "pending",
                    pictures: [],
                    createdAt: "2025-05-15T08:00:00Z",
                    updatedAt: nil,
                    fixPlannedAt: nil,
                    fixedAt: nil,
                    leaseId: "lease_001",
                    propertyId: "",
                    propertyName: "Condo",
                    tenantName: "John & Mary Doe",
                    read: true
                ),
                DamageResponse(
                    id: "damage_004",
                    comment: "Scratched floor in bedroom",
                    priority: "low",
                    roomName: "Bedroom",
                    fixStatus: "pending",
                    pictures: [],
                    createdAt: "2025-05-15T08:00:00Z",
                    updatedAt: nil,
                    fixPlannedAt: nil,
                    fixedAt: nil,
                    leaseId: "lease_001",
                    propertyId: "",
                    propertyName: "Condo",
                    tenantName: "John & Mary Doe",
                    read: true
                )
            ]
        )
        
        let viewModel = PropertyViewModel()
        let loginViewModel = LoginViewModel()
        
        return PropertyDetailView(property: .constant(property), viewModel: viewModel)
            .environmentObject(viewModel)
            .environmentObject(loginViewModel)
    }
}
