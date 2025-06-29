//
//  PropertyDetailView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/11/2024.
//

import SwiftUI
import PDFKit
import UniformTypeIdentifiers
import MobileCoreServices

struct PropertyDetailView: View {
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
    @State private var showError: Bool = false
    @State private var errorMessage: String?
    @State private var isLoading = false
    @State private var isImageLoading = false
    @State private var selectedTab: String = "Details".localized()
    @State private var isEntryInventory: Bool = true
    @Binding private var navigateToInventory: Bool
    @State private var rooms: [PropertyRoomsTenant] = []
    @State private var activeLeaseId: String?
    @State private var hasLoaded = false
    @Environment(\.dismiss) var dismiss
    @State private var damageFilter: Bool? = false
    @State private var selectedDamageId: String?
    @State private var selectedDamage: DamageResponse?
    @State private var showDocumentPicker = false
    private let tabs = ["Details".localized(), "Documents".localized(), "Damages".localized()]
    let propertyId: String
    private let debouncer = Debouncer(delay: 0.5)
    @State private var showDeleteDocumentPopUp = false
    @State private var documentToDeleteId: String?
    @State private var decodedDocuments: [String: PDFDocument] = [:]

    init(
        propertyId: String,
        viewModel: PropertyViewModel,
        navigateToReportDamage: Binding<Bool>,
        navigateToInventory: Binding<Bool>
    ) {
        self.propertyId = propertyId
        self.viewModel = viewModel
        self._inventoryViewModel = StateObject(wrappedValue: InventoryViewModel(property: viewModel.properties.first(where: { $0.id == propertyId }) ?? Property(id: "", ownerID: "", name: "", address: "", city: "", postalCode: "", country: "", photo: nil, monthlyRent: 0, deposit: 0, surface: 0, isAvailable: "", tenantName: nil, leaseId: nil, leaseStartDate: nil, leaseEndDate: nil, documents: [], createdAt: "", rooms: [], damages: [])))
        self._navigateToReportDamage = navigateToReportDamage
        self._navigateToInventory = navigateToInventory
    }

    private var property: Property? {
        viewModel.properties.first(where: { $0.id == propertyId })
    }

    private var hasActiveLease: Bool {
        if loginViewModel.userRole == "tenant" {
            return activeLeaseId != nil
        } else if loginViewModel.userRole == "owner" {
            return property?.leaseId != nil
        }
        return false
    }

    private func fetchDocuments(forceRefresh: Bool = false) async {
        guard let property = property, property.leaseId != nil else {
            return
        }
        do {
            _ = try await viewModel.fetchPropertyDocuments(propertyId: propertyId, forceRefresh: forceRefresh)
        } catch {
            errorMessage = "Error fetching documents: \(error.localizedDescription)".localized()
            showError = true
        }
    }

