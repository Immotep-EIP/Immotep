package com.example.immotep.addOrEditPropertyModal

import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.PickVisualMediaRequest
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.aspectRatio
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.wrapContentHeight
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
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.Text
import androidx.compose.material3.carousel.HorizontalUncontainedCarousel
import androidx.compose.material3.carousel.rememberCarouselState
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
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
import com.example.immotep.R
import com.example.immotep.apiClient.AddPropertyInput
import com.example.immotep.realProperty.Property
import com.example.immotep.ui.components.OutlinedTextField

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun AddPropertyModal(
    open : Boolean, close : () -> Unit,
    onSubmit : suspend (property : AddPropertyInput) -> Unit
) {
    val viewModel: AddPropertyViewModelViewModel = viewModel()
    val sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = true)
    val surfaceColor = MaterialTheme.colors.onBackground
    val form = viewModel.propertyForm.collectAsState()
    val photoPickerLauncher = rememberLauncherForActivityResult(
        contract = ActivityResultContracts.PickVisualMedia(),
        onResult = { uri ->
            if (uri != null) {
                viewModel.addPicture(uri)
            }
        }
    )
    val onClose : () -> Unit = { viewModel.reset(); close() }
    if (open) {
        ModalBottomSheet (
            onDismissRequest = onClose,
            sheetState = sheetState,
            modifier = Modifier
                .fillMaxWidth()
                .fillMaxHeight(1f)
                .testTag("addPropertyModal")

        ) {
            Column(
                modifier = Modifier
                    .fillMaxWidth()
                    .fillMaxHeight()
                    .verticalScroll(rememberScrollState())
                    .weight(weight = 1f, fill = false)
                    .background(color = MaterialTheme.colors.background),
                horizontalAlignment = Alignment.CenterHorizontally
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
                    Text(stringResource(R.string.create_new_property), fontWeight = FontWeight.Bold, fontSize = 20.sp)
                    IconButton(onClick = onClose) {
                        Icon(Icons.Filled.Close, contentDescription = "Close")
                    }
                }
                Column(
                    modifier = Modifier
                        .fillMaxWidth()
                        .fillMaxHeight()
                        .padding(10.dp),
                    horizontalAlignment = Alignment.CenterHorizontally
                ) {
                    Text(stringResource(R.string.fill_property_infos))
                    OutlinedTextField(
                        value = form.value.name,
                        onValueChange = { value -> viewModel.setName(value)},
                        label = "${stringResource(R.string.name)}*",
                        modifier = Modifier
                            .fillMaxWidth()
                            .padding(top = 10.dp)
                    )
                    OutlinedTextField(
                        value = form.value.address,
                        onValueChange = { value -> viewModel.setAddress(value)},
                        label = "${stringResource(R.string.address)}*",
                        modifier = Modifier
                            .fillMaxWidth()
                            .padding(top = 10.dp)
                    )
                    OutlinedTextField(
                        value = form.value.city,
                        onValueChange = { value -> viewModel.setCity(value)},
                        label = "${stringResource(R.string.city)}*",
                        modifier = Modifier
                            .fillMaxWidth()
                            .padding(top = 10.dp)
                    )
                    OutlinedTextField(
                        value = form.value.postal_code,
                        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number),
                        onValueChange =
                        {
                            value -> viewModel.setZipCode(value)
                        },
                        label = "${stringResource(R.string.zip_code)}*",
                        modifier = Modifier
                            .fillMaxWidth()
                            .padding(top = 10.dp)
                    )
                    OutlinedTextField(
                        value = form.value.country,
                        onValueChange = { value -> viewModel.setCountry(value) },
                        label = "${stringResource(R.string.country)}*",
                        modifier = Modifier
                            .fillMaxWidth()
                            .padding(top = 10.dp)
                    )
                    OutlinedTextField(
                        value = form.value.area_sqm.toString(),
                        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number),
                        onValueChange =
                        {
                            value ->
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
                    )
                    OutlinedTextField(
                        value = form.value.rental_price_per_month.toString(),
                        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number),
                        onValueChange =
                        {
                            value ->
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
                    )
                    OutlinedTextField(
                        value = form.value.deposit_price.toString(),
                        keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Number),
                        onValueChange =
                        {
                            value -> run {
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
                    HorizontalUncontainedCarousel(
                        state = rememberCarouselState {
                            viewModel.pictures.size
                        },
                        itemWidth = 150.dp,
                        itemSpacing = 12.dp,
                        contentPadding = PaddingValues(start = 12.dp),
                        modifier = Modifier
                            .fillMaxWidth()
                            .wrapContentHeight()
                            .padding(top = 12.dp, bottom = 12.dp)
                    )
                        { index ->
                            AsyncImage(
                                modifier = Modifier
                                    .fillMaxWidth()
                                    .aspectRatio(1f)
                                    .padding(top = 10.dp),
                                model = viewModel.pictures[index],
                                contentDescription = "Preview of the added picture at index $index"
                            )
                        }
                    Button(
                        onClick = {  viewModel.onSubmit(onClose, onSubmit) },
                        colors = ButtonDefaults.buttonColors(containerColor = androidx.compose.material3.MaterialTheme.colorScheme.tertiary),
                        modifier = Modifier
                            .fillMaxWidth()
                            .padding(top = 10.dp)

                            .clip(RectangleShape)
                    ) {
                        Icon(Icons.Outlined.Add, contentDescription = "Add property")
                        Text(stringResource(R.string.add_prop))
                    }

                }

            }
        }
    }
}
