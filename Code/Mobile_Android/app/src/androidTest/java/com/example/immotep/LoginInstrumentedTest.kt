package com.example.immotep

import android.content.res.Resources
import androidx.compose.ui.test.assert
import androidx.compose.ui.test.assertIsDisplayed
import androidx.compose.ui.test.hasText
import androidx.compose.ui.test.junit4.createAndroidComposeRule
import androidx.compose.ui.test.onNodeWithTag
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.performClick
import androidx.compose.ui.test.performTextInput
import androidx.test.ext.junit.runners.AndroidJUnit4
import androidx.test.platform.app.InstrumentationRegistry
import com.example.immotep.RetrofitTestClient.BASE_URL
import com.example.immotep.apiClient.ApiService
import com.example.immotep.apiClient.LoginResponse
import com.example.immotep.apiClient.RetrofitClient
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import com.google.gson.Gson
import com.google.gson.GsonBuilder
import kotlinx.coroutines.runBlocking
import okhttp3.OkHttpClient
import okhttp3.mockwebserver.MockResponse
import okhttp3.mockwebserver.MockWebServer
import org.junit.After
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import java.util.concurrent.TimeUnit


object RetrofitTestClient {
    // private val BASE_URL = "https://test1.icytree-5b429d30.eastus.azurecontainerapps.io"
    private const val BASE_URL = "/"



}

@RunWith(AndroidJUnit4::class)
class LoginInstrumentedTest {
    @get:Rule
    val composeTestRule = createAndroidComposeRule<MainActivity>()
    private val res: Resources = InstrumentationRegistry.getInstrumentation().targetContext.resources
    private val server = MockWebServer()
    private lateinit var api: Retrofit

    private fun removeToken() {
        val dataStore = InstrumentationRegistry.getInstrumentation().targetContext.dataStore
        val authServ = AuthService(dataStore)
        runBlocking {
            authServ.deleteToken()
        }
    }

    @Before
    fun setup() {
        server.start(8080)
        val okHttpClient: OkHttpClient = OkHttpClient().newBuilder()
            .connectTimeout(60, TimeUnit.SECONDS)
            .readTimeout(60, TimeUnit.SECONDS)
            .writeTimeout(60, TimeUnit.SECONDS)
            .build()
        val retrofit: Retrofit by lazy {
            Retrofit
                .Builder()
                .baseUrl(server.url("/"))
                .client(okHttpClient)
                .addConverterFactory(GsonConverterFactory.create())
                .build()
        }
        retrofit.create(ApiService::class.java)
        /*
        val dataStore = InstrumentationRegistry.getInstrumentation().targetContext.dataStore
        val authServ = AuthService(dataStore)
        try {
            runBlocking {
                authServ.getToken()
                composeTestRule.onNodeWithTag("loggedTopBarImage").assertIsDisplayed().performClick()
                Thread.sleep(5000)
            }
        } catch (e: Exception) {
            return
        }
        */
    }
    @After
    fun after() {
        server.shutdown()
    }
    /*
    @Test
    fun useAppContext() {
        val appContext = InstrumentationRegistry.getInstrumentation().targetContext
        assertEquals("com.example.immotep", appContext.packageName)
        this.removeToken()
    }

    @Test
    fun hasTheHeader() {
        composeTestRule.onNodeWithTag("header").assertIsDisplayed()
    }

    @Test
    fun loginTextDisplayed() {
        composeTestRule
            .onNodeWithText(res.getString(R.string.login_hello))
            .assertIsDisplayed()
        composeTestRule.onNodeWithText(res.getString(R.string.login_details)).assertIsDisplayed()
    }

    @Test
    fun canChangeViewToRegister() {
        composeTestRule.onNodeWithTag("loginScreenToRegisterButton").assertIsDisplayed().performClick()
        composeTestRule.onNodeWithTag("registerScreen").assertIsDisplayed()
    }

    @Test
    fun passwordAndToggleExists() {
        composeTestRule.onNodeWithTag("loginPasswordInput").assertIsDisplayed()
        composeTestRule.onNodeWithTag("togglePasswordVisibility").assertIsDisplayed()
    }

    @Test
    fun emailAndKeepSignedExists() {
        composeTestRule.onNodeWithTag("loginEmailInput").assertIsDisplayed()
        composeTestRule.onNodeWithTag("keepSignedCheckbox").assertIsDisplayed()
    }

    @Test
    fun canTogglePasswordVisibility() {
        composeTestRule.onNodeWithTag("loginPasswordInput").performClick().performTextInput("test123")
        composeTestRule.onNodeWithText("test123").assertDoesNotExist()
        composeTestRule.onNodeWithTag("togglePasswordVisibility").performClick()
        composeTestRule.onNodeWithText("test123").assertIsDisplayed()
    }

    @Test
    fun canSetValueToLoginScreenInputs() {
        composeTestRule.onNodeWithTag("loginEmailInput").performClick().performTextInput("robin.denni@epitech.eu")
        composeTestRule.onNodeWithTag("loginPasswordInput").performClick().performTextInput("test99")
        composeTestRule.onNodeWithTag("togglePasswordVisibility").performClick()
        composeTestRule.onNodeWithTag("loginEmailInput").assert(hasText("robin.denni@epitech.eu", ignoreCase = true))
        composeTestRule.onNodeWithTag("loginPasswordInput").assert(hasText("test99", ignoreCase = true))
    }

    @Test
    fun handlesErrorOnInvalidEmail() {
        composeTestRule.onNodeWithTag("loginEmailInput").performClick().performTextInput("robin.denni")
        composeTestRule.onNodeWithTag("loginPasswordInput").performClick().performTextInput("test99")
        composeTestRule.onNodeWithTag("loginButton").performClick()
        composeTestRule.onNodeWithText(res.getString(R.string.email_error)).assertIsDisplayed()
    }

    @Test
    fun handlesErrorOnInvalidPassword() {
        composeTestRule.onNodeWithTag("loginEmailInput").performClick().performTextInput("test@gmail.com")
        composeTestRule.onNodeWithTag("loginPasswordInput").performClick().performTextInput("te")
        composeTestRule.onNodeWithTag("loginButton").performClick()
        composeTestRule.onNodeWithText(res.getString(R.string.password_error)).assertIsDisplayed()
    }
    */
    /* for this test you need to be connected to the internet, to have a server running and to register a user with the right email and password */
    @Test
    fun canGoToDashboard() {
        val responseBody = LoginResponse(
            access_token = "test",
            refresh_token = "test",
            token_type = "access",
            expires_in = 100000,
            properties = mapOf("test" to "test")
        )
        server.enqueue(MockResponse().setResponseCode(200).setBody(Gson().toJson(responseBody)))
        //this.removeToken()
        composeTestRule.onNodeWithTag("loginEmailInput").performClick().performTextInput("robin.denni@epitech.eu")
        composeTestRule.onNodeWithTag("loginPasswordInput").performClick().performTextInput("Ttest99&")
        composeTestRule.onNodeWithTag("loginButton").performClick()
        Thread.sleep(10000)
        composeTestRule.onNodeWithTag("dashboardScreen").assertIsDisplayed()
        //this.removeToken()
    }
    /*
    @Test
    fun triggersErrorOnUnknownUser() {
        composeTestRule.onNodeWithTag("loginEmailInput").performClick().performTextInput("error@gmail.com")
        composeTestRule.onNodeWithTag("loginPasswordInput").performClick().performTextInput("testError")
        composeTestRule.onNodeWithTag("loginButton").performClick()
        Thread.sleep(4000)
        composeTestRule.onNodeWithTag("errorAlert").assertIsDisplayed()
    */

}
