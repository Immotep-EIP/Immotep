package com.example.immotep.realProperty.details

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.wrapContentSize
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AccountBox
import androidx.compose.material.icons.outlined.AllOut
import androidx.compose.material.icons.outlined.AttachFile
import androidx.compose.material.icons.outlined.CalendarMonth
import androidx.compose.material.icons.outlined.CalendarViewMonth
import androidx.compose.material.icons.outlined.EditNote
import androidx.compose.material.icons.outlined.MailOutline
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.R
import com.example.immotep.realProperty.PropertyBox
import com.example.immotep.realProperty.PropertyBoxTextLine
import com.example.immotep.ui.components.BackButton
import java.text.SimpleDateFormat

@Composable
fun OneDocument(name: String) {
    Column(
        horizontalAlignment = Alignment.CenterHorizontally,
        modifier = Modifier
            .fillMaxWidth(0.33f)
            .padding(5.dp)
            .clickable { }
            .wrapContentSize(Alignment.Center)
            .testTag("OneDocument")
    ) {
        Box(
            modifier = Modifier
                .background(MaterialTheme.colorScheme.background)
                .border(1.dp, color = MaterialTheme.colorScheme.background, shape = RoundedCornerShape(5.dp))
                .padding(start = 25.dp, end = 25.dp, top = 10.dp, bottom = 10.dp)
        ) {
            Icon(Icons.Outlined.AttachFile, contentDescription = "document icon", modifier = Modifier.size(50.dp))
        }
        Text(text = name, textAlign = TextAlign.Center, modifier = Modifier.padding(start = 10.dp, end = 10.dp).fillMaxWidth())
    }
}

@OptIn(ExperimentalLayoutApi::class)
@Composable
fun RealPropertyDetailsScreen(navController: NavController, propertyId: String, getBack: () -> Unit) {
    val viewModel: RealPropertyDetailsViewModel = viewModel(factory = RealPropertyDetailsViewModelFactory(propertyId, navController))
    val property = viewModel.property.collectAsState()

    LaunchedEffect(Unit) {
        viewModel.loadProperty()
    }
    Column(modifier = Modifier.padding(5.dp).testTag("realPropertyDetailsScreen")) {
        Row(
            verticalAlignment = Alignment.CenterVertically,
            modifier = Modifier.fillMaxWidth(),
        ) {
            BackButton(getBack)
            Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.End) {
                Button(onClick = { navController.navigate("messages") }) {
                    Text(stringResource(R.string.open_in_messages), modifier = Modifier.padding(end = 5.dp))
                    Icon(Icons.Outlined.MailOutline, contentDescription = stringResource(R.string.open_in_messages))
                }
            }
        }
        PropertyBox(property.value.toProperty())
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
                    if (property.value.startDate != null)
                        SimpleDateFormat("dd/MM/yyyy").format(property.value.startDate)
                    else
                        "---------------------",
                    Icons.Outlined.CalendarMonth
                )
                PropertyBoxTextLine(
                    (
                            if (property.value.endDate != null)
                                SimpleDateFormat("dd/MM/yyyy").format(property.value.endDate)
                            else
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
        Text(text = stringResource(R.string.documents))
        Box(
            modifier = Modifier.fillMaxWidth()
                .border(1.dp, color = MaterialTheme.colorScheme.onBackground, shape = RoundedCornerShape(5.dp))
                .background(MaterialTheme.colorScheme.tertiaryContainer)
                .padding(5.dp)
        ) {
            FlowRow {
                property.value.documents.forEach {
                    item ->
                    OneDocument(item)
                }
            }
        }
        Button(
            onClick = { navController.navigate("inventory/$propertyId") },
            colors = ButtonDefaults.buttonColors(containerColor = MaterialTheme.colorScheme.tertiary),
            modifier = Modifier
                .clip(RoundedCornerShape(5.dp))
                .padding(5.dp)
                .fillMaxWidth()
                .testTag("startInventory")
        ) {
            Text(
                stringResource(R.string.start_inventory),
                color = MaterialTheme.colorScheme.onTertiary
            )
        }
    }
}