    private func fetchInitialData() async {
        guard let property = property else {
            return
        }
        isLoading = true
        defer { isLoading = false }
        
        do {
            if property.photo == nil {
                isImageLoading = true
                do {
                    let image = try await viewModel.fetchPropertiesPicture(propertyId: propertyId)
                    if let index = viewModel.properties.firstIndex(where: { $0.id == propertyId }) {
                        var updatedProperty = viewModel.properties[index]
                        updatedProperty.photo = image
                        viewModel.properties[index] = updatedProperty
                    }
                } catch {
                    errorMessage = "Error fetching image: \(error.localizedDescription)".localized()
                    showError = true
                }
                isImageLoading = false
            }

            if loginViewModel.userRole == "tenant" {
                let token = try await TokenStorage.getValidAccessToken()
                if rooms.isEmpty {
                    do {
                        rooms = try await viewModel.fetchPropertyRooms(propertyId: propertyId, token: token)
                    } catch {
                        errorMessage = "Error fetching rooms: \(error.localizedDescription)".localized()
                        showError = true
                    }
                }
                if activeLeaseId == nil {
                    do {
                        activeLeaseId = try await viewModel.fetchActiveLeaseIdForProperty(propertyId: propertyId, token: token)
                        if activeLeaseId == nil {
                            errorMessage = "No active lease found for this property.".localized()
                            showError = true
                        }
                    } catch {
                        errorMessage = "Error fetching lease ID: \(error.localizedDescription)".localized()
                        showError = true
                    }
                }
                if activeLeaseId != nil {
                    if property.damages.isEmpty {
                        try await viewModel.fetchPropertyDamages(propertyId: propertyId, fixed: damageFilter)
                    }
                    if property.documents.isEmpty {
                        await fetchDocuments()
                    }
                }
            } else if loginViewModel.userRole == "owner" {
                if let leaseId = property.leaseId {
                    if property.damages.isEmpty {
                        try await viewModel.fetchPropertyDamages(propertyId: propertyId, fixed: damageFilter)
                    }
                    if property.documents.isEmpty {
                        await fetchDocuments()
                    }
                    do {
                        if let lastReport = try await viewModel.fetchLastInventoryReport(propertyId: propertyId, leaseId: leaseId) {
                            isEntryInventory = lastReport.type == "end" || lastReport.type == "middle"
                        } else {
                            isEntryInventory = true
                        }
                    } catch {
                        errorMessage = "Error fetching inventory report: \(error.localizedDescription)".localized()
                        showError = true
                    }
                }
            }
        } catch {
            errorMessage = "Error fetching property data: \(error.localizedDescription)".localized()
            showError = true
        }
    }

