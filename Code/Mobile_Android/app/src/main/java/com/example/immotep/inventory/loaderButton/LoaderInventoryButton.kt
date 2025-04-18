package com.example.immotep.inventory.loaderButton

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.ui.res.stringResource
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.LocalApiService
import com.example.immotep.R
import com.example.immotep.ui.components.StyledButton

@Composable
fun LoaderInventoryButton(
    navController: NavController,
    propertyId: String,
    setIsLoading: (Boolean) -> Unit,
    viewModel: LoaderInventoryViewModel
) {

    LaunchedEffect(propertyId) {
        viewModel.loadInventory(propertyId)
    }

    StyledButton(
        onClick = { viewModel.onClick(setIsLoading, propertyId) },
        text = stringResource(R.string.inventory_title)
    )
}