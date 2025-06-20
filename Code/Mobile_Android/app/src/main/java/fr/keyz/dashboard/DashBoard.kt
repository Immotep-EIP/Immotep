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

/*

{
  "reminders": [
    {
      "id": "1",
      "priority": "high",
      "title": "Lease of property Nouvel appart is ending in 7 days.",
      "advice": "Plan an inventory appointment with the tenant to fill the inventory report.",
      "link": "/real-property/details/cmc0jghdt0001183toy4y77gw"
    },
    {
      "id": "6",
      "priority": "high",
      "title": "New medium damage reported in room Chambre of property Nouvel appart.",
      "advice": "Please review and plan a fix date.",
      "link": "/real-property/details/cmc0jghdt0001183toy4y77gw/damage/cmc2prtne000z1bx29bcupmm1"
    },
    {
      "id": "6",
      "priority": "high",
      "title": "New medium damage reported in room Chambre of property Nouvel appart.",
      "advice": "Please review and plan a fix date.",
      "link": "/real-property/details/cmc0jghdt0001183toy4y77gw/damage/cmc2prvnt00121bx2225b3le0"
    }
  ],
  "properties": {
    "nbr_total": 2,
    "nbr_archived": 0,
    "nbr_occupied": 2,
    "nbr_available": 0,
    "nbr_pending_invites": 0,
    "list_recently_added": [
      {
        "id": "cmc0jghdt0001183toy4y77gw",
        "name": "Nouvel appart",
        "address": "52 rue de la sinne",
        "apartment_number": "21",
        "city": "Mulhouse",
        "postal_code": "68100",
        "country": "France",
        "area_sqm": 70,
        "rental_price_per_month": 1100,
        "deposit_price": 1100,
        "created_at": "2025-06-17T13:07:59.872Z",
        "archived": false,
        "owner_id": "cmc0ag8v50000183tpb3k7llo"
      },
      {
        "id": "cmc3c6ksc00141bx22dgk0bvt",
        "name": "House 1",
        "address": "51b Baker Street",
        "apartment_number": "2",
        "city": "London",
        "postal_code": "57300",
        "country": "United KIngdom",
        "area_sqm": 30,
        "rental_price_per_month": 3000,
        "deposit_price": 3000,
        "created_at": "2025-06-19T12:07:38.938Z",
        "archived": false,
        "picture_id": "cmc4627rr001o1bx2vix6td5b",
        "owner_id": "cmc0ag8v50000183tpb3k7llo"
      }
    ]
  },
  "open_damages": {
    "nbr_total": 2,
    "nbr_urgent": 0,
    "nbr_high": 0,
    "nbr_medium": 2,
    "nbr_low": 0,
    "nbr_planned_to_fix_this_week": 0,
    "list_to_fix": [
      {
        "id": "cmc2prtne000z1bx29bcupmm1",
        "lease_id": "cmc1keivw0013183t0vvp7m6l",
        "tenant_name": "TenantR MyTen",
        "property_id": "cmc0jghdt0001183toy4y77gw",
        "property_name": "Nouvel appart",
        "room_id": "cmc1m478n0014183ttqlm91ow",
        "room_name": "Chambre",
        "comment": "application en ligne !",
        "priority": "medium",
        "read": false,
        "created_at": "2025-06-19T01:40:19.034Z",
        "updated_at": "2025-06-19T01:40:19.034Z",
        "fix_status": "pending",
        "fix_planned_at": null
      },
      {
        "id": "cmc2prvnt00121bx2225b3le0",
        "lease_id": "cmc1keivw0013183t0vvp7m6l",
        "tenant_name": "TenantR MyTen",
        "property_id": "cmc0jghdt0001183toy4y77gw",
        "property_name": "Nouvel appart",
        "room_id": "cmc1m478n0014183ttqlm91ow",
        "room_name": "Chambre",
        "comment": "application en ligne !",
        "priority": "medium",
        "read": false,
        "created_at": "2025-06-19T01:40:21.642Z",
        "updated_at": "2025-06-19T01:40:21.642Z",
        "fix_status": "pending",
        "fix_planned_at": null
      }
    ]
  }
}
 */










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
