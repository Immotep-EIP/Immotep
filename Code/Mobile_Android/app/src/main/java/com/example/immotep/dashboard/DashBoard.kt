package com.example.immotep.dashboard

import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.LocalIsOwner

@Composable
fun DashBoardScreen(
    navController: NavController,
    viewModel: DashBoardViewModel = viewModel()
) {
    DashBoardLayout(navController, "dashboardScreen") {
        Button(onClick = {
            navController.navigate("login")
        }) {
            Text("Dashboard")
        }
    }
}
