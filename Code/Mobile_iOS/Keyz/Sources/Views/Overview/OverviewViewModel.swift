//
//  OverviewViewModel.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 21/10/2024.
//

import SwiftUI
import Foundation

class OverviewViewModel: ObservableObject {
    @Published var dashboardData: DashboardResponse?
    @Published var isLoading: Bool = false
    @Published var errorMessage: String?
    @AppStorage("lang") private var lang: String = "en"

    init() {}

    @MainActor
    func fetchDashboardData() async {
        isLoading = true
        errorMessage = nil

        guard var urlComponents = URLComponents(string: "\(APIConfig.baseURL)/owner/dashboard/") else {
            errorMessage = "Invalid URL".localized()
            isLoading = false
            return
        }
        urlComponents.queryItems = [URLQueryItem(name: "lang", value: lang)]
        guard let url = urlComponents.url else {
            errorMessage = "Invalid URL with lang parameter".localized()
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
            decoder.dateDecodingStrategy = .iso8601

            do {
                let dashboardResponse = try decoder.decode(DashboardResponse.self, from: data)
                let remindersWithUniqueIds = dashboardResponse.reminders.enumerated().map { index, reminder in
                    var uniqueReminder = reminder
                    uniqueReminder.id = UUID().uuidString
                    return uniqueReminder
                }

                let uniqueDashboardResponse = DashboardResponse(
                    reminders: remindersWithUniqueIds,
                    properties: dashboardResponse.properties,
                    openDamages: dashboardResponse.openDamages
                )
                dashboardData = uniqueDashboardResponse
            } catch {
                throw error
            }
            
        } catch {
            errorMessage = "Error fetching dashboard data: \(error.localizedDescription)".localized()
        }

        isLoading = false
    }
}
