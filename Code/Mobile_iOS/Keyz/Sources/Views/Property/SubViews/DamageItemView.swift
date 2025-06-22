//
//  DamageItemView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/06/2025.
//

import SwiftUI
import Foundation

struct DamageItemView: View {
    let damage: DamageResponse
    @Binding var selectedDamageId: String?

    var body: some View {
        Button(action: {
            selectedDamageId = damage.id
        }) {
            VStack(alignment: .leading, spacing: 8) {
                HStack {
                    Text(damage.roomName)
                        .font(.headline)
                        .foregroundColor(Color("textColor"))
                    Spacer()
                    Text(damage.priority.capitalized.localized())
                        .font(.caption)
                        .fontWeight(.medium)
                        .foregroundColor(.white)
                        .padding(.horizontal, 8)
                        .padding(.vertical, 4)
                        .background(
                            RoundedRectangle(cornerRadius: 8)
                                .fill(priorityColor(damage.priority))
                        )
                }
                Text(damage.comment)
                    .font(.subheadline)
                    .foregroundColor(.gray)
                HStack {
                    Text(String(format: "Status: %@".localized(), damage.fixStatus.replacingOccurrences(of: "_", with: " ").capitalized.localized()))
                        .font(.caption)
                        .foregroundColor(statusColor(damage.fixStatus))
                    Spacer()
                    Text(formatDateString(damage.createdAt))
                        .font(.caption)
                        .foregroundColor(.gray)
                }
            }
            .padding()
            .background(Color("basicWhiteBlack"))
            .cornerRadius(10)
            .shadow(color: Color.black.opacity(0.1), radius: 5, x: 0, y: 2)
        }
        .buttonStyle(PlainButtonStyle())
        .accessibilityLabel("damage_item_\(damage.id)")
    }
    
    private func priorityColor(_ priority: String) -> Color {
        switch priority.lowercased() {
        case "low":
            return Color.green
        case "medium":
            return Color.yellow
        case "high":
            return Color.orange
        case "urgent":
            return Color.red
        default:
            return Color.gray
        }
    }
    
    private func statusColor(_ status: String) -> Color {
        switch status.lowercased() {
        case "fixed":
            return Color.green
        case "planned", "awaiting_tenant_confirmation":
            return Color.orange
        case "pending":
            return Color.red
        default:
            return Color.gray
        }
    }
    
    private func formatDateString(_ dateString: String) -> String {
        let isoFormatter = ISO8601DateFormatter()
        isoFormatter.formatOptions = [.withInternetDateTime, .withFractionalSeconds]
        
        let displayFormatter = DateFormatter()
        displayFormatter.dateFormat = "dd/MM/yyyy"
        displayFormatter.locale = Locale(identifier: "fr_FR")

        if let date = isoFormatter.date(from: dateString) {
            return displayFormatter.string(from: date)
        } else {
            return "Invalid Date".localized()
        }
    }
}

struct DamageItemView_Previews: PreviewProvider {
    @State static var selectedDamageId: String? = nil
    
    static var previews: some View {
        NavigationView {
            DamageItemView(
                damage: DamageResponse(
                    id: "damage_001",
                    comment: "Cracked window in the living room",
                    priority: "high",
                    roomName: "Living Room",
                    fixStatus: "pending",
                    pictures: ["base64_image_1"],
                    createdAt: "2025-05-10T21:03:53.293Z",
                    updatedAt: nil,
                    fixPlannedAt: "2025-05-25T14:00:00Z",
                    fixedAt: nil,
                    leaseId: "lease_001",
                    propertyId: "",
                    propertyName: "Condo",
                    tenantName: "John & Mary Doe",
                    read: true
                ),
                selectedDamageId: $selectedDamageId
            )
        }
    }
}
