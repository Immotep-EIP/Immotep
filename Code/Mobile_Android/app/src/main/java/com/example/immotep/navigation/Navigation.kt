package com.example.immotep.navigation

import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.navigation.NavController
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.example.immotep.authService.AuthService
import com.example.immotep.dashboard.DashBoardScreen
import com.example.immotep.inventory.InventoryScreen
import com.example.immotep.login.LoginScreen
import com.example.immotep.login.dataStore
import com.example.immotep.profile.ProfileScreen
import com.example.immotep.realProperty.RealPropertyScreen
import com.example.immotep.register.RegisterScreen
import kotlinx.coroutines.runBlocking

fun checkIfTokenIsPresent(navController: NavController) {
    val currentRoute = navController.currentBackStackEntry?.destination?.route
    if (currentRoute != null && currentRoute != "login")
        return
    runBlocking {
        try {
            AuthService(navController.context.dataStore).getToken()
            navController.navigate("dashboard")
        } catch (e: Exception) {
            navController.navigate("dashboard")
        }
    }
}

@Composable
fun Navigation() {
    val navController = rememberNavController()
    LaunchedEffect(Unit) {
        checkIfTokenIsPresent(navController)
    }
    NavHost(navController = navController, startDestination = "login") {
        composable("login") { LoginScreen(navController) }
        composable("dashboard") { DashBoardScreen(navController) }
        composable("register") { RegisterScreen(navController) }
        composable("profile") { ProfileScreen(navController) }
        composable("realProperty") { RealPropertyScreen(navController) }
        composable("inventory/{propertyId}") {
            navBackStackEntry ->
            val propertyId = navBackStackEntry.arguments?.getString("propertyId")
            propertyId?.let {
                InventoryScreen(navController, propertyId)
            }
        }
    }
}
