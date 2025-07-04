package fr.keyz.navigation

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.MutableState
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import fr.keyz.LocalApiService
import fr.keyz.LocalIsOwner
import fr.keyz.apiClient.ApiService
import fr.keyz.authService.AuthService
import fr.keyz.damageDetails.DamageDetailsScreen
import fr.keyz.dashboard.DashBoardScreen
import fr.keyz.inventory.InventoryScreen
import fr.keyz.inventory.loaderButton.LoaderInventoryViewModel
import fr.keyz.login.LoginScreen
import fr.keyz.login.dataStore
import fr.keyz.messages.Messages
import fr.keyz.profile.ProfileScreen
import fr.keyz.realProperty.RealPropertyScreen
import fr.keyz.register.RegisterScreen
import kotlinx.coroutines.runBlocking

fun checkIfTokenIsPresent(navController: NavController, apiService: ApiService, isOwner: MutableState<Boolean>) {
    val authServ = AuthService(navController.context.dataStore, apiService)
    val currentRoute = navController.currentBackStackEntry?.destination?.route
    if (currentRoute != null && currentRoute != "dashboard")
        return
    runBlocking {
        try {
            authServ.getToken()
            isOwner.value = authServ.isUserOwner()
        } catch (e: Exception) {
            authServ.onLogout(navController)
        }
    }
}

@Composable
fun Navigation() {
    val navController = rememberNavController()
    val apiService = LocalApiService.current
    val isOwner = LocalIsOwner.current
    val loaderInventory = viewModel {
        LoaderInventoryViewModel(navController, apiService)
    }
    LaunchedEffect(Unit) {
        checkIfTokenIsPresent(navController, apiService, isOwner)
    }
    NavHost(navController = navController, startDestination = "dashboard") {
        composable("login") { LoginScreen(navController) }
        composable("dashboard") { DashBoardScreen(navController) }
        composable("register") { RegisterScreen(navController) }
        composable("profile") { ProfileScreen(navController) }
        composable("messages") { Messages(navController) }
        composable("realProperty") { RealPropertyScreen(navController, loaderInventory) }
        composable("inventory/{propertyId}/{leaseId}") { navBackStackEntry ->
            val propertyId = navBackStackEntry.arguments?.getString("propertyId")
            val currentLeaseId = navBackStackEntry.arguments?.getString("leaseId")
            if (propertyId != null && currentLeaseId != null) {
                InventoryScreen(
                    navController = navController,
                    propertyId = propertyId,
                    loaderViewModel = loaderInventory,
                    leaseId = currentLeaseId
                )
            }
        }
        composable("damage/{propertyId}/{leaseId}/{damageId}") { navBackStackEntry ->
            val propertyId = navBackStackEntry.arguments?.getString("propertyId")
            val currentLeaseId = navBackStackEntry.arguments?.getString("leaseId")
            val damageId = navBackStackEntry.arguments?.getString("damageId")
            if (currentLeaseId != null && damageId != null) {
                DamageDetailsScreen(
                    navController = navController,
                    propertyId = propertyId,
                    leaseId = currentLeaseId,
                    damageId = damageId
                )
            }

        }
    }
}
