package com.example.keyz


import kotlin.jvm.java


import android.net.Uri
import androidx.navigation.NavController
import com.example.keyz.apiCallerServices.AICallerService
import com.example.keyz.apiClient.ApiService
import com.example.keyz.apiClient.mockApi.fakeAiCallOutput
import com.example.keyz.inventory.Cleanliness
import com.example.keyz.inventory.RoomDetail
import com.example.keyz.inventory.State
import com.example.keyz.inventory.roomDetails.OneDetail.OneDetailViewModel
import com.example.keyz.inventory.roomDetails.OneDetail.RoomDetailsError
import com.example.keyz.utils.Base64Utils
import io.mockk.Runs
import io.mockk.coEvery
import io.mockk.coVerify
import io.mockk.every
import io.mockk.just
import io.mockk.mockk
import io.mockk.spyk
import io.mockk.unmockkStatic
import io.mockk.verify
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.test.UnconfinedTestDispatcher
import kotlinx.coroutines.test.runTest
import kotlinx.coroutines.test.setMain
import org.junit.After
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertTrue
import org.junit.Before
import org.junit.Test

@ExperimentalCoroutinesApi
class OneDetailViewModelTest {

    private val apiService: ApiService = mockk()
    private val navController: NavController = mockk {
        every { context } returns mockk() // Mock the context if needed for Base64Utils
    }
    private val aiCallerService: AICallerService = mockk() // Mock AICallerService
    private lateinit var viewModel: OneDetailViewModel
    private val testDispatcher = UnconfinedTestDispatcher()

    private val detail1 = RoomDetail(id = "d1", name = "Sofa")
    private val uri1: Uri = spyk()

    @Before
    fun setUp() {
        Dispatchers.setMain(testDispatcher)
        viewModel = OneDetailViewModel(apiService, navController)

        val aiCallerField = viewModel::class.java.getDeclaredField("aiCaller")
        aiCallerField.isAccessible = true
        aiCallerField.set(viewModel, aiCallerService)
        every { uri1.scheme } returns "custom"
        every { uri1.equals(any()) } returns true
    }

    @After
    fun tearDown() {
        unmockkStatic(Base64Utils::class)
    }

    @Test
    fun `reset with newDetail initializes correctly`() = runTest {
        viewModel.reset(detail1)

        assertEquals(detail1, viewModel.detail.first())
        assertTrue(viewModel.picture.isEmpty())
        assertTrue(viewModel.entryPictures.isEmpty())
        assertEquals(RoomDetailsError(), viewModel.errors.first())
    }

    @Test
    fun `reset with null newDetail initializes with empty detail`() = runTest {
        viewModel.reset(null)

        assertTrue(viewModel.detail.first().name.isEmpty())
        assertTrue(viewModel.detail.first().comment.isEmpty())
        assertTrue(!viewModel.detail.first().completed)
        assertTrue(viewModel.detail.first().cleanliness == Cleanliness.not_set)
        assertTrue(viewModel.detail.first().status == State.not_set)
        assertTrue(viewModel.picture.isEmpty())
        assertTrue(viewModel.detail.first().pictures.isEmpty())
        assertTrue(viewModel.entryPictures.isEmpty())
        assertFalse(viewModel.errors.first().name)
        assertFalse(viewModel.errors.first().comment)
    }

    @Test
    fun `setName updates name and clears name error`() = runTest {
        val newName = "Updated Sofa"
        viewModel.setName(newName)

        assertEquals(newName, viewModel.detail.first().name)
        assertFalse(viewModel.errors.first().name)
    }

    @Test
    fun `setName with long name does not update name`() = runTest {
        val longName = "a".repeat(51)
        val initialName = viewModel.detail.first().name
        viewModel.setName(longName)

        assertEquals(initialName, viewModel.detail.first().name)
    }

    @Test
    fun `setComment updates comment and clears comment error`() = runTest {
        val newComment = "This is a new comment."
        viewModel.setComment(newComment)

        assertEquals(newComment, viewModel.detail.first().comment)
        assertFalse(viewModel.errors.first().comment)
    }

    @Test
    fun `setComment with long comment does not update comment`() = runTest {
        val longComment = "a".repeat(501)
        val initialComment = viewModel.detail.first().comment
        viewModel.setComment(longComment)

        assertEquals(initialComment, viewModel.detail.first().comment)
    }

    @Test
    fun `setCleanliness updates cleanliness and clears cleanliness error`() = runTest {
        val newCleanliness = Cleanliness.clean
        viewModel.setCleanliness(newCleanliness)

        assertEquals(newCleanliness, viewModel.detail.first().cleanliness)
        assertFalse(viewModel.errors.first().cleanliness)
    }

    @Test
    fun `setStatus updates status and clears status error`() = runTest {
        val newStatus = State.good
        viewModel.setStatus(newStatus)

        assertEquals(newStatus, viewModel.detail.first().status)
        assertFalse(viewModel.errors.first().status)
    }

    @Test
    fun `addPicture adds a picture and clears picture error`() = runTest {
        viewModel.addPicture(uri1)

        assertEquals(listOf(uri1), viewModel.picture)
        assertFalse(viewModel.errors.first().picture)
    }

    @Test
    fun `removePicture removes a picture`() = runTest {
        viewModel.addPicture(uri1)
        viewModel.removePicture(0)

        assertTrue(viewModel.picture.isEmpty())
    }

