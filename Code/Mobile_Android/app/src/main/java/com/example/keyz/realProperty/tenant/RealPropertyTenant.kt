package com.example.keyz.realProperty.tenant

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.keyz.LocalApiService
import com.example.keyz.components.InternalLoading
import com.example.keyz.dashboard.DashBoardLayout
import com.example.keyz.inventory.loaderButton.LoaderInventoryViewModel
import com.example.keyz.realProperty.details.RealPropertyDetailsScreen

@Composable
fun RealPropertyTenant(navController: NavController, loaderInventoryViewModel: LoaderInventoryViewModel) {
    val apiService = LocalApiService.current
    val viewModel: RealPropertyTenantViewModel = viewModel {
        RealPropertyTenantViewModel(apiService, navController)
    }
    val property = viewModel.property.collectAsState()
    val isLoading = viewModel.isLoading.collectAsState()
    LaunchedEffect(Unit) {
        viewModel.loadProperty()
    }
    DashBoardLayout(navController, "realPropertyTenant") {
        if (property.value != null && !isLoading.value) {
            RealPropertyDetailsScreen(
                navController = navController,
                getBack = {},
                loaderInventoryViewModel = loaderInventoryViewModel,
                newProperty = property.value!!
            )
        } else {
            InternalLoading()
        }
    }
}