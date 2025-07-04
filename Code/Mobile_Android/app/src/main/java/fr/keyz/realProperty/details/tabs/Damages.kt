package fr.keyz.realProperty.details.tabs

import androidx.compose.foundation.Image
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowColumn
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.aspectRatio
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController
import fr.keyz.LocalIsOwner
import fr.keyz.R
import fr.keyz.components.addDamageModal.AddDamageModal
import fr.keyz.apiCallerServices.Damage
import fr.keyz.ui.components.StyledButton
import fr.keyz.utils.Base64Utils
import fr.keyz.utils.DateFormatter


@Composable
fun OneDamage(damage: Damage, goToDamageDetails : (Damage) -> Unit) {
    val dateCreationAsString = DateFormatter.formatOffsetDateTime(damage.createdAt)
    val bitmap = try {
        Base64Utils.decodeBase64ToImage(damage.pictures.first())
    } catch (e : Exception) {
        null
    }
    Row(
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.SpaceBetween,
        modifier = Modifier
            .fillMaxWidth()
            .testTag("oneDamage ${damage.id}")
            .clickable { goToDamageDetails(damage) }
            .padding(5.dp).drawBehind {
            val y = size.height - 2.dp.toPx() / 2
            drawLine(
                Color.LightGray,
                Offset(0f, y),
                Offset(size.width, y),
                2.dp.toPx()
            )
        }
    ) {
        Column(modifier = Modifier
            .fillMaxWidth(0.7f)
            .padding(top = 5.dp, bottom = 5.dp, end = 5.dp)
        ) {
            Text(
                damage.roomName,
                fontSize = 15.sp,
                fontWeight = FontWeight.SemiBold,
                color = MaterialTheme.colorScheme.onPrimaryContainer
            )
            Text(
                damage.comment,
                fontSize = 12.sp,
                color = MaterialTheme.colorScheme.onTertiary
            )
        }
        Column(modifier = Modifier.fillMaxWidth().padding(bottom = 5.dp),
            horizontalAlignment = Alignment.CenterHorizontally) {
            Text(dateCreationAsString?: "", fontSize = 12.sp, color = MaterialTheme.colorScheme.onTertiary)
            if (bitmap != null) {
                Image(
                    modifier = Modifier
                        .fillMaxWidth()
                        .aspectRatio(1f)
                        .padding(top = 2.dp),
                    bitmap = bitmap,
                    contentDescription = "First Image of the damage ${damage.comment}"
                )
            } else {
                Text(
                    stringResource(R.string.picture_not_supported),
                    modifier = Modifier.padding(top = 10.dp)
                )
            }
        }
    }
}

@OptIn(ExperimentalLayoutApi::class)
@Composable
fun Damages(
    damageList : List<Damage>,
    addDamage : (Damage) -> Unit,
    navController: NavController,
    propertyId : String
) {
    val isOwner = LocalIsOwner.current.value
    var addDamageOpen by rememberSaveable { mutableStateOf(false) }
    if (!isOwner) {
        AddDamageModal(
            open = addDamageOpen,
            onClose = { addDamageOpen = false },
            addDamage = { addDamage(it) },
            navController = navController
        )
    }
    FlowColumn(
        modifier = Modifier
            .testTag("realPropertyDetailsDamagesTab")
            .fillMaxSize()
            .verticalScroll(rememberScrollState())
    ) {
        if (!isOwner) {
            StyledButton(
                onClick = { addDamageOpen = true },
                text = stringResource(R.string.report_claim),
                testTag = "reportClaimButton"
            )
        }
        damageList.forEach { item ->
            OneDamage(
                damage = item,
                goToDamageDetails = {
                    damage -> navController.navigate("damage/${propertyId}/${damage.leaseId}/${damage.id}" )
                }
            )
        }
    }
}