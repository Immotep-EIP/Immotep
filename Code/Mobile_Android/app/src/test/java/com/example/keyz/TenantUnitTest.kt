package com.example.keyz

import androidx.navigation.NavController
import com.example.keyz.apiCallerServices.ApiCallerServiceException
import com.example.keyz.apiCallerServices.RealPropertyCallerService
import com.example.keyz.apiClient.ApiService
import com.example.keyz.apiClient.mockApi.parisFakeProperty
import com.example.keyz.login.dataStore
import com.example.keyz.realProperty.tenant.RealPropertyTenantViewModel
import io.mockk.coEvery
import io.mockk.every
import io.mockk.mockk
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.test.UnconfinedTestDispatcher
import kotlinx.coroutines.test.runTest
import kotlinx.coroutines.test.setMain
import org.junit.Before
import org.junit.Test

@ExperimentalCoroutinesApi
class TenantUnitTest {
    private val navController: NavController = mockk()
    private val apiService: ApiService = mockk()
    private val apiCaller: RealPropertyCallerService = mockk()
    private lateinit var viewModel: RealPropertyTenantViewModel
    private val testDispatcher = UnconfinedTestDispatcher()

    @Before
    fun setUp() {
        Dispatchers.setMain(testDispatcher)
        every { navController.context.dataStore } returns mockk(relaxed = true)
        every { navController.context } returns mockk(relaxed = true)
        viewModel = RealPropertyTenantViewModel(apiService, navController)

        val apiCallerField = viewModel::class.java.getDeclaredField("apiCaller")
        apiCallerField.isAccessible = true
        apiCallerField.set(viewModel, apiCaller)
    }

    @Test
    fun loadPropertyLoadsWell() = runTest {
        val property = parisFakeProperty.toDetailedProperty()
        coEvery { apiCaller.getPropertyWithDetails() } returns property

        viewModel.loadProperty()

        coEvery { apiCaller.getPropertyWithDetails() }
        assert(viewModel.property.first() == property)
    }

    @Test
    fun loadsPropertyWithAnErrorSetsErrorAndDontSetValue() = runTest {
        coEvery { apiCaller.getPropertyWithDetails() } throws ApiCallerServiceException("400")

        viewModel.loadProperty()

        coEvery { apiCaller.getPropertyWithDetails() }
        assert(viewModel.property.first() == null)
        assert(viewModel.loadingError.first() == 400)
    }
}