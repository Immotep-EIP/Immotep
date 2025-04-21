package com.example.immotep

import android.content.Context
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.safeDrawingPadding
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.compositionLocalOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.ApiService
import com.example.immotep.apiClient.mockApi.MockedApiService
import com.example.immotep.login.dataStore
import com.example.immotep.navigation.Navigation
import com.example.immotep.ui.theme.AppTheme
import com.example.immotep.utils.LanguageSetter
import kotlinx.coroutines.runBlocking
import java.util.Locale


val LocalApiService = compositionLocalOf<ApiService> {
    error("ApiService not provided")
}

var isTesting = false

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        val languageSetter  = LanguageSetter(this.dataStore)
        val language = runBlocking {
            languageSetter.getLanguage()
        }
        val locale = Locale(language)
        Locale.setDefault(locale)

        val config = resources.configuration
        config.setLocale(locale)
        config.setLayoutDirection(locale)
        resources.updateConfiguration(config, resources.displayMetrics)
        setContent {
            AppTheme {
                val apiService = if (isTesting) MockedApiService() else ApiClient.apiService
                CompositionLocalProvider(
                    LocalApiService provides apiService
                ) {
                    Box(Modifier.safeDrawingPadding()) {
                        Navigation()
                    }
                }
            }
        }
    }
}

