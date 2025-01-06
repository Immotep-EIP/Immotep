package com.example.immotep.layouts


import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.immotep.components.LoggedBottomBar
import com.example.immotep.components.LoggedTopBar



import androidx.compose.foundation.Image
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AccountCircle
import androidx.compose.material.icons.outlined.Close
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.ui.Alignment
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.immotep.authService.AuthService
import com.example.immotep.R
import com.example.immotep.login.dataStore
import kotlinx.coroutines.launch


@Composable
fun InventoryTopBar(onExit: () -> Unit) {
    Row(
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier
            .testTag("inventoryTopBar")
            .height(35.dp)
            .padding(start = 10.dp, end = 10.dp)
            .drawBehind {
                val y = size.height - 2.dp.toPx() / 2
                drawLine(
                    Color.LightGray,
                    Offset(0f, y),
                    Offset(size.width, y),
                    2.dp.toPx()
                )
            }
    ) {
        Image(
            painter = painterResource(id = R.drawable.immotep_png_logo),
            contentDescription = stringResource(id = R.string.immotep_logo_desc),
            modifier = Modifier
                .size(35.dp)
                .padding(end = 10.dp)
                .testTag("inventoryTopBarImage")
                .clickable {
                },
        )
        Text(
            stringResource(R.string.inventory_title),
            fontSize = 20.sp,
            color = MaterialTheme.colorScheme.primary,
            fontWeight = FontWeight(500),
            modifier = Modifier.testTag("inventoryTopBarText")
        )
        Spacer(Modifier.weight(1f).fillMaxHeight())
        IconButton(
            onClick = onExit,
            modifier = Modifier.testTag("inventoryTopBarCloseIcon")
        ) {
            Icon(Icons.Outlined.Close, contentDescription = "Close inventory page")
        }
    }
}


@Composable
fun InventoryLayout(
    testTag: String,
    onExit: () -> Unit,
    content: @Composable () -> Unit
) {
    Column(modifier = Modifier.testTag(testTag)) {
        InventoryTopBar(onExit)
        Column(modifier = Modifier.weight(1f).padding(10.dp).testTag("dashboardLayout")) {
            content()
        }
    }
}