    @Test
    fun `onConfirm with valid data calls onModifyDetail and resets`() = runTest {
        val onModifyDetail: (RoomDetail) -> Unit = mockk()
        every { onModifyDetail(any()) } just Runs
        viewModel.setName("Valid Name")
        viewModel.setComment("Valid Comment")
        viewModel.setStatus(State.good)
        viewModel.setCleanliness(Cleanliness.clean)
        viewModel.addPicture(uri1)

        viewModel.onConfirm(onModifyDetail, false)

        verify { onModifyDetail(any()) }
        assertTrue(viewModel.detail.first().name.isEmpty())
        assertTrue(viewModel.detail.first().comment.isEmpty())
        assertTrue(!viewModel.detail.first().completed)
        assertTrue(viewModel.detail.first().cleanliness == Cleanliness.not_set)
        assertTrue(viewModel.detail.first().status == State.not_set)
        assertTrue(viewModel.picture.isEmpty())
        assertTrue(viewModel.detail.first().pictures.isEmpty())
    }

    @Test
    fun `onConfirm with invalid data sets errors`() = runTest {
        val onModifyDetail: (RoomDetail) -> Unit = mockk()
        every { onModifyDetail(any()) } just Runs

        viewModel.onConfirm(onModifyDetail, false)

        verify(exactly = 0) { onModifyDetail(any()) }
        val errors = viewModel.errors.first()
        assertTrue(errors.name)
        assertTrue(errors.comment)
        assertTrue(errors.status)
        assertTrue(errors.cleanliness)
        assertTrue(errors.picture)
    }

    @Test
    fun `onClose calls onModifyDetail and resets`() = runTest {
        val onModifyDetail: (RoomDetail) -> Unit = mockk()
        every { onModifyDetail(any()) } just Runs
        viewModel.setName("Valid Name")
        viewModel.setComment("Valid Comment")
        viewModel.setStatus(State.good)
        viewModel.setCleanliness(Cleanliness.clean)
        viewModel.addPicture(uri1)

        viewModel.onClose(onModifyDetail, false)

        verify { onModifyDetail(any()) }
        assertTrue(viewModel.detail.first().name.isEmpty())
        assertTrue(viewModel.detail.first().comment.isEmpty())
        assertTrue(!viewModel.detail.first().completed)
        assertTrue(viewModel.detail.first().cleanliness == Cleanliness.not_set)
        assertTrue(viewModel.detail.first().status == State.not_set)
        assertTrue(viewModel.picture.isEmpty())
        assertTrue(viewModel.detail.first().pictures.isEmpty())
    }

    @Test
    fun `summarizeOrCompare with empty pictures sets picture error`() = runTest {
        viewModel.summarizeOrCompare(null, "propertyId", "leaseId", true, false)

        assertTrue(viewModel.errors.first().picture)
    }

    @Test
    fun `summarizeOrCompare with oldReportId null calls summarize`() = runTest {
        val propertyId = "propertyId"
        val leaseId = "leaseId"
        val isRoom = true
        viewModel.addPicture(uri1)
        coEvery { aiCallerService.summarize(any(), any(), any()) } returns fakeAiCallOutput

        viewModel.summarizeOrCompare(null, propertyId, leaseId, isRoom, false)

        coVerify {
            aiCallerService.summarize(
                propertyId = propertyId,
                leaseId = leaseId,
                input = any(),
            )
        }
    }

    @Test
    fun `summarizeOrCompare with oldReportId not null calls compare`() = runTest {
        val oldReportId = "oldReportId"
        val propertyId = "propertyId"
        val leaseId = "leaseId"
        val isRoom = false
        viewModel.addPicture(uri1)
        coEvery { aiCallerService.compare(any(), any(), any(), any()) } returns fakeAiCallOutput

        viewModel.summarizeOrCompare(oldReportId, propertyId, leaseId, isRoom, true)

        coVerify {
            aiCallerService.compare(
                propertyId = propertyId,
                oldReportId = oldReportId,
                leaseId = leaseId,
                input = any(),
            )
        }
    }

    @Test
    fun `summarizeOrCompare sets aiCallError on summarize error`() = runTest {
        val propertyId = "propertyId"
        val isRoom = true
        viewModel.addPicture(uri1)
        coEvery { aiCallerService.summarize(any(), any(), any()) } answers {
            thirdArg<() -> Unit>().invoke()
            fakeAiCallOutput
        }

        viewModel.summarizeOrCompare(null, propertyId, "leaseId", isRoom, false)

        assertTrue(viewModel.aiCallError.first())
    }

    @Test
    fun `summarize updates detail with aiResponse`() = runTest {
        val propertyId = "propertyId"
        val aiResponse = fakeAiCallOutput
        viewModel.addPicture(uri1)
        coEvery { aiCallerService.summarize(any(), any(), any()) } returns aiResponse

        viewModel.summarizeOrCompare(null, propertyId, "leaseId",true, false)

        assertTrue(viewModel.detail.first().status == fakeAiCallOutput.state)
        assertTrue(viewModel.detail.first().cleanliness == fakeAiCallOutput.cleanliness)
        assertTrue(viewModel.detail.first().comment == fakeAiCallOutput.note)
        assertEquals(RoomDetailsError(), viewModel.errors.first())
    }

    @Test
    fun `compare updates detail with aiResponse`() = runTest {
        val oldReportId = "oldReportId"
        val propertyId = "propertyId"
        viewModel.addPicture(uri1)
        coEvery { aiCallerService.compare(any(), any(), any(), any()) } returns fakeAiCallOutput

        viewModel.summarizeOrCompare(oldReportId, propertyId, "leaseId",false, true)

        assertTrue(viewModel.detail.first().status == fakeAiCallOutput.state)
        assertTrue(viewModel.detail.first().cleanliness == fakeAiCallOutput.cleanliness)
        assertTrue(viewModel.detail.first().comment == fakeAiCallOutput.note)
    }
}
