import SwiftUI

struct DamageDetailView: View {
    let damage: DamageResponse
    @EnvironmentObject var propertyViewModel: PropertyViewModel
    @EnvironmentObject var loginViewModel: LoginViewModel
    @Environment(\.dismiss) var dismiss
    @State private var isLoading: Bool = false
    @State private var errorMessage: String?
    @State private var selectedFixStatus: String = ""
    @State private var showStatusPicker: Bool = false
    @State private var showDatePicker: Bool = false
    @State private var showConfirmFixPicker: Bool = false
    @State private var selectedFixPlannedDate: Date = Date().addingTimeInterval(7 * 24 * 60 * 60)

    var body: some View {
        VStack(spacing: 0) {
            TopBar(title: "Damage Details".localized())
                .overlay(
                    HStack {
                        Button(action: {
                            dismiss()
                        }) {
                            Image(systemName: "chevron.left")
                                .font(.title3)
                                .foregroundColor(Color("textColor"))
                                .frame(width: 40, height: 40)
                                .background(Color.black.opacity(0.2))
                                .clipShape(Circle())
                        }
                        .padding(.trailing, 16)
                    },
                    alignment: .trailing
                )
            
            ScrollView {
                VStack(alignment: .leading, spacing: 16) {
                    HStack {
                        Text(damage.roomName)
                            .font(.title2)
                            .fontWeight(.bold)
                            .foregroundColor(Color("textColor"))
                        Spacer()
                        Text(damage.priority.capitalized.localized())
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
                    .padding(.horizontal)
                    
                    VStack(alignment: .leading, spacing: 8) {
                        Text("Comment".localized())
                            .font(.headline)
                            .foregroundColor(Color("textColor"))
                        Text(damage.comment)
                            .font(.subheadline)
                            .foregroundColor(.gray)
                            .padding()
                            .background(Color.gray.opacity(0.1))
                            .cornerRadius(10)
                    }
                    .padding(.horizontal)
                    
                    VStack(alignment: .leading, spacing: 8) {
                        Text("Status".localized())
                            .font(.headline)
                            .foregroundColor(Color("textColor"))
                        if loginViewModel.userRole == "owner" && canOwnerChangeStatus() {
                            Button(action: {
                                selectedFixStatus = damage.fixStatus
                                showStatusPicker.toggle()
                            }) {
                                HStack {
                                    Text(damage.fixStatus.replacingOccurrences(of: "_", with: " ").capitalized.localized())
                                        .font(.subheadline)
                                        .foregroundColor(.white)
                                    Spacer()
                                    Image(systemName: "chevron.right")
                                        .font(.caption)
                                        .foregroundColor(.white)
                                }
                                .padding()
                                .frame(maxWidth: .infinity)
                                .background(statusColor(damage.fixStatus))
                                .cornerRadius(10)
                            }
                            .padding(.horizontal)
                        } else if loginViewModel.userRole == "tenant" && damage.fixStatus == "awaiting_tenant_confirmation" {
                            Button(action: {
                                showConfirmFixPicker.toggle()
                            }) {
                                HStack {
                                    Text("Confirm Fix".localized())
                                        .font(.subheadline)
                                        .foregroundColor(.white)
                                    Spacer()
                                    Image(systemName: "chevron.right")
                                        .font(.caption)
                                        .foregroundColor(.white)
                                }
                                .padding()
                                .frame(maxWidth: .infinity)
                                .background(statusColor(damage.fixStatus))
                                .cornerRadius(10)
                            }
                            .padding(.horizontal)
                        } else {
                            Text(damage.fixStatus.replacingOccurrences(of: "_", with: " ").capitalized.localized())
                                .font(.subheadline)
                                .foregroundColor(.white)
                                .padding()
                                .frame(maxWidth: .infinity)
                                .background(statusColor(damage.fixStatus))
                                .cornerRadius(10)
                                .padding(.horizontal)
                        }
                    }
                    .padding()
                    
                    if let pictures = damage.pictures, !pictures.isEmpty {
                        VStack(alignment: .leading, spacing: 8) {
                            Text("Pictures".localized())
                                .font(.headline)
                                .foregroundColor(Color("textColor"))
                            ScrollView(.horizontal, showsIndicators: false) {
                                HStack(spacing: 10) {
                                    ForEach(pictures, id: \.self) { picture in
                                        if let image = base64ToImage(picture) {
                                            Image(uiImage: image)
                                                .resizable()
                                                .scaledToFill()
                                                .frame(width: 100, height: 100)
                                                .clipShape(RoundedRectangle(cornerRadius: 10))
                                                .clipped()
                                        }
                                    }
                                }
                                .padding(.horizontal)
                            }
                        }
                        .padding()
                    }
                    
                    VStack(alignment: .leading, spacing: 8) {
                        Text("Dates".localized())
                            .font(.headline)
                            .foregroundColor(Color("textColor"))
                        dateRow(label: "Created At".localized(), date: damage.createdAt)
                        if let updatedAt = damage.updatedAt {
                            dateRow(label: "Updated At".localized(), date: updatedAt)
                        }
                        if let fixPlannedAt = damage.fixPlannedAt {
                            dateRow(label: "Fix Planned At".localized(), date: fixPlannedAt)
                        }
                        if let fixedAt = damage.fixedAt {
                            dateRow(label: "Fixed At".localized(), date: fixedAt)
                        }
                    }
                    .padding(.horizontal)
                    
                    if let errorMessage = errorMessage {
                        Text(errorMessage)
                            .foregroundColor(.red)
                            .padding()
                    }
                }
                .padding(.vertical, 20)
            }
        }
        .confirmationDialog(
            "Select Fix Status".localized(),
            isPresented: $showStatusPicker,
            titleVisibility: .visible
        ) {
            if damage.fixStatus == "pending" {
                Button("In Progress".localized()) {
                    showDatePicker = true
                }
                Button("Fixed".localized()) {
                    updateFixStatus(newStatus: "fixed")
                }
            } else if damage.fixStatus == "planned" {
                Button("Fixed".localized()) {
                    updateFixStatus(newStatus: "fixed")
                }
            }
            Button("Cancel".localized(), role: .cancel) {}
        } message: {
            Text("Choose a new status for the damage".localized())
        }
        .confirmationDialog(
            "Confirm Fix".localized(),
            isPresented: $showConfirmFixPicker,
            titleVisibility: .visible
        ) {
            Button("Confirm Fix".localized()) {
                confirmFix()
            }
            Button("Cancel".localized(), role: .cancel) {}
        } message: {
            Text("Are you sure you want to confirm this damage is fixed?".localized())
        }
        .sheet(isPresented: $showDatePicker) {
            VStack {
                Text("Select Fix Planned Date".localized())
                    .font(.headline)
                    .padding()
                DatePicker(
                    "Fix Planned Date".localized(),
                    selection: $selectedFixPlannedDate,
                    in: Date()...,
                    displayedComponents: [.date, .hourAndMinute]
                )
                .datePickerStyle(.graphical)
                .padding()
                Button(action: {
                    updateFixStatus(newStatus: "planned")
                    showDatePicker = false
                }) {
                    Text("Confirm".localized())
                        .font(.subheadline)
                        .foregroundColor(.white)
                        .padding()
                        .frame(maxWidth: .infinity)
                        .background(Color("LightBlue"))
                        .cornerRadius(10)
                }
                .padding(.horizontal)
                Button(action: {
                    showDatePicker = false
                }) {
                    Text("Cancel".localized())
                        .font(.subheadline)
                        .foregroundColor(.red)
                        .padding()
                }
            }
            .padding()
            .presentationDetents([.medium])
        }
        .navigationBarBackButtonHidden(true)
    }
    
    private func canOwnerChangeStatus() -> Bool {
        let canChange = damage.fixStatus == "pending" || damage.fixStatus == "planned"
        return canChange
    }
    
    private func dateRow(label: String, date: String) -> some View {
        HStack {
            Text(label)
                .font(.subheadline)
                .foregroundColor(.gray)
            Spacer()
            Text(formatDateString(date))
                .font(.subheadline)
                .foregroundColor(.gray)
        }
        .padding(.horizontal)
    }
    
    private func priorityColor(_ priority: String) -> Color {
        switch priority.lowercased() {
        case "low":
            return Color.green
        case "medium":
            return Color.yellow
        case "high":
            return Color.orange
        case "urgent":
            return Color.red
        default:
            return Color.gray
        }
    }

    private func statusColor(_ status: String) -> Color {
        switch status.lowercased() {
        case "fixed":
            return Color.green
        case "planned", "awaiting_tenant_confirmation":
            return Color.orange
        case "pending":
            return Color.red
        default:
            return Color.gray
        }
    }
    
    private func formatDateString(_ dateString: String) -> String {
        let isoFormatter = ISO8601DateFormatter()
        let formatOptions: [ISO8601DateFormatter.Options] = [
            [.withInternetDateTime, .withFractionalSeconds],
            [.withInternetDateTime]
        ]
        
        let displayFormatter = DateFormatter()
        displayFormatter.dateFormat = "dd/MM/yyyy HH:mm"
        displayFormatter.locale = Locale(identifier: "fr_FR")
        
        for options in formatOptions {
            isoFormatter.formatOptions = options
            if let date = isoFormatter.date(from: dateString) {
                return displayFormatter.string(from: date)
            }
        }
        
        return "Invalid Date".localized()
    }
    
    private func base64ToImage(_ base64String: String) -> UIImage? {
        var cleanBase64 = base64String
        if base64String.contains(",") {
            cleanBase64 = base64String.components(separatedBy: ",").last ?? base64String
        }
        if let data = Data(base64Encoded: cleanBase64, options: [.ignoreUnknownCharacters]) {
            return UIImage(data: data)
        }
        return nil
    }
    
    private func updateFixStatus(newStatus: String) {
        Task {
            isLoading = true
            do {
                let token = try await TokenStorage.getValidAccessToken()
                if newStatus == "planned" {
                    let isoFormatter = ISO8601DateFormatter()
                    isoFormatter.formatOptions = [.withInternetDateTime]
                    let fixPlannedAt = isoFormatter.string(from: selectedFixPlannedDate)
                    try await propertyViewModel.updateDamageStatus(
                        propertyId: damage.propertyId,
                        damageId: damage.id,
                        fixPlannedAt: fixPlannedAt,
                        read: true,
                        token: token
                    )
                } else if newStatus == "fixed" {
                    try await propertyViewModel.fixDamage(
                        propertyId: damage.propertyId,
                        damageId: damage.id,
                        token: token
                    )
                } else {
                    try await propertyViewModel.updateDamageStatus(
                        propertyId: damage.propertyId,
                        damageId: damage.id,
                        fixPlannedAt: nil,
                        read: true,
                        token: token
                    )
                }
                try await propertyViewModel.fetchPropertyDamages(propertyId: damage.propertyId)
                dismiss()
            } catch {
                errorMessage = "Error updating status: \(error.localizedDescription)".localized()
            }
            isLoading = false
        }
    }
    
    private func confirmFix() {
        Task {
            isLoading = true
            do {
                let token = try await TokenStorage.getValidAccessToken()
                try await propertyViewModel.fixDamage(
                    propertyId: damage.propertyId,
                    damageId: damage.id,
                    token: token
                )
                try await propertyViewModel.fetchPropertyDamages(propertyId: damage.propertyId)
                dismiss()
            } catch {
                errorMessage = "Error confirming fix: \(error.localizedDescription)".localized()
            }
            isLoading = false
        }
    }
}

struct DamageDetailView_Previews: PreviewProvider {
    static var previews: some View {
        NavigationView {
            DamageDetailView(damage: DamageResponse(
                id: "damage_001",
                comment: "Cracked window in the living room",
                priority: "high",
                roomName: "Living Room",
                fixStatus: "pending",
                pictures: [],
                createdAt: "2025-05-10T21:03:53.293Z",
                updatedAt: nil,
                fixPlannedAt: nil,
                fixedAt: nil,
                leaseId: "lease_001",
                propertyId: "property_001",
                propertyName: "Condo",
                tenantName: "John & Mary Doe",
                read: true
            ))
            .environmentObject(PropertyViewModel(loginViewModel: LoginViewModel()))
            .environmentObject(LoginViewModel())
        }
    }
}
