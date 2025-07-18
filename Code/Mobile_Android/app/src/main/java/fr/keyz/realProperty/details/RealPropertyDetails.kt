package fr.keyz.realProperty.details

import androidx.activity.compose.BackHandler
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.isSystemInDarkTheme
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.DropdownMenu
import androidx.compose.material3.DropdownMenuItem
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.EditNote
import androidx.compose.material.icons.outlined.MoreVert
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.IconButtonDefaults
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.State
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import coil.compose.AsyncImage
import fr.keyz.LocalApiService
import fr.keyz.LocalIsOwner
import fr.keyz.R
import fr.keyz.components.addOrEditPropertyModal.AddOrEditPropertyModal
import fr.keyz.apiCallerServices.DetailedProperty
import fr.keyz.apiCallerServices.PropertyStatus
import fr.keyz.components.DeletePopUp
import fr.keyz.components.DeletePopUpConfirmation
import fr.keyz.components.ErrorAlert
import fr.keyz.components.InitialFadeIn
import fr.keyz.components.InternalLoading
import fr.keyz.inventory.loaderButton.LoaderInventoryViewModel
import fr.keyz.components.inviteTenantModal.InviteTenantModal
import fr.keyz.layouts.TabsLayout
import fr.keyz.realProperty.PropertyStatusBox
import fr.keyz.realProperty.details.tabs.AboutPropertyTab
import fr.keyz.realProperty.details.tabs.Damages
import fr.keyz.realProperty.details.tabs.DocumentBox
import fr.keyz.ui.components.BackButton
import fr.keyz.utils.ThemeUtils

@Composable
fun RealPropertyDropDownMenuItem(
    name : String,
    onClick : (() -> Unit)?,
    disabled : Boolean = false,
    color : Color = MaterialTheme.colorScheme.onBackground,
    closeDropDown : () -> Unit,
    testTag : String
) {
    val endColor = if (disabled) color.copy(alpha = 0.4f) else color
    DropdownMenuItem(
        onClick = if (disabled || onClick == null) {
            closeDropDown
        } else {
            { onClick(); closeDropDown() }
        },
        text = { Text(name, color = endColor) },
        modifier = Modifier.testTag(testTag)
    )
}

@Composable
fun RealPropertyImageWithTopButtonsAndDropdown(
    getBack : ((DetailedProperty) -> Unit)?,
    property : State<DetailedProperty>,
    openAddTenant:  (() -> Unit)?,
    endLease : (() -> Unit)?,
    cancelInvitation : (() -> Unit)?,
    openEdit : () -> Unit,
    openDelete : () -> Unit,
    isOwner : Boolean
) {
    var expanded by rememberSaveable { mutableStateOf(false) }

    Box(modifier = Modifier.fillMaxWidth()) {
        if (property.value.picture == null) {
            AsyncImage(
                model = null,
                placeholder = painterResource(id = ThemeUtils.getIcon(isSystemInDarkTheme())),
                error = painterResource(id = ThemeUtils.getIcon(isSystemInDarkTheme())),
                contentDescription = "picture of the ${property.value.name} property",
                modifier = Modifier
                    .fillMaxWidth()
                    .height(150.dp)
                    .clip(
                        RoundedCornerShape(50.dp)
                    )
            )
        } else {
            Image(
                property.value.picture!!,
                contentDescription = "picture of the ${property.value.name} property",
                modifier = Modifier
                    .fillMaxWidth()
                    .height(150.dp)
                    .clip(
                        RoundedCornerShape(50.dp)
                    )
            )
        }
        if (isOwner) {
            Box(
                modifier = Modifier
                    .align(Alignment.TopStart)
                    .padding(top = 5.dp, start = 5.dp)
            ) {
                Row(
                    modifier = Modifier.fillMaxWidth(),
                    horizontalArrangement = Arrangement.SpaceBetween,
                    verticalAlignment = Alignment.CenterVertically
                ) {
                    if (getBack != null) {
                        BackButton { getBack(property.value) }
                    }
                    Box {
                        IconButton(
                            onClick = { expanded = true },
                            colors = IconButtonDefaults.iconButtonColors(containerColor = MaterialTheme.colorScheme.background),
                            modifier = Modifier.testTag("moreVertOptions"),
                        ) {
                            Icon(
                                Icons.Outlined.MoreVert,
                                contentDescription = "More options",
                                tint = MaterialTheme.colorScheme.onBackground
                            )
                        }
                        DropdownMenu(
                            expanded = expanded,
                            onDismissRequest = { expanded = false }
                        ) {
                            RealPropertyDropDownMenuItem(
                                name = stringResource(R.string.add_tenant),
                                onClick = openAddTenant,
                                disabled = openAddTenant == null,
                                closeDropDown = { expanded = false },
                                testTag = "inviteTenantBtn"
                            )
                            RealPropertyDropDownMenuItem(
                                name = stringResource(R.string.end_lease),
                                onClick = endLease,
                                disabled = endLease == null,
                                color = MaterialTheme.colorScheme.error,
                                closeDropDown = { expanded = false },
                                testTag = "endLeaseBtn"
                            )
                            RealPropertyDropDownMenuItem(
                                name = stringResource(R.string.cancel_invitation),
                                onClick = cancelInvitation,
                                disabled = cancelInvitation == null,
                                closeDropDown = {
                                    expanded = false
                                },
                                testTag = "cancelInvitationBtn"
                            )
                            RealPropertyDropDownMenuItem(
                                name = stringResource(R.string.mod_property),
                                onClick = openEdit,
                                closeDropDown = { expanded = false },
                                testTag = "editPropertyBtn"
                            )
                            RealPropertyDropDownMenuItem(
                                name = stringResource(R.string.delete_property),
                                onClick = openDelete,
                                color = MaterialTheme.colorScheme.error,
                                closeDropDown = { expanded = false },
                                testTag = "deletePropertyBtn"
                            )
                        }
                    }
                }
            }
        }
    }
    Spacer(modifier = Modifier.height(10.dp))
}

