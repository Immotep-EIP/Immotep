package com.example.immotep

import androidx.compose.ui.test.assertIsDisplayed
import androidx.compose.ui.test.junit4.createAndroidComposeRule
import androidx.compose.ui.test.onNodeWithTag
import androidx.compose.ui.test.performClick
import androidx.compose.ui.test.performTextInput
import androidx.test.ext.junit.runners.AndroidJUnit4
import androidx.test.platform.app.InstrumentationRegistry
import com.example.immotep.apiClient.mockApi.MockedApiService
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith

/*
@RunWith(AndroidJUnit4::class)
class DashBoardInstrumentedTest {
    constructor() {
        isTesting = true
    }
    @get:Rule
    val mainAct = createAndroidComposeRule<MainActivity>()

    @Before
    fun setup() {
        val dataStore = InstrumentationRegistry.getInstrumentation().targetContext.dataStore
        val authServ = AuthService(dataStore, apiService = MockedApiService())
        try {
            runBlocking {
                authServ.getToken()
            }
        } catch (e: Exception) {
            runBlocking {
                mainAct.onNodeWithTag("loginEmailInput").performClick().performTextInput("robin.denni@epitech.eu")
                mainAct.onNodeWithTag("loginPasswordInput").performClick().performTextInput("Ttest99&")
                mainAct.onNodeWithTag("loginButton").performClick()
                Thread.sleep(10000)
            }
        }
    }

    @Test
    fun useAppContext() {
        val appContext = InstrumentationRegistry.getInstrumentation().targetContext
        assertEquals("com.example.immotep", appContext.packageName)
    }

    @Test
    fun canGoToDashboard() {
        mainAct.onNodeWithTag("dashboardScreen").assertIsDisplayed()
    }

    @Test
    fun layoutIsPresent() {
        mainAct.onNodeWithTag("dashboardLayout").assertIsDisplayed()
    }

    @Test
    fun topBarIsPresent() {
        mainAct.onNodeWithTag("loggedTopBar").assertIsDisplayed()
    }

    @Test
    fun headerContainsImage() {
        mainAct.onNodeWithTag("loggedTopBarImage").assertIsDisplayed()
    }

    @Test
    fun headerContainsText() {
        mainAct.onNodeWithTag("loggedTopBarText").assertIsDisplayed()
    }

    @Test
    fun headerContainsClickableIcons() {
        mainAct.onNodeWithTag("loggedTopBarClickableIcon").assertIsDisplayed()
    }

    @Test
    fun canDisconnect() {
        mainAct.onNodeWithTag("loggedTopBarImage").performClick()
        mainAct.onNodeWithTag("loginEmailInput").assertIsDisplayed()
    }

    @Test
    fun bottomBarIsPresent() {
        mainAct.onNodeWithTag("loggedBottomBar").assertIsDisplayed()
    }

    @Test
    fun bottomBarContainsAllElements() {
        mainAct.onNodeWithTag("loggedBottomBarElement dashboard").assertIsDisplayed()
        mainAct.onNodeWithTag("loggedBottomBarElement realProperty").assertIsDisplayed()
        mainAct.onNodeWithTag("loggedBottomBarElement messages").assertIsDisplayed()
        mainAct.onNodeWithTag("loggedBottomBarElement settings").assertIsDisplayed()
    }

    @Test
    fun canGoToRealProperty() {
        mainAct.onNodeWithTag("loggedBottomBarElement realProperty").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("realPropertyScreen").assertIsDisplayed()
    }

    @Test
    fun canGoToRealPropertyAndGoBackToDashboard() {
        mainAct.onNodeWithTag("loggedBottomBarElement realProperty").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("loggedBottomBarElement dashboard").performClick()
        mainAct.onNodeWithTag("dashboardScreen").assertIsDisplayed()
    }

    @Test
    fun stayOnDashBoardIfDashboardIsClicked() {
        mainAct.onNodeWithTag("loggedBottomBarElement dashboard").performClick()
        mainAct.onNodeWithTag("dashboardScreen").assertIsDisplayed()
    }
}
*/