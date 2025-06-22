//
//  PropertyDetailView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/11/2024.
//

import SwiftUI
import PDFKit

struct PropertyDetailView: View {
    @State private var property: Property
    @ObservedObject var viewModel: PropertyViewModel
    @EnvironmentObject var loginViewModel: LoginViewModel
    @StateObject private var tenantViewModel = TenantViewModel()
    @StateObject private var inventoryViewModel: InventoryViewModel
    @State private var showInviteTenantSheet = false
    @State private var showEndLeasePopUp = false
    @State private var showCancelInvitePopUp = false
    @State private var showDeletePropertyPopUp = false
    @State private var showEditPropertyPopUp = false
    @Binding private var navigateToReportDamage: Bool
    @State private var errorMessage: String?
    @State private var isLoading = false
    @State private var selectedTab: String = "Details".localized()
    @State private var isEntryInventory: Bool = true
    @Binding private var navigateToInventory: Bool
    @State private var rooms: [PropertyRoomsTenant] = []
    @State private var activeLeaseId: String?
    @State private var hasLoaded = false
    @Environment(\.dismiss) var dismiss
    @State private var damageFilter: Bool? = false
    @State private var selectedDamageId: String? = nil
    private let tabs = ["Details".localized(), "Documents".localized(), "Damages".localized()]
    
    init(
        property: Property,
        viewModel: PropertyViewModel,
        navigateToReportDamage: Binding<Bool>,
        navigateToInventory: Binding<Bool>
    ) {
        self._property = State(initialValue: property)
        self.viewModel = viewModel
        self._inventoryViewModel = StateObject(wrappedValue: InventoryViewModel(property: property))
        self._navigateToReportDamage = navigateToReportDamage
        self._navigateToInventory = navigateToInventory
    }

    private func fetchDocuments() async {
        do {
            let fetchedDocuments = try await viewModel.fetchPropertyDocuments(propertyId: property.id)
            await MainActor.run {
                property.documents = fetchedDocuments
            }
            if let index = viewModel.properties.firstIndex(where: { $0.id == property.id }) {
                var updatedProperty = viewModel.properties[index]
                updatedProperty.documents = fetchedDocuments
                viewModel.properties[index] = updatedProperty
            }
        } catch {
            errorMessage = "Error fetching documents: \(error.localizedDescription)".localized()
        }
    }
    
