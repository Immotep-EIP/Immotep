//
//  APIConfig.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 24/03/2025.
//

import Foundation

enum AppEnvironment {
    case local
    case online

    var baseURL: URL {
        switch self {
        case .local:
            return URL(string: "http://localhost:3001/api/v1")!
        case .online:
            return URL(string: "https://dev.backend.immotep.mazettt.fr/api/v1")!
        }
    }
}

struct APIConfig {
    static let currentEnvironment: AppEnvironment = .online // .online for prod
    static let baseURL = currentEnvironment.baseURL
}
