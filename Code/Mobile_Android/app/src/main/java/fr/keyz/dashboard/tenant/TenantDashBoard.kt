package fr.keyz.dashboard.tenant

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
import fr.keyz.components.ErrorAlert
import fr.keyz.components.InitialFadeIn
import fr.keyz.components.InternalLoading
import fr.keyz.dashboard.DashBoardLayout
import fr.keyz.dashboard.widgets.DamagesListWidget
import fr.keyz.dashboard.widgets.HelloTenant
import fr.keyz.dashboard.widgets.PropertyOverview

@Composable
fun TenantDashBoard(navController: NavController) {
    val apiService = LocalApiService.current
    val viewModel = viewModel {
        TenantDashBoardViewModel(navController, apiService)
    }
    val isLoading = viewModel.isLoading.collectAsState()
    val property = viewModel.property.collectAsState()
    val userName = viewModel.userName.collectAsState()
    val apiError = viewModel.apiError.collectAsState()

    LaunchedEffect(Unit) {
        viewModel.loadDashBoard()
    }

    DashBoardLayout(navController, "dashboardTenantScreen") {
        if (isLoading.value) {
            InternalLoading()
            return@DashBoardLayout
        }
        ErrorAlert(apiError.value)
        InitialFadeIn(durationMs = 200) {
            Column(modifier = Modifier.verticalScroll(rememberScrollState())) {
                HelloTenant(userName = userName.value)
                if (property.value != null) {
                    PropertyOverview(property.value!!)
                }
                DamagesListWidget(viewModel.damages.toTypedArray())
            }
        }
    }
}