package fr.keyz


import fr.keyz.apiCallerServices.RoomType
import fr.keyz.inventory.Room
import fr.keyz.inventory.rooms.RoomsViewModel
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
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Before
import org.junit.Test

@ExperimentalCoroutinesApi
class RoomsViewModelTest {

    private val getRooms: () -> Array<Room> = mockk()
    private val addRoom: suspend (String, RoomType) -> String? = mockk()
    private val removeRoom: (String) -> Unit = mockk()
    private val closeInventory: () -> Unit = mockk()
    private val editRoom: (Room) -> Unit = mockk()
    private val confirmInventory: () -> Boolean = mockk()

    private lateinit var viewModel: RoomsViewModel
    private val testDispatcher = UnconfinedTestDispatcher()

    private val room1 = Room(id = "1", name = "Living Room")
    private val room2 = Room(id = "2", name = "Bedroom")

    @Before
    fun setUp() {
        Dispatchers.setMain(testDispatcher)
        viewModel = RoomsViewModel(
            getRooms,
            addRoom,
            removeRoom,
            closeInventory,
            editRoom,
            confirmInventory
        )
    }

    @Test
    fun `handleBaseRooms initializes allRooms correctly`() {
        every { getRooms() } returns arrayOf(room1, room2)

        viewModel.handleBaseRooms()

        assertEquals(listOf(room1, room2), viewModel.allRooms)
        verify { getRooms() }
    }

    @Test
    fun `onClose clears allRooms and calls closeInventory`() {
        viewModel.allRooms.addAll(listOf(room1, room2))
        every { closeInventory() } returns Unit

        viewModel.onClose()

        assertTrue(viewModel.allRooms.isEmpty())
        verify { closeInventory() }
    }
    /*
    Commented to wait for a good response of the server
    @Test
    fun `onConfirmInventory with successful confirmation calls onClose`() {
        every { confirmInventory() } returns true
        every { closeInventory() } returns Unit

        viewModel.onConfirmInventory()

        verify { confirmInventory() }
        verify { closeInventory() }
    }

     */

    @Test
    fun `onConfirmInventory with failed confirmation sets showNotCompletedRooms to true`() = runTest {
        every { confirmInventory() } returns false

        viewModel.onConfirmInventory()

        verify { confirmInventory() }
        assertTrue(viewModel.showNotCompletedRooms.first())
    }

    @Test
    fun `addARoom adds a room to allRooms`() = runTest {
        val newRoomName = "Kitchen"
        val roomType = RoomType.kitchen
        val newRoomId = "3"
        coEvery { addRoom(newRoomName, roomType) } returns newRoomId

        viewModel.addARoom(newRoomName, roomType)
        testDispatcher.scheduler.advanceUntilIdle()
        coVerify { addRoom(newRoomName, roomType) }
        assertEquals(1, viewModel.allRooms.size)
        assertTrue(viewModel.allRooms.first().name == newRoomName)
        assertTrue(viewModel.allRooms.first().id == newRoomId)
    }

    @Test
    fun `addARoom with null roomId does not add a room`() = runTest {
        val newRoomName = "Kitchen"
        val roomType = RoomType.kitchen
        coEvery { addRoom(newRoomName, roomType) } returns null

        try {
            viewModel.addARoom(newRoomName, roomType)
            assertTrue(false)
        } catch (e : Exception) {
            assertEquals("impossible_to_add_room", e.message)
            coVerify { addRoom(newRoomName, roomType) }
            assertTrue(viewModel.allRooms.isEmpty())
        }
    }

    @Test
    fun `openRoomPanel sets currentlyOpenRoom`() = runTest {
        viewModel.openRoomPanel(room1)

        assertEquals(room1, viewModel.currentlyOpenRoom.first())
    }

    @Test
    fun `closeRoomPanel updates room in allRooms and sets currentlyOpenRoom to null`() = runTest {
        val updatedRoom = room1.copy(name = "Living Room - Updated")
        viewModel.allRooms.addAll(listOf(room1, room2))
        every { editRoom(updatedRoom) } returns Unit

        viewModel.closeRoomPanel(updatedRoom)

        verify { editRoom(updatedRoom) }
        assertEquals(updatedRoom, viewModel.allRooms.first { it.id == room1.id })
        assertNull(viewModel.currentlyOpenRoom.first())
    }

    @Test
    fun `closeRoomPanel with invalid room index does nothing`() = runTest {
        val updatedRoom = Room(id = "999", name = "NonExistent Room")
        viewModel.allRooms.addAll(listOf(room1, room2))

        viewModel.closeRoomPanel(updatedRoom)

        verify(exactly = 0) { editRoom(any()) }
        assertEquals(listOf(room1, room2), viewModel.allRooms)
        assertNull(viewModel.currentlyOpenRoom.first())
    }
}
