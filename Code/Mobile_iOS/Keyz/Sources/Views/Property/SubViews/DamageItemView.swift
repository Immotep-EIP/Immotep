//
//  DamageItemView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 08/06/2025.
//

import SwiftUI

struct DamageItemView: View {
    let damage: DamageResponse

    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            HStack {
                Text(damage.roomName)
                    .font(.headline)
                    .foregroundColor(Color("textColor"))
                Spacer()
                Text(damage.priority.capitalized)
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
                Text("Status: \(damage.fixStatus.replacingOccurrences(of: "_", with: " ").capitalized)")
                    .font(.caption)
                    .foregroundColor(damage.fixStatus == "fixed" ? Color("GreenAlert") : Color("RedAlert"))
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
    
    private func formatDateString(_ dateString: String) -> String {
        let formatter = ISO8601DateFormatter()
        if let date = formatter.date(from: dateString) {
            let displayFormatter = DateFormatter()
            displayFormatter.dateStyle = .medium
            displayFormatter.timeStyle = .none
            return displayFormatter.string(from: date)
        }
        return dateString
    }
}

struct DamageItemView_Previews: PreviewProvider {
    static var previews: some View {
        DamageItemView(damage: DamageResponse(
            id: "damage_001",
            comment: "Cracked window in the living room",
            priority: "high",
            roomName: "Living Room",
            fixStatus: "pending",
            pictures: ["base64_image_1"],
            createdAt: "2025-05-10T09:00:00Z",
            updatedAt: nil,
            fixPlannedAt: "2025-05-25T14:00:00Z",
            fixedAt: nil,
            leaseId: "lease_001",
            propertyId: "",
            propertyName: "Condo",
            tenantName: "John & Mary Doe",
            read: true
        ))
    }
}
