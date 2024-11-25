package com.example.immotep

import android.content.res.Resources
import androidx.compose.ui.test.assertCountEquals
import androidx.compose.ui.test.assertIsDisplayed
import androidx.compose.ui.test.junit4.createAndroidComposeRule
import androidx.compose.ui.test.onAllNodesWithTag
import androidx.compose.ui.test.onAllNodesWithText
import androidx.compose.ui.test.onNodeWithTag
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.performClick
import androidx.compose.ui.test.performTextInput
import androidx.test.ext.junit.runners.AndroidJUnit4
import androidx.test.platform.app.InstrumentationRegistry
import com.example.immotep.AuthService.AuthService
import com.example.immotep.login.dataStore
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith

@RunWith(AndroidJUnit4::class)
class RealPropertyInstrumentedTest {
    @get:Rule
    val mainAct = createAndroidComposeRule<MainActivity>()
    private val res: Resources = InstrumentationRegistry.getInstrumentation().targetContext.resources

    @Before
    fun setup() {
        val dataStore = InstrumentationRegistry.getInstrumentation().targetContext.dataStore
        val authServ = AuthService(dataStore)
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
                Thread.sleep(10000)
                mainAct.onNodeWithTag("loggedBottomBarElement realProperty").assertIsDisplayed().performClick()
            }
        }
    }

    private fun goToDetails() {
        mainAct.onNodeWithTag("propertyBox last").assertIsDisplayed().performClick()
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

    @Test
    fun allTheTestInfosArePresent() {
        mainAct.onNodeWithText("19 rue de la paix, Paris 75000").assertIsDisplayed()
        mainAct.onNodeWithText("John Doe").assertIsDisplayed()
        mainAct.onNodeWithText("1 rue de la companie des indes, Marseille 13000").assertIsDisplayed()
        mainAct.onNodeWithText("Crash Bandicoot").assertIsDisplayed()
        mainAct.onNodeWithText("30 rue de la source, Lyon 69000").assertIsDisplayed()
        mainAct.onNodeWithText("Tom Nook").assertIsDisplayed()
    }

    @Test
    fun asGoodCountOfTopLeftElementsAndGoodValues() {
        mainAct.onAllNodesWithTag("topRightPropertyBoxInfo").assertCountEquals(3)
        mainAct.onAllNodesWithText(res.getString(R.string.available)).assertCountEquals(1)
        mainAct.onAllNodesWithText(res.getString(R.string.busy)).assertCountEquals(2)
    }

    @Test
    fun canClickOnProperty() {
        this.goToDetails()
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
        mainAct.onNodeWithText("19 rue de la paix, Paris 75000").assertIsDisplayed()
        mainAct.onAllNodesWithText("John Doe").assertCountEquals(2)
        mainAct.onAllNodesWithText(res.getString(R.string.available)).assertCountEquals(1)
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
    fun aboutBoxContainsAvailable() {
        this.goToDetails()
        mainAct.onAllNodesWithText(res.getString(R.string.available)).assertCountEquals(1)
    }

    @Test
    fun documentBoxIsPresentAndContainsGoodInfos() {
        this.goToDetails()
        mainAct.onAllNodesWithTag("OneDocument").assertCountEquals(4)
        mainAct.onNodeWithText("july quittance").assertIsDisplayed()
        mainAct.onNodeWithText("old inventory").assertIsDisplayed()
        mainAct.onNodeWithText("oven invoice").assertIsDisplayed()
        mainAct.onNodeWithText("august quittance").assertIsDisplayed()
    }

    @Test
    fun startInventoryButtonIsPresent() {
        this.goToDetails()
        mainAct.onNodeWithTag("startInventory").assertIsDisplayed()
    }
}
