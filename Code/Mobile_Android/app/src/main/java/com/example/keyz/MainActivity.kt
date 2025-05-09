package com.example.keyz

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.safeDrawingPadding
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.MutableState
import androidx.compose.runtime.compositionLocalOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier
import com.example.keyz.apiClient.ApiClient
import com.example.keyz.apiClient.ApiService
import com.example.keyz.apiClient.mockApi.MockedApiService
import com.example.keyz.login.dataStore
import com.example.keyz.navigation.Navigation
import com.example.keyz.ui.theme.AppTheme
import com.example.keyz.utils.LanguageSetter
import kotlinx.coroutines.runBlocking
import java.util.Locale


val LocalApiService = compositionLocalOf<ApiService> {
    error("ApiService not provided")
}

val LocalIsOwner = compositionLocalOf<MutableState<Boolean>> {
    error("No local is owner provided")
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
                val isOwner = remember { mutableStateOf(false) }
                CompositionLocalProvider(
                    LocalApiService provides apiService,
                    LocalIsOwner provides isOwner
                ) {
                    Box(Modifier.background(color = MaterialTheme.colorScheme.onPrimary).safeDrawingPadding()) {
                        Navigation()
                    }
                }
            }
        }
    }
}

