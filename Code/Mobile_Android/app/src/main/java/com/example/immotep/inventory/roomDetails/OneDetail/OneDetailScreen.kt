package com.example.immotep.inventory.roomDetails.OneDetail

import androidx.compose.foundation.layout.Column
import androidx.compose.material.Button
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.immotep.components.AddingPicturesCarousel
import com.example.immotep.components.InitialFadeIn
import com.example.immotep.inventory.RoomDetail
import com.example.immotep.layouts.InventoryLayout

@Composable
fun OneDetailScreen(onModifyDetail : (detailIndex : Int, detail : RoomDetail) -> Unit, index : Int, baseDetail : RoomDetail) {
    val viewModel : OneDetailViewModel = viewModel()
    LaunchedEffect(Unit) {
        viewModel.reset(baseDetail)
    }
    InventoryLayout(
        testTag = "oneDetailScreen",
        { viewModel.onClose(onModifyDetail, index) }
    ) {
        InitialFadeIn {
            Column {
                Text("GROSINJE")
                AddingPicturesCarousel(pictures = viewModel.picture, addPicture = { uri -> viewModel.addPicture(uri) })
            }
        }
    }
}