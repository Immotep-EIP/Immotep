package com.example.immotep.realProperty.details.tabs

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowColumn
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import com.example.immotep.apiCallerServices.Damage
import com.example.immotep.apiCallerServices.DamageInput
import com.example.immotep.ui.components.StyledButton

@OptIn(ExperimentalLayoutApi::class)
@Composable
fun Damages(
    damageList : List<Damage>,
    addDamage : ((DamageInput) -> Unit)?
) {
    FlowColumn(
        modifier = Modifier
            .testTag("realPropertyDetailsDamagesTab")
            .fillMaxSize()
            .verticalScroll(rememberScrollState())
    ) {
        if (addDamage != null) {
            StyledButton(
                onClick = {},
                text = "Add Damage",
            )
        }
        damageList.forEach { item ->
            Text(item.id)
        }
    }
}