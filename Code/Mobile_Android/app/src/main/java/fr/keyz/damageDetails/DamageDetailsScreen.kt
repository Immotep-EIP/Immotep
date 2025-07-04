package fr.keyz.damageDetails

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import fr.keyz.LocalApiService
import fr.keyz.LocalIsOwner

@Composable
fun DetailsScreen(navController: NavController, propertyId: String?, leaseId: String, damageId: String) {
    val isOwner = LocalIsOwner.current
    val apiService = LocalApiService.current
    val viewModel = viewModel {
        DamageDetailsViewModel(apiService, navController)
    }
    val isLoading = viewModel.isLoading.collectAsState()
    val apiError = viewModel.apiError.collectAsState()
    val currentDamage = viewModel.currentDamage.collectAsState()

    LaunchedEffect(damageId) {
        viewModel.getDamage(propertyId, leaseId, damageId)
    }
}