package fr.keyz.inventory.roomDetails.OneDetail


import androidx.activity.compose.BackHandler
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.Button
import androidx.compose.material.ButtonDefaults
import androidx.compose.material3.MaterialTheme
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import fr.keyz.LocalApiService
import fr.keyz.R
import fr.keyz.components.AddingPicturesCarousel
import fr.keyz.components.ErrorAlert
import fr.keyz.components.InitialFadeIn
import fr.keyz.components.LoadingDialog
import fr.keyz.inventory.Cleanliness
import fr.keyz.inventory.RoomDetail
import fr.keyz.inventory.State
import fr.keyz.layouts.InventoryLayout
import fr.keyz.ui.components.DropDown
import fr.keyz.ui.components.DropDownItem
import fr.keyz.ui.components.OutlinedTextField

@Composable
fun OneDetailScreen(
    onModifyDetail : (detail : RoomDetail) -> Unit,
    baseDetail : RoomDetail,
    oldReportId : String?,
    navController : NavController,
    propertyId : String,
    leaseId : String,
    isRoom : Boolean = false,
) {
    val apiService = LocalApiService.current
    val viewModel : OneDetailViewModel = viewModel {
        OneDetailViewModel(apiService, navController)
    }
    val detailValue = viewModel.detail.collectAsState()
    val detailError = viewModel.errors.collectAsState()
    val isLoading = viewModel.aiLoading.collectAsState()
    val callError = viewModel.aiCallError.collectAsState()
    val isExit = !detailValue.value.newItem && oldReportId != null

    BackHandler {
        viewModel.onClose(onModifyDetail, isExit)
    }

    LaunchedEffect(Unit) {
        viewModel.reset(baseDetail)
    }

    InventoryLayout(
        testTag = "oneDetailScreen",
        { viewModel.onClose(onModifyDetail, isExit) }
    ) {
        InitialFadeIn {
            LoadingDialog(isOpen = isLoading.value)
            Column(modifier = Modifier.verticalScroll(rememberScrollState())) {
                ErrorAlert(null, null, if (callError.value) stringResource(R.string.ai_call_error) else null)
                if (isExit) {
                    Text(stringResource(R.string.entry_pictures))
                    AddingPicturesCarousel(stringPictures = viewModel.entryPictures)
                }
                Text(if (isExit) stringResource(R.string.exit_pictures) else stringResource(R.string.pictures))
                AddingPicturesCarousel(
                    uriPictures = viewModel.picture,
                    addPicture = { uri -> viewModel.addPicture(uri) },
                    removePicture = { index -> viewModel.removePicture(index) },
                    error = if (detailError.value.picture) stringResource(R.string.add_picture_error) else null,
                )
                OutlinedTextField(
                    value = detailValue.value.comment,
                    onValueChange = { newVal -> viewModel.setComment(newVal) },
                    label = stringResource(R.string.comment),
                    minLines = 4,
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 10.dp)
                        .testTag("oneDetailComment"),
                    errorMessage = if (detailError.value.comment) stringResource(R.string.comment_error) else null
                )
                Text(
                    stringResource(R.string.status), modifier = Modifier.padding(top = 10.dp),
                    color = MaterialTheme.colorScheme.onPrimaryContainer
                )
                DropDown(
                    items = listOf(
                        DropDownItem(stringResource(R.string.new_state), State.new),
                        DropDownItem(stringResource(R.string.good_state), State.good),
                        DropDownItem(stringResource(R.string.medium_state), State.medium),
                        DropDownItem(stringResource(R.string.bad_state), State.bad),
                        DropDownItem(stringResource(R.string.needs_repair), State.needsRepair),
                        DropDownItem(stringResource(R.string.broken_state), State.broken)
                    ),
                    selectedItem = detailValue.value.status,
                    onItemSelected = { newVal -> viewModel.setStatus(newVal) },
                    error = if (detailError.value.status) stringResource(R.string.status_error) else null,
                    testTag = "dropDownState"
                )
                Text(
                    stringResource(R.string.cleaniness), modifier = Modifier.padding(top = 10.dp),
                    color = MaterialTheme.colorScheme.onPrimaryContainer
                )
                DropDown(
                    items = listOf(
                        DropDownItem(stringResource(R.string.clean), Cleanliness.clean),
                        DropDownItem(stringResource(R.string.ok_state), Cleanliness.medium),
                        DropDownItem(stringResource(R.string.dirty), Cleanliness.dirty),
                    ),
                    selectedItem = detailValue.value.cleanliness,
                    onItemSelected = { newVal -> viewModel.setCleanliness(newVal) },
                    error = if (detailError.value.cleanliness) stringResource(R.string.cleaniness_error) else null,
                    testTag = "dropDownCleanliness"
                )

                Column(modifier = Modifier.fillMaxWidth(), horizontalAlignment = Alignment.CenterHorizontally) {
                    Button(
                        shape = RoundedCornerShape(5.dp),
                        modifier = Modifier.padding(top = 10.dp).testTag("aiCallButton"),
                        colors = ButtonDefaults.buttonColors(
                            backgroundColor = MaterialTheme.colorScheme.secondary,
                            contentColor = MaterialTheme.colorScheme.onSecondary
                        ),
                        onClick = { viewModel.summarizeOrCompare(
                            oldReportId = oldReportId,
                            propertyId = propertyId,
                            leaseId = leaseId,
                            isRoom = isRoom,
                            isExit = isExit
                        ) },
                    ) {
                        Text(
                            stringResource(if (isExit) R.string.compare_images else R.string.analyze_pictures),
                            color = MaterialTheme.colorScheme.onPrimaryContainer
                        )
                    }
                    Button(
                        shape = RoundedCornerShape(5.dp),
                        modifier = Modifier.padding(top = 10.dp).testTag("validateButton"),
                        colors = ButtonDefaults.buttonColors(
                            backgroundColor = MaterialTheme.colorScheme.secondary,
                            contentColor = MaterialTheme.colorScheme.onSecondary
                        ),
                        onClick = { viewModel.onConfirm(onModifyDetail, isExit) },
                    ) {
                        Text(
                            stringResource(if (isRoom) R.string.validate_room else  R.string.validate_detail),
                            color = MaterialTheme.colorScheme.onPrimaryContainer
                        )
                    }
                }
            }
        }
    }
}
