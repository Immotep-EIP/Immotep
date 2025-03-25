//
//  SessionStorage.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 27/10/2024.
//

import Foundation

actor SessionStorage {
    private static var accessToken: String?
    private static var refreshToken: String?
    private static let queue = DispatchQueue(label: "com.yourapp.sessionStorageQueue")

    static func setAccessToken(_ token: String?) {
        queue.sync {
            accessToken = token
        }
    }

    static func getAccessToken() -> String? {
        return queue.sync {
            return accessToken
        }
    }

    static func setRefreshToken(_ token: String?) {
        queue.sync {
            refreshToken = token
        }
    }

    static func getRefreshToken() -> String? {
        return queue.sync {
            return refreshToken
        }
    }
}
