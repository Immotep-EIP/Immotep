package com.example.immotep

import android.content.res.Resources
import androidx.compose.ui.test.assertIsDisplayed
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
import kotlin.random.Random

@RunWith(AndroidJUnit4::class)
class RegisterInstrumentedTests {
    @get:Rule
    val mainAct = createAndroidComposeRule<MainActivity>()
    private val res: Resources = InstrumentationRegistry.getInstrumentation().targetContext.resources

    @Test
    fun useAppContext() {
        val appContext = InstrumentationRegistry.getInstrumentation().targetContext
        assertEquals("com.example.immotep", appContext.packageName)
    }

    @Test
    fun hasTheHeader() {
        mainAct.onNodeWithTag("header").assertIsDisplayed()
    }

    @Test
    fun canGoToRegisterScreen() {
        mainAct.onNodeWithTag("loginScreenToRegisterButton").assertIsDisplayed().performClick()
        mainAct.onNodeWithText(res.getString(R.string.create_account)).assertIsDisplayed()
    }

    @Test
    fun canGoBackToLoginScreen() {
        this.canGoToRegisterScreen()
        mainAct.onNodeWithTag("registerScreenToLoginButton").assertIsDisplayed().performClick()
        mainAct.onNodeWithText(res.getString(R.string.login_hello)).assertIsDisplayed()
    }

    @Test
    fun hasAllTheRequiredFields() {
        this.canGoToRegisterScreen()
        mainAct.onNodeWithTag("registerLastName").assertIsDisplayed()
        mainAct.onNodeWithTag("registerFirstName").assertIsDisplayed()
        mainAct.onNodeWithTag("registerEmail").assertIsDisplayed()
        mainAct.onNodeWithTag("registerPassword").assertIsDisplayed()
        mainAct.onNodeWithTag("registerPasswordConfirm").assertIsDisplayed()
        mainAct.onNodeWithTag("registerAgreeToTerm").assertIsDisplayed()
    }

    @Test
    fun hasAllTheRequiredButtons() {
        this.canGoToRegisterScreen()
        mainAct.onNodeWithTag("registerButton").assertIsDisplayed()
        mainAct.onNodeWithTag("registerScreenToLoginButton").assertIsDisplayed()
    }

    @Test
    fun errorWhenInfosIsNotPresent() {
        this.canGoToRegisterScreen()
        mainAct.onNodeWithTag("registerButton").performClick()
        mainAct.onNodeWithText(res.getString(R.string.agree_terms_error)).assertIsDisplayed()
        mainAct.onNodeWithText(res.getString(R.string.first_name_error)).assertIsDisplayed()
        mainAct.onNodeWithText(res.getString(R.string.last_name_error)).assertIsDisplayed()
        mainAct.onNodeWithText(res.getString(R.string.email_error)).assertIsDisplayed()
        mainAct.onNodeWithText(res.getString(R.string.register_password_error)).assertIsDisplayed()
        mainAct.onNodeWithText(res.getString(R.string.password_confirm_error)).assertIsDisplayed()
    }

    /* for this test you need to have a server running */
    @Test
    fun canRegisterUser() {
        this.canGoToRegisterScreen()
        mainAct.onNodeWithTag("registerLastName").performClick().performTextInput("test")
        mainAct.onNodeWithTag("registerFirstName").performClick().performTextInput("android")
        mainAct.onNodeWithTag("registerEmail").performClick().performTextInput("test${Random.nextInt(0, 10000)}@gmail.com")
        mainAct.onNodeWithTag("registerPassword").performClick().performTextInput("test123&")
        mainAct.onNodeWithTag("registerPasswordConfirm").performClick().performTextInput("test123&")
        mainAct.onNodeWithTag("registerAgreeToTerm").performClick()
        mainAct.onNodeWithTag("registerButton").performClick()
        Thread.sleep(10000)
        mainAct.onNodeWithText(res.getString(R.string.login_hello)).assertIsDisplayed()
    }
}
