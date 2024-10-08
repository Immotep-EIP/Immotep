package com.example.immotep.dashboard

import androidx.compose.foundation.layout.Column
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController

@Composable
fun DashBoardScreen(
    navController: NavController,
    viewModel: DashBoardViewModel = viewModel(),
) {
    Column(modifier = Modifier.testTag("dashboardScreen")) {
        Button(onClick = {
            navController.navigate("login")
        }) {
            Text("Dashboard")
        }
    }
}
