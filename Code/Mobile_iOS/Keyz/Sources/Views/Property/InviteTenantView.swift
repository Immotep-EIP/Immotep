//
//  InviteTenantView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/04/2025.
//

import SwiftUI

struct InviteTenantView: View {
    @ObservedObject var tenantViewModel: TenantViewModel
    @ObservedObject var propertyViewModel: PropertyViewModel
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
                        do {
                            try await tenantViewModel.inviteTenant(
                                propertyId: property.id,
                                email: email,
                                startDate: startDate,
                                endDate: endDate
                            )
                            await propertyViewModel.fetchProperties()
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

struct InviteTenantView_Previews: PreviewProvider {
    static var viewModel = TenantViewModel()
    static var propViewModel = PropertyViewModel(loginViewModel: LoginViewModel())
    static var previews: some View {
        InviteTenantView(
            tenantViewModel: viewModel,
            propertyViewModel: propViewModel,
            property: exampleDataProperty
        )
    }
}

