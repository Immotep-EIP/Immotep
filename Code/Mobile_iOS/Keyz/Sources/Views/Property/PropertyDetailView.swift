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
    @StateObject private var tenantViewModel = TenantViewModel()
    @State private var showInviteTenantSheet = false
    @State private var showEndLeasePopUp = false
    @State private var showCancelInvitePopUp = false
    @State private var showDeletePropertyPopUp = false
    @State private var showEditPropertyPopUp = false
    @Environment(\.dismiss) var dismiss

    var body: some View {
        ZStack {
            VStack(spacing: 0) {
                TopBar(title: "Property Details".localized())
                
                Form {
                    PropertyCardView(property: $property)
                        .padding(.vertical, 4)
                    
                    Section(header: Text("About the property".localized())) {
                        AboutCardView(property: $property)
                    }
                    
                    Section(header: Text("Documents")
                        .accessibilityIdentifier("documents_header")) {
                            DocumentsGrid(documents: $property.documents)
                        }
                }
                
                Menu {
                    Button(action: {
                        showInviteTenantSheet = true
                    }) {
                        Label("Invite Tenant".localized(), systemImage: "person.crop.circle.badge.plus")
                    }
                    
                    Button(action: {
                        showEndLeasePopUp = true
                    }) {
                        Label("End Lease".localized(), systemImage: "xmark.circle")
                    }
                    
                    Button(action: {
                        showCancelInvitePopUp = true
                    }) {
                        Label("Cancel Invite".localized(), systemImage: "person.crop.circle.badge.xmark")
                    }
                    
                    Button(action: {
                        showEditPropertyPopUp = true
                    }) {
                        Label("Edit Property".localized(), systemImage: "pencil")
                    }
                    
                    Button(action: {
                        showDeletePropertyPopUp = true
                    }) {
                        Label("Delete Property".localized(), systemImage: "trash")
                    }
                } label: {
                    Text("Actions")
                        .frame(maxWidth: .infinity)
                        .padding(.vertical, 10)
                        .background(Color.gray.opacity(0.2))
                        .cornerRadius(10)
                        .foregroundStyle(.blue)
                }
                .padding(.top, 20)
                .padding(.horizontal)
                
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
                            if let updatedProperty = viewModel.properties.first(where: { $0.id == property.id }) {
                                property = updatedProperty
                            }
                        } catch {
                            print("Error fetching documents: \(error.localizedDescription)")
                        }
                    }
                }
            }
            .sheet(isPresented: $showInviteTenantSheet) {
                InviteTenantView(tenantViewModel: tenantViewModel, property: property)
            }
            if showCancelInvitePopUp {
                CustomAlertTwoButtons(
                    isActive: $showCancelInvitePopUp,
                    title: "Cancel Invite".localized(),
                    message: "Are you sure you want to cancel the pending invite?".localized(),
                    buttonTitle: "Confirm".localized(),
                    secondaryButtonTitle: "Cancel".localized(),
                    action: {
                        
                    },
                    secondaryAction: {
                    }
                )
                .accessibilityIdentifier("InviteTenantAlert")
            }
            if showEndLeasePopUp {
                CustomAlertTwoButtons(
                    isActive: $showEndLeasePopUp,
                    title: "End Lease".localized(),
                    message: "Are you sure you want to end the current lease?".localized(),
                    buttonTitle: "Confirm".localized(),
                    secondaryButtonTitle: "Cancel".localized(),
                    action: {
                        
                    },
                    secondaryAction: {
                    }
                )
                .accessibilityIdentifier("EndLeaseAlert")
            }
            if showDeletePropertyPopUp {
                CustomAlertTwoButtons(
                    isActive: $showDeletePropertyPopUp,
                    title: "Delete Property".localized(),
                    message: "Are you sure you want to delete this property?".localized(),
                    buttonTitle: "Confirm".localized(),
                    secondaryButtonTitle: "Cancel".localized(),
                    action: {
                        
                    },
                    secondaryAction: {
                    }
                )
                .accessibilityIdentifier("DeletePropertyAlert")
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
                leftText: String(
                    format: "start_date".localized(),
                    property.leaseStartDate != nil ? formatDateString(property.leaseStartDate!) : "No start date assigned".localized()
                ),
                rightIcon: "coloncurrencysign.arrow.trianglehead.counterclockwise.rotate.90",
                rightText: String(format: "rent_month".localized(), property.monthlyRent)
            )
            .accessibilityIdentifier("lease_start_date")

            buildRow(
                icon: "calendar",
                leftText: String(
                    format: "end_date".localized(),
                    property.leaseEndDate != nil ? formatDateString(property.leaseEndDate!) : "No end date assigned".localized()
                ),
                rightIcon: "eurosign.bank.building",
                rightText: String(format: "deposit_value".localized(), property.deposit)
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
            isAvailable: "Busy",
            tenantName: "John & Mary Doe",
            leaseStartDate: "13/04/2025",
            leaseEndDate: "13/04/2025",
            documents: [],
            rooms: []
        )
        
        let viewModel = PropertyViewModel()
        
        PropertyDetailView(property: .constant(property), viewModel: viewModel)
            .environmentObject(viewModel)
    }
}
