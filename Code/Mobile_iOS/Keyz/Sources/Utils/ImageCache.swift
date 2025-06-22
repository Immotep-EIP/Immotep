//
//  ImageCache.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 22/06/2025.
//

import SwiftUI

class ImageCache {
    static let shared = ImageCache()
    private var cache: [String: UIImage] = [:]
    private let lock = NSLock()

    private init() {}

    func getImage(forKey key: String) -> UIImage? {
        lock.lock()
        defer { lock.unlock() }
        return cache[key]
    }

    func setImage(_ image: UIImage?, forKey key: String) {
        lock.lock()
        defer { lock.unlock() }
        cache[key] = image
    }

    func clearCache() {
        lock.lock()
        defer { lock.unlock() }
        cache.removeAll()
    }
}
