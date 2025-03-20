package com.example.immotep

import android.content.res.Resources
import androidx.compose.ui.semantics.setSelection
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
import androidx.compose.ui.test.performScrollToNode
import androidx.compose.ui.test.performTextInput
import androidx.navigation.activity
import androidx.test.ext.junit.runners.AndroidJUnit4
import androidx.test.platform.app.InstrumentationRegistry
import com.example.immotep.apiClient.mockApi.MockedApiService
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import java.time.LocalDate
import java.time.ZoneOffset


@RunWith(AndroidJUnit4::class)
class RealPropertyInstrumentedTest {
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
                mainAct.onNodeWithTag("loginEmailInput").performClick().performTextInput("robin.denni@epitech.eu")
                mainAct.onNodeWithTag("loginPasswordInput").performClick().performTextInput("Ttest99&")
                mainAct.onNodeWithTag("loginButton").performClick()
                Thread.sleep(2000)
                mainAct.onNodeWithTag("loggedBottomBarElement realProperty").assertIsDisplayed().performClick()
            }
        }
    }

    private fun goToDetails() {
        mainAct.onNodeWithTag("propertyBox parisFakeProperty").assertIsDisplayed().performClick()
    }

    private fun goToDetailsOfEmpty() {
        mainAct.onNodeWithTag("propertyBox emptyFakeProperty").assertIsDisplayed().performClick()
    }
    @Test
    fun useAppContext() {
        val appContext = InstrumentationRegistry.getInstrumentation().targetContext
        assertEquals("com.example.immotep", appContext.packageName)
    }

    @Test
    fun canGoToRealPropertyScreen() {
        mainAct.onNodeWithTag("realPropertyScreen").assertIsDisplayed()
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
    fun headerContainsImage() {
        mainAct.onNodeWithTag("loggedTopBarImage").assertIsDisplayed()
    }

    @Test
    fun addPropertyIsPresent() {
        mainAct.onNodeWithTag("addAPropertyBtn").assertIsDisplayed()
    }

    @Test
    fun containAllTheTestProperties() {
        mainAct.onAllNodesWithTag("propertyBoxRow").assertCountEquals(4)
    }

    @Test
    fun allTheTestInfosArePresent() {
        mainAct.onNodeWithText("19 rue de la paix").assertIsDisplayed()
        mainAct.onNodeWithText("test@gmail.com").assertIsDisplayed()
        mainAct.onNodeWithText("1 rue de la companie des indes").assertIsDisplayed()
        mainAct.onNodeWithText("crashbandicoot@gmail.com").assertIsDisplayed()
        mainAct.onNodeWithText("30 rue de la source").assertIsDisplayed()
        mainAct.onNodeWithText("tomnook@gmail.com").assertIsDisplayed()
        mainAct.onNodeWithText("30 rue du test").assertIsDisplayed()

    }

    @Test
    fun asGoodCountOfTopLeftElementsAndGoodValues() {
        mainAct.onAllNodesWithTag("topRightPropertyBoxInfo").assertCountEquals(4)
        mainAct.onAllNodesWithText(res.getString(R.string.available)).assertCountEquals(1)
        mainAct.onAllNodesWithText(res.getString(R.string.busy)).assertCountEquals(3)
    }

    @Test
    fun canClickOnProperty() {
        this.goToDetails()
        mainAct.onNodeWithTag("realPropertyDetailsScreen").assertIsDisplayed()
    }

    @Test
    fun canClickOnEmptyProperty() {
        this.goToDetailsOfEmpty()
        mainAct.onNodeWithTag("realPropertyDetailsScreen").assertIsDisplayed()
    }

    @Test
    fun canGoBack() {
        this.goToDetails()
        mainAct.onNodeWithTag("backButton").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("realPropertyScreen").assertIsDisplayed()
    }

    @Test
    fun allTheInfosInDetailsBoxAreGoodAndDisplayed() {
        this.goToDetails()
        mainAct.onNodeWithText("19 rue de la paix").assertIsDisplayed()
        mainAct.onAllNodesWithText("test@gmail.com").assertCountEquals(2)
        mainAct.onAllNodesWithText(res.getString(R.string.busy)).assertCountEquals(1)
    }

    @Test
    fun aboutBoxIsPresentAndContainsGoodInfos() {
        this.goToDetails()
        mainAct.onNodeWithText(res.getString(R.string.about_the_property)).assertIsDisplayed()
    }

    @Test
    fun aboutBoxContainsArea() {
        this.goToDetails()
        mainAct.onNodeWithText("Area: 45 m²").assertIsDisplayed()
    }

    @Test
    fun aboutBoxContainsDeposit() {
        this.goToDetails()
        mainAct.onNodeWithText("Deposit: 2000€").assertIsDisplayed()
    }

    @Test
    fun aboutBoxNotContainsAvailable() {
        this.goToDetails()
        mainAct.onAllNodesWithText(res.getString(R.string.available)).assertCountEquals(0)
    }

    @Test
    fun detailsContainsTheDocuments() {
        this.goToDetails()
        mainAct.onNodeWithTag("OneDocument test").assertIsDisplayed()
        mainAct.onNodeWithText("testDocName").assertIsDisplayed()
    }

    @Test
    fun canClickOnDocument() {
        this.goToDetails()
        //mainAct.onNodeWithTag("OneDocument test").assertIsDisplayed().performClick()
    }

    @Test
    fun modifyModalButtonIsPresent() {
        this.goToDetails()
        mainAct.onNodeWithTag("editProperty").assertIsDisplayed()
    }

    @Test
    fun modifyModalDoesOpen() {
        this.goToDetails()
        mainAct.onNodeWithTag("editProperty").assertIsDisplayed().performClick()
    }

    @ExperimentalTestApi
    @Test
    fun modifyModalContainsAllInfos() {
        this.goToDetails()
        mainAct.onNodeWithTag("editProperty").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("addOrEditDeposit"), timeoutMillis = 2000)
        mainAct.onNodeWithTag("addOrEditScrollContainer").assertIsDisplayed().performScrollToNode(
            hasTestTag("addOrEditDeposit")
        )
        mainAct.onNodeWithTag("addOrEditDeposit").assertIsDisplayed()
        mainAct.onNodeWithTag("addOrEditArea").assertIsDisplayed()
        mainAct.onNodeWithTag("addOrEditRental").assertIsDisplayed()
        mainAct.onNodeWithTag("addOrEditName").assertIsDisplayed()
        mainAct.onNodeWithTag("addOrEditAddress").assertIsDisplayed()
        mainAct.onNodeWithTag("addOrEditNumber").assertIsDisplayed()
        mainAct.onNodeWithTag("addOrEditCity").assertIsDisplayed()
        mainAct.onNodeWithTag("addOrEditPostalCode").assertIsDisplayed()
        mainAct.onNodeWithTag("addOrEditCountry").assertIsDisplayed()
        mainAct.onNodeWithTag("addOrEditScrollContainer").assertIsDisplayed().performScrollToNode(
            hasTestTag("addOrEditSubmit")
        )
        mainAct.onNodeWithTag("addOrEditSubmit").assertIsDisplayed()
    }

    @ExperimentalTestApi
    @Test
    fun canModifyProperty() {
        this.goToDetails()
        mainAct.onNodeWithTag("editProperty").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("addOrEditAddress").assertIsDisplayed().performClick().performTextInput("ZZ")
        mainAct.onNodeWithTag("addOrEditScrollContainer").assertIsDisplayed().performScrollToNode(
            hasTestTag("addOrEditSubmit")
        )
        mainAct.onNodeWithTag("addOrEditSubmit").assertIsDisplayed().performClick()
        mainAct.waitUntilDoesNotExist(hasTestTag("addOrEditPropertyModal"), timeoutMillis = 2000)
        mainAct.onNodeWithTag("addOrEditPropertyModal").assertIsNotDisplayed()
        mainAct.onNodeWithText("19 rue de la paixZZ").assertIsDisplayed()
    }

    @ExperimentalTestApi
    @Test
    fun modifyPropertyIsGoodOnRealPropertyView() {
        this.goToDetails()
        mainAct.onNodeWithTag("editProperty").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("addOrEditAddress").assertIsDisplayed().performClick().performTextInput("ZZ")
        mainAct.onNodeWithTag("addOrEditScrollContainer").assertIsDisplayed().performScrollToNode(
            hasTestTag("addOrEditSubmit")
        )
        mainAct.onNodeWithTag("addOrEditSubmit").assertIsDisplayed().performClick()
        mainAct.waitUntilDoesNotExist(hasTestTag("addOrEditPropertyModal"), timeoutMillis = 2000)
        mainAct.onNodeWithTag("addOrEditPropertyModal").assertIsNotDisplayed()
        mainAct.onNodeWithTag("backButton").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("realPropertyScreen"), timeoutMillis = 2000)
        mainAct.onNodeWithText("19 rue de la paixZZ").assertIsDisplayed()
    }

    @Test
    fun inviteUserButtonIsPresent() {
        this.goToDetailsOfEmpty()
        mainAct.onNodeWithTag("inviteTenantBtn").assertIsDisplayed()
    }

    @Test
    fun inviteUserModalDoesOpen() {
        this.goToDetailsOfEmpty()
        mainAct.onNodeWithTag("inviteTenantBtn").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("inviteTenantModal").assertIsDisplayed()
    }
    /*
    @Test
    fun inviteUserModalDoesContainsTheGoodInputsAndButton() {
        this.goToDetailsOfEmpty()
        mainAct.onNodeWithTag("inviteTenantBtn").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("tenantEmail").assertIsDisplayed()
        mainAct.onNodeWithTag( "textField startDateInput").assertIsDisplayed()
        mainAct.onNodeWithTag("textField endDateInput").assertIsDisplayed()
    }

    @Test
    fun canInviteUser() {
        this.goToDetailsOfEmpty()
        mainAct.onNodeWithTag("inviteTenantBtn").assertIsDisplayed().performClick()

    }

    @Test
    fun inviteUserIsGoodOnRealPropertyView() {
        this.goToDetailsOfEmpty()
    }
    */

    @Test
    fun startInventoryButtonIsPresent() {
        this.goToDetails()
        mainAct.onNodeWithTag("startInventory").assertIsDisplayed()
    }
}
