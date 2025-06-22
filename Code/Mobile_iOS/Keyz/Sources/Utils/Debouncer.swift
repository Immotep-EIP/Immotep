//
//  Debouncer.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 22/06/2025.
//

import SwiftUI

class Debouncer {
    private let delay: TimeInterval
    private var workItem: DispatchWorkItem?

    init(delay: TimeInterval) {
        self.delay = delay
    }

    func debounce(_ block: @escaping () -> Void) {
        workItem?.cancel()
        let newWorkItem = DispatchWorkItem(block: block)
        workItem = newWorkItem
        DispatchQueue.main.asyncAfter(deadline: .now() + delay, execute: newWorkItem)
    }
}
