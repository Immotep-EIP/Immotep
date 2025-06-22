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
    @State private var showStartDatePicker: Bool = false
    @State private var showEndDatePicker: Bool = false
    @State private var tempEndDate: Date? = nil
    @State private var didCancelEndDate: Bool = false
    @Environment(\.dismiss) var dismiss
    
    private var dateFormatter: DateFormatter {
        let formatter = DateFormatter()
        formatter.dateStyle = .medium
        formatter.timeStyle = .none
        return formatter
    }
    
    var body: some View {
        NavigationStack {
            Form {
                Section(header: Text(String(format: "Invite Tenant to %@".localized(), property.name))) {
                    TextField("Tenant Email".localized(), text: $email)
                        .keyboardType(.emailAddress)
                        .autocapitalization(.none)
                        .disableAutocorrection(true)
                    
                    HStack {
                        Button(action: {
                            showStartDatePicker = true
                        }) {
                            HStack {
                                Text("Start Date".localized())
                                    .foregroundColor(.primary)
                                Spacer()
                                Text(dateFormatter.string(from: startDate))
                                    .foregroundColor(.gray)
                            }
                            .padding(.vertical, 8)
                            .contentShape(Rectangle())
                        }
                        .accessibilityLabel("Select Start Date")
                    }
                    
                    HStack {
                        Button(action: {
                            tempEndDate = endDate
                            showEndDatePicker = true
                        }) {
                            HStack {
                                Text("End Date (Optional)".localized())
                                    .foregroundColor(endDate == nil ? .gray : .primary)
                                Spacer()
                                Text(endDate.map { dateFormatter.string(from: $0) } ?? "Not Set")
                                    .foregroundColor(.gray)
                            }
                            .padding(.vertical, 8)
                            .contentShape(Rectangle())
                        }
                        .accessibilityLabel("Select End Date")
                    }
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
                            print("Error inviting tenant: \(error.localizedDescription)".localized())
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
            .onChange(of: startDate) {
                if let currentEndDate = endDate, startDate > currentEndDate {
                    endDate = nil
                }
            }
            .sheet(isPresented: $showStartDatePicker) {
                VStack {
                    Text("Select Start Date".localized())
                        .font(.headline)
                        .padding()
                    
                    DatePicker(
                        "Start Date".localized(),
                        selection: $startDate,
                        in: Date()...,
                        displayedComponents: [.date]
                    )
                    .datePickerStyle(.graphical)
                    .padding()
                    
                    Button("Done".localized()) {
                        showStartDatePicker = false
                    }
                    .font(.headline)
                    .foregroundColor(.white)
                    .padding()
                    .background(Color.blue)
                    .clipShape(RoundedRectangle(cornerRadius: 10))
                    .padding(.bottom)
                }
                .presentationDetents([.medium])
                .presentationDragIndicator(.visible)
            }
            .sheet(isPresented: $showEndDatePicker, onDismiss: {
                if !didCancelEndDate, let selected = tempEndDate, selected >= startDate {
                    endDate = selected
                }
                didCancelEndDate = false
            }) {
                VStack {
                    Text("Select End Date".localized())
                        .font(.headline)
                        .padding()
                    
                    DatePicker(
                        "End Date".localized(),
                        selection: Binding(
                            get: { endDate ?? max(startDate, Date()) },
                            set: { endDate = $0 }
                        ),
                        in: startDate...,
                        displayedComponents: [.date]
                    )
                    .datePickerStyle(.graphical)

                    HStack {
                        Button("Reset".localized()) {
                            endDate = nil
                            tempEndDate = nil
                            showEndDatePicker = false
                        }
                        .font(.headline)
                        .foregroundColor(.red)
                        .padding()
                        .disabled(endDate == nil)
                                                
                        Button("Done".localized()) {
                            if let newEndDate = tempEndDate, newEndDate >= startDate {
                                endDate = newEndDate
                            }
                            showEndDatePicker = false
                        }
                        .font(.headline)
                        .foregroundColor(.white)
                        .padding()
                        .background(Color.blue)
                        .clipShape(RoundedRectangle(cornerRadius: 10))
                    }
                    .padding(.bottom)
                }
                .presentationDetents([.medium])
                .presentationDragIndicator(.visible)
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
