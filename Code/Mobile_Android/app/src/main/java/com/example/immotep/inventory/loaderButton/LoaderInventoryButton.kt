package com.example.immotep.inventory.loaderButton

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.ui.res.stringResource
import com.example.immotep.R
import com.example.immotep.ui.components.StyledButton

@Composable
fun LoaderInventoryButton(
    propertyId: String,
    currentLeaseId: String,
    setIsLoading: (Boolean) -> Unit,
    viewModel: LoaderInventoryViewModel
) {

    LaunchedEffect(propertyId) {
        viewModel.loadInventory(propertyId)
    }

    StyledButton(
        onClick = { viewModel.onClick(setIsLoading, propertyId, currentLeaseId) },
        text = stringResource(R.string.inventory_title),
        testTag = "startInventory"
    )
}