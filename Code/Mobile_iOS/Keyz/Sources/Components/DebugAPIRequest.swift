//
//  DebugAPIRequest.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 20/06/2025.
//

import Foundation

func debugPrintAPIRequest(_ request: URLRequest) {
    print("===== Debug API Request =====")
    print("URL: \(request.url?.absoluteString ?? "No URL")")
    print("HTTP Method: \(request.httpMethod ?? "No Method")")
    print("Headers:")
    if let headers = request.allHTTPHeaderFields, !headers.isEmpty {
        for (key, value) in headers {
            print("  \(key): \(value)")
        }
    } else {
        print("  No headers")
    }
    print("Body:")
    if let body = request.httpBody, let bodyString = String(data: body, encoding: .utf8) {
        do {
            let jsonObject = try JSONSerialization.jsonObject(with: body, options: [])
            let prettyData = try JSONSerialization.data(withJSONObject: jsonObject, options: [.prettyPrinted])
            if let prettyString = String(data: prettyData, encoding: .utf8) {
                print(prettyString)
            } else {
                print(bodyString)
            }
        } catch {
            print(bodyString)
        }
    } else {
        print("  No body")
    }
    print("============================")
}

func debugPrintAPIResponse(_ data: Data?, response: URLResponse?, error: Error?) {
    print("===== Debug API Response =====")
    if let httpResponse = response as? HTTPURLResponse {
        print("Status Code: \(httpResponse.statusCode)")
    }
    if let data = data, let responseString = String(data: data, encoding: .utf8) {
        print("Response Body: \(responseString)")
    } else {
        print("No response body")
    }
    if let error = error {
        print("Error: \(error.localizedDescription)")
    }
    print("================ =============")
}
