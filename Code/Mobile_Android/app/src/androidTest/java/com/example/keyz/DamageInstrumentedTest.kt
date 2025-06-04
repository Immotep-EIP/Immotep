package com.example.keyz

import android.content.res.Resources
import androidx.compose.ui.test.ExperimentalTestApi
import androidx.compose.ui.test.assertCountEquals
import androidx.compose.ui.test.assertIsDisplayed
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
import com.example.keyz.apiClient.mockApi.fakeDamagesArray
import com.example.keyz.authService.AuthService
import com.example.keyz.login.dataStore
import kotlinx.coroutines.runBlocking
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith

@ExperimentalTestApi
@RunWith(AndroidJUnit4::class)
class DamageInstrumentedTest {
    constructor() {
        isTesting = true
    }

    @get:Rule
    val mainAct = createAndroidComposeRule<MainActivity>()
    private val res: Resources =
        InstrumentationRegistry.getInstrumentation().targetContext.resources

    private fun goToDamageTab() {
        mainAct.waitUntilAtLeastOneExists(
            hasTestTag("loggedBottomBarElement realProperty"),
            2000
        )
        mainAct.onNodeWithTag("loggedBottomBarElement realProperty").assertIsDisplayed()
            .performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("propertyBox parisFakeProperty"), 2000)
        mainAct.onNodeWithTag("propertyBox parisFakeProperty").assertIsDisplayed()
            .performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("tab 2"), 2000)
        mainAct.onNodeWithTag("tab 2").assertIsDisplayed().performClick()
    }

    @Before
    fun setup() {
        val dataStore = InstrumentationRegistry.getInstrumentation().targetContext.dataStore
        val authServ = AuthService(dataStore, apiService = MockedApiService())
        try {
            runBlocking {
                authServ.getToken()
                goToDamageTab()
            }
        } catch (e: Exception) {
            runBlocking {
                mainAct.onNodeWithTag("loginEmailInput").performClick()
                    .performTextInput("robin.denni@epitech.eu")
                mainAct.onNodeWithTag("loginPasswordInput").performClick()
                    .performTextInput("Ttest99&")
                mainAct.onNodeWithTag("loginButton").performClick()
                goToDamageTab()
            }
        }
    }

    @Test
    fun canGoToDamageTab() {
        mainAct.onNodeWithTag("realPropertyDetailsDamagesTab").assertIsDisplayed()
    }

    @Test
    fun goodNumberOfDamage() {
        fakeDamagesArray.forEach {
            mainAct.onNodeWithTag("oneDamage ${it.id}").assertIsDisplayed()
        }
    }

    @Test
    fun damageContainsGoodInfos() {
        fakeDamagesArray.forEach {
            mainAct.onNodeWithText(it.room_name).assertIsDisplayed()
        }
        mainAct.onAllNodesWithText("fakeComment").assertCountEquals(2)
        mainAct.onAllNodesWithText("2025/03/09").assertCountEquals(2)
    }
}
