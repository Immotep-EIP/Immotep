package fr.keyz.components.inventory

import androidx.compose.runtime.Composable
import androidx.compose.ui.res.stringResource
import fr.keyz.R
import fr.keyz.layouts.BigModalLayout
import fr.keyz.ui.components.StyledButton

@Composable
fun EditRoomOrDetailModal(
    currentRoomOrDetailId : String?,
    //editRoomOrDetail: suspend (name : String, roomType : RoomType?) -> Unit,
    deleteRoomOrDetail: (id : String) -> Unit,
    //editRoomType : Boolean = false,
    close: () -> Unit,
    //isRoom : Boolean
) {
    /*
    var editModalOpen by rememberSaveable { mutableStateOf(false) }

    AddRoomOrDetailModal(
        open = editModalOpen,
        addRoomOrDetail = editRoomOrDetail,
        close = { editModalOpen = false },
        isRoom = isRoom,
        addRoomType = editRoomType
    )

     */
    BigModalLayout(
        height = 0.25f,
        open = currentRoomOrDetailId != null,
        close = close,
        testTag = "editRoomOrDetailModal"
    ) {
        /*
        StyledButton(
            onClick = {
                close()
                //editModalOpen = true
            },
            text = stringResource(R.string.mod_room)
        )
         */
        StyledButton(
            onClick = {
                close()
                deleteRoomOrDetail(currentRoomOrDetailId!!)
            },
            text = stringResource(R.string.delete_room),
            error = true
        )
    }
}