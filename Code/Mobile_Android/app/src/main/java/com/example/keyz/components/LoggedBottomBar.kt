package com.example.keyz.components


import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.Home
import androidx.compose.material.icons.outlined.HomeWork
import androidx.compose.material.icons.outlined.MailOutline
import androidx.compose.material.icons.outlined.Place
import androidx.compose.material.icons.outlined.Settings
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.NavigationBarItemDefaults
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

data class NavItem(
    val name: String,
    val icon: ImageVector,
    val iconDescription: String,
    val pageName: String
)

@Composable
fun LoggedBottomBar(navController: NavController) {
    val navigationItems = listOf(
        NavItem(
            stringResource(R.string.overview),
            Icons.Outlined.Home,
            "Home icon, go to the dashboard",
            "dashboard"
        ),
        NavItem(
            stringResource(R.string.RealProperty),
            Icons.Outlined.HomeWork,
            "Place icon, go to the real property page",
            "realProperty"
        ),
        NavItem(
            stringResource(R.string.messages),
            Icons.Outlined.MailOutline,
            "Message icon, go to the messages page",
            "messages"
        ),
        NavItem(
            stringResource(R.string.settings),
            Icons.Outlined.Settings,
            "Settings icon, go to the profile page",
            "profile"
        )
    )

    NavigationBar(
        containerColor = MaterialTheme.colorScheme.background,
        modifier = Modifier.testTag("loggedBottomBar")
    ) {
        navigationItems.forEachIndexed { _, (name, icon, iconDescription, pageName) ->
            val selected = navController.currentBackStackEntry?.destination?.route == pageName
            NavigationBarItem(
                selected = selected,
                onClick = {
                    navController.navigate(pageName)
                },
                icon = {
                    Icon(imageVector = icon, contentDescription = iconDescription)
                },
                label = {
                    Text(
                        name,
                        color = if (selected)
                            MaterialTheme.colorScheme.onPrimaryContainer
                        else Color.Gray
                    )
                },
                colors = NavigationBarItemDefaults.colors(
                    selectedIconColor = MaterialTheme.colorScheme.onPrimary,
                    unselectedIconColor = Color.Gray,
                    indicatorColor = MaterialTheme.colorScheme.primary,
                ),
                modifier = Modifier.testTag("loggedBottomBarElement $pageName")
            )
        }

    }
}
