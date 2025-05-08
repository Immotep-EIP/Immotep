package com.example.immotep.addDamageModal

import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material.Text
import androidx.compose.material3.OutlinedTextField
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.LocalApiService
import com.example.immotep.apiCallerServices.DamageInput
import com.example.immotep.apiCallerServices.DamagePriority
import com.example.immotep.components.AddingPicturesCarousel
import com.example.immotep.layouts.BigModalLayout
import com.example.immotep.ui.components.OutlinedTextField
import com.example.immotep.ui.components.DropDown
import com.example.immotep.ui.components.DropDownItem
import com.example.immotep.ui.components.StyledButton
import com.example.immotep.R

@Composable
fun AddDamageModal(
    open : Boolean,
    onClose : () -> Unit,
    addDamage : (DamageInput) -> Unit,
    navController: NavController
) {
    val apiService = LocalApiService.current
    val viewModel = viewModel {
        AddDamageModalViewModel(apiService, navController)
    }
    val form = viewModel.form.collectAsState()
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
            removePicture = { viewModel.removePicture(it) }
        )
        OutlinedTextField(
            value = form.value.comment,
            onValueChange = { viewModel.setComment(it) },
            label = stringResource(R.string.comment),
            modifier = Modifier.fillMaxWidth()
        )
        Text(stringResource(R.string.priority), modifier = Modifier.padding(top = 10.dp))
        DropDown(
            items = listOf(
                DropDownItem(stringResource(R.string.low), DamagePriority.LOW),
                DropDownItem(stringResource(R.string.medium), DamagePriority.MEDIUM),
                DropDownItem(stringResource(R.string.high), DamagePriority.HIGH),
                DropDownItem(stringResource(R.string.urgent), DamagePriority.URGENT)
            ),
            selectedItem = form.value.priority,
            onItemSelected = { viewModel.setPriority(it) }
        )
        Text(stringResource(R.string.room), modifier = Modifier.padding(top = 10.dp))
        DropDown(
            items = viewModel.rooms.map { DropDownItem(it.name, it.id) },
            selectedItem = form.value.room_id,
            onItemSelected = { it?.let { viewModel.setRoomId(it) } }
        )
        StyledButton(
            text = stringResource(R.string.submit),
            onClick = {
                viewModel.submit()
            }
        )
    }
}