package com.example.immotep.realProperty.details.tabs

import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AccountBox
import androidx.compose.material.icons.outlined.AllOut
import androidx.compose.material.icons.outlined.CalendarMonth
import androidx.compose.material.icons.outlined.CalendarViewMonth
import androidx.compose.material.icons.outlined.EditNote
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.State
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.example.immotep.R
import com.example.immotep.apiCallerServices.DetailedProperty
import com.example.immotep.realProperty.PropertyBoxTextLine
import com.example.immotep.utils.DateFormatter

@Composable
fun AboutPropertyTab(property : State<DetailedProperty>, openEdit : () -> Unit) {
    Text(text = stringResource(R.string.about_the_property))
    Row(
        modifier = Modifier
            .fillMaxWidth()
            .border(1.dp, color = MaterialTheme.colorScheme.onBackground, shape = RoundedCornerShape(5.dp))
            .padding(5.dp)
    ) {
        Column(modifier = Modifier.fillMaxWidth(0.5f)) {
            PropertyBoxTextLine(property.value.tenant?: "", Icons.Outlined.AccountBox)
            PropertyBoxTextLine(
                DateFormatter.formatOffsetDateTime(property.value.startDate) ?:
                "---------------------",
                Icons.Outlined.CalendarMonth
            )
            PropertyBoxTextLine(
                (
                        DateFormatter.formatOffsetDateTime(property.value.endDate) ?:
                        "---------------------"
                        ),
                Icons.Outlined.CalendarMonth
            )
        }
        Column(modifier = Modifier.fillMaxWidth()) {
            PropertyBoxTextLine("${stringResource(R.string.area)}: ${property.value.area} m²", Icons.Outlined.AllOut)
            PropertyBoxTextLine(
                "${stringResource(R.string.rentMonth)}: ${property.value.rent}€",
                Icons.Outlined.CalendarViewMonth,
            )
            PropertyBoxTextLine(
                "${stringResource(R.string.deposit)}: ${property.value.deposit}€",
                Icons.Outlined.EditNote,
            )
        }
    }
    Spacer(modifier = Modifier.height(10.dp))
    Button(
        onClick = { openEdit() },
        colors = ButtonDefaults.buttonColors(containerColor = MaterialTheme.colorScheme.secondary),
        modifier = Modifier
            .clip(RoundedCornerShape(5.dp))
            .padding(5.dp)
            .fillMaxWidth()
            .testTag("editProperty")
    ) {
        Text(
            stringResource(R.string.edit_property),
            color = MaterialTheme.colorScheme.onTertiary
        )
    }
    Spacer(modifier = Modifier.height(10.dp))
}