package com.example.immotep.navigation

import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.MutableState
import androidx.compose.runtime.compositionLocalOf
import androidx.compose.runtime.remember
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.example.immotep.LocalApiService
import com.example.immotep.LocalIsOwner
import com.example.immotep.apiClient.ApiService
import com.example.immotep.authService.AuthService
import com.example.immotep.dashboard.DashBoardScreen
import com.example.immotep.inventory.InventoryScreen
import com.example.immotep.inventory.loaderButton.LoaderInventoryViewModel
import com.example.immotep.login.LoginScreen
import com.example.immotep.login.dataStore
import com.example.immotep.profile.ProfileScreen
import com.example.immotep.realProperty.RealPropertyScreen
import com.example.immotep.register.RegisterScreen
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
