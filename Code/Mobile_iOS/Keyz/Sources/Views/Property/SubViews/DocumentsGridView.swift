//
//  DocumentsGridView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/06/2025.
//

import SwiftUI
import PDFKit
import UniformTypeIdentifiers
import MobileCoreServices

struct DocumentsGridView: View {
    @Binding var documents: [PropertyDocument]
    @EnvironmentObject var propertyViewModel: PropertyViewModel
    @EnvironmentObject var loginViewModel: LoginViewModel
    @State private var errorMessage: String?
    var onDelete: ((String) -> Void)?
    
    var body: some View {
        VStack {
            if let errorMessage = errorMessage {
                Text(errorMessage)
                    .foregroundColor(.red)
                    .padding()
            }
            
            if documents.isEmpty {
                Text("No documents available".localized())
                    .foregroundColor(.gray)
                    .padding()
            } else {
                ScrollView {
                    LazyVGrid(
                        columns: Array(repeating: GridItem(.flexible(), spacing: 15), count: 3),
                        spacing: 15
                    ) {
                        ForEach(documents, id: \.id) { document in
                            ZStack(alignment: .topTrailing) {
                                NavigationLink(destination: PDFViewer(base64String: document.data)) {
                                    documentThumbnailView(document: document)
                                }
                                
                                if loginViewModel.userRole == "owner" {
                                    Menu {
                                        Button(role: .destructive) {
                                            onDelete?(document.id)
                                        } label: {
                                            Label("Delete".localized(), systemImage: "trash")
                                        }
                                    } label: {
                                        Image(systemName: "ellipsis")
                                            .font(.caption)
                                            .foregroundColor(.gray)
                                            .padding(6)
                                            .background(Color.white.opacity(0.8))
                                            .clipShape(Circle())
                                    }
                                    .padding([.top, .trailing], 4)
                                    .accessibilityLabel("document_options_\(document.id)")
                                }
                            }
                        }
                    }
                    .padding()
                }
            }
        }
    }
    
    @ViewBuilder
    private func documentThumbnailView(document: PropertyDocument) -> some View {
        VStack {
            Image(systemName: "doc.text")
                .resizable()
                .scaledToFit()
                .frame(width: 50, height: 50)
            Text(document.title)
                .font(.caption)
                .lineLimit(3)
                .truncationMode(.tail)
                .multilineTextAlignment(.center)
                .fixedSize(horizontal: false, vertical: true)
        }
        .padding()
        .frame(width: 100, height: 150)
        .background(Color.gray.opacity(0.1))
        .cornerRadius(8)
    }
}

struct DocumentPicker: UIViewControllerRepresentable {
    var onDocumentPicked: (URL) -> Void
    
    func makeUIViewController(context: Context) -> UIDocumentPickerViewController {
        let picker = UIDocumentPickerViewController(forOpeningContentTypes: [
            .pdf,
            UTType(filenameExtension: "docx") ?? .data,
            UTType(filenameExtension: "xlsx") ?? .data
        ])
        picker.delegate = context.coordinator
        return picker
    }
    
    func updateUIViewController(_ uiViewController: UIDocumentPickerViewController, context: Context) {}
    
    func makeCoordinator() -> Coordinator {
        Coordinator(self)
    }
    
    class Coordinator: NSObject, UIDocumentPickerDelegate {
        let parent: DocumentPicker
        
        init(_ parent: DocumentPicker) {
            self.parent = parent
        }
        
        func documentPicker(_ controller: UIDocumentPickerViewController, didPickDocumentsAt urls: [URL]) {
            if let url = urls.first {
                parent.onDocumentPicked(url)
            }
        }
        
        func documentPickerWasCancelled(_ controller: UIDocumentPickerViewController) {}
    }
}

struct DocumentsGridView_Previews: PreviewProvider {
    static var previews: some View {
        DocumentsGridView(
            documents: .constant([
                PropertyDocument(id: "doc1", title: "Lease Agreement", fileName: "test.pdf", data: ""),
                PropertyDocument(id: "doc2", title: "Inventory Report", fileName: "inventory.pdf", data: "")
            ])
        )
        .environmentObject(PropertyViewModel(loginViewModel: LoginViewModel()))
        .environmentObject(LoginViewModel())
    }
}
