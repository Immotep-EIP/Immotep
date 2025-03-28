package com.example.immotep.realProperty.details

import android.content.Context
import android.widget.Toast
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
import androidx.compose.foundation.layout.defaultMinSize
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
import androidx.compose.runtime.State
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.core.content.FileProvider
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.LocalApiService
import com.example.immotep.R
import com.example.immotep.addOrEditPropertyModal.AddOrEditPropertyModal
import com.example.immotep.apiCallerServices.DetailedProperty
import com.example.immotep.apiCallerServices.Document
import com.example.immotep.apiCallerServices.PropertyStatus
import com.example.immotep.components.ErrorAlert
import com.example.immotep.components.InitialFadeIn
import com.example.immotep.components.InternalLoading
import com.example.immotep.components.LoadingDialog
import com.example.immotep.inviteTenantModal.InviteTenantModal
import com.example.immotep.realProperty.PropertyBox
import com.example.immotep.realProperty.PropertyBoxTextLine
import com.example.immotep.ui.components.BackButton
import com.example.immotep.utils.Base64Utils
import com.example.immotep.utils.DateFormatter
import com.example.immotep.utils.PdfsUtils

@Composable
fun OneDocument(document: Document, openPdf: (String) -> Unit) {
    InitialFadeIn(durationMs = 500) {
        Column(
            horizontalAlignment = Alignment.CenterHorizontally,
            modifier = Modifier
                .fillMaxWidth(0.33f)
                .padding(5.dp)
                .clickable {
                    openPdf(document.id)
                }
                .wrapContentSize(Alignment.Center)
                .testTag("OneDocument ${document.id}")
        ) {
            Box(
                modifier = Modifier
                    .background(MaterialTheme.colorScheme.background)
                    .border(
                        1.dp,
                        color = MaterialTheme.colorScheme.background,
                        shape = RoundedCornerShape(5.dp)
                    )
                    .padding(start = 25.dp, end = 25.dp, top = 10.dp, bottom = 10.dp)
            ) {
                Icon(
                    Icons.Outlined.AttachFile,
                    contentDescription = "document icon",
                    modifier = Modifier.size(50.dp)
                )
            }
            Text(
                text = document.name,
                textAlign = TextAlign.Center,
                modifier = Modifier.padding(start = 10.dp, end = 10.dp).fillMaxWidth(),
                maxLines = 2,
                overflow = TextOverflow.Ellipsis
            )
        }
    }
}

@OptIn(ExperimentalLayoutApi::class)
@Composable
fun DocumentBox(property : State<DetailedProperty>, openPdf: (String) -> Unit) {
    Text(text = stringResource(R.string.documents))
    Box(
        modifier = Modifier.fillMaxWidth()
            .border(
                1.dp,
                color = MaterialTheme.colorScheme.onBackground,
                shape = RoundedCornerShape(5.dp)
            )
            .background(MaterialTheme.colorScheme.tertiaryContainer)
            .padding(5.dp)

    ) {
        FlowRow(modifier = Modifier.defaultMinSize(minHeight = 125.dp)) {
            property.value.documents.forEach { item ->
                OneDocument(item, openPdf = { openPdf(it)})
            }
        }
    }
}

@Composable
fun AboutThePropertyBox(property : State<DetailedProperty>, openEdit : () -> Unit) {
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
        colors = ButtonDefaults.buttonColors(containerColor = MaterialTheme.colorScheme.tertiary),
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

@Composable
fun RealPropertyDetailsScreen(navController: NavController, newProperty : DetailedProperty, getBack: (DetailedProperty) -> Unit) {
    val apiService = LocalApiService.current
    val context = LocalContext.current

    val viewModel: RealPropertyDetailsViewModel = viewModel {
        RealPropertyDetailsViewModel(navController, apiService)
    }
    val property = viewModel.property.collectAsState()
    var editOpen by rememberSaveable { mutableStateOf(false) }
    var inviteTenantOpen by rememberSaveable { mutableStateOf(false) }

    val apiErrors = viewModel.apiError.collectAsState()
    val isLoading = viewModel.isLoading.collectAsState()
    val errorAlertVal = when (apiErrors.value) {
        RealPropertyDetailsViewModel.ApiErrors.GET_PROPERTY -> stringResource(R.string.api_error_get_property)
        RealPropertyDetailsViewModel.ApiErrors.UPDATE_PROPERTY -> stringResource(R.string.api_error_edit_property)
        else -> null
    }
    LaunchedEffect(newProperty) {
        viewModel.loadProperty(newProperty)
    }
    AddOrEditPropertyModal(
        open = editOpen,
        close = { editOpen = false },
        onSubmit = { viewModel.editProperty(it, newProperty.id) },
        popupName = stringResource(R.string.edit_property),
        baseValue = property.value.toAddPropertyInput(),
        submitButtonText = stringResource(R.string.save),
        submitButtonIcon = { Icon(Icons.Outlined.EditNote, contentDescription = "Edit property") }
    )
    InviteTenantModal(
        open = inviteTenantOpen,
        close = { inviteTenantOpen = false },
        navController = navController,
        propertyId = newProperty.id,
        onSubmit = {email, startDate, endDate -> viewModel.onSubmitInviteTenant(email, startDate, endDate) }
    )
    if (isLoading.value) {
        InternalLoading()
        return
    }
    InitialFadeIn(300) {
        Column(modifier = Modifier.padding(5.dp).testTag("realPropertyDetailsScreen")) {
            Row(
                verticalAlignment = Alignment.CenterVertically,
                modifier = Modifier.fillMaxWidth(),
            ) {
                BackButton {
                    getBack(property.value)
                }
                Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.End) {
                    Button(onClick = { navController.navigate("messages") }) {
                        Text(
                            stringResource(R.string.open_in_messages),
                            modifier = Modifier.padding(end = 5.dp)
                        )
                        Icon(
                            Icons.Outlined.MailOutline,
                            contentDescription = stringResource(R.string.open_in_messages)
                        )
                    }
                }
            }
            ErrorAlert(null, null, errorAlertVal)
            PropertyBox(property.value)
            if (property.value.status == PropertyStatus.available) {
                Button(
                    onClick = { inviteTenantOpen = true },
                    colors = ButtonDefaults.buttonColors(containerColor = MaterialTheme.colorScheme.tertiary),
                    modifier = Modifier
                        .clip(RoundedCornerShape(5.dp))
                        .padding(5.dp)
                        .fillMaxWidth()
                        .testTag("inviteTenantBtn")
                ) {
                    Text(
                        stringResource(R.string.invite_tenant),
                        color = MaterialTheme.colorScheme.onTertiary
                    )
                }
            }
            AboutThePropertyBox(property, openEdit = { editOpen = true })
            DocumentBox(
                property = property,
                openPdf = { viewModel.openPdf(it, context) }
            )
            Button(
                onClick = { navController.navigate("inventory/${newProperty.id}") },
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
}
