//
//  APIConfig.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 24/03/2025.
//

import Foundation

enum AppEnvironment {
    case local
    case dev
    case prod

    var baseURL: URL {
        switch self {
        case .local:
            return URL(string: "http://localhost:3001/v1")!
        case .dev:
            return URL(string: "https://dev.space.keyz-app.fr/api/v1")!
        case .prod:
            return URL(string: "https://space.keyz-app.fr/api/v1")!
        }
    }
}

struct APIConfig {
    static let currentEnvironment: AppEnvironment = .dev // dev or prod for online
    static let baseURL = currentEnvironment.baseURL
}
