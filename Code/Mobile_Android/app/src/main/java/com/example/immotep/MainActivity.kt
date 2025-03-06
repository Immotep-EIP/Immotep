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
import com.example.immotep.navigation.Navigation
import com.example.immotep.ui.theme.AppTheme

val LocalApiService = compositionLocalOf<ApiService> {
    error("ApiService not provided")
}

class MainActivity(val apiService: ApiService = ApiClient.apiService) : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            AppTheme {
                CompositionLocalProvider(LocalApiService provides apiService) {
                    Box(Modifier.safeDrawingPadding()) {
                        Navigation(apiService)
                    }
                }
            }
        }
    }
}
