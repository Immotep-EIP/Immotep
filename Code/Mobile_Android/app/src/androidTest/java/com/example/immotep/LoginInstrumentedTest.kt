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
import org.junit.Assert.assertEquals
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith

@RunWith(AndroidJUnit4::class)
class LoginInstrumentedTest {
    @get:Rule
    val composeTestRule = createAndroidComposeRule<MainActivity>()
    private val res: Resources = InstrumentationRegistry.getInstrumentation().targetContext.resources

    @Test
    fun useAppContext() {
        val appContext = InstrumentationRegistry.getInstrumentation().targetContext
        assertEquals("com.example.immotep", appContext.packageName)
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
        composeTestRule.onNodeWithText(res.getString(R.string.create_account)).assertIsDisplayed()
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

    /* for this test you need to be connected to the internet, to have a server running and to register a user with the right email and password */
    @Test
    fun canGoToDashboard() {
        composeTestRule.onNodeWithTag("loginEmailInput").performClick().performTextInput("robin.denni@epitech.eu")
        composeTestRule.onNodeWithTag("loginPasswordInput").performClick().performTextInput("test99")
        composeTestRule.onNodeWithTag("loginButton").performClick()
        Thread.sleep(10000)
        composeTestRule.onNodeWithTag("dashboardScreen").assertIsDisplayed()
    }

    @Test
    fun triggersErrorOnUnknownUser() {
        composeTestRule.onNodeWithTag("loginEmailInput").performClick().performTextInput("error@gmail.com")
        composeTestRule.onNodeWithTag("loginPasswordInput").performClick().performTextInput("testError")
        composeTestRule.onNodeWithTag("loginButton").performClick()
        Thread.sleep(4000)
        composeTestRule.onNodeWithTag("errorAlert").assertIsDisplayed()
    }
}
