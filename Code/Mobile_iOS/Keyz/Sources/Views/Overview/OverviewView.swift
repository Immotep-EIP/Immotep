//
//  OverviewView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 21/10/2024.
//

import SwiftUI

struct OverviewView: View {
    @EnvironmentObject private var loginViewModel: LoginViewModel
    @StateObject private var viewModel: OverviewViewModel
    @StateObject private var propertyViewModel: PropertyViewModel
    @AppStorage("isLoggedIn") var isLoggedIn: Bool = false
    @State private var navigateToProperty: String? = nil
    @State private var navigateToDamage: (propertyId: String, damageId: String)? = nil
    @State private var selectedDamage: DamageResponse? = nil
    @State private var showError = false
    @State private var errorMessage: String?
    @State private var navigateToReportDamage: Bool = false
    @State private var navigateToInventory: Bool = false
    @State private var isTenantLoading: Bool = false

    init() {
        self._viewModel = StateObject(wrappedValue: OverviewViewModel())
        self._propertyViewModel = StateObject(wrappedValue: PropertyViewModel(loginViewModel: LoginViewModel()))
    }

    var body: some View {
        NavigationStack {
            ZStack {
                VStack(spacing: 0) {
                    TopBar(title: "Keyz")
                    ScrollView {
                        if showError, let _ = errorMessage {
                            EmptyView()
                        } else if !(viewModel.isLoading || propertyViewModel.isFetchingDamages || isTenantLoading) {
                            contentView
                        }
                    }
                }
                if viewModel.isLoading || propertyViewModel.isFetchingDamages || isTenantLoading {
                    ProgressView()
                        .frame(maxWidth: .infinity, maxHeight: .infinity)
                }
                if showError, let message = errorMessage {
                    ErrorNotificationView(message: message)
                        .onDisappear {
                            showError = false
                            errorMessage = nil
                        }
                }
            }
            .navigationBarBackButtonHidden(true)
            .navigationDestination(isPresented: Binding(
                get: { navigateToProperty != nil },
                set: { if !$0 { navigateToProperty = nil } }
            )) {
                if let propertyId = navigateToProperty,
                   let _ = propertyViewModel.properties.first(where: { $0.id == propertyId }) {
                    PropertyDetailView(
                        propertyId: propertyId,
                        viewModel: propertyViewModel,
                        navigateToReportDamage: $navigateToReportDamage,
                        navigateToInventory: $navigateToInventory
                    )
                    .environmentObject(loginViewModel)
                } else {
                    Text("Property not found (ID: \(navigateToProperty ?? "nil"))".localized())
                        .foregroundColor(.red)
                        .padding()
                        .onAppear {
                            print("Property not found for ID: \(navigateToProperty ?? "nil")")
                            print("Available properties: \(propertyViewModel.properties.map { $0.id })")
                        }
                }
            }
            .navigationDestination(isPresented: Binding(
                get: { navigateToDamage != nil },
                set: { if !$0 { navigateToDamage = nil; selectedDamage = nil } }
            )) {
                if let damage = selectedDamage {
                    DamageDetailView(damage: damage)
                        .environmentObject(propertyViewModel)
                        .environmentObject(loginViewModel)
                } else {
                    Text("Damage not found".localized())
                        .foregroundColor(.red)
                        .padding()
                }
            }
            .onAppear {
                loginViewModel.loadUser()
                Task {
                    if loginViewModel.userRole == "owner" {
                        if viewModel.dashboardData == nil {
                            await viewModel.fetchDashboardData()
                            if let dashboardData = viewModel.dashboardData {
                                isTenantLoading = true
                                await fetchOwnerProperties(dashboardData: dashboardData)
                                isTenantLoading = false
                            } else {
                                errorMessage = "No dashboard data available".localized()
                                showError = true
                            }
                        }
                    } else if loginViewModel.userRole == "tenant" {
                        if propertyViewModel.properties.isEmpty {
                            isTenantLoading = true
                            await fetchTenantData()
                            isTenantLoading = false
                        }
                    }
                }
            }
            .onChange(of: viewModel.errorMessage) {
                if let error = viewModel.errorMessage {
                    errorMessage = error
                    showError = true
                } else {
                    showError = false
                    errorMessage = nil
                }
            }
            .onChange(of: propertyViewModel.damagesError) {
                if let error = propertyViewModel.damagesError {
                    errorMessage = error
                    showError = true
                } else {
                    showError = false
                    errorMessage = nil
                }
            }
        }
    }

