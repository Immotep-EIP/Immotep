package com.example.keyz.dashboard

import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.IconButton
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.MoreVert
import androidx.compose.material3.Button
import androidx.compose.material3.DropdownMenu
import androidx.compose.material3.DropdownMenuItem
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButtonDefaults
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.runtime.getValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.keyz.LocalApiService
import com.example.keyz.components.LoadingDialog
import com.example.keyz.R
import com.example.keyz.components.InitialFadeIn
import com.example.keyz.realProperty.details.RealPropertyDropDownMenuItem

data class WidgetMenuItem(
    val label : String,
    val onClick : (() -> Unit)?,
    val disabled : Boolean = false,
    val testTag : String
)

@Composable
fun WidgetBase(
    title : String? = null,
    dropDownItems : Array<WidgetMenuItem>,
    testTag : String,
    content : @Composable () -> Unit
) {
    var expanded by rememberSaveable { mutableStateOf(false) }
    Column(modifier = Modifier.fillMaxWidth().padding(top = 5.dp, bottom = 10.dp, start = 5.dp, end = 5.dp)) {
        if (title != null) {
            Text(title)
        }
        Box(
            modifier = Modifier
                .fillMaxWidth()
                .border(1.dp, MaterialTheme.colorScheme.onBackground, RoundedCornerShape(5.dp))
        ) {
            Box(modifier = Modifier.padding(10.dp)) {
                content()
            }
            Box(
                modifier = Modifier
                    .align(Alignment.TopStart)
                    .padding(top = 0.dp, start = 0.dp)
            ) {
                Row(
                    modifier = Modifier.fillMaxWidth(),
                    horizontalArrangement = Arrangement.End,
                    verticalAlignment = Alignment.CenterVertically
                ) {
                    Box {
                        IconButton(
                            onClick = { expanded = true },
                            colors = IconButtonDefaults.iconButtonColors(containerColor = MaterialTheme.colorScheme.background),
                            modifier = Modifier.testTag("moreVertWidget$testTag"),
                        ) {
                            Icon(
                                Icons.Outlined.MoreVert,
                                contentDescription = "More options",
                                tint = MaterialTheme.colorScheme.onBackground
                            )
                        }
                        DropdownMenu(
                            expanded = expanded,
                            onDismissRequest = { expanded = false }
                        ) {
                            dropDownItems.forEach {
                                RealPropertyDropDownMenuItem(
                                    name = it.label,
                                    onClick = it.onClick,
                                    disabled = it.disabled,
                                    closeDropDown = { expanded = false },
                                    testTag = it.testTag
                                )
                            }
                        }
                    }
                }
            }
        }
    }
}

@Composable
fun HelloWidget(nbOfProperties : Int) {
    WidgetBase(dropDownItems = arrayOf(), testTag = "helloWidget") {
        Column(modifier = Modifier.fillMaxWidth(), verticalArrangement = Arrangement.SpaceBetween) {
            Text(stringResource(R.string.welcome))
            Spacer(modifier = Modifier.height(5.dp))
            Text("${stringResource(R.string.here_overview)} $nbOfProperties ${stringResource(R.string.properties)}")
        }
    }
}

@Composable
fun UnreadMessagesWidget() {
    WidgetBase(title = stringResource(R.string.unread_messages), dropDownItems = arrayOf(), testTag = "unreadMessageWidget") {
        Box(modifier = Modifier.height(100.dp), contentAlignment = Alignment.Center) {

        }
    }
}

@Composable
fun ScheduledInventoryWidget() {
    WidgetBase(title = stringResource(R.string.unread_messages), dropDownItems = arrayOf(), testTag = "scheduledInventoryWidget") {
        Box(modifier = Modifier.height(100.dp), contentAlignment = Alignment.Center) {

        }
    }
}

@Composable
fun DamageInProgressWidget() {
    WidgetBase(title = stringResource(R.string.damage_in_progress), dropDownItems = arrayOf(), testTag = "damageInProgressWidget") {
        Box(modifier = Modifier.height(100.dp), contentAlignment = Alignment.Center) {

        }
    }
}

@Composable
fun AvailablePropertiesWidget() {
    WidgetBase(title = stringResource(R.string.available_properties), dropDownItems = arrayOf(), testTag = "availablePropertiesWidget") {
        Box(modifier = Modifier.height(100.dp), contentAlignment = Alignment.Center) {

        }
    }
}


@Composable
fun DashBoardScreen(
    navController: NavController,
) {
    val apiService = LocalApiService.current
    val viewModel : DashBoardViewModel = viewModel {
        DashBoardViewModel(navController, apiService)
    }
    val isLoading = viewModel.isLoading.collectAsState()
    val dashBoard = viewModel.dashBoard.collectAsState()


    LaunchedEffect(Unit) {
        viewModel.getDashBoard()
    }

    DashBoardLayout(navController, "dashboardScreen") {
        LoadingDialog(isLoading.value)
        InitialFadeIn(durationMs = 300) {
            Column(modifier = Modifier.verticalScroll(rememberScrollState())) {
                HelloWidget(dashBoard.value.properties.nbrTotal)
                UnreadMessagesWidget()
                ScheduledInventoryWidget()
                DamageInProgressWidget()
                AvailablePropertiesWidget()
            }
        }
    }
}
