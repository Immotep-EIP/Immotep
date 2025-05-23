package com.example.keyz


import androidx.navigation.NavController
import com.example.keyz.apiCallerServices.FurnitureCallerService
import com.example.keyz.apiCallerServices.InventoryCallerService
import com.example.keyz.apiCallerServices.RoomCallerService
import com.example.keyz.apiCallerServices.RoomType
import com.example.keyz.apiClient.ApiService
import com.example.keyz.apiClient.mockApi.fakeFurniture
import com.example.keyz.apiClient.mockApi.fakeFurnitureOutputValue
import com.example.keyz.apiClient.mockApi.fakeRoom
import com.example.keyz.apiClient.mockApi.fakeRoomOutputValue
import com.example.keyz.inventory.InventoryViewModel
import com.example.keyz.inventory.Room
import io.mockk.coEvery
import io.mockk.coVerify
import io.mockk.every
import io.mockk.mockk
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.test.StandardTestDispatcher
import kotlinx.coroutines.test.resetMain
import kotlinx.coroutines.test.runTest
import kotlinx.coroutines.test.setMain
import org.junit.After
import org.junit.Before
import org.junit.Test


@OptIn(ExperimentalCoroutinesApi::class)
class InventoryViewModelTest {

    private val navController: NavController = mockk(relaxed = true)
    private val apiService: ApiService = mockk(relaxed = true)
    private val inventoryCallerService: InventoryCallerService = mockk(relaxed = true)
    private val roomCallerService: RoomCallerService = mockk(relaxed = true)
    private val furnitureCallerService: FurnitureCallerService = mockk(relaxed = true)
    private lateinit var viewModel: InventoryViewModel
    private val testDispatcher = StandardTestDispatcher()

    @Before
    fun setup() {
        Dispatchers.setMain(testDispatcher)
        every { navController.context } returns mockk(relaxed = true)
        viewModel = InventoryViewModel(navController, apiService)
        viewModel.javaClass.getDeclaredField("inventoryApiCaller").apply {
            isAccessible = true
            set(viewModel, inventoryCallerService)
        }
        viewModel.javaClass.getDeclaredField("roomApiCaller").apply {
            isAccessible = true
            set(viewModel, roomCallerService)
        }
        viewModel.javaClass.getDeclaredField("furnitureApiCaller").apply {
            isAccessible = true
            set(viewModel, furnitureCallerService)
        }
        viewModel.setPropertyIdAndLeaseId(
            "1",
            "1"
        )
    }

    @After
    fun tearDown() {
        Dispatchers.resetMain()
    }

    /*
    @Test
    fun `getRooms with exit`() = runTest {
        coEvery { inventoryCallerService.getLastInventoryReport(any(), any()) } returns fakeInventoryReport
        viewModel.setInventoryOpen(InventoryOpenValues.EXIT)
        viewModel.getBaseRooms("1")
        viewModel.inventoryErrors.first()
        assert(viewModel.getRooms().isNotEmpty())
    }
    */


    @Test
    fun addRoom() = runTest {
        coEvery { roomCallerService.addRoom(any(), any()) } returns fakeRoomOutputValue
        viewModel.addRoom("testRoom", RoomType.playroom) { }
        assert(viewModel.getRooms().isNotEmpty())
        assert(viewModel.getRooms().find { it.id == "testRoom" } != null)
    }

    @Test
    fun addFurniture() = runTest {
        coEvery { furnitureCallerService.addFurniture(any(), any(), any()) } returns fakeFurnitureOutputValue
        coEvery { roomCallerService.addRoom(any(), any()) } returns fakeRoomOutputValue
        viewModel.addRoom("testRoom", RoomType.playroom) { }
        assert(viewModel.addFurnitureCall("testRoom", "testFurniture", {}) == "testFurniture")
        coVerify { furnitureCallerService.addFurniture(any(), any(), any()) }
    }

    @Test
    fun removeRoom() = runTest {
        coEvery { roomCallerService.addRoom(any(), any()) } returns fakeRoomOutputValue
        viewModel.addRoom("testRoom", RoomType.bedroom) { }
        viewModel.removeRoom("testRoom")
        assert(viewModel.getRooms().isEmpty())
    }