@Composable
fun RealPropertyDetailsScreen(
    navController: NavController,
    newProperty : DetailedProperty,
    getBack: ((DetailedProperty) -> Unit)? = null,
    loaderInventoryViewModel: LoaderInventoryViewModel,
    deleteProperty : (String) -> Unit
) {
    val isOwner = LocalIsOwner.current.value
    val apiService = LocalApiService.current
    val context = navController.context
    val tabs = listOf(
        stringResource(R.string.about),
        stringResource(R.string.documents),
        stringResource(R.string.damages)
    )

    val viewModel: RealPropertyDetailsViewModel = viewModel {
        RealPropertyDetailsViewModel(navController, apiService)
    }
    val property = viewModel.property.collectAsState()
    var editOpen by rememberSaveable { mutableStateOf(false) }
    var tabIndex by rememberSaveable { mutableIntStateOf(0) }
    var deleteOpen by rememberSaveable { mutableStateOf(false) }
    var inviteTenantOpen by rememberSaveable { mutableStateOf(false) }

    val apiErrors = viewModel.apiError.collectAsState()
    val isLoading = viewModel.isLoading.collectAsState()

    val errorAlertVal = when (apiErrors.value) {
        RealPropertyDetailsViewModel.ApiErrors.GET_PROPERTY -> stringResource(R.string.api_error_get_property)
        RealPropertyDetailsViewModel.ApiErrors.UPDATE_PROPERTY -> stringResource(R.string.api_error_edit_property)
        else -> null
    }


    BackHandler {
        if (getBack != null) {
            getBack(property.value)
        }
    }
    LaunchedEffect(newProperty) {
        viewModel.loadProperty(newProperty)
    }

    AddOrEditPropertyModal(
        open = editOpen,
        close = { editOpen = false },
        onSubmit = { viewModel.editProperty(it, newProperty.id) },
        popupName = stringResource(R.string.edit_property),
        baseValue = property.value.toAddPropertyInput(),
        submitButtonText = stringResource(R.string.save),
        submitButtonIcon = { Icon(Icons.Outlined.EditNote, contentDescription = "Edit property") },
        navController = navController,
        onSubmitPicture = { propertyId, picture -> viewModel.onSubmitPicture(picture) }
    )
    InviteTenantModal(
        open = inviteTenantOpen,
        close = { inviteTenantOpen = false },
        navController = navController,
        propertyId = newProperty.id,
        onSubmit = {email, startDate, endDate -> viewModel.onSubmitInviteTenant(email, startDate, endDate) },
        setIsLoading = { viewModel.setIsLoading(it) }
    )

    if (isLoading.value) {
        InternalLoading()
        return
    }
    InitialFadeIn(300) {
        Column(modifier = Modifier.testTag("realPropertyDetailsScreen")) {
            if (deleteOpen) {
                DeletePopUpConfirmation(
                    delete = { deleteProperty(property.value.id) },
                    close = { deleteOpen = false },
                    globalName = stringResource(R.string.property),
                    detailedName = property.value.name
                )
            }
            RealPropertyImageWithTopButtonsAndDropdown(
                getBack,
                property,
                openAddTenant = if (property.value.status == PropertyStatus.available) { {inviteTenantOpen = true} } else null ,
                endLease = if (property.value.status == PropertyStatus.unavailable) { {} } else null,
                cancelInvitation = if (property.value.status == PropertyStatus.invite_sent) {
                    { viewModel.onCancelInviteTenant() }
                } else null,
                openEdit = { editOpen = true },
                openDelete = { deleteOpen = true },
                isOwner = isOwner
            )
            ErrorAlert(null, null, errorAlertVal)
            Column(modifier = Modifier.background(MaterialTheme.colorScheme.background).padding(20.dp)) {
                Box(modifier = Modifier.fillMaxWidth()) {
                    PropertyStatusBox(property.value.status, modifier = Modifier.padding(end = 5.dp).align(Alignment.TopEnd))
                }
                Text(property.value.name, fontSize = 15.sp, fontWeight = FontWeight.Bold, color = MaterialTheme.colorScheme.onPrimaryContainer)
                Text("${property.value.address}, ${property.value.zipCode} ${property.value.city}, ${property.value.country}",
                    fontSize = 15.sp,
                    fontWeight = FontWeight.Light,
                    color = MaterialTheme.colorScheme.onPrimaryContainer
                )
            }
            TabsLayout(tabIndex, tabs, { tabIndex = it }) {
                when (tabIndex) {
                    0 -> AboutPropertyTab(property, navController, { viewModel.setIsLoading(it) }, loaderInventoryViewModel)
                    1 -> DocumentBox(
                        openPdf = { viewModel.openPdf(it, context)},
                        documents = viewModel.documents.toList(),
                        addDocument = { viewModel.addDocument(it, context) }
                    )
                    2 -> Damages(
                        damageList = viewModel.damages.toList(),
                        addDamage = { viewModel.addDamage(it) },
                        navController = navController,
                        propertyId = property.value.id
                    )
                }
            }
        }
    }
}
