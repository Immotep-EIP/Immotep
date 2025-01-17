package com.example.immotep.components

import android.net.Uri
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.PickVisualMediaRequest
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.aspectRatio
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.wrapContentHeight
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.Add
import androidx.compose.material3.Button
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.Text
import androidx.compose.material3.carousel.HorizontalUncontainedCarousel
import androidx.compose.material3.carousel.rememberCarouselState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import coil.compose.AsyncImage
import com.example.immotep.R

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun AddingPicturesCarousel(pictures : List<Uri>, addPicture : ((picture : Uri) -> Unit)?, maxPictures : Int = 10, error : String? = null) {

    var chooseOpen by rememberSaveable { mutableStateOf(false) }
    val onClose = { chooseOpen = false }

    val photoPickerLauncher = rememberLauncherForActivityResult(
        contract = ActivityResultContracts.PickVisualMedia(),
        onResult = { uri ->
            if (uri != null && addPicture != null) {
                addPicture(uri)
            }
        }
    )
    Row(
        verticalAlignment = Alignment.CenterVertically,
        modifier = Modifier
            .clip(RoundedCornerShape(5.dp))
            .background(color = MaterialTheme.colorScheme.surfaceDim)
            .border(1.dp, if (error == null) MaterialTheme.colorScheme.surfaceDim else MaterialTheme.colorScheme.error, RoundedCornerShape(5.dp))

    ) {
        if (chooseOpen && addPicture != null) {
            ModalBottomSheet(
                onDismissRequest = onClose,
                modifier = Modifier
                    .testTag("addingImagesModal")

            ) {
                Column(horizontalAlignment = Alignment.CenterHorizontally, modifier = Modifier.fillMaxWidth()) {
                    Button(onClick =
                    {
                        photoPickerLauncher.launch(
                            PickVisualMediaRequest(
                                ActivityResultContracts.PickVisualMedia.ImageAndVideo
                            )
                        )
                        onClose()
                    }) { Text(stringResource(R.string.add_picture_from_gallery)) }
                    TakePhotoButton(onImageCaptured = { uri ->
                        addPicture(uri)
                        onClose()
                    },
                        onAfterImageModalIsShow = { onClose() }
                        ) { }
                }
            }
        }

        HorizontalUncontainedCarousel(
            state = rememberCarouselState {
                pictures.size + 1
            },
            itemWidth = 150.dp,
            itemSpacing = 6.dp,
            contentPadding = PaddingValues(start = 6.dp),
            modifier = Modifier
                .fillMaxWidth()
                .wrapContentHeight()
                .padding(top = 12.dp, bottom = 12.dp)
        )
        { index ->
            if (index < pictures.size) {
                AsyncImage(
                    modifier = Modifier
                        .fillMaxWidth()
                        .aspectRatio(1f)
                        .padding(top = 10.dp),
                    model = pictures[index],
                    contentDescription = "Preview of the added picture at index $index"
                )
            } else if (addPicture != null && pictures.size < maxPictures) {
                Button(onClick = {
                    chooseOpen = true
                },
                    colors = androidx.compose.material3.ButtonDefaults.buttonColors(
                        containerColor = MaterialTheme.colorScheme.primaryContainer,
                        contentColor = MaterialTheme.colorScheme.surfaceDim
                    ),
                    modifier = Modifier
                        .fillMaxWidth()
                        .aspectRatio(1f)
                ) {
                    Icon(Icons.Outlined.Add, contentDescription = "Add picture")
                }
            }
        }
    }
    if (error != null) {
        Text(
            error,
            color = MaterialTheme.colorScheme.error,
            modifier = Modifier.padding(top = 10.dp)
        )
    }
}