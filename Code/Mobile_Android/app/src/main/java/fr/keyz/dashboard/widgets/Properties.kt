package fr.keyz.dashboard.widgets

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import fr.keyz.R
import fr.keyz.apiCallerServices.DashBoardProperties

@Composable
fun PropertiesWidget(properties: DashBoardProperties) {
    WidgetBase(title = stringResource(R.string.properties_cap), dropDownItems = arrayOf(), testTag = "availablePropertiesWidget") {
        Row(
            modifier = Modifier.fillMaxWidth(),
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.SpaceBetween
        ) {
            WidgetNumberBase(
                title = stringResource(R.string.busy),
                value = properties.nbrOccupied,
                titleColor = MaterialTheme.colorScheme.error
            )
            WidgetNumberBase(
                title = stringResource(R.string.pending),
                value = properties.nbrPendingInvites,
                titleColor = MaterialTheme.colorScheme.inversePrimary
            )
            WidgetNumberBase(
                title = stringResource(R.string.available),
                value = properties.nbrAvailable,
                titleColor = MaterialTheme.colorScheme.surfaceVariant
            )
        }
    }
}