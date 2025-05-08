package com.example.immotep.realProperty.details.tabs

import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowColumn
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.navigation.NavController
import com.example.immotep.addDamageModal.AddDamageModal
import com.example.immotep.apiCallerServices.Damage
import com.example.immotep.apiCallerServices.DamageInput
import com.example.immotep.layouts.BigModalLayout
import com.example.immotep.ui.components.StyledButton




@OptIn(ExperimentalLayoutApi::class)
@Composable
fun Damages(
    damageList : List<Damage>,
    addDamage : ((Damage) -> Unit)?,
    navController: NavController
) {
    var addDamageOpen by rememberSaveable { mutableStateOf(false) }
    AddDamageModal(
        open = addDamageOpen,
        onClose = { addDamageOpen = false },
        addDamage = { addDamage?.invoke(it) },
        navController = navController
    )
    FlowColumn(
        modifier = Modifier
            .testTag("realPropertyDetailsDamagesTab")
            .fillMaxSize()
            .verticalScroll(rememberScrollState())
    ) {
        if (addDamage != null) {
            StyledButton(
                onClick = { addDamageOpen = true },
                text = "Add Damage",
            )
        }
        damageList.forEach { item ->
            Text(item.id)
        }
    }
}