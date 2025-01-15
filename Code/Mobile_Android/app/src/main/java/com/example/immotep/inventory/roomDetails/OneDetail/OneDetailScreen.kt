package com.example.immotep.inventory.roomDetails.OneDetail

import androidx.compose.foundation.Image
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.text.BasicTextField
import androidx.compose.material.Button
import androidx.compose.material.DropdownMenu
import androidx.compose.material.DropdownMenuItem
import androidx.compose.material.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material.Text
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.ArrowDropDown
import androidx.compose.material3.TextField
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.immotep.R
import com.example.immotep.components.AddingPicturesCarousel
import com.example.immotep.components.InitialFadeIn
import com.example.immotep.inventory.RoomDetail
import com.example.immotep.layouts.InventoryLayout
import com.example.immotep.ui.components.DropDown
import com.example.immotep.ui.components.DropDownItem
import com.example.immotep.ui.components.OutlinedTextField



@Composable
fun OneDetailScreen(onModifyDetail : (detailIndex : Int, detail : RoomDetail) -> Unit, index : Int, baseDetail : RoomDetail, isExit : Boolean) {
    val viewModel : OneDetailViewModel = viewModel()
    val detailValue = viewModel.detail.collectAsState()
    val detailError = viewModel.errors.collectAsState()
    LaunchedEffect(Unit) {
        viewModel.reset(baseDetail)
    }
    InventoryLayout(
        testTag = "oneDetailScreen",
        { viewModel.onClose(onModifyDetail, index, isExit) }
    ) {
        InitialFadeIn {
            Column {
                Text(if (isExit) stringResource(R.string.entry_pictures) else stringResource(R.string.pictures))
                AddingPicturesCarousel(pictures = viewModel.picture, addPicture = { uri -> viewModel.addPicture(uri) })
                if (detailError.value.picture) Text(stringResource(R.string.add_picture_error),
                    color = MaterialTheme.colorScheme.error,
                    modifier = Modifier.padding(top = 10.dp))
                if (isExit) {
                    Text(stringResource(R.string.exit_pictures))
                    AddingPicturesCarousel(pictures = viewModel.exitPicture, addPicture = { uri -> viewModel.addExitPicture(uri) })
                    if (detailError.value.exitPicture) {
                        Text(
                            stringResource(R.string.add_picture_error),
                            color = MaterialTheme.colorScheme.error,
                            modifier = Modifier.padding(top = 10.dp)
                        )
                    }
                }
                OutlinedTextField(
                    value = detailValue.value.comment,
                    onValueChange = { newVal -> viewModel.setComment(newVal) },
                    label = stringResource(R.string.comment),
                    minLines = 4,
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 10.dp),
                    errorMessage = if (detailError.value.comment) stringResource(R.string.comment_error) else null
                )
                Text(stringResource(R.string.status), modifier = Modifier.padding(top = 10.dp))
                DropDown(
                    items = listOf(
                        DropDownItem(stringResource(R.string.new_state), "new"),
                        DropDownItem(stringResource(R.string.very_good_state), "very_good"),
                        DropDownItem(stringResource(R.string.good_state), "good_state"),
                        DropDownItem(stringResource(R.string.ok_state), "ok"),
                        DropDownItem(stringResource(R.string.degraded_state), "degraded"),
                        DropDownItem(stringResource(R.string.bad_state), "bad"),
                        DropDownItem(stringResource(R.string.broken_state), "broken"),
                    ),
                    selectedItem = detailValue.value.status,
                    onItemSelected = { newVal -> viewModel.setStatus(newVal) }
                )
                if (detailError.value.status) Text(stringResource(R.string.status_error),
                    modifier = Modifier.padding(top = 10.dp),
                    color = MaterialTheme.colorScheme.error)
                Column(modifier = Modifier.fillMaxWidth(), horizontalAlignment = Alignment.CenterHorizontally) {
                    Button(
                        shape = RoundedCornerShape(5.dp),
                        modifier = Modifier.padding(top = 10.dp),
                        colors = androidx.compose.material.ButtonDefaults.buttonColors(
                            backgroundColor = MaterialTheme.colorScheme.tertiary,
                            contentColor = MaterialTheme.colorScheme.onPrimary
                        ),
                        onClick = { },
                    ) {
                        Text(stringResource(if (isExit) R.string.compare_images else R.string.analyze_pictures))
                    }
                    Button(
                        shape = RoundedCornerShape(5.dp),
                        modifier = Modifier.padding(top = 10.dp),
                        colors = androidx.compose.material.ButtonDefaults.buttonColors(
                            backgroundColor = MaterialTheme.colorScheme.tertiary,
                            contentColor = MaterialTheme.colorScheme.onPrimary
                        ),
                        onClick = { viewModel.onConfirm(onModifyDetail, index, isExit) },
                    ) {
                        Text(stringResource(R.string.validate))
                    }
                }
            }
        }
    }
}
