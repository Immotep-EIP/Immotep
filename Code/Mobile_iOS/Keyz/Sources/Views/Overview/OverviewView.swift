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
    @AppStorage("isLoggedIn") var isLoggedIn: Bool = false
    @State private var navigateToProperty: Property? = nil
    @State private var navigateToDamage: (propertyId: String, damageId: String)? = nil
    @State private var showError = false
    @State private var errorMessage: String?
    @State private var navigateToReportDamage: Bool = false
    @State private var navigateToInventory: Bool = false

    init() {
        self._viewModel = StateObject(wrappedValue: OverviewViewModel())
    }

    var body: some View {
        NavigationStack {
            ZStack {
                VStack(spacing: 0) {
                    TopBar(title: "Keyz")
                    ScrollView {
                        if viewModel.isLoading {
                            ProgressView()
                                .padding()
                        } else if viewModel.errorMessage != nil {
                            EmptyView()
                        } else if let dashboardData = viewModel.dashboardData {
                            LazyVGrid(
                                columns: [GridItem(.flexible())],
                                spacing: 16
                            ) {
                                WelcomeWidget(
                                    firstName: loginViewModel.user?.firstname ?? "User"
                                )
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
                                    navigateToDamage: $navigateToDamage
                                )
                            }
                            .padding()
                        } else {
                            Text("No data available".localized())
                                .foregroundColor(.gray)
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
            .navigationBarBackButtonHidden(true)
            .navigationDestination(isPresented: Binding(
                get: { navigateToProperty != nil },
                set: { if !$0 { navigateToProperty = nil } }
            )) {
                if let property = navigateToProperty {
                    PropertyDetailView(
                        property: property,
                        viewModel: PropertyViewModel(loginViewModel: loginViewModel),
                        navigateToReportDamage: $navigateToReportDamage,
                        navigateToInventory: $navigateToInventory
                    )
                    .environmentObject(loginViewModel)
                }
            }
            .navigationDestination(isPresented: Binding(
                get: { navigateToDamage != nil },
                set: { if !$0 { navigateToDamage = nil } }
            )) {
                if let damage = navigateToDamage {
                    PropertyDetailView(
                        property: Property(
                            id: damage.propertyId,
                            ownerID: "",
                            name: "",
                            address: "",
                            city: "",
                            postalCode: "",
                            country: "",
                            photo: nil,
                            monthlyRent: 0,
                            deposit: 0,
                            surface: 0,
                            isAvailable: "",
                            tenantName: nil,
                            leaseId: nil,
                            leaseStartDate: nil,
                            leaseEndDate: nil,
                            documents: [],
                            createdAt: "",
                            rooms: [],
                            damages: []
                        ),
                        viewModel: PropertyViewModel(loginViewModel: loginViewModel),
                        navigateToReportDamage: $navigateToReportDamage,
                        navigateToInventory: $navigateToInventory
                    )
                    .environmentObject(loginViewModel)
                }
            }
            .onAppear {
                loginViewModel.loadUser()
                Task {
                    if loginViewModel.userRole == "owner" {                        
                        await viewModel.fetchDashboardData()
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
        }
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
    @Binding var navigateToProperty: Property?

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
                           let property = properties.recentlyAdded?.first(where: { $0.id == components.propertyId }) {
                            navigateToProperty = Property(
                                id: property.id,
                                ownerID: property.ownerId,
                                name: property.name,
                                address: property.address,
                                city: property.city,
                                postalCode: property.postalCode,
                                country: property.country,
                                photo: nil,
                                monthlyRent: property.rentalPricePerMonth,
                                deposit: property.depositPrice,
                                surface: property.areaSqm,
                                isAvailable: property.archived ? "archived" : (properties.available > 0 ? "available" : "occupied"),
                                tenantName: nil,
                                leaseId: nil,
                                leaseStartDate: nil,
                                leaseEndDate: nil,
                                documents: [],
                                createdAt: property.createdAt,
                                rooms: [],
                                damages: []
                            )
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
    @Binding var navigateToProperty: Property?

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
                    navigateToProperty = Property(
                        id: recentProperty.id,
                        ownerID: recentProperty.ownerId,
                        name: recentProperty.name,
                        address: recentProperty.address,
                        city: recentProperty.city,
                        postalCode: recentProperty.postalCode,
                        country: recentProperty.country,
                        photo: nil,
                        monthlyRent: recentProperty.rentalPricePerMonth,
                        deposit: recentProperty.depositPrice,
                        surface: recentProperty.areaSqm,
                        isAvailable: recentProperty.archived ? "archived" : (stats.available > 0 ? "available" : "occupied"),
                        tenantName: nil,
                        leaseId: nil,
                        leaseStartDate: nil,
                        leaseEndDate: nil,
                        documents: [],
                        createdAt: recentProperty.createdAt,
                        rooms: [],
                        damages: []
                    )
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
                Text("$\(property.rentalPricePerMonth)/month".localized())
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

struct DamagesWidget: View {
    let stats: OpenDamageStats
    @Binding var navigateToDamage: (propertyId: String, damageId: String)?

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
                        navigateToDamage = (propertyId: damage.propertyId, damageId: damage.id)
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

struct OverviewView_Previews: PreviewProvider {
    static var previews: some View {
        let loginViewModel = LoginViewModel()
        loginViewModel.user = User(id: "", email: "john@example.com", firstname: "John", lastname: "Doe", role: "owner")
        return OverviewView()
            .environmentObject(loginViewModel)
    }
}
