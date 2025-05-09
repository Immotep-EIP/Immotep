package com.example.keyz.addDamageModal

import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.keyz.LocalApiService
import com.example.keyz.apiCallerServices.DamagePriority
import com.example.keyz.components.AddingPicturesCarousel
import com.example.keyz.layouts.BigModalLayout
import com.example.keyz.ui.components.OutlinedTextField
import com.example.keyz.ui.components.DropDown
import com.example.keyz.ui.components.DropDownItem
import com.example.keyz.ui.components.StyledButton
import com.example.keyz.R
import com.example.keyz.apiCallerServices.Damage

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
        close = onClose
    ) {
        AddingPicturesCarousel(
            uriPictures = viewModel.pictures,
            addPicture = { viewModel.addPicture(it) },
            removePicture = { viewModel.removePicture(it) },
            error = if (errors.value.pictures) stringResource(R.string.add_picture_error) else null
        )
        OutlinedTextField(
            value = form.value.comment,
            onValueChange = { viewModel.setComment(it) },
            label = stringResource(R.string.comment),
            modifier = Modifier.fillMaxWidth(),
            errorMessage = if (errors.value.comment) stringResource(R.string.comment_error) else null
        )
        Text(stringResource(R.string.priority), modifier = Modifier.padding(top = 10.dp))
        DropDown(
            items = listOf(
                DropDownItem(stringResource(R.string.low), DamagePriority.low),
                DropDownItem(stringResource(R.string.medium), DamagePriority.medium),
                DropDownItem(stringResource(R.string.high), DamagePriority.high),
                DropDownItem(stringResource(R.string.urgent), DamagePriority.urgent)
            ),
            selectedItem = form.value.priority,
            onItemSelected = { viewModel.setPriority(it) }
        )
        Text(stringResource(R.string.room), modifier = Modifier.padding(top = 10.dp))
        DropDown(
            items = viewModel.rooms.map { DropDownItem(it.name, it.id) },
            selectedItem = form.value.room_id,
            onItemSelected = { it?.let { viewModel.setRoomId(it) } },
            error = if (errors.value.room) stringResource(R.string.select_an_element) else null
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
            }
        )
    }
}