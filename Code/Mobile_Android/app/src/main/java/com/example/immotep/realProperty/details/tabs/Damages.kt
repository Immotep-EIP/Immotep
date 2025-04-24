package com.example.immotep.realProperty.details.tabs

import androidx.compose.foundation.layout.Column
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag

@Composable
fun Damages() {
    Column(
        modifier = Modifier.testTag("realPropertyDetailsDamagesTab")
    ) {
        Text("Damages")
    }
}