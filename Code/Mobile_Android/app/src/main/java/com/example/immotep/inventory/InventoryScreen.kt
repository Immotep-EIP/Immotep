package com.example.immotep.inventory

import androidx.compose.foundation.BorderStroke
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.Button
import androidx.compose.material.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.outlined.ArrowForwardIos
import androidx.compose.material.icons.outlined.ArrowForwardIos
import androidx.compose.material.icons.outlined.CameraIndoor
import androidx.compose.material.icons.outlined.TurnLeft
import androidx.compose.material.icons.outlined.TurnRight
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.layouts.InventoryLayout
import com.example.immotep.R
import com.example.immotep.components.NextInventoryButton
import com.example.immotep.inventory.rooms.RoomsScreen


@Composable
fun InventoryScreen(
    navController: NavController,
    propertyId: String,
) {
    val viewModel: InventoryViewModel = viewModel(factory = InventoryViewModelFactory(navController, propertyId))
    var inventoryOpen = viewModel.inventoryOpen.collectAsState()

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
            isExit = inventoryOpen.value == InventoryOpenValues.EXIT,
            confirmInventory = { viewModel.sendInventory() },
            addDetail = { roomId, name -> viewModel.addFurniture(roomId, name) },
            navController = navController,
            propertyId = propertyId
        )
    }
}