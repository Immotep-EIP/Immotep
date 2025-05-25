package com.example.keyz.inventory


import android.widget.Toast
import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.keyz.LocalApiService
import com.example.keyz.R
import com.example.keyz.components.ErrorAlert
import com.example.keyz.inventory.loaderButton.LoaderInventoryViewModel
import com.example.keyz.inventory.rooms.RoomsScreen


@Composable
fun InventoryScreen(
    navController: NavController,
    propertyId: String,
    leaseId: String,
    loaderViewModel: LoaderInventoryViewModel
) {
    val apiService = LocalApiService.current
    val context = navController.context

    val viewModel: InventoryViewModel =
        viewModel {
            InventoryViewModel(navController, apiService = apiService)
        }

    val inventoryErrors = viewModel.inventoryErrors.collectAsState()
    val isLoading = loaderViewModel.isLoading.collectAsState()
    val oldReportId = loaderViewModel.oldReportId.collectAsState()

    val cannotAddRoomText = stringResource(R.string.cannot_add_room)
    val cannotAddDetailText = stringResource(R.string.cannot_add_detail)

    LaunchedEffect(propertyId, loaderViewModel, isLoading) {
        viewModel.setPropertyIdAndLeaseId(propertyId, leaseId)
        if (!isLoading.value) {
            val rooms = loaderViewModel.getRooms()
            viewModel.loadInventoryFromRooms(rooms)
        }
    }

    Column(
        modifier = Modifier.testTag("inventoryScreen")
    ) {
        if (inventoryErrors.value.getAllRooms) {
            ErrorAlert(null, null, stringResource(R.string.error_get_all_rooms))
        }
        if (inventoryErrors.value.getLastInventoryReport) {
            ErrorAlert(null, null, stringResource(R.string.error_get_last_inventory_report))
        }
        if (inventoryErrors.value.createInventoryReport) {
            ErrorAlert(null, null, stringResource(R.string.error_create_inventory_report))
        }
        if (inventoryErrors.value.errorRoomName != null) {
            Toast.makeText(context, inventoryErrors.value.errorRoomName, Toast.LENGTH_LONG)
                .show()
        }
        RoomsScreen(
            getRooms = { viewModel.getRooms() },
            addRoom =
            { name, type ->
                viewModel.addRoom(
                    name = name,
                    roomType = type,
                    onError = { Toast.makeText(context, cannotAddRoomText, Toast.LENGTH_LONG).show() }
                )
            },
            removeRoom = { viewModel.removeRoom(it) },
            editRoom = { room -> viewModel.editRoom(room) },
            closeInventory = {
                viewModel.onClose()
                navController.popBackStack()
            },
            oldReportId = oldReportId.value,
            confirmInventory = {
                viewModel.sendInventory(
                    oldReportId.value,
                    { rooms, reportId ->
                        loaderViewModel.setNewValueSetByCompletedInventory(rooms, reportId, navController.context)
                    }
                )
            },
            addDetail = { roomId, name ->
                viewModel.addFurnitureCall(roomId, name,
                    {
                        Toast.makeText(context, cannotAddDetailText, Toast.LENGTH_LONG).show()
                    })
            },
            navController = navController,
            propertyId = propertyId,
            leaseId = leaseId
        )
    }
}
