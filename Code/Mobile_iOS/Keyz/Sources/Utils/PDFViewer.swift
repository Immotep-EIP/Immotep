//
//  PDFViewer.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 23/04/2025.
//

import SwiftUI
import PDFKit

struct PDFViewer: View {
    let base64String: String

    var body: some View {
        if let pdfData = pdfData(from: base64String),
           let pdfDocument = PDFDocument(data: pdfData) {
            PDFKitView(pdfDocument: pdfDocument)
        } else {
            Text("Unable to load PDF".localized())
        }
    }

    private func pdfData(from base64String: String) -> Data? {
        let base64Content = base64String.replacingOccurrences(of: "data:application/pdf;base64,", with: "")
        return Data(base64Encoded: base64Content)
    }
}

struct PDFKitView: UIViewRepresentable {
    let pdfDocument: PDFDocument

    func makeUIView(context: Context) -> PDFView {
        let pdfView = PDFView()
        pdfView.document = pdfDocument
        pdfView.autoScales = true
        return pdfView
    }

    func updateUIView(_ uiView: PDFView, context: Context) {
        uiView.document = pdfDocument
    }
}
