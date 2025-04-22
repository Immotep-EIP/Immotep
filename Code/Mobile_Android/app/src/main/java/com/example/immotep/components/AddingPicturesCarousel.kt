package com.example.immotep.components

import android.content.Context
import android.graphics.BitmapFactory
import android.media.Image
import android.net.Uri
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.PickVisualMediaRequest
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
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
import androidx.compose.ui.graphics.ImageBitmap
import androidx.compose.ui.graphics.asImageBitmap
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import coil.compose.AsyncImage
import com.example.immotep.R
import com.example.immotep.ui.components.StyledButton
import com.example.immotep.utils.Base64Utils
import java.util.Base64

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun DeleteOrSeePictureModal(
    currentImage: Pair<Uri, Int>?,
    onClose: () -> Unit,
    removePicture: ((index: Int) -> Unit)?,
    openImageInFullScreen: (uri: Uri) -> Unit
) {
    if (currentImage != null) {

        ModalBottomSheet(
            onDismissRequest = onClose,
            modifier = Modifier
                .testTag("addRoomModal")

        ) {
            Column(
                horizontalAlignment = Alignment.CenterHorizontally,
                modifier = Modifier.padding(start = 10.dp, end = 10.dp, top = 20.dp, bottom = 20.dp)
            ) {
                StyledButton(
                    onClick = {
                        onClose()
                        openImageInFullScreen(currentImage.first)
                    },
                    text = stringResource(R.string.see_picture)
                )
                if (removePicture != null) {
                    StyledButton(
                        onClick = {
                            removePicture(currentImage.second)
                            onClose()
                        },
                        text = stringResource(R.string.delete_picture)
                    )
                }
            }
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun AddingPicturesCarousel(
    uriPictures : List<Uri>? = null,
    addPicture : ((picture : Uri) -> Unit)? = null,
    removePicture : ((index : Int) -> Unit)? = null,
    stringPictures : List<String>? = null,
    maxPictures : Int = 10,
    error : String? = null,
) {
    var chooseOpen by rememberSaveable { mutableStateOf(false) }
    var pictureSelected by rememberSaveable { mutableStateOf<Pair<Uri, Int>?>(null) }
    var pictureFullScreen by rememberSaveable { mutableStateOf<Uri?>(null) }
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
            .testTag("addingPicturesCarousel")

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
                    TakePhotoButton(
                        onImageCaptured = { uri ->
                            addPicture(uri)
                            onClose()
                        },
                        onAfterImageModalIsShow = { onClose() },
                        onError = {},
                        )
                }
            }
        }
        DeleteOrSeePictureModal(
            currentImage = pictureSelected,
            onClose = { pictureSelected = null },
            removePicture = removePicture,
            openImageInFullScreen = { uri -> pictureFullScreen = uri }
        )
        FullscreenZoomableImage(uri = pictureFullScreen, onClose = { pictureFullScreen = null })
        if (uriPictures != null && stringPictures == null) {
            HorizontalUncontainedCarousel(
                state = rememberCarouselState {
                    uriPictures.size + 1
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
                if (index < uriPictures.size) {
                    AsyncImage(
                        modifier = Modifier
                            .fillMaxWidth()
                            .aspectRatio(1f)
                            .padding(top = 10.dp)
                            .clickable (
                                onClick = {
                                    pictureSelected = Pair(uriPictures[index], index)
                                }
                            ),
                        model = uriPictures[index],
                        contentDescription = "Preview of the added picture at index $index"
                    )
                } else if (addPicture != null && uriPictures.size < maxPictures) {
                    Button(
                        onClick = {
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
        } else if (stringPictures != null) {
            HorizontalUncontainedCarousel(
                state = rememberCarouselState {
                    stringPictures.size
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
                val bitmap = Base64Utils.decodeBase64ToImage(stringPictures[index])
                if (bitmap != null) {
                    Image(
                        modifier = Modifier
                            .fillMaxWidth()
                            .aspectRatio(1f)
                            .padding(top = 10.dp),
                        bitmap = bitmap,
                        contentDescription = "Preview of the added picture at index $index"
                    )
                } else {
                    Text(
                        stringResource(R.string.picture_not_supported),
                        modifier = Modifier.padding(top = 10.dp)
                    )
                }
            }
        } else {
            Text(
                stringResource(R.string.no_pictures_added),
                modifier = Modifier.padding(top = 10.dp)
            )
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