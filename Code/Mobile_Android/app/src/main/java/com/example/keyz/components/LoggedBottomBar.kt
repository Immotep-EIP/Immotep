package com.example.keyz.components

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.material.Icon
import androidx.compose.material.Text
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.Home
import androidx.compose.material.icons.outlined.HomeWork
import androidx.compose.material.icons.outlined.MailOutline
import androidx.compose.material.icons.outlined.Place
import androidx.compose.material.icons.outlined.Settings
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.navigation.NavController
import com.example.keyz.R

@Composable
fun LoggedBottomBarElement(navController: NavController, name: String, icon: ImageVector, iconDescription: String, pageName: String) {
    val selected = navController.currentBackStackEntry?.destination?.route == pageName
    val lineColor = MaterialTheme.colorScheme.secondary
    Column(
        horizontalAlignment = Alignment.CenterHorizontally,
        modifier = Modifier
            .drawBehind {
                if (!selected) {
                    return@drawBehind
                }
                val y = size.height + 2.dp.toPx()
                drawLine(
                    lineColor,
                    Offset(size.width / 2 + size.width / 6, y),
                    Offset(size.width / 3, y),
                    2.dp.toPx()
                )
            }
            .clickable {
                if (!selected) {
                    navController.navigate(pageName)
                }
            }
            .testTag("loggedBottomBarElement $pageName")
    ) {
        Icon(
            imageVector = icon,
            contentDescription = iconDescription,
            tint = MaterialTheme.colorScheme.onPrimaryContainer
        )
        Text(
            text = name,
            color = MaterialTheme.colorScheme.onPrimaryContainer
        )
    }
}

@Composable
fun LoggedBottomBar(navController: NavController) {
    Spacer(
        modifier = Modifier.fillMaxWidth().padding(start = 10.dp, end = 10.dp).height(1.dp).drawBehind {
            val y = size.height - 2.dp.toPx() / 2
            drawLine(
                Color.LightGray,
                Offset(0f, y),
                Offset(size.width, y),
                2.dp.toPx()
            )
        }
    )
    Spacer(modifier = Modifier.height(5.dp))
    Row(
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier.fillMaxWidth().padding(start = 10.dp, end = 10.dp).testTag("loggedBottomBar"),
        horizontalArrangement = Arrangement.SpaceBetween
    ) {

        LoggedBottomBarElement(navController, stringResource(R.string.overview), Icons.Outlined.Home, "Home icon, go to the dashboard", "dashboard")
        LoggedBottomBarElement(navController, stringResource(R.string.RealProperty), Icons.Outlined.HomeWork, "Place icon, go to the real property page", "realProperty")
        LoggedBottomBarElement(navController, stringResource(R.string.messages), Icons.Outlined.MailOutline, "Message icon, go to the messages page", "messages")
        LoggedBottomBarElement(navController, stringResource(R.string.settings), Icons.Outlined.Settings, "Settings icon, go to the profile page", "profile")
    }
    Spacer(modifier = Modifier.height(10.dp))
}
