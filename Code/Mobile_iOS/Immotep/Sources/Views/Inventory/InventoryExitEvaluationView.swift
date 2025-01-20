//
//  InventoryExitEvaluationView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 19/01/2025.
//

import SwiftUI

struct InventoryExitEvaluationView: View {
    @EnvironmentObject var inventoryViewModel: InventoryViewModel

    @State private var showSheet = false
    @State private var sourceType: UIImagePickerController.SourceType = .photoLibrary
    @State private var replaceIndex: Int?

    var body: some View {
        NavigationView {
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
                                ForEach(["Available", "Unavailable", "Maintenance", "Retired"], id: \.self) { status in
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
                }

                TaskBar()
            }
            .fullScreenCover(isPresented: $showSheet) {
                ImagePicker(sourceType: $sourceType, selectedImage: createImagePickerBinding())
            }
        }
        .navigationBarBackButtonHidden(true)
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
}

struct InventoryExitEvaluationView_Previews: PreviewProvider {
    static var previews: some View {
        let fakeProperty = exampleDataProperty
        let viewModel = InventoryViewModel(property: fakeProperty)
        InventoryExitEvaluationView()
            .environmentObject(viewModel)
    }
}
