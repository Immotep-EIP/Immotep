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
import com.example.immotep.apiClient.mockApi.parisFakeProperty
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
        mainAct.onNodeWithTag("propertyBoxLazyColumn").assertIsDisplayed().performScrollToNode(
            hasTestTag("propertyBox emptyFakeProperty")
        )
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
        mainAct.onAllNodesWithTag("propertyBoxRow").assertCountEquals(3)
    }
    /*
    @Test
    fun allTheTestInfosArePresent() {
        mainAct.onNodeWithText("19 rue de la paix").assertIsDisplayed()
        mainAct.onNodeWithText("parisFake").assertIsDisplayed()
        mainAct.onNodeWithText("1 rue de la companie des indes").assertIsDisplayed()
        mainAct.onNodeWithText("marsFake").assertIsDisplayed()
        mainAct.onNodeWithText("lyonFake").assertIsDisplayed()
    }
     */

    @Test
    fun asGoodCountOfTopLeftElementsAndGoodValues() {
        mainAct.onAllNodesWithTag("topRightPropertyBoxInfo").assertCountEquals(3)
        mainAct.onAllNodesWithText(res.getString(R.string.busy)).assertCountEquals(3)
        mainAct.onNodeWithTag("propertyBoxLazyColumn").assertIsDisplayed().performScrollToNode(
            hasTestTag("propertyBox emptyFakeProperty")
        )
        mainAct.onAllNodesWithText(res.getString(R.string.available)).assertCountEquals(1)
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
        mainAct.onNodeWithText("parisFake").assertIsDisplayed()
        mainAct.onNodeWithText("19 rue de la paix, 75000 Paris, France").assertIsDisplayed()
        mainAct.onNodeWithText("2025/03/09 - 2026/03/09").assertIsDisplayed()
        mainAct.onNodeWithText(res.getString(R.string.busy)).assertIsDisplayed()
    }

    @Test
    fun aboutBoxContainsArea() {
        this.goToDetails()
        mainAct.onNodeWithText("45 m²").assertIsDisplayed()
        mainAct.onNodeWithText("Area").assertIsDisplayed()
    }

    @Test
    fun aboutBoxContainsDeposit() {
        this.goToDetails()
        mainAct.onNodeWithText("2000 €").assertIsDisplayed()
        mainAct.onNodeWithText("Deposit").assertIsDisplayed()
    }

    @Test
    fun aboutBoxContainsRent() {
        this.goToDetails()
        mainAct.onNodeWithText("1500 €").assertIsDisplayed()
        mainAct.onNodeWithText("Rent / Month").assertIsDisplayed()
    }

    @Test
    fun aboutBoxNotContainsAvailable() {
        this.goToDetails()
        mainAct.onAllNodesWithText(res.getString(R.string.available)).assertCountEquals(0)
    }

    @Test
    fun containsAllTheTabs() {
        this.goToDetails()
        mainAct.onNodeWithText("About").assertIsDisplayed()
        mainAct.onNodeWithText("Documents").assertIsDisplayed()
        mainAct.onNodeWithText("Damages").assertIsDisplayed()
    }

    @Test
    fun canGoToDocumentTab() {
        this.goToDetails()
        mainAct.onNodeWithTag("tab 1").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("realPropertyDocumentsTab").assertIsDisplayed()
    }

    @Test
    fun canGoToDamagesTab() {
        this.goToDetails()
        mainAct.onNodeWithTag("tab 2").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("realPropertyDetailsDamagesTab").assertIsDisplayed()
    }

    @Test
    fun canGoBackToDetailsTab() {
        this.goToDetails()
        mainAct.onNodeWithTag("tab 1").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("tab 0").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("realPropertyDetailsAboutTab").assertIsDisplayed()
    }

    @Test
    fun canClickOnDocument() {
        this.canGoToDocumentTab()
        mainAct.onNodeWithTag("OneDocument test").assertIsDisplayed().performClick()
    }

    @ExperimentalTestApi
    @Test
    fun modifyModalContainsAllInfos() {
        this.goToDetails()
        mainAct.onNodeWithTag("moreVertOptions").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("editPropertyBtn"))
        mainAct.onNodeWithTag("editPropertyBtn").assertIsDisplayed().performClick()
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
        mainAct.onNodeWithTag("moreVertOptions").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("editPropertyBtn"))
        mainAct.onNodeWithTag("editPropertyBtn").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("addOrEditName").assertIsDisplayed().performClick().performTextInput("ZZ")
        mainAct.onNodeWithTag("addOrEditScrollContainer").assertIsDisplayed().performScrollToNode(
            hasTestTag("addOrEditSubmit")
        )
        mainAct.onNodeWithTag("addOrEditSubmit").assertIsDisplayed().performClick()
        mainAct.waitUntilDoesNotExist(hasTestTag("addOrEditPropertyModal"), timeoutMillis = 2000)
        mainAct.onNodeWithTag("addOrEditPropertyModal").assertIsNotDisplayed()
        mainAct.onNodeWithText("parisFakeZZ").assertIsDisplayed()
    }

    @ExperimentalTestApi
    @Test
    fun modifyPropertyIsGoodOnRealPropertyView() {
        this.goToDetails()
        mainAct.onNodeWithTag("moreVertOptions").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("editPropertyBtn"))
        mainAct.onNodeWithTag("editPropertyBtn").assertIsDisplayed().performClick()
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

    @ExperimentalTestApi
    @Test
    fun dropDownContainsAllTheOptionsOnEmpty() {
        this.goToDetailsOfEmpty()
        mainAct.onNodeWithTag("moreVertOptions").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("inviteTenantBtn"))
        mainAct.onNodeWithTag("endLeaseBtn").assertIsDisplayed()
        mainAct.onNodeWithTag("cancelInvitationBtn").assertIsDisplayed()
        mainAct.onNodeWithTag("editPropertyBtn").assertIsDisplayed()
        mainAct.onNodeWithTag("deletePropertyBtn").assertIsDisplayed()
    }

    @ExperimentalTestApi
    @Test
    fun canInviteUser() {
        this.goToDetailsOfEmpty()
        mainAct.onNodeWithTag("moreVertOptions").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("inviteTenantBtn"))
        mainAct.onNodeWithTag("inviteTenantBtn").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("tenantEmail").assertIsDisplayed().performClick().performTextInput("newTenant@gmail.com")
        mainAct.onNodeWithTag("sendInvitation").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("realPropertyDetailsScreen"), timeoutMillis = 2000)
        mainAct.onNodeWithText("newTenant@gmail.com").assertIsDisplayed()

    }

    @ExperimentalTestApi
    @Test
    fun canCancelInvite() {
        this.canInviteUser()
        mainAct.onNodeWithTag("moreVertOptions").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("cancelInvitationBtn"))
        mainAct.onNodeWithTag("cancelInvitationBtn").assertIsDisplayed().performClick()
        //add confirm + other btns
    }

    @ExperimentalTestApi
    @Test
    fun inviteUserIsGoodOnRealPropertyView() {
        this.goToDetailsOfEmpty()
        mainAct.onNodeWithTag("moreVertOptions").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("inviteTenantBtn"))
        mainAct.onNodeWithTag("inviteTenantBtn").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("tenantEmail").assertIsDisplayed().performClick().performTextInput("newTenant@gmail.com")
        mainAct.onNodeWithTag("sendInvitation").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("backButton"), timeoutMillis = 2000)
        mainAct.onNodeWithTag("backButton").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("realPropertyScreen"), timeoutMillis = 2000)
        mainAct.onNodeWithTag("propertyBoxLazyColumn").assertIsDisplayed().performScrollToNode(
            hasTestTag("propertyBox emptyFakeProperty")
        )
        mainAct.onNodeWithTag("propertyBox emptyFakeProperty").assertIsDisplayed()
        mainAct.onNodeWithText("Pending").assertIsDisplayed()
    }


    @Test
    fun startInventoryButtonIsPresent() {
        this.goToDetails()
        mainAct.onNodeWithTag("startInventory").assertIsDisplayed()
    }
}

