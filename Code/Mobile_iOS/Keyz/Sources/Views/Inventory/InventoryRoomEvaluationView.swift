//
//  InventoryRoomEvaluationView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 09/06/2025.
//

import SwiftUI
import UIKit

struct InventoryRoomEvaluationView: View {
    @EnvironmentObject var inventoryViewModel: InventoryViewModel
    @Environment(\.dismiss) var dismiss
    let selectedRoom: LocalRoom

    @State private var showSheet = false
    @State private var sourceType: UIImagePickerController.SourceType = .photoLibrary
    @State private var replaceIndex: Int?
    @State private var isLoading: Bool = false
    @State private var showError = false
    @State private var errorMessage: String?
    @State private var isReportSent: Bool = false

    let stateMapping: [String: String] = [
        "not_set": "Select your room status".localized(),
        "broken": "Broken".localized(),
        "needsRepair": "Needs Repair".localized(),
        "bad": "Bad".localized(),
        "medium": "Medium".localized(),
        "good": "Good".localized(),
        "new": "New".localized()
    ]

    var body: some View {
        ZStack {
            VStack(spacing: 0) {
                TopBar(title: "Room Analysis".localized())
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
                    Section {
                        PicturesSegment(selectedImages: $inventoryViewModel.selectedImages, showImagePickerOptions: showImagePickerOptions)
                    }

                    VStack {
                        HStack {
                            Text("Comment".localized())
                                .font(.headline)
                            Spacer()
                        }
                        TextEditor(text: $inventoryViewModel.comment)
                            .frame(height: 100)
                            .padding()
                            .cornerRadius(20)
                            .overlay(
                                RoundedRectangle(cornerRadius: 10)
                                    .stroke(Color.gray, lineWidth: 1)
                            )
                    }
                    .padding()

                    VStack {
                        HStack {
                            Text("Status".localized())
                                .font(.headline)
                            Spacer()
                        }
                        HStack {
                            Picker("Select a status".localized(), selection: $inventoryViewModel.roomStatus) {
                                ForEach(Array(stateMapping.keys.sorted()), id: \.self) { key in
                                    Text(stateMapping[key] ?? key).tag(key)
                                }
                            }
                            .frame(maxWidth: .infinity)
                            .pickerStyle(MenuPickerStyle())
                            .padding()
                            .background(Color.gray.opacity(0.1))
                            .cornerRadius(10)
                            .overlay(
                                RoundedRectangle(cornerRadius: 10)
                                    .stroke(Color.gray, lineWidth: 1)
                            )
                            .accessibilityIdentifier("RoomStatusPicker")
                            Spacer()
                        }
                    }
                    .padding()

                    if isReportSent {
                        Button(action: {
                            Task {
                                await validateReport()
                            }
                        }, label: {
                            Text("Validate".localized())
                                .padding()
                                .frame(maxWidth: .infinity)
                                .background(Color("LightBlue"))
                                .foregroundColor(.white)
                                .cornerRadius(10)
                        })
                        .padding()
                    } else {
                        Button(action: {
                            Task {
                                isLoading = true
                                await sendRoomReport()
                                isLoading = false
                            }
                        }, label: {
                            ZStack {
                                if isLoading {
                                    ProgressView()
                                        .progressViewStyle(CircularProgressViewStyle())
                                        .tint(.white)
                                } else {
                                    Text("Send Room Report".localized())
                                }
                            }
                            .padding()
                            .frame(maxWidth: .infinity)
                            .background(inventoryViewModel.selectedImages.isEmpty ? Color.gray : Color("LightBlue"))
                            .foregroundColor(.white)
                            .cornerRadius(10)
                            .scaleEffect(isLoading ? 0.95 : 1.0)
                            .animation(.easeInOut(duration: 0.2), value: isLoading)
                        })
                        .disabled(isLoading || inventoryViewModel.selectedImages.isEmpty)
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
        .fullScreenCover(isPresented: $showSheet) {
            ImagePicker(sourceType: $sourceType, selectedImage: createImagePickerBinding())
        }
        .onAppear {
            inventoryViewModel.selectRoom(selectedRoom)
            inventoryViewModel.selectedImages = selectedRoom.images.isEmpty ? [] : selectedRoom.images
            inventoryViewModel.comment = selectedRoom.comment.isEmpty ? "" : selectedRoom.comment
            let validStates = stateMapping.keys
            inventoryViewModel.roomStatus = validStates.contains(selectedRoom.status.lowercased()) ? selectedRoom.status.lowercased() : "not_set"
        }
    }

    private func createImagePickerBinding() -> Binding<UIImage?> {
        Binding(
            get: { nil },
            set: { image in
                if let image = image {
                    if let index = replaceIndex {
                        inventoryViewModel.selectedImages[index] = image
                    } else {
                        inventoryViewModel.selectedImages.append(image)
                    }
                }
            }
        )
    }

    private func showImagePickerOptions(replaceIndex: Int?) {
        self.replaceIndex = replaceIndex
        let actionSheet = UIAlertController(title: "Select Image Source".localized(), message: nil, preferredStyle: .actionSheet)
        actionSheet.addAction(UIAlertAction(title: "Take Photo".localized(), style: .default, handler: { _ in
            self.sourceType = .camera
            self.showSheet.toggle()
        }))
        actionSheet.addAction(UIAlertAction(title: "Choose from Library".localized(), style: .default, handler: { _ in
            self.sourceType = .photoLibrary
            self.showSheet.toggle()
        }))
        actionSheet.addAction(UIAlertAction(title: "Cancel".localized(), style: .cancel, handler: nil))
        if let windowScene = UIApplication.shared.connectedScenes.first as? UIWindowScene,
           let rootViewController = windowScene.windows.first?.rootViewController {
            rootViewController.present(actionSheet, animated: true, completion: nil)
        }
    }

    private func sendRoomReport() async {
        do {
            try await inventoryViewModel.sendRoomReport()
            isReportSent = true
        } catch let error as NSError {
            switch error.code {
            case 404:
                errorMessage = "Property or lease not found.".localized()
            case 403:
                errorMessage = "You do not have permission to access this property.".localized()
            case 400:
                if error.localizedDescription.contains("datauri") {
                    errorMessage = "Invalid image format. Please ensure all images are valid JPEGs.".localized()
                } else {
                    errorMessage = "Invalid request: \(error.localizedDescription)".localized()
                }
            case 0 where error.localizedDescription.contains("No active lease found"):
                errorMessage = "No active lease found.".localized()
            default:
                errorMessage = "Error sending report.".localized()
            }
            showError = true
        }
    }

    private func validateReport() async {
        if let roomIndex = inventoryViewModel.localRooms.firstIndex(where: { $0.id == selectedRoom.id }) {
            inventoryViewModel.localRooms[roomIndex].checked = true
            inventoryViewModel.localRooms[roomIndex].images = inventoryViewModel.selectedImages
            inventoryViewModel.localRooms[roomIndex].status = inventoryViewModel.roomStatus
            inventoryViewModel.localRooms[roomIndex].comment = inventoryViewModel.comment
            inventoryViewModel.selectedRoom = inventoryViewModel.localRooms[roomIndex]
        }
        await inventoryViewModel.markRoomAsChecked(selectedRoom)
        dismiss()
    }
}
