package fr.keyz

import android.content.res.Resources
import androidx.compose.ui.test.ExperimentalTestApi
import androidx.compose.ui.test.assertIsDisplayed
import androidx.compose.ui.test.assertIsNotDisplayed
import androidx.compose.ui.test.hasTestTag
import androidx.compose.ui.test.hasText
import androidx.compose.ui.test.junit4.createAndroidComposeRule
import androidx.compose.ui.test.onNodeWithTag
import androidx.compose.ui.test.onNodeWithText
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
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith


@ExperimentalTestApi
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
                mainAct.waitUntilAtLeastOneExists(hasTestTag("loggedBottomBarElement realProperty"), 2000)
                mainAct.onNodeWithTag("loggedBottomBarElement realProperty").assertIsDisplayed()
                    .performClick()
                mainAct.onNodeWithTag("propertyBox parisFakeProperty").assertIsDisplayed().performClick()
                mainAct.onNodeWithTag("startInventory").assertIsDisplayed().performClick()
            }
        }
    }

    @Test
    fun canGoToInventoryPage() {
        mainAct.onNodeWithTag("inventoryScreen").assertIsDisplayed()
    }

    @Test
    fun inventoryPageContainsAllGoodInfos() {
        mainAct.onNodeWithTag("inventoryScreen").assertIsDisplayed()
    }

    @Test
    fun inventoryLayoutContainsAllGoodInfos() {
        mainAct.onNodeWithTag("inventoryTopBar").assertIsDisplayed()
        mainAct.onNodeWithTag("inventoryTopBarImage").assertIsDisplayed()
        mainAct.onNodeWithTag("inventoryTopBarText").assertIsDisplayed()
        mainAct.onNodeWithTag("inventoryTopBarCloseIcon").assertIsDisplayed()
        mainAct.onNodeWithTag("inventoryLayout").assertIsDisplayed()
    }

    @Test
    fun isRoomsScreenDisplayed() {
        mainAct.waitUntilAtLeastOneExists(hasTestTag("roomsScreen"), timeoutMillis = 2000)
    }

    @Test
    fun inventoryRoomPageContainsAllTheGoodInfos() {
        mainAct.onNodeWithTag("roomsScreen").assertIsDisplayed()
        mainAct.onNodeWithTag("confirmInventoryButton").assertIsDisplayed()
        mainAct.onNodeWithTag("editInventoryButton").assertIsDisplayed()
        mainAct.onNodeWithTag("roomButton testRoom").assertIsDisplayed()
        mainAct.onNodeWithText("BedRoom").assertIsDisplayed()
    }

    @Test
    fun doesAddRoomModalOpens() {
        this.canTriggerModifyForRooms()
        mainAct.onNodeWithTag("addRoomButton").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("addRoomOrDetailModal"), 2000)
    }

    @Test
    fun doesAddRoomModalContainsAllTheGoodInfos() {
        this.canTriggerModifyForRooms()
        mainAct.onNodeWithTag("addRoomButton").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("addRoomOrDetailModal"), 2000)
        mainAct.onNodeWithTag("roomNameTextField").assertIsDisplayed()
        mainAct.onNodeWithTag("addRoomModalCancel").assertIsDisplayed()
        mainAct.onNodeWithTag("addRoomModalConfirm").assertIsDisplayed()
    }

    @Test
    fun canCloseAddRoomModal() {
        this.canTriggerModifyForRooms()
        mainAct.onNodeWithTag("addRoomButton").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("addRoomModalCancel").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("addRoomModal").assertDoesNotExist()
    }

    @Test
    fun canTriggerModifyForRooms() {
        mainAct.onNodeWithTag("editInventoryButton").assertIsDisplayed().performClick()
    }

    @Test
    fun canAddARoom() {
        this.canTriggerModifyForRooms()
        mainAct.onNodeWithTag("addRoomButton").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("roomNameTextField").assertIsDisplayed().performClick().performTextInput("new Test Room")
        mainAct.onNodeWithTag("addRoomModalConfirm").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("addRoomModal").assertDoesNotExist()
        mainAct.onNodeWithText("new Test Room").assertIsDisplayed()
    }

    @Test
    fun canGoToDetailsPage() {
        mainAct.onNodeWithTag("roomButton testRoom").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("roomsDetailsScreen"), timeoutMillis = 2000)
    }

    @Test
    fun detailsPageContainsAllTheGoodInfos() {
        this.canGoToDetailsPage()
        mainAct.onNodeWithTag("editRoomsDetails").assertIsDisplayed()
        mainAct.onNodeWithTag("detailButton testFurniture").assertIsDisplayed()
    }

    @Test
    fun canTriggerModifyForDetails() {
        this.canGoToDetailsPage()
        mainAct.onNodeWithTag("editRoomsDetails").assertIsDisplayed().performClick()
    }

    @Test
    fun canAddANewDetail() {
        this.canTriggerModifyForDetails()
        mainAct.onNodeWithTag("addDetailsButton").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("roomNameTextField").assertIsDisplayed().performClick().performTextInput("new Test Detail")
        mainAct.onNodeWithTag("addRoomModalConfirm").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("addRoomModal").assertDoesNotExist()
        mainAct.onNodeWithText("new Test Detail").assertIsDisplayed()
    }


    @Test
    fun canExitDetailPage() {
        this.canGoToDetailsPage()
        mainAct.onNodeWithTag("inventoryTopBarCloseIcon").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("roomsDetailsScreen").assertIsNotDisplayed()
        mainAct.onNodeWithTag("roomsScreen").assertIsDisplayed()
    }

    @Test
    fun canGoToOneDetailPage() {
        this.canGoToDetailsPage()
        mainAct.onNodeWithTag("detailButton testFurniture").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("oneDetailScreen").assertIsDisplayed()
    }

    @Test
    fun oneDetailPageContainsAllGoodInfos() {
        this.canGoToOneDetailPage()
        mainAct.onNodeWithTag("addingPicturesCarousel").assertIsDisplayed()
        mainAct.onNodeWithTag("validateButton").assertIsDisplayed()
        mainAct.onNodeWithTag("aiCallButton").assertIsDisplayed()
        mainAct.onNodeWithTag("dropDownState").assertIsDisplayed()
        mainAct.onNodeWithTag("dropDownCleanliness").assertIsDisplayed()
        mainAct.onNodeWithTag("oneDetailComment").assertIsDisplayed()
    }
    /*
    @Test
    fun canFillOneDetailAllInfos() {
        this.canGoToOneDetailPage()
        mainAct.onAllNodesWithTag("addingPicturesCarousel").assertCountEquals(2)
        mainAct.onNodeWithTag("validateButton").assertIsDisplayed()
        mainAct.onNodeWithTag("aiCallButton").assertIsDisplayed()
        mainAct.onNodeWithTag("dropDownState").assertIsDisplayed()
        mainAct.onNodeWithTag("oneDetailComment").assertIsDisplayed()
        mainAct.onNodeWithTag("inventoryTopBarCloseIcon").assertIsDisplayed()
    }

     */

    @Test
    fun canExitOneDetail() {
        this.canGoToOneDetailPage()
        mainAct.onNodeWithTag("inventoryTopBarCloseIcon").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("roomsDetailsScreen"), timeoutMillis = 2000)
    }

    @Test
    fun canGoBack() {
        mainAct.onNodeWithTag("inventoryTopBarCloseIcon").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasText("Exit"), timeoutMillis = 2000)
        mainAct.onNodeWithText("Exit").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("realPropertyDetailsScreen"), timeoutMillis = 2000)
    }

}
