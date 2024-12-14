//
//  TaskBar.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 26/10/2024.
//

import SwiftUI
import NavigationTransitions

struct TaskBar: View {
    @AppStorage("lang") private var lang: String = "en"

    var body: some View {
            VStack {
                Divider()
                    .background(Color("textColor"))
                HStack {
                    Spacer()

                    NavigationLink(destination: OverviewView()) {
                        VStack {
                            Image(systemName: "house")
                                .font(.system(size: 25))
                                .frame(width: 30, height: 30)
                            Text("Overview".localized())
                                .frame(width: 90)
                                .font(.footnote)
                        }
                    }
                    Spacer()

                    NavigationLink(destination: PropertyView()) {
                        VStack {
                            Image(systemName: "building.2")
                                .font(.system(size: 25))
                                .frame(width: 30, height: 30)
                            Text("Real Property".localized())
                                .frame(width: 90)
                                .font(.footnote)
                        }
                    }
                    Spacer()

                    NavigationLink(destination: MessagesView()) {
                        VStack {
                            Image(systemName: "ellipsis.message")
                                .font(.system(size: 25))
                                .frame(width: 30, height: 30)
                            Text("Messages".localized())
                                .frame(width: 90)
                                .font(.footnote)
                        }
                    }
                    Spacer()

                    NavigationLink(destination: SettingsView()) {
                        VStack {
                            Image(systemName: "gear")
                                .font(.system(size: 25))
                                .frame(width: 30, height: 30)
                            Text("Settings".localized())
                                .frame(width: 90)
                                .font(.footnote)
                        }
                    }
                    Spacer()
                }
                .foregroundStyle(Color("textColor"))
                .frame(maxWidth: .infinity, alignment: .center)
            }
        }
}
