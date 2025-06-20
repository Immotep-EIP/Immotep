package fr.keyz.dashboard.widgets

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.height
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import fr.keyz.R

@Composable
fun DamageInProgressWidget() {
    WidgetBase(title = stringResource(R.string.damage_in_progress), dropDownItems = arrayOf(), testTag = "damageInProgressWidget") {
        Box(modifier = Modifier.height(100.dp), contentAlignment = Alignment.Center) {

        }
    }
}