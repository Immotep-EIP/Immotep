package com.example.keyz


import androidx.navigation.NavController
import com.example.keyz.apiClient.ApiService
import com.example.keyz.authService.AuthService
import com.example.keyz.login.LoginViewModel
import com.example.keyz.login.dataStore
import io.mockk.coEvery
import io.mockk.coVerify
import io.mockk.every
import io.mockk.mockk
import io.mockk.unmockkAll
import io.mockk.verify
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.test.StandardTestDispatcher
import kotlinx.coroutines.test.advanceUntilIdle
import kotlinx.coroutines.test.runTest
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.test.resetMain
import kotlinx.coroutines.test.setMain
import org.junit.After
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertTrue
import org.junit.Before
import org.junit.Test

@OptIn(ExperimentalCoroutinesApi::class)
class LoginViewModelTest {

    private lateinit var viewModel: LoginViewModel
    private val navController : NavController = mockk(relaxed = true)
    private val apiService : ApiService= mockk(relaxed = true)
    private val authService : AuthService = mockk(relaxed = true)
    private val testDispatcher = StandardTestDispatcher()

    @Before
    fun setup() {
        Dispatchers.setMain(testDispatcher)
        every { navController.context.dataStore } returns mockk(relaxed = true)
        every { navController.context } returns mockk(relaxed = true)
        viewModel = LoginViewModel(navController, apiService)
        viewModel.javaClass.getDeclaredField("authService").apply {
            isAccessible = true
            set(viewModel, authService)
        }
    }

    @After
    fun tearDown() {
        Dispatchers.resetMain()
        unmockkAll()
    }


    @Test
    fun isTestingMustBeSetToFalse() = runTest {
        assertFalse(isTesting)
    }

    @Test
    fun `updateEmailAndPassword should update emailAndPassword state`() = runTest {
        // Arrange
        val email = "test@example.com"
        val password = "password123"
        val keepSigned = true

        // Act
        viewModel.updateEmailAndPassword(email, password, keepSigned)

        // Assert
        val state = viewModel.emailAndPassword.first()
        assertEquals(email, state.email)
        assertEquals(password, state.password)
        assertEquals(keepSigned, state.keepSigned)
    }

    @Test
    fun `login should set email error if email is invalid`() = runTest {
        // Arrange
        viewModel.updateEmailAndPassword("invalid-email", "password", false)

        // Act
        viewModel.login({})
        advanceUntilIdle()

        // Assert
        val errorState = viewModel.errors.first()
        assertTrue(errorState.email)
        assertFalse(errorState.password)
        assertEquals(null, errorState.apiError)
    }

    @Test
    fun `login should set password error if password is too short`() = runTest {
        viewModel.updateEmailAndPassword("test@example.com", "12", false)

        viewModel.login({})
        advanceUntilIdle()

        val errorState = viewModel.errors.first()
        assertFalse(errorState.email)
        assertTrue(errorState.password)
        assertEquals(null, errorState.apiError)
    }

    @Test
    fun `login should call AuthService onLogin and navigate to dashboard if credentials are valid`() = runTest {
        val email = "test@example.com"
        val password = "password123"
        viewModel.updateEmailAndPassword(email, password, false)
        coEvery {
            authService.onLogin(username = email, password = password)
        } returns Unit

        viewModel.login({})
        advanceUntilIdle()

        coVerify {
            authService.onLogin(any(), any())
        }
        verify { navController.navigate("dashboard") }
    }

    @Test
    fun `login should set apiError if AuthService onLogin throws an exception`() = runTest {
        val email = "test@example.com"
        val password = "password123"
        val errorMessage = "Login failed,401"
        viewModel.updateEmailAndPassword(email, password, false)
        coEvery {
            authService.onLogin(username = email, password = password)
        } throws Exception(errorMessage)

        viewModel.login({})
        advanceUntilIdle()

        val errorState = viewModel.errors.first()
        assertFalse(errorState.email)
        assertFalse(errorState.password)
        assertEquals(401, errorState.apiError)
    }

    @Test
    fun `login should not call AuthService onLogin if there are errors`() = runTest {
        viewModel.updateEmailAndPassword("invalid-email", "12", false)

        viewModel.login({})
        advanceUntilIdle()

        coVerify(exactly = 0) {
            authService.onLogin(any(), any())
        }
    }
}
