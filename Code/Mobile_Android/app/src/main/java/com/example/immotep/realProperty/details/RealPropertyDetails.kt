package com.example.immotep.realProperty.details

import androidx.compose.foundation.background
import androidx.compose.foundation.border
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
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AccountBox
import androidx.compose.material.icons.outlined.AllOut
import androidx.compose.material.icons.outlined.CalendarMonth
import androidx.compose.material.icons.outlined.CalendarViewMonth
import androidx.compose.material.icons.outlined.EditNote
import androidx.compose.material.icons.outlined.MoreVert
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.IconButtonDefaults
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.State
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import coil.compose.AsyncImage
import com.example.immotep.LocalApiService
import com.example.immotep.R
import com.example.immotep.addOrEditPropertyModal.AddOrEditPropertyModal
import com.example.immotep.apiCallerServices.DetailedProperty
import com.example.immotep.components.ErrorAlert
import com.example.immotep.components.InitialFadeIn
import com.example.immotep.components.InternalLoading
import com.example.immotep.inviteTenantModal.InviteTenantModal
import com.example.immotep.layouts.TabsLayout
import com.example.immotep.realProperty.PropertyBoxTextLine
import com.example.immotep.realProperty.PropertyStatusBox
import com.example.immotep.realProperty.details.tabs.AboutPropertyTab
import com.example.immotep.realProperty.details.tabs.Damages
import com.example.immotep.realProperty.details.tabs.DocumentBox
import com.example.immotep.realProperty.details.tabs.OneDocument
import com.example.immotep.ui.components.BackButton
import com.example.immotep.utils.DateFormatter


@Composable
fun RealPropertyDetailsScreen(navController: NavController, newProperty : DetailedProperty, getBack: (DetailedProperty) -> Unit) {
    val apiService = LocalApiService.current
    val context = LocalContext.current
    val tabs = listOf(
        stringResource(R.string.about),
        stringResource(R.string.documents),
        stringResource(R.string.damages)
    )


    val viewModel: RealPropertyDetailsViewModel = viewModel {
        RealPropertyDetailsViewModel(navController, apiService)
    }
    val property = viewModel.property.collectAsState()
    var editOpen by rememberSaveable { mutableStateOf(false) }
    var tabIndex by rememberSaveable { mutableIntStateOf(0) }
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
        Column(modifier = Modifier.testTag("realPropertyDetailsScreen")) {
            Box(modifier = Modifier.fillMaxWidth()) {
                AsyncImage(
                    model = property.value.image,
                    placeholder = painterResource(id = R.drawable.immotep_png_logo),
                    error = painterResource(id = R.drawable.immotep_png_logo),
                    contentDescription = "picture of the ${property.value.tenant} property",
                    modifier = Modifier
                        .fillMaxWidth()
                        .height(200.dp)
                        .clip(
                            RoundedCornerShape(50.dp)
                        )
                )
                Box(
                    modifier = Modifier
                        .align(Alignment.TopStart)
                        .padding(top = 5.dp, start = 5.dp)
                ) {
                    Row(
                        modifier = Modifier.fillMaxWidth(),
                        horizontalArrangement = Arrangement.SpaceBetween,
                        verticalAlignment = Alignment.CenterVertically
                    ) {
                        BackButton { getBack(property.value) }
                        IconButton(
                            onClick = {},
                            colors = IconButtonDefaults.iconButtonColors(containerColor = MaterialTheme.colorScheme.background),
                            modifier = Modifier.testTag("backButton"),
                        ) {
                            Icon(Icons.Outlined.MoreVert, contentDescription = "More vert", tint = MaterialTheme.colorScheme.onBackground)
                        }
                    }
                }
            }
            Spacer(modifier = Modifier.height(10.dp))
            ErrorAlert(null, null, errorAlertVal)
            Column(modifier = Modifier.background(MaterialTheme.colorScheme.background).padding(20.dp)) {
                Box(modifier = Modifier.fillMaxWidth()) {
                    PropertyStatusBox(property.value.status, modifier = Modifier.padding(end = 5.dp).align(Alignment.TopEnd))
                }
                Text(property.value.name, fontSize = 15.sp, fontWeight = FontWeight.Bold)
                Text("${property.value.address}, ${property.value.zipCode} ${property.value.city}, ${property.value.country}",
                    fontSize = 15.sp,
                    fontWeight = FontWeight.Thin
                )
            }
            TabsLayout(tabIndex, tabs, { tabIndex = it }) {
                when (tabIndex) {
                    0 -> AboutPropertyTab(property, openEdit = { editOpen = true })
                    1 -> DocumentBox(property = property, openPdf = { viewModel.openPdf(it, context)})
                    2 -> Damages()
                }
            }


            /*
            if (property.value.status == PropertyStatus.available) {
                Button(
                    onClick = { inviteTenantOpen = true },
                    colors = ButtonDefaults.buttonColors(containerColor = MaterialTheme.colorScheme.secondary),
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
            if (property.value.status == PropertyStatus.unavailable) {
                Button(
                    onClick = { navController.navigate("inventory/${newProperty.id}") },
                    colors = ButtonDefaults.buttonColors(containerColor = MaterialTheme.colorScheme.secondary),
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
             */
        }
    }
}
