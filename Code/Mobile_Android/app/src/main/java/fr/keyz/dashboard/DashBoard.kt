package fr.keyz.dashboard

import androidx.compose.runtime.Composable
import androidx.navigation.NavController
import fr.keyz.LocalIsOwner
import fr.keyz.dashboard.tenant.TenantDashBoard


@Composable
fun DashBoardScreen(
    navController: NavController,
) {
    val isOwner = LocalIsOwner.current.value
    if (isOwner) {
        DashBoardScreenOwner(navController)
    } else {
        TenantDashBoard(navController)
    }
}