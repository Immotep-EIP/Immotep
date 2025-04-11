package com.example.immotep.realProperty

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.gestures.detectTapGestures
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.offset
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.layout.wrapContentHeight
import androidx.compose.foundation.layout.wrapContentSize
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AccountCircle
import androidx.compose.material.icons.outlined.Add
import androidx.compose.material.icons.outlined.DateRange
import androidx.compose.material.icons.outlined.Place
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.IconButtonDefaults
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.RectangleShape
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.input.pointer.pointerInput
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
import com.example.immotep.apiCallerServices.PropertyStatus
import com.example.immotep.components.DeletePopUp
import com.example.immotep.components.ErrorAlert
import com.example.immotep.components.InitialFadeIn
import com.example.immotep.dashboard.DashBoardLayout
import com.example.immotep.realProperty.details.RealPropertyDetailsScreen
import com.example.immotep.utils.DateFormatter

@Composable
fun PropertyBoxTextLine(text: String, icon: ImageVector) {
    Row(
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier.padding(top = 5.dp)
    ) {
        Icon(icon, contentDescription = "icon", tint = MaterialTheme.colorScheme.secondary)
        Text(text = text, modifier = Modifier.padding(start = 5.dp), color = MaterialTheme.colorScheme.secondary)
    }
}

@Composable
fun PropertyBox(property: DetailedProperty, onClick: (() -> Unit)? = null, onDelete: (() -> Unit)? = null) {
    val modifierColumn = if (onClick != null && onDelete != null) {
        Modifier.pointerInput(Unit) {
            detectTapGestures(
                onLongPress = {
                    onDelete()
                },
                onPress = {
                    val isReleased = tryAwaitRelease()
                    if (isReleased) {
                        onClick()
                    }
                }
            )
        }
    }
    else if (onClick != null) {
       Modifier.clickable {  }
    } else {
        Modifier
    }
    Box(modifier = Modifier.testTag("propertyBoxRow")
            .padding(10.dp)
        .fillMaxWidth()
        .wrapContentHeight()

    ) {
        Column(
            verticalArrangement = Arrangement.Center,
            modifier = modifierColumn
                .fillMaxWidth()
                .fillMaxHeight()
                .shadow(
                    elevation = 6.dp,
                    shape = RoundedCornerShape(10.dp),
                    clip = false
                )
                .clip(RoundedCornerShape(10.dp))
                .background(color = MaterialTheme.colorScheme.primaryContainer, shape = RoundedCornerShape(10.dp))
                .padding(15.dp)
                .testTag("propertyBox ${property.id}")
        ) {
            AsyncImage(
                model = property.image,
                placeholder = painterResource(id = R.drawable.immotep_png_logo),
                error = painterResource(id = R.drawable.immotep_png_logo),
                contentDescription = "picture of the ${property.tenant} property",
                modifier = Modifier
                    .fillMaxWidth()
                    .height(150.dp)
                    .clip(
                        RoundedCornerShape(50.dp)
                    )
            )
            Spacer(modifier = Modifier.height(10.dp))
            Text(property.name, fontSize = 20.sp, fontWeight = FontWeight.Bold)
            Spacer(modifier = Modifier.height(8.dp))
            PropertyBoxTextLine(property.address, Icons.Outlined.Place)
            /*
            PropertyBoxTextLine(property.tenant?: "", Icons.Outlined.AccountCircle)
            PropertyBoxTextLine(
                DateFormatter.formatOffsetDateTime(property.startDate)?: "---------------------",
                Icons.Outlined.DateRange
            )

             */
        }

        Box(
            modifier = Modifier
                .align(Alignment.TopEnd)
                .padding(top = 5.dp, end = 5.dp)
                .width(100.dp)
                .height(30.dp)
                .clip(RoundedCornerShape(10.dp))
                .background(
                    when(property.status) {
                        PropertyStatus.available -> MaterialTheme.colorScheme.surfaceVariant
                        PropertyStatus.invite_sent -> MaterialTheme.colorScheme.inversePrimary
                        else -> MaterialTheme.colorScheme.error
                    }
                )
                .padding(3.dp)
                .testTag("topRightPropertyBoxInfo")
        ) {
            Text(
                color = MaterialTheme.colorScheme.onError,
                text = when(property.status) {
                    PropertyStatus.available -> stringResource(R.string.available)
                    PropertyStatus.invite_sent -> stringResource(R.string.pending)
                    else -> stringResource(R.string.busy)
                },
                modifier = Modifier
                    .fillMaxWidth()
                    .wrapContentSize(Alignment.Center)
                    .padding(1.dp)
            )
        }

    }
}

