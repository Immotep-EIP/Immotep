package com.example.immotep.dashboard

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.components.LoggedBottomBar
import com.example.immotep.components.LoggedTopBar

@Composable
fun DashBoardLayout(
    navController: NavController,
    testTag: String,
    content: @Composable () -> Unit
) {
    Column(modifier = Modifier.testTag(testTag)) {
        LoggedTopBar(navController)
        Column(modifier = Modifier.weight(1f).padding(2.dp)) {
            content()
        }
        LoggedBottomBar(navController)
    }
}