    private func fetchOwnerProperties(dashboardData: DashboardResponse) async {
        let allProperties = dashboardData.properties.recentlyAdded ?? []
        
        for propertySummary in allProperties {
            let propertyId = propertySummary.id
            do {
                if let property = try await propertyViewModel.fetchPropertyById(propertyId) {
                    if property.leaseId != nil {
                        try await propertyViewModel.fetchPropertyDamages(propertyId: propertyId, fixed: false)
                    }
                }
            } catch {
                print("Failed to fetch property \(propertyId): \(error.localizedDescription)")
            }
        }
    }

    private func fetchTenantData() async {
        do {
            try await propertyViewModel.fetchProperties()
            if let propertyId = propertyViewModel.properties.first?.id {
                try await propertyViewModel.fetchPropertyDamages(propertyId: propertyId, fixed: false)
                let token = try await TokenStorage.getValidAccessToken()
                let rooms = try await propertyViewModel.fetchPropertyRooms(propertyId: propertyId, token: token)
                if let index = propertyViewModel.properties.firstIndex(where: { $0.id == propertyId }) {
                    var updatedProperty = propertyViewModel.properties[index]
                    updatedProperty.rooms = rooms.map { PropertyRooms(id: $0.id, name: $0.name, checked: false, inventory: []) }
                    propertyViewModel.properties[index] = updatedProperty
                }
                if propertyViewModel.properties.first?.leaseId != nil {
                    _ = try await propertyViewModel.fetchPropertyDocuments(propertyId: propertyId)
                }
            }
        } catch {
            errorMessage = "Error fetching tenant data: \(error.localizedDescription)".localized()
            showError = true
        }
    }

    @ViewBuilder
    private var contentView: some View {
        if loginViewModel.userRole == "owner", let dashboardData = viewModel.dashboardData {
            OwnerOverview(
                dashboardData: dashboardData,
                firstName: loginViewModel.user?.firstname ?? "User",
                navigateToProperty: $navigateToProperty,
                navigateToDamage: $navigateToDamage,
                onDamageSelected: { propertyId, damageId in
                    Task {
                        do {
                            let token = try await TokenStorage.getValidAccessToken()
                            let damage = try await propertyViewModel.fetchDamageByID(propertyId: propertyId, damageId: damageId, token: token)
                            if let index = propertyViewModel.properties.firstIndex(where: { $0.id == propertyId }) {
                                var updatedProperty = propertyViewModel.properties[index]
                                if let damageIndex = updatedProperty.damages.firstIndex(where: { $0.id == damageId }) {
                                    updatedProperty.damages[damageIndex] = damage
                                } else {
                                    updatedProperty.damages.append(damage)
                                }
                                propertyViewModel.properties[index] = updatedProperty
                                propertyViewModel.objectWillChange.send()
                            }
                            self.selectedDamage = damage
                            self.navigateToDamage = (propertyId, damageId)
                        } catch {
                            errorMessage = "Error fetching damage details: \(error.localizedDescription)".localized()
                            showError = true
                        }
                    }
                }
            )
        } else if loginViewModel.userRole == "tenant", let property = propertyViewModel.properties.first {
            TenantOverview(
                property: property,
                firstName: loginViewModel.user?.firstname ?? "User",
                navigateToProperty: $navigateToProperty,
                navigateToDamage: $navigateToDamage,
                onDamageSelected: { propertyId, damageId in
                    Task {
                        do {
                            let token = try await TokenStorage.getValidAccessToken()
                            let damage = try await propertyViewModel.fetchDamageByID(propertyId: propertyId, damageId: damageId, token: token)
                            if let index = propertyViewModel.properties.firstIndex(where: { $0.id == propertyId }) {
                                var updatedProperty = propertyViewModel.properties[index]
                                if let damageIndex = updatedProperty.damages.firstIndex(where: { $0.id == damageId }) {
                                    updatedProperty.damages[damageIndex] = damage
                                } else {
                                    updatedProperty.damages.append(damage)
                                }
                                propertyViewModel.properties[index] = updatedProperty
                                propertyViewModel.objectWillChange.send()
                            }
                            self.selectedDamage = damage
                            self.navigateToDamage = (propertyId, damageId)
                        } catch {
                            errorMessage = "Error fetching damage details: \(error.localizedDescription)".localized()
                            showError = true
                        }
                    }
                }
            )
        } else {
            Text("No data available".localized())
                .foregroundColor(.gray)
                .padding()
        }
    }
}

