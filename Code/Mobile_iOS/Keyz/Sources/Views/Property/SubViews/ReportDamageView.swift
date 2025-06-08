//
//  ReportDamageView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/06/2025.
//

import SwiftUI

struct ReportDamageView: View {
    @ObservedObject var viewModel: TenantPropertyViewModel
    let propertyId: String
    @Environment(\.dismiss) var dismiss
    @State private var comment = ""
    @State private var priority = "low"
    @State private var selectedRoomId = ""
    @State private var pictures: [String] = []
    @State private var errorMessage: String?
    @State private var offset: CGFloat = 1000
    @State private var isLoading = false

    private let priorities = ["low", "medium", "high", "urgent"]

    var body: some View {
        ZStack {
            Color.black.opacity(0.5)
                .onTapGesture { close() }

            VStack(spacing: 16) {
                Text("Report a Damage".localized())
                    .font(.title2)
                    .bold()
                    .padding(.top)
                    .foregroundStyle(Color("textColor"))

                Text("Please provide details about the damage.".localized())
                    .font(.body)
                    .multilineTextAlignment(.center)
                    .padding(.horizontal)
                    .foregroundStyle(Color("textColor"))

                Picker("Select Room".localized(), selection: $selectedRoomId) {
                    Text("Choose a room...".localized()).tag("")
                    ForEach(viewModel.rooms, id: \.id) { room in
                        Text(room.name.capitalized).tag(room.id)
                    }
                }
                .pickerStyle(.menu)
                .padding(.horizontal)
                .padding(.vertical, 8)
                .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
                .overlay(RoundedRectangle(cornerRadius: 8).stroke(Color.gray.opacity(0.5), lineWidth: 1))

                TextField("Comment".localized(), text: $comment)
                    .padding()
                    .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
                    .overlay(RoundedRectangle(cornerRadius: 8).stroke(Color.gray.opacity(0.5), lineWidth: 1))
                    .padding(.horizontal)

                Picker("Priority".localized(), selection: $priority) {
                    ForEach(priorities, id: \.self) { priority in
                        Text(priority.capitalized).tag(priority)
                    }
                }
                .pickerStyle(.menu)
                .padding(.horizontal)
                .padding(.vertical, 8)
                .background(RoundedRectangle(cornerRadius: 8).fill(Color("textfieldBackground")))
                .overlay(RoundedRectangle(cornerRadius: 8).stroke(Color.gray.opacity(0.5), lineWidth: 1))

                Text("Add Photos (not implemented yet)".localized())
                    .foregroundColor(.gray)
                    .padding(.horizontal)

                if let errorMessage {
                    Text(errorMessage)
                        .foregroundColor(.red)
                        .padding(.horizontal)
                        .transition(.opacity)
                }

                HStack(spacing: 16) {
                    Button {
                        close()
                    } label: {
                        Text("Cancel".localized())
                            .font(.system(size: 16, weight: .bold))
                            .foregroundColor(.white)
                            .padding()
                            .frame(maxWidth: .infinity)
                            .background(RoundedRectangle(cornerRadius: 20).fill(.gray))
                    }

                    Button {
                        submitDamage()
                    } label: {
                        if isLoading {
                            ProgressView()
                                .progressViewStyle(CircularProgressViewStyle(tint: .white))
                                .padding()
                                .frame(maxWidth: .infinity)
                                .background(RoundedRectangle(cornerRadius: 20).fill(Color("LightBlue").opacity(0.5)))
                        } else {
                            Text("Submit".localized())
                                .font(.system(size: 16, weight: .bold))
                                .foregroundColor(.white)
                                .padding()
                                .frame(maxWidth: .infinity)
                                .background(RoundedRectangle(cornerRadius: 20).fill(Color("LightBlue")))
                        }
                    }
                    .disabled(selectedRoomId.isEmpty || comment.isEmpty)
                }
                .padding(.horizontal)
                .padding(.bottom)
            }
            .fixedSize(horizontal: false, vertical: true)
            .padding()
            .background(Color("basicWhiteBlack"))
            .clipShape(RoundedRectangle(cornerRadius: 20))
            .shadow(radius: 20)
            .padding(30)
            .offset(y: offset)
            .overlay(alignment: .topTrailing) {
                Button {
                    close()
                } label: {
                    Image(systemName: "xmark")
                        .font(.title2)
                        .fontWeight(.medium)
                        .tint(Color("textColor"))
                        .padding()
                }
            }
            .onAppear {
                withAnimation(.spring()) {
                    offset = 0
                }
                fetchRooms()
            }
        }
        .ignoresSafeArea()
    }

    private func fetchRooms() {
        Task {
            do {
                let token = try await TokenStorage.getValidAccessToken()
                let rooms = try await viewModel.fetchPropertyRooms(token: token)
                viewModel.rooms = rooms
                if !rooms.isEmpty {
                    selectedRoomId = rooms.first!.id
                }
            } catch {
                errorMessage = "Error fetching rooms: \(error.localizedDescription)".localized()
            }
        }
    }

    private func submitDamage() {
        isLoading = true
        errorMessage = nil
        Task {
            do {
                let token = try await TokenStorage.getValidAccessToken()
                if let leaseId = try await viewModel.fetchActiveLeaseIdForProperty(propertyId: propertyId, token: token) {
                    let damageRequest = DamageRequest(
                        comment: comment,
                        priority: priority,
                        roomId: selectedRoomId,
                        pictures: pictures.isEmpty ? nil : pictures
                    )
                    let damageId = try await viewModel.createDamage(
                        propertyId: propertyId,
                        leaseId: leaseId,
                        damage: damageRequest,
                        token: token
                    )
                    print("Damage reported with ID: \(damageId)")
                    close()
                } else {
                    errorMessage = "No active lease found.".localized()
                }
            } catch {
                errorMessage = "Error reporting damage: \(error.localizedDescription)".localized()
                print("Error reporting damage: \(error.localizedDescription)")
            }
            isLoading = false
        }
    }

    private func close() {
        withAnimation(.spring()) {
            offset = 1000
            dismiss()
        }
    }
}

struct ReportDamageView_Previews: PreviewProvider {
    static var previews: some View {
        ReportDamageView(viewModel: TenantPropertyViewModel(), propertyId: "test_property")
            .environment(\.locale, .init(identifier: "en"))
    }
}
