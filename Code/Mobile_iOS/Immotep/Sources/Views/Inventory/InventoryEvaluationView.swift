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
                    
                    Button(action: {
                        Task {
                            isLoading = true
                            await markStuffAsCheckedAndSendReport()
                            isLoading = false
                        }
                    }) {
                        Text("Validate")
                    }
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
            print("markStuffChecked")
            try await sendInventoryReport()
        } catch {
            errorMessage = "Error: \(error.localizedDescription)"
        }
    }
    
    //    private func sendInventoryReport() async throws {
    //        guard let url = URL(string: "\(baseURL)/owner/properties/\(inventoryViewModel.property.id)/inventory-reports/summarize/") else {
    //            throw URLError(.badURL)
    //        }
    //
    //        guard let token = await inventoryViewModel.getToken() else {
    //            throw URLError(.userAuthenticationRequired)
    //        }
    //
    //        var request = URLRequest(url: url)
    //        request.httpMethod = "POST"
    //        request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
    //        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
    //
    //        let body: [String: Any] = [
    //            "type": "furniture",
    //            "id": selectedStuff.id,
    ////            "pictures": inventoryViewModel.selectedImages.map { $0.jpegData(compressionQuality: 0.8)?.base64EncodedString() ?? "" }
    ////            "pictures": convertImageToBase64(image: inventoryViewModel.selectedImages[0]) ?? ""
    //            "pictures" :[ "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"]
    //        ]
    //        print("token: \(token)")
    //        print("body done: \(body)")
    //
    //        request.httpBody = try JSONSerialization.data(withJSONObject: body)
    //
    //        print("request done: \(request)")
    //
    //        let (data, response) = try await URLSession.shared.data(for: request)
    //
    ////        print("data: \(String(data: data, encoding: .utf8) ?? "")")
    ////        print("response: \(response)")
    //
    //        guard let httpResponse = response as? HTTPURLResponse, (200...299).contains(httpResponse.statusCode) else {
    //            print("httpResponse: \(response ?? URLResponse())")
    //            throw URLError(.badServerResponse)
    //        }
    //
    //        print("Report sent successfully: \(String(data: data, encoding: .utf8) ?? "")")
    //    }

    struct SummarizeRequest: Codable {
        let id: String
        let pictures: [String]
        let type: String
    }

    struct SummarizeResponse: Codable {
        let cleanliness: String
        let note: String
        let state: String
    }

    private func sendInventoryReport() async throws {
        guard let url = URL(string: "\(baseURL)/owner/properties/\(inventoryViewModel.property.id)/inventory-reports/summarize/") else {
            throw URLError(.badURL)
        }
        guard let token = await inventoryViewModel.getToken() else {
            throw URLError(.userAuthenticationRequired)
        }
//        let base64Images = inventoryViewModel.selectedImages.compactMap { convertImageToBase64(image: $0) }
//        print(base64Images)
        
        let base64Images = [ "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"]
        
        let body = SummarizeRequest(
            id: selectedStuff.id,
            pictures: base64Images,
            type: "furniture"
        )

        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")

        let encoder = JSONEncoder()
        request.httpBody = try encoder.encode(body)

        let (data, response) = try await URLSession.shared.data(for: request)

        guard let httpResponse = response as? HTTPURLResponse, (200...299).contains(httpResponse.statusCode) else {
            let responseBody = String(data: data, encoding: .utf8) ?? "No response body"
            print("HTTP Status Code: \((response as? HTTPURLResponse)?.statusCode ?? -1)")
            print("Response Body: \(responseBody)")
            throw URLError(.badServerResponse)
        }

        let decoder = JSONDecoder()
        let summarizeResponse = try decoder.decode(SummarizeResponse.self, from: data)
        print("Report sent successfully: \(summarizeResponse)")
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

func scaleImage(image: UIImage, scale: CGFloat) -> UIImage? {
    let newSize = CGSize(width: image.size.width * scale, height: image.size.height * scale)
    let renderer = UIGraphicsImageRenderer(size: newSize)
    return renderer.image { _ in
        image.draw(in: CGRect(origin: .zero, size: newSize))
    }
}

func convertImageToBase64(image: UIImage, scale: CGFloat = 0.2, quality: CGFloat = 0.1) -> String? {
    guard let resizedImage = scaleImage(image: image, scale: scale),
          let imageData = resizedImage.jpegData(compressionQuality: quality) else { return nil }
    return imageData.base64EncodedString()
}
