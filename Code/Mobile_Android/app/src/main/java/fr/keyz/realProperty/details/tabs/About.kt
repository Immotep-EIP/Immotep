package fr.keyz.realProperty.details.tabs

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AccountBalanceWallet
import androidx.compose.material.icons.outlined.RequestQuote
import androidx.compose.material.icons.outlined.Straighten
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
import androidx.navigation.NavController
import fr.keyz.LocalIsOwner
import fr.keyz.R
import fr.keyz.apiCallerServices.DetailedProperty
import fr.keyz.apiCallerServices.PropertyStatus
import fr.keyz.inventory.loaderButton.LoaderInventoryButton
import fr.keyz.inventory.loaderButton.LoaderInventoryViewModel
import fr.keyz.utils.DateFormatter

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
        Text(value, fontWeight = FontWeight.Bold, fontSize = 15.sp, color = MaterialTheme.colorScheme.onPrimaryContainer)
        Text(name, fontWeight = FontWeight.Thin, fontSize = 12.sp, color = MaterialTheme.colorScheme.onPrimaryContainer)
    }
}


@Composable
fun AboutPropertyTab(
    property : State<DetailedProperty>,
    navController: NavController,
    setIsLoading : (Boolean) -> Unit,
    loaderInventoryViewModel: LoaderInventoryViewModel
) {
    val isOwner = LocalIsOwner.current.value
    val invite = if (property.value.status == PropertyStatus.invite_sent) {
        property.value.invite
    } else {
        null
    }
    val lease = if (property.value.status == PropertyStatus.unavailable) {
        property.value.lease
    } else {
        null
    }
    Column(modifier = Modifier
        .fillMaxSize()
        .verticalScroll(rememberScrollState()
        )
    ) {
        Row(
            modifier = Modifier.fillMaxWidth().padding(16.dp)
                .testTag("realPropertyDetailsAboutTab"),
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
                Text(
                    "${stringResource(R.string.tenant)}:",
                    fontWeight = FontWeight.Bold,
                    fontSize = 15.sp,
                    color = MaterialTheme.colorScheme.onPrimaryContainer
                )
                Text(
                    text = lease?.tenantName ?: invite?.tenantEmail ?: "---------------------",
                    fontSize = 15.sp,
                    color = MaterialTheme.colorScheme.onPrimaryContainer
                )
            }
            Spacer(Modifier.width(15.dp))
            Column {
                Text(
                    "${stringResource(R.string.dates)}:",
                    fontWeight = FontWeight.Bold,
                    fontSize = 15.sp,
                    color = MaterialTheme.colorScheme.onPrimaryContainer
                )
                Text(
                    if (lease != null) {
                        "${DateFormatter.formatOffsetDateTime(lease.startDate)} - ${
                            DateFormatter.formatOffsetDateTime(
                                lease.endDate
                            )
                        }"
                    } else if (invite != null) {
                        "${DateFormatter.formatOffsetDateTime(invite.startDate)} - ${
                            DateFormatter.formatOffsetDateTime(
                                invite.endDate
                            )
                        }"
                    } else {
                        "---------------------"
                    },
                    fontSize = 15.sp,
                    color = MaterialTheme.colorScheme.onPrimaryContainer
                )

            }
        }
        if (isOwner && property.value.status == PropertyStatus.unavailable && lease != null) {
            LoaderInventoryButton(
                propertyId = property.value.id,
                currentLeaseId = lease.id,
                setIsLoading = setIsLoading,
                viewModel = loaderInventoryViewModel,
            )
        }
    }
}
