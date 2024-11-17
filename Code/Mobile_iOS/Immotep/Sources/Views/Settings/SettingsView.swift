//
//  SettingsView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 21/10/2024.
//

import SwiftUI

struct SettingsView: View {
    @AppStorage("lang") private var lang: String = "en"
    @AppStorage("theme") private var selectedTheme: String = ThemeOption.system.rawValue

    var body: some View {
        VStack(alignment: .leading) {
            TopBar(title: "Settings".localized())

            Form {
                Section(header: Text("Language")) {
                    Picker(selection: $lang, label: Text("Language")) {
                        Text("English").tag("en")
                        Text("Français").tag("fr")
                    }
                    .pickerStyle(SegmentedPickerStyle())
                    .onChange(of: lang) {
                        Task {
                            await Bundle.setLanguage(lang)
                        }
                    }
                }

                Section(header: Text("Theme")) {
                    Picker("Theme", selection: $selectedTheme) {
                        ForEach(ThemeOption.allCases, id: \.self) { theme in
                            Text(theme.rawValue)
                                .tag(theme.rawValue)
                        }
                    }
                    .pickerStyle(SegmentedPickerStyle())
                    .onChange(of: selectedTheme) {
                        applyTheme(theme: selectedTheme)
                    }
                }
            }
            .onAppear {
                applyTheme(theme: selectedTheme)
            }

            Spacer()
            TaskBar()
        }
        .navigationBarBackButtonHidden(true)
    }

    private func applyTheme(theme: String) {
        guard let windowScene = UIApplication.shared.connectedScenes.first as? UIWindowScene else { return }
        let rootViewController = windowScene.windows.first?.rootViewController

        switch theme {
        case ThemeOption.light.rawValue:
            rootViewController?.overrideUserInterfaceStyle = .light
        case ThemeOption.dark.rawValue:
            rootViewController?.overrideUserInterfaceStyle = .dark
        default:
            rootViewController?.overrideUserInterfaceStyle = .unspecified
        }
    }
}

struct SettingsView_Previews: PreviewProvider {
    static var previews: some View {
        SettingsView()
    }
}
