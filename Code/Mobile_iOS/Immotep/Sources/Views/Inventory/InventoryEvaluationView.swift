//
//  InventoryEvaluationView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI
import UIKit
import Foundation

struct InventoryEntryEvaluationView: View {
    @EnvironmentObject var inventoryViewModel: InventoryViewModel
    let selectedStuff: RoomInventory

    @State private var showSheet = false
    @State private var sourceType: UIImagePickerController.SourceType = .photoLibrary
    @State private var replaceIndex: Int?
    @State private var isLoading: Bool = false
    @State private var errorMessage: String?
    let stateMapping: [String: String] = [
        "not_set": "Select your equipment status",
        "broken": "Broken",
        "bad": "Bad",
        "good": "Good",
        "new": "New"
    ]

    var body: some View {
        NavigationView {
            VStack(spacing: 0) {
                TopBar(title: "Inventory")

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

                    Button(action: {
                        Task {
                            isLoading = true
                            await markStuffAsCheckedAndSendReport()
                            isLoading = false
                        }
                    }, label: {
                        Text("Validate")
                    })
                    .disabled(isLoading)

                    if let errorMessage = errorMessage {
                        Text(errorMessage)
                            .foregroundColor(.red)
                    }
                }

                TaskBar()
            }
            .fullScreenCover(isPresented: $showSheet) {
                ImagePicker(sourceType: $sourceType, selectedImage: createImagePickerBinding())
            }
        }
        .navigationBarBackButtonHidden(true)
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
            try await inventoryViewModel.sendStuffReport()
        } catch {
            errorMessage = "Error: \(error.localizedDescription)"
        }
    }
}

struct PicturesSegment: View {
    @Binding var selectedImages: [UIImage]
    var showImagePickerOptions: (Int?) -> Void

    @State private var showImageOptions = false
    @State private var selectedImageIndex: Int?

    var body: some View {
        VStack {
            HStack {
                Text("Picture(s)")
                    .font(.headline)
                Spacer()
            }
            ScrollView(.horizontal, showsIndicators: false) {
                HStack(spacing: 10) {
                    ForEach(selectedImages.indices, id: \.self) { index in
                        Image(uiImage: selectedImages[index])
                            .resizable()
                            .scaledToFill()
                            .frame(width: 100, height: 100)
                            .clipShape(RoundedRectangle(cornerRadius: 10))
                            .onTapGesture {
                                selectedImageIndex = index
                                showImageOptions = true
                            }
                    }
                    Button(action: {
                        showImagePickerOptions(nil)
                    }, label: {
                        ZStack {
                            Rectangle()
                                .fill(Color.gray.opacity(0.2))
                                .frame(width: 100, height: 100)
                                .clipShape(RoundedRectangle(cornerRadius: 10))
                            Image(systemName: "plus")
                                .font(.largeTitle)
                                .foregroundColor(.gray)
                        }
                    })
                }
                .padding(.horizontal)
            }
            Divider()
                .frame(maxWidth: .infinity)
                .background(Color("textColor"))
                .padding()
        }
        .padding(.horizontal)
        .padding(.top)
        .actionSheet(isPresented: $showImageOptions) {
            ActionSheet(
                title: Text("Options"),
                buttons: [
                    .default(Text("Replace")) {
                        if let index = selectedImageIndex {
                            showImagePickerOptions(index)
                        }
                    },
                    .destructive(Text("Delete")) {
                        if let index = selectedImageIndex {
                            selectedImages.remove(at: index)
                        }
                    },
                    .cancel()
                ]
            )
        }
        .navigationBarBackButtonHidden(true)
    }
}
