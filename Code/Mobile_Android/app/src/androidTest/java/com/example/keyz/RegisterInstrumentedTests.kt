package com.example.keyz

import androidx.compose.ui.test.ExperimentalTestApi
import androidx.compose.ui.test.assertIsDisplayed
import androidx.compose.ui.test.hasTestTag
import androidx.compose.ui.test.junit4.createAndroidComposeRule
import androidx.compose.ui.test.onNodeWithTag
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.performClick
import androidx.compose.ui.test.performTextInput
import androidx.test.ext.junit.runners.AndroidJUnit4
import androidx.test.platform.app.InstrumentationRegistry
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import kotlin.random.Random

@RunWith(AndroidJUnit4::class)
class RegisterInstrumentedTests {
    constructor() {
        isTesting = true
    }
    @get:Rule
    val mainAct = createAndroidComposeRule<MainActivity>()
    private val res = InstrumentationRegistry.getInstrumentation().targetContext.resources

    private fun removeToken() {
        try {
            mainAct.onNodeWithTag("loggedTopBarImage").performClick()
        } catch (e: Throwable) {
            println("Node loggedTopBarImage not found. Skipping click.")
        }
    }

    @Before
    fun init() {
        this.removeToken()
    }

    @Test
    fun useAppContext() {
        val appContext = InstrumentationRegistry.getInstrumentation().targetContext
        assertEquals("com.example.keyz", appContext.packageName)
    }

    @Test
    fun hasTheHeader() {
        this.removeToken()
        mainAct.onNodeWithTag("header").assertIsDisplayed()
    }

    @Test
    fun canGoToRegisterScreen() {
        this.removeToken()
        mainAct.onNodeWithTag("loginScreenToRegisterButton").assertIsDisplayed()
            .performClick()
        mainAct.onNodeWithText(res.getString(R.string.create_account)).assertIsDisplayed()
    }

    @Test
    fun canGoBackToLoginScreen() {
        this.canGoToRegisterScreen()
        mainAct.onNodeWithTag("registerScreenToLoginButton").assertIsDisplayed()
            .performClick()
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
    fun lastNameIsClickable() {
        this.canGoToRegisterScreen()
        mainAct.onNodeWithTag("registerLastName").assertIsDisplayed().performClick()
    }

    @Test
    fun firstNameIsClickable() {
        this.canGoToRegisterScreen()
        mainAct.onNodeWithTag("registerFirstName").assertIsDisplayed().performClick()
    }

    @Test
    fun emailIsClickable() {
        this.canGoToRegisterScreen()
        mainAct.onNodeWithTag("registerEmail").assertIsDisplayed().performClick()
    }

    @Test
    fun passwordIsClickable() {
        this.canGoToRegisterScreen()
        mainAct.onNodeWithTag("registerPassword").assertIsDisplayed().performClick()
    }

    @Test
    fun passwordConfirmIsClickable() {
        this.canGoToRegisterScreen()
        mainAct.onNodeWithTag("registerPasswordConfirm").assertIsDisplayed().performClick()
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

    @ExperimentalTestApi
    @Test
    fun canRegisterUser() {
        this.canGoToRegisterScreen()
        val email = "test${Random.nextInt(0, 10000)}@gmail.com"
        mainAct.onNodeWithTag("registerLastName").performClick()
            .performTextInput("test")
        mainAct.onNodeWithTag("registerFirstName").performClick()
            .performTextInput("android")
        mainAct.onNodeWithTag("registerEmail").performClick().performTextInput(email)
        mainAct.onNodeWithTag("registerPassword").performClick()
            .performTextInput("Ttest123&")
        mainAct.onNodeWithTag("registerPasswordConfirm").performClick()
            .performTextInput("Ttest123&")
        mainAct.onNodeWithTag("registerAgreeToTerm").performClick()
        mainAct.onNodeWithTag("registerButton").performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("loginScreen"), timeoutMillis = 5000)
        mainAct.onNodeWithTag("registerScreen").assertDoesNotExist()
        mainAct.onNodeWithTag("loginScreen").assertIsDisplayed()
    }
}

