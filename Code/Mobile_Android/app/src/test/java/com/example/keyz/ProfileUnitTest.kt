package fr.keyz


import androidx.navigation.NavController
import fr.keyz.apiCallerServices.ProfileCallerService
import fr.keyz.apiClient.ApiService
import fr.keyz.apiClient.mockApi.fakeProfileResponse
import fr.keyz.login.dataStore
import fr.keyz.profile.ProfileViewModel
import io.mockk.coEvery
import io.mockk.coVerify
import io.mockk.every
import io.mockk.mockk
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.test.UnconfinedTestDispatcher
import kotlinx.coroutines.test.runTest
import kotlinx.coroutines.test.setMain
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
