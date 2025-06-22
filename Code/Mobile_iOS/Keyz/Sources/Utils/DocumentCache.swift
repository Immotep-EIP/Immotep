//
//  DocumentCache.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 22/06/2025.
//


import Foundation
import PDFKit

class DocumentCache {
    static let shared = DocumentCache()
    private var cache: [String: PDFDocument] = [:]
    private let cacheQueue = DispatchQueue(label: "com.immotep.documentCache", attributes: .concurrent)

    private init() {}

    func getDocument(forKey key: String) -> PDFDocument? {
        cacheQueue.sync {
            return cache[key]
        }
    }

    func setDocument(_ document: PDFDocument?, forKey key: String) {
        cacheQueue.async(flags: .barrier) {
            self.cache[key] = document
        }
    }

    func clearCache() {
        cacheQueue.async(flags: .barrier) {
            self.cache.removeAll()
        }
    }

    func removeDocument(forKey key: String) {
        cacheQueue.async(flags: .barrier) {
            self.cache.removeValue(forKey: key)
        }
    }
}