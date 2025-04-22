package com.example.immotep.components.inventory

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.focus.FocusRequester
import androidx.compose.ui.focus.focusRequester
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.example.immotep.R
import com.example.immotep.ui.components.OutlinedTextField
import com.example.immotep.ui.components.StyledButton
import kotlinx.coroutines.coroutineScope
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch


@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun AddRoomOrDetailModal(
    open: Boolean,
    addRoomOrDetail: suspend (name : String) -> Unit,
    close: () -> Unit, isRoom : Boolean
) {
    if (open) {
        val focusRequester = remember { FocusRequester() }
        val scope = rememberCoroutineScope()
        var roomName by rememberSaveable { mutableStateOf("") }
        var error by rememberSaveable { mutableStateOf<String?>(null) }

        val onSubmit = {
            scope.launch {
                try {
                    error = null
                    addRoomOrDetail(roomName)
                } catch (e : Exception) {
                    error = e.message
                }
            }
        }

        LaunchedEffect(Unit) {
            try {
                delay(500)
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
                    label = stringResource(if (isRoom) R.string.room_name else R.string.detail_name),
                    modifier = Modifier
                        .fillMaxWidth().focusRequester(focusRequester).testTag("roomNameTextField"),
                    errorMessage = when (error) {
                        null -> null
                        "room_already_exists" -> stringResource(R.string.room_already_exists)
                        "detail_already_exists" -> stringResource(R.string.detail_already_exists)
                        "impossible_to_add_room" -> stringResource(R.string.impossible_to_add_room)
                        "impossible_to_add_detail" -> stringResource(R.string.impossible_to_add_detail)
                        else -> stringResource(R.string.basic_error)
                    }
                )
                Row(
                    modifier = Modifier.fillMaxWidth(),
                    horizontalArrangement = Arrangement.SpaceBetween,
                    verticalAlignment = Alignment.CenterVertically
                ) {
                    Button(
                        shape = RoundedCornerShape(5.dp),
                        colors = androidx.compose.material3.ButtonDefaults.buttonColors(
                            containerColor = MaterialTheme.colorScheme.errorContainer,
                            contentColor = MaterialTheme.colorScheme.onError
                        ),
                        onClick = { close() },
                        modifier = Modifier.testTag("addRoomModalCancel")
                    ) {
                        Text(stringResource(R.string.cancel))
                    }
                    StyledButton(
                        onClick = { onSubmit() },
                        modifier = Modifier.testTag("addRoomModalConfirm"),
                        text = stringResource(if (isRoom) R.string.add_room else R.string.add_detail)
                    )
                }
            }
        }
    }
}
