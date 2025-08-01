package fr.keyz

import android.content.res.Resources
import androidx.compose.ui.test.ExperimentalTestApi
import androidx.compose.ui.test.assert
import androidx.compose.ui.test.assertIsDisplayed
import androidx.compose.ui.test.hasTestTag
import androidx.compose.ui.test.hasText
import androidx.compose.ui.test.junit4.createAndroidComposeRule
import androidx.compose.ui.test.onNodeWithTag
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.performClick
import androidx.compose.ui.test.performTextInput
import androidx.test.ext.junit.runners.AndroidJUnit4
import androidx.test.platform.app.InstrumentationRegistry
import fr.keyz.MainActivity
import fr.keyz.isTesting
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith

@ExperimentalTestApi
@RunWith(AndroidJUnit4::class)
class LoginInstrumentedTest {
    constructor() {
        isTesting = true
    }
    @get:Rule
    val composeTestRule = createAndroidComposeRule<MainActivity>()

    private val res: Resources = InstrumentationRegistry.getInstrumentation().targetContext.resources

    private fun removeToken() {
        try {
            composeTestRule.onNodeWithTag("loggedTopBarImage").performClick()
        } catch (e: Throwable) {
            println("Node loggedTopBarImage not found. Skipping click.")
        }
    }

    @Before
    fun init() {
        this.removeToken()
    }

    @Test
    fun hasTheHeader() {
        this.removeToken()
        composeTestRule.onNodeWithTag("header").assertIsDisplayed()
    }

    @Test
    fun loginTextDisplayed() {
        this.removeToken()

        composeTestRule
            .onNodeWithText(res.getString(R.string.login_hello))
            .assertIsDisplayed()
        composeTestRule.onNodeWithText(res.getString(R.string.login_details)).assertIsDisplayed()
    }

    @Test
    fun canChangeViewToRegister() {
        this.removeToken()

        composeTestRule.onNodeWithTag("loginScreenToRegisterButton").assertIsDisplayed().performClick()
        composeTestRule.onNodeWithTag("registerScreen").assertIsDisplayed()
    }

    @Test
    fun passwordAndToggleExists() {
        this.removeToken()

        composeTestRule.onNodeWithTag("loginPasswordInput").assertIsDisplayed()
        composeTestRule.onNodeWithTag("togglePasswordVisibility").assertIsDisplayed()
    }

    @Test
    fun emailAndKeepSignedExists() {
        this.removeToken()

        composeTestRule.onNodeWithTag("loginEmailInput").assertIsDisplayed()
        composeTestRule.onNodeWithTag("keepSignedCheckbox").assertIsDisplayed()
    }

    @Test
    fun canTogglePasswordVisibility() {
        this.removeToken()

        composeTestRule.onNodeWithTag("loginPasswordInput").performClick().performTextInput("test123")
        composeTestRule.onNodeWithText("test123").assertDoesNotExist()
        composeTestRule.onNodeWithTag("togglePasswordVisibility").performClick()
        composeTestRule.onNodeWithText("test123").assertIsDisplayed()
    }

    @Test
    fun canSetValueToLoginScreenInputs() {
        this.removeToken()

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
        this.removeToken()

        composeTestRule.onNodeWithTag("loginEmailInput").performClick().performTextInput("test@gmail.com")
        composeTestRule.onNodeWithTag("loginPasswordInput").performClick().performTextInput("te")
        composeTestRule.onNodeWithTag("loginButton").performClick()
        composeTestRule.onNodeWithText(res.getString(R.string.password_error)).assertIsDisplayed()
    }

    @Test
    fun canGoToDashboard() {
        this.removeToken()
        composeTestRule.onNodeWithTag("loginEmailInput").performClick().performTextInput("robin.denni@epitech.eu")
        composeTestRule.onNodeWithTag("loginPasswordInput").performClick().performTextInput("Ttest99&")
        composeTestRule.onNodeWithTag("loginButton").performClick()
        composeTestRule.waitUntilAtLeastOneExists(hasTestTag("dashboardScreen"), 2000)
        composeTestRule.onNodeWithTag("dashboardScreen").assertIsDisplayed()
    }

    @Test
    fun triggersErrorOnUnknownUser() {
        this.removeToken()
        composeTestRule.onNodeWithTag("loginEmailInput").performClick()
            .performTextInput("error@gmail.com")
        composeTestRule.onNodeWithTag("loginPasswordInput").performClick()
            .performTextInput("testError")
        composeTestRule.onNodeWithTag("loginButton").performClick()
        composeTestRule.onNodeWithTag("errorAlert").assertIsDisplayed()
    }

}
