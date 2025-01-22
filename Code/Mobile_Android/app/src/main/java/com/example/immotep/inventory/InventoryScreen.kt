package com.example.immotep.inventory


import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.TurnLeft
import androidx.compose.material.icons.outlined.TurnRight
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.res.stringResource
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.layouts.InventoryLayout
import com.example.immotep.R
import com.example.immotep.components.inventory.NextInventoryButton
import com.example.immotep.inventory.rooms.RoomsScreen


@Composable
fun InventoryScreen(
    navController: NavController,
    propertyId: String,
) {
    val viewModel: InventoryViewModel = viewModel(factory = InventoryViewModelFactory(navController, propertyId))
    val inventoryOpen = viewModel.inventoryOpen.collectAsState()
    val oldReportId = viewModel.oldReportId.collectAsState()

    LaunchedEffect(Unit) {
        viewModel.getBaseRooms()
    }
    if (inventoryOpen.value == InventoryOpenValues.CLOSED) {
        InventoryLayout(testTag = "inventoryScreen", { navController.popBackStack() }) {
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
            addRoom = { viewModel.addRoom(it) },
            removeRoom = { viewModel.removeRoom(it) },
            editRoom = { index, room -> viewModel.editRoom(index, room) },
            closeInventory = {
                viewModel.onClose()
                viewModel.setInventoryOpen(InventoryOpenValues.CLOSED)
            },
            oldReportId = if (inventoryOpen.value == InventoryOpenValues.EXIT) oldReportId.value else null,
            confirmInventory = { viewModel.sendInventory() },
            addDetail = { roomId, name -> viewModel.addFurniture(roomId, name) },
            navController = navController,
            propertyId = propertyId
        )
    }
}