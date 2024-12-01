//
//  PropertyDetailView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 25/11/2024.
//

import SwiftUI

struct PropertyDetailView: View {
    @Binding var property: Property
    @StateObject private var keyboardObserver = KeyboardObserver()

    var body: some View {
        VStack(spacing: 0) {
            TopBar(title: "Property Details")

            Form {
                PropertyCardView(property: property)
                    .padding(.vertical, 4)

                Section(header: Text("About the property")) {
                    AboutCardView(property: $property)
                }

                Section(header: Text("Documents")) {
                    DocumentsGrid(documents: $property.documents)
                }
            }
            Button("Start Inventory") {
                // Action to define
            }
            .frame(maxWidth: .infinity)
            .padding(.vertical, 10)
            .background(.blue)
            .cornerRadius(10)
            .padding()
            .foregroundStyle(.white)

            if !keyboardObserver.isKeyboardVisible {
                TaskBar()
            }
        }
        .navigationBarBackButtonHidden(true)
    }
}

struct AboutCardView: View {
    @Binding var property: Property

    var body: some View {
        Grid(alignment: .leading, horizontalSpacing: 10, verticalSpacing: 10) {
            buildRow(
                icon: "person",
                leftText: property.tenantName ?? "No tenant assigned",
                rightIcon: "square.split.bottomrightquarter",
                rightText: "Area: \(formattedValue(property.surface)) m²"
            )

            buildRow(
                icon: "calendar",
                leftText: "Start: \(property.leaseStartDate?.formatted(.dateTime.day().month().year()) ?? "No start date assigned")",
                rightIcon: "coloncurrencysign.arrow.trianglehead.counterclockwise.rotate.90",
                rightText: "Rent / Months: \(property.monthlyRent)€"
            )

            buildRow(
                icon: "calendar",
                leftText: "End: \(property.leaseEndDate?.formatted(.dateTime.day().month().year()) ?? "No end date assigned")",
                rightIcon: "eurosign.bank.building",
                rightText: "Deposit: \(property.deposit)€"
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
        .padding()
    }
}

struct PropertyDetailView_Previews: PreviewProvider {
    static var previews: some View {
        let property = Property(
            id: "",
            ownerID: "",
            name: "Condo",
            address: "4391 Hedge Street",
            city: "New Jersey",
            postalCode: "07102",
            country: "USA",
            photo: nil,
            monthlyRent: 1200,
            deposit: 2400,
            surface: 80.0,
            isAvailable: false,
            tenantName: "John & Mary Doe",
            leaseStartDate: Date(),
            leaseEndDate: Calendar.current.date(byAdding: .year, value: 1, to: Date()),
            documents: [
                PropertyDocument(id: UUID(), title: "Lease Agreement", fileName: "lease_agreement.pdf"),
                PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
                PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
                PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
                PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf"),
                PropertyDocument(id: UUID(), title: "Inspection Report", fileName: "inspection_report.pdf")

            ]
        )
        PropertyDetailView(property: .constant(property))
    }
}
