package com.example.keyz.navigation

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.MutableState
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.example.keyz.LocalApiService
import com.example.keyz.LocalIsOwner
import com.example.keyz.apiClient.ApiService
import com.example.keyz.authService.AuthService
import com.example.keyz.dashboard.DashBoardScreen
import com.example.keyz.inventory.InventoryScreen
import com.example.keyz.inventory.loaderButton.LoaderInventoryViewModel
import com.example.keyz.login.LoginScreen
import com.example.keyz.login.dataStore
import com.example.keyz.profile.ProfileScreen
import com.example.keyz.realProperty.RealPropertyScreen
import com.example.keyz.register.RegisterScreen
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
    }
}
