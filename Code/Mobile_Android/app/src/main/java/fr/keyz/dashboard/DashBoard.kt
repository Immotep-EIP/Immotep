package fr.keyz.dashboard

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowColumn
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.IconButton
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.MoreVert
import androidx.compose.material3.DropdownMenu
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
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import fr.keyz.LocalApiService
import fr.keyz.components.LoadingDialog
import fr.keyz.R
import fr.keyz.apiCallerServices.DashBoardReminder
import fr.keyz.components.InitialFadeIn
import fr.keyz.components.InternalLoading
import fr.keyz.dashboard.widgets.DamageWidget
import fr.keyz.dashboard.widgets.DamagesListWidget
import fr.keyz.dashboard.widgets.HelloWidget
import fr.keyz.dashboard.widgets.PropertiesWidget
import fr.keyz.dashboard.widgets.RemindersWidget
import fr.keyz.dashboard.widgets.ScheduledInventoryWidget
import fr.keyz.dashboard.widgets.UnreadMessagesWidget
import fr.keyz.dashboard.widgets.WidgetMenuItem
import fr.keyz.realProperty.details.RealPropertyDropDownMenuItem
import fr.keyz.realProperty.details.tabs.OneDamage
import fr.keyz.ui.components.StyledButton

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
    val userName = viewModel.userName.collectAsState()

    LaunchedEffect(Unit) {
        viewModel.getDashBoard()
        viewModel.getName()
    }

    DashBoardLayout(navController, "dashboardScreen") {
        if (isLoading.value) {
            InternalLoading()
            return@DashBoardLayout
        }
        InitialFadeIn(durationMs = 200) {
            Column(modifier = Modifier.verticalScroll(rememberScrollState())) {
                HelloWidget(dashBoard.value.properties.nbrTotal, userName = userName.value)
                RemindersWidget(dashBoard.value.reminders)
                PropertiesWidget(dashBoard.value.properties)
                //UnreadMessagesWidget()
                //ScheduledInventoryWidget()
                DamageWidget(dashBoard.value.openDamages)
                DamagesListWidget(dashBoard.value.openDamages.listToFix)
            }
        }
    }
}
