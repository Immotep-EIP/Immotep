package fr.keyz.components

import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import fr.keyz.R
import fr.keyz.utils.Base64Utils

@Composable
fun Base64ImageView(image : String, description : String, modifier: Modifier) {
    val bitmap = try {
        Base64Utils.decodeBase64ToImage(image)
    } catch (e : Exception) {
        null
    }
    if (bitmap != null) {
        Image(
            modifier = modifier,
            bitmap = bitmap,
            contentDescription = description
        )
    } else {
        Text(
            stringResource(R.string.picture_not_supported),
            modifier = Modifier.padding(top = 10.dp)
        )
    }
}