struct OwnerOverview: View {
    let dashboardData: DashboardResponse
    let firstName: String
    @Binding var navigateToProperty: String?
    @Binding var navigateToDamage: (propertyId: String, damageId: String)?
    let onDamageSelected: (String, String) -> Void

    var body: some View {
        LazyVGrid(
            columns: [GridItem(.flexible())],
            spacing: 16
        ) {
            WelcomeWidget(firstName: firstName)
            RemindersWidget(
                reminders: dashboardData.reminders,
                properties: dashboardData.properties,
                navigateToProperty: $navigateToProperty
            )
            PropertiesWidget(
                stats: dashboardData.properties,
                navigateToProperty: $navigateToProperty
            )
            DamagesWidget(
                stats: dashboardData.openDamages,
                navigateToDamage: $navigateToDamage,
                onDamageSelected: onDamageSelected
            )
        }
        .padding()
    }
}

struct TenantOverview: View {
    let property: Property
    let firstName: String
    @Binding var navigateToProperty: String?
    @Binding var navigateToDamage: (propertyId: String, damageId: String)?
    let onDamageSelected: (String, String) -> Void

    private var damageStats: OpenDamageStats {
        let damages = property.damages.map { damage in
            DamageSummary(
                id: damage.id,
                leaseId: damage.leaseId,
                tenantName: damage.tenantName ?? "",
                propertyId: damage.propertyId,
                propertyName: damage.propertyName,
                roomId: "",
                roomName: damage.roomName,
                comment: damage.comment,
                priority: damage.priority,
                read: damage.read,
                createdAt: damage.createdAt,
                updatedAt: damage.updatedAt ?? "",
                fixStatus: damage.fixStatus,
                fixPlannedAt: damage.fixPlannedAt
            )
        }
        return OpenDamageStats(
            total: damages.count,
            urgent: damages.filter { $0.priority.lowercased() == "urgent" }.count,
            high: damages.filter { $0.priority.lowercased() == "high" }.count,
            medium: damages.filter { $0.priority.lowercased() == "medium" }.count,
            low: damages.filter { $0.priority.lowercased() == "low" }.count,
            plannedThisWeek: damages.filter { $0.fixPlannedAt != nil }.count,
            toFix: damages.filter { $0.fixStatus == "pending" }
        )
    }

    var body: some View {
        LazyVGrid(
            columns: [GridItem(.flexible())],
            spacing: 16
        ) {
            WelcomeWidget(firstName: firstName)
            TenantPropertyWidget(
                property: property,
                navigateToProperty: $navigateToProperty
            )
            DamagesWidget(
                stats: damageStats,
                navigateToDamage: $navigateToDamage,
                onDamageSelected: onDamageSelected
            )
            MessagesWidget()
        }
        .padding()
    }
}