    private func fetchInitialData() async {
        do {
            isLoading = true
            if property.photo == nil || property.damages.isEmpty {
                await viewModel.fetchProperties()
                if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                    property = updatedProperty
                }
            }
            await fetchDocuments()
            if loginViewModel.userRole == "tenant" {
                let token = try await TokenStorage.getValidAccessToken()
                do {
                    rooms = try await viewModel.fetchPropertyRooms(propertyId: property.id, token: token)
                } catch {
                    errorMessage = "Error fetching rooms: \(error.localizedDescription)".localized()
                }
                do {
                    activeLeaseId = try await viewModel.fetchActiveLeaseIdForProperty(propertyId: property.id, token: token)
                } catch {
                    errorMessage = "Error fetching active lease: \(error.localizedDescription)".localized()
                }
                if property.damages.isEmpty, activeLeaseId != nil {
                    do {
                        try await viewModel.fetchPropertyDamages(propertyId: property.id, fixed: damageFilter)
                        if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                            property = updatedProperty
                        }
                    } catch {
                        errorMessage = "Error fetching damages: \(error.localizedDescription)".localized()
                    }
                }
            }
            if loginViewModel.userRole == "owner", let leaseId = property.leaseId {
                if property.damages.isEmpty {
                    do {
                        try await viewModel.fetchPropertyDamages(propertyId: property.id, fixed: damageFilter)
                        if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                            property = updatedProperty
                        }
                    } catch {
                        errorMessage = "Error fetching damages: \(error.localizedDescription)".localized()
                    }
                }
                do {
                    if let lastReport = try await viewModel.fetchLastInventoryReport(propertyId: property.id, leaseId: leaseId) {
                        isEntryInventory = lastReport.type == "end" || lastReport.type == "middle"
                    } else {
                        isEntryInventory = true
                    }
                } catch {
                    errorMessage = "Error fetching inventory report: \(error.localizedDescription)".localized()
                }
            } else {
                isEntryInventory = true
            }
        } catch {
            errorMessage = "Error fetching property data: \(error.localizedDescription)".localized()
        }
        isLoading = false
    }
    
    var body: some View {
        ZStack {
            mainContentView
            errorMessageView
            
            if selectedTab == "Damages".localized() && loginViewModel.userRole == "tenant" {
                VStack {
                    Spacer()
                    HStack {
                        Spacer()
                        Button(action: { navigateToReportDamage = true }) {
                            Image(systemName: "plus")
                                .font(.system(size: 24, weight: .bold))
                                .foregroundColor(.white)
                                .padding()
                                .background(Color("LightBlue"))
                                .clipShape(Circle())
                                .shadow(radius: 4)
                        }
                        .accessibilityLabel("report_damage_btn")
                    }
                }
                .padding(.bottom, 20)
                .padding(.trailing, 20)
            }
        }
        .navigationBarBackButtonHidden(true)
        .onReceive(viewModel.$properties) { properties in
            if let updatedProperty = properties.first(where: { $0.id == property.id }) {
                if !updatedProperty.documents.isEmpty || updatedProperty.isAvailable != property.isAvailable {
                    self.property = updatedProperty
                }
            }
        }
        .onAppear {
            guard !hasLoaded else { return }
            hasLoaded = true
            Task {
                await fetchInitialData()
            }
            inventoryViewModel.onInventoryFinalized = {
                Task {
                    for _ in 1...3 {
                        await fetchDocuments()
                        if !property.documents.isEmpty {
                            break
                        }
                        try? await Task.sleep(nanoseconds: 2_000_000_000)
                    }
                    if property.documents.isEmpty {
                        errorMessage = "No documents found after inventory finalization".localized()
                    }
                }
            }
        }
        .sheet(isPresented: $showEditPropertyPopUp) {
            EditPropertyView(viewModel: viewModel, property: $property)
        }
        .sheet(isPresented: $showInviteTenantSheet) {
            InviteTenantView(tenantViewModel: tenantViewModel, property: property)
        }
        .overlay(alertsView)
        .navigationDestination(isPresented: $navigateToInventory) {
            InventoryRoomView()
                .environmentObject(inventoryViewModel)
        }
        .navigationDestination(isPresented: $navigateToReportDamage) {
            if activeLeaseId != nil {
                ReportDamageView(
                    viewModel: viewModel,
                    propertyId: property.id,
                    rooms: rooms,
                    leaseId: activeLeaseId,
                    onDamageCreated: {
                        Task {
                            try await viewModel.fetchPropertyDamages(propertyId: property.id, fixed: damageFilter)
                            if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                                property = updatedProperty
                            }
                        }
                    }
                )
            }
        }
    }
    
    private var damagesContentView: some View {
        ZStack {
            VStack(spacing: 16) {
                Picker("Filter Damages", selection: $damageFilter) {
                    Text("In Progress".localized()).tag(false)
                    Text("Fixed".localized()).tag(true)
                }
                .pickerStyle(.segmented)
                .padding(.horizontal)
                .onChange(of: damageFilter) {
                    Task {
                        do {
                            try await viewModel.fetchPropertyDamages(propertyId: property.id, fixed: damageFilter)
                            if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                                property = updatedProperty
                            }
                        } catch {
                            errorMessage = "Error fetching damages: \(error.localizedDescription)".localized()
                        }
                    }
                }
                
                ScrollView {
                    VStack(spacing: 16) {
                        if viewModel.isFetchingDamages {
                            ProgressView()
                                .progressViewStyle(.circular)
                                .padding()
                        } else if let damagesError = viewModel.damagesError {
                            Text(damagesError)
                                .foregroundColor(.red)
                                .padding()
                        } else if let currentProperty = viewModel.properties.first(where: { $0.id == property.id }), !currentProperty.damages.isEmpty {
                            LazyVStack(spacing: 10) {
                                ForEach(currentProperty.damages.sorted { $0.createdAt > $1.createdAt }, id: \.id) { damage in
                                    DamageItemView(damage: damage, selectedDamageId: $selectedDamageId)
                                        .id(damage.id)
                                }
                            }
                            .padding(.horizontal)
                        } else {
                            Text("no_damages_reported".localized())
                                .foregroundColor(.gray)
                                .padding()
                        }
                    }
                    .padding(.vertical, 20)
                }
            }
        }
        .navigationDestination(isPresented: Binding(
            get: { selectedDamageId != nil },
            set: { if !$0 { selectedDamageId = nil } }
        )) {
            if let damageId = selectedDamageId,
               let damage = viewModel.properties
                   .first(where: { $0.id == property.id })?
                   .damages
                   .first(where: { $0.id == damageId }) {
                DamageDetailView(damage: damage)
            }
        }
    }

    private var mainContentView: some View {
        VStack(spacing: 0) {
            TopBar(title: "Keyz".localized())
            headerView
            tabsView
            contentScrollView
        }
    }

    private var headerView: some View {
        VStack(alignment: .leading, spacing: 16) {
            ZStack(alignment: .topLeading) {
                imageView
                if loginViewModel.userRole == "owner" {
                    backButtonView
                    optionsMenuView
                }
            }
            propertyInfoView
        }
    }

    private var imageView: some View {
        Group {
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
        }
    }

    private var backButtonView: some View {
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
    }

    private var optionsMenuView: some View {
        Menu {
            if property.isAvailable == "available" {
                Button(action: { showInviteTenantSheet = true }) {
                    Label("Invite Tenant".localized(), systemImage: "person.crop.circle.badge.plus")
                }
            }
            if property.isAvailable == "pending" {
                Button(action: { showCancelInvitePopUp = true }) {
                    Label("Cancel Invite".localized(), systemImage: "person.crop.circle.badge.xmark")
                }
            }
            if property.isAvailable == "unavailable" {
                Button(action: { showEndLeasePopUp = true }) {
                    Label("End Lease".localized(), systemImage: "xmark.circle")
                }
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

    private var propertyInfoView: some View {
        VStack(alignment: .leading, spacing: 4) {
            HStack {
                Text(property.name)
                    .font(.title2)
                    .fontWeight(.bold)
                    .foregroundColor(Color("textColor"))
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
        .padding(.horizontal)
        .padding(.bottom, 8)
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
        case "pending": return Color.orange
        case "unavailable": return Color("RedAlert")
        default: return Color.gray
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

    private var tabsView: some View {
        HStack(spacing: 20) {
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

    private var contentScrollView: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 16) {
                switch selectedTab {
                case "Details".localized():
                    detailsContentView
                    actionButtonView
                case "Documents".localized():
                    if let errorMessage = errorMessage {
                        Text(errorMessage)
                            .foregroundColor(.red)
                            .padding()
                    } else {
                        DocumentsGridView(documents: $property.documents)
                            .padding(.horizontal)
                    }
                case "Damages".localized():
                    damagesContentView
                default:
                    EmptyView()
                }
            }
            .padding(.vertical, 20)
        }
    }

    private var detailsContentView: some View {
        VStack(alignment: .leading, spacing: 16) {
            LazyVGrid(
                columns: Array(repeating: GridItem(.flexible(), spacing: 15), count: 3),
                spacing: 15
            ) {
                DetailItemView(
                    icon: "square.split.bottomrightquarter",
                    value: "\(formattedValue(property.surface))m2",
                    label: "Area".localized()
                )
                DetailItemView(
                    icon: "coloncurrencysign.arrow.trianglehead.counterclockwise.rotate.90",
                    value: "\(property.monthlyRent)$",
                    label: "Rent /month".localized()
                )
                DetailItemView(
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
    }

    private var actionButtonView: some View {
        VStack {
            Group {
                if loginViewModel.userRole == "owner" {
                    Button(action: {
                        inventoryViewModel.isEntryInventory = isEntryInventory
                        navigateToInventory = true
                    }) {
                        Text(isEntryInventory ? "Start Entry Inventory".localized() : "Start Exit Inventory".localized())
                            .frame(maxWidth: .infinity)
                            .padding(.vertical, 15)
                            .background(property.leaseId == nil ? Color.gray : Color("LightBlue"))
                            .foregroundColor(.white)
                            .cornerRadius(10)
                            .padding(.horizontal)
                            .padding(.bottom, 10)
                    }
                    .disabled(property.leaseId == nil)
                    .accessibilityLabel(isEntryInventory ? "inventory_btn_entry" : "inventory_btn_exit")
                }
            }
        }
        .ignoresSafeArea(.container, edges: .bottom)
    }

    private var errorMessageView: some View {
        VStack {
            Spacer()
            Group {
                if let errorMessage = errorMessage {
                    Text(errorMessage)
                        .foregroundColor(.red)
                        .padding()
                }
            }
        }
    }

    private var alertsView: some View {
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
                                }
                            } catch {
                                errorMessage = "Error ending lease: \(error.localizedDescription)".localized()
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

struct PropertyDetailView_Previews: PreviewProvider {
    @State static var navigateToReportDamage: Bool = false
    @State static var navigateToInventory: Bool = false
    
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
            leaseId: "leaseId",
            leaseStartDate: "2025-04-08T22:00:00Z",
            leaseEndDate: nil,
            documents: [
                PropertyDocument(
                    id: "cmboeu0yh000mcxa1b83l5tnv",
                    title: "09-06-2025",
                    fileName: "inventory_report_2025-06-09_cmboeu0wx000hcxa1.pdf",
                    data: ""
                )
            ],
            createdAt: "",
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
        
        let viewModel = PropertyViewModel(loginViewModel: LoginViewModel())
        let loginViewModel = LoginViewModel()
        
        return PropertyDetailView(
            property: property,
            viewModel: viewModel,
            navigateToReportDamage: $navigateToReportDamage,
            navigateToInventory: $navigateToInventory
        )
        .environmentObject(viewModel)
        .environmentObject(loginViewModel)
    }
}
