package fr.keyz.components

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.AlertDialog
import androidx.compose.material.Button
import androidx.compose.material.Text
import androidx.compose.material.TextButton
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import fr.keyz.R


@Composable
fun DeletePopUpConfirmation(close: () -> Unit, delete: () -> Unit, globalName: String, detailedName: String) {
    AlertDialog(
        shape = RoundedCornerShape(10.dp),
        backgroundColor = MaterialTheme.colorScheme.background,
        onDismissRequest = { close() },
        confirmButton = {
            TextButton(onClick = { delete(); close() }) {
                Text(stringResource(R.string.delete), color = MaterialTheme.colorScheme.error)
            }
        },
        title = {
            Text(
                "${stringResource(R.string.delete)} $globalName ?",
                color = MaterialTheme.colorScheme.onPrimaryContainer
            )
        },
        text = {
            Text(
                "${stringResource(R.string.sure_to_delete)} $detailedName ?",
                color = MaterialTheme.colorScheme.onPrimaryContainer
            )
        },

        )
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun DeletePopUp(
    open : Boolean,
    delete : () -> Unit,
    close : () -> Unit,
    globalName : String,
    detailedName : String
) {
    var confirmOpen by rememberSaveable { mutableStateOf(false) }
    if (open) {
        if (confirmOpen) {
            DeletePopUpConfirmation({confirmOpen = false; close()}, delete, globalName, detailedName)
        }
        ModalBottomSheet(
            onDismissRequest = close,
            modifier = Modifier
                .testTag("deletePopUp")

        ) {
            Column(modifier = Modifier.padding(top = 5.dp, bottom = 5.dp).align(Alignment.CenterHorizontally)) {
                Button(
                    shape = RoundedCornerShape(5.dp),
                    colors = androidx.compose.material.ButtonDefaults.buttonColors(
                        backgroundColor = MaterialTheme.colorScheme.errorContainer,
                        contentColor = MaterialTheme.colorScheme.onError
                    ),
                    onClick = { confirmOpen = true }) {
                    Text("${stringResource(R.string.delete)} $globalName")
                }
            }
        }
    }
}