@Composable
fun RealPropertyScreen(navController: NavController) {
    val apiService = LocalApiService.current
    val viewModel: RealPropertyViewModel =
        viewModel {
            RealPropertyViewModel(
                navController,
                apiService
            )
        }
    var deleteOpen by rememberSaveable { mutableStateOf<Pair<String, String>?>(null) }
    var addPropertyModalOpen by rememberSaveable { mutableStateOf(false) }
    val apiErrors = viewModel.apiError.collectAsState()
    val propertySelectedDetails = viewModel.propertySelectedDetails.collectAsState()

    val errorAlertVal = when (apiErrors.value) {
        RealPropertyViewModel.WhichApiError.GET_PROPERTIES -> stringResource(R.string.api_error_get_properties)
        RealPropertyViewModel.WhichApiError.ADD_PROPERTY -> stringResource(R.string.api_error_add_property)
        RealPropertyViewModel.WhichApiError.DELETE_PROPERTY -> stringResource(R.string.api_error_delete_property)
        else -> null
    }
    LaunchedEffect(Unit) {
        viewModel.getProperties()
    }

    DashBoardLayout(navController, "realPropertyScreen") {
        if (propertySelectedDetails.value == null) {
            ErrorAlert(null, null, errorAlertVal)
            DeletePopUp(
                open = deleteOpen != null,
                delete = { viewModel.deleteProperty(deleteOpen!!.first) },
                close = { deleteOpen = null },
                globalName = stringResource(R.string.property),
                detailedName = deleteOpen?.second?: ""
            )
            Box(modifier = Modifier.fillMaxSize()) {

                LazyColumn {
                    items(viewModel.properties) { item ->
                        InitialFadeIn {
                            PropertyBox(
                                item,
                                onClick = {
                                    if (deleteOpen == null) {
                                        viewModel.setPropertySelectedDetails(item.id)
                                        viewModel.closeError()
                                    }
                                },
                                onDelete = { deleteOpen = Pair(item.id, item.address) })
                        }
                    }
                }
                Box(
                    modifier = Modifier.align(Alignment.BottomEnd)
                ) {
                    IconButton(
                        onClick = { addPropertyModalOpen = true },
                        colors = IconButtonDefaults.iconButtonColors(containerColor = MaterialTheme.colorScheme.secondary),
                        modifier = Modifier
                            .height(55.dp)
                            .width(55.dp)
                            .clip(CircleShape)
                            .testTag("addAPropertyBtn")
                    ) {
                        Icon(
                            Icons.Outlined.Add,
                            contentDescription = "add",
                            tint = MaterialTheme.colorScheme.onSecondary,
                            modifier = Modifier.size(45.dp)
                        )
                    }
                }
            }

        } else {
            RealPropertyDetailsScreen(
                navController,
                propertySelectedDetails.value!!,
                getBack = { viewModel.getBackFromDetails(it) }
            )
        }
        AddOrEditPropertyModal(
            open = addPropertyModalOpen,
            close = { addPropertyModalOpen = false },
            onSubmit = { property -> viewModel.addProperty(property) },
            popupName = stringResource(R.string.create_new_property),
            submitButtonText = stringResource(R.string.add_prop),
            submitButtonIcon = { Icon(Icons.Outlined.Add, contentDescription = "add") }
        )
    }
}
