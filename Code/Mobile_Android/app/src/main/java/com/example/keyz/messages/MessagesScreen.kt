package com.example.keyz.messages

import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.navigation.NavController
import com.example.keyz.dashboard.DashBoardLayout

@Composable
fun Messages(navController: NavController) {
    DashBoardLayout(
        navController = navController,
        testTag = "messagesScreen"
    ) {
        Text("Coming Soon...")
    }
}