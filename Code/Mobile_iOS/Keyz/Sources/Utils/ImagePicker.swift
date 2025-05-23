//
//  ImagePicker.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 30/11/2024.
//

import SwiftUI

struct ImagePicker: UIViewControllerRepresentable {
    @Environment(\.presentationMode) private var presentationMode
    @Binding var sourceType: UIImagePickerController.SourceType
    @Binding var selectedImage: UIImage?

    func makeUIViewController(context: UIViewControllerRepresentableContext<ImagePicker>) -> UIImagePickerController {
        let imagePicker = UIImagePickerController()
        imagePicker.allowsEditing = false
        imagePicker.sourceType = sourceType
        imagePicker.delegate = context.coordinator

        imagePicker.modalPresentationStyle = .fullScreen

        return imagePicker
    }

    func updateUIViewController(_ uiViewController: UIImagePickerController, context: UIViewControllerRepresentableContext<ImagePicker>) {
    }

    func makeCoordinator() -> Coordinator {
        Coordinator(self)
    }

    final class Coordinator: NSObject, UIImagePickerControllerDelegate, UINavigationControllerDelegate {

        var parent: ImagePicker

        init(_ parent: ImagePicker) {
            self.parent = parent
        }

        func imagePickerController(_ picker: UIImagePickerController, didFinishPickingMediaWithInfo info: [UIImagePickerController.InfoKey: Any]) {

            if let image = info[UIImagePickerController.InfoKey.originalImage] as? UIImage {
                parent.selectedImage = image
            }

            parent.presentationMode.wrappedValue.dismiss()
        }

    }
}

func convertUIImagesToBase64(_ images: [UIImage]) -> [String] {
    return images.compactMap { image in
        guard let jpegData = image.jpegData(compressionQuality: 0.8) else {
            return nil
        }
        return jpegData.base64EncodedString()
    }
}

func convertUIImageToBase64(_ image: UIImage) -> String {
    guard let imageData = image.jpegData(compressionQuality: 0.8) else {
        return ""
    }
    return imageData.base64EncodedString()
}

func convertBase64ToUIImage(_ base64: String) -> UIImage? {
    guard let data = Data(base64Encoded: base64) else {
        return nil
    }
    return UIImage(data: data)
}

