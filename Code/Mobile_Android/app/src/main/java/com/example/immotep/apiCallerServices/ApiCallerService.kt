package com.example.immotep.apiCallerServices

import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiService
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import retrofit2.HttpException

class ApiCallerServiceException(message:String): Exception(message) {
    fun getCode() : Int {
        return try {
                this.message?.toInt() ?: 400
        } catch (e : NumberFormatException) {
            return 400
        }
    }
}

sealed class ApiCallerService(
    protected val apiService: ApiService,
    protected val navController: NavController
) {
    private val authService = AuthService(dataStore =  navController.context.dataStore, apiService)
    protected suspend fun getBearerToken() : String {
        try {
            val bearerToken = authService.getBearerToken()
            return bearerToken
        } catch(e : Exception) {
            authService.onLogout(navController)
            throw e
        }
    }

    protected suspend fun <T>changeRetrofitExceptionByApiCallerException(
        logoutOnUnauthorized : Boolean = false,
        fnToRun : suspend () -> T,
    ) : T {
        try {
            return fnToRun()
        } catch (e : HttpException) {
            println("error response : ${e.response()}")
            if (e.code() == 401 && logoutOnUnauthorized) {
                authService.onLogout(navController)
            }
            throw ApiCallerServiceException(e.code().toString())
        } catch (e : Exception) {
            println("error : ${e.message}")
            throw ApiCallerServiceException("500")
        }
    }
}