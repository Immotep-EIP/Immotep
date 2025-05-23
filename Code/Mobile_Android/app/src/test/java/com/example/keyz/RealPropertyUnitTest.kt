package com.example.keyz

import androidx.navigation.NavController
import com.example.keyz.apiCallerServices.AddPropertyInput
import com.example.keyz.apiCallerServices.DetailedProperty
import com.example.keyz.apiCallerServices.RealPropertyCallerService
import com.example.keyz.apiClient.ApiService
import com.example.keyz.apiClient.CreateOrUpdateResponse
import com.example.keyz.apiClient.mockApi.fakeProperties
import com.example.keyz.login.dataStore
import com.example.keyz.realProperty.RealPropertyViewModel
import io.mockk.Runs
import io.mockk.coEvery
import io.mockk.coVerify
import io.mockk.every
import io.mockk.just
import io.mockk.mockk
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.test.UnconfinedTestDispatcher
import kotlinx.coroutines.test.runTest
import kotlinx.coroutines.test.setMain
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Before
import org.junit.Test
import kotlin.jvm.java

@ExperimentalCoroutinesApi
class RealPropertyViewModelTest {

    private val navController: NavController = mockk()
    private val apiService: ApiService = mockk()
    private val apiCaller: RealPropertyCallerService = mockk()
    private lateinit var viewModel: RealPropertyViewModel
    private val testDispatcher = UnconfinedTestDispatcher()

    private val property1 = DetailedProperty(id = "1", name = "Property 1")
    private val property2 = DetailedProperty(id = "2", name = "Property 2")
    private val addPropertyInput = AddPropertyInput(name = "New Property")

    @Before
    fun setUp() {
        Dispatchers.setMain(testDispatcher)
        every { navController.context.dataStore } returns mockk(relaxed = true)
        every { navController.context } returns mockk(relaxed = true)
        viewModel = RealPropertyViewModel(navController, apiService)

        val apiCallerField = viewModel::class.java.getDeclaredField("apiCaller")
        apiCallerField.isAccessible = true
        apiCallerField.set(viewModel, apiCaller)
    }

    @Test
    fun `closeError sets apiError to NONE`() = runTest {
        viewModel.closeError()
        assertEquals(RealPropertyViewModel.WhichApiError.NONE, viewModel.apiError.first())
    }

    @Test
    fun `getProperties success updates properties and sets isLoading`() = runTest {
        val propertiesList = fakeProperties.map {
            it.toDetailedProperty()
        }.toTypedArray()
        coEvery { apiCaller.getPropertiesAsDetailedProperties() } returns propertiesList

        viewModel.getProperties()

        coVerify { apiCaller.getPropertiesAsDetailedProperties() }
        assertEquals(propertiesList.size, viewModel.properties.size)
        assertEquals(propertiesList.first().id, viewModel.properties.first().id)
        assertFalse(viewModel.isLoading.first())
        assertEquals(RealPropertyViewModel.WhichApiError.NONE, viewModel.apiError.first())
    }

    @Test
    fun `getProperties api error sets apiError and isLoading`() = runTest {
        coEvery { apiCaller.getPropertiesAsDetailedProperties() } throws Exception()

        viewModel.getProperties()

        coVerify { apiCaller.getPropertiesAsDetailedProperties() }
        assertTrue(viewModel.apiError.first() == RealPropertyViewModel.WhichApiError.GET_PROPERTIES)
        assertFalse(viewModel.isLoading.first())
    }

    @Test
    fun `addProperty success adds property and clears error`() = runTest {
        coEvery { apiCaller.addProperty(addPropertyInput) } returns CreateOrUpdateResponse(
            id = property1.id
        )

        viewModel.addProperty(addPropertyInput)

        coVerify { apiCaller.addProperty(addPropertyInput) }
        assertEquals(1, viewModel.properties.size)
        assertEquals(property1.id, viewModel.properties.first().id)
        assertEquals(RealPropertyViewModel.WhichApiError.NONE, viewModel.apiError.first())
    }

    @Test
    fun `addProperty api error sets apiError`() = runTest {
        coEvery { apiCaller.addProperty(any()) } throws Exception()

        viewModel.addProperty(addPropertyInput)

        coVerify { apiCaller.addProperty(any()) }
        assertTrue(viewModel.apiError.first() == RealPropertyViewModel.WhichApiError.ADD_PROPERTY)
    }

    @Test
    fun `deleteProperty success removes property and clears error`() = runTest {
        viewModel.properties.addAll(listOf(property1, property2))
        coEvery { apiCaller.archiveProperty("1") } just Runs

        viewModel.deleteProperty("1")

        coVerify { apiCaller.archiveProperty("1") }
        assertEquals(listOf(property2), viewModel.properties)
        assertEquals(RealPropertyViewModel.WhichApiError.NONE, viewModel.apiError.first())
    }

    @Test
    fun `deleteProperty api error sets apiError`() = runTest {
        viewModel.properties.add(property1)
        coEvery { apiCaller.archiveProperty("1") } throws Exception()

        viewModel.deleteProperty("1")

        coVerify { apiCaller.archiveProperty("1") }
        assertTrue(viewModel.apiError.first() == RealPropertyViewModel.WhichApiError.DELETE_PROPERTY)
    }

    @Test
    fun `deleteProperty with non-existent id does nothing`() = runTest {
        viewModel.properties.addAll(listOf(property1, property2))

        viewModel.deleteProperty("999")

        coVerify(exactly = 0) { apiCaller.archiveProperty(any()) }
        assertEquals(listOf(property1, property2), viewModel.properties)
    }

    @Test
    fun `setPropertySelectedDetails sets propertySelectedDetails`() = runTest {
        viewModel.properties.addAll(listOf(property1, property2))

        viewModel.setPropertySelectedDetails("2")

        assertEquals(property2, viewModel.propertySelectedDetails.first())
    }

    @Test
    fun `setPropertySelectedDetails with non-existent id does nothing`() = runTest {
        viewModel.properties.addAll(listOf(property1, property2))

        viewModel.setPropertySelectedDetails("999")

        assertNull(viewModel.propertySelectedDetails.first())
    }

    @Test
    fun `getBackFromDetails updates property and clears selected details`() = runTest {
        viewModel.properties.addAll(listOf(property1, property2))
        val modifiedProperty = property1.copy(name = "Modified Property 1")

        viewModel.getBackFromDetails(modifiedProperty)

        assertEquals(modifiedProperty, viewModel.properties.first { it.id == "1" })
        assertNull(viewModel.propertySelectedDetails.first())
    }

    @Test
    fun `getBackFromDetails with non-existent id does nothing`() = runTest {
        viewModel.properties.addAll(listOf(property1, property2))
        val nonExistentProperty = DetailedProperty(id = "999", name = "Non-existent")

        viewModel.getBackFromDetails(nonExistentProperty)

        assertEquals(listOf(property1, property2), viewModel.properties)
    }
}
