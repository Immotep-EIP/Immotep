package com.example.immotep
/*
import android.net.Uri
import androidx.arch.core.executor.testing.InstantTaskExecutorRule
import com.example.immotep.apiCallerServices.AddPropertyInput
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.test.StandardTestDispatcher
import kotlinx.coroutines.test.resetMain
import kotlinx.coroutines.test.runTest
import kotlinx.coroutines.test.setMain
import org.junit.After
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertTrue
import org.junit.Before
import org.junit.Rule
import org.junit.Test
import org.mockito.Mockito.mock
import org.mockito.Mockito.verify

@ExperimentalCoroutinesApi
class AddOrEditPropertyViewModelTest {

    @get:Rule
    val instantTaskExecutorRule = InstantTaskExecutorRule()

    private val testDispatcher = StandardTestDispatcher()

    private lateinit var viewModel: AddOrEditPropertyViewModel

    @Before
    fun setup() {
        Dispatchers.setMain(testDispatcher)
        viewModel = AddOrEditPropertyViewModel()
    }

    @After
    fun tearDown() {
        Dispatchers.resetMain()
    }

    @Test
    fun `setBaseValue updates propertyForm`() = runTest {
        val property = AddPropertyInput(address = "Test Address")
        viewModel.setBaseValue(property)
        assertEquals(property, viewModel.propertyForm.first())
    }

    @Test
    fun `setAddress updates address in propertyForm`() = runTest {
        viewModel.setAddress("New Address")
        assertEquals("New Address", viewModel.propertyForm.first().address)
    }

    @Test
    fun `setZipCode updates postal_code in propertyForm`() = runTest {
        viewModel.setZipCode("12345")
        assertEquals("12345", viewModel.propertyForm.first().postal_code)
    }

    @Test
    fun `setCountry updates country in propertyForm`() = runTest {
        viewModel.setCountry("New Country")
        assertEquals("New Country", viewModel.propertyForm.first().country)
    }

    @Test
    fun `setArea updates area_sqm in propertyForm`() = runTest {
        viewModel.setArea(100.0)
        assertEquals(100.0, viewModel.propertyForm.first().area_sqm)
    }

    @Test
    fun `setRental updates rental_price_per_month in propertyForm`() = runTest {
        viewModel.setRental(500)
        assertEquals(500, viewModel.propertyForm.first().rental_price_per_month)
    }

    @Test
    fun `setDeposit updates deposit_price in propertyForm`() = runTest {
        viewModel.setDeposit(1000)
        assertEquals(1000, viewModel.propertyForm.first().deposit_price)
    }

    @Test
    fun `setName updates name in propertyForm`() = runTest {
        viewModel.setName("Test Name")
        assertEquals("Test Name", viewModel.propertyForm.first().name)
    }

    @Test
    fun `setCity updates city in propertyForm`() = runTest {
        viewModel.setCity("Test City")
        assertEquals("Test City", viewModel.propertyForm.first().city)
    }

    @Test
    fun `addPicture adds picture to pictures list`() {
        val uri = mock(Uri::class.java)
        viewModel.addPicture(uri)
        assertTrue(viewModel.pictures.contains(uri))
    }

    @Test
    fun `setAppartementNumber updates apartment_number in propertyForm`() = runTest {
        viewModel.setAppartementNumber("123")
        assertEquals("123", viewModel.propertyForm.first().apartment_number)
    }

    @Test
    fun `reset with no baseValue resets propertyForm to default`() = runTest {
        viewModel.setAddress("Some Address")
        viewModel.reset()
        assertEquals(AddPropertyInput(), viewModel.propertyForm.first())
    }

    @Test
    fun `reset with baseValue sets propertyForm to baseValue`() = runTest {
        val baseValue = AddPropertyInput(address = "Base Address")
        viewModel.reset(baseValue)
        assertEquals(baseValue, viewModel.propertyForm.first())
    }

    @Test
    fun `onSubmit with valid data calls sendFormFn and onClose`() = runTest {
        val mockOnClose = mock<() -> Unit>()
        val mockSendFormFn: suspend (AddPropertyInput) -> Unit = mock()
        val property = AddPropertyInput(
            address = "Valid Address",
            postal_code = "12345",
            country = "Valid Country",
            area_sqm = 10.0,
            rental_price_per_month = 100,
            deposit_price = 50
        )
        viewModel.setBaseValue(property)

        viewModel.onSubmit(mockOnClose, mockSendFormFn)
        testDispatcher.scheduler.advanceUntilIdle()

        verify(mockSendFormFn).invoke(property)
        verify(mockOnClose).invoke()
        assertEquals(AddPropertyInput(), viewModel.propertyForm.first())
    }

    @Test
    fun `onSubmit with invalid address sets address error`() = runTest {
        val mockOnClose = mock<() -> Unit>()
        val mockSendFormFn: suspend (AddPropertyInput) -> Unit = mock()
        viewModel.setAddress("aa")

        viewModel.onSubmit(mockOnClose, mockSendFormFn)
        testDispatcher.scheduler.advanceUntilIdle()

        assertTrue(viewModel.propertyFormError.first().address)
    }

    @Test
    fun `onSubmit with invalid zipCode sets zipCode error`() = runTest {
        val mockOnClose = mock<() -> Unit>()
        val mockSendFormFn: suspend (AddPropertyInput) -> Unit = mock()
        viewModel.setZipCode("123")

        viewModel.onSubmit(mockOnClose, mockSendFormFn)
        testDispatcher.scheduler.advanceUntilIdle()

        assertTrue(viewModel.propertyFormError.first().zipCode)
    }

    @Test
    fun `onSubmit with invalid country sets country error`() = runTest {
        val mockOnClose = mock<() -> Unit>()
        val mockSendFormFn: suspend (AddPropertyInput) -> Unit = mock()
        viewModel.setCountry("aa")

        viewModel.onSubmit(mockOnClose, mockSendFormFn)
        testDispatcher.scheduler.advanceUntilIdle()

        assertTrue(viewModel.propertyFormError.first().country)
    }

    @Test
    fun `onSubmit with invalid area sets area error`() = runTest {
        val mockOnClose = mock<() -> Unit>()
        val mockSendFormFn: suspend (AddPropertyInput) -> Unit = mock()
        viewModel.setArea(0.0)

        viewModel.onSubmit(mockOnClose, mockSendFormFn)
        testDispatcher.scheduler.advanceUntilIdle()

        assertTrue(viewModel.propertyFormError.first().area)
    }

    @Test
    fun `onSubmit with invalid rental sets rental error`() = runTest {
        val mockOnClose = mock<() -> Unit>()
        val mockSendFormFn: suspend (AddPropertyInput) -> Unit = mock()
        viewModel.setRental(0)

        viewModel.onSubmit(mockOnClose, mockSendFormFn)
        testDispatcher.scheduler.advanceUntilIdle()

        assertTrue(viewModel.propertyFormError.first().rental)
    }

    @Test
    fun `onSubmit with invalid deposit sets deposit error`() = runTest {
        val mockOnClose = mock<() -> Unit>()
        val mockSendFormFn: suspend (AddPropertyInput) -> Unit = mock()
        viewModel.setDeposit(0)

        viewModel.onSubmit(mockOnClose, mockSendFormFn)
        testDispatcher.scheduler.advanceUntilIdle()

        assertTrue(viewModel.propertyFormError.first().deposit)
    }

    @Test
    fun `onSubmit with multiple invalid fields sets multiple errors`() = runTest {
        val mockOnClose = mock<() -> Unit>()
        val mockSendFormFn: suspend (AddPropertyInput) -> Unit = mock()
        viewModel.setAddress("aa")
        viewModel.setZipCode("123")
        viewModel.setCountry("aa")
        viewModel.setArea(0.0)
        viewModel.setRental(0)
        viewModel.setDeposit(0)

        viewModel.onSubmit(mockOnClose, mockSendFormFn)
        testDispatcher.scheduler.advanceUntilIdle()
    }
*/