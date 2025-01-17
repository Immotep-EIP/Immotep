package com.example.immotep.inventory.rooms

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.lazy.itemsIndexed
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.AlertDialog
import androidx.compose.material.Button
import androidx.compose.material.Text
import androidx.compose.material.TextButton
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.Check
import androidx.compose.material.icons.outlined.Warning
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.OutlinedTextField
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.focus.FocusRequester
import androidx.compose.ui.focus.focusRequester
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.compose.ui.window.Popup
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.immotep.components.NextInventoryButton
import com.example.immotep.inventory.Room
import com.example.immotep.layouts.InventoryLayout
import com.example.immotep.realProperty.PropertyBox
import com.example.immotep.R
import com.example.immotep.components.InitialFadeIn
import com.example.immotep.components.InventoryCenterAddButton
import com.example.immotep.inventory.roomDetails.RoomDetailsScreen

fun roomIsCompleted(room: Room): Boolean {
    if (room.details.isEmpty()) return false
    for (detail in room.details) {
        if (!detail.completed) {
            return false
        }
    }
    return true
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun AddRoomOrDetailModal(open: Boolean, addRoomOrDetail: (name : String) -> Unit, close: () -> Unit, isRoom : Boolean) {
    if (open) {
        val focusRequester = remember { FocusRequester() }
        var roomName by rememberSaveable { mutableStateOf("") }
        LaunchedEffect(Unit) {
            focusRequester.requestFocus()
        }
        ModalBottomSheet(
            onDismissRequest = close,
            modifier = Modifier
                .testTag("addRoomModal")

        ) {
            Column(modifier = Modifier.padding(start = 10.dp, end = 10.dp, top = 20.dp, bottom = 20.dp)) {
                OutlinedTextField(
                    value = roomName,
                    onValueChange = { roomName = it },
                    label = { Text(stringResource(if (isRoom) R.string.room_name else R.string.detail_name)) },
                    modifier = Modifier
                        .fillMaxWidth().focusRequester(focusRequester)
                )
                Row(
                    modifier = Modifier.fillMaxWidth(),
                    horizontalArrangement = Arrangement.SpaceBetween
                ) {
                    Button(
                        shape = RoundedCornerShape(5.dp),
                        colors = androidx.compose.material.ButtonDefaults.buttonColors(
                            backgroundColor = MaterialTheme.colorScheme.errorContainer,
                            contentColor = MaterialTheme.colorScheme.onError
                        ),
                        onClick = { close() }) {
                        Text(stringResource(R.string.cancel))
                    }
                    Button(
                        shape = RoundedCornerShape(5.dp),
                        colors = androidx.compose.material.ButtonDefaults.buttonColors(
                            backgroundColor = MaterialTheme.colorScheme.tertiary,
                            contentColor = MaterialTheme.colorScheme.onPrimary
                        ),
                        onClick = { addRoomOrDetail(roomName) }) {
                        Text(stringResource(if (isRoom) R.string.add_room else R.string.add_detail))
                    }
                }
            }
        }
    }
}


@Composable
fun RoomsScreen(
    getRooms: () -> Array<Room>,
    addRoom: (String) -> Unit,
    removeRoom: (Int) -> Unit,
    editRoom: (Int, Room) -> Unit,
    closeInventory: () -> Unit,
    confirmInventory: () -> Boolean,
    isExit : Boolean
) {
    val viewModel: RoomsViewModel = viewModel(
        factory = RoomsViewModelFactory(getRooms, addRoom, removeRoom, editRoom, closeInventory, confirmInventory)
    )

    val currentlyOpenRoomIndex = viewModel.currentlyOpenRoomIndex.collectAsState()
    var exitPopUpOpen by rememberSaveable { mutableStateOf(false) }
    var confirmPopUpOpen by rememberSaveable { mutableStateOf(false) }
    var addRoomModalOpen by rememberSaveable { mutableStateOf(false) }

    val showNotCompletedRooms = viewModel.showNotCompletedRooms.collectAsState()

    LaunchedEffect(Unit) {
        viewModel.handleBaseRooms()
    }

    if (currentlyOpenRoomIndex.value == null) {
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
                            onClick = { confirmPopUpOpen = true }) {
                            Text(stringResource(R.string.confirm_inventory))
                        }
                        Button(
                            shape = RoundedCornerShape(5.dp),
                            colors = androidx.compose.material.ButtonDefaults.buttonColors(
                                backgroundColor = MaterialTheme.colorScheme.tertiary,
                                contentColor = MaterialTheme.colorScheme.onPrimary
                            ),
                            onClick = { }) {
                            Text(stringResource(R.string.edit))
                        }
                    }
                    Column {
                        LazyColumn {
                            itemsIndexed(viewModel.allRooms) { index, room ->
                                NextInventoryButton(
                                    leftIcon = if (roomIsCompleted(room)) Icons.Outlined.Check else null,
                                    leftText = room.name,
                                    onClick = {
                                        viewModel.openRoomPanel(index)
                                    },
                                    testTag = "roomButton $index",
                                    error = !roomIsCompleted(room) && showNotCompletedRooms.value
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
            closeRoomPanel = { roomIndex, details ->
                viewModel.closeRoomPanel(roomIndex, details)
            },
            roomDetails = viewModel.allRooms[currentlyOpenRoomIndex.value!!].details,
            roomIndex = currentlyOpenRoomIndex.value!!,
            isExit = isExit
        )
    }
}