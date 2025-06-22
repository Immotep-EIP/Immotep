//
//  ReportDamageView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/06/2025.
//

import SwiftUI
import Photos
import AVFoundation

struct ReportDamageView: View {
    @Environment(\.dismiss) var dismiss
    @ObservedObject var viewModel: PropertyViewModel
    let propertyId: String
    let rooms: [PropertyRoomsTenant]
    let leaseId: String?
    let onDamageCreated: (() -> Void)?

    @State private var description = ""
    @State private var selectedPriority = "medium"
    @State private var selectedRoomId: String?
    @State private var showError: Bool = false
    @State private var errorMessage: String?
    @State private var isLoading = false
    @State private var showImagePicker = false
    @State private var sourceType: UIImagePickerController.SourceType = .photoLibrary
    @State private var selectedImages: [UIImage] = []
    @State private var replaceIndex: Int?
    
    private let priorities = ["low", "medium", "high", "urgent"]
    private let maxImages = 5

    var body: some View {
        ZStack {
            NavigationView {
                Form {
                    Section(header: Text("Description")) {
                        TextEditor(text: $description)
                            .frame(height: 100)
                            .padding()
                            .overlay(
                                RoundedRectangle(cornerRadius: 10)
                                    .stroke(Color.gray.opacity(0.2), lineWidth: 1)
                            )
                    }
                    
                    Section(header: Text("Priority".localized())) {
                        Picker("Priority".localized(), selection: $selectedPriority) {
                            ForEach(priorities, id: \.self) { priority in
                                Text(priority.capitalized.localized()).tag(priority)
                            }
                        }
                        .pickerStyle(.segmented)
                    }
                    
                    Section(header: Text("Room".localized())) {
                        if rooms.isEmpty {
                            Text("No rooms available".localized()).foregroundColor(.red)
                        } else {
                            Picker("Room".localized(), selection: $selectedRoomId) {
                                Text("Select a room".localized()).tag(nil as String?)
                                ForEach(rooms) { room in
                                    Text(room.name).tag(room.id as String?)
                                }
                            }
                        }
                    }
                    Section {
                        PicturesSegmentDamage(selectedImages: $selectedImages, showImagePickerOptions: showImagePickerOptions, maxImages: maxImages)
                    }
                }
                .navigationTitle("Report Damage".localized())
                .toolbar {
                    ToolbarItem(placement: .navigationBarLeading) {
                        Button("Cancel".localized()) {
                            dismiss()
                        }
                    }
                    ToolbarItem(placement: .navigationBarTrailing) {
                        Button("Submit".localized()) {
                            isLoading = true
                            Task {
                                await submitDamage()
                                isLoading = false
                            }
                        }
                        .disabled(isLoading || description.isEmpty || selectedRoomId == nil || rooms.isEmpty)
                    }
                }
                .fullScreenCover(isPresented: $showImagePicker) {
                    ImagePicker(sourceType: $sourceType, selectedImage: createImagePickerBinding())
                }
                .onAppear {
                    if !rooms.isEmpty {
                        selectedRoomId = rooms.first!.id
                    }
                }
            }
            .navigationBarBackButtonHidden(true)

            if showError, let message = errorMessage {
                ErrorNotificationView(message: message)
                    .onDisappear {
                        showError = false
                        errorMessage = nil
                    }
            }
        }
    }
    
    private func createImagePickerBinding() -> Binding<UIImage?> {
        return Binding(
            get: { nil },
            set: { image in
                if let image = image {
                    if let index = replaceIndex {
                        selectedImages[index] = image
                    } else if selectedImages.count < maxImages {
                        selectedImages.append(image)
                    }
                    replaceIndex = nil
                }
                showImagePicker = false
            }
        )
    }
    