struct WelcomeWidget: View {
    let firstName: String

    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            Text(String(format: "welcome_message".localized(), firstName))
                .font(.title2)
                .fontWeight(.bold)
                .foregroundColor(Color("textColor"))
            Text("dashboard_overview".localized())
                .font(.subheadline)
                .foregroundColor(.gray)
        }
        .padding()
        .frame(maxWidth: .infinity, alignment: .leading)
        .background(Color("basicWhiteBlack"))
        .cornerRadius(12)
        .shadow(color: Color.black.opacity(0.2), radius: 4, x: 0, y: 2)
    }
}

struct RemindersWidget: View {
    let reminders: [Reminder]
    let properties: PropertyStats
    @Binding var navigateToProperty: String?

    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            HStack {
                Text("Reminders".localized())
                    .font(.headline)
                    .foregroundColor(Color("textColor"))
                Spacer()
                Image(systemName: "bell.fill")
                    .foregroundColor(.red)
            }

            if reminders.isEmpty {
                Text("No reminders".localized())
                    .font(.subheadline)
                    .foregroundColor(.gray)
                    .padding(.vertical, 8)
            } else {
                ForEach(reminders.prefix(3)) { reminder in
                    Button(action: {
                        if let components = parseLink(reminder.link),
                           properties.recentlyAdded?.first(where: { $0.id == components.propertyId }) != nil {
                            navigateToProperty = components.propertyId
                        }
                    }) {
                        ReminderItem(reminder: reminder)
                    }
                }
                if reminders.count > 3 {
                    Text(String(format: "+%d more".localized(), reminders.count - 3))
                        .font(.caption)
                        .foregroundColor(.gray)
                        .padding(.top, 4)
                    }
                }
            }
            .padding()
            .background(Color("basicWhiteBlack"))
            .cornerRadius(12)
            .shadow(color: Color.black.opacity(0.2), radius: 4, x: 0, y: 2)
    }

    private func parseLink(_ link: String) -> (propertyId: String, damageId: String)? {
        let components = link.components(separatedBy: "/")
        guard components.count >= 7,
              components[2] == "real-property",
              components[3] == "details",
              components[5] == "damage" else {
            return nil
        }
        return (propertyId: components[4], damageId: components[6])
    }
}

struct ReminderItem: View {
    let reminder: Reminder

    var body: some View {
        HStack {
            Circle()
                .fill(priorityColor(for: reminder.priority))
                .frame(width: 10, height: 10)
            VStack(alignment: .leading, spacing: 4) {
                Text(reminder.title)
                    .font(.subheadline)
                    .foregroundColor(Color("textColor"))
                    .lineLimit(2)
                    .multilineTextAlignment(.leading)
                Text(reminder.advice)
                    .font(.caption)
                    .foregroundColor(.gray)
                    .lineLimit(1)
            }
            Spacer()
        }
        .padding(.vertical, 4)
    }

    private func priorityColor(for priority: String) -> Color {
        switch priority.lowercased() {
        case "urgent": return .red
        case "high": return .orange
        case "medium": return .yellow
        case "low": return .green
        default: return .gray
        }
    }
}

struct PropertiesWidget: View {
    let stats: PropertyStats
    @Binding var navigateToProperty: String?

    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            HStack {
                Text("Properties".localized())
                    .font(.headline)
                    .foregroundColor(Color("textColor"))
                Spacer()
                Image(systemName: "house.fill")
                    .foregroundColor(.blue)
            }

            LazyVGrid(
                columns: [
                    GridItem(.flexible()),
                    GridItem(.flexible())
                ],
                spacing: 8
            ) {
                StatItem(label: "Total".localized(), value: "\(stats.total)")
                StatItem(label: "Occupied".localized(), value: "\(stats.occupied)")
                StatItem(label: "Available_property".localized(), value: "\(stats.available)")
                StatItem(label: "Archived".localized(), value: "\(stats.archived)")
                StatItem(label: "Pending Invites".localized(), value: "\(stats.pendingInvites)")
            }

