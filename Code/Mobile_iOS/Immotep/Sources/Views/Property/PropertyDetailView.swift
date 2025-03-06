//
//  PropertyDetailView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/11/2024.
//

import SwiftUI
import PDFKit

struct PropertyDetailView: View {
    @Binding var property: Property
    @ObservedObject var viewModel: PropertyViewModel

    var body: some View {
        VStack(spacing: 0) {
            TopBar(title: "Property Details".localized())

            Form {
                PropertyCardView(property: property)
                    .padding(.vertical, 4)

                Section(header: Text("About the property".localized())) {
                    AboutCardView(property: $property)
                }

                Section(header: Text("Documents").accessibilityLabel("documents_header")) {
                    DocumentsGrid(documents: $property.documents)
                }
            }

            NavigationLink {
                InventoryTypeView(property: $property)
            } label: {
                Text("Start Inventory".localized())
            }
            .frame(maxWidth: .infinity)
            .padding(.vertical, 10)
            .background(.blue)
            .cornerRadius(10)
            .padding()
            .foregroundStyle(.white)
            .accessibilityLabel("inventory_btn_start")
        }
        .navigationBarBackButtonHidden(true)
        .onAppear {
            Task {
                if !CommandLine.arguments.contains("-skipLogin") {
                    do {
                        try await viewModel.fetchPropertyDocuments(propertyId: property.id)
                        // Synchroniser property avec viewModel.properties
                        if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                            property = updatedProperty
                        }
                    } catch {
                        print("Error fetching documents: \(error.localizedDescription)")
                    }
                }
            }
        }
    }
}

struct AboutCardView: View {
    @Binding var property: Property

    var body: some View {
        Grid(alignment: .leading, horizontalSpacing: 10, verticalSpacing: 10) {
            buildRow(
                icon: "person",
                leftText: property.tenantName ?? "No tenant assigned".localized(),
                rightIcon: "square.split.bottomrightquarter",
                rightText: String(format: "area".localized(), formattedValue(property.surface))
            )

            buildRow(
                icon: "calendar",
                leftText:
                    String(format: "start_date".localized(),
                           property.leaseStartDate?.formatted(.dateTime.day().month().year()) ?? "No start date assigned".localized()),
                rightIcon: "coloncurrencysign.arrow.trianglehead.counterclockwise.rotate.90",
                rightText:
                    String(
                        format: "rent_month".localized(),
                        property.monthlyRent)
            )

            buildRow(
                icon: "calendar",
                leftText:
                    String(
                        format: "end_date".localized(), property.leaseEndDate?.formatted(.dateTime.day().month().year()) ??
                        "No end date assigned".localized()),
                rightIcon: "eurosign.bank.building",
                rightText:
                    String(
                        format: "deposit_value".localized(),
                        property.deposit)
            )
        }
        .padding(.vertical, 10)
    }

    private func buildRow(icon: String, leftText: String, rightIcon: String, rightText: String) -> some View {
        GridRow {
            buildHStack(icon: icon, text: leftText)
            buildHStack(icon: rightIcon, text: rightText)
        }
        .padding(.vertical, 10)
    }

    private func buildHStack(icon: String, text: String) -> some View {
        HStack {
            Image(systemName: icon)
            Text(text)
                .lineLimit(nil)
                .fixedSize(horizontal: false, vertical: true)
                .font(.system(size: 14))
        }
        .frame(maxWidth: .infinity, alignment: .leading)
    }

    private func formattedValue(_ value: Double) -> String {
        value == Double(Int(value)) ? String(format: "%.0f", value) : String(format: "%.2f", value)
    }
}

struct DocumentsGrid: View {
    @Binding var documents: [PropertyDocument]

    var body: some View {
        LazyVGrid(
            columns: Array(repeating: GridItem(.flexible(), spacing: 15), count: 3),
            spacing: 15
        ) {
            ForEach(documents) { document in
                NavigationLink(destination: PDFViewer(base64String: document.data)) {
                    VStack {
                        Image(systemName: "text.document")
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

struct PDFViewer: View {
    let base64String: String

    var body: some View {
        if let pdfData = pdfData(from: base64String),
           let pdfDocument = PDFDocument(data: pdfData) {
            PDFKitView(pdfDocument: pdfDocument)
        } else {
            Text("Unable to load PDF")
        }
    }

    private func pdfData(from base64String: String) -> Data? {
        // Supprimer le préfixe "data:application/pdf;base64,"
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

// struct PropertyDetailView_Previews: PreviewProvider {
//    static var previews: some View {
//        let property = Property(
//            id: "",
//            ownerID: "",
//            name: "Condo",
//            address: "4391 Hedge Street",
//            city: "New Jersey",
//            postalCode: "07102",
//            country: "USA",
//            photo: nil,
//            monthlyRent: 1200,
//            deposit: 2400,
//            surface: 80.0,
//            isAvailable: false,
//            tenantName: "John & Mary Doe",
//            leaseStartDate: Date(),
//            leaseEndDate: Calendar.current.date(byAdding: .year, value: 1, to: Date()),
//            documents: [
//                PropertyDocument(id: UUID(), title: "Lease Agreement", fileName: "lease_agreement.pdf"),
//                PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
//                PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
//                PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
//                PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
//                PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf")
//
//            ],
//            rooms: []
//        )
//        PropertyDetailView(property: .constant(property), viewModel: viewModel)
//    }
// }
