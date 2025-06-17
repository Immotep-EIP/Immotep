//
//  InventoryRoomExitEvaluationView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 09/06/2025.
//

import SwiftUI
import UIKit

struct InventoryRoomExitEvaluationView: View {
    @EnvironmentObject var inventoryViewModel: InventoryViewModel
    @Environment(\.dismiss) var dismiss
    let selectedRoom: LocalRoom

    @State private var showSheet: Bool = false
    @State private var sourceType: UIImagePickerController.SourceType = .photoLibrary
    @State private var replaceIndex: Int?
    @State private var isLoading: Bool = false
    @State private var errorMessage: String?
    @State private var isReportSent: Bool = false

    let stateMapping: [String: String] = [
        "broken": "Broken",
        "bad": "Bad",
        "good": "Good",
        "new": "Good",
        "medium": "Images"
    ]

    var body: some View {
        VStack(spacing: 0) {
            TopBar(title: "Room Analysis - Exit".localized())

            ScrollView {
                Section {
                    PicturesSegment(selectedImages: $inventoryViewModel.selectedImages, showImagePickerOptions: showImageSelection)
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
                        Picker("Select a status".localized(), selection: $inventoryViewModel.selectedStatus) {
                            Text("Select room status".localized()).tag("Select room status")
                            ForEach(Array(stateMapping.values), id: \.self) { status in
                                Text(status.localized()).tag(status)
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
                            await sendRoomComparisonReport()
                            isLoading = false
                        }
                    }, label: {
                        Text("Send Room Report".localized())
                            .padding()
                            .frame(maxWidth: .infinity)
                            .background(Color("LightBlue"))
                            .foregroundColor(.white)
                            .cornerRadius(10)
                    })
                    .disabled(isLoading || inventoryViewModel.selectedImages.isEmpty)
                    .padding()
                }

                if let errorMessage = errorMessage {
                    Text(errorMessage)
                        .foregroundColor(.red)
                        .padding()
                }
            }
        }
        .fullScreenCover(isPresented: $showSheet) {
            ImagePicker(sourceType: $sourceType, selectedImage: createImagePickerBinding())
        }
        .onAppear {
            inventoryViewModel.selectRoom(selectedRoom)
            inventoryViewModel.selectedImages = []
            inventoryViewModel.comment = ""
            inventoryViewModel.selectedStatus = "Select room status"
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

    private func showImageSelection(replaceIndex: Int?) {
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

    private func sendRoomComparisonReport() async {
        do {
            await inventoryViewModel.fetchLastInventoryReport()
            if let oldReportId = inventoryViewModel.lastReportId {
                try await inventoryViewModel.compareRoomReport(oldReportId: oldReportId)
                isReportSent = true
            } else {
                errorMessage = "No previous inventory report found for comparison.".localized()
            }
        } catch let error as NSError {
            switch error.code {
            case 404:
                errorMessage = "Property or old report not found. Please check the property details.".localized()
            case 403:
                errorMessage = "You do not have permission to access this property.".localized()
            case 400:
                if error.localizedDescription.contains("datauri") {
                    errorMessage = "Invalid image format. Please ensure all images are valid JPEGs.".localized()
                } else {
                    errorMessage = "Invalid request: \(error.localizedDescription)".localized()
                }
            case 0 where error.localizedDescription.contains("No active lease found"):
                errorMessage = "No active lease found for this property.".localized()
            default:
                errorMessage = "Error: \(error.localizedDescription)".localized()
            }
            print("Error sending room comparison report: \(error.localizedDescription)")
        }
    }

    private func validateReport() async {
        if let roomIndex = inventoryViewModel.localRooms.firstIndex(where: { $0.id == selectedRoom.id }) {
            inventoryViewModel.localRooms[roomIndex].checked = true
            inventoryViewModel.localRooms[roomIndex].images = inventoryViewModel.selectedImages
            inventoryViewModel.localRooms[roomIndex].status = inventoryViewModel.selectedStatus
            inventoryViewModel.localRooms[roomIndex].comment = inventoryViewModel.comment
        }
        await inventoryViewModel.markRoomAsChecked(selectedRoom)
        dismiss()
    }
}
