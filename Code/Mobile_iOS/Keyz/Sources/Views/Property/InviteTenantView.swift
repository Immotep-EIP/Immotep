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
    @Environment(\.dismiss) var dismiss
    
    var body: some View {
        NavigationStack {
            Form {
                Section(header: Text(String(format: "Invite Tenant to %@".localized(), property.name))) {
                    TextField("Tenant Email".localized(), text: $email)
                        .textFieldStyle(RoundedBorderTextFieldStyle())
                        .keyboardType(.emailAddress)
                        .autocapitalization(.none)
                    
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
                        do {
                            try await tenantViewModel.inviteTenant(
                                propertyId: property.id,
                                email: email,
                                startDate: startDate,
                                endDate: endDate
                            )
                            dismiss()
                        } catch {
                            print("Error inviting tenant: \(error)".localized())
                        }
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
    }
}
