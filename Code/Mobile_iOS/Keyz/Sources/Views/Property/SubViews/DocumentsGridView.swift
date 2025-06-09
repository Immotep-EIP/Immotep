//
//  DocumentsGridView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/06/2025.
//

import SwiftUI
import PDFKit

struct DocumentsGridView: View {
    @Binding var documents: [PropertyDocument]
    @State private var refreshID = UUID()
    @State private var isInitialized = false

    var body: some View {
        VStack {
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
}

struct DocumentsGridView_Previews: PreviewProvider {
    static var previews: some View {
        DocumentsGridView(documents: .constant([
            PropertyDocument(id: "doc1", title: "Lease Agreement", fileName: "test", data: "base64_data")
        ]))
    }
}
