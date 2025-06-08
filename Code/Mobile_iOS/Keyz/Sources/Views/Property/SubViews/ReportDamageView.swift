//
//  ReportDamageView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/06/2025.
//

import SwiftUI

struct ReportDamageView: View {
    @ObservedObject var viewModel: PropertyViewModel
    let propertyId: String
    @Environment(\.dismiss) var dismiss
    @State private var comment = ""
    @State private var priority = "low"
    @State private var roomName = ""
    @State private var pictures: [String] = []
    @State private var errorMessage: String?

    var body: some View {
        NavigationView {
            VStack(spacing: 16) {
                Text("Report a Damage")
                    .font(.title)
                    .padding()

                TextField("Room Name", text: $roomName)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .padding(.horizontal)

                TextField("Comment", text: $comment)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .padding(.horizontal)

                Picker("Priority", selection: $priority) {
                    Text("Low").tag("low")
                    Text("Medium").tag("medium")
                    Text("High").tag("high")
                    Text("Urgent").tag("urgent")
                }
                .pickerStyle(MenuPickerStyle())
                .padding(.horizontal)

                Text("Add Photos (not implemented yet)")
                    .foregroundColor(.gray)
                    .padding(.horizontal)

                Button(action: {
                    Task {
                        do {
                            let token = try await TokenStorage.getValidAccessToken()
                            if let leaseId = try await viewModel.fetchActiveLeaseIdForProperty(propertyId: propertyId, token: token) {
                                let damageRequest = DamageRequest(comment: comment, priority: priority, roomName: roomName, pictures: pictures.isEmpty ? nil : pictures)
                                let damageId = try await viewModel.createDamage(propertyId: propertyId, leaseId: leaseId, damage: damageRequest, token: token)
                                print("Damage reported with ID: \(damageId)")
                                dismiss()
                            } else {
                                errorMessage = "No active lease found.".localized()
                            }
                        } catch {
                            errorMessage = "Error reporting damage: \(error.localizedDescription)".localized()
                            print("Error reporting damage: \(error.localizedDescription)")
                        }
                    }
                }) {
                    Text("Submit")
                        .frame(maxWidth: .infinity)
                        .padding()
                        .background(Color("LightBlue"))
                        .foregroundColor(.white)
                        .cornerRadius(10)
                }
                .padding(.horizontal)

                if let errorMessage = errorMessage {
                    Text(errorMessage)
                        .foregroundColor(.red)
                        .padding()
                }

                Spacer()
            }
            .navigationBarItems(trailing: Button("Cancel") {
                dismiss()
            })
        }
    }
}

struct ReportDamageView_Previews: PreviewProvider {
    static var previews: some View {
        ReportDamageView(viewModel: PropertyViewModel(loginViewModel: LoginViewModel()), propertyId: "test_property")
    }
}
