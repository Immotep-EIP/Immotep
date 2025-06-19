package fr.keyz.components.addOrEditPropertyModal

import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.PickVisualMediaRequest
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.aspectRatio
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.IconButton
import androidx.compose.material.MaterialTheme
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Close
import androidx.compose.material.icons.outlined.Add
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.RectangleShape
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import coil.compose.AsyncImage
import fr.keyz.LocalApiService
import fr.keyz.R
import fr.keyz.apiCallerServices.AddPropertyInput
import fr.keyz.layouts.BigModalLayout
import fr.keyz.ui.components.OutlinedTextField

@Composable
fun AddOrEditPropertyModal(
    open : Boolean, close : () -> Unit,
    onSubmit : suspend (property : AddPropertyInput) -> String,
    onSubmitPicture : (String) -> Unit,
    popupName : String,
    submitButtonText : String,
    submitButtonIcon : @Composable () -> Unit,
    baseValue : AddPropertyInput? = null,
    navController: NavController
) {
    val apiService = LocalApiService.current
    val viewModel = viewModel {
        AddOrEditPropertyViewModel(
            apiService,
            navController
        )
    }
    val surfaceColor = MaterialTheme.colors.onBackground
    val form = viewModel.propertyForm.collectAsState()
    val picture = viewModel.picture.collectAsState()
    val photoPickerLauncher = rememberLauncherForActivityResult(
        contract = ActivityResultContracts.PickVisualMedia(),
        onResult = { uri ->
            if (uri != null) {
                viewModel.setPicture(uri)
            }
        },
    )
    val onClose: () -> Unit = { viewModel.reset(baseValue); close() }

    LaunchedEffect(baseValue) {
        if (baseValue != null) {
            viewModel.setBaseValue(baseValue)
        }
    }
    BigModalLayout(open = open, close = onClose, height = 0.95f, testTag = "addOrEditPropertyModal") {
        Column(
            modifier = Modifier
                .fillMaxWidth()
                .fillMaxHeight(0.95f)
                .verticalScroll(rememberScrollState())
                .testTag("addOrEditScrollContainer")
        ) {
            Row(
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(3.dp)
                    .drawBehind {
                        val y = size.height - 2.dp.toPx() / 2
                        drawLine(
                            surfaceColor,
                            Offset(0f, y),
                            Offset(size.width, y),
                            2.dp.toPx()
                        )
                    },
                horizontalArrangement = Arrangement.SpaceBetween,
                verticalAlignment = Alignment.CenterVertically
            ) {
                Text(popupName, fontWeight = FontWeight.Bold, fontSize = 20.sp)
                IconButton(onClick = onClose) {
                    Icon(Icons.Filled.Close, contentDescription = "Close")
                }
            }

            Column(
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(10.dp),
                horizontalAlignment = Alignment.CenterHorizontally
            ) {
                Text(stringResource(R.string.fill_property_infos))
                OutlinedTextField(
                    value = form.value.name,
                    onValueChange = { value -> viewModel.setName(value) },
                    label = "${stringResource(R.string.name)}*",
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 10.dp)
                        .testTag("addOrEditName")
                )
                OutlinedTextField(
                    value = form.value.address,
                    onValueChange = { value -> viewModel.setAddress(value) },
                    label = "${stringResource(R.string.address)}*",
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 10.dp)
                        .testTag("addOrEditAddress")

                )
                OutlinedTextField(
                    value = form.value.apartment_number,
                    onValueChange = { value -> viewModel.setAppartementNumber(value) },
                    label = "${stringResource(R.string.appartment_number)}*",
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 10.dp)
                        .testTag("addOrEditNumber")
                )
                OutlinedTextField(
                    value = form.value.city,
                    onValueChange = { value -> viewModel.setCity(value) },
                    label = "${stringResource(R.string.city)}*",
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 10.dp)
                        .testTag("addOrEditCity")
                )
                OutlinedTextField(
                    value = form.value.postal_code,
                    keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number),
                    onValueChange =
                    { value ->
                        viewModel.setZipCode(value)
                    },
                    label = "${stringResource(R.string.zip_code)}*",
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 10.dp)
                        .testTag("addOrEditPostalCode")
                )
                OutlinedTextField(
                    value = form.value.country,
                    onValueChange = { value -> viewModel.setCountry(value) },
                    label = "${stringResource(R.string.country)}*",
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 10.dp)
                        .testTag("addOrEditCountry")
                )
                OutlinedTextField(
                    value = form.value.area_sqm.toString(),
                    keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number),
                    onValueChange =
                    { value ->
                        run {
                            if (value.isEmpty()) {
                                viewModel.setArea(0.0)
                                return@run
                            }
                            val area = value.toDoubleOrNull() ?: return@run
                            viewModel.setArea(area)
                        }
                    },
                    label = "${stringResource(R.string.area)}*",
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 10.dp)
                        .testTag("addOrEditArea")
                )
                OutlinedTextField(
                    value = form.value.rental_price_per_month.toString(),
                    keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number),
                    onValueChange =
                    { value ->
                        run {
                            if (value.isEmpty()) {
                                viewModel.setRental(0)
                                return@run
                            }
                            val rental = value.toIntOrNull() ?: return@run
                            viewModel.setRental(rental)
                        }
                    },
                    label = "${stringResource(R.string.rental)}*",
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 10.dp)
                        .testTag("addOrEditRental")
                )
                OutlinedTextField(
                    value = form.value.deposit_price.toString(),
                    keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number),
                    onValueChange =
                    { value ->
                        run {
                            if (value.isEmpty()) {
                                viewModel.setDeposit(0)
                                return@run
                            }
                            val deposit = value.toIntOrNull() ?: return@run
                            viewModel.setDeposit(deposit)
                        }
                    },
                    label = "${stringResource(R.string.deposit)}*",
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 10.dp)
                        .testTag("addOrEditDeposit")
                )


                Button(
                    onClick =
                    {
                        photoPickerLauncher.launch(
                            PickVisualMediaRequest(ActivityResultContracts.PickVisualMedia.ImageAndVideo)
                        )
                    },
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 10.dp)
                        .clip(RectangleShape)
                ) {
                    Icon(Icons.Outlined.Add, contentDescription = "Add picture")
                    Text(stringResource(R.string.add_picture))
                }
                if (picture.value != null) {
                    AsyncImage(
                        modifier = Modifier
                            .fillMaxWidth()
                            .aspectRatio(1f)
                            .padding(top = 10.dp),
                        model = picture.value,
                        contentDescription = "Preview of the added picture"
                    )
                }
                Button(
                    onClick = { viewModel.onSubmit(onClose, onSubmit, onSubmitPicture, navController.context) },
                    colors = ButtonDefaults.buttonColors(containerColor = androidx.compose.material3.MaterialTheme.colorScheme.secondary),
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(top = 10.dp)
                        .clip(RectangleShape)
                        .testTag("addOrEditSubmit")
                ) {
                    submitButtonIcon()
                    Text(submitButtonText)
                }

            }
        }
    }
}
