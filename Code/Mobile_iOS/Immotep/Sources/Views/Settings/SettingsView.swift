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
        VStack(spacing: 0) {
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
                        Task { @MainActor in
                            ThemeManager.applyTheme(theme: selectedTheme)
                        }
                    }
                }
            }
            Spacer()
            TaskBar()
        }
        .navigationBarBackButtonHidden(true)
    }
}

struct SettingsView_Previews: PreviewProvider {
    static var previews: some View {
        SettingsView()
    }
}
