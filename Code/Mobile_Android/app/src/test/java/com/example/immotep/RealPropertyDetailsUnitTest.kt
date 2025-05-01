package com.example.immotep


import com.example.immotep.apiCallerServices.AddPropertyInput
import com.example.immotep.apiCallerServices.DetailedProperty
import com.example.immotep.apiCallerServices.PropertyStatus
import com.example.immotep.apiCallerServices.RealPropertyCallerService
import com.example.immotep.apiClient.ApiService
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.Document
import com.example.immotep.apiCallerServices.InviteDetailedProperty
import com.example.immotep.apiCallerServices.LeaseDetailedProperty
import com.example.immotep.apiClient.CreateOrUpdateResponse
import com.example.immotep.apiClient.mockApi.baseDateStr
import com.example.immotep.apiClient.mockApi.fakeDocument
import com.example.immotep.login.dataStore
import com.example.immotep.realProperty.details.RealPropertyDetailsViewModel
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
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Test
import java.time.Instant
import java.time.OffsetDateTime
import java.time.ZoneOffset

@ExperimentalCoroutinesApi
class RealPropertyDetailsViewModelTest {

    private val navController: NavController = mockk()
    private val apiService: ApiService = mockk()
    private val apiCaller: RealPropertyCallerService = mockk()
    private lateinit var viewModel: RealPropertyDetailsViewModel
    private val testDispatcher = UnconfinedTestDispatcher()

    private val property1 = DetailedProperty(id = "1", name = "Property 1")
    private val addPropertyInput = AddPropertyInput(name = "Updated Property")
    private val document1 = Document(
        id = "doc1",
        name = "Document 1",
        data = "base64Data",
        created_at = baseDateStr
    )

    @Before
    fun setUp() {
        Dispatchers.setMain(testDispatcher)
        every { navController.context.dataStore } returns mockk(relaxed = true)
        every { navController.context } returns mockk(relaxed = true)
        viewModel = RealPropertyDetailsViewModel(navController, apiService)

        val apiCallerField = viewModel::class.java.getDeclaredField("apiCaller")
        apiCallerField.isAccessible = true
        apiCallerField.set(viewModel, apiCaller)
    }

    @Test
    fun `loadProperty success updates property with documents and sets isLoading`() = runTest {
        coEvery { apiCaller.getPropertyDocuments("1", any()) } returns arrayOf(fakeDocument)

        viewModel.loadProperty(property1.copy(lease = LeaseDetailedProperty(
            id = "1",
            startDate = OffsetDateTime.now(),
            endDate = OffsetDateTime.now().plusMonths(1),
            tenantEmail = "tenant@example.com",
            tenantName = "John Doe"
        )))

        coVerify { apiCaller.getPropertyDocuments("1", any()) }
        assertEquals(fakeDocument.id, viewModel.documents.first().id)
        assertEquals(fakeDocument.data, viewModel.documents.first().data)
        assertEquals(fakeDocument.name, viewModel.documents.first().name)
        assertEquals(fakeDocument.created_at, viewModel.documents.first().created_at)
        assertEquals(RealPropertyDetailsViewModel.ApiErrors.NONE, viewModel.apiError.first())
    }

    @Test
    fun `loadProperty api error does not crash and sets isLoading`() = runTest {
        coEvery { apiCaller.getPropertyDocuments("1", any()) } throws Exception("API Error")

        viewModel.loadProperty(property1.copy(lease = LeaseDetailedProperty(
            id = "1",
            startDate = OffsetDateTime.now(),
            endDate = OffsetDateTime.now().plusMonths(1),
            tenantEmail = "tenant@example.com",
            tenantName = "John Doe"
        )))

        coVerify { apiCaller.getPropertyDocuments("1", any()) }
        assertEquals(property1.id, viewModel.property.first().id)
        assertEquals(RealPropertyDetailsViewModel.ApiErrors.NONE, viewModel.apiError.first())
        assertEquals(false, viewModel.isLoading.first())
    }

    @Test
    fun `editProperty success updates property and clears error`() = runTest {
        val updatedProperty = property1.copy(name = "Updated Property")
        coEvery { apiCaller.updateProperty(addPropertyInput, "1") } returns CreateOrUpdateResponse(updatedProperty.id)

        viewModel.editProperty(addPropertyInput, "1")

        coVerify { apiCaller.updateProperty(addPropertyInput, "1") }
        assertEquals(updatedProperty.name, viewModel.property.first().name)
        assertEquals(RealPropertyDetailsViewModel.ApiErrors.NONE, viewModel.apiError.first())
    }

    @Test
    fun `editProperty api error sets apiError`() = runTest {
        coEvery { apiCaller.updateProperty(addPropertyInput, "1") } throws Exception("API Error")

        viewModel.editProperty(addPropertyInput, "1")

        coVerify { apiCaller.updateProperty(addPropertyInput, "1") }
        assertEquals(RealPropertyDetailsViewModel.ApiErrors.UPDATE_PROPERTY, viewModel.apiError.first())
    }

    @Test
    fun `onSubmitInviteTenant updates property with tenant details`() = runTest {
        val email = "tenant@example.com"
        val startDate = 1678886400000L
        val endDate = 1681478400000L

        viewModel.loadProperty(property1)

        viewModel.onSubmitInviteTenant(email, startDate, endDate)

        val expectedStartDate = OffsetDateTime.ofInstant(Instant.ofEpochMilli(startDate), ZoneOffset.UTC)
        val expectedEndDate = OffsetDateTime.ofInstant(Instant.ofEpochMilli(endDate), ZoneOffset.UTC)
        val expectedProperty = property1.copy(
            invite = InviteDetailedProperty(
                tenantEmail = email,
                startDate = expectedStartDate,
                endDate = expectedEndDate
            ),
            status = PropertyStatus.invite_sent
        )
        assertEquals(viewModel.property.first().invite?.tenantEmail, expectedProperty.invite?.tenantEmail)
        assertEquals(viewModel.property.first().invite?.startDate, expectedProperty.invite?.startDate)
        assertEquals(viewModel.property.first().invite?.endDate, expectedProperty.invite?.endDate)
        assertEquals(viewModel.property.first().status, expectedProperty.status)
    }
}
