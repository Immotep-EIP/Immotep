//
//  EditPropertyView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 22/02/2025.
//

import SwiftUI

struct EditPropertyView: View {
    @Environment(\.dismiss) private var dismiss
    @ObservedObject var viewModel: PropertyViewModel
    @Binding var property: Property

    @State private var name: String
    @State private var address: String
    @State private var city: String
    @State private var postalCode: String
    @State private var country: String
    @State private var photo: UIImage?
    @State private var monthlyRent: NSNumber?
    @State private var deposit: NSNumber?
    @State private var surface: NSNumber?
    @State private var errorMessage: String?

    @State private var showSheet = false
    @State private var sourceType: UIImagePickerController.SourceType = .photoLibrary

    init(viewModel: PropertyViewModel, property: Binding<Property>) {
        self.viewModel = viewModel
        self._property = property
        self._name = State(initialValue: property.wrappedValue.name)
        self._address = State(initialValue: property.wrappedValue.address)
        self._city = State(initialValue: property.wrappedValue.city)
        self._postalCode = State(initialValue: property.wrappedValue.postalCode)
        self._country = State(initialValue: property.wrappedValue.country)
        self._photo = State(initialValue: property.wrappedValue.photo)
        self._monthlyRent = State(initialValue: NSNumber(value: property.wrappedValue.monthlyRent))
        self._deposit = State(initialValue: NSNumber(value: property.wrappedValue.deposit))
        self._surface = State(initialValue: NSNumber(value: property.wrappedValue.surface))
    }

    var body: some View {
        VStack {
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

            if let errorMessage = errorMessage {
                Text(errorMessage)
                    .foregroundColor(.red)
                    .padding()
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
                Button("Save Changes".localized()) {
                    Task {
                        await updateProperty()
                    }
                }
                .padding(.horizontal, 25)
                .padding(.vertical, 8)
                .background(Color.blue)
                .foregroundStyle(Color.white)
                .font(.headline)
                .cornerRadius(8)
                .accessibilityIdentifier("save_button")

                Spacer()
            }
        }
        .navigationBarBackButtonHidden(true)
        .fullScreenCover(isPresented: $showSheet) {
            ImagePicker(sourceType: $sourceType, selectedImage: $photo)
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

    private func updateProperty() async {
        let updatedProperty = Property(
            id: property.id,
            ownerID: property.ownerID,
            name: name,
            address: address,
            city: city,
            postalCode: postalCode,
            country: country,
            photo: photo,
            monthlyRent: monthlyRent?.intValue ?? 0,
            deposit: deposit?.intValue ?? 0,
            surface: surface?.doubleValue ?? 0.0,
            isAvailable: property.isAvailable,
            tenantName: property.tenantName,
            leaseId: property.leaseId,
            leaseStartDate: property.leaseStartDate,
            leaseEndDate: property.leaseEndDate,
            documents: property.documents,
            createdAt: property.createdAt,
            rooms: property.rooms,
            damages: []
        )

        guard let token = await TokenStorage.getAccessToken() else {
            errorMessage = "Failed to retrieve token.".localized()
            print("Token is nil")
            return
        }

        do {
            let propertyId = try await viewModel.updateProperty(request: updatedProperty, token: token)
            print("Property updated successfully with ID: \(propertyId)")

            if let newPhoto = photo, newPhoto != property.photo {
                do {
                    let res = try await viewModel.updatePropertyPicture(token: token, propertyPicture: newPhoto, propertyID: propertyId)
                    print("Property picture updated successfully: \(res)")
                } catch {
                    print("Failed to update property picture: \(error.localizedDescription)")
                }
            }

            if let updatedProperty = viewModel.properties.first(where: { $0.id == propertyId }) {
                property = updatedProperty
                print("Updated property binding in EditPropertyView for ID: \(propertyId)")
            } else {
                errorMessage = "Updated property not found in refreshed data.".localized()
                print("Property \(propertyId) not found in viewModel.properties after update")
            }

            dismiss()
        } catch {
            errorMessage = "Error updating property: \(error.localizedDescription)".localized()
            print("Error updating property: \(error)")
        }
    }
}

struct EditPropertyView_Previews: PreviewProvider {
    static var viewModel = PropertyViewModel(loginViewModel: LoginViewModel())
    static var previews: some View {
        EditPropertyView(viewModel: viewModel, property: .constant(exampleDataProperty))
    }
}
