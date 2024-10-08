package com.example.immotep

import android.content.res.Resources
import androidx.compose.ui.test.assertIsDisplayed
import androidx.compose.ui.test.junit4.createAndroidComposeRule
import androidx.compose.ui.test.onNodeWithTag
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.performClick
import androidx.test.ext.junit.runners.AndroidJUnit4
import androidx.test.platform.app.InstrumentationRegistry
import org.junit.Assert.assertEquals
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith

@RunWith(AndroidJUnit4::class)
class RegisterInstrumentedTests {
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
    fun canGoToRegisterScreen() {
        composeTestRule.onNodeWithTag("loginScreenToRegisterButton").assertIsDisplayed().performClick()
        composeTestRule.onNodeWithText(res.getString(R.string.create_account)).assertIsDisplayed()
    }

    @Test
    fun canGoBackToLoginScreen() {
        this.canGoToRegisterScreen()
        composeTestRule.onNodeWithTag("registerScreenToLoginButton").assertIsDisplayed().performClick()
        composeTestRule.onNodeWithText(res.getString(R.string.login_hello)).assertIsDisplayed()
    }
}
