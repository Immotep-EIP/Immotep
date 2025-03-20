package com.example.immotep

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.safeDrawingPadding
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.compositionLocalOf
import androidx.compose.ui.Modifier
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.ApiService
import com.example.immotep.apiClient.mockApi.MockedApiService
import com.example.immotep.navigation.Navigation
import com.example.immotep.ui.theme.AppTheme

val LocalApiService = compositionLocalOf<ApiService> {
    error("ApiService not provided")
}

var isTesting = false

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        val apiService = if (isTesting) MockedApiService() else ApiClient.apiService
        setContent {
            AppTheme {
                CompositionLocalProvider(LocalApiService provides apiService) {
                    Box(Modifier.safeDrawingPadding()) {
                        Navigation()
                    }
                }
            }
        }
    }
}

