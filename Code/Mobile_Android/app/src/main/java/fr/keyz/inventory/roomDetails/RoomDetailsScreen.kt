package fr.keyz.inventory.roomDetails

import androidx.activity.compose.BackHandler
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.Button
import androidx.compose.material.ButtonDefaults
import androidx.compose.material.Text
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.Check
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import fr.keyz.R
import fr.keyz.components.InitialFadeIn
import fr.keyz.components.InventoryCenterAddButton
import fr.keyz.components.inventory.AddRoomOrDetailModal
import fr.keyz.components.inventory.EditRoomOrDetailModal
import fr.keyz.components.inventory.NextInventoryButton
import fr.keyz.inventory.Room
import fr.keyz.inventory.roomDetails.EndRoomDetails.EndRoomDetailsScreen
import fr.keyz.inventory.roomDetails.OneDetail.OneDetailScreen
import fr.keyz.layouts.InventoryLayout

fun roomIsCompleted(room: Room): Boolean {
    for (detail in room.details) {
        if (!detail.completed) {
            return false
        }
    }
    return true
}

@Composable
fun RoomDetailsScreen(
    baseRoom: Room,
    closeRoomPanel : (room : Room) -> Unit,
    addDetail: suspend (roomId : String, name : String) -> String?,
    removeDetail: (String) -> Unit,
    oldReportId : String?,
    navController: NavController,
    propertyId: String,
    leaseId : String
) {
    val viewModel: RoomDetailsViewModel = viewModel {
        RoomDetailsViewModel(closeRoomPanel, addDetail, removeDetail)
    }

    val currentlyOpenDetail = viewModel.currentlyOpenDetail.collectAsState()
    var addDetailModalOpen by rememberSaveable { mutableStateOf(false) }
    var endRoomDetailsScreenOpen by rememberSaveable { mutableStateOf(false) }
    var editOpen by rememberSaveable { mutableStateOf(false) }
    var editDetailOpen by rememberSaveable { mutableStateOf<String?>(null) }

    BackHandler {
        viewModel.onClose(baseRoom)
    }

    LaunchedEffect(Unit) {
        viewModel.addBaseDetails(baseRoom.details)
    }

    AddRoomOrDetailModal(
        open = addDetailModalOpen,
        addRoomOrDetail =
        { name, _ ->
            viewModel.addDetailToRoomDetailPage(name, baseRoom.id)
            addDetailModalOpen = false
        },
        close = { addDetailModalOpen = false },
        isRoom = false
    )
    EditRoomOrDetailModal(
        currentRoomOrDetailId = editDetailOpen,
        deleteRoomOrDetail = { viewModel.handleRemoveDetail(it) },
        close = { editDetailOpen = null },
    )
    if (endRoomDetailsScreenOpen) {
        EndRoomDetailsScreen(
            room = baseRoom,
            closeRoomPanel = closeRoomPanel,
            oldReportId = oldReportId,
            propertyId = propertyId,
            navController = navController,
            isOpen = endRoomDetailsScreenOpen,
            setOpen = { endRoomDetailsScreenOpen = it },
            newDetails = viewModel.details.toTypedArray(),
            leaseId = leaseId
        )
        return
    }
    if (currentlyOpenDetail.value == null) {
        InventoryLayout(testTag = "roomsDetailsScreen", { viewModel.onClose(baseRoom) }) {
            InitialFadeIn {
                Column {
                    Row(
                        modifier = Modifier.fillMaxWidth(),
                        horizontalArrangement = Arrangement.End
                    ) {
                        Button(
                            shape = RoundedCornerShape(5.dp),
                            colors = ButtonDefaults.buttonColors(
                                backgroundColor =
                                    if (editOpen) MaterialTheme.colorScheme.error else MaterialTheme.colorScheme.secondary,
                                contentColor = MaterialTheme.colorScheme.onSecondary
                            ),
                            modifier = Modifier.testTag("editRoomsDetails"),
                            onClick = { editOpen = !editOpen }) {
                            Text(stringResource(if (editOpen) R.string.close_edit else R.string.edit))
                        }
                    }
                    Column {
                        LazyColumn {
                            items(viewModel.details) { detail ->
                                NextInventoryButton(
                                    leftIcon = if (detail.completed) Icons.Outlined.Check else null,
                                    leftText = detail.name,
                                    onClick = { viewModel.onOpenDetail(detail) },
                                    testTag = "detailButton ${detail.id}",
                                    editOpen = editOpen,
                                    onClickEdit = {}
                                )
                            }
                        }
                        if (editOpen) {
                            InventoryCenterAddButton(
                                onClick = { addDetailModalOpen = true },
                                testTag = "addDetailsButton"
                            )
                        }
                    }
                    if (roomIsCompleted(Room(id = baseRoom.id, details = viewModel.details.toTypedArray(), name = baseRoom.name))) {
                        Column(modifier = Modifier.fillMaxWidth(), horizontalAlignment = Alignment.CenterHorizontally) {
                            Button(
                                shape = RoundedCornerShape(5.dp),
                                modifier = Modifier.padding(top = 10.dp),
                                colors = ButtonDefaults.buttonColors(
                                    backgroundColor = MaterialTheme.colorScheme.secondary,
                                    contentColor = MaterialTheme.colorScheme.onSecondary
                                ),
                                onClick = { endRoomDetailsScreenOpen = true },
                            ) {
                                Text("${stringResource(R.string.complete_room)} ${baseRoom.name}")
                            }
                        }
                    }
                }
            }
        }
    } else {
        OneDetailScreen(
            onModifyDetail = { detail ->
                viewModel.onModifyDetail(detail)
            },
            baseDetail = currentlyOpenDetail.value!!,
            oldReportId = oldReportId,
            navController = navController,
            propertyId = propertyId,
            leaseId = leaseId,
        )
    }
}