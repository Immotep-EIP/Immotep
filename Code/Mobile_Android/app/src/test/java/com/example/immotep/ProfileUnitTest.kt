package com.example.immotep


import androidx.activity.result.launch
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import androidx.navigation.set
import com.example.immotep.apiCallerServices.ProfileCallerService
import com.example.immotep.apiCallerServices.ProfileResponse
import com.example.immotep.apiClient.ApiService
import com.example.immotep.apiClient.mockApi.fakeProfileResponse
import com.example.immotep.profile.ProfileState
import com.example.immotep.profile.ProfileViewModel
import io.mockk.Runs
import io.mockk.coEvery
import io.mockk.coVerify
import io.mockk.every
import io.mockk.just
import io.mockk.mockk
import io.mockk.verify
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.test.UnconfinedTestDispatcher
import kotlinx.coroutines.test.runTest
import kotlinx.coroutines.test.setMain
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertTrue
import org.junit.Before
import org.junit.Test
import kotlin.jvm.java

@ExperimentalCoroutinesApi
class ProfileViewModelTest {

    private val navController: NavController = mockk()
    private val apiService: ApiService = mockk()
    private val apiCaller: ProfileCallerService = mockk()
    private lateinit var viewModel: ProfileViewModel
    private val testDispatcher = UnconfinedTestDispatcher()

    @Before
    fun setUp() {
        Dispatchers.setMain(testDispatcher)
        viewModel = ProfileViewModel(navController, apiService)

        // Mock the private apiCaller
        val apiCallerField = viewModel::class.java.getDeclaredField("apiCaller")
        apiCallerField.isAccessible = true
        apiCallerField.set(viewModel, apiCaller)
    }

    @Test
    fun `initProfile success updates infos`() = runTest {
        coEvery { apiCaller.getProfile(any()) } returns fakeProfileResponse

        viewModel.initProfile()

        coVerify { apiCaller.getProfile(any()) }

        assertTrue(viewModel.infos.first().email == fakeProfileResponse.email)
        assertTrue(viewModel.infos.first().firstname == fakeProfileResponse.firstname)
        assertTrue(viewModel.infos.first().lastname == fakeProfileResponse.lastname)
        assertTrue(viewModel.infos.first().role == fakeProfileResponse.role)
        assertFalse(viewModel.apiError.first())
    }

    @Test
    fun `initProfile api error sets apiError`() = runTest {
        coEvery { apiCaller.getProfile(any()) } throws Exception()

        viewModel.initProfile()

        coVerify { apiCaller.getProfile(any()) }
        assertTrue(viewModel.apiError.first())
    }

    @Test
    fun `setEmail updates email in infos`() = runTest {
        viewModel.setEmail("newemail@example.com")

        assertEquals("newemail@example.com", viewModel.infos.first().email)
    }

    @Test
    fun `setFirstName updates firstname in infos`() = runTest {
        viewModel.setFirstName("Jane")

        assertEquals("Jane", viewModel.infos.first().firstname)
    }

    @Test
    fun `setLastName updates lastname in infos`() = runTest {
        viewModel.setLastName("Smith")

        assertEquals("Smith", viewModel.infos.first().lastname)
    }

    @Test
    fun `updateProfile success calls apiCaller updateProfile`() = runTest {
        coEvery { apiCaller.updateProfile(any(), any()) } returns Unit
        val initialState = ProfileState(
            email = "test@example.com",
            firstname = "John",
            lastname = "Doe",
            role = "user"
        )
        viewModel.setEmail(initialState.email)
        viewModel.setFirstName(initialState.firstname)
        viewModel.setLastName(initialState.lastname)

        viewModel.updateProfile()

        coVerify { apiCaller.updateProfile(initialState.toProfileUpdateInput(), any()) }
        assertFalse(viewModel.apiError.first())
        assertFalse(viewModel.isLoading.first())
    }

    @Test
    fun `updateProfile api error sets apiError`() = runTest {
        coEvery { apiCaller.updateProfile(any(), any()) } throws Exception()

        viewModel.updateProfile()

        coVerify { apiCaller.updateProfile(any(), any()) }
        assertTrue(viewModel.apiError.first())
        assertFalse(viewModel.isLoading.first())
    }
}