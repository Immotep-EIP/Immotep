package fr.keyz.components

import android.net.Uri
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.gestures.detectTransformGestures
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.runtime.Composable
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableFloatStateOf
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.input.pointer.pointerInput
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalConfiguration
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.unit.dp
import coil.compose.AsyncImage
import fr.keyz.layouts.BigModalLayout

@Composable
fun FullscreenZoomableImage(uri: Uri?, onClose: () -> Unit) {
    var scale by remember { mutableFloatStateOf(1f) }
    var offset by remember { mutableStateOf(Offset.Zero) }
    var rotationState by remember { mutableFloatStateOf(0f) }
    val height = 0.85f
    val configuration = LocalConfiguration.current
    val screenHeight = configuration.screenHeightDp.dp
    val innerHeight = screenHeight * height
    val imageModifier = Modifier
        .fillMaxWidth()
        .height(innerHeight)
        .graphicsLayer(
            scaleX = scale,
            scaleY = scale,
            translationX = offset.x,
            translationY = offset.y,
            rotationZ = rotationState
        )
        .pointerInput(Unit) {
            detectTransformGestures { _, pan, zoom, rotation ->
                scale = (scale * zoom).coerceIn(1f, 5f)
                offset += pan
                rotationState += rotation
            }
        }
    BigModalLayout(
        height = height,
        open = uri != null,
        close = {onClose(); scale = 1f; offset = Offset.Zero; rotationState = 0f;  },
        testTag = "fullscreenZoomableImage",
        backgroundColor = Color.Black
    ) {
        Box(
            modifier = Modifier
                .fillMaxWidth()
                .height(innerHeight)
                .background(Color.Black)
                .clickable { onClose() },
            contentAlignment = Alignment.Center
        ) {
            AsyncImage(
                model = uri,
                contentDescription = "Zoomable Image",
                contentScale = ContentScale.Fit,
                modifier = imageModifier
            )
        }
    }
}
