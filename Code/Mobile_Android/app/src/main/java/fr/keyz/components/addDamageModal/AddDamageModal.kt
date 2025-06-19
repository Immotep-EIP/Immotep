package fr.keyz.components.addDamageModal

import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material.Text
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import fr.keyz.LocalApiService
import fr.keyz.apiCallerServices.DamagePriority
import fr.keyz.components.AddingPicturesCarousel
import fr.keyz.layouts.BigModalLayout
import fr.keyz.ui.components.OutlinedTextField
import fr.keyz.ui.components.DropDown
import fr.keyz.ui.components.DropDownItem
import fr.keyz.ui.components.StyledButton
import fr.keyz.R
import fr.keyz.apiCallerServices.Damage

@Composable
fun AddDamageModal(
    open : Boolean,
    onClose : () -> Unit,
    addDamage : (Damage) -> Unit,
    navController: NavController
) {
    val apiService = LocalApiService.current
    val viewModel = viewModel {
        AddDamageModalViewModel(apiService, navController)
    }
    val form = viewModel.form.collectAsState()
    val errors = viewModel.formError.collectAsState()
    LaunchedEffect(open) {
        viewModel.reset()
    }
    BigModalLayout(
        height = 0.8f,
        open = open,
        close = onClose,
        testTag = "addDamageModal"
    ) {
        AddingPicturesCarousel(
            uriPictures = viewModel.pictures,
            addPicture = { viewModel.addPicture(it) },
            removePicture = { viewModel.removePicture(it) },
            error = if (errors.value.pictures) stringResource(R.string.add_picture_error) else null,
        )
        OutlinedTextField(
            value = form.value.comment,
            onValueChange = { viewModel.setComment(it) },
            label = stringResource(R.string.comment),
            modifier = Modifier.fillMaxWidth().testTag("addDamageCommentInput"),
            errorMessage = if (errors.value.comment) stringResource(R.string.comment_error) else null,
        )
        Text(stringResource(R.string.priority), modifier = Modifier.padding(top = 10.dp), color = MaterialTheme.colorScheme.onPrimaryContainer)
        DropDown(
            items = listOf(
                DropDownItem(stringResource(R.string.low), DamagePriority.low),
                DropDownItem(stringResource(R.string.medium), DamagePriority.medium),
                DropDownItem(stringResource(R.string.high), DamagePriority.high),
                DropDownItem(stringResource(R.string.urgent), DamagePriority.urgent)
            ),
            selectedItem = form.value.priority,
            onItemSelected = { viewModel.setPriority(it) },
            testTag = "addDamagePriorityDropDown"
        )
        Text(stringResource(R.string.room), modifier = Modifier.padding(top = 10.dp), color = MaterialTheme.colorScheme.onPrimaryContainer)
        DropDown(
            items = viewModel.rooms.map { DropDownItem(it.name, it.id) },
            selectedItem = form.value.room_id,
            onItemSelected = { it?.let { viewModel.setRoomId(it) } },
            error = if (errors.value.room) stringResource(R.string.select_an_element) else null,
            testTag = "addDamageRoomDropDown"
        )
        StyledButton(
            text = stringResource(R.string.submit),
            onClick = {
                viewModel.submit(
                    addDamage = {
                        addDamage(it)
                        onClose()
                    },
                    tenantName = ""
                )
            },
            testTag = "addDamageSubmitButton"
        )
    }
}