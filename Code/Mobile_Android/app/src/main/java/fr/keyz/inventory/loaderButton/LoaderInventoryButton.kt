package fr.keyz.inventory.loaderButton

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.res.stringResource
import fr.keyz.R
import fr.keyz.ui.components.StyledButton

@Composable
fun LoaderInventoryButton(
    propertyId: String,
    currentLeaseId: String,
    setIsLoading: (Boolean) -> Unit,
    viewModel: LoaderInventoryViewModel
) {
    val isLoading = viewModel.isLoading.collectAsState()
    LaunchedEffect(propertyId) {
        viewModel.loadInventory(propertyId)
    }

    StyledButton(
        onClick = { viewModel.onClick(setIsLoading, propertyId, currentLeaseId) },
        text = stringResource(R.string.inventory_title),
        testTag = "startInventory",
        isLoading = isLoading.value
    )
}