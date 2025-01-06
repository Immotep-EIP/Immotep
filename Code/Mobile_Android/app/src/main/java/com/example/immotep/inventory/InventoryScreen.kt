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

enum class InventoryOpenValues {
    ENTRY, EXIT, CLOSED
}

@Composable
fun InventoryScreen(
    navController: NavController,
    propertyId: String,
    viewModel: InventoryViewModel = viewModel()
) {
    var inventoryOpen by rememberSaveable { mutableStateOf(InventoryOpenValues.CLOSED) }
    if (inventoryOpen == InventoryOpenValues.CLOSED) {
        InventoryLayout(testTag = "inventoryScreen", { navController.popBackStack() }) {
            NextInventoryButton(
                Icons.Outlined.TurnRight,
                stringResource(R.string.entry_inventory),
                { inventoryOpen = InventoryOpenValues.ENTRY },
                testTag = "entryInventoryButton"
            )
            NextInventoryButton(
                Icons.Outlined.TurnLeft,
                stringResource(R.string.exit_inventory),
                { inventoryOpen = InventoryOpenValues.EXIT },
                testTag = "exitInventoryButton"
            )
        }
    } else {
        RoomsScreen(
            viewModel.rooms.toTypedArray(),
            { viewModel.addRoom(it) },
            { viewModel.removeRoom(it) },
            { index, room -> viewModel.editRoom(index, room) },
            {
                viewModel.onClose()
                inventoryOpen = InventoryOpenValues.CLOSED
            }
        )
    }
}