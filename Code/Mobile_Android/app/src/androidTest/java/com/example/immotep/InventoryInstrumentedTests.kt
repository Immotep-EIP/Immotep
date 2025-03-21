package com.example.immotep

import android.content.res.Resources
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
import com.example.immotep.apiClient.mockApi.MockedApiService
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import kotlinx.coroutines.runBlocking
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith


@RunWith(AndroidJUnit4::class)
class InventoryInstrumentedTests {
    constructor() {
        isTesting = true
    }

    @get:Rule
    val mainAct = createAndroidComposeRule<MainActivity>()
    private val res: Resources =
        InstrumentationRegistry.getInstrumentation().targetContext.resources

    @Before
    fun setup() {
        val dataStore = InstrumentationRegistry.getInstrumentation().targetContext.dataStore
        val authServ = AuthService(dataStore, apiService = MockedApiService())
        try {
            runBlocking {
                authServ.getToken()
                mainAct.onNodeWithTag("loggedBottomBarElement realProperty").assertIsDisplayed()
                    .performClick()
                mainAct.onNodeWithTag("propertyBox parisFakeProperty").assertIsDisplayed().performClick()
                mainAct.onNodeWithTag("startInventory").assertIsDisplayed().performClick()
            }
        } catch (e: Exception) {
            runBlocking {
                mainAct.onNodeWithTag("loginEmailInput").performClick()
                    .performTextInput("robin.denni@epitech.eu")
                mainAct.onNodeWithTag("loginPasswordInput").performClick()
                    .performTextInput("Ttest99&")
                mainAct.onNodeWithTag("loginButton").performClick()
                Thread.sleep(2000)
                mainAct.onNodeWithTag("loggedBottomBarElement realProperty").assertIsDisplayed()
                    .performClick()
            }
        }
    }

    private fun goToEntryInventory() {
        mainAct.onNodeWithTag("entryInventoryButton").assertIsDisplayed().performClick()
    }

    @Test
    fun canGoToInventoryPage() {
        mainAct.onNodeWithTag("inventoryScreen").assertIsDisplayed()
    }

    @Test
    fun inventoryPageContainsAllGoodInfos() {
        mainAct.onNodeWithTag("inventoryScreen").assertIsDisplayed()
        mainAct.onNodeWithTag("exitInventoryButton").assertIsDisplayed()
        mainAct.onNodeWithTag("entryInventoryButton").assertIsDisplayed()
    }

    @Test
    fun inventoryLayoutContainsAllGoodInfos() {
        mainAct.onNodeWithTag("inventoryTopBar").assertIsDisplayed()
        mainAct.onNodeWithTag("inventoryTopBarImage").assertIsDisplayed()
        mainAct.onNodeWithTag("inventoryTopBarText").assertIsDisplayed()
        mainAct.onNodeWithTag("inventoryTopBarCloseIcon").assertIsDisplayed()
        mainAct.onNodeWithTag("inventoryLayout").assertIsDisplayed()
    }

    @ExperimentalTestApi
    @Test
    fun canGoToEntryInventoryPage() {
        mainAct.onNodeWithTag("entryInventoryButton").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("roomsScreen"), timeoutMillis = 2000)
    }

    @ExperimentalTestApi
    @Test
    fun entryInventoryRoomPageContainsAllTheGoodInfos() {
        this.canGoToEntryInventoryPage()
        mainAct.onNodeWithTag("roomsScreen").assertIsDisplayed()
        mainAct.onNodeWithTag("addRoomButton").assertIsDisplayed()
        mainAct.onNodeWithTag("confirmInventoryButton").assertIsDisplayed()
        mainAct.onNodeWithTag("editInventoryButton").assertIsDisplayed()
        mainAct.onNodeWithTag("roomButton testRoom").assertIsDisplayed()
        mainAct.onNodeWithText("testRoomName").assertIsDisplayed()
    }

    @ExperimentalTestApi
    @Test
    fun doesAddRoomModalOpens() {
        this.canGoToEntryInventoryPage()
        mainAct.onNodeWithTag("addRoomButton").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("addRoomModal").assertIsDisplayed()
    }

    @ExperimentalTestApi
    @Test
    fun doesAddRoomModalContainsAllTheGoodInfos() {
        this.canGoToEntryInventoryPage()
        mainAct.onNodeWithTag("addRoomButton").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("addRoomModal").assertIsDisplayed()
        mainAct.onNodeWithTag("roomNameTextField").assertIsDisplayed()
        mainAct.onNodeWithTag("addRoomModalCancel").assertIsDisplayed()
        mainAct.onNodeWithTag("addRoomModalConfirm").assertIsDisplayed()
    }

    @ExperimentalTestApi
    @Test
    fun canCloseAddRoomModal() {
        this.canGoToEntryInventoryPage()
        mainAct.onNodeWithTag("addRoomButton").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("addRoomModalCancel").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("addRoomModal").assertDoesNotExist()
    }

    @ExperimentalTestApi
    @Test
    fun canAddARoom() {
        this.canGoToEntryInventoryPage()
        mainAct.onNodeWithTag("addRoomButton").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("roomNameTextField").assertIsDisplayed().performClick().performTextInput("new Test Room")
        mainAct.onNodeWithTag("addRoomModalConfirm").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("addRoomModal").assertDoesNotExist()
        mainAct.onNodeWithText("new Test Room").assertIsDisplayed()

    }

    @Test
    fun canGoToDetailsPage() {

    }

    @Test
    fun detailsPageContainsAllTheGoodInfos() {

    }

    @Test
    fun canAddANewDetail() {

    }

    @Test
    fun canExitDetailPage() {

    }

    @Test
    fun canGoToOneDetailPage() {

    }

    @Test
    fun oneDetailPageContainsAllGoodInfos() {

    }

    @Test
    fun canFillOneDetailAllInfos() {

    }

    @Test
    fun canMakeAnAISimpleCall() {

    }

    @Test
    fun canMakeAnCompareAiCall() {

    }

    @Test
    fun canExitOneDetail() {

    }

    @Test
    fun canCompleteARoom() {

    }

    @Test
    fun canCompleteAnInventoryReport() {

    }

    @Test
    fun canGoToExitInventory() {

    }

    @Test
    fun canCompleteExitInventoryReport() {

    }

    @Test
    fun canGoBack() {

    }

    @Test
    fun exitInventoryIsBlocked() {

    }

    @Test
    fun exitIn() {

    }
}