package com.example.immotep.realProperty.details.tabs

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AccountBalanceWallet
import androidx.compose.material.icons.outlined.AccountBox
import androidx.compose.material.icons.outlined.AllOut
import androidx.compose.material.icons.outlined.CalendarMonth
import androidx.compose.material.icons.outlined.CalendarViewMonth
import androidx.compose.material.icons.outlined.EditNote
import androidx.compose.material.icons.outlined.RequestQuote
import androidx.compose.material.icons.outlined.Straighten
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.State
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import com.example.immotep.R
import com.example.immotep.apiCallerServices.DetailedProperty
import com.example.immotep.realProperty.PropertyBoxTextLine
import com.example.immotep.utils.DateFormatter

@Composable
fun AboutThePropertyBox(name : String, value : String, icon : ImageVector, modifier: Modifier) {
    Column(
        horizontalAlignment = Alignment.CenterHorizontally,
        modifier = modifier
            .shadow(
                elevation = 6.dp,
                shape = RoundedCornerShape(10.dp),
                clip = false
            )
            .clip(RoundedCornerShape(10.dp))
            .background(color = MaterialTheme.colorScheme.primaryContainer, shape = RoundedCornerShape(10.dp))
            .padding(3.dp)
    ) {
        Icon(icon, contentDescription = "$name box icon", tint = MaterialTheme.colorScheme.secondary)
        Text(value, fontWeight = FontWeight.Bold, fontSize = 15.sp)
        Text(name, fontWeight = FontWeight.Thin, fontSize = 12.sp)
    }
}


@Composable
fun AboutPropertyTab(property : State<DetailedProperty>) {
    Row(
        modifier = Modifier.fillMaxWidth().padding(16.dp),
        horizontalArrangement = Arrangement.spacedBy(18.dp)
    ) {
        AboutThePropertyBox(
            stringResource(R.string.area),
            "${property.value.area} m²",
            Icons.Outlined.Straighten,
            modifier = Modifier.weight(1f)
        )
        AboutThePropertyBox(
            stringResource(R.string.rentPerMonth),
            "${property.value.rent} €",
            Icons.Outlined.RequestQuote,
            modifier = Modifier.weight(1f)
        )
        AboutThePropertyBox(
            stringResource(R.string.deposit),
            "${property.value.deposit} €",
            Icons.Outlined.AccountBalanceWallet,
            modifier = Modifier.weight(1f)
        )
    }
    Row(
        modifier = Modifier.fillMaxWidth().padding(16.dp)
    ) {
        Column {
            Text("${stringResource(R.string.tenant)}:", fontWeight = FontWeight.Bold, fontSize = 15.sp)
            Text(if (property.value.tenant != null) property.value.tenant!! else "---------------------", fontSize = 15.sp)
        }
        Spacer(Modifier.width(15.dp))
        Column {
            Text("${stringResource(R.string.dates)}:", fontWeight = FontWeight.Bold, fontSize = 15.sp)
            Text(
                if (property.value.startDate != null && property.value.endDate != null) {
                    "${DateFormatter.formatOffsetDateTime(property.value.startDate)} - ${DateFormatter.formatOffsetDateTime(property.value.endDate)}"
                }
                else {
                    "---------------------"
                },
                fontSize = 15.sp)
        }
    }
}