    var body: some View {
        ZStack {
            if let property = property {
                mainContentView(property: property)

                if selectedTab == "Damages".localized() && loginViewModel.userRole == "tenant" && hasActiveLease && !rooms.isEmpty && activeLeaseId != nil {
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

                if selectedTab == "Documents".localized() && (loginViewModel.userRole == "tenant" || loginViewModel.userRole == "owner") && hasActiveLease {
                    VStack {
                        Spacer()
                        HStack {
                            Spacer()
                            Button(action: { showDocumentPicker = true }) {
                                Image(systemName: "plus")
                                    .font(.system(size: 24, weight: .bold))
                                    .foregroundColor(.white)
                                    .padding()
                                    .background(Color("LightBlue"))
                                    .clipShape(Circle())
                                    .shadow(radius: 4)
                            }
                            .accessibilityLabel("upload_document_btn")
                        }
                    }
                    .padding(.bottom, 20)
                    .padding(.trailing, 20)
                }

                if showError, let message = errorMessage {
                    ErrorNotificationView(message: message)
                        .onDisappear {
                            showError = false
                            errorMessage = nil
                        }
                }
            } else {
                if isLoading {
                    ProgressView()
                        .progressViewStyle(.circular)
                        .scaleEffect(1.5)
                } else {
                    Text("Property not found".localized())
                        .foregroundColor(.red)
                        .padding()
                }
            }
        }
        .navigationBarBackButtonHidden(true)
        .onAppear {
            guard !hasLoaded else { return }
            hasLoaded = true
            Task {
                if property == nil {
                    isLoading = true
                    do {
                        _ = try await viewModel.fetchPropertyById(propertyId)
                    } catch {
                        errorMessage = "Error fetching property: \(error.localizedDescription)".localized()
                        showError = true
                    }
                    isLoading = false
                }
                await fetchInitialData()
            }
            inventoryViewModel.onDocumentsRefreshNeeded = {
                Task {
                    await self.fetchDocuments(forceRefresh: true)
                    if let property = self.property, let leaseId = property.leaseId {
                        do {
                            if let lastReport = try await self.viewModel.fetchLastInventoryReport(propertyId: property.id, leaseId: leaseId) {
                                self.isEntryInventory = lastReport.type == "end" || lastReport.type == "middle"
                            } else {
                                self.isEntryInventory = true
                            }
                        } catch {
                            self.errorMessage = "Error fetching inventory report: \(error.localizedDescription)".localized()
                            self.showError = true
                        }
                    }
                }
            }
        }
        .sheet(isPresented: $showEditPropertyPopUp) {
            if let property = property {
                EditPropertyView(viewModel: viewModel, property: .constant(property))
            }
        }
        .sheet(isPresented: $showInviteTenantSheet) {
            if let property = property {
                InviteTenantView(
                    tenantViewModel: tenantViewModel,
                    propertyViewModel: viewModel,
                    property: property
                )
            }
        }
        .sheet(isPresented: $showDocumentPicker) {
            DocumentPicker { url in
                Task {
                    await uploadDocument(url: url)
                }
            }
        }
        .overlay(alertsView)
        .navigationDestination(isPresented: $navigateToInventory) {
            InventoryRoomView()
                .environmentObject(inventoryViewModel)
        }
        .navigationDestination(isPresented: Binding(
            get: { selectedDamage != nil },
            set: { if !$0 {
                selectedDamage = nil
                selectedDamageId = nil
            } }
        )) {
            if let damage = selectedDamage {
                DamageDetailView(damage: damage)
                    .environmentObject(viewModel)
                    .environmentObject(loginViewModel)
            } else {
                Text("Damage not found".localized())
                    .foregroundColor(.red)
                    .padding()
            }
        }
        .navigationDestination(isPresented: $navigateToReportDamage) {
            if let _ = property, let leaseId = activeLeaseId {
                ReportDamageView(
                    viewModel: viewModel,
                    propertyId: propertyId,
                    rooms: rooms,
                    leaseId: leaseId,
                    onDamageCreated: {
                        Task {
                            try await viewModel.fetchPropertyDamages(propertyId: propertyId, fixed: damageFilter)
                        }
                    }
                )
            }
        }
        .onChange(of: selectedTab) {
            if selectedTab == "Documents".localized() {
                Task {
                    await fetchDocuments()
                }
            }
        }
    }

    private func uploadDocument(url: URL) async {
        do {
            guard let propertyId = viewModel.properties.first(where: { $0.id == self.propertyId })?.id else {
                throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "No property selected.".localized()])
            }

            let allowedTypes: [UTType] = [.pdf, .init(filenameExtension: "docx")!, .init(filenameExtension: "xlsx")!]
            guard let fileType = UTType(filenameExtension: url.pathExtension.lowercased()),
                  allowedTypes.contains(fileType) else {
                throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid file type. Only PDF, DOCX, and XLSX are supported.".localized()])
            }

            let mimeType: String
            switch fileType {
            case .pdf:
                mimeType = "application/pdf"
            case UTType(filenameExtension: "docx"):
                mimeType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
            case UTType(filenameExtension: "xlsx"):
                mimeType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
            default:
                mimeType = "application/octet-stream"
            }

            guard url.startAccessingSecurityScopedResource() else {
                throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Unable to access file.".localized()])
            }
            defer { url.stopAccessingSecurityScopedResource() }

            let data = try Data(contentsOf: url)
            let base64String = "data:\(mimeType);base64,\(data.base64EncodedString())"
            let fileName = url.lastPathComponent

            try await viewModel.uploadDocument(
                propertyId: propertyId,
                fileName: fileName,
                base64Data: base64String
            )
            await fetchDocuments(forceRefresh: true)
        } catch {
            errorMessage = "Error uploading document: \(error.localizedDescription)".localized()
            showError = true
        }
        await MainActor.run {
            showDocumentPicker = false
        }
    }

    private var damagesContentView: some View {
        ZStack {
            VStack(spacing: 16) {
                if hasActiveLease {
                    Picker("Filter Damages", selection: $damageFilter) {
                        Text("In Progress".localized()).tag(false as Bool?)
                        Text("Fixed".localized()).tag(true as Bool?)
                    }
                    .pickerStyle(.segmented)
                    .padding(.horizontal)
                    .onChange(of: damageFilter) {
                        Task {
                            do {
                                try await viewModel.fetchPropertyDamages(propertyId: propertyId, fixed: damageFilter)
                            } catch {
                                errorMessage = "Error fetching damages: \(error.localizedDescription)".localized()
                                showError = true
                            }
                        }
                    }
                }
                
                ScrollView {
                    VStack(spacing: 16) {
                        if !hasActiveLease {
                            Text("No lease is currently active".localized())
                                .foregroundColor(.gray)
                                .padding()
                        } else if viewModel.isFetchingDamages {
                            ProgressView()
                                .progressViewStyle(.circular)
                                .padding()
                        } else if let damagesError = viewModel.damagesError {
                            Text(damagesError)
                                .foregroundColor(.red)
                                .padding()
                        } else if let property = property, !property.damages.isEmpty {
                            LazyVStack(spacing: 10) {
                                ForEach(property.damages.sorted { $0.createdAt > $1.createdAt }, id: \.id) { damage in
                                    Button(action: {
                                        selectedDamageId = damage.id
                                        selectedDamage = damage
                                    }) {
                                        DamageItemView(damage: damage, selectedDamageId: $selectedDamageId)
                                            .id(damage.id)
                                    }
                                }
                            }
                            .padding(.horizontal)
                        } else {
                            Text(rooms.isEmpty ? "No rooms available for reporting damages".localized() : "No damages reported".localized())
                                .foregroundColor(.gray)
                                .padding()
                        }
                    }
                    .padding(.vertical, 20)
                }
            }
        }
    }
    
    private func mainContentView(property: Property) -> some View {
        VStack(spacing: 0) {
            TopBar(title: "Keyz".localized())
            headerView(property: property)
            tabsView
            contentScrollView(property: property)
        }
    }

    private func headerView(property: Property) -> some View {
        VStack(alignment: .leading, spacing: 16) {
            ZStack(alignment: .topLeading) {
                imageView(property: property)
                if loginViewModel.userRole == "owner" {
                    backButtonView
                    optionsMenuView(property: property)
                }
            }
            propertyInfoView(property: property)
        }
    }

    private func imageView(property: Property) -> some View {
        Group {
            if isImageLoading {
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

    private func optionsMenuView(property: Property) -> some View {
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

    private func propertyInfoView(property: Property) -> some View {
        VStack(alignment: .leading, spacing: 4) {
            HStack {
                Text(property.name)
                    .font(.title2)
                    .fontWeight(.bold)
                    .foregroundColor(Color("textColor"))
                Text(statusText(property: property))
                    .font(.caption)
                    .fontWeight(.medium)
                    .foregroundColor(.white)
                    .padding(.horizontal, 8)
                    .padding(.vertical, 4)
                    .background(
                        RoundedRectangle(cornerRadius: 8)
                            .fill(statusColor(property: property))
                    )
                    .accessibilityLabel(statusAccessibilityLabel(property: property))
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

    private func statusText(property: Property) -> String {
        switch property.isAvailable {
        case "available": return "Available".localized()
        case "pending": return "Pending".localized()
        case "unavailable": return "Unavailable".localized()
        default: return "Unknown".localized()
        }
    }

    private func statusColor(property: Property) -> Color {
        switch property.isAvailable {
        case "available": return Color("GreenAlert")
        case "pending": return Color.orange
        case "unavailable": return Color("RedAlert")
        default: return Color.gray
        }
    }

    private func statusAccessibilityLabel(property: Property) -> String {
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

    private func contentScrollView(property: Property) -> some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 16) {
                switch selectedTab {
                case "Details".localized():
                    detailsContentView(property: property)
                    actionButtonView(property: property)
                case "Documents".localized():
                    if !hasActiveLease {
                        Text("No lease is currently active".localized())
                            .foregroundColor(.gray)
                            .padding()
                    } else {
                        DocumentsGridView(
                            documents: .constant(property.documents),
                            onDelete: { documentId in
                                documentToDeleteId = documentId
                                showDeleteDocumentPopUp = true
                            }
                        )
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

    private func detailsContentView(property: Property) -> some View {
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
                    value: "\(property.monthlyRent)€",
                    label: "Rent /month".localized()
                )
                DetailItemView(
                    icon: "eurosign.bank.building",
                    value: "\(property.deposit)€",
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
                Text(formatLeaseDates(property: property))
                    .foregroundColor(.gray)
            }
            .padding(.horizontal)
        }
    }

    private func actionButtonView(property: Property) -> some View {
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
                                try await viewModel.cancelInvite(propertyId: propertyId, token: token)
                                try await viewModel.fetchProperties()
                            } catch {
                                errorMessage = "Error cancelling invite: \(error.localizedDescription)".localized()
                                showError = true
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
                                if let leaseId = try await viewModel.fetchActiveLease(propertyId: propertyId, token: token) {
                                    try await viewModel.endLease(propertyId: propertyId, leaseId: leaseId, token: token)
                                    try await viewModel.fetchProperties()
                                } else {
                                    errorMessage = "No active lease found.".localized()
                                    showError = true
                                }
                            } catch {
                                errorMessage = "Error ending lease: \(error.localizedDescription)".localized()
                                showError = true
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
                                try await viewModel.deleteProperty(propertyId: propertyId)
                                dismiss()
                            } catch {
                                errorMessage = "Error deleting property: \(error.localizedDescription)".localized()
                                showError = true
                            }
                        }
                    },
                    secondaryAction: {}
                )
                .accessibilityIdentifier("DeletePropertyAlert")
            }
            if showDeleteDocumentPopUp {
                CustomAlertTwoButtons(
                    isActive: $showDeleteDocumentPopUp,
                    title: "Delete Document".localized(),
                    message: "Are you sure you want to delete this document?".localized(),
                    buttonTitle: "Confirm".localized(),
                    secondaryButtonTitle: "Cancel".localized(),
                    action: {
                        Task {
                            do {
                                if let docId = documentToDeleteId {
                                    try await viewModel.deleteDocument(propertyId: propertyId, documentId: docId)
                                    DocumentCache.shared.removeDocument(forKey: docId)
                                    await fetchDocuments(forceRefresh: true)
                                }
                            } catch {
                                errorMessage = "Error deleting document: \(error.localizedDescription)".localized()
                                showError = true
                            }
                            await MainActor.run {
                                documentToDeleteId = nil
                                showDeleteDocumentPopUp = false
                            }
                        }
                    },
                    secondaryAction: {
                        documentToDeleteId = nil
                        showDeleteDocumentPopUp = false
                    }
                )
                .accessibilityIdentifier("DeleteDocumentAlert")
            }
        }
    }

    private func formattedValue(_ value: Double) -> String {
        value == Double(Int(value)) ? String(format: "%.0f", value) : String(format: "%.2f", value)
    }

    private func formatLeaseDates(property: Property) -> String {
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
            id: "property_001",
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
                    propertyId: "property_001",
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
                    propertyId: "property_001",
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
                    propertyId: "property_001",
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
                    propertyId: "property_001",
                    propertyName: "Condo",
                    tenantName: "John & Mary Doe",
                    read: true
                )
            ]
        )
        
        let viewModel = PropertyViewModel(loginViewModel: LoginViewModel())
        let loginViewModel = LoginViewModel()
        viewModel.properties = [property]
        
        return PropertyDetailView(
            propertyId: "property_001",
            viewModel: viewModel,
            navigateToReportDamage: $navigateToReportDamage,
            navigateToInventory: $navigateToInventory
        )
        .environmentObject(viewModel)
        .environmentObject(loginViewModel)
    }
}
