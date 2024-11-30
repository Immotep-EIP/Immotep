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
    @StateObject private var keyboardObserver = KeyboardObserver()

    @State private var address = ""
    @State private var postalCode = ""
    @State private var country = ""
    @State private var photo = UIImage(named: "DefaultImageProperty") ?? UIImage()
    @State private var monthlyRent: NSNumber?
    @State private var deposit: NSNumber?
    @State private var surface: NSNumber?

    @State private var image = UIImage(named: "DefaultImageProperty") ?? UIImage()
    @State private var showSheet = false
    @State private var sourceType: UIImagePickerController.SourceType = .photoLibrary

    var body: some View {
        VStack {
            TopBar(title: "New Property")

            Form {
                Section {
                    VStack {
                        Image(uiImage: self.image)
                            .resizable()
                            .scaledToFill()
                            .frame(width: 100, height: 100)
                            .clipShape(Circle())
                            .overlay(Circle().stroke(Color.black, lineWidth: 1))
                            .padding(.top, 10)
                            .onTapGesture {
                                showImagePickerOptions()
                            }

                        Text("Click on the image to change")
                            .font(.subheadline)
                            .foregroundColor(.gray)
                            .padding(.top, 8)
                    }
                    .frame(maxWidth: .infinity)
                    CustomTextInput(title: "Address", placeholder: "Enter address", text: $address, isSecure: false)
                    CustomTextInput(title: "Postal Code", placeholder: "Enter postal code", text: $postalCode, isSecure: false)
                    CustomTextInput(title: "Country", placeholder: "Enter country", text: $country, isSecure: false)
                    CustomTextInputNB(title: "Monthly Rent", placeholder: "Enter monthly rent", value: $monthlyRent, isSecure: false)
                    CustomTextInputNB(title: "Deposit", placeholder: "Enter deposit", value: $deposit, isSecure: false)
                    CustomTextInputNB(title: "Surface (m²)", placeholder: "Enter surface", value: $surface, isSecure: false)
                }
            }

            HStack {
                Spacer()
                Button("Cancel") {
                    dismiss()
                }
                .padding(.horizontal, 25)
                .padding(.vertical, 8)
                .background(Color.red)
                .foregroundStyle(Color.white)
                .font(.headline)
                .cornerRadius(8)
                Spacer()
                Button("Add Property") {
                    let newProperty = Property(
                        id: UUID(),
                        address: address,
                        postalCode: postalCode,
                        country: country,
                        photo: photo,
                        monthlyRent: monthlyRent?.doubleValue ?? 0.0,
                        deposit: deposit?.doubleValue ?? 0.0,
                        surface: surface?.doubleValue ?? 0.0,
                        isAvailable: true,
                        tenantName: nil,
                        leaseStartDate: nil,
                        leaseEndDate: nil,
                        documents: []
                    )
                    viewModel.addProperty(newProperty)
                    dismiss()
                }
                .padding(.horizontal, 25)
                .padding(.vertical, 8)
                .background(Color.blue)
                .foregroundStyle(Color.white)
                .font(.headline)
                .cornerRadius(8)
                Spacer()
            }

            if !keyboardObserver.isKeyboardVisible {
                TaskBar()
            }
        }
        .navigationBarBackButtonHidden(true)
        .fullScreenCover(isPresented: $showSheet) {
            ImagePicker(sourceType: $sourceType, selectedImage: $photo)
        }
    }

    private func showImagePickerOptions() {
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

struct CreatePropertyView_Previews: PreviewProvider {
    static var viewModel = PropertyViewModel()
    static var previews: some View {
        CreatePropertyView(viewModel: viewModel)
    }
}