    private func showImagePickerOptions(replaceIndex: Int?) {
        self.replaceIndex = replaceIndex
        let actionSheet = UIAlertController(title: "Select Image Source", message: nil, preferredStyle: .actionSheet)
        
        actionSheet.addAction(UIAlertAction(title: "Take Photo", style: .default, handler: { _ in
            AVCaptureDevice.requestAccess(for: .video) { granted in
                DispatchQueue.main.async {
                    if granted {
                        self.sourceType = .camera
                        self.showImagePicker = true
                    } else {
                        errorMessage = "Camera access denied. Please enable it in Settings.".localized()
                        showError = true
                    }
                }
            }
        }))
        
        actionSheet.addAction(UIAlertAction(title: "Choose from Library", style: .default, handler: { _ in
            PHPhotoLibrary.requestAuthorization { status in
                DispatchQueue.main.async {
                    if status == .authorized {
                        self.sourceType = .photoLibrary
                        self.showImagePicker = true
                    } else {
                        errorMessage = "Photo library access denied. Please enable it in Settings.".localized()
                        showError = true
                    }
                }
            }
        }))
        
        actionSheet.addAction(UIAlertAction(title: "Cancel", style: .cancel, handler: nil))
        
        if let windowScene = UIApplication.shared.connectedScenes.first as? UIWindowScene,
           let rootViewController = windowScene.windows.first?.rootViewController {
            rootViewController.present(actionSheet, animated: true, completion: nil)
        }
    }
    
    private func submitDamage() async {
        guard let roomId = selectedRoomId else {
            errorMessage = "Please select a room.".localized()
            showError = true
            return
        }
        
        guard let leaseId = leaseId else {
            errorMessage = "No active lease found.".localized()
            showError = true
            return
        }
        
        let base64Images = selectedImages.compactMap { convertUIImageToBase64($0) }
        
        let damageRequest = DamageRequest(
            comment: description,
            priority: selectedPriority,
            roomId: roomId,
            pictures: base64Images.isEmpty ? nil : base64Images
        )
        
        do {
            let token = try await TokenStorage.getValidAccessToken()
            let _ = try await viewModel.createDamage(
                propertyId: propertyId,
                leaseId: leaseId,
                damage: damageRequest,
                token: token
            )
            await MainActor.run {
                onDamageCreated?()
                dismiss()
            }
        } catch {
            await MainActor.run {
                errorMessage = "Error submitting damage: \(error.localizedDescription)".localized()
                showError = true
            }
        }
    }
    
    private func convertUIImageToBase64(_ image: UIImage) -> String? {
        guard let imageData = image.jpegData(compressionQuality: 0.8) else {
            return nil
        }
        let base64String = imageData.base64EncodedString()
        let dataURI = "data:image/jpeg;base64,\(base64String)"
        return dataURI
    }
}

struct PicturesSegmentDamage: View {
    @Binding var selectedImages: [UIImage]
    var showImagePickerOptions: (Int?) -> Void
    let maxImages: Int
    
    @State private var showImageOptions = false
    @State private var selectedImageIndex: Int?
    
    var body: some View {
        VStack {
            HStack {
                Text(String(format: "Photos (%@/%@)".localized(), "\(selectedImages.count)", "\(maxImages)"))
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
                    if selectedImages.count < maxImages {
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
                }
                .padding(.horizontal)
            }
        }
        .padding(.horizontal)
        .padding(.top)
        .actionSheet(isPresented: $showImageOptions) {
            ActionSheet(
                title: Text("Options"),
                buttons: [
                    .default(Text("Replace".localized())) {
                        if let index = selectedImageIndex {
                            showImagePickerOptions(index)
                        }
                    },
                    .destructive(Text("Delete".localized())) {
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

struct ReportDamageView_Previews: PreviewProvider {
    static var previews: some View {
        NavigationView {
            ReportDamageView(
                viewModel: PropertyViewModel(loginViewModel: LoginViewModel()),
                propertyId: "fakeid",
                rooms: [],
                leaseId: "fakeLeaseId",
                onDamageCreated: nil
            )
            .environment(\.locale, .init(identifier: "en"))
        }
    }
}
