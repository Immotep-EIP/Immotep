package com.example.immotep.inventory.rooms

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.Button
import androidx.compose.material.Text
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.Check
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.immotep.components.NextInventoryButton
import com.example.immotep.inventory.Room
import com.example.immotep.layouts.InventoryLayout
import com.example.immotep.realProperty.PropertyBox
import com.example.immotep.R
import com.example.immotep.components.InventoryCenterAddButton

fun roomIsCompleted(room: Room): Boolean {
    for (detail in room.details) {
        if (!detail.completed) {
            return false
        }
    }
    return true
}


//todo mettre une poppup etes vous surs de quitter pas sauvegardé etc etc

@Composable
fun RoomsScreen(
    rooms: Array<Room>,
    addRoom: (String) -> Unit,
    removeRoom: (Int) -> Unit,
    editRoom: (Int, Room) -> Unit,
    closeInventory: () -> Unit
) {
    val viewModel: RoomsViewModel = viewModel(factory = RoomsViewModelFactory(rooms, addRoom, removeRoom, editRoom))

    val currentlyOpenRoomIndex = viewModel.currentlyOpenRoomIndex.collectAsState()
    if (currentlyOpenRoomIndex.value == null) {
        InventoryLayout(testTag = "roomsScreen", closeInventory) {
            Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.End) {
                Button(
                    shape = RoundedCornerShape(5.dp),
                    colors = androidx.compose.material.ButtonDefaults.buttonColors(backgroundColor = MaterialTheme.colorScheme.tertiary),
                    onClick = { }) {
                        Text(stringResource(R.string.edit))
                    }
            }
            Column {
                LazyColumn {
                    items(rooms) { room ->
                        NextInventoryButton(
                            leftIcon = if (roomIsCompleted(room)) Icons.Outlined.Check else null,
                            leftText = room.name,
                            onClick = { viewModel.openRoomPanel(rooms.indexOf(room)) },
                            testTag = "roomButton ${rooms.indexOf(room)}"
                        )
                    }
                }
                InventoryCenterAddButton(onClick = { viewModel.addRoom("testRoom") }, testTag = "addRoomButton")
            }
        }
    }
}