            if let recentProperty = stats.recentlyAdded?.first {
                Button(action: {
                    navigateToProperty = recentProperty.id
                    print("Navigating to property ID: \(recentProperty.id)")
                }) {
                    PropertyItem(property: recentProperty)
                }
            } else {
                Text("No recent properties".localized())
                    .font(.subheadline)
                    .foregroundColor(.gray)
                    .padding(.vertical, 8)
            }
        }
        .padding()
        .background(Color("basicWhiteBlack"))
        .cornerRadius(12)
        .shadow(color: Color.black.opacity(0.2), radius: 4, x: 0, y: 2)
    }
}

struct StatItem: View {
    let label: String
    let value: String

    var body: some View {
        VStack {
            Text(value)
                .font(.title3)
                .fontWeight(.bold)
                .foregroundColor(Color("textColor"))
            Text(label)
                .font(.caption)
                .foregroundColor(.gray)
                .multilineTextAlignment(.center)
        }
        .frame(width: 150, height: 60)
        .padding(8)
        .background(Color.gray.opacity(0.1))
        .cornerRadius(8)
    }
}

struct PropertyItem: View {
    let property: PropertySummary

    var body: some View {
        HStack {
            VStack(alignment: .leading, spacing: 4) {
                Text(property.name)
                    .font(.subheadline)
                    .fontWeight(.medium)
                    .foregroundColor(Color("textColor"))
                Text("\(property.address), \(property.city)")
                    .font(.caption)
                    .foregroundColor(.gray)
                    .lineLimit(1)
                Text(String(format: "%d€/month".localized(), property.rentalPricePerMonth))
                    .font(.caption)
                    .foregroundColor(.gray)
            }
            Spacer()
            Image(systemName: "chevron.right")
                .foregroundColor(.gray)
        }
        .padding(.vertical, 8)
    }
}

struct TenantPropertyWidget: View {
    let property: Property
    @Binding var navigateToProperty: String?

    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            HStack {
                Text("My Property".localized())
                    .font(.headline)
                    .foregroundColor(Color("textColor"))
                Spacer()
                Image(systemName: "house.fill")
                    .foregroundColor(.blue)
            }

            Button(action: {
                navigateToProperty = property.id
            }) {
                VStack(alignment: .leading, spacing: 4) {
                    Text(property.name)
                        .font(.subheadline)
                        .fontWeight(.medium)
                        .foregroundColor(Color("textColor"))
                    Text("\(property.address), \(property.city), \(property.country)")
                        .font(.caption)
                        .foregroundColor(.gray)
                        .lineLimit(1)
                    Text(String(format: "%d€/month".localized(), property.monthlyRent))
                        .font(.caption)
                        .foregroundColor(.gray)
                    if let startDate = property.leaseStartDate, let endDate = property.leaseEndDate {
                        Text("\(formatDate(startDate)) - \(formatDate(endDate))")
                            .font(.caption)
                            .foregroundColor(.gray)
                    } else if let startDate = property.leaseStartDate {
                        Text("Started: \(formatDate(startDate))")
                            .font(.caption)
                            .foregroundColor(.gray)
                    } else {
                        Text("No active lease".localized())
                            .font(.caption)
                            .foregroundColor(.gray)
                    }
                }
                .frame(maxWidth: .infinity, alignment: .leading)
                .padding(.vertical, 8)
            }

