package com.example.immotep.inventory.rooms

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.AlertDialog
import androidx.compose.material.Button
import androidx.compose.material.Text
import androidx.compose.material.TextButton
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.Check
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.components.inventory.NextInventoryButton
import com.example.immotep.inventory.Room
import com.example.immotep.layouts.InventoryLayout
import com.example.immotep.R
import com.example.immotep.components.InitialFadeIn
import com.example.immotep.components.InventoryCenterAddButton
import com.example.immotep.components.inventory.AddRoomOrDetailModal
import com.example.immotep.inventory.roomDetails.RoomDetailsScreen


@Composable
fun RoomsScreen(
    getRooms: () -> Array<Room>,
    addRoom: suspend (String) -> String?,
    addDetail: suspend (roomId : String, name : String) -> String?,
    removeRoom: (String) -> Unit,
    editRoom: (Room) -> Unit,
    closeInventory: () -> Unit,
    confirmInventory: () -> Boolean,
    oldReportId : String?,
    navController: NavController,
    propertyId: String
) {
    val viewModel: RoomsViewModel = viewModel {
        RoomsViewModel(
            getRooms = getRooms,
            addRoom = addRoom,
            removeRoom = removeRoom,
            editRoom = editRoom,
            closeInventory = closeInventory,
            confirmInventory = confirmInventory
        )
    }

    val currentlyOpenRoom = viewModel.currentlyOpenRoom.collectAsState()
    var exitPopUpOpen by rememberSaveable { mutableStateOf(false) }
    var confirmPopUpOpen by rememberSaveable { mutableStateOf(false) }
    var addRoomModalOpen by rememberSaveable { mutableStateOf(false) }

    val showNotCompletedRooms = viewModel.showNotCompletedRooms.collectAsState()

    LaunchedEffect(Unit) {
        viewModel.handleBaseRooms()
    }

    if (currentlyOpenRoom.value == null) {
        InventoryLayout(testTag = "roomsScreen", { exitPopUpOpen = true }) {
            if (exitPopUpOpen) {
                AlertDialog(
                    shape = RoundedCornerShape(10.dp),
                    onDismissRequest = { exitPopUpOpen = false },
                    confirmButton = {
                        TextButton(onClick = { exitPopUpOpen = false; viewModel.onClose()}) {
                            Text(stringResource(R.string.exit))
                        }
                                    },
                    dismissButton = {
                        TextButton(onClick = { exitPopUpOpen = false }) {
                            Text(stringResource(R.string.cancel))
                        }
                    },
                    title = {
                        Text(stringResource(R.string.are_you_sure_exit))
                    },
                    text = {
                        Text(stringResource(R.string.not_saved_modifications))
                    },

                )
            }
            if (confirmPopUpOpen) {
                AlertDialog(
                    shape = RoundedCornerShape(10.dp),
                    onDismissRequest = { confirmPopUpOpen = false },
                    confirmButton = {
                        TextButton(onClick = { confirmPopUpOpen = false; viewModel.onConfirmInventory() }) {
                            Text(stringResource(R.string.confirm))
                        }
                    },
                    dismissButton = {
                        TextButton(onClick = { confirmPopUpOpen = false }) {
                            Text(stringResource(R.string.cancel))
                        }
                    },
                    title = {
                        Text(stringResource(R.string.confirm_inventory_end))
                    },
                    text = {
                        Text(stringResource(R.string.not_forget))
                    },

                    )
            }
            AddRoomOrDetailModal(
                open = addRoomModalOpen,
                addRoomOrDetail = { viewModel.addARoom(it); addRoomModalOpen = false },
                close = { addRoomModalOpen = false },
                isRoom = true
            )
            InitialFadeIn {
                Column {
                    Row(
                        modifier = Modifier.fillMaxWidth(),
                        horizontalArrangement = Arrangement.SpaceBetween
                    ) {
                        Button(
                            shape = RoundedCornerShape(5.dp),
                            colors = androidx.compose.material.ButtonDefaults.buttonColors(
                                backgroundColor = MaterialTheme.colorScheme.secondary,
                                contentColor = MaterialTheme.colorScheme.onPrimary
                            ),
                            onClick = { confirmPopUpOpen = true },
                            modifier = Modifier.testTag("confirmInventoryButton")
                            ) {
                            Text(stringResource(R.string.confirm_inventory))
                        }
                        Button(
                            shape = RoundedCornerShape(5.dp),
                            colors = androidx.compose.material.ButtonDefaults.buttonColors(
                                backgroundColor = MaterialTheme.colorScheme.tertiary,
                                contentColor = MaterialTheme.colorScheme.onPrimary
                            ),
                            modifier = Modifier.testTag("editInventoryButton"),
                            onClick = { }) {
                            Text(stringResource(R.string.edit))
                        }
                    }
                    Column {
                        LazyColumn {
                            items(viewModel.allRooms) { room ->
                                NextInventoryButton(
                                    leftIcon = if (room.completed) Icons.Outlined.Check else null,
                                    leftText = room.name,
                                    onClick = {
                                        viewModel.openRoomPanel(room)
                                    },
                                    testTag = "roomButton ${room.id}",
                                    error = !room.completed && showNotCompletedRooms.value
                                )
                            }
                        }
                        InventoryCenterAddButton(
                            onClick = { addRoomModalOpen = true },
                            testTag = "addRoomButton"
                        )
                    }
                }
            }
        }
    } else {
        RoomDetailsScreen(
            closeRoomPanel = { viewModel.closeRoomPanel(it) },
            baseRoom = currentlyOpenRoom.value!!,
            oldReportId = oldReportId,
            addDetail = addDetail,
            navController = navController,
            propertyId = propertyId
        )
    }
}