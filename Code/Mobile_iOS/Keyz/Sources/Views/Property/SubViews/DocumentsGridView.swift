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
    @State private var refreshID = UUID()
    @State private var isInitialized = false
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
                LazyVGrid(
                    columns: Array(repeating: GridItem(.flexible(), spacing: 15), count: 3),
                    spacing: 15
                ) {
                    ForEach(documents, id: \.id) { document in
                        ZStack(alignment: .topTrailing) {
                            NavigationLink(destination: PDFViewer(base64String: document.data)) {
                                VStack {
                                    Image(systemName: "doc.text")
                                        .resizable()
                                        .scaledToFit()
                                        .frame(width: 50, height: 50)
                                    Text(document.title)
                                        .font(.caption)
                                        .multilineTextAlignment(.center)
                                        .frame(maxWidth: .infinity)
                                }
                                .padding()
                                .frame(maxWidth: .infinity)
                                .background(Color.gray.opacity(0.1))
                                .cornerRadius(8)
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
                .id(refreshID)
            }
        }
        .onChange(of: documents) {
            refreshID = UUID()
        }
        .onAppear {
            isInitialized = true
        }
    }
    
    private func handleDocumentSelection(url: URL) async throws {
        guard let propertyId = propertyViewModel.properties.first?.id else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "No property selected.".localized()])
        }
        
        let allowedTypes: [UTType] = [.pdf, .init(filenameExtension: "docx")!, .init(filenameExtension: "xlsx")!]
        guard let fileType = UTType(filenameExtension: url.pathExtension.lowercased()),
              allowedTypes.contains(fileType) else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid file type. Only PDF, DOCX, and XLSX are supported.".localized()])
        }
        
        let mimeType: String
        switch fileType {
        case .pdf:
            mimeType = "application/pdf"
        case UTType(filenameExtension: "docx"):
            mimeType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
        case UTType(filenameExtension: "xlsx"):
            mimeType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
        default:
            mimeType = "application/octet-stream"
        }
        
        guard url.startAccessingSecurityScopedResource() else {
            throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Unable to access file.".localized()])
        }
        defer { url.stopAccessingSecurityScopedResource() }
        
        let data = try Data(contentsOf: url)
        let base64String = "data:\(mimeType);base64,\(data.base64EncodedString())"
        let fileName = url.lastPathComponent
        
        try await propertyViewModel.uploadDocument(
            propertyId: propertyId,
            fileName: fileName,
            base64Data: base64String
        )
        
        let updatedDocuments = try await propertyViewModel.fetchPropertyDocuments(propertyId: propertyId)
        documents = updatedDocuments
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
        DocumentsGridView(documents: .constant([
            PropertyDocument(id: "doc1", title: "Lease Agreement", fileName: "test.pdf", data: "base64_data")
        ]))
        .environmentObject(PropertyViewModel(loginViewModel: LoginViewModel()))
        .environmentObject(LoginViewModel())
    }
}
