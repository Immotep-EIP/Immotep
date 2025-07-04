package fr.keyz.damageDetails

import androidx.activity.compose.BackHandler
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import fr.keyz.LocalApiService
import fr.keyz.LocalIsOwner
import fr.keyz.layouts.InventoryLayout
import fr.keyz.R
import fr.keyz.components.InternalLoading
import fr.keyz.components.PriorityBox

@Composable
fun DamageDetailsScreen(navController: NavController, propertyId: String?, leaseId: String, damageId: String) {
    val isOwner = LocalIsOwner.current
    val apiService = LocalApiService.current
    val viewModel = viewModel {
        DamageDetailsViewModel(apiService, navController)
    }
    val isLoading = viewModel.isLoading.collectAsState()
    val apiError = viewModel.apiError.collectAsState()
    val damage = viewModel.currentDamage.collectAsState()

    LaunchedEffect(damageId) {
        viewModel.getDamage(propertyId, leaseId, damageId)
    }

    BackHandler {
        navController.popBackStack()
    }

    InventoryLayout(
        testTag = "damageDetails",
        onExit = { navController.popBackStack() },
        customTitle = stringResource(R.string.damage_detail)
    ) {
        if (isLoading.value || damage.value == null) {
            InternalLoading()
        } else {
            Row(
                modifier = Modifier.fillMaxWidth(),
                verticalAlignment = Alignment.CenterVertically,
                horizontalArrangement = Arrangement.SpaceBetween
            ) {
                Text(text = damage.value!!.roomName)
                PriorityBox(damage.value!!.priority)
            }
        }
    }

}