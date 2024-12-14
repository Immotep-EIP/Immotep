//
//  PropertyView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import SwiftUI

struct PropertyView: View {
    @StateObject private var viewModel = PropertyViewModel()
    @State private var isCreatingProperty = false

    var body: some View {
        NavigationView {
            VStack(spacing: 0) {
                TopBar(title: "Property".localized())

                HStack {
                    Spacer()
                    NavigationLink(destination: CreatePropertyView(viewModel: viewModel)) {
                        Text("Add a property".localized())
                            .font(.headline)
                            .foregroundColor(.white)
                            .padding(.horizontal)
                            .padding(.vertical, 8)
                            .background(Color.blue)
                            .cornerRadius(8)
                    }
                    .padding()
                    .accessibilityLabel("add_property")
                }

                ScrollView {
                    if !viewModel.properties.isEmpty {
                        ForEach($viewModel.properties) { $property in
                            NavigationLink(destination: PropertyDetailView(property: $property)) {
                                PropertyCardView(property: property)
                                    .padding(.horizontal)
                                    .padding(.vertical, 4)
                            }
                            .buttonStyle(PlainButtonStyle())
                        }
                    } else {
                        NavigationLink(destination: PropertyDetailView(property: Binding(
                            get: { exampleDataProperty },
                            set: { _ in }
                        ))) {
                            PropertyCardView(property: exampleDataProperty)
                                .padding(.horizontal)
                                .padding(.vertical, 4)
                        }
                        .accessibilityLabel("nav_link_details")
                    }
                }
                TaskBar()
            }
        }
        .navigationBarBackButtonHidden(true)
        .onAppear {
            Task {
                await viewModel.fetchProperties()
            }
        }
    }
}

struct PropertyCardView: View {
    let property: Property

    var body: some View {
        ZStack(alignment: .topLeading) {
            VStack {
                HStack {
                    if let uiImage = property.photo {
                        Image(uiImage: uiImage)
                            .resizable()
                            .scaledToFill()
                            .frame(width: 50, height: 50)
                            .clipShape(Circle())
                            .overlay(Circle().stroke(Color.black, lineWidth: 1))
                    } else {
                        Image("DefaultImageProperty")
                            .resizable()
                            .scaledToFit()
                            .frame(width: 50, height: 50)
                            .clipShape(Circle())
                            .overlay(Circle().stroke(Color("textColor"), lineWidth: 1))
                            .accessibilityLabel("image_property")
                    }

                    VStack(alignment: .leading, spacing: 4) {
                        Text(property.address)
                            .font(.headline)
                            .padding(.trailing, 25)
                            .lineLimit(2)
                            .accessibilityLabel("text_address")

                        if let tenant = property.tenantName {
                            Text(tenant)
                                .font(.subheadline)
                                .foregroundColor(.secondary)
                                .accessibilityLabel("text_tenant")
                        }

                        if let leaseStart = property.leaseStartDate {
                            Text(
                                String(
                                    format: "started_on".localized(),
                                    dateFormatter.string(from: leaseStart)
                                ))
                                .font(.subheadline)
                                .foregroundColor(.secondary)
                                .accessibilityLabel("text_started_on")
                        }
                    }
                }
                .padding(.trailing, 16)
            }

            if property.isAvailable {
                Text("Available".localized())
                    .font(.caption)
                    .foregroundColor(.green)
                    .padding(.horizontal, 8)
                    .padding(.vertical, 4)
                    .background(Capsule().fill(Color.green.opacity(0.2)))
                    .frame(maxWidth: .infinity, alignment: .topTrailing)
                    .accessibilityLabel("text_available")
            } else {
                Text("Busy".localized())
                    .font(.caption)
                    .foregroundColor(.red)
                    .padding(.horizontal, 8)
                    .padding(.vertical, 4)
                    .background(Capsule().fill(Color.red.opacity(0.2)))
                    .frame(maxWidth: .infinity, alignment: .topTrailing)
                    .accessibilityLabel("text_busy")
            }
        }
        .padding(10)
        .cornerRadius(10)
        .shadow(color: .black.opacity(0.1), radius: 4, x: 0, y: 2)
        .overlay(
            RoundedRectangle(cornerRadius: 10)
                .stroke(Color.gray.opacity(0.5), lineWidth: 1)
        )
        .navigationBarBackButtonHidden(true)
        .navigationTransition(
            .fade(.in).animation(.easeInOut(duration: 0))
        )
    }
}

private let dateFormatter: DateFormatter = {
    let formatter = DateFormatter()
    formatter.dateStyle = .medium
    return formatter
}()

struct PropertyView_Previews: PreviewProvider {
    static var previews: some View {
        PropertyView()
            .environmentObject(PropertyViewModel())
    }
}
