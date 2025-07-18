package fr.keyz

import androidx.compose.ui.test.ExperimentalTestApi
import androidx.compose.ui.test.assertIsDisplayed
import androidx.compose.ui.test.hasTestTag
import androidx.compose.ui.test.junit4.createAndroidComposeRule
import androidx.compose.ui.test.onNodeWithTag
import androidx.compose.ui.test.performClick
import androidx.compose.ui.test.performTextInput
import androidx.test.ext.junit.runners.AndroidJUnit4
import androidx.test.platform.app.InstrumentationRegistry
import fr.keyz.apiClient.mockApi.MockedApiService
import fr.keyz.authService.AuthService
import fr.keyz.MainActivity
import fr.keyz.isTesting
import fr.keyz.login.dataStore
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith


@ExperimentalTestApi
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
                mainAct.waitUntilAtLeastOneExists(hasTestTag("dashboardScreen"), 2000)
            }
        }
    }

    @Test
    fun useAppContext() {
        val appContext = InstrumentationRegistry.getInstrumentation().targetContext
        assertEquals("fr.keyz", appContext.packageName)
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
    fun canDisconnect() {
        mainAct.onNodeWithTag("loggedTopBarImage").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("loginEmailInput"), 2000)
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
        mainAct.onNodeWithTag("loggedBottomBarElement profile").assertIsDisplayed()
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
