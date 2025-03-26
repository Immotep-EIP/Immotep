//
//  InventoryExitEvaluationView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 19/01/2025.
//

import SwiftUI

struct InventoryExitEvaluationView: View {
    @EnvironmentObject var inventoryViewModel: InventoryViewModel
    @Environment(\.presentationMode) var presentationMode
    let selectedStuff: LocalInventory

    @State private var showSheet = false
    @State private var sourceType: UIImagePickerController.SourceType = .photoLibrary
    @State private var replaceIndex: Int?
    @State private var isLoading: Bool = false
    @State private var errorMessage: String?
    @State private var isReportSent: Bool = false

    let stateMapping: [String: String] = [
        "broken": "Broken",
        "needsRepair": "Needs Repair",
        "bad": "Bad",
        "medium": "Medium",
        "good": "Good",
        "new": "New"
    ]

    var body: some View {
        VStack(spacing: 0) {
            TopBar(title: "Inventory Exit")

            ScrollView {
                Section {
                    PicturesSegment(selectedImages: $inventoryViewModel.selectedImages, showImagePickerOptions: showImagePickerOptions)
                }

                VStack {
                    HStack {
                        Text("Comments")
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
                        Text("Status")
                            .font(.headline)
                        Spacer()
                    }
                    HStack {
                        Picker("Select Equipment Status", selection: $inventoryViewModel.selectedStatus) {
                            Text("Select your equipment status").tag("Select your equipment status")
                            ForEach(Array(stateMapping.values), id: \.self) { status in
                                Text(status).tag(status)
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
                        Text("Validate")
                            .padding()
                            .frame(maxWidth: .infinity)
                            .background(Color.blue)
                            .foregroundColor(.white)
                            .cornerRadius(10)
                    })
                    .padding()
                } else {
                    Button(action: {
                        Task {
                            isLoading = true
                            await markStuffAsCheckedAndSendReport()
                            isLoading = false
                        }
                    }, label: {
                        Text("Send Report")
                            .padding()
                            .frame(maxWidth: .infinity)
                            .background(Color.blue)
                            .foregroundColor(.white)
                            .cornerRadius(10)
                    })
                    .disabled(isLoading)
                    .padding()
                }

                if let errorMessage = errorMessage {
                    Text(errorMessage)
                        .foregroundColor(.red)
                }
            }
        }
        .fullScreenCover(isPresented: $showSheet) {
            ImagePicker(sourceType: $sourceType, selectedImage: createImagePickerBinding())
        }
//        .navigationBarBackButtonHidden(true)
        .onAppear {
            inventoryViewModel.selectStuff(selectedStuff)
        }
    }

    private func createImagePickerBinding() -> Binding<UIImage?> {
        return Binding(
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
        let actionSheet = UIAlertController(title: "Select Image Source", message: nil, preferredStyle: .actionSheet)
        actionSheet.addAction(UIAlertAction(title: "Take Photo", style: .default, handler: { _ in
            self.sourceType = .camera
            self.showSheet.toggle()
        }))
        actionSheet.addAction(UIAlertAction(title: "Choose from Library", style: .default, handler: { _ in
            self.sourceType = .photoLibrary
            self.showSheet.toggle()
        }))
        actionSheet.addAction(UIAlertAction(title: "Cancel", style: .cancel, handler: nil))
        if let windowScene = UIApplication.shared.connectedScenes.first as? UIWindowScene,
           let rootViewController = windowScene.windows.first?.rootViewController {
            rootViewController.present(actionSheet, animated: true, completion: nil)
        }
    }

    private func markStuffAsCheckedAndSendReport() async {
        do {
            try await inventoryViewModel.markStuffAsChecked(selectedStuff)

            await inventoryViewModel.fetchLastInventoryReport()

            if let oldReportId = inventoryViewModel.lastReportId {
                try await inventoryViewModel.compareStuffReport(oldReportId: oldReportId)
//                print("Comparison completed successfully")
            } else {
                errorMessage = "No previous inventory report found for comparison."
            }

            isReportSent = true
        } catch {
            errorMessage = "Error: \(error.localizedDescription)"
        }
    }

    private func validateReport() async {
        if let index = inventoryViewModel.selectedInventory.firstIndex(where: { $0.id == selectedStuff.id }) {
            inventoryViewModel.selectedInventory[index].checked = true
            inventoryViewModel.selectedInventory[index].images = inventoryViewModel.selectedImages
            inventoryViewModel.selectedInventory[index].status = inventoryViewModel.selectedStatus
            inventoryViewModel.selectedInventory[index].comment = inventoryViewModel.comment
            inventoryViewModel.updateRoomCheckedStatus()
        }

        if let roomIndex = inventoryViewModel.localRooms.firstIndex(where: { $0.id == inventoryViewModel.selectedRoom?.id }),
           let stuffIndex = inventoryViewModel.localRooms[roomIndex].inventory.firstIndex(where: { $0.id == selectedStuff.id }) {
            inventoryViewModel.localRooms[roomIndex].inventory[stuffIndex].checked = true
            inventoryViewModel.localRooms[roomIndex].inventory[stuffIndex].images = inventoryViewModel.selectedImages
            inventoryViewModel.localRooms[roomIndex].inventory[stuffIndex].status = inventoryViewModel.selectedStatus
            inventoryViewModel.localRooms[roomIndex].inventory[stuffIndex].comment = inventoryViewModel.comment
        }

        presentationMode.wrappedValue.dismiss()
    }
}

// struct InventoryExitEvaluationView_Previews: PreviewProvider {
//    static var previews: some View {
//        let fakeProperty = exampleDataProperty
//        let viewModel = InventoryViewModel(property: fakeProperty)
//        InventoryExitEvaluationView()
//            .environmentObject(viewModel)
//    }
// }
