//
//  SettingsView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 21/10/2024.
//

import SwiftUI

struct SettingsView: View {
    @AppStorage("lang") private var lang: String = "en"

    var body: some View {
        VStack(alignment: .leading) {
            TopBar()

            Text("Settings".localized())
                .font(.title)
                .fontWeight(.bold)
                .padding(.leading, 10)
                .padding(.bottom, 30)

            Text("Language".localized())
                .font(.headline)
                .fontWeight(.bold)
                .padding(.leading, 20)
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
            .padding()

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
