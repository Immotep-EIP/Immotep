package com.example.immotep.components

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AccountCircle
import androidx.compose.material.icons.outlined.Lock
import androidx.compose.material3.Button
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
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
import androidx.navigation.NavController
import com.example.immotep.AuthService.AuthService
import com.example.immotep.R
import com.example.immotep.login.dataStore
import kotlinx.coroutines.launch


class LoggedTopBarViewModel : ViewModel() {
    fun logout(navController: NavController) {
        viewModelScope.launch {
            AuthService(navController.context.dataStore).onLogout(navController)
        }
    }
}


@Composable
fun LoggedTopBar(navController: NavController) {
    val viewModel : LoggedTopBarViewModel = viewModel()
    Row(verticalAlignment = Alignment.CenterVertically, modifier = Modifier.testTag("header").height(35.dp).padding(start = 10.dp, end = 10.dp).drawBehind {
        val y = size.height - 2.dp.toPx() / 2

        drawLine(
            Color.LightGray,
            Offset(0f, y),
            Offset(size.width, y),
            2.dp.toPx()
        )
    }) {
        Image(
            painter = painterResource(id = R.drawable.immotep_png_logo),
            contentDescription = stringResource(id = R.string.immotep_logo_desc),
            modifier = Modifier.size(35.dp).padding(end = 10.dp),
        )
        Text(stringResource(R.string.app_name), fontSize = 20.sp, color = MaterialTheme.colorScheme.primary, fontWeight = FontWeight(500))
        Spacer(Modifier.weight(1f).fillMaxHeight())
        IconButton (
            onClick = {
                viewModel.logout(navController)
            },
        ) {
            Icon(Icons.Outlined.AccountCircle, contentDescription = "Account circle, go back to the login page")
        }
    }
}