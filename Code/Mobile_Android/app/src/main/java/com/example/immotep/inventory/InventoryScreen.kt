package com.example.immotep.inventory


import android.widget.Toast
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.AlertDialog
import androidx.compose.material.Text
import androidx.compose.material.TextButton
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.TurnLeft
import androidx.compose.material.icons.outlined.TurnRight
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.LocalApiService
import com.example.immotep.layouts.InventoryLayout
import com.example.immotep.R
import com.example.immotep.components.ErrorAlert
import com.example.immotep.components.inventory.NextInventoryButton
import com.example.immotep.inventory.rooms.RoomsScreen


@Composable
fun InventoryScreen(
    navController: NavController,
    propertyId: String,
) {
    val viewModel: InventoryViewModel =
        viewModel(
            factory = InventoryViewModelFactory(
                navController,
                propertyId,
                apiService = LocalApiService.current
            )
        )
    val context = LocalContext.current
    val inventoryOpen = viewModel.inventoryOpen.collectAsState()
    val oldReportId = viewModel.oldReportId.collectAsState()
    val cannotMakeExitInventory = viewModel.cannotMakeExitInventory.collectAsState()
    val inventoryErrors = viewModel.inventoryErrors.collectAsState()
    val cannotAddRoomText = stringResource(R.string.cannot_add_room)
    val cannotAddDetailText = stringResource(R.string.cannot_add_detail)
    LaunchedEffect(propertyId) {
        viewModel.getBaseRooms(
            propertyId = propertyId
        )
    }
    if (inventoryOpen.value == InventoryOpenValues.CLOSED) {
        InventoryLayout(testTag = "inventoryScreen", { navController.popBackStack() }) {
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
                Toast.makeText(context, inventoryErrors.value.errorRoomName, Toast.LENGTH_LONG).show()
            }
            if (cannotMakeExitInventory.value) {
                AlertDialog(
                    shape = RoundedCornerShape(10.dp),
                    onDismissRequest = { viewModel.closeCannotMakeExitInventory() },
                    confirmButton = {
                        TextButton(onClick = { viewModel.closeCannotMakeExitInventory()}) {
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
            NextInventoryButton(
                Icons.Outlined.TurnRight,
                stringResource(R.string.entry_inventory),
                { viewModel.setInventoryOpen(InventoryOpenValues.ENTRY) },
                testTag = "entryInventoryButton"
            )
            NextInventoryButton(
                Icons.Outlined.TurnLeft,
                stringResource(R.string.exit_inventory),
                { viewModel.setInventoryOpen(InventoryOpenValues.EXIT) },
                testTag = "exitInventoryButton"
            )
        }
    } else {
        RoomsScreen(
            getRooms = { viewModel.getRooms() },
            addRoom = { viewModel.addRoom(
                it,
                {
                    Toast.makeText(context, cannotAddRoomText, Toast.LENGTH_LONG).show()
                }) },
            removeRoom = { viewModel.removeRoom(it) },
            editRoom = { room -> viewModel.editRoom(room) },
            closeInventory = {
                viewModel.onClose()
                viewModel.setInventoryOpen(InventoryOpenValues.CLOSED)
            },
            oldReportId = if (inventoryOpen.value == InventoryOpenValues.EXIT) oldReportId.value else null,
            confirmInventory = { viewModel.sendInventory() },
            addDetail = { roomId, name -> viewModel.addFurnitureCall(roomId, name,
                {
                    Toast.makeText(context, cannotAddDetailText, Toast.LENGTH_LONG).show()
                }) },
            navController = navController,
            propertyId = propertyId
        )
    }
}