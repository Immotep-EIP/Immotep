package com.example.keyz.dashboard

import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.keyz.LocalApiService
import com.example.keyz.components.LoadingDialog

@Composable
fun DashBoardScreen(
    navController: NavController,
) {
    val apiService = LocalApiService.current

    val viewModel : DashBoardViewModel = viewModel {
        DashBoardViewModel(navController, apiService)
    }

    val isLoading = viewModel.isLoading.collectAsState()

    LaunchedEffect(Unit) {
        viewModel.getDashBoards()
    }

    DashBoardLayout(navController, "dashboardScreen") {
        LoadingDialog(isLoading.value)
        Button(onClick = {
            navController.navigate("login")
        }) {
            Text("Dashboard")
        }
    }
}
