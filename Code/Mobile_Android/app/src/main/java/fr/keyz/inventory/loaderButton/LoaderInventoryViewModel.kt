package fr.keyz.inventory.loaderButton

import android.content.Context
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import fr.keyz.apiCallerServices.InventoryCallerService
import fr.keyz.apiCallerServices.RoomCallerService
import fr.keyz.apiClient.ApiService
import fr.keyz.inventory.Room
import fr.keyz.inventory.RoomDetail
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import kotlinx.coroutines.sync.Mutex
import kotlinx.coroutines.sync.withLock


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


    private val inventoryApiCaller = InventoryCallerService(apiService, navController)
    private val roomApiCaller = RoomCallerService(apiService, navController)

    private var _newValueSetByCompletedInventory = false
    private val _loadingMutex = Mutex()
    private val _inventoryErrors = MutableStateFlow(InventoryApiErrors())
    private val _oldReportId = MutableStateFlow<String?>(null)
    private val _internalIsLoading = MutableStateFlow(false)

    private val rooms = mutableStateListOf<Room>()

    val oldReportId = _oldReportId.asStateFlow()
    val inventoryErrors = _inventoryErrors.asStateFlow()
    val isLoading = _internalIsLoading.asStateFlow()

    private fun setNewValuesForLastInventoryDetails(
        originalRoomDetails : Array<RoomDetail>,
        lastInventoryRoomDetails : Array<RoomDetail>
    ) : Array<RoomDetail> {
        return originalRoomDetails.map { detail ->
            return@map lastInventoryRoomDetails.find {
                lastInventoryDetail ->
                lastInventoryDetail.id == detail.id
            } ?: detail.copy(newItem = true)
        }.toTypedArray()
    }

    private suspend fun tryLoadLastInventory(propertyId: String, baseRooms : Array<Room>) {
        val inventoryReport = inventoryApiCaller.getLastInventoryReport(propertyId)
        val lastInventoryRoomsAsRooms = inventoryReport.getRoomsAsRooms(empty = true)
        _oldReportId.value = inventoryReport.id
        val newRooms = baseRooms.map { room ->
            val lastInventoryRoom = lastInventoryRoomsAsRooms.find { it.id == room.id }
            if (lastInventoryRoom == null) {
                return@map room.copy(
                    newItem = true,
                    details = room.details.map { detail -> detail.copy(newItem = true) }.toTypedArray()
                )
            }
            lastInventoryRoom.details = setNewValuesForLastInventoryDetails(room.details, lastInventoryRoom.details)
            return@map lastInventoryRoom
        }
        rooms.addAll(newRooms)
    }

    private suspend fun tryGetBaseRooms(propertyId: String) : Array<Room> {
        val newRooms = roomApiCaller.getAllRoomsWithFurniture(
            propertyId,
            { _inventoryErrors.value = _inventoryErrors.value.copy(errorRoomName = it) }
        )
        return newRooms
    }

    fun loadInventory(propertyId: String) {
        _inventoryErrors.value = InventoryApiErrors()
        if (_newValueSetByCompletedInventory) {
            _newValueSetByCompletedInventory = false
            return
        }
        viewModelScope.launch {
            _loadingMutex.withLock {
                rooms.clear()
                _internalIsLoading.value = true
                val newRooms = try {
                    tryGetBaseRooms(propertyId)
                } catch (e: Exception) {
                    println("Error during get base rooms ${e.message}")
                    _inventoryErrors.value = _inventoryErrors.value.copy(getAllRooms = true, getLastInventoryReport = true)
                    e.printStackTrace()
                    return@launch
                }
                try {
                    tryLoadLastInventory(propertyId, newRooms)
                } catch(e : Exception) {
                    rooms.addAll(newRooms)
                    e.printStackTrace()
                }
                _internalIsLoading.value = false
            }
        }
    }

    fun onClick(setIsLoading : (Boolean) -> Unit, propertyId : String, currentLeaseId : String) {
        viewModelScope.launch {
            _loadingMutex.withLock {
                try {
                    setIsLoading(true)
                    setIsLoading(false)
                    navController.navigate("inventory/${propertyId}/${currentLeaseId}")
                } catch (e: Exception) {
                    setIsLoading(false)
                    println("Error occured on the onClick of LoaderInventoryButtonViewModel ${e.message}")
                    e.printStackTrace()
                }
            }
        }
    }

    fun getRooms() : Array<Room> {
        if (_loadingMutex.isLocked) {
            return arrayOf()
        }
        return this.rooms.toTypedArray()
    }

    fun setNewValueSetByCompletedInventory(newRooms : Array<Room>, reportId : String, context: Context) {
        viewModelScope.launch {
            _loadingMutex.withLock {
                rooms.clear()
                _oldReportId.value = reportId
                rooms.addAll(newRooms.map {
                    it.resetAfterInventory(context)
                })
                _newValueSetByCompletedInventory = true
            }
        }
    }
}