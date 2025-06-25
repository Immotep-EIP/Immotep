package fr.keyz


import fr.keyz.apiCallerServices.AddPropertyInput
import fr.keyz.apiCallerServices.DetailedProperty
import fr.keyz.apiCallerServices.PropertyStatus
import fr.keyz.apiCallerServices.RealPropertyCallerService
import fr.keyz.apiClient.ApiService
import androidx.navigation.NavController
import fr.keyz.apiCallerServices.ApiCallerServiceException
import fr.keyz.apiCallerServices.Damage
import fr.keyz.apiCallerServices.DamageCallerService
import fr.keyz.apiCallerServices.DamageStatus
import fr.keyz.apiCallerServices.Document
import fr.keyz.apiCallerServices.InviteDetailedProperty
import fr.keyz.apiCallerServices.LeaseDetailedProperty
import fr.keyz.apiCallerServices.Priority
import fr.keyz.apiClient.CreateOrUpdateResponse
import fr.keyz.apiClient.mockApi.baseDateStr
import fr.keyz.apiClient.mockApi.fakeDamagesArray
import fr.keyz.apiClient.mockApi.fakeDocument
import fr.keyz.login.dataStore
import fr.keyz.realProperty.details.RealPropertyDetailsViewModel
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
import org.junit.Assert
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
    private val damageApiCaller: DamageCallerService = mockk()
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
        val damageApiCallerField = viewModel::class.java.getDeclaredField("damageApiCaller")
        damageApiCallerField.isAccessible = true
        damageApiCallerField.set(viewModel, damageApiCaller)
    }

    @Test
    fun `loadProperty success updates property with documents and sets isLoading`() = runTest {
        coEvery { apiCaller.getPropertyDocuments("1", any()) } returns arrayOf(fakeDocument)
        coEvery { damageApiCaller.getPropertyDamages(any(), any()) } returns arrayOf()

        viewModel.loadProperty(property1.copy(lease = LeaseDetailedProperty(
            id = "1",
            startDate = OffsetDateTime.now(),
            endDate = OffsetDateTime.now().plusMonths(1),
            tenantEmail = "tenant@example.com",
            tenantName = "John Doe"
        )
        ))

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
        )
        ))

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

    @Test
    fun `loadProperty success updates property with documents and damages and sets isLoading`() = runTest {
        val propertyWithLease = property1.copy(lease = LeaseDetailedProperty(
            id = "lease1", // Make sure lease has an ID
            startDate = OffsetDateTime.now(),
            endDate = OffsetDateTime.now().plusMonths(1),
            tenantEmail = "tenant@example.com",
            tenantName = "John Doe"
        )
        )

        val fakeDamagesArrayToTest = fakeDamagesArray.map { it.toDamage() }.toTypedArray()
        coEvery { apiCaller.getPropertyDocuments(propertyWithLease.id, propertyWithLease.lease!!.id) } returns arrayOf(
            fakeDocument
        )
        coEvery { damageApiCaller.getPropertyDamages(propertyWithLease.id, propertyWithLease.lease!!.id) } returns fakeDamagesArrayToTest

        viewModel.loadProperty(propertyWithLease)

        coVerify { apiCaller.getPropertyDocuments(propertyWithLease.id, propertyWithLease.lease!!.id) }
        coVerify { damageApiCaller.getPropertyDamages(propertyWithLease.id, propertyWithLease.lease!!.id) }

        assertEquals(fakeDocument.id, viewModel.documents.first().id)
        assertEquals(fakeDamagesArrayToTest.first().id, viewModel.damages.first().id)
        assertEquals(fakeDamagesArrayToTest.first().comment, viewModel.damages.first().comment)
        assertEquals(RealPropertyDetailsViewModel.ApiErrors.NONE, viewModel.apiError.value)
    }

    @Test
    fun `loadProperty with damages api error does not crash and sets isLoading`() = runTest {
        val propertyWithLease = property1.copy(lease = LeaseDetailedProperty(
            id = "lease1",
            startDate = OffsetDateTime.now(),
            endDate = OffsetDateTime.now().plusMonths(1),
            tenantEmail = "tenant@example.com",
            tenantName = "John Doe"
        )
        )
        coEvery { apiCaller.getPropertyDocuments(propertyWithLease.id, propertyWithLease.lease!!.id) } returns emptyArray()
        coEvery { damageApiCaller.getPropertyDamages(propertyWithLease.id, propertyWithLease.lease!!.id) } throws ApiCallerServiceException("400")

        viewModel.loadProperty(propertyWithLease)

        coVerify { damageApiCaller.getPropertyDamages(propertyWithLease.id, propertyWithLease.lease!!.id) }
        assertEquals(propertyWithLease.id, viewModel.property.value.id)
        Assert.assertTrue(viewModel.damages.isEmpty())
        assertEquals(RealPropertyDetailsViewModel.ApiErrors.NONE, viewModel.apiError.value)
    }

    @Test
    fun `addDamage successfully adds damage to the list`() {
        val initialDamageCount = viewModel.damages.size
        val newDamage = Damage(
            id = "newDamageId",
            comment = "Test damage",
            createdAt = OffsetDateTime.now(),
            fixPlannedAt = null,
            fixStatus = DamageStatus.PENDING,
            fixedAt = null,
            leaseId = "lease123",
            pictures = emptyArray(),
            priority = Priority.medium,
            read = false,
            roomId = "room1",
            roomName = "Living Room",
            tenantName = "Test Tenant",
            updatedAt = OffsetDateTime.now()
        )

        viewModel.addDamage(newDamage)

        assertEquals(initialDamageCount + 1, viewModel.damages.size)
        Assert.assertTrue(viewModel.damages.contains(newDamage))
        assertEquals(newDamage.id, viewModel.damages.last().id)
        assertEquals(newDamage.comment, viewModel.damages.last().comment)
    }

}
