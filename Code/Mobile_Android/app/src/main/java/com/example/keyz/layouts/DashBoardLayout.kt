package com.example.keyz.dashboard

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.unit.dp
import androidx.navigation.NavController
import com.example.keyz.components.LoggedBottomBar
import com.example.keyz.components.LoggedTopBar

@Composable
fun DashBoardLayout(
    navController: NavController,
    testTag: String,
    content: @Composable () -> Unit
) {
    Column(modifier = Modifier.testTag(testTag)) {
        LoggedTopBar(navController)
        Column(modifier = Modifier.weight(1f).padding(2.dp).testTag("dashboardLayout")) {
            content()
        }
        LoggedBottomBar(navController)
    }
}
