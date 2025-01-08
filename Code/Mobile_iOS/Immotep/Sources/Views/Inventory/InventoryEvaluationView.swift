//
//  InventoryEvaluationView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/12/2024.
//

import SwiftUI

struct InventoryEvaluationView: View {
    @Binding var stuff: RoomInventory

    @State private var selectedImages: [UIImage] = []
    @State private var comment: String = ""
    @State private var selectedStatus: String = "Select your equipment status"
    let statusOptions = ["Available", "Unavailable", "Maintenance", "Retired"]

    @State private var showSheet = false
    @State private var sourceType: UIImagePickerController.SourceType = .photoLibrary
    @State private var replaceIndex: Int?

    var body: some View {
        NavigationView {
            VStack(spacing: 0) {
                TopBar(title: "Inventory")

                ScrollView {
                    Section {
                        PicturesSegment(selectedImages: $selectedImages, showImagePickerOptions: showImagePickerOptions)
                    }

                    VStack {
                        HStack {
                            Text("Comments")
                                .font(.headline)
                            Spacer()
                        }
                        TextEditor(text: $comment)
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
                            Picker("Select Equipment Status", selection: $selectedStatus) {
                                Text("Select your equipment status").tag("Select your equipment status")
                                ForEach(statusOptions, id: \.self) { status in
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
                ImagePicker(sourceType: $sourceType, selectedImage: Binding(
                    get: { nil },
                    set: { image in
                        if let image = image {
                            if let index = replaceIndex {
                                selectedImages[index] = image
                            } else {
                                selectedImages.append(image)
                            }
                        }
                    }
                ))
            }
        }
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
    }
}

struct InventoryEvaluationView_Previews: PreviewProvider {
    static var previews: some View {
        let fakeProperty = exampleDataProperty
        InventoryEvaluationView(stuff: .constant(fakeProperty.rooms[1].inventory[1]))
    }
}
