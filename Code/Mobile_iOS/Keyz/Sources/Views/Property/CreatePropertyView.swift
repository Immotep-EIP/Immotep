//
//  CreatePropertyView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/11/2024.
//

import SwiftUI

struct CreatePropertyView: View {
    @Environment(\.dismiss) private var dismiss
    @ObservedObject var viewModel: PropertyViewModel

    @State private var name = ""
    @State private var address = ""
    @State private var city = ""
    @State private var postalCode = ""
    @State private var country = ""
    @State private var photo: UIImage? = UIImage(named: "DefaultImageProperty")
    @State private var monthlyRent: NSNumber?
    @State private var deposit: NSNumber?
    @State private var surface: NSNumber?
    @State private var showError: Bool = false
    @State private var errorMessage: String?
    @State private var showSheet = false
    @State private var sourceType: UIImagePickerController.SourceType = .photoLibrary

    var body: some View {
        ZStack {
            VStack {
                TopBar(title: "New Property".localized())

                Form {
                    Section {
                        VStack {
                            Image(uiImage: photo ?? UIImage(named: "DefaultImageProperty")!)
                                .resizable()
                                .scaledToFill()
                                .frame(width: 100, height: 100)
                                .clipShape(Circle())
                                .overlay(Circle().stroke(Color.black, lineWidth: 1))
                                .padding(.top, 10)
                                .onTapGesture {
                                    showImagePickerOptions()
                                }
                                .accessibilityIdentifier("image_property")

                            Text("Click on the image to change".localized())
                                .font(.subheadline)
                                .foregroundColor(.gray)
                                .padding(.top, 8)
                                .accessibilityIdentifier("touch_to_change_image")
                        }
                        .frame(maxWidth: .infinity)
                        CustomTextInput(title: "Name", placeholder: "Enter property name", text: $name, isSecure: false)
                        CustomTextInput(title: "Address", placeholder: "Enter address", text: $address, isSecure: false)
                        CustomTextInput(title: "City", placeholder: "Enter city", text: $city, isSecure: false)
                        CustomTextInput(title: "Postal Code", placeholder: "Enter postal code", text: $postalCode, isSecure: false)
                        CustomTextInput(title: "Country", placeholder: "Enter country", text: $country, isSecure: false)
                        CustomTextInputNB(title: "Monthly Rent", placeholder: "Enter monthly rent", value: $monthlyRent, isSecure: false)
                        CustomTextInputNB(title: "Deposit", placeholder: "Enter deposit", value: $deposit, isSecure: false)
                        CustomTextInputNB(title: "Surface (m²)", placeholder: "Enter surface", value: $surface, isSecure: false)
                    }
                }
                HStack {
                    Spacer()
                    Button("Cancel".localized()) {
                        dismiss()
                    }
                    .padding(.horizontal, 25)
                    .padding(.vertical, 8)
                    .background(Color.red)
                    .foregroundStyle(Color.white)
                    .font(.headline)
                    .cornerRadius(8)
                    .accessibilityIdentifier("cancel_button")

                    Spacer()
                    Button("Add Property".localized()) {
                        Task {
                            await addProperty()
                        }
                    }
                    .padding(.horizontal, 25)
                    .padding(.vertical, 8)
                    .background(Color.blue)
                    .foregroundStyle(Color.white)
                    .font(.headline)
                    .cornerRadius(8)
                    .accessibilityIdentifier("confirm_button")

                    Spacer()
                }
            }
            .navigationBarBackButtonHidden(true)
            .fullScreenCover(isPresented: $showSheet) {
                ImagePicker(sourceType: $sourceType, selectedImage: $photo)
            }

            if showError, let message = errorMessage {
                ErrorNotificationView(message: message)
                    .onDisappear {
                        showError = false
                        errorMessage = nil
                    }
            }
        }
    }

    private func showImagePickerOptions() {
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

    private func addProperty() async {
        let newProperty = Property(
            id: "",
            ownerID: "",
            name: name,
            address: address,
            city: city,
            postalCode: postalCode,
            country: country,
            photo: photo,
            monthlyRent: monthlyRent?.intValue ?? 0,
            deposit: deposit?.intValue ?? 0,
            surface: surface?.doubleValue ?? 0.0,
            isAvailable: "available",
            tenantName: nil,
            leaseId: nil,
            leaseStartDate: nil,
            leaseEndDate: nil,
            documents: [],
            createdAt: ISO8601DateFormatter().string(from: Date()),
            rooms: [],
            damages: []
        )

        guard let token = await TokenStorage.getAccessToken() else {
            errorMessage = "Failed to retrieve access token.".localized()
            showError = true
            return
        }

        do {
            let response = try await viewModel.createProperty(request: newProperty, token: token)
            await MainActor.run {
                errorMessage = "Property created successfully with ID: \(response)!".localized()
                showError = true
                DispatchQueue.main.asyncAfter(deadline: .now() + 5) {
                    dismiss()
                }
            }
            await viewModel.fetchProperties()
        } catch {
            await MainActor.run {
                if let nsError = error as NSError?, nsError.code == 409 {
                    errorMessage = "A property with this information already exists.".localized()
                } else {
                    errorMessage = "Error creating property: \(error.localizedDescription)".localized()
                }
                showError = true
            }
        }
    }
}

struct CreatePropertyView_Previews: PreviewProvider {
    static var viewModel = PropertyViewModel(loginViewModel: LoginViewModel())
    static var previews: some View {
        CreatePropertyView(viewModel: viewModel)
    }
}
