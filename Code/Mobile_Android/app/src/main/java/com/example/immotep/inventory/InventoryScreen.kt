package com.example.immotep.inventory


import android.widget.Toast
import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.LocalApiService
import com.example.immotep.layouts.InventoryLayout
import com.example.immotep.R
import com.example.immotep.components.ErrorAlert
import com.example.immotep.inventory.loaderButton.LoaderInventoryViewModel
import com.example.immotep.inventory.rooms.RoomsScreen


@Composable
fun InventoryScreen(
    navController: NavController,
    propertyId: String,
    leaseId: String,
    loaderViewModel: LoaderInventoryViewModel
) {
    val apiService = LocalApiService.current
    val context = LocalContext.current

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
            println("rooms length = ${rooms.size}")
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
        /*
        if (cannotMakeExitInventory.value) {
            AlertDialog(
                shape = RoundedCornerShape(10.dp),
                onDismissRequest = { viewModel.closeCannotMakeExitInventory() },
                confirmButton = {
                    TextButton(onClick = { viewModel.closeCannotMakeExitInventory() }) {
                        Text(stringResource(R.string.understand))
                    }
                },
                title = {
                    Text(stringResource(R.string.cannot_make_exit_inventory))
                },
                text = {
                    Text(stringResource(R.string.cannot_exit_inventory_text))
                },

                )
        }
         */
        RoomsScreen(
            getRooms = { viewModel.getRooms() },
            addRoom = {
                viewModel.addRoom(
                    it,
                    {
                        Toast.makeText(context, cannotAddRoomText, Toast.LENGTH_LONG).show()
                    })
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
            propertyId = propertyId
        )
    }
}
