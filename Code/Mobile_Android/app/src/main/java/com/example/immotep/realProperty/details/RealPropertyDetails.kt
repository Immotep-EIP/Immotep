package com.example.immotep.realProperty.details

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.realProperty.PropertyBox
import com.example.immotep.ui.components.BackButton

@Composable
fun RealPropertyDetailsScreen(navController: NavController, propertyId: String, getBack : () -> Unit) {
    val viewModel: RealPropertyDetailsViewModel = viewModel(factory = RealPropertyDetailsViewModelFactory(propertyId))
    val property = viewModel.property.collectAsState()
    Column {
        BackButton(getBack)
        PropertyBox(property.value.toProperty())
    }
}
