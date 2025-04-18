package com.example.immotep.inventory.loaderButton

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.InventoryCallerService
import com.example.immotep.apiCallerServices.RoomCallerService
import com.example.immotep.apiClient.ApiService
import com.example.immotep.inventory.Room
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import kotlinx.coroutines.sync.Mutex


class LoaderInventoryViewModel(
    private val navController: NavController,
    apiService: ApiService
) : ViewModel() {
    /**
     * InventoryApiErrors, made for store the errors that can happen during the api calls
     */
    data class InventoryApiErrors(
        var getAllRooms: Boolean = false,
        var getLastInventoryReport: Boolean = false,
        var errorRoomName: String? = null,
        var createInventoryReport: Boolean = false,
    )

    init {
        println("LoaderInventoryViewModel initialized")
    }

    private val inventoryApiCaller = InventoryCallerService(apiService, navController)
    private val roomApiCaller = RoomCallerService(apiService, navController)

    private val _loadingMutex = Mutex()
    private val _inventoryErrors = MutableStateFlow(InventoryApiErrors())
    private val _oldReportId = MutableStateFlow<String?>(null)
    private val _internalIsLoading = MutableStateFlow(false)

    private val rooms = mutableStateListOf<Room>()

    val oldReportId = _oldReportId.asStateFlow()
    val inventoryErrors = _inventoryErrors.asStateFlow()
    val isLoading = _internalIsLoading.asStateFlow()

    private suspend fun tryLoadLastInventory(propertyId: String) {
        val inventoryReport = inventoryApiCaller.getLastInventoryReport(propertyId)
        val lastInventoryRoomsAsRooms = inventoryReport.getRoomsAsRooms(empty = true)
        _oldReportId.value = inventoryReport.id
        this.rooms.addAll(lastInventoryRoomsAsRooms)
    }

    private suspend fun tryGetBaseRooms(propertyId: String) {
        try {
            val newRooms = roomApiCaller.getAllRoomsWithFurniture(
                propertyId,
                { _inventoryErrors.value = _inventoryErrors.value.copy(getAllRooms = true) },
                { _inventoryErrors.value = _inventoryErrors.value.copy(errorRoomName = it) }
            )
            rooms.addAll(newRooms)
        } catch (e: Exception) {
            println("Error during get base rooms ${e.message}")
            _inventoryErrors.value = _inventoryErrors.value.copy(getLastInventoryReport = true)
            e.printStackTrace()
        }
    }

    fun loadInventory(propertyId: String) {
        println("Load the inventory...")
        /*
        try {
            viewModelScope.cancel()
        } catch (e: Exception) {
            println("The scope does not contain any coroutines")
        }
        */
        _inventoryErrors.value = InventoryApiErrors()
        this.rooms.clear()
        viewModelScope.launch {
            _loadingMutex.lock()
            _internalIsLoading.value = true
            try {
                tryLoadLastInventory(propertyId)
            } catch (e: Exception) {
                println("Impossible to load the last inventory ${e.message}")
                tryGetBaseRooms(propertyId)
            } finally {
                _internalIsLoading.value = false
                _loadingMutex.unlock()
            }
        }
    }

    fun onClick(setIsLoading : (Boolean) -> Unit, propertyId : String) {
        viewModelScope.launch {
            try {
                setIsLoading(true)
                _loadingMutex.lock()
                setIsLoading(false)
                _loadingMutex.unlock()
                println("first size...")
                navController.navigate("inventory/${propertyId}")
            } catch (e : Exception) {
                setIsLoading(false)
                println("Error occured on the onClick of LoaderInventoryButtonViewModel ${e.message}")
                _loadingMutex.unlock()
                e.printStackTrace()
            }
        }
    }

    fun getRooms() : Array<Room> {
        if (!_loadingMutex.tryLock()) {
            println("LOCKEDD")
            return arrayOf()
        }
        _loadingMutex.unlock()
        println("get the loader rooms size ? ${rooms.size}")
        return this.rooms.toTypedArray()
    }
}