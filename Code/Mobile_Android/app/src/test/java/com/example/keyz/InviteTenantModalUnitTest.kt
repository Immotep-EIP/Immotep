package com.example.keyz


import androidx.navigation.NavController
import com.example.keyz.apiCallerServices.InviteTenantCallerService
import com.example.keyz.apiClient.ApiService
import com.example.keyz.apiClient.mockApi.fakeInviteOutputValue
import com.example.keyz.components.inviteTenantModal.InviteTenantViewModel
import io.mockk.coEvery
import io.mockk.coVerify
import io.mockk.every
import io.mockk.mockk
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.test.StandardTestDispatcher
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
import java.util.Date

@OptIn(ExperimentalCoroutinesApi::class)
class InviteTenantViewModelTest {

    private lateinit var viewModel: InviteTenantViewModel
    private val apiService: ApiService = mockk(relaxed = true)
    private val navController: NavController = mockk(relaxed = true)
    private val callerService: InviteTenantCallerService = mockk(relaxed = true)
    private val testDispatcher = StandardTestDispatcher()

    @Before
    fun setup() {
        Dispatchers.setMain(testDispatcher)
        every { navController.context } returns mockk(relaxed = true)
        viewModel = InviteTenantViewModel(apiService, navController)
        viewModel.javaClass.getDeclaredField("callerService").apply {
            isAccessible = true
            set(viewModel, callerService)
        }
    }

    @After
    fun tearDown() {
        Dispatchers.resetMain()
    }

    @Test
    fun `reset should clear invitationForm and invitationFormError`() = runTest {
        viewModel.setEmail("test@example.com")
        viewModel.setStartDate(Date().time + 1000)
        viewModel.setEndDate(Date().time + 2000)

        viewModel.reset()

        val form = viewModel.invitationForm.first()
        val formError = viewModel.invitationFormError.first()
        assertEquals("", form.email)
        val testDate = Date().time
        assertTrue(testDate == form.startDate || testDate + 1 == form.startDate || testDate - 1 == form.startDate)
        assertTrue(testDate - 1000 == form.endDate || testDate - 999 == form.endDate || testDate - 1001 == form.endDate)
        assertFalse(formError.email)
        assertFalse(formError.date)
    }

    @Test
    fun `setStartDate should update invitationForm`() = runTest {
        val newStartDate = Date().time + 1000

        viewModel.setStartDate(newStartDate)

        val form = viewModel.invitationForm.first()
        assertEquals(newStartDate, form.startDate)
    }

    @Test
    fun `setEndDate should update invitationForm`() = runTest {
        val newEndDate = Date().time + 2000

        viewModel.setEndDate(newEndDate)

        val form = viewModel.invitationForm.first()
        assertEquals(newEndDate, form.endDate)
    }

    @Test
    fun `setEmail should update invitationForm`() = runTest {
        val newEmail = "test@example.com"

        viewModel.setEmail(newEmail)

        val form = viewModel.invitationForm.first()
        assertEquals(newEmail, form.email)
    }

    @Test
    fun `inviteTenantValidator should return false if email is invalid`() = runTest {
        viewModel.setEmail("invalid-email")
        viewModel.inviteTenant(
            close = {},
            propertyId = "1",
            onError = {},
            onSubmit = { _, _, _ -> },
            setIsLoading = {}
        )
        val formError = viewModel.invitationFormError.first()
        assertTrue(formError.email)
    }

    @Test
    fun `inviteTenantValidator should return false if startDate is after endDate`() = runTest {
        val startDate = Date().time + 2000
        val endDate = Date().time + 1000

        viewModel.setEmail("test@example.com")
        viewModel.setStartDate(startDate)
        viewModel.setEndDate(endDate)

        viewModel.inviteTenant(
            close = {},
            propertyId = "1",
            onError = {},
            onSubmit = { _, _, _ -> },
            setIsLoading = {}
        )
        val formError = viewModel.invitationFormError.first()
        assertTrue(formError.date)
    }

    @Test
    fun `inviteTenantValidator should return true if email and dates are valid`() = runTest {
        val startDate = Date().time + 1000
        val endDate = Date().time + 2000
        viewModel.setEmail("test@example.com")
        viewModel.setStartDate(startDate)
        viewModel.setEndDate(endDate)


        viewModel.inviteTenant(
            close = {},
            propertyId = "1",
            onError = {},
            onSubmit = { _, _, _ -> },
            setIsLoading = {}
        )

        val formError = viewModel.invitationFormError.first()
        assertFalse(formError.date)
        assertFalse(formError.email)
    }

    @Test
    fun `inviteTenant should call callerService invite and onSubmit with correct data`() = runTest {
        coEvery { callerService.invite(any(), any()) } returns fakeInviteOutputValue
        val now = Date().time
        val startDate = now + 1000
        val endDate = now + 2000
        val email = "test@example.com"
        val propertyId = "1"
        viewModel.setEmail(email)
        viewModel.setStartDate(startDate)
        viewModel.setEndDate(endDate)
        var submittedEmail: String? = null
        var submittedStartDate: Long? = null
        var submittedEndDate: Long? = null
        val onSubmit: (String, Long, Long) -> Unit = { e, s, ed ->
            submittedEmail = e
            submittedStartDate = s
            submittedEndDate = ed
        }
        val close: () -> Unit = {}
        val onError: () -> Unit = {}

        viewModel.inviteTenant(
            close = close,
            propertyId = propertyId,
            onError = onError,
            onSubmit = onSubmit,
            setIsLoading = {}
        )
        testDispatcher.scheduler.advanceUntilIdle()
        coVerify { callerService.invite(any(), any()) }
        assertTrue(email == submittedEmail)
        assertTrue(startDate == submittedStartDate)
        assertTrue(endDate == submittedEndDate)
    }

    @Test
    fun `inviteTenant should call close and reset`() = runTest {
        coEvery { callerService.invite(any(), any()) } returns fakeInviteOutputValue
        val time = Date().time
        val startDate = time + 1000
        val endDate = time + 2000
        val email = "test@example.com"
        val propertyId = "1"
        viewModel.setEmail(email)
        viewModel.setStartDate(startDate)
        viewModel.setEndDate(endDate)
        var closeCalled = false
        val close: () -> Unit = { closeCalled = true }
        val onError: () -> Unit = {}
        val onSubmit: (String, Long, Long) -> Unit = { _, _, _ -> }

        viewModel.inviteTenant(close, propertyId, onError, onSubmit, {})
        testDispatcher.scheduler.advanceUntilIdle()

        assertTrue(closeCalled)
        val form = viewModel.invitationForm.first()
        val formError = viewModel.invitationFormError.first()
        assertEquals("", form.email)
        assertTrue(startDate < form.startDate + 1000)
        assertTrue(endDate < form.endDate + 3000)
        assertFalse(formError.email)
        assertFalse(formError.date)
    }
}
