package fr.keyz.realProperty.tenant

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import fr.keyz.LocalApiService
import fr.keyz.components.ErrorAlert
import fr.keyz.components.InternalLoading
import fr.keyz.dashboard.DashBoardLayout
import fr.keyz.inventory.loaderButton.LoaderInventoryViewModel
import fr.keyz.realProperty.details.RealPropertyDetailsScreen

@Composable
fun RealPropertyTenant(navController: NavController, loaderInventoryViewModel: LoaderInventoryViewModel) {
    val apiService = LocalApiService.current
    val viewModel: RealPropertyTenantViewModel = viewModel {
        RealPropertyTenantViewModel(apiService, navController)
    }
    val property = viewModel.property.collectAsState()
    val isLoading = viewModel.isLoading.collectAsState()
    val errorLoading = viewModel.loadingError.collectAsState()
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
        }
        else if (!isLoading.value && errorLoading.value != null) {
            ErrorAlert(errorLoading.value, null)
        } else {
            InternalLoading()
        }
    }
}