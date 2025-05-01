package com.example.immotep


import androidx.activity.result.launch
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import androidx.navigation.set
import com.example.immotep.apiCallerServices.ProfileCallerService
import com.example.immotep.apiCallerServices.ProfileResponse
import com.example.immotep.apiClient.ApiService
import com.example.immotep.apiClient.mockApi.fakeProfileResponse
import com.example.immotep.login.dataStore
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
        every { navController.context.dataStore } returns mockk(relaxed = true)
        every { navController.context } returns mockk(relaxed = true)
        viewModel = ProfileViewModel(navController, apiService)

        val apiCallerField = viewModel::class.java.getDeclaredField("apiCaller")
        apiCallerField.isAccessible = true
        apiCallerField.set(viewModel, apiCaller)
    }

    @Test
    fun `initProfile success updates infos`() = runTest {
        coEvery { apiCaller.getProfile() } returns fakeProfileResponse

        viewModel.initProfile()

        coVerify { apiCaller.getProfile() }

        assertTrue(viewModel.infos.first().email == fakeProfileResponse.email)
        assertTrue(viewModel.infos.first().firstname == fakeProfileResponse.firstname)
        assertTrue(viewModel.infos.first().lastname == fakeProfileResponse.lastname)
        assertTrue(viewModel.infos.first().role == fakeProfileResponse.role)
        assertFalse(viewModel.apiError.first())
    }

    @Test
    fun `initProfile api error sets apiError`() = runTest {
        coEvery { apiCaller.getProfile() } throws Exception()

        viewModel.initProfile()

        coVerify { apiCaller.getProfile() }
        assertTrue(viewModel.apiError.first())
    }

}
