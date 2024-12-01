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
                        Text("Add a property")
                            .font(.headline)
                            .foregroundColor(.white)
                            .padding(.horizontal)
                            .padding(.vertical, 8)
                            .background(Color.blue)
                            .cornerRadius(8)
                    }
                    .padding()
                }

                ScrollView {
                    ForEach($viewModel.properties) { $property in
                        NavigationLink(destination: PropertyDetailView(property: $property)) {
                            PropertyCardView(property: property)
                                .padding(.horizontal)
                                .padding(.vertical, 4)
                        }
                        .buttonStyle(PlainButtonStyle())
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
                            .overlay(Circle().stroke(Color.black, lineWidth: 1))
                    }

                    VStack(alignment: .leading, spacing: 4) {
                        Text(property.address)
                            .font(.headline)
                            .padding(.trailing, 25)
                            .lineLimit(2)

                        if let tenant = property.tenantName {
                            Text(tenant)
                                .font(.subheadline)
                                .foregroundColor(.secondary)
                        }

                        if let leaseStart = property.leaseStartDate {
                            Text("Started on \(leaseStart, formatter: dateFormatter)")
                                .font(.subheadline)
                                .foregroundColor(.secondary)
                        }
                    }
                }
                .padding(.trailing, 16)
            }

            if property.isAvailable {
                Text("Available")
                    .font(.caption)
                    .foregroundColor(.green)
                    .padding(.horizontal, 8)
                    .padding(.vertical, 4)
                    .background(Capsule().fill(Color.green.opacity(0.2)))
                    .frame(maxWidth: .infinity, alignment: .topTrailing)
            } else {
                Text("Busy")
                    .font(.caption)
                    .foregroundColor(.red)
                    .padding(.horizontal, 8)
                    .padding(.vertical, 4)
                    .background(Capsule().fill(Color.red.opacity(0.2)))
                    .frame(maxWidth: .infinity, alignment: .topTrailing)
            }
        }
        .padding(10)
        .background(Color.white)
        .cornerRadius(10)
        .shadow(color: .black.opacity(0.1), radius: 4, x: 0, y: 2)
        .overlay(
            RoundedRectangle(cornerRadius: 10)
                .stroke(Color.gray.opacity(0.5), lineWidth: 1)
        )
        .navigationBarBackButtonHidden(true)
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
