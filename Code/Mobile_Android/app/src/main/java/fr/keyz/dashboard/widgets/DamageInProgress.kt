package fr.keyz.dashboard.widgets

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import fr.keyz.R
import fr.keyz.apiCallerServices.DashBoardOpenDamage
import fr.keyz.apiCallerServices.Priority
import fr.keyz.utils.ThemeUtils

@Composable
fun DamageWidget(damages : DashBoardOpenDamage) {
    WidgetBase(title = stringResource(R.string.damages), dropDownItems = arrayOf(), testTag = "damageInProgressWidget") {
        Row(
            modifier = Modifier.fillMaxWidth(),
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.SpaceBetween
        ) {
            WidgetNumberBase(
                title = stringResource(R.string.urgent),
                value = damages.nbrUrgent,
                titleColor = ThemeUtils.getStatusColor(priority = Priority.urgent)
            )
            WidgetNumberBase(
                title = stringResource(R.string.high),
                value = damages.nbrHigh,
                titleColor = ThemeUtils.getStatusColor(priority = Priority.high)
            )
            WidgetNumberBase(
                title = stringResource(R.string.medium),
                value = damages.nbrMedium,
                titleColor = ThemeUtils.getStatusColor(priority = Priority.medium)
            )
            WidgetNumberBase(
                title = stringResource(R.string.low),
                value = damages.nbrLow,
                titleColor = ThemeUtils.getStatusColor(priority = Priority.low)
            )
        }
    }
}