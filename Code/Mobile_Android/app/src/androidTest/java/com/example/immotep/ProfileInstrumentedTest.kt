package com.example.immotep

import androidx.compose.ui.test.assert
import androidx.compose.ui.test.assertIsDisplayed
import androidx.compose.ui.test.hasText
import androidx.compose.ui.test.junit4.createAndroidComposeRule
import androidx.compose.ui.test.onNodeWithTag
import androidx.compose.ui.test.performClick
import androidx.compose.ui.test.performTextInput
import androidx.test.ext.junit.runners.AndroidJUnit4
import androidx.test.platform.app.InstrumentationRegistry
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith

@RunWith(AndroidJUnit4::class)
class ProfileInstrumentedTest {
    @get:Rule
    val mainAct = createAndroidComposeRule<MainActivity>()

    @Before
    fun setup() {
        val dataStore = InstrumentationRegistry.getInstrumentation().targetContext.dataStore
        val authServ = AuthService(dataStore)
        try {
            runBlocking {
                authServ.getToken()
                mainAct.onNodeWithTag("loggedTopBarClickableIcon").assertIsDisplayed().performClick()
            }
        } catch (e: Exception) {
            runBlocking {
                mainAct.onNodeWithTag("loginEmailInput").performClick().performTextInput("robin.denni@epitech.eu")
                mainAct.onNodeWithTag("loginPasswordInput").performClick().performTextInput("Ttest99&")
                mainAct.onNodeWithTag("loginButton").performClick()
                Thread.sleep(10000)
                mainAct.onNodeWithTag("loggedTopBarClickableIcon").assertIsDisplayed().performClick()
            }
        }
    }

    @Test
    fun useAppContext() {
        val appContext = InstrumentationRegistry.getInstrumentation().targetContext
        assertEquals("com.example.immotep", appContext.packageName)
    }

    @Test
    fun canGoToProfile() {
        mainAct.onNodeWithTag("profile").assertIsDisplayed()
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
    fun bottomBarIsPresent() {
        mainAct.onNodeWithTag("loggedBottomBar").assertIsDisplayed()
    }

    @Test
    fun lastNameTestFieldIsPresentAndClickable() {
        mainAct.onNodeWithTag("profileLastName").assertIsDisplayed().performClick()
    }

    @Test
    fun firstNameTestFieldIsPresentAndClickable() {
        mainAct.onNodeWithTag("profileFirstName").assertIsDisplayed().performClick()
    }

    @Test
    fun emailTestFieldIsPresentAndClickable() {
        mainAct.onNodeWithTag("profileEmail").assertIsDisplayed().performClick()
    }

    @Test
    fun lastNameTestFieldContainsGoodValue() {
        mainAct.onNodeWithTag("profileLastName").assert(hasText("User"))
    }

    @Test
    fun firstNameTestFieldContainsGoodValue() {
        mainAct.onNodeWithTag("profileFirstName").assert(hasText("Test"))
    }

    @Test
    fun emailTestFieldContainsGoodValue() {
        mainAct.onNodeWithTag("profileEmail").assert(hasText("robin.denni@epitech.eu"))
    }
}
