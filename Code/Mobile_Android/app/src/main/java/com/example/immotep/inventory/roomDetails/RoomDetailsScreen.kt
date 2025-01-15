package com.example.immotep.inventory.roomDetails

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.lazy.itemsIndexed
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.Button
import androidx.compose.material.Text
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.Check
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.immotep.R
import com.example.immotep.components.InitialFadeIn
import com.example.immotep.components.InventoryCenterAddButton
import com.example.immotep.components.NextInventoryButton
import com.example.immotep.inventory.RoomDetail
import com.example.immotep.inventory.roomDetails.OneDetail.OneDetailScreen
import com.example.immotep.inventory.rooms.roomIsCompleted
import com.example.immotep.layouts.InventoryLayout

@Composable
fun RoomDetailsScreen(
    closeRoomPanel : (roomIndex: Int, details: Array<RoomDetail>) -> Unit,
    roomDetails: Array<RoomDetail>,
    roomIndex: Int,
    isExit : Boolean
) {
    val viewModel: RoomDetailsViewModel = viewModel(factory = RoomDetailsViewModelFactory(closeRoomPanel))

    val currentlyOpenRoomIndex = viewModel.currentlyOpenDetailIndex.collectAsState()

    LaunchedEffect(Unit) {
        viewModel.addBaseDetails(roomDetails)
    }
    if (currentlyOpenRoomIndex.value == null) {
        InventoryLayout(testTag = "roomsScreen", { viewModel.onClose(roomIndex) }) {
            InitialFadeIn {
                Column {
                    Row(
                        modifier = Modifier.fillMaxWidth(),
                        horizontalArrangement = Arrangement.End
                    ) {
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
                            itemsIndexed(viewModel.details) { index, detail ->
                                NextInventoryButton(
                                    leftIcon = if (detail.completed) Icons.Outlined.Check else null,
                                    leftText = detail.name,
                                    onClick = { viewModel.onOpenDetail(index) },
                                    testTag = "detailButton $index"
                                )
                            }
                        }
                        InventoryCenterAddButton(
                            onClick = { viewModel.addDetail("testDetail") },
                            testTag = "addDetailsButton"
                        )
                    }
                }
            }
        }
    } else {
        OneDetailScreen(
            onModifyDetail = { detailIndex, detail ->
                viewModel.onModifyDetail(detailIndex, detail)
            },
            index = currentlyOpenRoomIndex.value!!,
            baseDetail = viewModel.details[currentlyOpenRoomIndex.value!!],
            isExit = isExit
        )
    }
}