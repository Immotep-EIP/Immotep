package com.example.immotep.apiCallerServices

import androidx.datastore.dataStore
import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.ApiService
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore

abstract class ApiCallerService(
    protected val apiService: ApiService,
    protected val navController: NavController
) {
    protected suspend fun getBearerToken() : String {
        val authService = AuthService(dataStore =  navController.context.dataStore, apiService)
        try {
            val bearerToken = authService.getBearerToken()
            return bearerToken
        } catch(e : Exception) {
            authService.onLogout(navController)
            throw e
        }
    }
}