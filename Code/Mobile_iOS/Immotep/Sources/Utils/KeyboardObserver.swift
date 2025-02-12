//
//  KeyboardObserver.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 17/11/2024.
//

import SwiftUI
import Combine

class KeyboardObserver: ObservableObject {
    @State var isKeyboardVisible = false
    private var keyboardPublisher: AnyCancellable?

    init() {
        self.keyboardPublisher = NotificationCenter.default.publisher(for: UIResponder.keyboardWillShowNotification)
            .merge(with: NotificationCenter.default.publisher(for: UIResponder.keyboardWillHideNotification))
            .sink { [weak self] notification in
                guard let self = self else { return }
                if notification.name == UIResponder.keyboardWillShowNotification {
                    self.isKeyboardVisible = true
                } else {
                    self.isKeyboardVisible = false
                }
            }
    }

    deinit {
        keyboardPublisher?.cancel()
    }
}
