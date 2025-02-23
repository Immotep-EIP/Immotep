package com.example.immotep.realProperty

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.gestures.detectTapGestures
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.layout.wrapContentSize
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AccountCircle
import androidx.compose.material.icons.outlined.DateRange
import androidx.compose.material.icons.outlined.Place
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Icon
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
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.input.pointer.pointerInput
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import coil.compose.AsyncImage
import com.example.immotep.R
import com.example.immotep.addPropertyModal.AddPropertyModal
import com.example.immotep.components.DeletePopUp
import com.example.immotep.components.InitialFadeIn
import com.example.immotep.dashboard.DashBoardLayout
import com.example.immotep.realProperty.details.RealPropertyDetailsScreen
import com.example.immotep.utils.DateFormatter
import java.text.SimpleDateFormat
import java.time.format.DateTimeFormatter

@Composable
fun PropertyBoxTextLine(text: String, icon: ImageVector) {
    Row(
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier.padding(top = 5.dp)
    ) {
        Icon(icon, contentDescription = "icon")
        Text(text = text, modifier = Modifier.padding(start = 5.dp))
    }
}

@Composable
fun PropertyBox(property: Property, onClick: (() -> Unit)? = null, onDelete: (() -> Unit)? = null) {
    val modifierRow = if (onClick != null && onDelete != null) {
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
    Box(modifier = Modifier.testTag("propertyBoxRow")) {
        Row(
            verticalAlignment = Alignment.CenterVertically,
            modifier = modifierRow
                .padding(10.dp)
                .border(
                    1.dp,
                    color = MaterialTheme.colorScheme.primary,
                    shape = RoundedCornerShape(5.dp)
                )
                .fillMaxWidth()
                .padding(start = 10.dp, end = 10.dp, top = 30.dp, bottom = 30.dp)
                .testTag("propertyBox ${property.id}")
        ) {
            AsyncImage(
                model = property.image,
                placeholder = painterResource(id = R.drawable.immotep_png_logo),
                error = painterResource(id = R.drawable.immotep_png_logo),
                contentDescription = "picture of the ${property.tenant} property",
                modifier = Modifier
                    .width(75.dp)
                    .height(75.dp)
                    .border(
                        1.dp,
                        color = MaterialTheme.colorScheme.primary,
                        shape = RoundedCornerShape(50.dp)
                    )
                    .clip(
                        RoundedCornerShape(50.dp)
                    )
            )
            Column(
                horizontalAlignment = Alignment.Start,
                modifier = Modifier.padding(start = 10.dp)
            ) {
                PropertyBoxTextLine(property.address, Icons.Outlined.Place)
                PropertyBoxTextLine(property.tenant?: "", Icons.Outlined.AccountCircle)
                PropertyBoxTextLine(
                    DateFormatter.formatOffsetDateTime(property.startDate)?: "---------------------",
                    Icons.Outlined.DateRange
                )
            }
        }
        Box(
            modifier = Modifier
                .align(Alignment.TopEnd)
                .padding(15.dp)
                .width(100.dp)
                .height(30.dp)
                .clip(RoundedCornerShape(30.dp))
                .background(if (property.available) MaterialTheme.colorScheme.surfaceVariant else MaterialTheme.colorScheme.error)
                .padding(3.dp)
                .testTag("topRightPropertyBoxInfo")
        ) {
            Text(
                color = MaterialTheme.colorScheme.onError,
                text = if (property.available) stringResource(R.string.available) else stringResource(R.string.busy),
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
    val viewModel: RealPropertyViewModel =
        viewModel(factory = RealPropertyViewModelFactory(navController))
    var detailsOpen by rememberSaveable { mutableStateOf<String?>(null) }
    var deleteOpen by rememberSaveable { mutableStateOf<Pair<String, String>?>(null) }
    var addPropertyModalOpen by rememberSaveable { mutableStateOf(false) }

    LaunchedEffect(Unit) {
        viewModel.getProperties()
    }

    DashBoardLayout(navController, "realPropertyScreen") {
        if (detailsOpen == null) {
            Row(horizontalArrangement = Arrangement.End, modifier = Modifier.fillMaxWidth()) {
                Button(
                    onClick = { addPropertyModalOpen = true },
                    colors = ButtonDefaults.buttonColors(containerColor = MaterialTheme.colorScheme.tertiary),
                    modifier = Modifier
                        .clip(RoundedCornerShape(5.dp))
                        .padding(5.dp)
                        .testTag("addAPropertyBtn")
                ) {
                    Text(
                        stringResource(R.string.add_prop),
                        color = MaterialTheme.colorScheme.onTertiary
                    )
                }
            }
            DeletePopUp(
                open = deleteOpen != null,
                delete = { viewModel.deleteProperty(deleteOpen!!.first) },
                close = { deleteOpen = null },
                globalName = stringResource(R.string.property),
                detailedName = deleteOpen?.second?: ""
            )
                LazyColumn {
                    items(viewModel.properties) { item ->
                        PropertyBox(item, onClick = { if (deleteOpen == null) detailsOpen = item.id }, onDelete = { deleteOpen = Pair(item.id, item.address) })
                    }
            }

        } else {
            RealPropertyDetailsScreen(
                navController,
                detailsOpen!!,
                getBack = { detailsOpen = null })
        }
        AddPropertyModal(
            addPropertyModalOpen,
            close = { addPropertyModalOpen = false },
            navController,
            { property -> viewModel.addProperty(property) }
        )
    }
}
