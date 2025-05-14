package com.example.keyz

import android.content.res.Resources
import androidx.compose.ui.test.ExperimentalTestApi
import androidx.compose.ui.test.assertCountEquals
import androidx.compose.ui.test.assertIsDisplayed
import androidx.compose.ui.test.assertIsNotDisplayed
import androidx.compose.ui.test.hasTestTag
import androidx.compose.ui.test.junit4.createAndroidComposeRule
import androidx.compose.ui.test.onAllNodesWithTag
import androidx.compose.ui.test.onAllNodesWithText
import androidx.compose.ui.test.onNodeWithTag
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.performClick
import androidx.compose.ui.test.performTextInput
import androidx.test.ext.junit.runners.AndroidJUnit4
import androidx.test.platform.app.InstrumentationRegistry
import com.example.keyz.apiClient.mockApi.MockedApiService
import com.example.keyz.authService.AuthService
import com.example.keyz.login.dataStore
import kotlinx.coroutines.runBlocking
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith

@ExperimentalTestApi
@RunWith(AndroidJUnit4::class)
class TenantInstrumentedTests {
    constructor() {
        isTesting = true
    }

    @get:Rule
    val mainAct = createAndroidComposeRule<MainActivity>()
    private val res: Resources = InstrumentationRegistry.getInstrumentation().targetContext.resources

    @Before
    fun setup() {
        val dataStore = InstrumentationRegistry.getInstrumentation().targetContext.dataStore
        val authServ = AuthService(dataStore, apiService = MockedApiService())
        try {
            runBlocking {
                authServ.getToken()
                mainAct.onNodeWithTag("loggedBottomBarElement realProperty").assertIsDisplayed().performClick()
            }
        } catch (e: Exception) {
            runBlocking {
                mainAct.onNodeWithTag("loginEmailInput").performClick().performTextInput("tenant@gmail.com")
                mainAct.onNodeWithTag("loginPasswordInput").performClick().performTextInput("Ttest99&")
                mainAct.onNodeWithTag("loginButton").performClick()
                mainAct.waitUntilAtLeastOneExists(hasTestTag("loggedBottomBarElement realProperty"), 2000)
                mainAct.onNodeWithTag("loggedBottomBarElement realProperty").assertIsDisplayed().performClick()
            }
        }
    }

    @Test
    fun canLoginGoToRealPropertyScreen() {
        mainAct.onNodeWithTag("realPropertyTenant").assertIsDisplayed()
    }

    @Test
    fun containNoProperties() {
        mainAct.onAllNodesWithTag("propertyBoxRow").assertCountEquals(0)
    }

    @Test
    fun allTheInfosInDetailsBoxAreGoodAndDisplayed() {
        mainAct.onNodeWithText("parisFake").assertIsDisplayed()
        mainAct.onNodeWithText("19 rue de la paix, 75000 Paris, France").assertIsDisplayed()
        mainAct.onNodeWithText("2025/03/09 - 2026/03/09").assertIsDisplayed()
        mainAct.onNodeWithText(res.getString(R.string.busy)).assertIsDisplayed()
    }

    @Test
    fun aboutBoxContainsArea() {
        mainAct.onNodeWithText("45 m²").assertIsDisplayed()
        mainAct.onNodeWithText("Area").assertIsDisplayed()
    }

    @Test
    fun aboutBoxContainsDeposit() {
        mainAct.onNodeWithText("2000 €").assertIsDisplayed()
        mainAct.onNodeWithText("Deposit").assertIsDisplayed()
    }

    @Test
    fun aboutBoxContainsRent() {
        mainAct.onNodeWithText("1500 €").assertIsDisplayed()
        mainAct.onNodeWithText("Rent / Month").assertIsDisplayed()
    }

    @Test
    fun aboutBoxNotContainsAvailable() {
        mainAct.onAllNodesWithText(res.getString(R.string.available)).assertCountEquals(0)
    }

    @Test
    fun containsAllTheTabs() {
        mainAct.onNodeWithText("About").assertIsDisplayed()
        mainAct.onNodeWithText("Documents").assertIsDisplayed()
        mainAct.onNodeWithText("Damages").assertIsDisplayed()
    }

    @Test
    fun canGoToDocumentTab() {
        mainAct.onNodeWithTag("tab 1").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("realPropertyDocumentsTab").assertIsDisplayed()
    }

    @Test
    fun canGoToDamagesTab() {
        mainAct.onNodeWithTag("tab 2").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("realPropertyDetailsDamagesTab").assertIsDisplayed()
    }

    @Test
    fun canGoBackToDetailsTab() {
        mainAct.onNodeWithTag("tab 1").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("tab 0").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("realPropertyDetailsAboutTab").assertIsDisplayed()
    }

    @Test
    fun canClickOnDocument() {
        this.canGoToDocumentTab()
        mainAct.onNodeWithTag("OneDocument test").assertIsDisplayed().performClick()
    }

    @Test
    fun moreVertIsNotDisplayed() {
        mainAct.onNodeWithTag("moreVertOptions").assertIsNotDisplayed()
    }

    @Test
    fun canGoToProfile() {
        mainAct.onNodeWithTag("loggedTopBarClickableIcon").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("profile").assertIsDisplayed()
    }

    @Test
    fun profileContainsGoodInfos() {
        mainAct.onNodeWithTag("loggedTopBarClickableIcon").assertIsDisplayed().performClick()
        mainAct.onNodeWithText("User's informations").assertIsDisplayed()
        mainAct.onNodeWithText("Test").assertIsDisplayed()
        mainAct.onNodeWithText("User").assertIsDisplayed()
        mainAct.onNodeWithText("Language").assertIsDisplayed()
        mainAct.onNodeWithText("Logout").assertIsDisplayed()
    }
}