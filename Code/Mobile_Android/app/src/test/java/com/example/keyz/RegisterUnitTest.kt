package fr.keyz


import androidx.navigation.NavController
import fr.keyz.apiClient.ApiService
import fr.keyz.apiClient.mockApi.fakeRegistrationResponse
import fr.keyz.authService.AuthService
import fr.keyz.authService.RegistrationInput
import fr.keyz.login.dataStore
import fr.keyz.register.RegisterFormError
import fr.keyz.register.RegisterViewModel
import io.mockk.coEvery
import io.mockk.coVerify
import io.mockk.every
import io.mockk.mockk
import io.mockk.verify
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.test.StandardTestDispatcher
import kotlinx.coroutines.test.runTest
import kotlinx.coroutines.test.setMain
import org.junit.Assert.assertEquals
import org.junit.Assert.assertTrue
import org.junit.Before
import org.junit.Test

@ExperimentalCoroutinesApi
class RegisterViewModelTest {

    private lateinit var viewModel: RegisterViewModel
    private val navController : NavController = mockk(relaxed = true)
    private val apiService : ApiService = mockk(relaxed = true)
    private val authService : AuthService = mockk(relaxed = true)
    private val testDispatcher = StandardTestDispatcher()

    @Before
    fun setup() {
        Dispatchers.setMain(testDispatcher)
        every { navController.context.dataStore } returns mockk(relaxed = true)
        every { navController.context } returns mockk(relaxed = true)
        viewModel = RegisterViewModel(navController, apiService)
        viewModel.javaClass.getDeclaredField("authService").apply {
            isAccessible = true
            set(viewModel, authService)
        }
    }

    @Test
    fun `setLastName updates lastName in registerForm`() = runTest {
        viewModel.setLastName("Smith")
        assertEquals("Smith", viewModel.regForm.first().lastName)
    }

    @Test
    fun `setFirstName updates firstName in registerForm`() = runTest {
        viewModel.setFirstName("John")
        assertEquals("John", viewModel.regForm.first().firstName)
    }

    @Test
    fun `setEmail updates email in registerForm`() = runTest {
        viewModel.setEmail("test@example.com")
        assertEquals("test@example.com", viewModel.regForm.first().email)
    }

    @Test
    fun `setPassword updates password in registerForm`() = runTest {
        viewModel.setPassword("password123")
        assertEquals("password123", viewModel.regForm.first().password)
    }

    @Test
    fun `setConfirmPassword updates password in registerConfirm`() = runTest {
        viewModel.setConfirmPassword("password456")
        assertEquals("password456", viewModel.regConfirm.first().password)
    }

    @Test
    fun `setAgreeToTerms updates agreeToTerms in registerConfirm`() = runTest {
        viewModel.setAgreeToTerms(true)
        assertTrue(viewModel.regConfirm.first().agreeToTerms)
    }

    @Test
    fun `onSubmit with valid form registers and navigates to login`() = runTest {
        coEvery { authService.register(any()) } returns fakeRegistrationResponse

        // Set valid form data
        viewModel.setLastName("Smith")
        viewModel.setFirstName("John")
        viewModel.setEmail("john.smith@example.com")
        viewModel.setPassword("Password12345678&")
        viewModel.setConfirmPassword("Password12345678&")
        viewModel.setAgreeToTerms(true)

        viewModel.onSubmit(navController)

        testDispatcher.scheduler.advanceUntilIdle()

        coVerify { authService.register(any()) }
        verify { navController.navigate("login") }

        // Verify form is reset
        val emptyForm = RegistrationInput("", "", "", "")
        assertEquals(emptyForm.email, viewModel.regForm.first().email)
        assertEquals(emptyForm.password, viewModel.regForm.first().password)
        assertEquals(emptyForm.firstName, viewModel.regForm.first().firstName)
        assertEquals(emptyForm.lastName, viewModel.regForm.first().lastName)
        assertEquals("", viewModel.regConfirm.first().password)
        assertEquals(false, viewModel.regConfirm.first().agreeToTerms)
        assertEquals(false, viewModel.regFormError.first().password)
        assertEquals(false, viewModel.regFormError.first().email)
        assertEquals(false, viewModel.regFormError.first().firstName)
        assertEquals(false, viewModel.regFormError.first().lastName)
        assertTrue(viewModel.regFormError.first().apiError == null)
        assertEquals(false, viewModel.regFormError.first().agreeToTerms)
        assertEquals(false, viewModel.regFormError.first().confirmPassword)
    }

    @Test
    fun `onSubmit with invalid form sets form errors`() = runTest {
        // Set invalid form data
        viewModel.setLastName("Sm") // Too short
        viewModel.setFirstName("Jo") // Too short
        viewModel.setEmail("invalid-email")
        viewModel.setPassword("123") // Too short
        viewModel.setConfirmPassword("456") // Doesn't match password
        viewModel.setAgreeToTerms(false)

        // Mock email validation to fail

        viewModel.onSubmit(navController)


        coVerify(exactly = 0) { authService.register(any()) }
        verify(exactly = 0) { navController.navigate(any<String>()) }

        val expectedErrors = RegisterFormError(
            lastName = true,
            firstName = true,
            email = true,
            password = true,
            confirmPassword = true,
            agreeToTerms = true,
            apiError = null
        )
        assertEquals(expectedErrors.lastName, viewModel.regFormError.first().lastName)
        assertEquals(expectedErrors.firstName, viewModel.regFormError.first().firstName)
        assertEquals(expectedErrors.email, viewModel.regFormError.first().email)
        assertEquals(expectedErrors.password, viewModel.regFormError.first().password)
        assertEquals(expectedErrors.confirmPassword, viewModel.regFormError.first().confirmPassword)
        assertEquals(expectedErrors.agreeToTerms, viewModel.regFormError.first().agreeToTerms)
    }

    @Test
    fun `onSubmit with api error sets apiError`() = runTest {
        coEvery { authService.register(any()) } throws Exception("HTTP 400 Unautorized")

        // Set valid form data
        viewModel.setLastName("Smith")
        viewModel.setFirstName("John")
        viewModel.setEmail("john.smith@example.com")
        viewModel.setPassword("Password12345678&")
        viewModel.setConfirmPassword("Password12345678&")
        viewModel.setAgreeToTerms(true)

        viewModel.onSubmit(navController)

        testDispatcher.scheduler.advanceUntilIdle()

        coVerify { authService.register(any()) }
        verify(exactly = 0) { navController.navigate(any<String>()) }

        assertEquals(400, viewModel.regFormError.first().apiError)

    }
}