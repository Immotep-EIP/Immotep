//
//  OverviewViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 21/10/2024.
//

import Foundation

class OverviewViewModel: ObservableObject {
    @Published var dashboardData: DashboardResponse?
    @Published var isLoading: Bool = false
    @Published var errorMessage: String?

    init() {}

    @MainActor
    func fetchDashboardData() async {
        isLoading = true
        errorMessage = nil

        guard let url = URL(string: "\(APIConfig.baseURL)/owner/dashboard/") else {
            errorMessage = "Invalid URL".localized()
            isLoading = false
            return
        }

        do {
            let token = try await TokenStorage.getValidAccessToken()
            var request = URLRequest(url: url)
            request.httpMethod = "GET"
            request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
            request.setValue("application/json", forHTTPHeaderField: "Accept")

            let (data, response) = try await URLSession.shared.data(for: request)
            guard let httpResponse = response as? HTTPURLResponse else {
                throw NSError(domain: "", code: 0, userInfo: [NSLocalizedDescriptionKey: "Invalid server response".localized()])
            }

            let responseBody = String(data: data, encoding: .utf8) ?? "Unable to decode response"

            guard (200...299).contains(httpResponse.statusCode) else {
                throw NSError(domain: "", code: httpResponse.statusCode, userInfo: [NSLocalizedDescriptionKey: "Failed with status code: \(httpResponse.statusCode) - \(responseBody)".localized()])
            }

            let decoder = JSONDecoder()
            let dashboardResponse = try decoder.decode(DashboardResponse.self, from: data)

            var reminderDict: [String: Reminder] = [:]
            for reminder in dashboardResponse.reminders {
                if reminderDict[reminder.id] == nil {
                    reminderDict[reminder.id] = reminder
                }
            }
            let uniqueReminders = reminderDict.values.sorted { $0.id < $1.id }

            let uniqueDashboardResponse = DashboardResponse(
                reminders: Array(uniqueReminders),
                properties: dashboardResponse.properties,
                openDamages: dashboardResponse.openDamages
            )
            dashboardData = uniqueDashboardResponse
            
        } catch {
            errorMessage = "Error fetching dashboard data: \(error.localizedDescription)".localized()
        }

        isLoading = false
    }
}
