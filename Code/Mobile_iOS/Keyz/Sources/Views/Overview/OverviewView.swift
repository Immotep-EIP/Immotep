//
//  OverviewView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 21/10/2024.
//

import SwiftUI

struct OverviewView: View {
    @EnvironmentObject private var loginViewModel: LoginViewModel
    @AppStorage("isLoggedIn") var isLoggedIn: Bool = false

    var body: some View {
        NavigationStack {
            VStack {
                TopBar(title: "Keyz")
                ScrollView {
                    VStack(spacing: 30) {
                        OverviewBox(
                            title: "",
                            content: [
                                String(
                                    format: "welcome_message".localized(),
                                    loginViewModel.user?.firstname ?? ""),
                                    "Here is an overview of your appartments".localized()
                            ])
                        OverviewBox(
                            title: "New messages".localized(),
                            content: ["Email 1: Subject", "Email 2: Subject", "Email 3: Subject"])
                        OverviewBox(
                            title: "Scheduled inventory".localized(),
                            content: ["Inventory 1: Scheduled", "Inventory 2: Scheduled"])
                        OverviewBox(
                            title: "Damage in progress".localized(),
                            content: ["Damage 1: In Progress", "Damage 2: In Progress"])
                        OverviewBox(
                            title: "Available properties".localized(),
                            content: ["Property 1: Available", "Property 2: Available"])
                    }
                    .padding()
                }
            }
            .navigationBarBackButtonHidden(true)
            .onAppear {
                loginViewModel.loadUser()
            }
        }
    }
}

struct OverviewBox: View {
    var title: String
    var content: [String]

    var body: some View {
        VStack(alignment: .leading) {
            HStack {
                Text(title)
                    .font(.headline)
                Spacer()
                if !title.isEmpty {
                    Menu {
                        Button("Option 1") { }
                        Button("Option 2") { }
                        Button("Option 3") { }
                    } label: {
                        Image(systemName: "ellipsis.circle")
                            .font(.title2)
                            .foregroundColor(.gray)
                    }
                }
            }
            .padding(.bottom, 5)

            VStack(alignment: .leading, spacing: 5) {
                ForEach(content, id: \.self) { item in
                    Text(item)
                        .font(.system(size: 14))
                        .frame(maxWidth: .infinity, alignment: .leading)
                }
            }
            .padding()
            .cornerRadius(10)
            .overlay(
                RoundedRectangle(cornerRadius: 10)
                    .stroke(Color.gray, lineWidth: 0.5)
            )
        }
        .padding(.horizontal)
    }
}

struct OverviewView_Previews: PreviewProvider {
    static var previews: some View {
        OverviewView()
            .environmentObject(LoginViewModel())
    }
}
