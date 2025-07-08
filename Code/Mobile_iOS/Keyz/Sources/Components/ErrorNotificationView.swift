//
//  ErrorNotificationView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 17/06/2025.
//

import SwiftUI

enum NotificationType {
    case error
    case success
}

struct ErrorNotificationView: View {
    let message: String
    let type: NotificationType
    @State private var isVisible = false
    @State private var opacity: Double = 0.0
    private var duration: Double {
        type == .success ? 2.0 : 5.0
    }

    init(message: String, type: NotificationType = .error) {
        self.message = message
        self.type = type
    }

    var body: some View {
        ZStack {
            if isVisible {
                VStack {
                    Spacer()
                    HStack {
                        Image(systemName: type == .success ? "checkmark.circle.fill" : "exclamationmark.triangle.fill")
                            .foregroundColor(.white)
                            .padding(.leading, 10)
                        Text(message)
                            .font(.system(size: 14, weight: .semibold))
                            .foregroundColor(.white)
                            .padding(.vertical)
                            .frame(maxWidth: .infinity, alignment: .leading)
                    }
                    .background(type == .success ? Color("GreenAlert") : Color("RedAlert"))
                    .cornerRadius(8)
                    .shadow(color: Color.black.opacity(0.2), radius: 4, x: 0, y: 2)
                    .padding(.horizontal, 16)
                    .padding(.bottom, 20)
                    .transition(.opacity)
                    .opacity(opacity)
                    .animation(.easeInOut(duration: 0.3), value: opacity)
                }
            }
        }
        .onAppear {
            isVisible = true
            opacity = 1.0
            DispatchQueue.main.asyncAfter(deadline: .now() + duration) {
                withAnimation {
                    opacity = 0.0
                }
                DispatchQueue.main.asyncAfter(deadline: .now() + 0.3) {
                    isVisible = false
                }
            }
        }
    }
}

struct ErrorNotificationView_Previews: PreviewProvider {
    static var previews: some View {
        VStack {
            ErrorNotificationView(message: "Error fetching data")
            ErrorNotificationView(message: "Data saved successfully!", type: .success)
        }
    }
}