            LazyVGrid(
                columns: [
                    GridItem(.flexible()),
                    GridItem(.flexible())
                ],
                spacing: 8
            ) {
                StatItem(label: "Area".localized(), value: "\(formattedValue(property.surface))m²")
                StatItem(label: "Deposit".localized(), value: "\(property.deposit)")
            }
        }
        .padding()
        .background(Color("basicWhiteBlack"))
        .cornerRadius(12)
        .shadow(color: Color.black.opacity(0.2), radius: 4, x: 0, y: 2)
    }

    private func formattedValue(_ value: Double) -> String {
        value == Double(Int(value)) ? String(format: "%.0f", value) : String(format: "%.2f", value)
    }

    private func formatDate(_ dateString: String) -> String {
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

struct DamagesWidget: View {
    let stats: OpenDamageStats
    @Binding var navigateToDamage: (propertyId: String, damageId: String)?
    let onDamageSelected: (String, String) -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            HStack {
                Text("Open Damages".localized())
                    .font(.headline)
                    .foregroundColor(Color("textColor"))
                Spacer()
                Image(systemName: "exclamationmark.triangle.fill")
                    .foregroundColor(.orange)
            }

            LazyVGrid(
                columns: [
                    GridItem(.flexible()),
                    GridItem(.flexible())
                ],
                spacing: 8
            ) {
                StatItem(label: "Total".localized(), value: "\(stats.total)")
                StatItem(label: "Urgent_damage".localized(), value: "\(stats.urgent)")
                StatItem(label: "High_damage".localized(), value: "\(stats.high)")
                StatItem(label: "Medium_damage".localized(), value: "\(stats.medium)")
                StatItem(label: "Low_damage".localized(), value: "\(stats.low)")
                StatItem(label: "Planned This Week".localized(), value: "\(stats.plannedThisWeek)")
            }

            if stats.toFix?.isEmpty ?? true {
                Text("No damages to fix".localized())
                    .font(.subheadline)
                    .foregroundColor(.gray)
                    .padding(.vertical, 8)
            } else if let toFix = stats.toFix {
                ForEach(toFix.prefix(3)) { damage in
                    Button(action: {
                        onDamageSelected(damage.propertyId, damage.id)
                    }) {
                        DamageItem(damage: damage)
                    }
                }
                if toFix.count > 3 {
                    Text(String(format: "+%d more".localized(), toFix.count - 3))
                        .font(.caption)
                        .foregroundColor(.gray)
                        .padding(.top, 4)
                }
            }
        }
        .padding()
        .background(Color("basicWhiteBlack"))
        .cornerRadius(12)
        .shadow(color: Color.black.opacity(0.2), radius: 4, x: 0, y: 2)
    }
}

struct DamageItem: View {
    let damage: DamageSummary

    var body: some View {
        HStack {
            Circle()
                .fill(priorityColor(for: damage.priority))
                .frame(width: 10, height: 10)
            VStack(alignment: .leading, spacing: 4) {
                Text(damage.comment)
                    .font(.subheadline)
                    .foregroundColor(Color("textColor"))
                    .lineLimit(2)
                Text("\(damage.propertyName) - \(damage.roomName)")
                    .font(.caption)
                    .foregroundColor(.gray)
                    .lineLimit(1)
            }
            Spacer()
            Image(systemName: "chevron.right")
                .foregroundColor(.gray)
        }
        .padding(.vertical, 4)
    }

    private func priorityColor(for priority: String) -> Color {
        switch priority.lowercased() {
        case "urgent": return .red
        case "high": return .orange
        case "medium": return .yellow
        case "low": return .green
        default: return .gray
        }
    }
}

struct MessagesWidget: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            HStack {
                Text("Messages".localized())
                    .font(.headline)
                    .foregroundColor(Color("textColor"))
                Spacer()
                Image(systemName: "envelope.fill")
                    .foregroundColor(.blue)
            }
            Text("Messaging coming soon!".localized())
                .font(.subheadline)
                .foregroundColor(.gray)
                .padding(.vertical, 8)
        }
        .padding()
        .background(Color("basicWhiteBlack"))
        .cornerRadius(12)
        .shadow(color: Color.black.opacity(0.2), radius: 4, x: 0, y: 2)
    }
}

struct OverviewView_Previews: PreviewProvider {
    static var previews: some View {
        let loginViewModel = LoginViewModel()
        loginViewModel.user = User(id: "", email: "john@example.com", firstname: "John", lastname: "Doe", role: "tenant")
        return OverviewView()
            .environmentObject(loginViewModel)
    }
}
