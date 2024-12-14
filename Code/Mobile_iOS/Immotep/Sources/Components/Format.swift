//
//  Format.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 14/12/2024.
//

import Foundation

func stringToDate(_ dateString: String, format: String? = nil) -> Date? {
    let dateFormatter = DateFormatter()

    if let format = format {
        dateFormatter.dateFormat = format
    } else {
        dateFormatter.dateFormat = "dd/MM/yyyy"
    }

    return dateFormatter.date(from: dateString)
}
