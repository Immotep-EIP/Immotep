package fr.keyz.dashboard

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Modifier
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import fr.keyz.LocalApiService
import fr.keyz.components.InitialFadeIn
import fr.keyz.components.InternalLoading
import fr.keyz.dashboard.widgets.DamageWidget
import fr.keyz.dashboard.widgets.DamagesListWidget
import fr.keyz.dashboard.widgets.HelloWidget
import fr.keyz.dashboard.widgets.PropertiesWidget
import fr.keyz.dashboard.widgets.RemindersWidget

@Composable
fun DashBoardScreenOwner(
    navController: NavController,
) {
    val apiService = LocalApiService.current
    val viewModel : DashBoardViewModel = viewModel {
        DashBoardViewModel(navController, apiService)
    }
    val isLoading = viewModel.isLoading.collectAsState()
    val dashBoard = viewModel.dashBoard.collectAsState()
    val userName = viewModel.userName.collectAsState()

    LaunchedEffect(Unit) {
        viewModel.getDashBoard()
        viewModel.getName()
    }

    DashBoardLayout(navController, "dashboardScreen") {
        if (isLoading.value) {
            InternalLoading()
            return@DashBoardLayout
        }
        InitialFadeIn(durationMs = 200) {
            Column(modifier = Modifier.verticalScroll(rememberScrollState())) {
                HelloWidget(dashBoard.value.properties.nbrTotal, userName = userName.value)
                RemindersWidget(dashBoard.value.reminders)
                PropertiesWidget(dashBoard.value.properties)
                //UnreadMessagesWidget()
                //ScheduledInventoryWidget()
                DamageWidget(dashBoard.value.openDamages)
                DamagesListWidget(dashBoard.value.openDamages.listToFix)
            }
        }
    }
}