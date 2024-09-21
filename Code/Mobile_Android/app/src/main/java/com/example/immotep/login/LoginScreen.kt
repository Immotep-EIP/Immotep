package com.example.immotep.login

import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.navigation.NavController
import androidx.lifecycle.viewmodel.compose.viewModel

@Composable
fun LoginScreen(navController: NavController, viewModel: LoginViewModel = viewModel()) {
    Button(onClick = {
        navController.navigate("dashboard")
    }) {
        Text("Login")
    }
}
