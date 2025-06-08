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

    var body: some View {
        LazyVGrid(
            columns: Array(repeating: GridItem(.flexible(), spacing: 15), count: 3),
            spacing: 15
        ) {
            ForEach(documents) { document in
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
    }
}

struct DocumentsGridView_Previews: PreviewProvider {
    static var previews: some View {
        DocumentsGridView(documents: .constant([
            PropertyDocument(id: "doc1", title: "Lease Agreement", fileName: "test", data: "base64_data")
        ]))
    }
}
