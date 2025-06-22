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
    @State private var showError: Bool = false
    @State private var errorMessage: String?
    @Environment(\.dismiss) var dismiss

    private var dateFormatter: DateFormatter {
        let formatter = DateFormatter()
        formatter.dateStyle = .medium
        formatter.timeStyle = .none
        return formatter
    }

    var body: some View {
        ZStack {
            NavigationStack {
                Form {
                    Section(header: Text(String(format: "Invite Tenant to %@".localized(), property.name))) {
                        TextField("Tenant Email".localized(), text: $email)
                            .keyboardType(.emailAddress)
                            .autocapitalization(.none)
                            .disableAutocorrection(true)

                        HStack {
                            Button {
                                showStartDatePicker = true
                            } label: {
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
                        }

                        HStack {
                            Button {
                                tempEndDate = endDate
                                showEndDatePicker = true
                            } label: {
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
                        }
                    }

                    Button("Send Invitation".localized()) {
                        Task { await inviteTenant() }
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

                DatePicker("Start Date".localized(), selection: $startDate, in: Date()..., displayedComponents: [.date])
                    .datePickerStyle(.graphical)
                    .accentColor(Color("LightBlue"))
                    .padding()

                Button("Done".localized()) {
                    showStartDatePicker = false
                }
                .font(.headline)
                .foregroundColor(.white)
                .padding()
                .background(Color("LightBlue"))
                .clipShape(RoundedRectangle(cornerRadius: 10))
                .padding(.bottom, 70)
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
                        get: { tempEndDate ?? max(startDate, Date()) },
                        set: { tempEndDate = $0 }
                    ),
                    in: startDate...,
                    displayedComponents: [.date]
                )
                .datePickerStyle(.graphical)
                .accentColor(Color("LightBlue"))
                .padding()

                HStack {
                    Button("Reset".localized()) {
                        endDate = nil
                        tempEndDate = nil
                        showEndDatePicker = false
                    }
                    .font(.headline)
                    .foregroundColor(Color("RedAlert"))
                    .padding()
                    .disabled(endDate == nil)

                    Spacer()

                    Button("Cancel".localized()) {
                        didCancelEndDate = true
                        showEndDatePicker = false
                    }
                    .font(.headline)
                    .foregroundColor(.gray)
                    .padding()

                    Button("Done".localized()) {
                        if let newEndDate = tempEndDate, newEndDate >= startDate {
                            endDate = newEndDate
                        }
                        showEndDatePicker = false
                    }
                    .font(.headline)
                    .foregroundColor(.white)
                    .padding()
                    .background(Color("LightBlue"))
                    .clipShape(RoundedRectangle(cornerRadius: 10))
                }
                .padding(.bottom, 50)
            }
            .presentationDetents([.medium])
            .presentationDragIndicator(.visible)
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
            await propertyViewModel.fetchProperties()
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
    static var propViewModel = PropertyViewModel(loginViewModel: LoginViewModel())
    static var previews: some View {
        InviteTenantView(
            tenantViewModel: viewModel,
            propertyViewModel: propViewModel,
            property: exampleDataProperty
        )
    }
}