    @Test
    fun editRoom() = runTest {
        coEvery { roomCallerService.addRoom(any(), any()) } returns fakeRoomOutputValue
        viewModel.addRoom("fakeRoom", RoomType.bedroom) { }
        viewModel.editRoom(Room(id = "testRoom", name = "room2"))
        assert(viewModel.getRooms()[0].name == "room2")
    }

    @Test
    fun `onClose with entry`() = runTest {
        coEvery { roomCallerService.addRoom(any(), any()) } returns fakeRoomOutputValue
        viewModel.addRoom("room1", RoomType.bedroom) { }
        viewModel.onClose()
        assert(viewModel.getRooms().isEmpty())
    }
    /*
    @Test
    fun `onClose with exit`() = runTest {
        coEvery { inventoryCallerService.getLastInventoryReport(any(), any()) } returns fakeInventoryReport
        viewModel.getBaseRooms("1")
        viewModel.inventoryErrors.first()
        viewModel.setInventoryOpen(InventoryOpenValues.EXIT)
        viewModel.onClose()
        assert(viewModel.getRooms().isNotEmpty())
    }
    */


    @Test
    fun getBaseRooms() = runTest {
        viewModel.loadInventoryFromRooms(arrayOf(fakeRoom.toRoom(arrayOf(fakeFurniture.toRoomDetail()))))
        assert(viewModel.getRooms().isNotEmpty())
        assert(viewModel.getRooms().size == 1)
    }

    /*
    a refaire après pr fill les bonnes infos
    @Test
    fun sendInventory() = runTest {
        coEvery { roomCallerService.getAllRooms(any(), any()) } returns arrayOf(fakeRoom)
        coEvery { inventoryCallerService.createInventoryReport(any(), any(), any()) } returns Unit
        viewModel.getBaseRooms("1")
        viewModel.inventoryErrors.first()
        viewModel.sendInventory()
        coVerify { inventoryCallerService.createInventoryReport(any(), any(), {}) }
    }

     */


    @Test
    fun `addRoom with error`() = runTest {
        coEvery { roomCallerService.addRoom(any(), any()) } throws Exception()
        viewModel.addRoom("room1", RoomType.office) { }
        assert(viewModel.getRooms().isEmpty())
    }

    @Test
    fun `addFurniture with error`() = runTest {
        coEvery { furnitureCallerService.addFurniture(any(), any(), any()) } throws Exception()
        assert(viewModel.addFurnitureCall("testRoom", "furniture1") {} == null)
        coVerify { furnitureCallerService.addFurniture(any(), any(), any()) }
    }

    @Test
    fun `removeRoom with wrong id`() = runTest {
        coEvery { roomCallerService.addRoom(any(), any()) } returns fakeRoomOutputValue
        viewModel.addRoom("testRoom", RoomType.bedroom) { }
        viewModel.removeRoom("2")
        assert(viewModel.getRooms().isNotEmpty())
    }

    @Test
    fun `editRoom with wrong id`() = runTest {
        coEvery { roomCallerService.addRoom(any(), any()) } returns fakeRoomOutputValue
        viewModel.addRoom("fakeRoomName", RoomType.garage) { }
        viewModel.editRoom(Room(id = "2", name = "room2"))
        assert(viewModel.getRooms()[0].name == "fakeRoomName")
    }

    @Test
    fun `onClose with entry and no rooms`() = runTest {
        viewModel.onClose()
        assert(viewModel.getRooms().isEmpty())
    }
    /*
    @Test
    fun `onClose with exit and no rooms`() = runTest {
        every { inventoryCallerService.getLastInventoryReport(any(), any()) } returns InventoryReport(id = "1", propertyId = "1", type = "start", rooms = Vector())
        viewModel.getBaseRooms("1")
        viewModel.inventoryErrors.first()
        viewModel.setInventoryOpen(InventoryOpenValues.EXIT)
        viewModel.onClose()
        assert(viewModel.getRooms().isEmpty())
    }

    @Test
    fun `sendInventory with error`() = runTest {
        coEvery { roomCallerService.getAllRoomsWithFurniture(any(), any(), any()) } returns mutableStateListOf(
            androidx.room.Room(id = "1", name = "room1"))
        coEvery { inventoryCallerService.createInventoryReport(any(), any(), any()) } throws Exception()
        viewModel.getBaseRooms("1")
        viewModel.inventoryErrors.first()
        viewModel.sendInventory()
        assert(viewModel.inventoryErrors.first().createInventoryReport)
    }
     */
}