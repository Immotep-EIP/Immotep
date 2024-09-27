package com.example.immotep.dashboard

import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController

@Composable
fun DashBoardScreen(
    navController: NavController,
    viewModel: DashBoardViewModel = viewModel(),
) {
    Button(onClick = {
        navController.navigate("login")
    }) {
        Text("Dashboard")
    }
}
