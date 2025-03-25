package com.example.immotep.components.inventory

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.Button
import androidx.compose.material.Text
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.OutlinedTextField
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.focus.FocusRequester
import androidx.compose.ui.focus.focusRequester
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.example.immotep.R


@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun AddRoomOrDetailModal(open: Boolean, addRoomOrDetail: (name : String) -> Unit, close: () -> Unit, isRoom : Boolean) {
    if (open) {
        val focusRequester = remember { FocusRequester() }
        var roomName by rememberSaveable { mutableStateOf("") }
        LaunchedEffect(Unit) {
            try {
                focusRequester.requestFocus()
            } catch (e : Exception) {
                println("Impossible to request focus")
            }
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
                        .fillMaxWidth().focusRequester(focusRequester).testTag("roomNameTextField")
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
                        onClick = { close() },
                        modifier = Modifier.testTag("addRoomModalCancel")
                    ) {
                        Text(stringResource(R.string.cancel))
                    }
                    Button(
                        shape = RoundedCornerShape(5.dp),
                        colors = androidx.compose.material.ButtonDefaults.buttonColors(
                            backgroundColor = MaterialTheme.colorScheme.tertiary,
                            contentColor = MaterialTheme.colorScheme.onPrimary
                        ),
                        onClick = { addRoomOrDetail(roomName) },
                        modifier = Modifier.testTag("addRoomModalConfirm")
                    ) {
                        Text(stringResource(if (isRoom) R.string.add_room else R.string.add_detail))
                    }
                }
            }
        }
    }
}
