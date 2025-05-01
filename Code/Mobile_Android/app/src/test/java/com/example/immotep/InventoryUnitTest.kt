package com.example.immotep

/*
import androidx.compose.runtime.mutableStateListOf
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.FurnitureCallerService
import com.example.immotep.apiCallerServices.FurnitureInput
import com.example.immotep.apiCallerServices.InventoryCallerService
import com.example.immotep.apiCallerServices.InventoryReportInput
import com.example.immotep.apiCallerServices.RoomCallerService
import com.example.immotep.apiCallerServices.RoomOutput
import com.example.immotep.apiClient.AddRoomInput
import com.example.immotep.apiClient.ApiService
import com.example.immotep.apiClient.mockApi.fakeFurniture
import com.example.immotep.apiClient.mockApi.fakeInventoryReport
import com.example.immotep.apiClient.mockApi.fakeRoom
import com.example.immotep.inventory.InventoryOpenValues
import com.example.immotep.inventory.InventoryReportOutput
import com.example.immotep.inventory.InventoryViewModel
import com.example.immotep.inventory.Room
import io.mockk.coEvery
import io.mockk.coVerify
import io.mockk.every
import io.mockk.mockk
import io.mockk.verify
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.ExperimentalCoroutinesApi
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.test.StandardTestDispatcher
import kotlinx.coroutines.test.resetMain
import kotlinx.coroutines.test.runTest
import kotlinx.coroutines.test.setMain
import okhttp3.internal.wait
import org.junit.After
import org.junit.Before
import org.junit.Test
import java.util.Vector

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
        viewModel = InventoryViewModel(navController, "1", apiService)
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
    }

    @After
    fun tearDown() {
        Dispatchers.resetMain()
    }

    @Test
    fun `setInventoryOpen with errors`() = runTest {
        viewModel.getBaseRooms("1")
        viewModel.inventoryErrors.first()
        viewModel.setInventoryOpen(InventoryOpenValues.ENTRY)
        coVerify(exactly = 0) { inventoryCallerService.getLastInventoryReport(any(), any()) }
    }

    @Test
    fun `setInventoryOpen with exit and no old report`() = runTest {
        coEvery { inventoryCallerService.getLastInventoryReport(any(), any()) } throws Exception()
        viewModel.setInventoryOpen(InventoryOpenValues.EXIT)
        assert(viewModel.cannotMakeExitInventory.value)
    }

    @Test
    fun `setInventoryOpen with exit and empty last inventory`() = runTest {
        coEvery { inventoryCallerService.getLastInventoryReport(any(), any()) } returns fakeInventoryReport
        viewModel.setInventoryOpen(InventoryOpenValues.EXIT)
        assert(viewModel.cannotMakeExitInventory.first())
    }

    @Test
    fun `setInventoryOpen with exit and error on last inventory`() = runTest {
        coEvery { inventoryCallerService.getLastInventoryReport(any(), any()) } throws Exception()
        viewModel.setInventoryOpen(InventoryOpenValues.EXIT)
        assert(viewModel.cannotMakeExitInventory.first())
    }

    @Test
    fun closeCannotMakeExitInventory() = runTest {
        viewModel.closeCannotMakeExitInventory()
        assert(!viewModel.cannotMakeExitInventory.first())
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
        coEvery { roomCallerService.addRoom(any(), any(), any()) } returns fakeRoom
        viewModel.addRoom("testRoom") { }
        assert(viewModel.getRooms().isNotEmpty())
        assert(viewModel.getRooms().find { it.id == "testRoom" } != null)
    }

    @Test
    fun addFurniture() = runTest {
        coEvery { furnitureCallerService.addFurniture(any(), any(), any(), any()) } returns fakeFurniture
        coEvery { roomCallerService.addRoom(any(), any(), any()) } returns fakeRoom
        viewModel.addRoom("testRoom") { }
        assert(viewModel.addFurnitureCall("testRoom", "testFurniture", {}) == "testFurniture")
        coVerify { furnitureCallerService.addFurniture(any(), any(), any(), any()) }
    }

    @Test
    fun removeRoom() = runTest {
        coEvery { roomCallerService.addRoom(any(), any(), any()) } returns fakeRoom
        viewModel.addRoom("testRoom") { }
        viewModel.removeRoom("testRoom")
        assert(viewModel.getRooms().isEmpty())
    }

    @Test
    fun editRoom() = runTest {
        coEvery { roomCallerService.addRoom(any(), any(), any()) } returns fakeRoom
        viewModel.addRoom("fakeRoom") { }
        viewModel.editRoom(Room(id = "testRoom", name = "room2"))
        assert(viewModel.getRooms()[0].name == "room2")
    }

    @Test
    fun `onClose with entry`() = runTest {
        coEvery { roomCallerService.addRoom(any(), any(), any()) } returns fakeRoom
        viewModel.addRoom("room1") { }
        viewModel.setInventoryOpen(InventoryOpenValues.ENTRY)
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
        coEvery { roomCallerService.getAllRoomsWithFurniture(any(), any(), any()) } returns arrayOf(fakeRoom.toRoom(arrayOf(fakeFurniture.toRoomDetail())))
        viewModel.getBaseRooms("1")
        testDispatcher.scheduler.advanceUntilIdle()
        coVerify { roomCallerService.getAllRoomsWithFurniture(any(), any(), any()) }
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
    fun `getBaseRooms with error`() = runTest {
        coEvery { roomCallerService.getAllRoomsWithFurniture(any(), any(), any()) } throws Exception()
        viewModel.getBaseRooms("1")
        testDispatcher.scheduler.advanceUntilIdle()
        assert(viewModel.getRooms().isEmpty())
    }

    @Test
    fun `getLastInventory with error`() = runTest {
        coEvery { inventoryCallerService.getLastInventoryReport(any(), any()) } throws Exception()
        viewModel.getBaseRooms("1")
        testDispatcher.scheduler.advanceUntilIdle()
        assert(!viewModel.cannotMakeExitInventory.first())
    }

    @Test
    fun `addRoom with error`() = runTest {
        coEvery { roomCallerService.addRoom(any(), any(), any()) } throws Exception()
        viewModel.addRoom("room1") { }
        assert(viewModel.getRooms().isEmpty())
    }

    @Test
    fun `addFurniture with error`() = runTest {
        coEvery { furnitureCallerService.addFurniture(any(), any(), any(), any()) } throws Exception()
        assert(viewModel.addFurnitureCall("testRoom", "furniture1") {} == null)
        coVerify { furnitureCallerService.addFurniture(any(), any(), any(), any()) }
    }

    @Test
    fun `removeRoom with wrong id`() = runTest {
        coEvery { roomCallerService.addRoom(any(), any(), any()) } returns fakeRoom
        viewModel.addRoom("testRoom") { }
        viewModel.removeRoom("2")
        assert(viewModel.getRooms().isNotEmpty())
    }

    @Test
    fun `editRoom with wrong id`() = runTest {
        coEvery { roomCallerService.addRoom(any(), any(), any()) } returns fakeRoom
        viewModel.addRoom("fakeRoomName") { }
        viewModel.editRoom(Room(id = "2", name = "room2"))
        assert(viewModel.getRooms()[0].name == "fakeRoomName")
    }

    @Test
    fun `onClose with entry and no rooms`() = runTest {
        viewModel.setInventoryOpen(InventoryOpenValues.ENTRY)
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

 */