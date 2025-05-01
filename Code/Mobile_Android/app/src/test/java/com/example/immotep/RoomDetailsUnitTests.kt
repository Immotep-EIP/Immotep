package com.example.immotep

import androidx.compose.foundation.layout.size
import androidx.compose.ui.geometry.isEmpty
import kotlin.collections.addAll
import com.example.immotep.inventory.Room
import com.example.immotep.inventory.RoomDetail
import com.example.immotep.inventory.roomDetails.RoomDetailsViewModel
import io.mockk.coEvery
import io.mockk.coVerify
import io.mockk.every
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
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Before
import org.junit.Test

@ExperimentalCoroutinesApi
class RoomDetailsViewModelTest {

    private val closeRoomPanel: (room: Room) -> Unit = mockk()
    private val addDetail: suspend (roomId: String, name: String) -> String? = mockk()

    private lateinit var viewModel: RoomDetailsViewModel
    private val testDispatcher = UnconfinedTestDispatcher()

    private val room1 = Room(id = "1", name = "Living Room")
    private val detail1 = RoomDetail(id = "d1", name = "Sofa")
    private val detail2 = RoomDetail(id = "d2", name = "TV")

    @Before
    fun setUp() {
        Dispatchers.setMain(testDispatcher)
        viewModel = RoomDetailsViewModel(closeRoomPanel, addDetail)
    }

    @Test
    fun `addBaseDetails initializes details correctly`() {
        viewModel.addBaseDetails(arrayOf(detail1, detail2))

        assertEquals(listOf(detail1, detail2), viewModel.details)
    }

    @Test
    fun `onClose updates room details, calls closeRoomPanel, and clears details`() {
        val updatedRoom = room1.copy()
        viewModel.details.addAll(listOf(detail1, detail2))
        every { closeRoomPanel(any()) } returns Unit

        viewModel.onClose(updatedRoom)

        verify { closeRoomPanel(updatedRoom) }
        assertEquals(listOf(detail1, detail2), updatedRoom.details?.toList())
        assertTrue(viewModel.details.isEmpty())
    }

    @Test
    fun `addDetailToRoomDetailPage adds a detail to details`() = runTest {
        val newDetailName = "Table"
        val newDetailId = "d3"
        val roomId = "1"
        coEvery { addDetail(roomId, newDetailName) } returns newDetailId

        viewModel.addDetailToRoomDetailPage(newDetailName, roomId)

        coVerify { addDetail(roomId, newDetailName) }
        assertEquals(1, viewModel.details.size)
        assertTrue(viewModel.details.first().id == newDetailId)
        assertTrue(viewModel.details.first().name == newDetailName)
        assertFalse(viewModel.details.first().completed)
    }

    @Test
    fun `addDetailToRoomDetailPage with null detailId does not add a detail`() = runTest {
        val newDetailName = "Table"
        val roomId = "1"
        coEvery { addDetail(roomId, newDetailName) } returns null
        try {
            viewModel.addDetailToRoomDetailPage(newDetailName, roomId)
            assertTrue(false)
        } catch (e: Exception) {
            coVerify { addDetail(roomId, newDetailName) }
            assertTrue(viewModel.details.isEmpty())
        }
    }

    @Test
    fun `removeDetail removes a detail from details`() {
        viewModel.details.addAll(listOf(detail1, detail2))

        viewModel.removeDetail(detail1)

        assertEquals(listOf(detail2), viewModel.details)
    }

    @Test
    fun `removeDetail with invalid detail index does nothing`() {
        viewModel.details.addAll(listOf(detail1, detail2))
        val nonExistentDetail = RoomDetail(id = "d999", name = "NonExistent")

        viewModel.removeDetail(nonExistentDetail)

        assertEquals(listOf(detail1, detail2), viewModel.details)
    }

    @Test
    fun `onModifyDetail updates detail in details and sets currentlyOpenDetail to null`() = runTest {
        val updatedDetail = detail1.copy(name = "Updated Sofa")
        viewModel.details.addAll(listOf(detail1, detail2))

        viewModel.onModifyDetail(updatedDetail)

        assertEquals(updatedDetail, viewModel.details.first { it.id == detail1.id })
        assertNull(viewModel.currentlyOpenDetail.first())
    }

    @Test
    fun `onModifyDetail with invalid detail index does nothing`() = runTest {
        val nonExistentDetail = RoomDetail(id = "d999", name = "NonExistent")
        viewModel.details.addAll(listOf(detail1, detail2))

        viewModel.onModifyDetail(nonExistentDetail)

        assertEquals(listOf(detail1, detail2), viewModel.details)
        assertNull(viewModel.currentlyOpenDetail.first())
    }

    @Test
    fun `onOpenDetail sets currentlyOpenDetail`() = runTest {
        viewModel.onOpenDetail(detail1)

        assertEquals(detail1, viewModel.currentlyOpenDetail.first())
    }
}