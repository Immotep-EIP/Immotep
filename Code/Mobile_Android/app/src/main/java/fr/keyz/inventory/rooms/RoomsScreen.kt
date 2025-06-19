package fr.keyz.inventory.rooms

import androidx.activity.compose.BackHandler
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.AlertDialog
import androidx.compose.material.Button
import androidx.compose.material.ButtonDefaults
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
import fr.keyz.components.inventory.NextInventoryButton
import fr.keyz.inventory.Room
import fr.keyz.layouts.InventoryLayout
import fr.keyz.R
import fr.keyz.apiCallerServices.RoomType
import fr.keyz.components.InitialFadeIn
import fr.keyz.components.InventoryCenterAddButton
import fr.keyz.components.inventory.AddRoomOrDetailModal
import fr.keyz.components.inventory.EditRoomOrDetailModal
import fr.keyz.inventory.roomDetails.RoomDetailsScreen


@Composable
fun RoomsScreen(
    getRooms: () -> Array<Room>,
    addRoom: suspend (String, RoomType) -> String?,
    addDetail: suspend (roomId : String, name : String) -> String?,
    removeRoom: (String) -> Unit,
    editRoom: (Room) -> Unit,
    closeInventory: () -> Unit,
    confirmInventory: () -> Boolean,
    oldReportId : String?,
    navController: NavController,
    propertyId: String,
    leaseId : String
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
    var editOpen by rememberSaveable { mutableStateOf(false) }
    var editRoomOpen by rememberSaveable { mutableStateOf<String?>(null) }

    val showNotCompletedRooms = viewModel.showNotCompletedRooms.collectAsState()

    BackHandler {
        exitPopUpOpen = true
    }

    LaunchedEffect(Unit) {
        viewModel.handleBaseRooms()
    }

    if (currentlyOpenRoom.value == null) {
        InventoryLayout(testTag = "roomsScreen", { exitPopUpOpen = true }) {
            if (exitPopUpOpen) {
                AlertDialog(
                    shape = RoundedCornerShape(10.dp),
                    backgroundColor = MaterialTheme.colorScheme.background,
                    onDismissRequest = { exitPopUpOpen = false },
                    confirmButton = {
                        TextButton(onClick = { exitPopUpOpen = false; viewModel.onClose()}) {
                            Text(
                                stringResource(R.string.exit),
                                color = MaterialTheme.colorScheme.secondary
                            )
                        }
                                    },
                    dismissButton = {
                        TextButton(onClick = { exitPopUpOpen = false }) {
                            Text(
                                stringResource(R.string.cancel),
                                color = MaterialTheme.colorScheme.secondary
                            )
                        }
                    },
                    title = {
                        Text(
                            stringResource(R.string.are_you_sure_exit),
                            color = MaterialTheme.colorScheme.onPrimaryContainer
                        )
                    },
                    text = {
                        Text(
                            stringResource(R.string.not_saved_modifications),
                            color = MaterialTheme.colorScheme.onPrimaryContainer
                        )
                    },

                )
            }
            if (confirmPopUpOpen) {
                AlertDialog(
                    shape = RoundedCornerShape(10.dp),
                    onDismissRequest = { confirmPopUpOpen = false },
                    backgroundColor = MaterialTheme.colorScheme.background,
                    confirmButton = {
                        TextButton(onClick = { confirmPopUpOpen = false; viewModel.onConfirmInventory() }) {
                            Text(
                                stringResource(R.string.confirm),
                                color = MaterialTheme.colorScheme.secondary
                                )
                        }
                    },
                    dismissButton = {
                        TextButton(onClick = { confirmPopUpOpen = false }) {
                            Text(
                                stringResource(R.string.cancel),
                                color = MaterialTheme.colorScheme.secondary
                            )
                        }
                    },
                    title = {
                        Text(
                            stringResource(R.string.confirm_inventory_end),
                            color = MaterialTheme.colorScheme.onPrimaryContainer
                        )
                    },
                    text = {
                        Text(
                            stringResource(R.string.not_forget),
                            color = MaterialTheme.colorScheme.onPrimaryContainer
                        )
                    },

                    )
            }
            AddRoomOrDetailModal(
                open = addRoomModalOpen,
                addRoomOrDetail =
                { name, type ->
                    type?: return@AddRoomOrDetailModal
                    viewModel.addARoom(name, type)
                    addRoomModalOpen = false
                },
                close = { addRoomModalOpen = false },
                isRoom = true,
                addRoomType = true
            )
            EditRoomOrDetailModal(
                currentRoomOrDetailId = editRoomOpen,
                deleteRoomOrDetail = { viewModel.handleRemoveRoom(it) },
                close = { editRoomOpen = null },
            )
            InitialFadeIn {
                Column {
                    Row(
                        modifier = Modifier.fillMaxWidth(),
                        horizontalArrangement = Arrangement.SpaceBetween
                    ) {
                        Button(
                            shape = RoundedCornerShape(5.dp),
                            colors = ButtonDefaults.buttonColors(
                                backgroundColor = MaterialTheme.colorScheme.secondary,
                                contentColor = MaterialTheme.colorScheme.onSecondary
                            ),
                            onClick = { confirmPopUpOpen = true },
                            modifier = Modifier.testTag("confirmInventoryButton")
                            ) {
                            Text(
                                stringResource(R.string.confirm_inventory),
                                )
                        }
                        Button(
                            shape = RoundedCornerShape(5.dp),
                            colors = ButtonDefaults.buttonColors(
                                backgroundColor =
                                if (editOpen) MaterialTheme.colorScheme.error else MaterialTheme.colorScheme.secondary,
                                contentColor = MaterialTheme.colorScheme.onSecondary
                            ),
                            modifier = Modifier.testTag("editInventoryButton"),
                            onClick = { editOpen = !editOpen }) {
                            Text(stringResource(if (editOpen) R.string.close_edit else R.string.edit))
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
                                    error = !room.completed && showNotCompletedRooms.value,
                                    editOpen = editOpen,
                                    onClickEdit = { editRoomOpen = room.id }
                                )
                            }
                        }
                        if (editOpen) {
                            InventoryCenterAddButton(
                                onClick = { addRoomModalOpen = true },
                                testTag = "addRoomButton"
                            )
                        }
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
            propertyId = propertyId,
            leaseId = leaseId,
            removeDetail = {}
        )
    }
}