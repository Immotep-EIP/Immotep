package com.example.immotep.login

import androidx.compose.material3.Button
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController

@Composable
fun LoginScreen(
    navController: NavController,
    viewModel: LoginViewModel = viewModel(),
) {
    val post = viewModel.post.collectAsState()

    if (post.value != null) {
        Text("Post title: ${post.value}")
        // Display other post data as needed
    } else {
        CircularProgressIndicator() // Show loading indicator while fetching data
    }

    Button(onClick = {
        navController.navigate("dashboard")
    }) {
        Text("Login")
    }
}
