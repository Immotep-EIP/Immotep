package com.example.immotep.inventory.roomDetails.EndRoomDetails

import androidx.compose.runtime.Composable
import androidx.navigation.NavController
import com.example.immotep.inventory.Room
import com.example.immotep.inventory.RoomDetail
import com.example.immotep.inventory.roomDetails.OneDetail.OneDetailScreen

@Composable
fun EndRoomDetailsScreen(
    room : Room,
    closeRoomPanel : (room: Room) -> Unit,
    oldReportId : String?,
    propertyId : String,
    newDetails : Array<RoomDetail>,
    navController : NavController,
    isOpen : Boolean,
    setOpen : (Boolean) -> Unit
) {
    if (isOpen) {
        OneDetailScreen(
            onModifyDetail = { detail ->
                setOpen(false)
                if (detail.completed) {
                    val tmpNewRoom = detail.toRoom()
                    tmpNewRoom.details = newDetails
                    closeRoomPanel(tmpNewRoom)
                }
            },
            baseDetail = room.toRoomDetail(),
            oldReportId = oldReportId,
            navController = navController,
            propertyId = propertyId,
            isRoom = true
        )
    }
}