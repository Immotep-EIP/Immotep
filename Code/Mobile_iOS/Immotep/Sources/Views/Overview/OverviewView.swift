//
//  OverviewView.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 21/10/2024.
//

import SwiftUI

struct OverviewView: View {
    var body: some View {
        NavigationStack {
            VStack {
                TopBar()
                Spacer()

                TaskBar()
            }
            .navigationBarBackButtonHidden(true)
        }
    }
}

struct OverviewView_Previews: PreviewProvider {
    static var previews: some View {
        OverviewView()
    }
}
