package fr.keyz.damageDetails

import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.aspectRatio
import androidx.compose.foundation.layout.defaultMinSize
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.DatePicker
import androidx.compose.material3.DatePickerDialog
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.Text
import androidx.compose.material3.TextButton
import androidx.compose.material3.rememberDatePickerState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import fr.keyz.LocalApiService
import fr.keyz.LocalIsOwner
import fr.keyz.layouts.InventoryLayout
import fr.keyz.R
import fr.keyz.apiCallerServices.DamageStatus
import fr.keyz.components.Base64ImageView
import fr.keyz.components.InternalLoading
import fr.keyz.components.PriorityBox
import fr.keyz.layouts.BigModalLayout
import fr.keyz.ui.components.StyledButton

@OptIn(ExperimentalLayoutApi::class)
@Composable
fun DamageImagesPanel(
    images : Array<String>
) {
    Column(
        modifier = Modifier.fillMaxWidth(),
        horizontalAlignment = Alignment.Start
    ) {
        Text(stringResource(R.string.state), fontSize = 24.sp, fontWeight = FontWeight.SemiBold)
        Spacer(modifier = Modifier.height(5.dp))
        FlowRow(modifier = Modifier.defaultMinSize(minHeight = 125.dp)) {
            images.forEachIndexed { index, image ->
                Base64ImageView(
                    image = image,
                    description = "Damage image $index",
                    modifier = Modifier.fillMaxWidth(0.33f).aspectRatio(1f)
                )
            }
        }
    }
}

@Composable
fun DamageStatusPanel(
    damageStatus: DamageStatus,
    confirm: () -> Unit,
    pendingClick: () -> Unit,
    isOwner : Boolean
) {
    val confirmOwner = damageStatus == DamageStatus.AWAITING_OWNER_CONFIRMATION && isOwner
    val confirmTenant = damageStatus == DamageStatus.AWAITING_TENANT_CONFIRMATION && !isOwner
    val statusColor = when (damageStatus) {
        DamageStatus.FIXED-> MaterialTheme.colorScheme.surfaceVariant
        DamageStatus.PLANNED -> MaterialTheme.colorScheme.secondary
        else -> MaterialTheme.colorScheme.inversePrimary
    }
    val statusText = when (damageStatus) {
        DamageStatus.FIXED-> stringResource(R.string.fixed)
        DamageStatus.PLANNED -> stringResource(R.string.planned)
        DamageStatus.PENDING -> stringResource(R.string.pending)
        DamageStatus.AWAITING_OWNER_CONFIRMATION -> stringResource(if (isOwner) R.string.click_here_to_confirm else R.string.awaiting_owner_confirmation)
        DamageStatus.AWAITING_TENANT_CONFIRMATION -> stringResource(if (!isOwner) R.string.click_here_to_confirm else R.string.awaiting_tenant_confirmation)
    }
    val onClick = {
        if (confirmOwner || confirmTenant) {
            confirm()
        } else if (damageStatus == DamageStatus.PENDING || damageStatus == DamageStatus.PLANNED) {
            pendingClick()
        }
    }
    Column(
        modifier = Modifier.fillMaxWidth(),
        horizontalAlignment = Alignment.Start
    ) {
        Text(stringResource(R.string.state), fontSize = 24.sp, fontWeight = FontWeight.SemiBold)
        Spacer(modifier = Modifier.height(5.dp))
        Box(
            modifier = Modifier
                .fillMaxWidth()
                .shadow(
                    elevation = 6.dp,
                    shape = RoundedCornerShape(10.dp),
                    clip = false
                )
                .clip(RoundedCornerShape(10.dp))
                .background(color = statusColor, shape = RoundedCornerShape(10.dp))
                .clickable(onClick = onClick)
                .padding(10.dp),

            contentAlignment = Alignment.Center
        ) {
            Text(statusText)
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun DamageResolveBottomModal(open: Boolean, close: () -> Unit, onSubmit : (Long?) -> Unit) {
    var showDatePicker by remember { mutableStateOf(false) }
    val datePickerState = rememberDatePickerState()
    if (showDatePicker) {
        DatePickerDialog(
            onDismissRequest = { showDatePicker = false },
            confirmButton = {
                TextButton(onClick = {
                    onSubmit(datePickerState.selectedDateMillis)
                    showDatePicker = false
                }) {
                    Text(stringResource(R.string.confirm))
                }
            },
            dismissButton = {
                TextButton(onClick = { showDatePicker = false }) {
                    Text(stringResource(R.string.cancel))
                }
            }
        ) {
            DatePicker(state = datePickerState)
        }
    }
    BigModalLayout(height = 0.3f, open = open, close = close, testTag = "damageResolveBottomModal") {
        Column(
            modifier = Modifier
                .fillMaxWidth()
                .fillMaxHeight(0.95f)
                .verticalScroll(rememberScrollState())
                .testTag("damageResolveBottomModalInternalContainer")
        ) {
            StyledButton(
                onClick = {
                    close()
                    showDatePicker = true
                },
                text = stringResource(R.string.in_progress)
            )
            StyledButton(
                onClick = {
                    close()
                    onSubmit(null)
                },
                text = stringResource(R.string.fixed)
            )
        }
    }
}

@Composable
fun DamageDetailsScreen(navController: NavController, propertyId: String?, leaseId: String, damageId: String) {
    val isOwner = LocalIsOwner.current
    val apiService = LocalApiService.current
    val viewModel = viewModel {
        DamageDetailsViewModel(apiService, navController)
    }
    val isLoading = viewModel.isLoading.collectAsState()
    val apiError = viewModel.apiError.collectAsState()
    val damage = viewModel.currentDamage.collectAsState()

    var damageResolveBottomModalIsOpen by remember { mutableStateOf(false) }


    LaunchedEffect(damageId) {
        viewModel.getDamage(propertyId, leaseId, damageId)
    }

    BackHandler {
        navController.popBackStack()
    }

    InventoryLayout(
        testTag = "damageDetails",
        onExit = { navController.popBackStack() },
        customTitle = stringResource(R.string.damage_detail)
    ) {
        DamageResolveBottomModal(
            open = damageResolveBottomModalIsOpen,
            close = { damageResolveBottomModalIsOpen = false },
            onSubmit = { date -> viewModel.onSubmitUpdateDamageResolution(date, propertyId) }
        )
        if (isLoading.value || damage.value == null) {
            InternalLoading()
        } else {
            Row(
                modifier = Modifier.fillMaxWidth(),
                verticalAlignment = Alignment.CenterVertically,
                horizontalArrangement = Arrangement.SpaceBetween
            ) {
                Text(text = damage.value!!.roomName, fontSize = 24.sp, fontWeight = FontWeight.Bold)
                PriorityBox(damage.value!!.priority)
            }
            Spacer(modifier = Modifier.height(20.dp))
            Column(
                modifier = Modifier.fillMaxWidth(),
                horizontalAlignment = Alignment.Start
            ) {
                Text(stringResource(R.string.comment), fontSize = 24.sp, fontWeight = FontWeight.SemiBold)
                Spacer(modifier = Modifier.height(5.dp))
                Text(damage.value!!.comment)
            }
            Spacer(modifier = Modifier.height(20.dp))
            DamageStatusPanel(damage.value!!.fixStatus, {}, {}, isOwner.value)
            Spacer(modifier = Modifier.height(20.dp))
            DamageImagesPanel(damage.value!!.pictures)
        }
    }

}