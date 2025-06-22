//
//  InviteTenantView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/04/2025.
//

import SwiftUI

struct InviteTenantView: View {
    @ObservedObject var tenantViewModel: TenantViewModel
    let property: Property
    @State private var email: String = ""
    @State private var startDate: Date = Date()
    @State private var endDate: Date? = nil
    @State private var showError: Bool = false
    @State private var errorMessage: String?
    @Environment(\.dismiss) var dismiss
    
    var body: some View {
        ZStack {
            NavigationStack {
                Form {
                    Section(header: Text(String(format: "Invite Tenant to %@".localized(), property.name))) {
                        TextField("Tenant Email".localized(), text: $email)
                            .keyboardType(.emailAddress)
                            .autocapitalization(.none)
                            .disableAutocorrection(true)
                
                        DatePicker("Start Date".localized(),
                                  selection: $startDate,
                                  displayedComponents: .date)
                            .datePickerStyle(.compact)
                    
                        DatePicker("End Date (Optional)".localized(),
                                  selection: Binding(
                                    get: { endDate ?? Date() },
                                    set: { endDate = $0 }
                                  ),
                                  displayedComponents: .date)
                            .datePickerStyle(.compact)
                            .foregroundColor(endDate == nil ? .gray : .primary)
                            .overlay(
                                Button(action: {
                                    endDate = nil
                                }) {
                                    Image(systemName: "xmark.circle.fill")
                                        .foregroundColor(.gray)
                                }
                                .opacity(endDate != nil ? 1 : 0),
                                alignment: .trailing
                            )
                    }
                
                    Button("Send Invitation".localized()) {
                        Task {
                            await inviteTenant()
                        }
                    }
                    .disabled(email.isEmpty || !email.contains("@"))
                }
                .navigationTitle("Invite Tenant".localized())
                .toolbar {
                    ToolbarItem(placement: .cancellationAction) {
                        Button("Cancel".localized()) {
                            dismiss()
                        }
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
    }

    private func inviteTenant() async {
        do {
            try await tenantViewModel.inviteTenant(
                propertyId: property.id,
                email: email,
                startDate: startDate,
                endDate: endDate
            )
            await MainActor.run {
                errorMessage = "Invitation sent successfully!".localized()
                showError = true
                DispatchQueue.main.asyncAfter(deadline: .now() + 5) {
                    dismiss()
                }
            }
        } catch {
            await MainActor.run {
                errorMessage = "Error inviting tenant: \(error.localizedDescription)".localized()
                showError = true
            }
        }
    }
}

struct InviteTenantView_Previews: PreviewProvider {
    static var viewModel = TenantViewModel()
    static var previews: some View {
        InviteTenantView(
            tenantViewModel: viewModel,
            property: exampleDataProperty
        )
    }
}
