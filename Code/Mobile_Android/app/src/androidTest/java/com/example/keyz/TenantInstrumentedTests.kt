package fr.keyz

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
import fr.keyz.apiClient.mockApi.MockedApiService
import fr.keyz.apiClient.mockApi.fakeDamagesArray
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
                if (authServ.isUserOwner()) {
                    mainAct.onNodeWithTag("loggedTopBarImage").performClick()
                    mainAct.waitUntilAtLeastOneExists(hasTestTag("loginEmailInput"))
                    throw Exception("User is not a tenant")
                }
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
    fun startInventoryButtonIsNotPresent() {
        mainAct.onNodeWithTag("startInventory").assertIsNotDisplayed()
    }

    @Test
    fun canGoToProfile() {
        mainAct.onNodeWithTag("loggedBottomBarElement profile").assertIsDisplayed().performClick()
        mainAct.onNodeWithTag("profile").assertIsDisplayed()
    }

    @Test
    fun profileContainsGoodInfos() {
        mainAct.onNodeWithTag("loggedBottomBarElement profile").assertIsDisplayed().performClick()
        mainAct.onNodeWithText("User's informations").assertIsDisplayed()
        mainAct.onNodeWithText("Test").assertIsDisplayed()
        mainAct.onNodeWithText("User").assertIsDisplayed()
        mainAct.onNodeWithText("Language").assertIsDisplayed()
        mainAct.onNodeWithText("Logout").assertIsDisplayed()
    }

    @Test
    fun goodNumberOfDamage() {
        this.canGoToDamagesTab()
        fakeDamagesArray.forEach {
            mainAct.onNodeWithTag("oneDamage ${it.id}").assertIsDisplayed()
        }
    }

    @Test
    fun damageContainsGoodInfos() {
        this.canGoToDamagesTab()
        fakeDamagesArray.forEach {
            mainAct.onNodeWithText(it.room_name).assertIsDisplayed()
        }
        mainAct.onAllNodesWithText("fakeComment").assertCountEquals(2)
        mainAct.onAllNodesWithText("2025/03/09").assertCountEquals(2)
    }

    @Test
    fun buttonToReportAClaimIsPresentAndClickable() {
        this.canGoToDamagesTab()
        mainAct.onNodeWithTag("reportClaimButton").assertIsDisplayed().performClick()
        mainAct.waitUntilAtLeastOneExists(hasTestTag("addDamageModal"), 2000)
        mainAct.onNodeWithTag("addDamageModal").assertIsDisplayed()
    }

    @Test
    fun reportClaimModalContainsAllTheGoodInputs() {
        this.buttonToReportAClaimIsPresentAndClickable()
        mainAct.onNodeWithTag("addDamageCommentInput").assertIsDisplayed()
        mainAct.onNodeWithTag("addDamagePriorityDropDown").assertIsDisplayed()
        mainAct.onNodeWithTag("addDamageRoomDropDown").assertIsDisplayed()
        mainAct.onNodeWithTag("addDamageSubmitButton").assertIsDisplayed()
        mainAct.onNodeWithText(res.getString(R.string.priority)).assertIsDisplayed()
        mainAct.onNodeWithText(res.getString(R.string.room)).assertIsDisplayed()
    }
}
