package com.example.immotep.apiClient

import android.content.Context
import androidx.compose.ui.platform.LocalContext
import androidx.navigation.NavController
import androidx.navigation.compose.rememberNavController
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import kotlinx.coroutines.runBlocking
import okhttp3.Interceptor
import okhttp3.OkHttpClient
import okhttp3.Response
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory




object RetrofitClient {
    // private val BASE_URL = "https://test1.icytree-5b429d30.eastus.azurecontainerapps.io"
    private const val BASE_URL = "http://10.0.2.2:8080/"


    val retrofit: Retrofit by lazy {
        Retrofit
            .Builder()
            .baseUrl(BASE_URL)
            .addConverterFactory(GsonConverterFactory.create())
            .build()
    }
}

object ApiClient {
    val apiService: ApiService by lazy {
        RetrofitClient.retrofit.create(ApiService::class.java)
    }
}
