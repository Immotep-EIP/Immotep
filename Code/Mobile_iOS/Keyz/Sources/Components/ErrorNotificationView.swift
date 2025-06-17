//
//  ErrorNotificationView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 17/06/2025.
//


//
//  ErrorNotificationView.swift
//  Keyz
//
//  Created by Liebenguth Alessio on 17/06/2025.
//


import SwiftUI

struct ErrorNotificationView: View {
    let message: String
    @State private var isVisible = false
    @State private var opacity: Double = 0.0
    private let duration: TimeInterval = 5.0

    var body: some View {
        ZStack {
            if isVisible {
                VStack {
                    Spacer()
                    HStack {
                        Image(systemName: "exclamationmark.triangle.fill")
                            .foregroundColor(.white)
                            .padding(.leading, 10)
                        Text(message)
                            .font(.system(size: 14, weight: .semibold))
                            .foregroundColor(.white)
                            .padding(.vertical)
                            .frame(maxWidth: .infinity, alignment: .leading)
                    }
                    .background(Color("RedAlert"))
                    .cornerRadius(12)
                    .shadow(color: Color.black.opacity(0.2), radius: 4, x: 0, y: 2)
                    .transition(.opacity)
                    .opacity(opacity)
                    .padding(.horizontal, 16)
                    .padding(.bottom, 20)
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
        ErrorNotificationView(message: "Error fetching data")
    }
}
